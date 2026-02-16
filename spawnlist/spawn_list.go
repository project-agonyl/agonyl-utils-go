// Package spawnlist reads and writes spawn list binary format:
// a contiguous sequence of little-endian spawn entries (position, orientation, etc.).
package spawnlist

import (
	"bytes"
	"encoding/binary"
	"io"
)

// SpawnListItem is a single spawn entry as stored in the spawn list file.
type SpawnListItem struct {
	Id          uint16 // Spawn/npc identifier
	X           byte   // X coordinate
	Y           byte   // Y coordinate
	Unknown1    uint16 // Reserved
	Orientation byte   // Facing direction
	SpwanStep   byte   // Spawn step
}

// SpawnList is a slice of spawn entries as stored in the spawn list file.
type SpawnList []SpawnListItem

// Read reads a spawn list from r.
// The entire stream is decoded as a contiguous sequence of SpawnListItem values until EOF.
// Returns the decoded list or an error if the stream is truncated or invalid.
func Read(r io.Reader) (SpawnList, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	itemSize := binary.Size(SpawnListItem{})
	if len(b)%itemSize != 0 {
		return nil, io.ErrUnexpectedEOF
	}

	n := len(b) / itemSize
	data := make(SpawnList, n)
	if n == 0 {
		return data, nil
	}

	if err := binary.Read(bytes.NewReader(b), binary.LittleEndian, &data); err != nil {
		return nil, err
	}

	return data, nil
}

// Write writes data to w in spawn list binary format.
func Write(w io.Writer, data SpawnList) error {
	if err := binary.Write(w, binary.LittleEndian, data); err != nil {
		return err
	}

	return nil
}
