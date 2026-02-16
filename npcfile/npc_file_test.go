package npcfile

import (
	"bytes"
	"encoding/binary"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRead_ValidRecord(t *testing.T) {
	data := makeNPCWithName("Guard")
	data.Id = 42
	data.Level = 5
	data.HP = 1000

	var buf bytes.Buffer
	require.NoError(t, Write(&buf, data))

	read, err := Read(&buf)
	require.NoError(t, err)
	assert.Equal(t, data, read)
	assert.Equal(t, "Guard", read.GetName())
	assert.Equal(t, uint16(42), read.Id)
	assert.Equal(t, byte(5), read.Level)
	assert.Equal(t, uint32(1000), read.HP)
}

func TestRead_Truncated(t *testing.T) {
	// NPC record is fixed-size; need at least binary.Size(NPCFileData{}) bytes
	size := binary.Size(NPCFileData{})
	truncated := bytes.Repeat([]byte{0xFF}, size-1)
	buf := bytes.NewBuffer(truncated)

	_, err := Read(buf)
	assert.Error(t, err)
}

func TestRead_EmptyStream(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	_, err := Read(buf)
	assert.Error(t, err)
}

func TestRead_InvalidReader(t *testing.T) {
	var r errReader
	_, err := Read(&r)
	assert.Error(t, err)
}

type errReader struct{}

func (errReader) Read([]byte) (int, error) {
	return 0, io.ErrClosedPipe
}

func TestWrite_InvalidWriter(t *testing.T) {
	var w errWriter
	err := Write(&w, NPCFileData{})
	assert.Error(t, err)
}

type errWriter struct{}

func (errWriter) Write([]byte) (int, error) {
	return 0, io.ErrShortWrite
}

func TestWriteThenRead_RoundTrip(t *testing.T) {
	original := NPCFileData{
		Id:                  100,
		RespawnRate:         30,
		AttackTypeInfo:      1,
		TargetSelectionInfo: 2,
		Defense:             10,
		AdditionalDefense:   5,
		AttackSpeedLow:      100,
		AttackSpeedHigh:     200,
		MovementSpeed:       150,
		Level:               10,
		PlayerExp:           500,
		Appearance:          3,
		HP:                  5000,
		BlueAttackDefense:   20,
		RedAttackDefense:    25,
		GreyAttackDefense:   15,
		MercenaryExp:        100,
		Unknown:             0xAB,
	}
	copy(original.Name[:], "Boss\x00\x00")
	original.Attacks[0] = NPCAttack{Range: 50, Area: 10, Damage: 100, AdditionalDamage: 20}
	original.Attacks[1] = NPCAttack{Range: 100, Area: 25, Damage: 200, AdditionalDamage: 50}

	var buf bytes.Buffer
	require.NoError(t, Write(&buf, original))

	read, err := Read(&buf)
	require.NoError(t, err)
	assert.Equal(t, original, read)
	assert.Equal(t, "Boss", read.GetName())
	assert.Equal(t, original.Attacks[0], read.Attacks[0])
	assert.Equal(t, original.Attacks[1], read.Attacks[1])
}

func TestWriteThenRead_ZeroValue(t *testing.T) {
	var data NPCFileData

	var buf bytes.Buffer
	require.NoError(t, Write(&buf, data))

	read, err := Read(&buf)
	require.NoError(t, err)
	assert.Equal(t, data, read)
	assert.Equal(t, "", read.GetName())
}

func TestWrite_DeterministicSize(t *testing.T) {
	expectedSize := binary.Size(NPCFileData{})
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, NPCFileData{}))
	assert.Equal(t, expectedSize, buf.Len(), "written bytes must match struct size")
}

func TestGetName_EmptyName(t *testing.T) {
	var n NPCFileData
	assert.Equal(t, "", n.GetName())
}

func TestGetName_ShortName(t *testing.T) {
	n := NPCFileData{}
	copy(n.Name[:], "Merchant")
	assert.Equal(t, "Merchant", n.GetName())
}

func TestGetName_NullPadded(t *testing.T) {
	n := NPCFileData{}
	copy(n.Name[:], "A\x00\x00\x00")
	assert.Equal(t, "A", n.GetName())
}

func TestGetName_FullLengthNoNull(t *testing.T) {
	var n NPCFileData
	for i := range 0x14 {
		n.Name[i] = 'X'
	}
	name := n.GetName()
	assert.Len(t, name, 0x14)
	assert.Equal(t, "XXXXXXXXXXXXXXXXXXXX", name)
}

func TestGetName_ExactBufferSize(t *testing.T) {
	// Name is 0x14 (20) bytes; use exactly 20 printable chars
	n := NPCFileData{}
	for i := range 0x14 {
		n.Name[i] = byte('A' + (i % 26))
	}
	name := n.GetName()
	assert.Len(t, name, 0x14)
	assert.Equal(t, "ABCDEFGHIJKLMNOPQRST", name)
}

func TestRead_LittleEndian(t *testing.T) {
	// Manually build a minimal record with known byte order
	// Id=0x0102 (LE: 02 01), RespawnRate=0x0304 (LE: 04 03), etc.
	var buf bytes.Buffer
	name := [0x14]byte{}
	copy(name[:], "LE")
	_ = binary.Write(&buf, binary.LittleEndian, name)
	_ = binary.Write(&buf, binary.LittleEndian, uint16(0x1234))
	_ = binary.Write(&buf, binary.LittleEndian, uint16(0x5678))
	// Rest of struct: AttackTypeInfo, TargetSelectionInfo, Defense, AdditionalDefense
	buf.Write([]byte{1, 2, 3, 4})
	// Attacks[3]
	for i := 0; i < 3; i++ {
		_ = binary.Write(&buf, binary.LittleEndian, uint16(0))
		_ = binary.Write(&buf, binary.LittleEndian, uint16(0))
		_ = binary.Write(&buf, binary.LittleEndian, uint16(0))
		_ = binary.Write(&buf, binary.LittleEndian, uint16(0))
	}
	_ = binary.Write(&buf, binary.LittleEndian, uint16(0))
	_ = binary.Write(&buf, binary.LittleEndian, uint16(0))
	_ = binary.Write(&buf, binary.LittleEndian, uint32(0))
	buf.WriteByte(0)
	_ = binary.Write(&buf, binary.LittleEndian, uint16(0))
	buf.WriteByte(0)
	_ = binary.Write(&buf, binary.LittleEndian, uint32(0))
	_ = binary.Write(&buf, binary.LittleEndian, uint16(0))
	_ = binary.Write(&buf, binary.LittleEndian, uint16(0))
	_ = binary.Write(&buf, binary.LittleEndian, uint16(0))
	_ = binary.Write(&buf, binary.LittleEndian, uint16(0))
	_ = binary.Write(&buf, binary.LittleEndian, uint16(0))

	data, err := Read(&buf)
	require.NoError(t, err)
	assert.Equal(t, uint16(0x1234), data.Id)
	assert.Equal(t, uint16(0x5678), data.RespawnRate)
	assert.Equal(t, "LE", data.GetName())
}

func makeNPCWithName(name string) NPCFileData {
	var n NPCFileData
	copy(n.Name[:], name)
	return n
}
