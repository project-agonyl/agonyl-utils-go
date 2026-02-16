package questfile

import (
	"bytes"
	"encoding/binary"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// --- Test helpers ---

func minimalValidQuestFile() QuestFile {
	var q QuestFile
	q.Header.SetQuestID(1)
	q.Header.SetGivenNPCID(100)
	// TargetNPCBlock stays zero
	q.Header.MinLevel = 10
	q.Header.MaxLevel = 50
	q.Header.EXP = 1000
	q.Header.Woonz = 500
	q.Header.Lore = 100
	// Unused reward slots: 0xFFFF in first 2 bytes of each slot
	binary.LittleEndian.PutUint16(q.Header.RewardSlot1[:2], UnusedRewardItemCode)
	binary.LittleEndian.PutUint16(q.Header.RewardSlot2[:2], UnusedRewardItemCode)
	binary.LittleEndian.PutUint16(q.Header.RewardSlot3[:2], UnusedRewardItemCode)
	for i := range q.Objectives {
		q.Objectives[i].Block[0] = TypeKILL
	}
	q.Continuation[0] = UnusedContinuation
	q.Continuation[1] = UnusedContinuation
	q.Continuation[2] = UnusedContinuation
	return q
}

// --- 1. Header tests ---

func TestRead_HeaderTooShort(t *testing.T) {
	buf := bytes.NewBuffer(make([]byte, 50))
	_, err := Read(buf)
	require.Error(t, err)
	assert.ErrorIs(t, err, io.ErrUnexpectedEOF)
}

func TestHeader_QuestIDParsing(t *testing.T) {
	// 0x00001234 -> QuestID = 0x1234
	var h QuestHeader
	binary.LittleEndian.PutUint32(h.QuestIDRaw[:], 0x00001234)
	assert.Equal(t, uint16(0x1234), h.QuestID())
}

func TestHeader_QuestIDPaddingPreserved(t *testing.T) {
	// Raw 0x1234ABCD: ID=0x1234, padding preserved on round-trip
	q := minimalValidQuestFile()
	q.Header.QuestIDRaw = [4]byte{0x34, 0x12, 0xCD, 0xAB}
	var buf bytes.Buffer
	err := Write(&buf, q)
	require.NoError(t, err)
	read, err := Read(&buf)
	require.NoError(t, err)
	assert.Equal(t, [4]byte{0x34, 0x12, 0xCD, 0xAB}, read.Header.QuestIDRaw)
	assert.Equal(t, uint16(0x1234), read.Header.QuestID())
}

func TestHeader_GivenNPCPaddingPreserved(t *testing.T) {
	q := minimalValidQuestFile()
	q.Header.GivenNPCRaw = [4]byte{0x78, 0x56, 0x34, 0x12}
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, q))
	read, err := Read(&buf)
	require.NoError(t, err)
	assert.Equal(t, [4]byte{0x78, 0x56, 0x34, 0x12}, read.Header.GivenNPCRaw)
}

func TestHeader_TargetNPCBlockPreserved(t *testing.T) {
	q := minimalValidQuestFile()
	for i := range q.Header.TargetNPCBlock {
		q.Header.TargetNPCBlock[i] = byte(i)
	}
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, q))
	read, err := Read(&buf)
	require.NoError(t, err)
	assert.Equal(t, q.Header.TargetNPCBlock, read.Header.TargetNPCBlock)
}

func TestHeader_MinMaxLevel(t *testing.T) {
	q := minimalValidQuestFile()
	q.Header.MinLevel = 1
	q.Header.MaxLevel = 200
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, q))
	read, err := Read(&buf)
	require.NoError(t, err)
	assert.Equal(t, uint8(1), read.Header.MinLevel)
	assert.Equal(t, uint8(200), read.Header.MaxLevel)
}

func TestHeader_MinGreaterThanMaxAllowed(t *testing.T) {
	q := minimalValidQuestFile()
	q.Header.MinLevel = 100
	q.Header.MaxLevel = 50
	var buf bytes.Buffer
	err := Write(&buf, q)
	require.NoError(t, err)
	_, err = Read(&buf)
	require.NoError(t, err)
}

