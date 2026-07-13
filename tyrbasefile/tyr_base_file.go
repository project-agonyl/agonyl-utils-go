// Package tyrbasefile reads and writes Tyr BaseInfo.tyr files.
package tyrbasefile

import (
	"encoding/binary"
	"fmt"
	"io"
)

const RecordSize = 0x0a
const MaxRecords = 16

type Record struct{ Raw [RecordSize]byte }
type Data []Record

func Read(r io.Reader) (Data, error) {
	b, e := io.ReadAll(r)
	if e != nil {
		return nil, e
	}
	if len(b)%RecordSize != 0 || len(b)/RecordSize > MaxRecords {
		return nil, fmt.Errorf("tyr base file: got %d bytes", len(b))
	}
	d := make(Data, len(b)/RecordSize)
	seen := map[uint16]bool{}
	for i := range d {
		copy(d[i].Raw[:], b[i*RecordSize:(i+1)*RecordSize])
		if e := validate(d[i], seen); e != nil {
			return nil, e
		}
	}
	return d, nil
}
func Write(w io.Writer, d Data) error {
	if len(d) > MaxRecords {
		return fmt.Errorf("tyr base file: too many records")
	}
	seen := map[uint16]bool{}
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
func validate(r Record, seen map[uint16]bool) error {
	i, g, n := r.Index(), r.Grade(), r.Nation()
	if i >= MaxRecords || g > 4 || (n >= 2 && n != 0xff) || seen[i] {
		return fmt.Errorf("tyr base file: invalid index=%d grade=%d nation=%d", i, g, n)
	}
	seen[i] = true
	return nil
}
func (r Record) Index() uint16         { return binary.LittleEndian.Uint16(r.Raw[0:2]) }
func (r Record) Grade() uint16         { return binary.LittleEndian.Uint16(r.Raw[2:4]) }
func (r Record) WarPointValue() uint16 { return binary.LittleEndian.Uint16(r.Raw[4:6]) }
func (r Record) MoraleValue() uint16   { return binary.LittleEndian.Uint16(r.Raw[6:8]) }
func (r Record) Nation() uint16        { return binary.LittleEndian.Uint16(r.Raw[8:10]) }
func (r *Record) SetField(i int, v uint16) error {
	if i < 0 || i >= 5 {
		return fmt.Errorf("tyr base file: field index %d", i)
	}
	binary.LittleEndian.PutUint16(r.Raw[i*2:i*2+2], v)
	return nil
}
