package towertreasurefile

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRoundTripPreservesPartialTail(t *testing.T) {
	raw := make([]byte, RecordSize+1)
	raw[RecordSize] = 9
	d, e := Read(bytes.NewReader(raw))
	require.NoError(t, e)
	d.Records[0].SetWeight(25)
	var out bytes.Buffer
	require.NoError(t, Write(&out, d))
	require.Equal(t, byte(9), out.Bytes()[RecordSize])
}
func TestReadPreservesRowsBeyondRuntimeCapAsOpaqueTail(t *testing.T) {
	raw := make([]byte, RecordSize*(MaxRecords+1))
	raw[len(raw)-1] = 9
	data, err := Read(bytes.NewReader(raw))
	require.NoError(t, err)
	require.Len(t, data.Records, MaxRecords)
	require.Equal(t, raw[MaxRecords*RecordSize:], data.Trailing)
	var encoded bytes.Buffer
	require.NoError(t, Write(&encoded, data))
	require.Equal(t, raw, encoded.Bytes())
}
