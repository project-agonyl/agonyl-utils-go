package questexfile

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRoundTripAndKnownFieldUpdate(t *testing.T) {
	var d Data
	require.NoError(t, d.SetInt32(0, 42))
	require.NoError(t, d.SetContentString(0, 0, "prompt"))
	var encoded bytes.Buffer
	require.NoError(t, Write(&encoded, d))
	parsed, e := Read(bytes.NewReader(encoded.Bytes()))
	require.NoError(t, e)
	var roundTrip bytes.Buffer
	require.NoError(t, Write(&roundTrip, parsed))
	require.Equal(t, encoded.Bytes(), roundTrip.Bytes())
	require.NoError(t, parsed.SetInt32(4, 1000))
	require.Equal(t, int32(1000), parsed.Int32(4))
}
func TestReadRejectsTruncation(t *testing.T) {
	_, e := Read(bytes.NewReader([]byte{1}))
	require.Error(t, e)
}
