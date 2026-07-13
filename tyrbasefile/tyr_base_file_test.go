package tyrbasefile

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRoundTripAndFieldUpdate(t *testing.T) {
	raw := make([]byte, RecordSize)
	d, e := Read(bytes.NewReader(raw))
	require.NoError(t, e)
	require.NoError(t, d[0].SetField(3, 7))
	var out bytes.Buffer
	require.NoError(t, Write(&out, d))
	require.Equal(t, []byte{7, 0}, out.Bytes()[6:8])
}
func TestReadRejectsMisalignment(t *testing.T) {
	_, e := Read(bytes.NewReader(make([]byte, RecordSize+1)))
	require.Error(t, e)
}
