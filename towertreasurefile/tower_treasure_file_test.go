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
func TestReadRejectsTooManyRows(t *testing.T) {
	_, e := Read(bytes.NewReader(make([]byte, RecordSize*(MaxRecords+1))))
	require.Error(t, e)
}
