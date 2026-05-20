package itemcombinationdata

import (
	"bytes"
	"encoding/binary"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRead_EmptyStream(t *testing.T) {
	data, err := Read(bytes.NewReader(nil))
	require.NoError(t, err)
	assert.NotNil(t, data)
	assert.Len(t, data, 0)
}

func TestRead_SingleFormula(t *testing.T) {
	formula := sampleFormula()
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, ItemCombinationData{formula}))

	data, err := Read(&buf)
	require.NoError(t, err)
	require.Len(t, data, 1)
	assert.Equal(t, formula, data[0])
}

func TestRead_MultipleFormulas(t *testing.T) {
	formulas := ItemCombinationData{
		sampleFormula(),
		{
			Item1: 100, Item2: 101, Item3: 102, Item4: 103, Item5: 104,
			Item6: 105, Item7: 106, Item8: 107, Item9: 108, Item10: 109,
			SuccessRate: 120,
			Outcome:     999,
			Unknown1:    1,
			Unknown2:    2,
			Unknown3:    3,
			Unknown4:    4,
		},
	}
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, formulas))

	data, err := Read(&buf)
	require.NoError(t, err)
	assert.Equal(t, formulas, data)
}

func TestRead_KnownBytesDecode(t *testing.T) {
	raw := []byte{
		0x01, 0x00, // Item1
		0x02, 0x00, // Item2
		0x03, 0x00, // Item3
		0x04, 0x00, // Item4
		0x05, 0x00, // Item5
		0x06, 0x00, // Item6
		0x07, 0x00, // Item7
		0x08, 0x00, // Item8
		0x09, 0x00, // Item9
		0x0A, 0x00, // Item10
		0x78, 0x00, // SuccessRate
		0x34, 0x12, // Outcome
		0xCD, 0xAB, // Unknown1
		0xDC, 0xBA, // Unknown2
		0xEF, 0xBE, // Unknown3
		0xFE, 0xCA, // Unknown4
	}

	data, err := Read(bytes.NewReader(raw))
	require.NoError(t, err)
	require.Len(t, data, 1)
	assert.Equal(t, CraftFormula{
		Item1:       1,
		Item2:       2,
		Item3:       3,
		Item4:       4,
		Item5:       5,
		Item6:       6,
		Item7:       7,
		Item8:       8,
		Item9:       9,
		Item10:      10,
		SuccessRate: 120,
		Outcome:     0x1234,
		Unknown1:    0xABCD,
		Unknown2:    0xBADC,
		Unknown3:    0xBEEF,
		Unknown4:    0xCAFE,
	}, data[0])
}

func TestWrite_SingleFormula_ByteLayout(t *testing.T) {
	formula := CraftFormula{
		Item1:       0x0102,
		Item2:       0x0304,
		Item3:       0x0506,
		Item4:       0x0708,
		Item5:       0x090A,
		Item6:       0x0B0C,
		Item7:       0x0D0E,
		Item8:       0x0F10,
		Item9:       0x1112,
		Item10:      0x1314,
		SuccessRate: 0x1516,
		Outcome:     0x1718,
		Unknown1:    0x191A,
		Unknown2:    0x1B1C,
		Unknown3:    0x1D1E,
		Unknown4:    0x1F20,
	}

	var buf bytes.Buffer
	require.NoError(t, Write(&buf, ItemCombinationData{formula}))

	expected := []byte{
		0x02, 0x01, // Item1
		0x04, 0x03, // Item2
		0x06, 0x05, // Item3
		0x08, 0x07, // Item4
		0x0A, 0x09, // Item5
		0x0C, 0x0B, // Item6
		0x0E, 0x0D, // Item7
		0x10, 0x0F, // Item8
		0x12, 0x11, // Item9
		0x14, 0x13, // Item10
		0x16, 0x15, // SuccessRate
		0x18, 0x17, // Outcome
		0x1A, 0x19, // Unknown1
		0x1C, 0x1B, // Unknown2
		0x1E, 0x1D, // Unknown3
		0x20, 0x1F, // Unknown4
	}
	assert.Equal(t, expected, buf.Bytes())
	assert.Equal(t, FormulaSize, buf.Len())
}

func TestWriteThenRead_RoundTrip(t *testing.T) {
	original := ItemCombinationData{
		sampleFormula(),
		{
			Item1:       0,
			Item2:       2000,
			Item10:      65535,
			SuccessRate: 1,
			Outcome:     3000,
			Unknown1:    0xAAAA,
			Unknown2:    0xBBBB,
			Unknown3:    0xCCCC,
			Unknown4:    0xDDDD,
		},
	}

	var buf bytes.Buffer
	require.NoError(t, Write(&buf, original))

	read, err := Read(&buf)
	require.NoError(t, err)
	assert.Equal(t, original, read)
}

func TestRead_TruncatedStream(t *testing.T) {
	_, err := Read(bytes.NewReader(bytes.Repeat([]byte{0}, FormulaSize-1)))
	require.Error(t, err)
	assert.ErrorIs(t, err, io.ErrUnexpectedEOF)
}

func TestRead_PartialSecondFormula(t *testing.T) {
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, ItemCombinationData{sampleFormula()}))
	buf.Write(bytes.Repeat([]byte{0}, 2))

	_, err := Read(&buf)
	require.Error(t, err)
	assert.ErrorIs(t, err, io.ErrUnexpectedEOF)
}

func TestRead_InvalidReader(t *testing.T) {
	_, err := Read(errReader{})
	assert.Error(t, err)
}

func TestWrite_Empty(t *testing.T) {
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, nil))
	assert.Empty(t, buf.Bytes())

	data, err := Read(&buf)
	require.NoError(t, err)
	assert.NotNil(t, data)
	assert.Len(t, data, 0)
}

func TestWrite_EmptySlice(t *testing.T) {
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, ItemCombinationData{}))
	assert.Empty(t, buf.Bytes())
}

func TestWrite_InvalidWriter(t *testing.T) {
	err := Write(errWriter{}, ItemCombinationData{sampleFormula()})
	assert.Error(t, err)
}

func TestCraftFormula_Size(t *testing.T) {
	assert.Equal(t, FormulaSize, binary.Size(CraftFormula{}))
}

func sampleFormula() CraftFormula {
	return CraftFormula{
		Item1:       101,
		Item2:       102,
		Item3:       103,
		Item4:       104,
		Item5:       105,
		Item6:       106,
		Item7:       107,
		Item8:       108,
		Item9:       109,
		Item10:      110,
		SuccessRate: 75,
		Outcome:     500,
		Unknown1:    0x1111,
		Unknown2:    0x2222,
		Unknown3:    0x3333,
		Unknown4:    0x4444,
	}
}

type errReader struct{}

func (errReader) Read([]byte) (int, error) {
	return 0, io.ErrClosedPipe
}

type errWriter struct{}

func (errWriter) Write([]byte) (int, error) {
	return 0, io.ErrShortWrite
}
