// Package questexfile reads and writes one Zone Server QuestEx binary record.
package questexfile

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

const (
	RecordSize           = 0x670
	HeaderSize           = 0x10
	NeedsOffset          = 0x10
	NeedsSize            = 0x1c
	RewardOffset         = 0x2c
	RewardSize           = 0x34
	ContentOffset        = 0x60
	ContentCount         = 7
	ContentStride        = 0xdc
	ContentFixedSize     = 0x5c
	ContentStringSize    = 0x40
	ContentString0Offset = 0xbc
	ContentString1Offset = 0xfc
	NextQuestOffset      = 0x664
	NextQuestCount       = 3
)

type Data struct {
	Raw      [RecordSize]byte
	Disabled bool
}

func Read(r io.Reader) (Data, error) {
	var data Data
	b, err := io.ReadAll(r)
	if err != nil {
		return Data{}, err
	}
	reader := bytes.NewReader(b)
	if _, err = io.ReadFull(reader, data.Raw[:HeaderSize]); err != nil {
		return Data{}, fmt.Errorf("questex file: header: %w", err)
	}
	if data.Int32(0) == -1 {
		data.Disabled = true
		if reader.Len() != 0 {
			return Data{}, fmt.Errorf("questex file: disabled record has trailing bytes")
		}
		return data, nil
	}
	if _, err = io.ReadFull(reader, data.Raw[NeedsOffset:NeedsOffset+NeedsSize]); err != nil {
		return Data{}, fmt.Errorf("questex file: needs: %w", err)
	}
	if _, err = io.ReadFull(reader, data.Raw[RewardOffset:RewardOffset+RewardSize]); err != nil {
		return Data{}, fmt.Errorf("questex file: reward: %w", err)
	}
	for index := 0; index < ContentCount; index++ {
		base := ContentOffset + index*ContentStride
		if _, err = io.ReadFull(reader, data.Raw[base:base+ContentFixedSize]); err != nil {
			return Data{}, fmt.Errorf("questex file: content %d: %w", index, err)
		}
		if err = readString(reader, data.Raw[ContentString0Offset+index*ContentStride:ContentString0Offset+index*ContentStride+ContentStringSize]); err != nil {
			return Data{}, fmt.Errorf("questex file: content %d string 0: %w", index, err)
		}
		if err = readString(reader, data.Raw[ContentString1Offset+index*ContentStride:ContentString1Offset+index*ContentStride+ContentStringSize]); err != nil {
			return Data{}, fmt.Errorf("questex file: content %d string 1: %w", index, err)
		}
	}
	if _, err = io.ReadFull(reader, data.Raw[NextQuestOffset:NextQuestOffset+NextQuestCount*4]); err != nil {
		return Data{}, fmt.Errorf("questex file: next quests: %w", err)
	}
	if reader.Len() != 0 {
		return Data{}, fmt.Errorf("questex file: %d trailing bytes", reader.Len())
	}
	return data, nil
}

func Write(w io.Writer, data Data) error {
	if _, err := w.Write(data.Raw[:HeaderSize]); err != nil {
		return err
	}
	if data.Disabled || data.Int32(0) == -1 {
		return nil
	}
	if _, err := w.Write(data.Raw[NeedsOffset : NeedsOffset+NeedsSize]); err != nil {
		return err
	}
	if _, err := w.Write(data.Raw[RewardOffset : RewardOffset+RewardSize]); err != nil {
		return err
	}
	for index := 0; index < ContentCount; index++ {
		base := ContentOffset + index*ContentStride
		if _, err := w.Write(data.Raw[base : base+ContentFixedSize]); err != nil {
			return err
		}
		if err := writeString(w, data.Raw[ContentString0Offset+index*ContentStride:ContentString0Offset+index*ContentStride+ContentStringSize]); err != nil {
			return err
		}
		if err := writeString(w, data.Raw[ContentString1Offset+index*ContentStride:ContentString1Offset+index*ContentStride+ContentStringSize]); err != nil {
			return err
		}
	}
	_, err := w.Write(data.Raw[NextQuestOffset : NextQuestOffset+NextQuestCount*4])
	return err
}

func (d Data) Int32(offset int) int32 {
	if offset < 0 || offset+4 > len(d.Raw) {
		return 0
	}
	return int32(binary.LittleEndian.Uint32(d.Raw[offset : offset+4]))
}
func (d *Data) SetInt32(offset int, value int32) error {
	if offset < 0 || offset+4 > len(d.Raw) {
		return fmt.Errorf("questex file: offset %d", offset)
	}
	binary.LittleEndian.PutUint32(d.Raw[offset:offset+4], uint32(value))
	d.Disabled = d.Int32(0) == -1
	return nil
}
func (d Data) ContentString(content int, slot int) (string, error) {
	offset, err := stringOffset(content, slot)
	if err != nil {
		return "", err
	}
	b := d.Raw[offset : offset+ContentStringSize]
	if i := bytes.IndexByte(b, 0); i >= 0 {
		b = b[:i]
	}
	return string(b), nil
}
func (d *Data) SetContentString(content int, slot int, value string) error {
	offset, err := stringOffset(content, slot)
	if err != nil {
		return err
	}
	if len(value) >= ContentStringSize {
		return fmt.Errorf("questex file: string length %d exceeds %d", len(value), ContentStringSize-1)
	}
	clear(d.Raw[offset : offset+ContentStringSize])
	copy(d.Raw[offset:offset+ContentStringSize], value)
	return nil
}
func stringOffset(content int, slot int) (int, error) {
	if content < 0 || content >= ContentCount || slot < 0 || slot > 1 {
		return 0, fmt.Errorf("questex file: string %d/%d", content, slot)
	}
	base := ContentString0Offset
	if slot == 1 {
		base = ContentString1Offset
	}
	return base + content*ContentStride, nil
}
func readString(r io.Reader, dst []byte) error {
	var size [2]byte
	if _, err := io.ReadFull(r, size[:]); err != nil {
		return err
	}
	length := int(binary.LittleEndian.Uint16(size[:]))
	if length >= ContentStringSize {
		return fmt.Errorf("length %d exceeds %d", length, ContentStringSize-1)
	}
	if _, err := io.ReadFull(r, dst[:length]); err != nil {
		return err
	}
	dst[length] = 0
	return nil
}
func writeString(w io.Writer, src []byte) error {
	length := bytes.IndexByte(src, 0)
	if length < 0 {
		length = len(src)
	}
	if length >= ContentStringSize {
		return fmt.Errorf("questex file: string length %d", length)
	}
	var size [2]byte
	binary.LittleEndian.PutUint16(size[:], uint16(length))
	if _, err := w.Write(size[:]); err != nil {
		return err
	}
	_, err := w.Write(src[:length])
	return err
}
