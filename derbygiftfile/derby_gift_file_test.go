package derbygiftfile

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRoundTripAndWeightUpdate(t *testing.T) {
	original := make([]byte, RecordSize)
	data, err := Read(bytes.NewReader(original))
	require.NoError(t, err)
	data[0].SetWeight(25)
	var encoded bytes.Buffer
	require.NoError(t, Write(&encoded, data))
	require.Equal(t, append(make([]byte, 8), 25, 0, 0, 0), encoded.Bytes())
}

func TestReadRejectsCumulativeOverflow(t *testing.T) {
	var first, second Record
	first.SetWeight(MaxProbability)
	second.SetWeight(1)
	var encoded bytes.Buffer
	encoded.Write(first.Raw[:])
	encoded.Write(second.Raw[:])
	_, err := Read(bytes.NewReader(encoded.Bytes()))
	require.Error(t, err)
}
