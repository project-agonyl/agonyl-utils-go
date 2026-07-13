// Package hiredsoldierskillfile reads and writes Zone Server HSST skill files.
package hiredsoldierskillfile

import (
	"encoding/binary"
	"fmt"
	"io"
)

const (
	// RecordSize is the encoded byte size of one hired-soldier skill row.
	RecordSize = 0x2d0
	// MaxRecordCount is the maximum number of stored nonzero skill slots.
	MaxRecordCount = 63
	// LevelCount is the maximum encoded level count.
	LevelCount = 10
	// LevelStride is the encoded byte size of one level row.
	LevelStride = 0x47
)

// Record preserves one hired-soldier skill row.
type Record struct {
	Raw [RecordSize]byte
}

// Data is one HSST class file.
type Data []Record

// Level contains the recovered cost fields for one level.
type Level struct {
	RequiredItemCode uint16
	SkillPointCost   byte
	MoneyCost        uint32
	LoreCost         uint32
}

// Read reads one aligned HSST class file.
func Read(r io.Reader) (Data, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	if len(b)%RecordSize != 0 || len(b)/RecordSize > MaxRecordCount {
		return nil, fmt.Errorf("hired-soldier skill file: got %d bytes, record size %d, max records %d", len(b), RecordSize, MaxRecordCount)
	}

	data := make(Data, len(b)/RecordSize)
	for index := range data {
		copy(data[index].Raw[:], b[index*RecordSize:(index+1)*RecordSize])
	}

	return data, nil
}

// Write writes one HSST class file.
func Write(w io.Writer, data Data) error {
	if len(data) > MaxRecordCount {
		return fmt.Errorf("hired-soldier skill file: got %d records, max %d", len(data), MaxRecordCount)
	}

	for _, record := range data {
		if _, err := w.Write(record.Raw[:]); err != nil {
			return err
		}
	}

	return nil
}

// Level returns the recovered cost fields for one one-based level.
func (r Record) Level(level int) (Level, bool) {
	if level < 1 || level > LevelCount {
		return Level{}, false
	}

	offset := (level - 1) * LevelStride
	return Level{
		RequiredItemCode: binary.LittleEndian.Uint16(r.Raw[offset+0x0a:]),
		SkillPointCost:   r.Raw[offset+0x0c],
		MoneyCost:        binary.LittleEndian.Uint32(r.Raw[offset+0x0d:]),
		LoreCost:         binary.LittleEndian.Uint32(r.Raw[offset+0x11:]),
	}, true
}

// SetLevel updates only the recovered cost bytes for one one-based level.
func (r *Record) SetLevel(level int, value Level) error {
	if level < 1 || level > LevelCount {
		return fmt.Errorf("hired-soldier skill level %d out of range", level)
	}

	offset := (level - 1) * LevelStride
	binary.LittleEndian.PutUint16(r.Raw[offset+0x0a:], value.RequiredItemCode)
	r.Raw[offset+0x0c] = value.SkillPointCost
	binary.LittleEndian.PutUint32(r.Raw[offset+0x0d:], value.MoneyCost)
	binary.LittleEndian.PutUint32(r.Raw[offset+0x11:], value.LoreCost)
	return nil
}
