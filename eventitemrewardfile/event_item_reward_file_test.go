package eventitemrewardfile

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRoundTripAndMessageUpdate(t *testing.T) {
	original := make([]byte, RecordSize)
	for i := range original {
		original[i] = byte(i)
	}
	data, err := Read(bytes.NewReader(original))
	require.NoError(t, err)
	var encoded bytes.Buffer
	require.NoError(t, Write(&encoded, data))
	require.Equal(t, original, encoded.Bytes())
	require.NoError(t, data[0].SetMessage("reward"))
	require.Equal(t, original[:4], data[0].Raw[:4])
}

func TestReadRejectsTooManyRecords(t *testing.T) {
	_, err := Read(bytes.NewReader(make([]byte, RecordSize*(MaxRecords+1))))
	require.Error(t, err)
}
