package spawnlist

import (
	"bytes"
	"encoding/binary"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRead_EmptyStream(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	data, err := Read(buf)
	require.NoError(t, err)
	assert.NotNil(t, data)
	assert.Len(t, data, 0)
}

func TestRead_SingleItem(t *testing.T) {
	item := SpawnListItem{
		Id:          1,
		X:           10,
		Y:           20,
		Unknown1:    0,
		Orientation: 3,
		SpwanStep:   1,
	}
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, SpawnList{item}))

	data, err := Read(&buf)
	require.NoError(t, err)
	require.Len(t, data, 1)
	assert.Equal(t, item, data[0])
}

func TestRead_MultipleItems(t *testing.T) {
	items := SpawnList{
		{Id: 1, X: 1, Y: 1, Orientation: 0, SpwanStep: 0},
		{Id: 2, X: 2, Y: 2, Orientation: 1, SpwanStep: 1},
		{Id: 3, X: 3, Y: 3, Orientation: 2, SpwanStep: 2},
	}
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, items))

	data, err := Read(&buf)
	require.NoError(t, err)
	require.Len(t, data, 3)
	assert.Equal(t, items[0], data[0])
	assert.Equal(t, items[1], data[1])
	assert.Equal(t, items[2], data[2])
}

func TestRead_TruncatedStream(t *testing.T) {
	// One full item is 8 bytes; 5 bytes is truncated
	buf := bytes.NewBuffer([]byte{0x01, 0x00, 0x0A, 0x14, 0x00})
	_, err := Read(buf)
	assert.Error(t, err)
	assert.ErrorIs(t, err, io.ErrUnexpectedEOF)
}

func TestRead_PartialSecondItem(t *testing.T) {
	// 8 + 4 = 12 bytes (one full item + half of second)
	var buf bytes.Buffer
	item := SpawnListItem{Id: 1, X: 0, Y: 0}
	require.NoError(t, binary.Write(&buf, binary.LittleEndian, item))
	buf.Write([]byte{0x02, 0x00, 0x01, 0x02}) // partial second item

	_, err := Read(&buf)
	assert.Error(t, err)
	assert.ErrorIs(t, err, io.ErrUnexpectedEOF)
}

func TestRead_InvalidReader(t *testing.T) {
	var r errReader
	_, err := Read(&r)
	assert.Error(t, err)
}

type errReader struct{}

func (errReader) Read([]byte) (int, error) {
	return 0, io.ErrClosedPipe
}

func TestWrite_Empty(t *testing.T) {
	var buf bytes.Buffer
	err := Write(&buf, nil)
	require.NoError(t, err)
	assert.Empty(t, buf.Bytes())

	data, err := Read(&buf)
	require.NoError(t, err)
	assert.Len(t, data, 0)
}

func TestWrite_EmptySlice(t *testing.T) {
	var buf bytes.Buffer
	err := Write(&buf, SpawnList{})
	require.NoError(t, err)
	assert.Empty(t, buf.Bytes())
}

func TestWrite_SingleItem_ByteLayout(t *testing.T) {
	item := SpawnListItem{
		Id:          0x0102, // little-endian: 02 01
		X:           0x0A,
		Y:           0x14,
		Unknown1:    0x0304, // 04 03
		Orientation: 0x05,
		SpwanStep:   0x06,
	}
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, SpawnList{item}))

	expected := []byte{
		0x02, 0x01, // Id
		0x0A,       // X
		0x14,       // Y
		0x04, 0x03, // Unknown1
		0x05, // Orientation
		0x06, // SpwanStep
	}
	assert.Equal(t, expected, buf.Bytes())
	assert.Equal(t, binary.Size(SpawnListItem{}), buf.Len())
}

func TestWrite_MultipleItems_Size(t *testing.T) {
	items := SpawnList{
		{Id: 1},
		{Id: 2},
		{Id: 3},
	}
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, items))

	expectedLen := 3 * binary.Size(SpawnListItem{})
	assert.Equal(t, expectedLen, buf.Len())
}

func TestWrite_InvalidWriter(t *testing.T) {
	var w errWriter
	err := Write(&w, SpawnList{{Id: 1}})
	assert.Error(t, err)
}

type errWriter struct{}

func (errWriter) Write([]byte) (int, error) {
	return 0, io.ErrShortWrite
}

func TestWriteThenRead_RoundTrip(t *testing.T) {
	original := SpawnList{
		{Id: 100, X: 10, Y: 20, Unknown1: 0x1234, Orientation: 2, SpwanStep: 1},
		{Id: 200, X: 30, Y: 40, Unknown1: 0x5678, Orientation: 0, SpwanStep: 0},
		{Id: 300, X: 0, Y: 0, Unknown1: 0, Orientation: 15, SpwanStep: 255},
	}

	var buf bytes.Buffer
	require.NoError(t, Write(&buf, original))

	data, err := Read(&buf)
	require.NoError(t, err)
	require.Len(t, data, 3)
	assert.Equal(t, original[0], data[0])
	assert.Equal(t, original[1], data[1])
	assert.Equal(t, original[2], data[2])
}

func TestRead_KnownBytesDecode(t *testing.T) {
	// Manually build bytes for one item and decode
	raw := []byte{
		0x34, 0x12, // Id = 0x1234
		0x0A,       // X = 10
		0x14,       // Y = 20
		0xCD, 0xAB, // Unknown1 = 0xABCD
		0x03,       // Orientation = 3
		0x02,       // SpwanStep = 2
	}
	buf := bytes.NewReader(raw)
	data, err := Read(buf)
	require.NoError(t, err)
	require.Len(t, data, 1)
	assert.Equal(t, uint16(0x1234), data[0].Id)
	assert.Equal(t, byte(10), data[0].X)
	assert.Equal(t, byte(20), data[0].Y)
	assert.Equal(t, uint16(0xABCD), data[0].Unknown1)
	assert.Equal(t, byte(3), data[0].Orientation)
	assert.Equal(t, byte(2), data[0].SpwanStep)
}

func TestSpawnListItem_Size(t *testing.T) {
	// Ensure struct size is as expected (8 bytes) for layout tests
	size := binary.Size(SpawnListItem{})
	assert.Equal(t, 8, size, "SpawnListItem must be 8 bytes for binary format")
}
