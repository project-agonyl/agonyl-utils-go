// Package skilldatafile reads and writes one Zone Server class skill-data file.
package skilldatafile

import (
	"encoding/binary"
	"fmt"
	"io"
)

const (
	// RecordSize is the encoded byte size of one skill row.
	RecordSize = 0x9d
	// RecordCount is the number of stored nonzero skill slots.
	RecordCount = 63
	// LevelCount is the number of encoded skill levels.
	LevelCount = 3
	// LevelStride is the encoded byte size of one level row.
	LevelStride = 0x30

	levelOffset = 0x0d
)

// Record preserves one class skill row.
type Record struct {
	Raw [RecordSize]byte
}

// Data is one complete class skill-data file.
type Data []Record

// Level describes the named fields in one skill-level row.
type Level struct {
	HPConsume      uint16
	MPConsume      uint16
	Delay          uint16
	Range          byte
	Scope          byte
	LastingTime    byte
	AddDamage      uint16
	AddIceDamage   uint16
	AddFireDamage  uint16
	AddLightDamage uint16
	EffectRate     uint16
	HitRate        byte
	RecoveryRate   uint16
	AttackRate     uint16
	MagicRate      uint16
	DefenseRate    uint16
	IceAttack      uint16
	FireAttack     uint16
	LightAttack    uint16
	IceDefense     uint16
	FireDefense    uint16
	LightDefense   uint16
	MaxHPRate      uint16
	MaxMPRate      uint16
}

// Read reads one complete class skill-data file.
func Read(r io.Reader) (Data, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	want := RecordCount * RecordSize
	if len(b) != want {
		return nil, fmt.Errorf("skill data file: got %d bytes, want %d", len(b), want)
	}

	data := make(Data, RecordCount)
	for index := range data {
		copy(data[index].Raw[:], b[index*RecordSize:(index+1)*RecordSize])
	}

	return data, nil
}

// Write writes one complete class skill-data file.
func Write(w io.Writer, data Data) error {
	if len(data) != RecordCount {
		return fmt.Errorf("skill data file: got %d records, want %d", len(data), RecordCount)
	}

	for _, record := range data {
		if _, err := w.Write(record.Raw[:]); err != nil {
			return err
		}
	}

	return nil
}

func (r Record) Code() byte       { return r.Raw[0] }
func (r Record) Type() byte       { return r.Raw[1] }
func (r Record) SubType() byte    { return r.Raw[2] }
func (r Record) TargetType() byte { return r.Raw[3] }
func (r Record) NeedItem() byte   { return r.Raw[4] }
func (r Record) Reaction() byte   { return r.Raw[5] }
func (r Record) Abnormalcy() byte { return r.Raw[6] }

func (r *Record) SetCode(value byte)       { r.Raw[0] = value }
func (r *Record) SetType(value byte)       { r.Raw[1] = value }
func (r *Record) SetSubType(value byte)    { r.Raw[2] = value }
func (r *Record) SetTargetType(value byte) { r.Raw[3] = value }
func (r *Record) SetNeedItem(value byte)   { r.Raw[4] = value }
func (r *Record) SetReaction(value byte)   { r.Raw[5] = value }
func (r *Record) SetAbnormalcy(value byte) { r.Raw[6] = value }

// MonsterRate returns one of the six recovered monster-rate bytes.
func (r Record) MonsterRate(index int) (byte, bool) {
	if index < 0 || index >= 6 {
		return 0, false
	}

	return r.Raw[7+index], true
}

// SetMonsterRate updates one recovered monster-rate byte.
func (r *Record) SetMonsterRate(index int, value byte) error {
	if index < 0 || index >= 6 {
		return fmt.Errorf("monster rate index %d out of range", index)
	}

	r.Raw[7+index] = value
	return nil
}

