// Package tyrskilllayerfile reads and writes Tyr SkillLayer.tyr files.
package tyrskilllayerfile

import (
	"encoding/binary"
	"fmt"
	"io"
)

const RecordSize = 0x18
const FlagCount = 5

type Record struct{ Raw [RecordSize]byte }
type Data []Record

func Read(r io.Reader) (Data, error) {
	b, e := io.ReadAll(r)
	if e != nil {
		return nil, e
	}
	if len(b)%RecordSize != 0 {
		return nil, fmt.Errorf("tyr skill layer file: got %d bytes", len(b))
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
	c, s := r.ClassIndex(), r.SkillIndex()
	if (c != 0xff || s != 0xff) && (c >= 4 || s >= 64) {
		return fmt.Errorf("tyr skill layer file: class=%d skill=%d", c, s)
	}
	return nil
}
func (r Record) ClassIndex() uint16 { return binary.LittleEndian.Uint16(r.Raw[0:2]) }
func (r Record) SkillIndex() uint16 { return binary.LittleEndian.Uint16(r.Raw[2:4]) }
func (r Record) PlayerToTarget(i int) (bool, bool) {
	if i < 0 || i >= FlagCount {
		return false, false
	}
	return int16(binary.LittleEndian.Uint16(r.Raw[4+i*2:6+i*2])) != 0, true
}
func (r Record) TargetToEffect(i int) (bool, bool) {
	if i < 0 || i >= FlagCount {
		return false, false
	}
	return int16(binary.LittleEndian.Uint16(r.Raw[0x0e+i*2:0x10+i*2])) != 0, true
}
func (r *Record) SetClassIndex(v uint16) { binary.LittleEndian.PutUint16(r.Raw[0:2], v) }
func (r *Record) SetSkillIndex(v uint16) { binary.LittleEndian.PutUint16(r.Raw[2:4], v) }
func (r *Record) SetPlayerToTarget(i int, v bool) error {
	if i < 0 || i >= FlagCount {
		return fmt.Errorf("tyr skill layer file: flag index %d", i)
	}
	n := uint16(0)
	if v {
		n = 1
	}
	binary.LittleEndian.PutUint16(r.Raw[4+i*2:6+i*2], n)
	return nil
}
func (r *Record) SetTargetToEffect(i int, v bool) error {
	if i < 0 || i >= FlagCount {
		return fmt.Errorf("tyr skill layer file: flag index %d", i)
	}
	n := uint16(0)
	if v {
		n = 1
	}
	binary.LittleEndian.PutUint16(r.Raw[0x0e+i*2:0x10+i*2], n)
	return nil
}
