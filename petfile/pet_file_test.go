package petfile

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRoundTripAndNameUpdate(t *testing.T) {
	original := make([]byte, RecordSize)
	original[2], original[3] = 0xe8, 0x03
	copy(original[4:36], "Shue")
	data, err := Read(bytes.NewReader(original))
	require.NoError(t, err)
	var encoded bytes.Buffer
	require.NoError(t, Write(&encoded, data))
	require.Equal(t, original, encoded.Bytes())
	require.NoError(t, data[0].SetName("Special Shue"))
	require.Equal(t, original[:4], data[0].Raw[:4])
	require.Equal(t, original[36:], data[0].Raw[36:])
}

func TestReadRejectsInvalidCodeAndOverflow(t *testing.T) {
	invalid := make([]byte, RecordSize)
	invalid[0] = 1
	_, err := Read(bytes.NewReader(invalid))
	require.Error(t, err)
	_, err = Read(bytes.NewReader(make([]byte, RecordSize*(MaxRecords+1))))
	require.Error(t, err)
}
