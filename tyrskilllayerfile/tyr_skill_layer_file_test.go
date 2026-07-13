package tyrskilllayerfile

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRoundTripAndFlagMask(t *testing.T) {
	d, e := Read(bytes.NewReader(make([]byte, RecordSize)))
	require.NoError(t, e)
	before := d[0].Raw
	require.NoError(t, d[0].SetTargetToEffect(2, true))
	require.Equal(t, before[:0x12], d[0].Raw[:0x12])
	require.Equal(t, []byte{1, 0}, d[0].Raw[0x12:0x14])
	require.Equal(t, before[0x14:], d[0].Raw[0x14:])
	var out bytes.Buffer
	require.NoError(t, Write(&out, d))
}
func TestReadRejectsMisalignment(t *testing.T) {
	_, e := Read(bytes.NewReader(make([]byte, RecordSize+1)))
	require.Error(t, e)
}
