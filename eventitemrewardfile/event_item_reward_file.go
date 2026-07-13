// Package eventitemrewardfile reads and writes binary reward files referenced by EventItemInfo.ini.
package eventitemrewardfile

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

const (
	RecordSize  = 0x44
	MaxRecords  = 30
	MessageSize = 0x40
)

type Record struct{ Raw [RecordSize]byte }
type Data []Record

func Read(r io.Reader) (Data, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	if len(b)%RecordSize != 0 || len(b)/RecordSize > MaxRecords {
		return nil, fmt.Errorf("event item reward file: got %d bytes", len(b))
	}
	data := make(Data, len(b)/RecordSize)
	for i := range data {
		copy(data[i].Raw[:], b[i*RecordSize:(i+1)*RecordSize])
	}
	return data, nil
}
func Write(w io.Writer, data Data) error {
	if len(data) > MaxRecords {
		return fmt.Errorf("event item reward file: got %d records, max %d", len(data), MaxRecords)
	}
	for _, record := range data {
		if _, err := w.Write(record.Raw[:]); err != nil {
			return err
		}
	}
	return nil
}
func (r Record) ItemCode() uint16 { return binary.LittleEndian.Uint16(r.Raw[0:2]) }
func (r Record) Weight() uint16   { return binary.LittleEndian.Uint16(r.Raw[2:4]) }
func (r Record) Message() string {
	b := r.Raw[4:]
	if i := bytes.IndexByte(b, 0); i >= 0 {
		b = b[:i]
	}
	return string(b)
}
func (r *Record) SetItemCode(v uint16) { binary.LittleEndian.PutUint16(r.Raw[0:2], v) }
func (r *Record) SetWeight(v uint16)   { binary.LittleEndian.PutUint16(r.Raw[2:4], v) }
func (r *Record) SetMessage(v string) error {
	if len(v) > MessageSize {
		return fmt.Errorf("event item reward file: message length %d exceeds %d", len(v), MessageSize)
	}
	clear(r.Raw[4:])
	copy(r.Raw[4:], v)
	return nil
}
