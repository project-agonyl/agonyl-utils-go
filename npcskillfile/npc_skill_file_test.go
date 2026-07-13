package npcskillfile

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
	data[0].SetEffectCode(0xbeef)
	require.Equal(t, before[:14], data[0].Raw[:14])
	require.Equal(t, []byte{0xef, 0xbe}, data[0].Raw[14:16])
	require.Equal(t, before[16:], data[0].Raw[16:])
}

func TestReadRejectsMisalignedData(t *testing.T) {
	_, err := Read(bytes.NewReader(make([]byte, RecordSize+1)))
	require.Error(t, err)
}
