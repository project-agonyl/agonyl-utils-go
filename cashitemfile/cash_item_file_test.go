package cashitemfile

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRoundTripAndKnownFieldUpdate(t *testing.T) {
	original := make([]byte, RecordSize*2)
	for i := range original {
		original[i] = byte(i)
	}
	data, err := Read(bytes.NewReader(original))
	require.NoError(t, err)
	var encoded bytes.Buffer
	require.NoError(t, Write(&encoded, data))
	require.Equal(t, original, encoded.Bytes())
	before := data[0].Raw
	data[0].SetPrice(0x12345678)
	require.Equal(t, before[:4], data[0].Raw[:4])
	require.Equal(t, []byte{0x78, 0x56, 0x34, 0x12}, data[0].Raw[4:8])
	require.Equal(t, before[8:], data[0].Raw[8:])
}

func TestReadRejectsMisalignment(t *testing.T) {
	_, err := Read(bytes.NewReader(make([]byte, RecordSize+1)))
	require.Error(t, err)
}
