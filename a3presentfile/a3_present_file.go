// Package a3presentfile reads and writes Zone Server Present and Reload_Present files.
package a3presentfile

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

const (
	RecordSize  = 0x2f
	NameSize    = 0x15
	RewardCount = 3
)

type Record struct{ Raw [RecordSize]byte }
type Data []Record

func Read(r io.Reader) (Data, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	if len(b)%RecordSize != 0 {
		return nil, fmt.Errorf("a3 present file: got %d bytes, record size %d", len(b), RecordSize)
	}
	data := make(Data, len(b)/RecordSize)
	for i := range data {
		copy(data[i].Raw[:], b[i*RecordSize:(i+1)*RecordSize])
	}
	return data, nil
}
func Write(w io.Writer, data Data) error {
	for _, record := range data {
		if _, err := w.Write(record.Raw[:]); err != nil {
			return err
		}
	}
	return nil
}
func (r Record) Name() string {
	b := r.Raw[:NameSize]
	if i := bytes.IndexByte(b, 0); i >= 0 {
		b = b[:i]
	}
	return string(b)
}
func (r Record) Reward(index int) (uint16, uint16, bool) {
	if index < 0 || index >= RewardCount {
		return 0, 0, false
	}
	offset := NameSize + index*4
	return binary.LittleEndian.Uint16(r.Raw[offset : offset+2]), binary.LittleEndian.Uint16(r.Raw[offset+2 : offset+4]), true
}
func (r Record) Money() uint32      { return binary.LittleEndian.Uint32(r.Raw[0x21:0x25]) }
func (r Record) Lore() uint32       { return binary.LittleEndian.Uint32(r.Raw[0x25:0x29]) }
func (r Record) Experience() uint32 { return binary.LittleEndian.Uint32(r.Raw[0x29:0x2d]) }
func (r Record) Offered() uint16    { return binary.LittleEndian.Uint16(r.Raw[0x2d:0x2f]) }
func (r *Record) SetName(v string) error {
	if len(v) > NameSize {
		return fmt.Errorf("a3 present file: name length %d exceeds %d", len(v), NameSize)
	}
	clear(r.Raw[:NameSize])
	copy(r.Raw[:NameSize], v)
	return nil
}
func (r *Record) SetReward(index int, count uint16, itemCode uint16) error {
	if index < 0 || index >= RewardCount {
		return fmt.Errorf("a3 present file: reward index %d", index)
	}
	offset := NameSize + index*4
	binary.LittleEndian.PutUint16(r.Raw[offset:offset+2], count)
	binary.LittleEndian.PutUint16(r.Raw[offset+2:offset+4], itemCode)
	return nil
}
func (r *Record) SetMoney(v uint32)      { binary.LittleEndian.PutUint32(r.Raw[0x21:0x25], v) }
func (r *Record) SetLore(v uint32)       { binary.LittleEndian.PutUint32(r.Raw[0x25:0x29], v) }
func (r *Record) SetExperience(v uint32) { binary.LittleEndian.PutUint32(r.Raw[0x29:0x2d], v) }
func (r *Record) SetOffered(v uint16)    { binary.LittleEndian.PutUint16(r.Raw[0x2d:0x2f], v) }
