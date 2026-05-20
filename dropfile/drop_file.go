// Package dropfile reads and writes A3 monster .itm drop files:
// a contiguous sequence of little-endian drop records.
package dropfile

import (
	"bytes"
	"encoding/binary"
	"io"
)

// DropSize is the encoded byte size of one Drop.
const DropSize = 6

// EmptyItemID is the item ID used for an empty drop slot.
const EmptyItemID uint16 = 0xFFFF

// ItemIDMask strips item flag bits before item-name lookup.
const ItemIDMask uint16 = 0x3FFF

// Drop is one monster drop record as stored in a .itm file.
type Drop struct {
	ItemID    uint16
	DropRate  uint16
	DropGroup uint16
}

// DropFile is a slice of monster drop records as stored in the file.
type DropFile []Drop

// Read reads a monster .itm drop file from r.
// The entire stream is decoded as contiguous Drop values.
// Returns io.ErrUnexpectedEOF when the byte length is not a multiple of 6.
func Read(r io.Reader) (DropFile, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	if len(b)%DropSize != 0 {
		return nil, io.ErrUnexpectedEOF
	}

	data := make(DropFile, len(b)/DropSize)
	if len(data) == 0 {
		return data, nil
	}

	if err := binary.Read(bytes.NewReader(b), binary.LittleEndian, &data); err != nil {
		return nil, err
	}

	return data, nil
}

// Write writes data to w in monster .itm drop file format.
func Write(w io.Writer, data DropFile) error {
	return binary.Write(w, binary.LittleEndian, data)
}
