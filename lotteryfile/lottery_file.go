// Package lotteryfile reads and writes Zone Server Event/LotteryItem.dat files.
package lotteryfile

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

const (
	RecordSize   = 0x144
	MessageCount = 5
	MessageSize  = 0x40
)

type Record struct{ Raw [RecordSize]byte }
type Data []Record

func Read(r io.Reader) (Data, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	if len(b)%RecordSize != 0 {
		return nil, fmt.Errorf("lottery file: got %d bytes, record size %d", len(b), RecordSize)
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
func (r Record) EventItemID() byte      { return r.Raw[0] }
func (r Record) Enabled() byte          { return r.Raw[1] }
func (r Record) RewardItemCode() uint16 { return binary.LittleEndian.Uint16(r.Raw[2:4]) }
func (r Record) Message(index int) (string, bool) {
	if index < 0 || index >= MessageCount {
		return "", false
	}
	b := r.Raw[4+index*MessageSize : 4+(index+1)*MessageSize]
	if i := bytes.IndexByte(b, 0); i >= 0 {
		b = b[:i]
	}
	return string(b), true
}
func (r *Record) SetEventItemID(v byte)      { r.Raw[0] = v }
func (r *Record) SetEnabled(v byte)          { r.Raw[1] = v }
func (r *Record) SetRewardItemCode(v uint16) { binary.LittleEndian.PutUint16(r.Raw[2:4], v) }
func (r *Record) SetMessage(index int, v string) error {
	if index < 0 || index >= MessageCount {
		return fmt.Errorf("lottery file: message index %d", index)
	}
	if len(v) > MessageSize {
		return fmt.Errorf("lottery file: message length %d exceeds %d", len(v), MessageSize)
	}
	dst := r.Raw[4+index*MessageSize : 4+(index+1)*MessageSize]
	clear(dst)
	copy(dst, v)
	return nil
}
