// Package tyrgiftfile reads and writes Tyr TyrGift.dat files.
package tyrgiftfile

import (
	"encoding/binary"
	"fmt"
	"io"
)

const RecordSize = 0x0c
const MaxRecords = 100
const MaxWeight = 2_000_000_000

type Record struct{ Raw [RecordSize]byte }
type Data []Record

func Read(r io.Reader) (Data, error) {
	b, e := io.ReadAll(r)
	if e != nil {
		return nil, e
	}
	if len(b)%RecordSize != 0 || len(b)/RecordSize > MaxRecords {
		return nil, fmt.Errorf("tyr gift file: got %d bytes", len(b))
	}
	d := make(Data, len(b)/RecordSize)
	var total uint64
	for i := range d {
		copy(d[i].Raw[:], b[i*RecordSize:(i+1)*RecordSize])
		total += uint64(d[i].Weight())
		if total > MaxWeight {
			return nil, fmt.Errorf("tyr gift file: cumulative weight %d", total)
		}
	}
	return d, nil
}
func Write(w io.Writer, d Data) error {
	if len(d) > MaxRecords {
		return fmt.Errorf("tyr gift file: too many records")
	}
	var total uint64
	for _, r := range d {
		total += uint64(r.Weight())
		if total > MaxWeight {
			return fmt.Errorf("tyr gift file: cumulative weight %d", total)
		}
		if _, e := w.Write(r.Raw[:]); e != nil {
			return e
		}
	}
	return nil
}
func (r Record) ItemCode() uint32      { return binary.LittleEndian.Uint32(r.Raw[0:4]) }
func (r Record) Count() uint32         { return binary.LittleEndian.Uint32(r.Raw[4:8]) }
func (r Record) Weight() uint32        { return binary.LittleEndian.Uint32(r.Raw[8:12]) }
func (r *Record) SetItemCode(v uint32) { binary.LittleEndian.PutUint32(r.Raw[0:4], v) }
func (r *Record) SetCount(v uint32)    { binary.LittleEndian.PutUint32(r.Raw[4:8], v) }
func (r *Record) SetWeight(v uint32)   { binary.LittleEndian.PutUint32(r.Raw[8:12], v) }
