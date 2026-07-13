package setitemfile

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRoundTripAndBonusUpdate(t *testing.T) {
	original := make([]byte, FileSize)
	for i := range original {
		original[i] = byte(i)
	}
	data, err := Read(bytes.NewReader(original))
	require.NoError(t, err)
	var encoded bytes.Buffer
	require.NoError(t, Write(&encoded, data))
	require.Equal(t, original, encoded.Bytes())
	before := data[0].Raw
	require.NoError(t, data[0].SetBonus(2, 3, 7, 0xbeef))
	offset := 0x1e + 0x1e + 9
	require.Equal(t, before[:offset], data[0].Raw[:offset])
	require.Equal(t, []byte{7, 0xef, 0xbe}, data[0].Raw[offset:offset+3])
	require.Equal(t, before[offset+3:], data[0].Raw[offset+3:])
}

func TestReadRequiresExactSize(t *testing.T) {
	_, err := Read(bytes.NewReader(make([]byte, FileSize-1)))
	require.Error(t, err)
}
