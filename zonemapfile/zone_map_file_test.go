package zonemapfile

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRoundTripAndBitFieldUpdates(t *testing.T) {
	original := make([]byte, HeaderSize+WarpSize+CellCount*4+3)
	original[20], original[21], original[22] = 'A', '3', 1
	binary.LittleEndian.PutUint16(original[23:25], 7)
	binary.LittleEndian.PutUint32(original[HeaderSize+WarpSize:], 0xf0f0f0f0)
	copy(original[len(original)-3:], []byte{1, 2, 3})

	data, err := Read(bytes.NewReader(original))
	require.NoError(t, err)
	var encoded bytes.Buffer
	require.NoError(t, Write(&encoded, data))
	require.Equal(t, original, encoded.Bytes())
	require.Equal(t, "A3", data.Name())

	raw := data.Cells[0]
	updated := SetCanMove(raw, !CanMove(raw))
	require.Equal(t, raw&^moveMask, updated&^moveMask)
	updated, err = SetPKLevel(raw, 2)
	require.NoError(t, err)
	require.Equal(t, byte(2), PKLevel(updated))
	require.Equal(t, raw&^(pkMask<<pkShift), updated&^(pkMask<<pkShift))
}

func TestReadRejectsTruncatedAndWriteRejectsWarpOverflow(t *testing.T) {
	_, err := Read(bytes.NewReader(make([]byte, HeaderSize-1)))
	require.Error(t, err)
	_, err = Read(bytes.NewReader(make([]byte, HeaderSize+CellCount*4-1)))
	require.Error(t, err)
	require.Error(t, Write(&bytes.Buffer{}, Data{Warps: make([]Warp, 256)}))
}
