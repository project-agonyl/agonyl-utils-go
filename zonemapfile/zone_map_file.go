// Package zonemapfile reads and writes Zone Server map files.
package zonemapfile

import (
	"encoding/binary"
	"fmt"
	"io"
	"strings"
)

const (
	// Width is the fixed horizontal cell count.
	Width = 256
	// Height is the fixed vertical cell count.
	Height = 256
	// CellCount is the fixed cell count.
	CellCount = Width * Height
	// HeaderSize is the encoded map header size.
	HeaderSize = 23
	// WarpSize is the encoded warp row size.
	WarpSize = 6

	nameOffset             = 20
	nameSize               = 2
	warpCountOffset        = 22
	moveMask        uint32 = 1
	warpShift              = 11
	warpMask        uint32 = 0x0f
	pkShift                = 15
	pkMask          uint32 = 0x03
)

// Warp preserves one warp row.
type Warp struct {
	Raw [WarpSize]byte
}

// Data preserves one complete map file, including opaque header and trailing bytes.
type Data struct {
	Header   [HeaderSize]byte
	Warps    []Warp
	Cells    [CellCount]uint32
	Trailing []byte
}

// Read reads one Zone Server map file.
func Read(r io.Reader) (Data, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return Data{}, err
	}

	if len(b) < HeaderSize {
		return Data{}, fmt.Errorf("zone map file: got %d bytes, need at least %d", len(b), HeaderSize)
	}

	warpCount := int(b[warpCountOffset])
	meshOffset := HeaderSize + warpCount*WarpSize
	want := meshOffset + CellCount*4
	if len(b) < want {
		return Data{}, fmt.Errorf("zone map file: got %d bytes, need at least %d", len(b), want)
	}

	var data Data
	copy(data.Header[:], b[:HeaderSize])
	data.Warps = make([]Warp, warpCount)
	for index := range data.Warps {
		offset := HeaderSize + index*WarpSize
		copy(data.Warps[index].Raw[:], b[offset:offset+WarpSize])
	}

	for index := range data.Cells {
		offset := meshOffset + index*4
		data.Cells[index] = binary.LittleEndian.Uint32(b[offset : offset+4])
	}

	data.Trailing = append([]byte(nil), b[want:]...)
	return data, nil
}

// Write writes one Zone Server map file.
func Write(w io.Writer, data Data) error {
	if len(data.Warps) > 255 {
		return fmt.Errorf("zone map file: got %d warps, max 255", len(data.Warps))
	}

	data.Header[warpCountOffset] = byte(len(data.Warps))
	if _, err := w.Write(data.Header[:]); err != nil {
		return err
	}

	for _, warp := range data.Warps {
		if _, err := w.Write(warp.Raw[:]); err != nil {
			return err
		}
	}

	var cell [4]byte
	for _, raw := range data.Cells {
		binary.LittleEndian.PutUint32(cell[:], raw)
		if _, err := w.Write(cell[:]); err != nil {
			return err
		}
	}

	_, err := w.Write(data.Trailing)
	return err
}

// Name returns the recovered two-byte map name.
func (d Data) Name() string {
	return strings.TrimRight(string(d.Header[nameOffset:nameOffset+nameSize]), "\x00")
}

// SetName updates the recovered two-byte map name.
func (d *Data) SetName(value string) error {
	if len(value) > nameSize {
		return fmt.Errorf("zone map name: got %d bytes, max %d", len(value), nameSize)
	}

	clear(d.Header[nameOffset : nameOffset+nameSize])
	copy(d.Header[nameOffset:nameOffset+nameSize], value)
	return nil
}

func (w Warp) MapID() uint16   { return binary.LittleEndian.Uint16(w.Raw[0:2]) }
func (w Warp) Cell() uint16    { return binary.LittleEndian.Uint16(w.Raw[2:4]) }
func (w Warp) Unknown() uint16 { return binary.LittleEndian.Uint16(w.Raw[4:6]) }

func (w *Warp) SetMapID(value uint16)   { binary.LittleEndian.PutUint16(w.Raw[0:2], value) }
func (w *Warp) SetCell(value uint16)    { binary.LittleEndian.PutUint16(w.Raw[2:4], value) }
func (w *Warp) SetUnknown(value uint16) { binary.LittleEndian.PutUint16(w.Raw[4:6], value) }

// CellIndex returns the zero-based cell index for a coordinate.
func CellIndex(x byte, y byte) int {
	return int(y)*Width + int(x)
}

// CanMove reports whether the recovered movement bit is set.
func CanMove(raw uint32) bool { return raw&moveMask != 0 }

// SetCanMove updates only the recovered movement bit.
func SetCanMove(raw uint32, value bool) uint32 {
	if value {
		return raw | moveMask
	}

	return raw &^ moveMask
}

// PKLevel returns the recovered two-bit PK level.
func PKLevel(raw uint32) byte { return byte((raw >> pkShift) & pkMask) }

// SetPKLevel updates only the recovered PK level bits.
func SetPKLevel(raw uint32, value byte) (uint32, error) {
	if value > byte(pkMask) {
		return raw, fmt.Errorf("zone map PK level %d out of range", value)
	}

	return (raw &^ (pkMask << pkShift)) | uint32(value)<<pkShift, nil
}

// WarpIndex returns the recovered four-bit warp index and whether it is populated.
func WarpIndex(raw uint32) (byte, bool) {
	value := byte((raw >> warpShift) & warpMask)
	return value, value != byte(warpMask)
}

// SetWarpIndex updates only the recovered warp-index bits. Nil stores the no-warp sentinel.
func SetWarpIndex(raw uint32, value *byte) (uint32, error) {
	index := byte(warpMask)
	if value != nil {
		if *value >= byte(warpMask) {
			return raw, fmt.Errorf("zone map warp index %d out of range", *value)
		}

		index = *value
	}

	return (raw &^ (warpMask << warpShift)) | uint32(index)<<warpShift, nil
}
