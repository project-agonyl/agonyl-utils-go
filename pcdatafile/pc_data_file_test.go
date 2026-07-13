package pcdatafile

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRoundTripAndFieldUpdate(t *testing.T) {
	original := make([]byte, Size)
	for index := range original {
		original[index] = byte(index)
	}

	data, err := Read(bytes.NewReader(original))
	require.NoError(t, err)
	var encoded bytes.Buffer
	require.NoError(t, Write(&encoded, data))
	require.Equal(t, original, encoded.Bytes())

	require.NoError(t, data.SetValue(4, 0xbeef))
	require.Equal(t, original[:8], data.Raw[:8])
	require.Equal(t, []byte{0xef, 0xbe}, data.Raw[8:10])
	require.Equal(t, original[10:], data.Raw[10:])
}

func TestReadRejectsWrongSize(t *testing.T) {
	_, err := Read(bytes.NewReader(make([]byte, Size-1)))
	require.Error(t, err)
}
