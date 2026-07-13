// Package tyrportalfile reads and writes Tyr WarpPortal.tyr files.
package tyrportalfile

import (
	"encoding/binary"
	"fmt"
	"io"
)

const RecordSize = 8

type Record struct{ Raw [RecordSize]byte }
type Data []Record

func Read(r io.Reader) (Data, error) {
	b, e := io.ReadAll(r)
	if e != nil {
		return nil, e
	}
	if len(b)%RecordSize != 0 {
		return nil, fmt.Errorf("tyr portal file: got %d bytes", len(b))
	}
	d := make(Data, len(b)/RecordSize)
	seen := map[uint32]bool{}
	for i := range d {
		copy(d[i].Raw[:], b[i*RecordSize:(i+1)*RecordSize])
		key := uint32(d[i].SourceY())<<16 | uint32(d[i].SourceX())
		if seen[key] {
			return nil, fmt.Errorf("tyr portal file: duplicate source")
		}
		seen[key] = true
	}
	return d, nil
}
func Write(w io.Writer, d Data) error {
	seen := map[uint32]bool{}
	for _, r := range d {
		key := uint32(r.SourceY())<<16 | uint32(r.SourceX())
		if seen[key] {
			return fmt.Errorf("tyr portal file: duplicate source")
		}
		seen[key] = true
		if _, e := w.Write(r.Raw[:]); e != nil {
			return e
		}
	}
	return nil
}
func (r Record) SourceX() uint16      { return binary.LittleEndian.Uint16(r.Raw[0:2]) }
func (r Record) SourceY() uint16      { return binary.LittleEndian.Uint16(r.Raw[2:4]) }
func (r Record) DestinationX() uint16 { return binary.LittleEndian.Uint16(r.Raw[4:6]) }
func (r Record) DestinationY() uint16 { return binary.LittleEndian.Uint16(r.Raw[6:8]) }
func (r *Record) SetField(i int, v uint16) error {
	if i < 0 || i >= 4 {
		return fmt.Errorf("tyr portal file: field index %d", i)
	}
	binary.LittleEndian.PutUint16(r.Raw[i*2:i*2+2], v)
	return nil
}
