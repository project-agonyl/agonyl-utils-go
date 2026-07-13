package skilldatafile

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRoundTripAndKnownFieldUpdate(t *testing.T) {
	original := make([]byte, RecordCount*RecordSize)
	for index := range original {
		original[index] = byte(index)
	}

	data, err := Read(bytes.NewReader(original))
	require.NoError(t, err)
	var encoded bytes.Buffer
	require.NoError(t, Write(&encoded, data))
	require.Equal(t, original, encoded.Bytes())

	before := data[0].Raw
	data[0].SetCode(0xff)
	require.Equal(t, byte(0xff), data[0].Raw[0])
	require.Equal(t, before[1:], data[0].Raw[1:])
}

func TestReadAndWriteValidateRecordCount(t *testing.T) {
	_, err := Read(bytes.NewReader(make([]byte, RecordSize)))
	require.Error(t, err)
	require.Error(t, Write(&bytes.Buffer{}, make(Data, 1)))
}

func TestSetLevelPreservesOpaqueTail(t *testing.T) {
	var record Record
	for index := range record.Raw {
		record.Raw[index] = 0xaa
	}
	before := record.Raw
	require.NoError(t, record.SetLevel(1, Level{HPConsume: 7, MPConsume: 8}))
	require.Equal(t, before[0x38:0x3d], record.Raw[0x38:0x3d])
	require.Error(t, record.SetLevel(0, Level{}))
}
