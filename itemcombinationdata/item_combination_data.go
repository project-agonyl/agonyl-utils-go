// Package itemcombinationdata reads and writes A3 ItemCombinationData crafting
// files: a contiguous sequence of little-endian crafting formula records.
package itemcombinationdata

import (
	"bytes"
	"encoding/binary"
	"io"
)

// FormulaSize is the encoded byte size of one CraftFormula.
const FormulaSize = 32

// CraftFormula is one ItemCombinationData crafting recipe record.
type CraftFormula struct {
	Item1       uint16
	Item2       uint16
	Item3       uint16
	Item4       uint16
	Item5       uint16
	Item6       uint16
	Item7       uint16
	Item8       uint16
	Item9       uint16
	Item10      uint16
	SuccessRate uint16
	Outcome     uint16
	Unknown1    uint16
	Unknown2    uint16
	Unknown3    uint16
	Unknown4    uint16
}

// ItemCombinationData is a slice of crafting formulas as stored in the file.
type ItemCombinationData []CraftFormula

// Read reads an ItemCombinationData crafting file from r.
// The entire stream is decoded as a contiguous sequence of CraftFormula values.
// Returns io.ErrUnexpectedEOF when the byte length is not a multiple of 32.
func Read(r io.Reader) (ItemCombinationData, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	if len(b)%FormulaSize != 0 {
		return nil, io.ErrUnexpectedEOF
	}

	data := make(ItemCombinationData, len(b)/FormulaSize)
	if len(data) == 0 {
		return data, nil
	}

	if err := binary.Read(bytes.NewReader(b), binary.LittleEndian, &data); err != nil {
		return nil, err
	}

	return data, nil
}

// Write writes data to w in ItemCombinationData crafting file format.
func Write(w io.Writer, data ItemCombinationData) error {
	return binary.Write(w, binary.LittleEndian, data)
}
