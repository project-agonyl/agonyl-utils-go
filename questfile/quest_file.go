// Package questfile reads and writes the A3 binary quest file format:
// a 96-byte header, exactly 7 objective blocks (each 96 bytes + optional name),
// and a 12-byte continuation section. All multi-byte values are little-endian.
package questfile

import (
	"encoding/binary"
	"errors"
	"io"
)

// Format constants.
const (
	HeaderSize         = 96
	ObjectiveBlockSize = 96
	NumObjectives      = 7
	ContinuationSize   = 12
	MinFileSize        = HeaderSize + NumObjectives*ObjectiveBlockSize + ContinuationSize // 780
)

// Objective type constants (value at offset 0 in each objective block).
const (
	TypeKILL      = 0
	TypeQUESTITEM = 1
	TypeBRINGNPC  = 2
	TypeDROP      = 3
	TypeFIND      = 4

	// TypeUnused is the sentinel value (0xFF) used to mark empty/unused objective
	// slots. Real quest files always contain exactly 7 objective blocks; unused
	// slots are filled with 0xFF bytes rather than a valid type code.
	TypeUnused = 0xFF
)

// Sentinel values.
const (
	UnusedRewardItemCode = 0xFFFF
	UnusedContinuation   = 0xFFFFFFFF
)

// Sentinel errors.
var (
	// ErrInvalidObjectiveType is returned when an objective block's type byte is
	// not one of the five defined types (0–4) and is not the unused sentinel
	// (0xFF).
	ErrInvalidObjectiveType = errors.New("questfile: invalid objective type")

	// ErrNameLengthForType is returned when an objective whose type does not
	// support names (KILL, QUESTITEM, BRINGNPC) has a non-zero name-length byte.
	ErrNameLengthForType = errors.New("questfile: name length must be 0 for this objective type")

	// ErrTrailingBytes is returned when extra bytes are found after the
	// continuation section.
	ErrTrailingBytes = errors.New("questfile: trailing bytes after continuation")
)

// QuestHeader is the fixed 96-byte quest file header.
// Layout preserves padding for exact round-trip.
type QuestHeader struct {
	QuestIDRaw     [4]byte  // 0–3:   Quest ID (lower 16 bits) + 2 padding
	GivenNPCRaw    [4]byte  // 4–7:   Given NPC ID (lower 16 bits) + 2 padding
	TargetNPCBlock [24]byte // 8–31:  Target NPC (first 2 bytes = ID) + 22 bytes
	MinLevel       uint8    // 32
	MinLevelPad    [3]byte  // 33–35
	MaxLevel       uint8    // 36
	MaxLevelPad    [3]byte  // 37–39
	QuestFlags     uint32   // 40–43
	RewardSlot1    [4]byte  // 44–47: item code (UInt16) + 2 padding
	RewardSlot2    [4]byte  // 48–51
	RewardSlot3    [4]byte  // 52–55
	RewardSlot4Pad [4]byte  // 56–59: padding (not used as reward)
	RewardAreaPad  [8]byte  // 60–67
	Count1         uint8    // 68
	Count1Pad      [3]byte  // 69–71
	Count2         uint8    // 72
	Count2Pad      [3]byte  // 73–75
	Count3         uint8    // 76
	Count3Pad      [3]byte  // 77–79
	EXP            uint32   // 80–83
	Woonz          uint32   // 84–87
	Lore           uint32   // 88–91
	HeaderTail     [4]byte  // 92–95
}

// Objective is one of exactly 7 objectives: a 96-byte block plus optional name
// bytes. Unused slots have type byte 0xFF and all remaining bytes set to 0xFF
// (except the last four bytes which are 0x00, holding NameLength = 0).
type Objective struct {
	Block [96]byte // fixed block; NameLength at offset 92
	Name  []byte   // exactly NameLength bytes after block (only DROP/FIND, when > 0)
}

// QuestFile is the in-memory representation of an A3 quest file.
type QuestFile struct {
	Header       QuestHeader
	Objectives   [NumObjectives]Objective
	Continuation [3]uint32 // 0xFFFFFFFF = unused
}

