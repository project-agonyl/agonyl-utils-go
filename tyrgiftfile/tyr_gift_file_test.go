package tyrgiftfile

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRoundTripAndFieldUpdate(t *testing.T) {
	d, e := Read(bytes.NewReader(make([]byte, RecordSize)))
	require.NoError(t, e)
	d[0].SetWeight(7)
	var out bytes.Buffer
	require.NoError(t, Write(&out, d))
	require.Equal(t, []byte{7, 0, 0, 0}, out.Bytes()[8:12])
}
func TestReadRejectsTooManyRows(t *testing.T) {
	_, e := Read(bytes.NewReader(make([]byte, RecordSize*(MaxRecords+1))))
	require.Error(t, e)
}
