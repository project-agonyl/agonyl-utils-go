// Package monsterbin reads and writes the A3 clientmonster bin binary format:
// a little-endian uint32 entry count followed by fixed-size monster entries.
package monsterbin

import (
	"encoding/binary"
	"io"

	"github.com/cyberinferno/go-utils/utils"
)

// MonsterBinItem is a single monster record (ID, name, and reserved bytes).
// Name is 0x1F bytes; Unknown is 0x3D bytes of reserved/padding data.
type MonsterBinItem struct {
	ID      uint32
	Name    [0x1F]byte
	Unknown [0x3D]byte
}

// MonsterBin is a slice of monster entries as stored in the bin file.
type MonsterBin []MonsterBinItem

// Read reads a monster bin from r: entry count then each MonsterBinItem.
// Returns the decoded slice or an error if the stream is truncated or invalid.
func Read(r io.Reader) (MonsterBin, error) {
	var entryCount uint32
	if err := binary.Read(r, binary.LittleEndian, &entryCount); err != nil {
		return nil, err
	}

	monsterData := make(MonsterBin, entryCount)
	for i := range monsterData {
		if err := binary.Read(r, binary.LittleEndian, &monsterData[i]); err != nil {
			return nil, err
		}
	}

	return monsterData, nil
}

// Write writes data to w in monster bin format: entry count then each item.
func Write(w io.Writer, data MonsterBin) error {
	entryCount := uint32(len(data))
	if err := binary.Write(w, binary.LittleEndian, entryCount); err != nil {
		return err
	}

	for i := range data {
		if err := binary.Write(w, binary.LittleEndian, &data[i]); err != nil {
			return err
		}
	}

	return nil
}

// GetName returns the name of the monster as a string.
func (m *MonsterBinItem) GetName() string {
	return utils.ReadStringFromBytes(m.Name[:])
}
