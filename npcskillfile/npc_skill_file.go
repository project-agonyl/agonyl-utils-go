// Package npcskillfile reads and writes Zone Server NPCSkill files.
package npcskillfile

import (
	"encoding/binary"
	"fmt"
	"io"
)

// RecordSize is the encoded byte size of one NPC skill row.
const RecordSize = 0x12

// Record preserves one NPC skill row.
type Record struct {
	Raw [RecordSize]byte
}

// Data is a complete NPCSkill file.
type Data []Record

// Read reads all aligned NPC skill rows.
func Read(r io.Reader) (Data, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	if len(b)%RecordSize != 0 {
		return nil, fmt.Errorf("npc skill file: got %d bytes, record size %d", len(b), RecordSize)
	}

	data := make(Data, len(b)/RecordSize)
	for index := range data {
		copy(data[index].Raw[:], b[index*RecordSize:(index+1)*RecordSize])
	}

	return data, nil
}

// Write writes all NPC skill rows.
func Write(w io.Writer, data Data) error {
	for _, record := range data {
		if _, err := w.Write(record.Raw[:]); err != nil {
			return err
		}
	}

	return nil
}

func (r Record) NPCType() uint16         { return binary.LittleEndian.Uint16(r.Raw[0:2]) }
func (r Record) Kind() byte              { return r.Raw[2] }
func (r Record) AttackType() byte        { return r.Raw[3] }
func (r Record) OneTargetRange() byte    { return r.Raw[8] }
func (r Record) RangeRadius() byte       { return r.Raw[9] }
func (r Record) CooldownSeconds() uint16 { return binary.LittleEndian.Uint16(r.Raw[10:12]) }
func (r Record) EffectParam() uint16     { return binary.LittleEndian.Uint16(r.Raw[12:14]) }
func (r Record) EffectCode() uint16      { return binary.LittleEndian.Uint16(r.Raw[14:16]) }
func (r Record) EffectValue() uint16     { return binary.LittleEndian.Uint16(r.Raw[16:18]) }

func (r *Record) SetNPCType(value uint16)         { binary.LittleEndian.PutUint16(r.Raw[0:2], value) }
func (r *Record) SetKind(value byte)              { r.Raw[2] = value }
func (r *Record) SetAttackType(value byte)        { r.Raw[3] = value }
func (r *Record) SetOneTargetRange(value byte)    { r.Raw[8] = value }
func (r *Record) SetRangeRadius(value byte)       { r.Raw[9] = value }
func (r *Record) SetCooldownSeconds(value uint16) { binary.LittleEndian.PutUint16(r.Raw[10:12], value) }
func (r *Record) SetEffectParam(value uint16)     { binary.LittleEndian.PutUint16(r.Raw[12:14], value) }
func (r *Record) SetEffectCode(value uint16)      { binary.LittleEndian.PutUint16(r.Raw[14:16], value) }
func (r *Record) SetEffectValue(value uint16)     { binary.LittleEndian.PutUint16(r.Raw[16:18], value) }
