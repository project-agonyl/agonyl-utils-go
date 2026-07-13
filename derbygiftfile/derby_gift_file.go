// Package derbygiftfile reads and writes Zone Server npc/DerbyGift.dat files.
package derbygiftfile

import (
	"encoding/binary"
	"fmt"
	"io"
)

const RecordSize = 0x0c
const MaxProbability = 2_000_000_000

type Record struct{ Raw [RecordSize]byte }
type Data []Record

func Read(r io.Reader) (Data, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	if len(b)%RecordSize != 0 {
		return nil, fmt.Errorf("derby gift file: got %d bytes, record size %d", len(b), RecordSize)
	}
	data := make(Data, len(b)/RecordSize)
	var total uint64
	for i := range data {
		copy(data[i].Raw[:], b[i*RecordSize:(i+1)*RecordSize])
		total += uint64(data[i].Weight())
		if total > MaxProbability {
			return nil, fmt.Errorf("derby gift file: cumulative probability %d exceeds %d", total, MaxProbability)
		}
	}
	return data, nil
}
func Write(w io.Writer, data Data) error {
	var total uint64
	for _, record := range data {
		total += uint64(record.Weight())
		if total > MaxProbability {
			return fmt.Errorf("derby gift file: cumulative probability %d exceeds %d", total, MaxProbability)
		}
		if _, err := w.Write(record.Raw[:]); err != nil {
			return err
		}
	}
	return nil
}
func (r Record) ItemCode() uint32      { return binary.LittleEndian.Uint32(r.Raw[0:4]) }
func (r Record) Quantity() uint32      { return binary.LittleEndian.Uint32(r.Raw[4:8]) }
func (r Record) Weight() uint32        { return binary.LittleEndian.Uint32(r.Raw[8:12]) }
func (r *Record) SetItemCode(v uint32) { binary.LittleEndian.PutUint32(r.Raw[0:4], v) }
func (r *Record) SetQuantity(v uint32) { binary.LittleEndian.PutUint32(r.Raw[4:8], v) }
func (r *Record) SetWeight(v uint32)   { binary.LittleEndian.PutUint32(r.Raw[8:12], v) }
