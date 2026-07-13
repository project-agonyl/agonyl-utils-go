// Package towertreasurefile reads and writes Tower treasure item files.
package towertreasurefile

import (
	"encoding/binary"
	"fmt"
	"io"
)

const RecordSize = 6
const MaxRecords = 0x3e

type Record struct{ Raw [RecordSize]byte }
type Data struct {
	Records  []Record
	Trailing []byte
}

func Read(r io.Reader) (Data, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return Data{}, err
	}
	count := len(b) / RecordSize
	if count > MaxRecords {
		count = MaxRecords
	}
	data := Data{Records: make([]Record, count), Trailing: append([]byte(nil), b[count*RecordSize:]...)}
	for i := range data.Records {
		copy(data.Records[i].Raw[:], b[i*RecordSize:(i+1)*RecordSize])
	}
	return data, nil
}
func Write(w io.Writer, data Data) error {
	if len(data.Records) > MaxRecords {
		return fmt.Errorf("tower treasure file: count %d exceeds %d", len(data.Records), MaxRecords)
	}
	for _, record := range data.Records {
		if _, err := w.Write(record.Raw[:]); err != nil {
			return err
		}
	}
	_, err := w.Write(data.Trailing)
	return err
}
func (r Record) ItemCode() uint16      { return binary.LittleEndian.Uint16(r.Raw[0:2]) }
func (r Record) Weight() uint32        { return binary.LittleEndian.Uint32(r.Raw[2:6]) }
func (r *Record) SetItemCode(v uint16) { binary.LittleEndian.PutUint16(r.Raw[0:2], v) }
func (r *Record) SetWeight(v uint32)   { binary.LittleEndian.PutUint32(r.Raw[2:6], v) }
