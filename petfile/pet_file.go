// Package petfile reads and writes Zone Server item/pet files.
package petfile

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

const (
	RecordSize = 0x30
	MaxRecords = 0x100
	NameSize   = 0x20
	CodeMin    = 1000
	CodeMax    = CodeMin + MaxRecords - 1
)

type Record struct{ Raw [RecordSize]byte }
type Data []Record

func Read(r io.Reader) (Data, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	if len(b)%RecordSize != 0 || len(b)/RecordSize > MaxRecords {
		return nil, fmt.Errorf("pet file: got %d bytes", len(b))
	}
	data := make(Data, len(b)/RecordSize)
	for i := range data {
		copy(data[i].Raw[:], b[i*RecordSize:(i+1)*RecordSize])
		if !data[i].Empty() && (data[i].Code() < CodeMin || data[i].Code() > CodeMax) {
			return nil, fmt.Errorf("pet file: code %d outside %d..%d", data[i].Code(), CodeMin, CodeMax)
		}
	}
	return data, nil
}

func Write(w io.Writer, data Data) error {
	if len(data) > MaxRecords {
		return fmt.Errorf("pet file: got %d records, max %d", len(data), MaxRecords)
	}
	for _, record := range data {
		if !record.Empty() && (record.Code() < CodeMin || record.Code() > CodeMax) {
			return fmt.Errorf("pet file: code %d outside %d..%d", record.Code(), CodeMin, CodeMax)
		}
		if _, err := w.Write(record.Raw[:]); err != nil {
			return err
		}
	}
	return nil
}

func (r Record) Empty() bool        { return bytes.Equal(r.Raw[:], make([]byte, RecordSize)) }
func (r Record) Group() uint16      { return binary.LittleEndian.Uint16(r.Raw[0:2]) }
func (r Record) Code() uint16       { return binary.LittleEndian.Uint16(r.Raw[2:4]) }
func (r Record) Name() string       { return fixedString(r.Raw[4:36]) }
func (r Record) Price() uint32      { return binary.LittleEndian.Uint32(r.Raw[36:40]) }
func (r Record) Auth() uint16       { return binary.LittleEndian.Uint16(r.Raw[40:42]) }
func (r Record) Limit() uint16      { return binary.LittleEndian.Uint16(r.Raw[42:44]) }
func (r Record) Data1() uint16      { return binary.LittleEndian.Uint16(r.Raw[44:46]) }
func (r Record) Data2() uint16      { return binary.LittleEndian.Uint16(r.Raw[46:48]) }
func (r *Record) SetGroup(v uint16) { binary.LittleEndian.PutUint16(r.Raw[0:2], v) }
func (r *Record) SetCode(v uint16) error {
	if v < CodeMin || v > CodeMax {
		return fmt.Errorf("pet file: code %d outside %d..%d", v, CodeMin, CodeMax)
	}
	binary.LittleEndian.PutUint16(r.Raw[2:4], v)
	return nil
}
func (r *Record) SetName(v string) error { return putFixedString(r.Raw[4:36], v) }
func (r *Record) SetPrice(v uint32)      { binary.LittleEndian.PutUint32(r.Raw[36:40], v) }
func (r *Record) SetAuth(v uint16)       { binary.LittleEndian.PutUint16(r.Raw[40:42], v) }
func (r *Record) SetLimit(v uint16)      { binary.LittleEndian.PutUint16(r.Raw[42:44], v) }
func (r *Record) SetData1(v uint16)      { binary.LittleEndian.PutUint16(r.Raw[44:46], v) }
func (r *Record) SetData2(v uint16)      { binary.LittleEndian.PutUint16(r.Raw[46:48], v) }

func fixedString(b []byte) string {
	if i := bytes.IndexByte(b, 0); i >= 0 {
		b = b[:i]
	}
	return string(bytes.TrimRight(b, " "))
}
func putFixedString(dst []byte, v string) error {
	if len(v) > len(dst) {
		return fmt.Errorf("pet file: string length %d exceeds %d", len(v), len(dst))
	}
	clear(dst)
	copy(dst, v)
	return nil
}
