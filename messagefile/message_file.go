// Package messagefile reads and writes Zone Server message table files.
package messagefile

import (
	"encoding/binary"
	"fmt"
	"io"
)

const MaxRecords = 400
const MaxTextLength = 0x3c

type Record struct{ Raw []byte }
type Data []Record

func Read(r io.Reader) (Data, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	if len(b) < 2 {
		return nil, fmt.Errorf("message file: missing count")
	}
	count := int(binary.LittleEndian.Uint16(b[:2]))
	if count > MaxRecords {
		return nil, fmt.Errorf("message file: count %d exceeds %d", count, MaxRecords)
	}
	data := make(Data, 0, count)
	offset := 2
	for i := 0; i < count; i++ {
		if offset+6 > len(b) {
			return nil, fmt.Errorf("message file: truncated record %d", i)
		}
		maximum := int(binary.LittleEndian.Uint16(b[offset+2 : offset+4]))
		length := int(binary.LittleEndian.Uint16(b[offset+4 : offset+6]))
		if maximum < length || length > MaxTextLength || offset+6+length > len(b) {
			return nil, fmt.Errorf("message file: invalid record %d", i)
		}
		raw := append([]byte(nil), b[offset:offset+6+length]...)
		data = append(data, Record{Raw: raw})
		offset += len(raw)
	}
	if offset != len(b) {
		return nil, fmt.Errorf("message file: %d trailing bytes", len(b)-offset)
	}
	return data, nil
}

func Write(w io.Writer, data Data) error {
	if len(data) > MaxRecords {
		return fmt.Errorf("message file: count %d exceeds %d", len(data), MaxRecords)
	}
	var count [2]byte
	binary.LittleEndian.PutUint16(count[:], uint16(len(data)))
	if _, err := w.Write(count[:]); err != nil {
		return err
	}
	for index, record := range data {
		if len(record.Raw) < 6 || int(binary.LittleEndian.Uint16(record.Raw[4:6])) != len(record.Raw)-6 {
			return fmt.Errorf("message file: invalid record %d", index)
		}
		if _, err := w.Write(record.Raw); err != nil {
			return err
		}
	}
	return nil
}

func (r Record) Index() uint16         { return binary.LittleEndian.Uint16(r.Raw[0:2]) }
func (r Record) MaximumLength() uint16 { return binary.LittleEndian.Uint16(r.Raw[2:4]) }
func (r Record) Text() string          { return string(r.Raw[6:]) }
func (r *Record) SetIndex(v uint16)    { binary.LittleEndian.PutUint16(r.Raw[0:2], v) }
func (r *Record) SetText(v string) error {
	if len(v) > int(r.MaximumLength()) || len(v) > MaxTextLength {
		return fmt.Errorf("message file: text length %d exceeds maximum", len(v))
	}
	raw := make([]byte, 6+len(v))
	copy(raw[:4], r.Raw[:4])
	binary.LittleEndian.PutUint16(raw[4:6], uint16(len(v)))
	copy(raw[6:], v)
	r.Raw = raw
	return nil
}
