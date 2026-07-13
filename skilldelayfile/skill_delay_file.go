// Package skilldelayfile reads and writes Zone Server SkillDelay.dat files.
package skilldelayfile

import (
	"encoding/binary"
	"fmt"
	"io"
)

// RecordSize is the encoded byte size of one skill-delay row.
const RecordSize = 5

// Record preserves one skill-delay row.
type Record struct {
	Raw [RecordSize]byte
}

// Data is a complete SkillDelay.dat file.
type Data []Record

// Read reads all aligned skill-delay rows.
func Read(r io.Reader) (Data, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	if len(b)%RecordSize != 0 {
		return nil, fmt.Errorf("skill delay file: got %d bytes, record size %d", len(b), RecordSize)
	}

	data := make(Data, len(b)/RecordSize)
	for index := range data {
		copy(data[index].Raw[:], b[index*RecordSize:(index+1)*RecordSize])
	}

	return data, nil
}

// Write writes all skill-delay rows.
func Write(w io.Writer, data Data) error {
	for _, record := range data {
		if _, err := w.Write(record.Raw[:]); err != nil {
			return err
		}
	}

	return nil
}

func (r Record) ClassIndex() byte   { return r.Raw[0] }
func (r Record) SkillIndex() uint16 { return binary.LittleEndian.Uint16(r.Raw[1:3]) }
func (r Record) DelayMS() uint16    { return binary.LittleEndian.Uint16(r.Raw[3:5]) }
func (r Record) DelayInfo() byte    { return r.Raw[4] }

func (r *Record) SetClassIndex(value byte)   { r.Raw[0] = value }
func (r *Record) SetSkillIndex(value uint16) { binary.LittleEndian.PutUint16(r.Raw[1:3], value) }
func (r *Record) SetDelayMS(value uint16)    { binary.LittleEndian.PutUint16(r.Raw[3:5], value) }
func (r *Record) SetDelayInfo(value byte)    { r.Raw[4] = value }
