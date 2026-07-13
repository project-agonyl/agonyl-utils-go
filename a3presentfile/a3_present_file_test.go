package a3presentfile

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRoundTripAndRewardUpdate(t *testing.T) {
	original := make([]byte, RecordSize)
	for i := range original {
		original[i] = byte(i)
	}
	data, err := Read(bytes.NewReader(original))
	require.NoError(t, err)
	var encoded bytes.Buffer
	require.NoError(t, Write(&encoded, data))
	require.Equal(t, original, encoded.Bytes())
	require.NoError(t, data[0].SetReward(2, 7, 0xbeef))
	offset := NameSize + 8
	require.Equal(t, original[:offset], data[0].Raw[:offset])
	require.Equal(t, []byte{7, 0, 0xef, 0xbe}, data[0].Raw[offset:offset+4])
	require.Equal(t, original[offset+4:], data[0].Raw[offset+4:])
}

func TestReadRejectsMisalignment(t *testing.T) {
	_, err := Read(bytes.NewReader(make([]byte, RecordSize+1)))
	require.Error(t, err)
}
