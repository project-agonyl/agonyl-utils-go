// Package squestquizfile reads and writes SQuest QuizTable.dat files.
package squestquizfile

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

const (
	RecordSize   = 0x262
	RecordCount  = 100
	QuestionSize = 201
	AnswerCount  = 5
	AnswerSize   = 81
	FileSize     = RecordSize * RecordCount
)

type Record struct{ Raw [RecordSize]byte }
type Data [RecordCount]Record

func Read(r io.Reader) (Data, error) {
	var data Data
	b, err := io.ReadAll(r)
	if err != nil {
		return data, err
	}
	if len(b) != FileSize {
		return data, fmt.Errorf("squest quiz file: got %d bytes, expected %d", len(b), FileSize)
	}
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
func (r Record) ID() uint16       { return binary.LittleEndian.Uint16(r.Raw[0:2]) }
func (r Record) Question() string { return fixedString(r.Raw[2:0xcb]) }
func (r Record) Answer(index int) (string, bool) {
	if index < 0 || index >= AnswerCount {
		return "", false
	}
	offset := 0xcb + index*AnswerSize
	return fixedString(r.Raw[offset : offset+AnswerSize]), true
}
func (r Record) Correct() uint16             { return binary.LittleEndian.Uint16(r.Raw[0x260:0x262]) }
func (r *Record) SetID(v uint16)             { binary.LittleEndian.PutUint16(r.Raw[0:2], v) }
func (r *Record) SetQuestion(v string) error { return putFixed(r.Raw[2:0xcb], v) }
func (r *Record) SetAnswer(index int, v string) error {
	if index < 0 || index >= AnswerCount {
		return fmt.Errorf("squest quiz file: answer index %d", index)
	}
	offset := 0xcb + index*AnswerSize
	return putFixed(r.Raw[offset:offset+AnswerSize], v)
}
func (r *Record) SetCorrect(v uint16) error {
	if v >= AnswerCount {
		return fmt.Errorf("squest quiz file: correct answer %d", v)
	}
	binary.LittleEndian.PutUint16(r.Raw[0x260:0x262], v)
	return nil
}
func fixedString(b []byte) string {
	if i := bytes.IndexByte(b, 0); i >= 0 {
		b = b[:i]
	}
	return string(b)
}
func putFixed(dst []byte, v string) error {
	if len(v) > len(dst) {
		return fmt.Errorf("squest quiz file: text length %d exceeds %d", len(v), len(dst))
	}
	clear(dst)
	copy(dst, v)
	return nil
}
