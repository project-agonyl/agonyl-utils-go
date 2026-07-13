// Package passiveskillfile reads and writes Zone Server PsvSkill.dat files.
package passiveskillfile

import (
	"encoding/binary"
	"fmt"
	"io"
)

const (
	// RecordSize is the encoded byte size of one passive-skill level row.
	RecordSize = 0x44
	// EffectCount is the number of effect values in a row.
	EffectCount = 11
)

// Record preserves one passive-skill level row.
type Record struct {
	Raw [RecordSize]byte
}

// Data is a complete PsvSkill.dat file.
type Data []Record

// Read reads all aligned passive-skill rows.
func Read(r io.Reader) (Data, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	if len(b)%RecordSize != 0 {
		return nil, fmt.Errorf("passive skill file: got %d bytes, record size %d", len(b), RecordSize)
	}

	data := make(Data, len(b)/RecordSize)
	for index := range data {
		copy(data[index].Raw[:], b[index*RecordSize:(index+1)*RecordSize])
	}

	return data, nil
}

// Write writes all passive-skill rows.
func Write(w io.Writer, data Data) error {
	for _, record := range data {
		if _, err := w.Write(record.Raw[:]); err != nil {
			return err
		}
	}

	return nil
}

func (r Record) PassiveID() uint32        { return binary.LittleEndian.Uint32(r.Raw[0x00:0x04]) }
func (r Record) ClassRestriction() uint32 { return binary.LittleEndian.Uint32(r.Raw[0x04:0x08]) }
func (r Record) EffectKind() uint32       { return binary.LittleEndian.Uint32(r.Raw[0x08:0x0c]) }
func (r Record) Level() uint32            { return binary.LittleEndian.Uint32(r.Raw[0x0c:0x10]) }
func (r Record) RequiredPoints() uint32   { return binary.LittleEndian.Uint32(r.Raw[0x10:0x14]) }
func (r Record) Money() uint32            { return binary.LittleEndian.Uint32(r.Raw[0x14:0x18]) }

func (r *Record) SetPassiveID(value uint32) { binary.LittleEndian.PutUint32(r.Raw[0x00:0x04], value) }
func (r *Record) SetClassRestriction(value uint32) {
	binary.LittleEndian.PutUint32(r.Raw[0x04:0x08], value)
}
func (r *Record) SetEffectKind(value uint32) { binary.LittleEndian.PutUint32(r.Raw[0x08:0x0c], value) }
func (r *Record) SetLevel(value uint32)      { binary.LittleEndian.PutUint32(r.Raw[0x0c:0x10], value) }
func (r *Record) SetRequiredPoints(value uint32) {
	binary.LittleEndian.PutUint32(r.Raw[0x10:0x14], value)
}
func (r *Record) SetMoney(value uint32) { binary.LittleEndian.PutUint32(r.Raw[0x14:0x18], value) }

// Effect returns one recovered effect dword.
func (r Record) Effect(index int) (uint32, bool) {
	if index < 0 || index >= EffectCount {
		return 0, false
	}

	offset := 0x18 + index*4
	return binary.LittleEndian.Uint32(r.Raw[offset : offset+4]), true
}

// SetEffect updates one recovered effect dword.
func (r *Record) SetEffect(index int, value uint32) error {
	if index < 0 || index >= EffectCount {
		return fmt.Errorf("passive skill effect index %d out of range", index)
	}

	offset := 0x18 + index*4
	binary.LittleEndian.PutUint32(r.Raw[offset:offset+4], value)
	return nil
}
