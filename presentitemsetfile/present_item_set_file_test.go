package presentitemsetfile

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRoundTripAndRewardUpdate(t *testing.T) {
	original := make([]byte, RecordSize)
	for i := range original {
		original[i] = byte(i)
	}
	data, err := Read(bytes.NewReader(original))
	require.NoError(t, err)
	var encoded bytes.Buffer
	require.NoError(t, Write(&encoded, data))
	require.Equal(t, original, encoded.Bytes())
	require.NoError(t, data[0].SetReward(1, 3, 0xbeef))
	require.Equal(t, original[:6], data[0].Raw[:6])
	require.Equal(t, []byte{3, 0, 0xef, 0xbe}, data[0].Raw[6:10])
	require.Equal(t, original[10:], data[0].Raw[10:])
}

func TestSetRewardRejectsIndex(t *testing.T) {
	var record Record
	require.Error(t, record.SetReward(RewardCount, 0, 0))
}