func TestHeader_RewardItemsAndUnusedSlot(t *testing.T) {
	q := minimalValidQuestFile()
	binary.LittleEndian.PutUint16(q.Header.RewardSlot1[:2], 100)
	binary.LittleEndian.PutUint16(q.Header.RewardSlot2[:2], 200)
	binary.LittleEndian.PutUint16(q.Header.RewardSlot3[:2], 300)
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, q))
	read, err := Read(&buf)
	require.NoError(t, err)
	assert.Equal(t, uint16(100), binary.LittleEndian.Uint16(read.Header.RewardSlot1[:2]))
	assert.Equal(t, uint16(200), binary.LittleEndian.Uint16(read.Header.RewardSlot2[:2]))
	assert.Equal(t, uint16(300), binary.LittleEndian.Uint16(read.Header.RewardSlot3[:2]))
}

func TestHeader_UnusedSlot0xFFFF(t *testing.T) {
	q := minimalValidQuestFile()
	binary.LittleEndian.PutUint16(q.Header.RewardSlot2[:2], UnusedRewardItemCode)
	q.Header.RewardSlot2[2] = 0xAB
	q.Header.RewardSlot2[3] = 0xCD
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, q))
	read, err := Read(&buf)
	require.NoError(t, err)
	assert.Equal(t, uint16(UnusedRewardItemCode), binary.LittleEndian.Uint16(read.Header.RewardSlot2[:2]))
	assert.Equal(t, byte(0xAB), read.Header.RewardSlot2[2])
	assert.Equal(t, byte(0xCD), read.Header.RewardSlot2[3])
}

func TestHeader_FourthSlotUntouched(t *testing.T) {
	q := minimalValidQuestFile()
	q.Header.RewardSlot4Pad = [4]byte{0x11, 0x22, 0x33, 0x44}
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, q))
	read, err := Read(&buf)
	require.NoError(t, err)
	assert.Equal(t, [4]byte{0x11, 0x22, 0x33, 0x44}, read.Header.RewardSlot4Pad)
}

func TestHeader_RewardCounts(t *testing.T) {
	q := minimalValidQuestFile()
	q.Header.Count1 = 1
	q.Header.Count2 = 255
	q.Header.Count3 = 0
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, q))
	read, err := Read(&buf)
	require.NoError(t, err)
	assert.Equal(t, uint8(1), read.Header.Count1)
	assert.Equal(t, uint8(255), read.Header.Count2)
	assert.Equal(t, uint8(0), read.Header.Count3)
}

func TestHeader_RewardCountsPaddingPreserved(t *testing.T) {
	q := minimalValidQuestFile()
	q.Header.Count1Pad = [3]byte{0xA, 0xB, 0xC}
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, q))
	read, err := Read(&buf)
	require.NoError(t, err)
	assert.Equal(t, [3]byte{0xA, 0xB, 0xC}, read.Header.Count1Pad)
}

func TestHeader_NumericRewardsMaxUint32(t *testing.T) {
	q := minimalValidQuestFile()
	q.Header.EXP = 0xFFFFFFFF
	q.Header.Woonz = 0xFFFFFFFF
	q.Header.Lore = 0xFFFFFFFF
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, q))
	read, err := Read(&buf)
	require.NoError(t, err)
	assert.Equal(t, uint32(0xFFFFFFFF), read.Header.EXP)
	assert.Equal(t, uint32(0xFFFFFFFF), read.Header.Woonz)
	assert.Equal(t, uint32(0xFFFFFFFF), read.Header.Lore)
}

func TestHeader_TailPaddingPreserved(t *testing.T) {
	q := minimalValidQuestFile()
	q.Header.HeaderTail = [4]byte{0xDD, 0xEE, 0xFF, 0x00}
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, q))
	read, err := Read(&buf)
	require.NoError(t, err)
	assert.Equal(t, [4]byte{0xDD, 0xEE, 0xFF, 0x00}, read.Header.HeaderTail)
}

// --- 2. Objectives tests ---

func TestRead_IncompleteObjectiveBlock(t *testing.T) {
	q := minimalValidQuestFile()
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, q))
	raw := buf.Bytes()
	// Truncate after header + 1 full objective + 50 bytes of second (incomplete)
	truncated := raw[:HeaderSize+ObjectiveBlockSize+50]
	_, err := Read(bytes.NewReader(truncated))
	require.Error(t, err)
	assert.ErrorIs(t, err, io.ErrUnexpectedEOF)
}

