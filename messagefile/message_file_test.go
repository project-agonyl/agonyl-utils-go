package messagefile

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRoundTripAndFieldUpdate(t *testing.T) {
	raw := []byte{1, 0, 7, 0, 10, 0, 3, 0, 'a', 'b', 'c'}
	d, e := Read(bytes.NewReader(raw))
	require.NoError(t, e)
	var out bytes.Buffer
	require.NoError(t, Write(&out, d))
	require.Equal(t, raw, out.Bytes())
	d[0].SetIndex(9)
	require.Equal(t, uint16(9), d[0].Index())
}
func TestReadRejectsTruncation(t *testing.T) {
	_, e := Read(bytes.NewReader([]byte{1, 0}))
	require.Error(t, e)
}
