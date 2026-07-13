package passiveskillfile

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRoundTripAndKnownFieldUpdate(t *testing.T) {
	original := make([]byte, RecordSize*2)
	for index := range original {
		original[index] = byte(index)
	}

	data, err := Read(bytes.NewReader(original))
	require.NoError(t, err)
	var encoded bytes.Buffer
	require.NoError(t, Write(&encoded, data))
	require.Equal(t, original, encoded.Bytes())

	before := data[0].Raw
	require.NoError(t, data[0].SetEffect(4, 0xdeadbeef))
	require.Equal(t, before[:0x28], data[0].Raw[:0x28])
	require.Equal(t, []byte{0xef, 0xbe, 0xad, 0xde}, data[0].Raw[0x28:0x2c])
	require.Equal(t, before[0x2c:], data[0].Raw[0x2c:])
}

func TestReadRejectsMisalignedData(t *testing.T) {
	_, err := Read(bytes.NewReader(make([]byte, RecordSize-1)))
	require.Error(t, err)
}
