// Package tyrstartpointfile reads and writes Tyr StartPoint.tyr files.
package tyrstartpointfile

import (
	"encoding/binary"
	"fmt"
	"io"
)

const RecordSize = 0x0c
const FieldCount = 6

type Record struct{ Raw [RecordSize]byte }
type Data []Record

func Read(r io.Reader) (Data, error) {
	b, e := io.ReadAll(r)
	if e != nil {
		return nil, e
	}
	if len(b)%RecordSize != 0 {
		return nil, fmt.Errorf("tyr start point file: got %d bytes", len(b))
	}
	d := make(Data, len(b)/RecordSize)
	for i := range d {
		copy(d[i].Raw[:], b[i*RecordSize:(i+1)*RecordSize])
		if e := validate(d[i]); e != nil {
			return nil, e
		}
	}
	return d, nil
}
func Write(w io.Writer, d Data) error {
	for _, r := range d {
		if e := validate(r); e != nil {
			return e
		}
		if _, e := w.Write(r.Raw[:]); e != nil {
			return e
		}
	}
	return nil
}
func validate(r Record) error {
	rank, unit, nation, dir := r.Field(0), r.Field(1), r.Field(2), r.Field(5)
	if rank > 2 || unit > 12 || nation >= 2 || dir == 0 || dir > 8 || (rank != 0 && unit == 0) {
		return fmt.Errorf("tyr start point file: invalid row")
	}
	return nil
}
func (r Record) Field(i int) uint16 {
	if i < 0 || i >= FieldCount {
		return 0
	}
	return binary.LittleEndian.Uint16(r.Raw[i*2 : i*2+2])
}
func (r *Record) SetField(i int, v uint16) error {
	if i < 0 || i >= FieldCount {
		return fmt.Errorf("tyr start point file: field index %d", i)
	}
	binary.LittleEndian.PutUint16(r.Raw[i*2:i*2+2], v)
	return nil
}
