// Package mapbin reads and writes the A3 client map bin binary format:
// a little-endian uint32 entry count followed by fixed-size map entries.
package mapbin

import (
	"encoding/binary"
	"io"

	"github.com/cyberinferno/go-utils/utils"
)

// MapBinItem is a single map record (ID, unknown fields, and name).
// Name is 0x20 bytes; Unknown1–Unknown5 are reserved uint32 values.
type MapBinItem struct {
	ID       uint32
	Unknown1 uint32
	Unknown2 uint32
	Unknown3 uint32
	Unknown4 uint32
	Unknown5 uint32
	Name     [0x20]byte
}

// MapBin is a slice of map entries as stored in the bin file.
type MapBin []MapBinItem

// Read reads a map bin from r: entry count then each MapBinItem.
// Returns the decoded slice or an error if the stream is truncated or invalid.
func Read(r io.Reader) (MapBin, error) {
	var entryCount uint32
	if err := binary.Read(r, binary.LittleEndian, &entryCount); err != nil {
		return nil, err
	}

	mapData := make(MapBin, entryCount)
	for i := range mapData {
		if err := binary.Read(r, binary.LittleEndian, &mapData[i]); err != nil {
			return nil, err
		}
	}
	return mapData, nil
}

// Write writes data to w in map bin format: entry count then each item.
func Write(w io.Writer, data MapBin) error {
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

// GetName returns the name of the map as a string.
func (m *MapBinItem) GetName() string {
	return utils.ReadStringFromBytes(m.Name[:])
}
