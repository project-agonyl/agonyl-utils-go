package hiredsoldierskillfile

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRoundTripAndKnownFieldUpdate(t *testing.T) {
	original := make([]byte, RecordSize)
	for index := range original {
		original[index] = byte(index)
	}

	data, err := Read(bytes.NewReader(original))
	require.NoError(t, err)
	var encoded bytes.Buffer
	require.NoError(t, Write(&encoded, data))
	require.Equal(t, original, encoded.Bytes())

	before := data[0].Raw
	require.NoError(t, data[0].SetLevel(2, Level{RequiredItemCode: 42, SkillPointCost: 3, MoneyCost: 4, LoreCost: 5}))
	offset := LevelStride
	require.Equal(t, before[:offset+0x0a], data[0].Raw[:offset+0x0a])
	require.Equal(t, before[offset+0x15:], data[0].Raw[offset+0x15:])
}

func TestReadRejectsMisalignedAndOversizedData(t *testing.T) {
	_, err := Read(bytes.NewReader(make([]byte, RecordSize-1)))
	require.Error(t, err)
	_, err = Read(bytes.NewReader(make([]byte, (MaxRecordCount+1)*RecordSize)))
	require.Error(t, err)
}
