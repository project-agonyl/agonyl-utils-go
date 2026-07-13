package squestquizfile

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRoundTripAndKnownFieldUpdate(t *testing.T) {
	raw := make([]byte, FileSize)
	d, e := Read(bytes.NewReader(raw))
	require.NoError(t, e)
	require.NoError(t, d[0].SetAnswer(2, "answer"))
	before := d[0].Raw
	d[0].SetID(7)
	require.Equal(t, before[2:], d[0].Raw[2:])
	var out bytes.Buffer
	require.NoError(t, Write(&out, d))
	require.Len(t, out.Bytes(), FileSize)
}
func TestReadRequiresExactSize(t *testing.T) {
	_, e := Read(bytes.NewReader(make([]byte, FileSize-1)))
	require.Error(t, e)
}
