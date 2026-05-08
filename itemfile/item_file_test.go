package itemfile

import (
	"bytes"
	"encoding/binary"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRawRecordSizes(t *testing.T) {
	assert.Equal(t, 242, binary.Size(IT0Raw{}))
	assert.Equal(t, 92, binary.Size(IT0ExRaw{}))
	assert.Equal(t, 52, binary.Size(IT1Raw{}))
	assert.Equal(t, 48, binary.Size(IT2Raw{}))
	assert.Equal(t, 48, binary.Size(IT3Raw{}))
}

func TestReadRawFiles_Empty(t *testing.T) {
	it0, err := ReadIT0(bytes.NewReader(nil))
	require.NoError(t, err)
	assert.NotNil(t, it0)
	assert.Empty(t, it0)

	it0Ex, err := ReadIT0Ex(bytes.NewReader(nil))
	require.NoError(t, err)
	assert.NotNil(t, it0Ex)
	assert.Empty(t, it0Ex)

	it1, err := ReadIT1(bytes.NewReader(nil))
	require.NoError(t, err)
	assert.NotNil(t, it1)
	assert.Empty(t, it1)

	it2, err := ReadIT2(bytes.NewReader(nil))
	require.NoError(t, err)
	assert.NotNil(t, it2)
	assert.Empty(t, it2)

	it3, err := ReadIT3(bytes.NewReader(nil))
	require.NoError(t, err)
	assert.NotNil(t, it3)
	assert.Empty(t, it3)
}

func TestReadWriteIT0_RoundTrip(t *testing.T) {
	data := IT0File{
		{
			ItemCodeBase: 1,
			Row:          0,
			Slot:         2,
			Type:         3,
			NPCPrice:     400,
			Unknown2:     [9]uint16{1, 2, 3},
		},
		{
			ItemCodeBase: 2,
			Row:          1,
			Slot:         4,
			Type:         5,
			NPCPrice:     600,
		},
	}
	copy(data[0].Name[:], "Sword")
	data[0].Levels[0] = IT0RawLevelProperties{
		AdditionalAttribute: 1,
		Strength:            2,
		Dexterity:           3,
		Intelligence:        4,
		Attribute:           5,
		Range:               6,
		BlueOption:          7,
		RedOption:           8,
		GreyOption:          9,
	}
	copy(data[1].Name[:], "Shield")

	var buf bytes.Buffer
	require.NoError(t, WriteIT0(&buf, data))
	assert.Equal(t, 2*binary.Size(IT0Raw{}), buf.Len())

	read, err := ReadIT0(&buf)
	require.NoError(t, err)
	assert.Equal(t, data, read)
	assert.Equal(t, "Sword", read[0].GetName())
	assert.Equal(t, "Shield", read[1].GetName())
}

func TestReadWriteIT0Ex_RoundTrip(t *testing.T) {
	data := IT0ExFile{{Row: 1}, {Row: 2}}
	data[0].Levels[0].Strength = 11
	data[1].Levels[4].GreyOption = 22

	var buf bytes.Buffer
	require.NoError(t, WriteIT0Ex(&buf, data))

	read, err := ReadIT0Ex(&buf)
	require.NoError(t, err)
	assert.Equal(t, data, read)
}

func TestReadWriteIT1_RoundTrip(t *testing.T) {
	data := IT1File{
		{Type: 4, Row: 0, NPCPrice: 100, Unknown1: 1, RequiredLevel: 10, Attribute: 20, BlueOption: 30, RedOption: 40, GreyOption: 50},
		{Type: 5, Row: 1, NPCPrice: 200, Unknown1: 2, RequiredLevel: 11, Attribute: 21, BlueOption: 31, RedOption: 41, GreyOption: 51},
	}
	copy(data[0].Name[:], "Potion")
	copy(data[1].Name[:], "Charm")

	var buf bytes.Buffer
	require.NoError(t, WriteIT1(&buf, data))
	assert.Equal(t, 2*binary.Size(IT1Raw{}), buf.Len())

	read, err := ReadIT1(&buf)
	require.NoError(t, err)
	assert.Equal(t, data, read)
	assert.Equal(t, "Potion", read[0].GetName())
	assert.Equal(t, "Charm", read[1].GetName())
}

func TestReadWriteIT2_RoundTrip(t *testing.T) {
	data := IT2File{
		{Type: 6, Row: 0, NPCPrice: 100, Class: 1, RequiredLevel: 2, Unknown1: 3, SkillLevel: 4},
		{Type: 7, Row: 1, NPCPrice: 200, Class: 5, RequiredLevel: 6, Unknown1: 7, SkillLevel: 8},
	}
	copy(data[0].Name[:], "Skill One")
	copy(data[1].Name[:], "Skill Two")

	var buf bytes.Buffer
	require.NoError(t, WriteIT2(&buf, data))

	read, err := ReadIT2(&buf)
	require.NoError(t, err)
	assert.Equal(t, data, read)
	assert.Equal(t, "Skill One", read[0].GetName())
	assert.Equal(t, "Skill Two", read[1].GetName())
}

func TestReadWriteIT3_RoundTrip(t *testing.T) {
	data := IT3File{
		{Type: 8, Row: 0, NPCPrice: 100, Unknown1: 1, Unknown2: 2, Unknown3: 3, Unknown4: 4},
		{Type: 9, Row: 1, NPCPrice: 200, Unknown1: 5, Unknown2: 6, Unknown3: 7, Unknown4: 8},
	}
	copy(data[0].Name[:], "Misc One")
	copy(data[1].Name[:], "Misc Two")

	var buf bytes.Buffer
	require.NoError(t, WriteIT3(&buf, data))

	read, err := ReadIT3(&buf)
	require.NoError(t, err)
	assert.Equal(t, data, read)
	assert.Equal(t, "Misc One", read[0].GetName())
	assert.Equal(t, "Misc Two", read[1].GetName())
}

func TestRawReaders_TruncatedInput(t *testing.T) {
	readers := []struct {
		name string
		read func(io.Reader) error
	}{
		{name: "IT0", read: func(r io.Reader) error { _, err := ReadIT0(r); return err }},
		{name: "IT0Ex", read: func(r io.Reader) error { _, err := ReadIT0Ex(r); return err }},
		{name: "IT1", read: func(r io.Reader) error { _, err := ReadIT1(r); return err }},
		{name: "IT2", read: func(r io.Reader) error { _, err := ReadIT2(r); return err }},
		{name: "IT3", read: func(r io.Reader) error { _, err := ReadIT3(r); return err }},
	}

	for _, tc := range readers {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.read(bytes.NewReader([]byte{0x01, 0x02, 0x03}))
			assert.ErrorIs(t, err, io.ErrUnexpectedEOF)
		})
	}
}