func TestRead_ValidObjectiveTypes0to4(t *testing.T) {
	q := minimalValidQuestFile()
	q.Objectives[0].Block[0] = TypeKILL
	q.Objectives[1].Block[0] = TypeQUESTITEM
	q.Objectives[2].Block[0] = TypeBRINGNPC
	q.Objectives[3].Block[0] = TypeDROP
	q.Objectives[4].Block[0] = TypeFIND
	q.Objectives[5].Block[0] = TypeKILL
	q.Objectives[6].Block[0] = TypeKILL
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, q))
	read, err := Read(&buf)
	require.NoError(t, err)
	assert.Equal(t, uint8(TypeKILL), read.Objectives[0].ObjectiveType())
	assert.Equal(t, uint8(TypeQUESTITEM), read.Objectives[1].ObjectiveType())
	assert.Equal(t, uint8(TypeBRINGNPC), read.Objectives[2].ObjectiveType())
	assert.Equal(t, uint8(TypeDROP), read.Objectives[3].ObjectiveType())
	assert.Equal(t, uint8(TypeFIND), read.Objectives[4].ObjectiveType())
}

func TestRead_InvalidObjectiveType(t *testing.T) {
	q := minimalValidQuestFile()
	q.Objectives[2].Block[0] = 9
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, q))
	_, err := Read(&buf)
	require.Error(t, err)
	assert.ErrorIs(t, err, ErrInvalidObjectiveType)
}

func TestRead_ObjectiveFieldsParsed(t *testing.T) {
	q := minimalValidQuestFile()
	// Objective 0: set MapID, LocationID, Radius, MonsterID, KillCount, etc.
	binary.LittleEndian.PutUint16(q.Objectives[0].Block[4:6], 15)
	binary.LittleEndian.PutUint16(q.Objectives[0].Block[8:10], 100)
	q.Objectives[0].Block[12] = 25
	binary.LittleEndian.PutUint16(q.Objectives[0].Block[16:18], 3001)
	binary.LittleEndian.PutUint16(q.Objectives[0].Block[20:22], 20)
	binary.LittleEndian.PutUint16(q.Objectives[0].Block[24:26], 4500)
	binary.LittleEndian.PutUint16(q.Objectives[0].Block[56:58], 10)
	q.Objectives[0].Block[76] = 50
	q.Objectives[0].Block[80] = 25
	q.Objectives[0].Block[84] = 10
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, q))
	read, err := Read(&buf)
	require.NoError(t, err)
	assert.Equal(t, uint16(15), binary.LittleEndian.Uint16(read.Objectives[0].Block[4:6]))
	assert.Equal(t, uint16(100), binary.LittleEndian.Uint16(read.Objectives[0].Block[8:10]))
	assert.Equal(t, uint8(25), read.Objectives[0].Block[12])
	assert.Equal(t, uint16(3001), binary.LittleEndian.Uint16(read.Objectives[0].Block[16:18]))
	assert.Equal(t, uint16(20), binary.LittleEndian.Uint16(read.Objectives[0].Block[20:22]))
	assert.Equal(t, uint16(4500), binary.LittleEndian.Uint16(read.Objectives[0].Block[24:26]))
	assert.Equal(t, uint16(10), binary.LittleEndian.Uint16(read.Objectives[0].Block[56:58]))
	assert.Equal(t, uint8(50), read.Objectives[0].Block[76])
	assert.Equal(t, uint8(25), read.Objectives[0].Block[80])
	assert.Equal(t, uint8(10), read.Objectives[0].Block[84])
}

func TestRead_NameLengthZeroForKILL(t *testing.T) {
	q := minimalValidQuestFile()
	q.Objectives[0].Block[0] = TypeKILL
	q.Objectives[0].Block[92] = 0
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, q))
	_, err := Read(&buf)
	require.NoError(t, err)
}

