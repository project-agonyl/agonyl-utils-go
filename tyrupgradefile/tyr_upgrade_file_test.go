package tyrupgradefile

import (
	"bytes"
	"encoding/binary"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRoundTripAndFieldUpdate(t *testing.T) {
	raw := make([]byte, RecordSize)
	binary.LittleEndian.PutUint16(raw[0:2], 1)
	binary.LittleEndian.PutUint16(raw[2:4], 1)
	d, e := Read(bytes.NewReader(raw))
	require.NoError(t, e)
	require.NoError(t, d[0].SetField(4, 9))
	var out bytes.Buffer
	require.NoError(t, Write(&out, d))
	require.Equal(t, []byte{9, 0}, out.Bytes()[8:10])
}
func TestReadRejectsInvalidCode(t *testing.T) {
	_, e := Read(bytes.NewReader(make([]byte, RecordSize)))
	require.Error(t, e)
}