// Level returns one one-based skill-level row.
func (r Record) Level(level int) (Level, bool) {
	row, ok := r.levelBytes(level)
	if !ok {
		return Level{}, false
	}

	return Level{
		HPConsume: binary.LittleEndian.Uint16(row[0x00:0x02]), MPConsume: binary.LittleEndian.Uint16(row[0x02:0x04]),
		Delay: binary.LittleEndian.Uint16(row[0x04:0x06]), Range: row[0x06], Scope: row[0x07], LastingTime: row[0x08],
		AddDamage: binary.LittleEndian.Uint16(row[0x09:0x0b]), AddIceDamage: binary.LittleEndian.Uint16(row[0x0b:0x0d]),
		AddFireDamage: binary.LittleEndian.Uint16(row[0x0d:0x0f]), AddLightDamage: binary.LittleEndian.Uint16(row[0x0f:0x11]),
		EffectRate: binary.LittleEndian.Uint16(row[0x11:0x13]), HitRate: row[0x13], RecoveryRate: binary.LittleEndian.Uint16(row[0x13:0x15]),
		AttackRate: binary.LittleEndian.Uint16(row[0x15:0x17]), MagicRate: binary.LittleEndian.Uint16(row[0x17:0x19]),
		DefenseRate: binary.LittleEndian.Uint16(row[0x19:0x1b]), IceAttack: binary.LittleEndian.Uint16(row[0x1b:0x1d]),
		FireAttack: binary.LittleEndian.Uint16(row[0x1d:0x1f]), LightAttack: binary.LittleEndian.Uint16(row[0x1f:0x21]),
		IceDefense: binary.LittleEndian.Uint16(row[0x21:0x23]), FireDefense: binary.LittleEndian.Uint16(row[0x23:0x25]),
		LightDefense: binary.LittleEndian.Uint16(row[0x25:0x27]), MaxHPRate: binary.LittleEndian.Uint16(row[0x27:0x29]),
		MaxMPRate: binary.LittleEndian.Uint16(row[0x29:0x2b]),
	}, true
}

// SetLevel updates the named bytes in one one-based skill-level row.
func (r *Record) SetLevel(level int, value Level) error {
	row, ok := r.levelBytes(level)
	if !ok {
		return fmt.Errorf("skill level %d out of range", level)
	}

	binary.LittleEndian.PutUint16(row[0x00:0x02], value.HPConsume)
	binary.LittleEndian.PutUint16(row[0x02:0x04], value.MPConsume)
	binary.LittleEndian.PutUint16(row[0x04:0x06], value.Delay)
	row[0x06], row[0x07], row[0x08] = value.Range, value.Scope, value.LastingTime
	binary.LittleEndian.PutUint16(row[0x09:0x0b], value.AddDamage)
	binary.LittleEndian.PutUint16(row[0x0b:0x0d], value.AddIceDamage)
	binary.LittleEndian.PutUint16(row[0x0d:0x0f], value.AddFireDamage)
	binary.LittleEndian.PutUint16(row[0x0f:0x11], value.AddLightDamage)
	binary.LittleEndian.PutUint16(row[0x11:0x13], value.EffectRate)
	row[0x13] = value.HitRate
	binary.LittleEndian.PutUint16(row[0x13:0x15], value.RecoveryRate)
	binary.LittleEndian.PutUint16(row[0x15:0x17], value.AttackRate)
	binary.LittleEndian.PutUint16(row[0x17:0x19], value.MagicRate)
	binary.LittleEndian.PutUint16(row[0x19:0x1b], value.DefenseRate)
	binary.LittleEndian.PutUint16(row[0x1b:0x1d], value.IceAttack)
	binary.LittleEndian.PutUint16(row[0x1d:0x1f], value.FireAttack)
	binary.LittleEndian.PutUint16(row[0x1f:0x21], value.LightAttack)
	binary.LittleEndian.PutUint16(row[0x21:0x23], value.IceDefense)
	binary.LittleEndian.PutUint16(row[0x23:0x25], value.FireDefense)
	binary.LittleEndian.PutUint16(row[0x25:0x27], value.LightDefense)
	binary.LittleEndian.PutUint16(row[0x27:0x29], value.MaxHPRate)
	binary.LittleEndian.PutUint16(row[0x29:0x2b], value.MaxMPRate)
	return nil
}

func (r *Record) levelBytes(level int) ([]byte, bool) {
	if level < 1 || level > LevelCount {
		return nil, false
	}

	offset := levelOffset + (level-1)*LevelStride
	return r.Raw[offset : offset+LevelStride], true
}
