package lotteryfile

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
	require.NoError(t, data[0].SetMessage(2, "winner"))
	offset := 4 + 2*MessageSize
	require.Equal(t, original[:offset], data[0].Raw[:offset])
	require.Equal(t, original[offset+MessageSize:], data[0].Raw[offset+MessageSize:])
}

func TestSetMessageRejectsOverflow(t *testing.T) {
	var record Record
	require.Error(t, record.SetMessage(0, string(make([]byte, MessageSize+1))))
}