func TestRawReaders_InvalidReader(t *testing.T) {
	_, err := ReadIT0(errReader{})
	assert.Error(t, err)
}

func TestRawWriters_EmptyAndInvalidWriter(t *testing.T) {
	var buf bytes.Buffer
	require.NoError(t, WriteIT0(&buf, nil))
	assert.Empty(t, buf.Bytes())

	require.NoError(t, WriteIT1(&buf, IT1File{}))
	assert.Empty(t, buf.Bytes())

	assert.Error(t, WriteIT3(errWriter{}, IT3File{{Type: 1}}))
}

func TestParseIT0Items_WithExtensionLevels(t *testing.T) {
	it0 := IT0File{
		{
			ItemCodeBase: 2,
			Row:          0,
			Slot:         3,
			Type:         4,
			NPCPrice:     500,
		},
	}
	copy(it0[0].Name[:], "Blade")
	it0[0].Levels[0] = IT0RawLevelProperties{
		AdditionalAttribute: 1,
		Strength:            2,
		Dexterity:           3,
		Intelligence:        4,
		Attribute:           5,
		Range:               6,
		BlueOption:          7,
		RedOption:           8,
		GreyOption:          9,
	}

	it0Ex := IT0ExFile{{Row: 0}}
	it0Ex[0].Levels[0] = IT0RawLevelProperties{
		AdditionalAttribute: 10,
		Strength:            11,
		Dexterity:           12,
		Intelligence:        13,
		Attribute:           14,
		Range:               15,
		BlueOption:          16,
		RedOption:           17,
		GreyOption:          18,
	}

	items, err := ParseIT0Items(it0, it0Ex)
	require.NoError(t, err)
	require.Len(t, items, 1)

	item := items[0]
	assert.Equal(t, uint32((2<<10)+0), item.ItemCode)
	assert.Equal(t, byte(3), item.SlotIndex)
	assert.Equal(t, "Blade", item.ItemName)
	assert.Equal(t, byte(4), item.Itemtype)
	assert.Equal(t, uint32(500), item.NPCPrice)
	require.NotNil(t, item.IT0Property)
	require.Len(t, item.IT0Property.Levels, 15)

	base := item.IT0Property.Levels[0]
	assert.Equal(t, byte(1), base.Level)
	assert.Equal(t, uint16(6), base.AttributeRange)
	assert.Equal(t, uint16(5), base.Attribute)
	assert.Equal(t, uint16(2), base.Strength)
	assert.Equal(t, uint16(4), base.Intelligence)
	assert.Equal(t, uint16(3), base.Dexterity)
	assert.Equal(t, uint16(1), base.AdditionalAttribute)
	assert.Equal(t, uint16(8), base.RedOption)
	assert.Equal(t, uint16(9), base.GreyOption)
	assert.Equal(t, uint16(7), base.BlueOption)

	extra := item.IT0Property.Levels[10]
	assert.Equal(t, byte(1), extra.Level)
	assert.Equal(t, uint16(15), extra.AttributeRange)
	assert.Equal(t, uint16(14), extra.Attribute)
	assert.Equal(t, uint16(11), extra.Strength)
	assert.Equal(t, uint16(13), extra.Intelligence)
	assert.Equal(t, uint16(12), extra.Dexterity)
	assert.Equal(t, uint16(10), extra.AdditionalAttribute)
	assert.Equal(t, uint16(17), extra.RedOption)
	assert.Equal(t, uint16(18), extra.GreyOption)
	assert.Equal(t, uint16(16), extra.BlueOption)
}

