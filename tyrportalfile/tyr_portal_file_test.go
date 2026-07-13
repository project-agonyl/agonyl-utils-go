package tyrportalfile

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRoundTripAndFieldUpdate(t *testing.T) {
	d, e := Read(bytes.NewReader(make([]byte, RecordSize)))
	require.NoError(t, e)
	require.NoError(t, d[0].SetField(2, 9))
	var out bytes.Buffer
	require.NoError(t, Write(&out, d))
	require.Equal(t, []byte{9, 0}, out.Bytes()[4:6])
}
func TestReadRejectsDuplicateSource(t *testing.T) {
	_, e := Read(bytes.NewReader(make([]byte, RecordSize*2)))
	require.Error(t, e)
}
