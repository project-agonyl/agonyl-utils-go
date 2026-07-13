// Package oxquizfile reads and writes OXQuizTable.dat files.
package oxquizfile

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

const RecordSize = 0x85
const TextSize = 0x40

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
	data := Data{Records: make([]Record, count), Trailing: append([]byte(nil), b[count*RecordSize:]...)}
	for i := range data.Records {
		copy(data.Records[i].Raw[:], b[i*RecordSize:(i+1)*RecordSize])
		a := data.Records[i].Answer()
		if a != 'O' && a != 'X' {
			return Data{}, fmt.Errorf("ox quiz file: row %d answer %#x", i, a)
		}
	}
	return data, nil
}
func Write(w io.Writer, data Data) error {
	for i, record := range data.Records {
		a := record.Answer()
		if a != 'O' && a != 'X' {
			return fmt.Errorf("ox quiz file: row %d answer %#x", i, a)
		}
		if _, err := w.Write(record.Raw[:]); err != nil {
			return err
		}
	}
	_, err := w.Write(data.Trailing)
	return err
}
func (r Record) RewardCode() uint16       { return binary.LittleEndian.Uint16(r.Raw[0:2]) }
func (r Record) RewardCount() uint16      { return binary.LittleEndian.Uint16(r.Raw[2:4]) }
func (r Record) Answer() byte             { return r.Raw[4] }
func (r Record) Question() string         { return oxString(r.Raw[5:0x45]) }
func (r Record) Explanation() string      { return oxString(r.Raw[0x45:0x85]) }
func (r *Record) SetRewardCode(v uint16)  { binary.LittleEndian.PutUint16(r.Raw[0:2], v) }
func (r *Record) SetRewardCount(v uint16) { binary.LittleEndian.PutUint16(r.Raw[2:4], v) }
func (r *Record) SetAnswer(v byte) error {
	if v != 'O' && v != 'X' {
		return fmt.Errorf("ox quiz file: invalid answer %#x", v)
	}
	r.Raw[4] = v
	return nil
}
func (r *Record) SetQuestion(v string) error    { return putOXString(r.Raw[5:0x45], v) }
func (r *Record) SetExplanation(v string) error { return putOXString(r.Raw[0x45:0x85], v) }
func oxString(b []byte) string {
	if i := bytes.IndexByte(b, 0); i >= 0 {
		b = b[:i]
	}
	return string(bytes.TrimRight(b, " "))
}
func putOXString(dst []byte, v string) error {
	if len(v) > len(dst) {
		return fmt.Errorf("ox quiz file: text length %d exceeds %d", len(v), len(dst))
	}
	clear(dst)
	copy(dst, v)
	return nil
}
