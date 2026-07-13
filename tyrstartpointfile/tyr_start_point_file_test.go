package tyrstartpointfile

import (
	"bytes"
	"encoding/binary"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRoundTripAndFieldUpdate(t *testing.T) {
	raw := make([]byte, RecordSize)
	binary.LittleEndian.PutUint16(raw[10:12], 1)
	d, e := Read(bytes.NewReader(raw))
	require.NoError(t, e)
	require.NoError(t, d[0].SetField(3, 8))
	var out bytes.Buffer
	require.NoError(t, Write(&out, d))
	require.Equal(t, []byte{8, 0}, out.Bytes()[6:8])
}
func TestReadRejectsDirection(t *testing.T) {
	_, e := Read(bytes.NewReader(make([]byte, RecordSize)))
	require.Error(t, e)
}
