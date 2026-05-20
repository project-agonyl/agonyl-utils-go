package dropfile

import (
	"bytes"
	"encoding/binary"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRead_EmptyStream(t *testing.T) {
	data, err := Read(bytes.NewReader(nil))
	require.NoError(t, err)
	assert.NotNil(t, data)
	assert.Empty(t, data)
}

func TestRead_KnownBytesDecode(t *testing.T) {
	raw := []byte{
		0x02, 0x01, // ItemID
		0x04, 0x03, // DropRate
		0x06, 0x05, // DropGroup
	}

	data, err := Read(bytes.NewReader(raw))
	require.NoError(t, err)
	require.Len(t, data, 1)
	assert.Equal(t, Drop{
		ItemID:    0x0102,
		DropRate:  0x0304,
		DropGroup: 0x0506,
	}, data[0])
}

func TestRead_MultipleDrops(t *testing.T) {
	drops := DropFile{
		{ItemID: 100, DropRate: 10, DropGroup: 1},
		{ItemID: 200, DropRate: 20, DropGroup: 2},
		{ItemID: EmptyItemID, DropRate: 0, DropGroup: 0},
	}
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, drops))

	data, err := Read(&buf)
	require.NoError(t, err)
	assert.Equal(t, drops, data)
}

func TestWrite_SingleDrop_ByteLayout(t *testing.T) {
	drop := Drop{
		ItemID:    0x0102,
		DropRate:  0x0304,
		DropGroup: 0x0506,
	}

	var buf bytes.Buffer
	require.NoError(t, Write(&buf, DropFile{drop}))

	expected := []byte{
		0x02, 0x01, // ItemID
		0x04, 0x03, // DropRate
		0x06, 0x05, // DropGroup
	}
	assert.Equal(t, expected, buf.Bytes())
	assert.Equal(t, DropSize, buf.Len())
}

func TestWriteThenRead_RoundTrip(t *testing.T) {
	original := DropFile{
		{ItemID: 0x4001, DropRate: 100, DropGroup: 1},
		{ItemID: 0x8002, DropRate: 200, DropGroup: 2},
		{ItemID: EmptyItemID, DropRate: 0, DropGroup: 0},
	}

	var buf bytes.Buffer
	require.NoError(t, Write(&buf, original))

	read, err := Read(&buf)
	require.NoError(t, err)
	assert.Equal(t, original, read)
}

func TestRead_TruncatedStream(t *testing.T) {
	_, err := Read(bytes.NewReader(bytes.Repeat([]byte{0}, DropSize-1)))
	require.Error(t, err)
	assert.ErrorIs(t, err, io.ErrUnexpectedEOF)
}

func TestRead_PartialSecondDrop(t *testing.T) {
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, DropFile{{ItemID: 1}}))
	buf.Write(bytes.Repeat([]byte{0}, 2))

	_, err := Read(&buf)
	require.Error(t, err)
	assert.ErrorIs(t, err, io.ErrUnexpectedEOF)
}

func TestRead_InvalidReader(t *testing.T) {
	_, err := Read(errReader{})
	assert.Error(t, err)
}

func TestWrite_Empty(t *testing.T) {
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, nil))
	assert.Empty(t, buf.Bytes())

	data, err := Read(&buf)
	require.NoError(t, err)
	assert.NotNil(t, data)
	assert.Empty(t, data)
}

func TestWrite_EmptySlice(t *testing.T) {
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, DropFile{}))
	assert.Empty(t, buf.Bytes())
}

func TestWrite_InvalidWriter(t *testing.T) {
	err := Write(errWriter{}, DropFile{{ItemID: 1}})
	assert.Error(t, err)
}

func TestDrop_Size(t *testing.T) {
	assert.Equal(t, DropSize, binary.Size(Drop{}))
}

func TestConstants(t *testing.T) {
	assert.Equal(t, uint16(0xFFFF), EmptyItemID)
	assert.Equal(t, uint16(0x3FFF), ItemIDMask)
	assert.Equal(t, uint16(0x0001), uint16(0x4001)&ItemIDMask)
}

type errReader struct{}

func (errReader) Read([]byte) (int, error) {
	return 0, io.ErrClosedPipe
}

type errWriter struct{}

func (errWriter) Write([]byte) (int, error) {
	return 0, io.ErrShortWrite
}
