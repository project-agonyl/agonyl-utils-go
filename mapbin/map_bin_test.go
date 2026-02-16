package mapbin

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRead_EmptyBin(t *testing.T) {
	// Count = 0, no items
	buf := bytes.NewBuffer([]byte{0x00, 0x00, 0x00, 0x00})
	data, err := Read(buf)
	require.NoError(t, err)
	assert.NotNil(t, data)
	assert.Len(t, data, 0)
}

func TestRead_SingleItem(t *testing.T) {
	item := MapBinItem{ID: 1}
	copy(item.Name[:], "Forest")

	var buf bytes.Buffer
	require.NoError(t, Write(&buf, MapBin{item}))

	data, err := Read(&buf)
	require.NoError(t, err)
	require.Len(t, data, 1)
	assert.Equal(t, uint32(1), data[0].ID)
	assert.Equal(t, "Forest", data[0].GetName())
}

func TestRead_MultipleItems(t *testing.T) {
	items := MapBin{
		{ID: 1},
		{ID: 2},
		{ID: 3},
	}
	copy(items[0].Name[:], "Forest")
	copy(items[1].Name[:], "Dungeon")
	copy(items[2].Name[:], "Town")

	var buf bytes.Buffer
	require.NoError(t, Write(&buf, items))

	data, err := Read(&buf)
	require.NoError(t, err)
	require.Len(t, data, 3)
	assert.Equal(t, uint32(1), data[0].ID)
	assert.Equal(t, "Forest", data[0].GetName())
	assert.Equal(t, uint32(2), data[1].ID)
	assert.Equal(t, "Dungeon", data[1].GetName())
	assert.Equal(t, uint32(3), data[2].ID)
	assert.Equal(t, "Town", data[2].GetName())
}

func TestRead_TruncatedCount(t *testing.T) {
	// Only 2 bytes instead of 4 for count
	buf := bytes.NewBuffer([]byte{0x01, 0x00})
	_, err := Read(buf)
	assert.Error(t, err)
}

func TestRead_TruncatedItem(t *testing.T) {
	// Count = 1 but stream ends before full item (need 4*6 + 0x20 = 56 bytes after count)
	var buf bytes.Buffer
	count := []byte{0x01, 0x00, 0x00, 0x00}
	buf.Write(count)
	buf.Write(bytes.Repeat([]byte{0}, 40)) // only 40 bytes, need 56

	_, err := Read(&buf)
	assert.Error(t, err)
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
	assert.Equal(t, []byte{0x00, 0x00, 0x00, 0x00}, buf.Bytes())

	data, err := Read(&buf)
	require.NoError(t, err)
	assert.Len(t, data, 0)
}

func TestWrite_EmptySlice(t *testing.T) {
	var buf bytes.Buffer
	err := Write(&buf, MapBin{})
	require.NoError(t, err)
	assert.Len(t, buf.Bytes(), 4)
	assert.Equal(t, []byte{0x00, 0x00, 0x00, 0x00}, buf.Bytes())
}

func TestWriteThenRead_RoundTrip(t *testing.T) {
	original := MapBin{
		{ID: 100, Unknown1: 1, Unknown2: 2},
		{ID: 200, Unknown3: 3, Unknown4: 4, Unknown5: 5},
	}
	copy(original[0].Name[:], "Alpha")
	copy(original[1].Name[:], "Beta")

	var buf bytes.Buffer
	require.NoError(t, Write(&buf, original))

	data, err := Read(&buf)
	require.NoError(t, err)
	require.Len(t, data, 2)
	assert.Equal(t, original[0].ID, data[0].ID)
	assert.Equal(t, original[0].Unknown1, data[0].Unknown1)
	assert.Equal(t, original[0].Unknown2, data[0].Unknown2)
	assert.Equal(t, original[0].Name, data[0].Name)
	assert.Equal(t, original[1].ID, data[1].ID)
	assert.Equal(t, original[1].Unknown3, data[1].Unknown3)
	assert.Equal(t, original[1].Unknown4, data[1].Unknown4)
	assert.Equal(t, original[1].Unknown5, data[1].Unknown5)
	assert.Equal(t, original[1].Name, data[1].Name)
	assert.Equal(t, "Alpha", data[0].GetName())
	assert.Equal(t, "Beta", data[1].GetName())
}

func TestGetName_EmptyName(t *testing.T) {
	var m MapBinItem
	// Name is zero-initialized
	assert.Equal(t, "", m.GetName())
}

func TestGetName_ShortName(t *testing.T) {
	var m MapBinItem
	copy(m.Name[:], "Forest")
	assert.Equal(t, "Forest", m.GetName())
}

func TestGetName_NullPadded(t *testing.T) {
	var m MapBinItem
	copy(m.Name[:], "A\x00\x00\x00")
	assert.Equal(t, "A", m.GetName())
}

func TestGetName_NoNullInBuffer(t *testing.T) {
	var m MapBinItem
	for i := range 0x20 {
		m.Name[i] = 'X'
	}
	// Behavior depends on utils.ReadStringFromBytes; expect at least no panic
	name := m.GetName()
	assert.Len(t, name, 32)
	assert.Equal(t, "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", name)
}
