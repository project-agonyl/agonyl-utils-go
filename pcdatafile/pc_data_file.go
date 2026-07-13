// Package pcdatafile reads and writes Zone Server PC data files.
package pcdatafile

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Size is the encoded byte size of a PC data file.
const Size = 0x1c

// Data preserves one PC data row, including fields not yet named.
type Data struct {
	Raw [Size]byte
}

// Read reads one PC data file.
func Read(r io.Reader) (Data, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return Data{}, err
	}

	if len(b) != Size {
		return Data{}, fmt.Errorf("pc data file: got %d bytes, want %d", len(b), Size)
	}

	var data Data
	copy(data.Raw[:], b)
	return data, nil
}

// Write writes one PC data file.
func Write(w io.Writer, data Data) error {
	_, err := w.Write(data.Raw[:])
	return err
}

// Value returns the two-byte field at index.
func (d Data) Value(index int) (uint16, bool) {
	if index < 0 || index >= Size/2 {
		return 0, false
	}

	offset := index * 2
	return binary.LittleEndian.Uint16(d.Raw[offset : offset+2]), true
}

// SetValue updates the two-byte field at index.
func (d *Data) SetValue(index int, value uint16) error {
	if index < 0 || index >= Size/2 {
		return fmt.Errorf("pc data field index %d out of range", index)
	}

	offset := index * 2
	binary.LittleEndian.PutUint16(d.Raw[offset:offset+2], value)
	return nil
}
