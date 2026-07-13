package npcfavorfile

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRoundTripAndKnownFieldUpdate(t *testing.T) {
	original := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	data, err := Read(bytes.NewReader(original))
	require.NoError(t, err)
	var encoded bytes.Buffer
	require.NoError(t, Write(&encoded, data))
	require.Equal(t, original, encoded.Bytes())

	data[0].SetNPCType(42)
	require.Equal(t, original[:4], data[0].Raw[:4])
	require.Equal(t, []byte{42, 0, 0, 0}, data[0].Raw[4:])
}

func TestReadRejectsMisalignedData(t *testing.T) {
	_, err := Read(bytes.NewReader(make([]byte, RecordSize-1)))
	require.Error(t, err)
}
