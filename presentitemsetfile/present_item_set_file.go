// Package presentitemsetfile reads and writes Zone Server PresentItemSet.dat files.
package presentitemsetfile

import (
	"encoding/binary"
	"fmt"
	"io"
)

const RecordSize = 0x0e
const RewardCount = 3

type Record struct{ Raw [RecordSize]byte }
type Data []Record

func Read(r io.Reader) (Data, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	if len(b)%RecordSize != 0 {
		return nil, fmt.Errorf("present item set file: got %d bytes, record size %d", len(b), RecordSize)
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

func (r Record) PresentSetCode() uint16 { return binary.LittleEndian.Uint16(r.Raw[0:2]) }
func (r Record) Reward(index int) (uint16, uint16, bool) {
	if index < 0 || index >= RewardCount {
		return 0, 0, false
	}
	offset := 2 + index*4
	return binary.LittleEndian.Uint16(r.Raw[offset : offset+2]), binary.LittleEndian.Uint16(r.Raw[offset+2 : offset+4]), true
}
func (r *Record) SetPresentSetCode(v uint16) { binary.LittleEndian.PutUint16(r.Raw[0:2], v) }
func (r *Record) SetReward(index int, count uint16, itemCode uint16) error {
	if index < 0 || index >= RewardCount {
		return fmt.Errorf("present item set file: reward index %d", index)
	}
	offset := 2 + index*4
	binary.LittleEndian.PutUint16(r.Raw[offset:offset+2], count)
	binary.LittleEndian.PutUint16(r.Raw[offset+2:offset+4], itemCode)
	return nil
}
