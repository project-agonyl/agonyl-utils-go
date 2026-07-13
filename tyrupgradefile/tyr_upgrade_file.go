// Package tyrupgradefile reads and writes Tyr Upgrade.tyr files.
package tyrupgradefile

import (
	"encoding/binary"
	"fmt"
	"io"
)

const RecordSize = 0x32
const FieldCount = 25

type Record struct{ Raw [RecordSize]byte }
type Data []Record

func Read(r io.Reader) (Data, error) {
	b, e := io.ReadAll(r)
	if e != nil {
		return nil, e
	}
	if len(b)%RecordSize != 0 {
		return nil, fmt.Errorf("tyr upgrade file: got %d bytes", len(b))
	}
	d := make(Data, len(b)/RecordSize)
	seen := map[uint32]bool{}
	for i := range d {
		copy(d[i].Raw[:], b[i*RecordSize:(i+1)*RecordSize])
		if e := validate(d[i], seen); e != nil {
			return nil, e
		}
	}
	return d, nil
}
func Write(w io.Writer, d Data) error {
	seen := map[uint32]bool{}
	for _, r := range d {
		if e := validate(r, seen); e != nil {
			return e
		}
		if _, e := w.Write(r.Raw[:]); e != nil {
			return e
		}
	}
	return nil
}
func validate(r Record, seen map[uint32]bool) error {
	code, level := r.Field(0), r.Field(1)
	valid := (code >= 1 && code <= 16) || (code >= 101 && code <= 116)
	key := uint32(code)<<16 | uint32(level)
	if !valid || level < 1 || level > 10 || seen[key] {
		return fmt.Errorf("tyr upgrade file: invalid code=%d level=%d", code, level)
	}
	seen[key] = true
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
		return fmt.Errorf("tyr upgrade file: field index %d", i)
	}
	binary.LittleEndian.PutUint16(r.Raw[i*2:i*2+2], v)
	return nil
}
