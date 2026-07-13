package tyrnpcregenfile

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRoundTripCompactAndFull(t *testing.T) {
	for _, size := range []int{CompactRecordSize, FullRecordSize} {
		d, e := Read(bytes.NewReader(make([]byte, size)))
		require.NoError(t, e)
		require.NoError(t, d.Records[0].SetField(1, 9))
		var out bytes.Buffer
		require.NoError(t, Write(&out, d))
		require.Len(t, out.Bytes(), size)
	}
}
func TestReadRejectsMisalignment(t *testing.T) {
	_, e := Read(bytes.NewReader(make([]byte, 11)))
	require.Error(t, e)
}
