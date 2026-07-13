package shuecombinationfile

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRoundTripAndFieldUpdate(t *testing.T) {
	original := make([]byte, RecordSize)
	for i := range original {
		original[i] = byte(i)
	}
	data, err := Read(bytes.NewReader(original))
	require.NoError(t, err)
	var encoded bytes.Buffer
	require.NoError(t, Write(&encoded, data))
	require.Equal(t, original, encoded.Bytes())
	require.NoError(t, data[0].SetField(11, 0xbeef))
	require.Equal(t, original[:22], data[0].Raw[:22])
	require.Equal(t, []byte{0xef, 0xbe}, data[0].Raw[22:24])
	require.Equal(t, original[24:], data[0].Raw[24:])
}

func TestReadRejectsMisalignment(t *testing.T) {
	_, err := Read(bytes.NewReader(make([]byte, RecordSize+1)))
	require.Error(t, err)
}
