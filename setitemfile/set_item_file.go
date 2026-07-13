// Package setitemfile reads and writes one Zone Server SIT0 through SIT3 file.
package setitemfile

import (
	"encoding/binary"
	"fmt"
	"io"
)

const (
	FileSize        = 0x0ce4
	RecordSize      = 0x014a
	RecordCount     = 10
	PieceCount      = 10
	BonusPieceCount = 11
	BonusCount      = 10
)

type Record struct{ Raw [RecordSize]byte }
type Data [RecordCount]Record

func Read(r io.Reader) (Data, error) {
	var data Data
	b, err := io.ReadAll(r)
	if err != nil {
		return data, err
	}
	if len(b) != FileSize {
		return data, fmt.Errorf("set item file: got %d bytes, expected %d", len(b), FileSize)
	}
	for i := range data {
		copy(data[i].Raw[:], b[i*RecordSize:(i+1)*RecordSize])
	}
	return data, nil
}

func Write(w io.Writer, data Data) error {
	for _, record := range data {
		if _, err := w.Write(record.Raw[:]); err != nil {
			return err
		}
	}
	return nil
}

func (r Record) Piece(index int) (byte, uint16, bool) {
	if index < 0 || index >= PieceCount {
		return 0, 0, false
	}
	offset := index * 3
	return r.Raw[offset], binary.LittleEndian.Uint16(r.Raw[offset+1 : offset+3]), true
}
func (r *Record) SetPiece(index int, wearSlot byte, code uint16) error {
	if index < 0 || index >= PieceCount {
		return fmt.Errorf("set item file: piece index %d", index)
	}
	offset := index * 3
	r.Raw[offset] = wearSlot
	binary.LittleEndian.PutUint16(r.Raw[offset+1:offset+3], code)
	return nil
}
func (r Record) Bonus(pieceCount int, index int) (byte, uint16, bool) {
	if pieceCount < 1 || pieceCount >= BonusPieceCount || index < 0 || index >= BonusCount {
		return 0, 0, false
	}
	offset := 0x1e + (pieceCount-1)*0x1e + index*3
	return r.Raw[offset], binary.LittleEndian.Uint16(r.Raw[offset+1 : offset+3]), true
}
func (r *Record) SetBonus(pieceCount int, index int, optionID byte, value uint16) error {
	if pieceCount < 1 || pieceCount >= BonusPieceCount || index < 0 || index >= BonusCount {
		return fmt.Errorf("set item file: bonus index %d/%d", pieceCount, index)
	}
	offset := 0x1e + (pieceCount-1)*0x1e + index*3
	r.Raw[offset] = optionID
	binary.LittleEndian.PutUint16(r.Raw[offset+1:offset+3], value)
	return nil
}
