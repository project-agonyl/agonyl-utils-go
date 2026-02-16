// Package npcfile reads and writes NPC file binary format:
// a fixed-size little-endian record containing NPC stats, attacks, and name.
package npcfile

import (
	"encoding/binary"
	"io"

	"github.com/cyberinferno/go-utils/utils"
)

// NPCFileData is a single NPC record as stored in the NPC file.
// Name is 0x14 bytes; Attacks holds up to 3 attack definitions.
type NPCFileData struct {
	Name                [0x14]byte // Null-padded display name
	Id                  uint16
	RespawnRate         uint16
	AttackTypeInfo      byte
	TargetSelectionInfo byte
	Defense             byte
	AdditionalDefense   byte
	Attacks             [0x3]NPCAttack
	AttackSpeedLow      uint16
	AttackSpeedHigh     uint16
	MovementSpeed       uint32
	Level               byte
	PlayerExp           uint16
	Appearance          byte
	HP                  uint32
	BlueAttackDefense   uint16
	RedAttackDefense    uint16
	GreyAttackDefense   uint16
	MercenaryExp        uint16
	Unknown             uint16
}

// NPCAttack describes one attack slot for an NPC (range, area, damage).
type NPCAttack struct {
	Range            uint16
	Area             uint16
	Damage           uint16
	AdditionalDamage uint16
}

// Read reads a single NPC record from r in little-endian binary format.
// Returns the decoded NPCFileData or an error if the stream is truncated or invalid.
func Read(r io.Reader) (NPCFileData, error) {
	var data NPCFileData
	if err := binary.Read(r, binary.LittleEndian, &data); err != nil {
		return NPCFileData{}, err
	}

	return data, nil
}

// Write writes data to w in NPC file binary format (little-endian).
func Write(w io.Writer, data NPCFileData) error {
	if err := binary.Write(w, binary.LittleEndian, data); err != nil {
		return err
	}

	return nil
}

// GetName returns the NPC display name as a string (trimmed of null padding).
func (n *NPCFileData) GetName() string {
	return utils.ReadStringFromBytes(n.Name[:])
}