func TestParseIT1Items(t *testing.T) {
	it1 := IT1File{
		{Type: 4, Row: 0, NPCPrice: 100, RequiredLevel: 10, Attribute: 20, BlueOption: 30, RedOption: 40, GreyOption: 50},
		{Type: 5, Row: 1, NPCPrice: 200, RequiredLevel: 11, Attribute: 21, BlueOption: 31, RedOption: 41, GreyOption: 51},
	}
	copy(it1[0].Name[:], "Potion")
	copy(it1[1].Name[:], "Scroll")

	items, err := ParseIT1Items(it1)
	require.NoError(t, err)
	require.Len(t, items, 2)

	assert.Equal(t, uint32((4<<10)+0), items[0].ItemCode)
	assert.Equal(t, byte(8), items[0].SlotIndex)
	assert.Equal(t, "Potion", items[0].ItemName)
	assert.Equal(t, byte(4), items[0].Itemtype)
	assert.Equal(t, uint32(100), items[0].NPCPrice)
	require.NotNil(t, items[0].IT1Property)
	assert.Equal(t, uint16(10), items[0].IT1Property.RequiredLevel)
	assert.Equal(t, uint16(20), items[0].IT1Property.Attribute)
	assert.Equal(t, uint16(40), items[0].IT1Property.RedOption)
	assert.Equal(t, uint16(50), items[0].IT1Property.GreyOption)
	assert.Equal(t, uint16(30), items[0].IT1Property.BlueOption)

	assert.Equal(t, uint32((5<<10)+1), items[1].ItemCode)
	assert.Equal(t, byte(9), items[1].SlotIndex)
	assert.Equal(t, "Scroll", items[1].ItemName)
}

