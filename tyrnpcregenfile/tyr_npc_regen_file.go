// Package tyrnpcregenfile reads and writes Tyr NPCRegen.tyr files.
package tyrnpcregenfile

import (
	"encoding/binary"
	"fmt"
	"io"
)

const FullRecordSize = 0x0c
const CompactRecordSize = 0x0a

type Record struct{ Raw [FullRecordSize]byte }
type Data struct {
	RecordSize int
	Records    []Record
}

func Read(r io.Reader) (Data, error) {
	b, e := io.ReadAll(r)
	if e != nil {
		return Data{}, e
	}
	size := FullRecordSize
	if len(b)%size != 0 {
		size = CompactRecordSize
	}
	if len(b)%size != 0 {
		return Data{}, fmt.Errorf("tyr npc regen file: got %d bytes", len(b))
	}
	d := Data{RecordSize: size, Records: make([]Record, len(b)/size)}
	for i := range d.Records {
		copy(d.Records[i].Raw[:], b[i*size:(i+1)*size])
		if e := validate(d.Records[i]); e != nil {
			return Data{}, e
		}
	}
	return d, nil
}
func Write(w io.Writer, d Data) error {
	if d.RecordSize != FullRecordSize && d.RecordSize != CompactRecordSize {
		return fmt.Errorf("tyr npc regen file: invalid record size %d", d.RecordSize)
	}
	for _, r := range d.Records {
		if e := validate(r); e != nil {
			return e
		}
		if _, e := w.Write(r.Raw[:d.RecordSize]); e != nil {
			return e
		}
	}
	return nil
}
func validate(r Record) error {
	if r.BaseIndex() >= 16 || r.Direction() > 8 {
		return fmt.Errorf("tyr npc regen file: invalid base=%d direction=%d", r.BaseIndex(), r.Direction())
	}
	return nil
}
func (r Record) BaseIndex() uint16 { return binary.LittleEndian.Uint16(r.Raw[0:2]) }
func (r Record) NPCType() uint16   { return binary.LittleEndian.Uint16(r.Raw[2:4]) }
func (r Record) X() uint16         { return binary.LittleEndian.Uint16(r.Raw[4:6]) }
func (r Record) Y() uint16         { return binary.LittleEndian.Uint16(r.Raw[6:8]) }
func (r Record) Direction() uint16 { return binary.LittleEndian.Uint16(r.Raw[8:10]) }
func (r Record) Step() byte        { return r.Raw[10] }
func (r *Record) SetField(i int, v uint16) error {
	if i < 0 || i >= 5 {
		return fmt.Errorf("tyr npc regen file: field index %d", i)
	}
	binary.LittleEndian.PutUint16(r.Raw[i*2:i*2+2], v)
	return nil
}
func (r *Record) SetStep(v byte) { r.Raw[10] = v }
