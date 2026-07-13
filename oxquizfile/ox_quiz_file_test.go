package oxquizfile

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRoundTripAndTrailingBytes(t *testing.T) {
	raw := make([]byte, RecordSize+2)
	raw[4] = 'O'
	raw[RecordSize] = 9
	d, e := Read(bytes.NewReader(raw))
	require.NoError(t, e)
	require.NoError(t, d.Records[0].SetQuestion("question"))
	var out bytes.Buffer
	require.NoError(t, Write(&out, d))
	require.Equal(t, raw[RecordSize:], out.Bytes()[RecordSize:])
}
func TestReadRejectsAnswer(t *testing.T) {
	_, e := Read(bytes.NewReader(make([]byte, RecordSize)))
	require.Error(t, e)
}