func TestParseIT2Items(t *testing.T) {
	it2 := IT2File{{Type: 6, Row: 0, NPCPrice: 300, Class: 1, RequiredLevel: 2, SkillLevel: 3}}
	copy(it2[0].Name[:], "Skill")

	items, err := ParseIT2Items(it2)
	require.NoError(t, err)
	require.Len(t, items, 1)

	assert.Equal(t, uint32((6<<10)+0), items[0].ItemCode)
	assert.Equal(t, "Skill", items[0].ItemName)
	assert.Equal(t, byte(6), items[0].Itemtype)
	assert.Equal(t, uint32(300), items[0].NPCPrice)
	require.NotNil(t, items[0].IT2Property)
	assert.Equal(t, uint16(1), items[0].IT2Property.Class)
	assert.Equal(t, uint16(2), items[0].IT2Property.RequiredLevel)
	assert.Equal(t, uint16(3), items[0].IT2Property.SkillLevel)
}

func TestParseIT3Items(t *testing.T) {
	it3 := IT3File{{Type: 7, Row: 0, NPCPrice: 400}}
	copy(it3[0].Name[:], "Misc")

	items, err := ParseIT3Items(it3)
	require.NoError(t, err)
	require.Len(t, items, 1)

	assert.Equal(t, uint32((7<<10)+0), items[0].ItemCode)
	assert.Equal(t, "Misc", items[0].ItemName)
	assert.Equal(t, byte(7), items[0].Itemtype)
	assert.Equal(t, uint32(400), items[0].NPCPrice)
	assert.Nil(t, items[0].IT0Property)
	assert.Nil(t, items[0].IT1Property)
	assert.Nil(t, items[0].IT2Property)
}

func TestParseItems_BadRows(t *testing.T) {
	_, err := ParseIT0Items(IT0File{{Row: 1}}, nil)
	assert.Error(t, err)

	_, err = ParseIT0Items(IT0File{{Row: 0}, {Row: 1}}, IT0ExFile{{Row: 2}})
	assert.Error(t, err)

	_, err = ParseIT0Items(IT0File{{Row: 1}, {Row: 1}}, IT0ExFile{{Row: 0}})
	assert.Error(t, err)

	_, err = ParseIT1Items(IT1File{{Row: 1}})
	assert.Error(t, err)

	_, err = ParseIT2Items(IT2File{{Row: 1}})
	assert.Error(t, err)

	_, err = ParseIT3Items(IT3File{{Row: 1}})
	assert.Error(t, err)
}

func TestReadItemHelpers(t *testing.T) {
	it0 := IT0File{{Row: 0}}
	copy(it0[0].Name[:], "Blade")
	var it0Buf bytes.Buffer
	require.NoError(t, WriteIT0(&it0Buf, it0))
	var it0ExBuf bytes.Buffer
	require.NoError(t, WriteIT0Ex(&it0ExBuf, nil))

	items, err := ReadIT0Items(&it0Buf, &it0ExBuf)
	require.NoError(t, err)
	require.Len(t, items, 1)
	assert.Equal(t, "Blade", items[0].ItemName)

	it1 := IT1File{{Row: 0}}
	copy(it1[0].Name[:], "Potion")
	var it1Buf bytes.Buffer
	require.NoError(t, WriteIT1(&it1Buf, it1))

	items, err = ReadIT1Items(&it1Buf)
	require.NoError(t, err)
	require.Len(t, items, 1)
	assert.Equal(t, "Potion", items[0].ItemName)

	it2 := IT2File{{Row: 0}}
	copy(it2[0].Name[:], "Skill")
	var it2Buf bytes.Buffer
	require.NoError(t, WriteIT2(&it2Buf, it2))

	items, err = ReadIT2Items(&it2Buf)
	require.NoError(t, err)
	require.Len(t, items, 1)
	assert.Equal(t, "Skill", items[0].ItemName)

	it3 := IT3File{{Row: 0}}
	copy(it3[0].Name[:], "Misc")
	var it3Buf bytes.Buffer
	require.NoError(t, WriteIT3(&it3Buf, it3))

	items, err = ReadIT3Items(&it3Buf)
	require.NoError(t, err)
	require.Len(t, items, 1)
	assert.Equal(t, "Misc", items[0].ItemName)
}

type errReader struct{}

func (errReader) Read([]byte) (int, error) {
	return 0, io.ErrClosedPipe
}

type errWriter struct{}

func (errWriter) Write([]byte) (int, error) {
	return 0, io.ErrShortWrite
}