func TestRead_NameLengthGreaterThanZeroForKILLError(t *testing.T) {
	q := minimalValidQuestFile()
	q.Objectives[0].Block[0] = TypeKILL
	q.Objectives[0].Block[92] = 5
	q.Objectives[0].Name = make([]byte, 5)
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, q))
	_, err := Read(&buf)
	require.Error(t, err)
	assert.ErrorIs(t, err, ErrNameLengthForType)
}

func TestRead_DROPWithNameLength10(t *testing.T) {
	q := minimalValidQuestFile()
	q.Objectives[0].Block[0] = TypeDROP
	q.Objectives[0].Block[92] = 10
	q.Objectives[0].Name = make([]byte, 10)
	for i := range q.Objectives[0].Name {
		q.Objectives[0].Name[i] = byte('A' + i)
	}
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, q))
	read, err := Read(&buf)
	require.NoError(t, err)
	assert.Equal(t, uint8(10), read.Objectives[0].NameLength())
	require.Len(t, read.Objectives[0].Name, 10)
	assert.Equal(t, []byte("ABCDEFGHIJ"), read.Objectives[0].Name)
}

func TestRead_Type3WithNameLengthZeroValid(t *testing.T) {
	q := minimalValidQuestFile()
	q.Objectives[0].Block[0] = TypeDROP
	q.Objectives[0].Block[92] = 0
	q.Objectives[0].Name = nil
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, q))
	read, err := Read(&buf)
	require.NoError(t, err)
	assert.Equal(t, uint8(TypeDROP), read.Objectives[0].ObjectiveType())
	assert.Len(t, read.Objectives[0].Name, 0)
}

func TestRead_TruncatedName(t *testing.T) {
	// Build raw: header + obj1 with NameLength=20 but only 10 bytes after block
	q := minimalValidQuestFile()
	q.Objectives[0].Block[0] = TypeDROP
	q.Objectives[0].Block[92] = 20
	q.Objectives[0].Name = make([]byte, 20) // we write 20, then truncate the buffer
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, q))
	raw := buf.Bytes()
	// Truncate so only 10 name bytes are present (after the first objective block)
	truncated := raw[:HeaderSize+ObjectiveBlockSize+10]
	_, err := Read(bytes.NewReader(truncated))
	require.Error(t, err)
	assert.ErrorIs(t, err, io.ErrUnexpectedEOF)
}

func TestRead_MultipleNamedObjectives(t *testing.T) {
	q := minimalValidQuestFile()
	q.Objectives[0].Block[0] = TypeDROP
	q.Objectives[0].Block[92] = 5
	q.Objectives[0].Name = []byte("AAAAA")
	q.Objectives[1].Block[0] = TypeFIND
	q.Objectives[1].Block[92] = 7
	q.Objectives[1].Name = []byte("BBBBBBB")
	q.Objectives[2].Block[0] = TypeKILL
	q.Objectives[2].Block[92] = 0
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, q))
	read, err := Read(&buf)
	require.NoError(t, err)
	assert.Equal(t, []byte("AAAAA"), read.Objectives[0].Name)
	assert.Equal(t, []byte("BBBBBBB"), read.Objectives[1].Name)
	assert.Len(t, read.Objectives[2].Name, 0)
}

func TestRead_NameLength255(t *testing.T) {
	q := minimalValidQuestFile()
	q.Objectives[0].Block[0] = TypeFIND
	q.Objectives[0].Block[92] = 255
	q.Objectives[0].Name = make([]byte, 255)
	for i := range q.Objectives[0].Name {
		q.Objectives[0].Name[i] = byte(i)
	}
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, q))
	read, err := Read(&buf)
	require.NoError(t, err)
	require.Len(t, read.Objectives[0].Name, 255)
	for i := range read.Objectives[0].Name {
		assert.Equal(t, byte(i), read.Objectives[0].Name[i])
	}
}

// --- 3. Continuation tests ---

func TestRead_ContinuationMissing(t *testing.T) {
	q := minimalValidQuestFile()
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, q))
	raw := buf.Bytes()
	truncated := raw[:len(raw)-4] // 4 bytes short of 12
	_, err := Read(bytes.NewReader(truncated))
	require.Error(t, err)
	assert.ErrorIs(t, err, io.ErrUnexpectedEOF)
}

