// Package shuecombinationfile reads and writes Zone Server ShueCombinationData files.
package shuecombinationfile

import (
	"encoding/binary"
	"fmt"
	"io"
)

const RecordSize = 0x20
const FieldCount = 16

type Record struct{ Raw [RecordSize]byte }
type Data []Record

func Read(r io.Reader) (Data, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	if len(b)%RecordSize != 0 {
		return nil, fmt.Errorf("shue combination file: got %d bytes, record size %d", len(b), RecordSize)
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
func (r Record) Field(index int) (uint16, bool) {
	if index < 0 || index >= FieldCount {
		return 0, false
	}
	return binary.LittleEndian.Uint16(r.Raw[index*2 : index*2+2]), true
}
func (r *Record) SetField(index int, value uint16) error {
	if index < 0 || index >= FieldCount {
		return fmt.Errorf("shue combination file: field index %d", index)
	}
	binary.LittleEndian.PutUint16(r.Raw[index*2:index*2+2], value)
	return nil
}
func (r Record) BaseRune1() uint16    { v, _ := r.Field(0); return v }
func (r Record) BaseRune2() uint16    { v, _ := r.Field(1); return v }
func (r Record) SuccessRatio() uint16 { v, _ := r.Field(10); return v }
func (r Record) SuccessItem() uint16  { v, _ := r.Field(11); return v }
