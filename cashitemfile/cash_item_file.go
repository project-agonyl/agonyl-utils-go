// Package cashitemfile reads and writes Zone Server CashItemTbl.dat files.
package cashitemfile

import (
	"encoding/binary"
	"fmt"
	"io"
)

const RecordSize = 0x0a

type Record struct{ Raw [RecordSize]byte }
type Data []Record

func Read(r io.Reader) (Data, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	if len(b)%RecordSize != 0 {
		return nil, fmt.Errorf("cash item file: got %d bytes, record size %d", len(b), RecordSize)
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

func (r Record) NPCType() uint16       { return binary.LittleEndian.Uint16(r.Raw[0:2]) }
func (r Record) ItemCode() uint16      { return binary.LittleEndian.Uint16(r.Raw[2:4]) }
func (r Record) Price() uint32         { return binary.LittleEndian.Uint32(r.Raw[4:8]) }
func (r Record) Count() uint16         { return binary.LittleEndian.Uint16(r.Raw[8:10]) }
func (r *Record) SetNPCType(v uint16)  { binary.LittleEndian.PutUint16(r.Raw[0:2], v) }
func (r *Record) SetItemCode(v uint16) { binary.LittleEndian.PutUint16(r.Raw[2:4], v) }
func (r *Record) SetPrice(v uint32)    { binary.LittleEndian.PutUint32(r.Raw[4:8], v) }
func (r *Record) SetCount(v uint16)    { binary.LittleEndian.PutUint16(r.Raw[8:10], v) }