func TestRead_ContinuationParsed(t *testing.T) {
	q := minimalValidQuestFile()
	q.Continuation[0] = 5001
	q.Continuation[1] = UnusedContinuation
	q.Continuation[2] = 102
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, q))
	read, err := Read(&buf)
	require.NoError(t, err)
	assert.Equal(t, uint32(5001), read.Continuation[0])
	assert.Equal(t, uint32(UnusedContinuation), read.Continuation[1])
	assert.Equal(t, uint32(102), read.Continuation[2])
}

func TestRead_ContinuationAllUnused(t *testing.T) {
	q := minimalValidQuestFile()
	q.Continuation[0] = UnusedContinuation
	q.Continuation[1] = UnusedContinuation
	q.Continuation[2] = UnusedContinuation
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, q))
	read, err := Read(&buf)
	require.NoError(t, err)
	assert.Equal(t, uint32(UnusedContinuation), read.Continuation[0])
	assert.Equal(t, uint32(UnusedContinuation), read.Continuation[1])
	assert.Equal(t, uint32(UnusedContinuation), read.Continuation[2])
}

func TestRead_TrailingBytesError(t *testing.T) {
	q := minimalValidQuestFile()
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, q))
	raw := buf.Bytes()
	withExtra := append(raw, 0x00, 0x01, 0x02)
	_, err := Read(bytes.NewReader(withExtra))
	require.Error(t, err)
	assert.ErrorIs(t, err, ErrTrailingBytes)
}

// --- 4. Structure tests ---

func TestRead_MinimalValidFile(t *testing.T) {
	q := minimalValidQuestFile()
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, q))
	assert.Equal(t, MinFileSize, buf.Len(), "minimal file must be 780 bytes")
	read, err := Read(&buf)
	require.NoError(t, err)
	assert.Equal(t, uint16(1), read.Header.QuestID())
	assert.Equal(t, uint16(100), read.Header.GivenNPCID())
}

func TestRead_MaximalFileSize(t *testing.T) {
	q := minimalValidQuestFile()
	for i := range q.Objectives {
		q.Objectives[i].Block[0] = TypeDROP
		q.Objectives[i].Block[92] = 255
		q.Objectives[i].Name = make([]byte, 255)
	}
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, q))
	expectedSize := MinFileSize + 7*255
	assert.Equal(t, expectedSize, buf.Len())
	read, err := Read(&buf)
	require.NoError(t, err)
	for i := range read.Objectives {
		require.Len(t, read.Objectives[i].Name, 255)
	}
}

// --- 5. Round-trip tests ---

func TestRoundTrip_BinaryIdentityMinimal(t *testing.T) {
	q := minimalValidQuestFile()
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, q))
	original := buf.Bytes()
	read, err := Read(bytes.NewReader(original))
	require.NoError(t, err)
	var buf2 bytes.Buffer
	require.NoError(t, Write(&buf2, read))
	newBytes := buf2.Bytes()
	assert.Equal(t, original, newBytes)
}

func TestRoundTrip_BinaryIdentityWithNames(t *testing.T) {
	q := minimalValidQuestFile()
	q.Objectives[0].Block[0] = TypeDROP
	q.Objectives[0].Block[92] = 5
	q.Objectives[0].Name = []byte("HELLO")
	q.Objectives[3].Block[0] = TypeFIND
	q.Objectives[3].Block[92] = 3
	q.Objectives[3].Name = []byte("XYZ")
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, q))
	original := buf.Bytes()
	read, err := Read(bytes.NewReader(original))
	require.NoError(t, err)
	var buf2 bytes.Buffer
	require.NoError(t, Write(&buf2, read))
	assert.Equal(t, original, buf2.Bytes())
}

func TestRoundTrip_StructEquality(t *testing.T) {
	q := minimalValidQuestFile()
	q.Header.QuestIDRaw[2] = 0xAB
	q.Header.QuestIDRaw[3] = 0xCD
	q.Objectives[1].Block[4] = 99
	q.Objectives[1].Block[92] = 0
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, q))
	read, err := Read(&buf)
	require.NoError(t, err)
	assert.Equal(t, q.Header, read.Header)
	assert.Equal(t, q.Continuation, read.Continuation)
	for i := range q.Objectives {
		assert.Equal(t, q.Objectives[i].Block, read.Objectives[i].Block)
		assert.Equal(t, q.Objectives[i].Name, read.Objectives[i].Name)
	}
}