// Read reads a complete quest file from r.
//
// Error conditions:
//   - io.ErrUnexpectedEOF  – file is truncated
//   - ErrInvalidObjectiveType – type byte is not 0–4 or 0xFF
//   - ErrNameLengthForType    – KILL/QUESTITEM/BRINGNPC block has non-zero name length
//   - ErrTrailingBytes        – extra data follows the continuation section
func Read(r io.Reader) (QuestFile, error) {
	var q QuestFile

	// ── Header: 96 bytes ────────────────────────────────────────────────────
	if err := binary.Read(r, binary.LittleEndian, &q.Header); err != nil {
		if err == io.EOF {
			return QuestFile{}, io.ErrUnexpectedEOF
		}

		return QuestFile{}, err
	}

	// ── Exactly 7 objectives ────────────────────────────────────────────────
	for i := range q.Objectives {
		if _, err := io.ReadFull(r, q.Objectives[i].Block[:]); err != nil {
			// io.ReadFull already converts EOF → ErrUnexpectedEOF when 0 bytes
			// were read, but we normalise both cases for clarity.
			if err == io.EOF {
				return QuestFile{}, io.ErrUnexpectedEOF
			}

			return QuestFile{}, err
		}

		objType := q.Objectives[i].Block[0]
		nameLen := q.Objectives[i].Block[92]

		// ErrInvalidObjectiveType. Real files fill unused objective slots with
		// 0xFF, so TypeUnused (0xFF) must be accepted as a valid no-op slot.
		// Any other out-of-range value (5–254) is still an error.
		if objType > TypeFIND && objType != TypeUnused {
			return QuestFile{}, ErrInvalidObjectiveType
		}

		// The name-length guard must also cover the unused (0xFF)
		// slot. An unused slot should always have nameLen == 0; if it somehow
		// does not, that is a malformed file. The original condition
		// (objType <= TypeBRINGNPC) silently skipped unused slots, which
		// could have caused a spurious name read on a junk byte at offset 92.
		// We now require nameLen == 0 for every type that does not support
		// names: KILL, QUESTITEM, BRINGNPC, and the unused sentinel.
		if objType != TypeDROP && objType != TypeFIND && nameLen != 0 {
			return QuestFile{}, ErrNameLengthForType
		}

		if nameLen > 0 {
			q.Objectives[i].Name = make([]byte, nameLen)
			if _, err := io.ReadFull(r, q.Objectives[i].Name); err != nil {
				if err == io.EOF {
					return QuestFile{}, io.ErrUnexpectedEOF
				}

				return QuestFile{}, err
			}
		}
	}

	// ── Continuation: 12 bytes (3 × uint32) ─────────────────────────────────
	for i := range q.Continuation {
		if err := binary.Read(r, binary.LittleEndian, &q.Continuation[i]); err != nil {
			if err == io.EOF {
				return QuestFile{}, io.ErrUnexpectedEOF
			}

			return QuestFile{}, err
		}
	}

	// The second clause fires when err is non-nil AND not io.EOF, which would
	// incorrectly return ErrTrailingBytes for legitimate read errors (e.g.
	// a network timeout). A read error here means we successfully parsed the
	// whole file; the error is on a speculative extra read and should be
	// ignored. We only care whether any bytes were actually returned.
	var one [1]byte
	n, _ := r.Read(one[:])
	if n > 0 {
		return QuestFile{}, ErrTrailingBytes
	}

	return q, nil
}

// Write writes q to w in A3 quest file binary format.
func Write(w io.Writer, q QuestFile) error {
	if err := binary.Write(w, binary.LittleEndian, &q.Header); err != nil {
		return err
	}

	for i := range q.Objectives {
		if _, err := w.Write(q.Objectives[i].Block[:]); err != nil {
			return err
		}

		if len(q.Objectives[i].Name) > 0 {
			if _, err := w.Write(q.Objectives[i].Name); err != nil {
				return err
			}
		}
	}

	if err := binary.Write(w, binary.LittleEndian, &q.Continuation); err != nil {
		return err
	}

	return nil
}

// QuestID returns the quest ID (lower 16 bits of the first header field).
func (h *QuestHeader) QuestID() uint16 {
	return binary.LittleEndian.Uint16(h.QuestIDRaw[:2])
}

// SetQuestID sets the quest ID while preserving the upper 2 padding bytes.
func (h *QuestHeader) SetQuestID(id uint16) {
	binary.LittleEndian.PutUint16(h.QuestIDRaw[:2], id)
}

// GivenNPCID returns the given NPC ID (lower 16 bits).
func (h *QuestHeader) GivenNPCID() uint16 {
	return binary.LittleEndian.Uint16(h.GivenNPCRaw[:2])
}

// SetGivenNPCID sets the given NPC ID while preserving padding.
func (h *QuestHeader) SetGivenNPCID(id uint16) {
	binary.LittleEndian.PutUint16(h.GivenNPCRaw[:2], id)
}

// ObjectiveType returns the objective type byte at offset 0 in the block.
func (o *Objective) ObjectiveType() uint8 {
	return o.Block[0]
}

// IsUnused reports whether this objective slot is an unused (0xFF-filled) slot.
func (o *Objective) IsUnused() bool {
	return o.Block[0] == TypeUnused
}

// NameLength returns the name length byte at offset 92 in the block.
func (o *Objective) NameLength() uint8 {
	return o.Block[92]
}