// --- 6. Padding preservation (covered above; one more) ---

func TestRoundTrip_NonZeroPaddingPreserved(t *testing.T) {
	q := minimalValidQuestFile()
	q.Header.MinLevelPad = [3]byte{1, 2, 3}
	q.Header.RewardAreaPad = [8]byte{0x10, 0x20, 0x30, 0x40, 0x50, 0x60, 0x70, 0x80}
	// Objective internal padding (40-55, 60-75, 88-91)
	q.Objectives[0].Block[40] = 0xAA
	q.Objectives[0].Block[55] = 0xBB
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, q))
	read, err := Read(&buf)
	require.NoError(t, err)
	assert.Equal(t, [3]byte{1, 2, 3}, read.Header.MinLevelPad)
	assert.Equal(t, [8]byte{0x10, 0x20, 0x30, 0x40, 0x50, 0x60, 0x70, 0x80}, read.Header.RewardAreaPad)
	assert.Equal(t, byte(0xAA), read.Objectives[0].Block[40])
	assert.Equal(t, byte(0xBB), read.Objectives[0].Block[55])
}

// --- 7. Robustness ---

func TestRead_TruncatedAtVariousOffsets(t *testing.T) {
	q := minimalValidQuestFile()
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, q))
	raw := buf.Bytes()
	for cut := 0; cut < len(raw); cut++ {
		truncated := raw[:cut]
		_, err := Read(bytes.NewReader(truncated))
		require.Error(t, err, "truncation at offset %d must error", cut)
	}
}

func TestRead_RandomGarbageErrors(t *testing.T) {
	rng := bytes.NewReader(bytes.Repeat([]byte{0xDE, 0xAD, 0xBE, 0xEF}, 200))
	_, err := Read(rng)
	require.Error(t, err)
}

func TestRead_MalformedNameOverflow(t *testing.T) {
	// Valid header + 7 objectives; one objective claims NameLength=200 but we provide only 5 bytes
	q := minimalValidQuestFile()
	q.Objectives[0].Block[0] = TypeDROP
	q.Objectives[0].Block[92] = 200
	q.Objectives[0].Name = make([]byte, 5) // only 5 bytes
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, q))
	raw := buf.Bytes()
	// Actually we wrote 5 bytes; so file is 780 + 5. If we corrupt to say 200 and only have 5...
	// Better: build raw with first objective having nameLen=200 in block but only 5 bytes in stream
	raw[HeaderSize+92] = 200
	_, err := Read(bytes.NewReader(raw))
	require.Error(t, err)
	assert.ErrorIs(t, err, io.ErrUnexpectedEOF)
}

// --- 8. Optional semantic (skip or simple) ---
// Optional: KillCount>0 => Type=KILL etc. Omitted for format-only compliance.

// --- 9. Concurrency: Read/Write use no shared mutable state; safe. ---

func TestRead_ConcurrentReads(t *testing.T) {
	q := minimalValidQuestFile()
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, q))
	data := buf.Bytes()
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			_, _ = Read(bytes.NewReader(data))
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}

// --- 10. Little-endian ---

func TestRead_LittleEndianQuestID(t *testing.T) {
	q := minimalValidQuestFile()
	q.Header.QuestIDRaw = [4]byte{0x34, 0x12, 0, 0}
	var buf bytes.Buffer
	require.NoError(t, Write(&buf, q))
	read, err := Read(&buf)
	require.NoError(t, err)
	assert.Equal(t, uint16(0x1234), read.Header.QuestID())
}

func TestWrite_InvalidWriter(t *testing.T) {
	err := Write(&errWriter{}, minimalValidQuestFile())
	assert.Error(t, err)
}

type errWriter struct{}

func (errWriter) Write([]byte) (int, error) {
	return 0, io.ErrShortWrite
}

func TestHeader_Size(t *testing.T) {
	assert.Equal(t, HeaderSize, binary.Size(QuestHeader{}))
}

func TestQuestFile_MinFileSizeConstant(t *testing.T) {
	assert.Equal(t, 780, MinFileSize)
}
