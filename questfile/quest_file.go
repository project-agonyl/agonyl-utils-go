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
)

// Sentinel values.
const (
	UnusedRewardItemCode = 0xFFFF
	UnusedContinuation   = 0xFFFFFFFF
)

// Sentinel errors.
var (
	ErrInvalidObjectiveType = errors.New("questfile: invalid objective type")
	ErrNameLengthForType    = errors.New("questfile: name length must be 0 for this objective type")
	ErrTrailingBytes        = errors.New("questfile: trailing bytes after continuation")
)

// QuestHeader is the fixed 96-byte quest file header.
// Layout preserves padding for exact round-trip.
type QuestHeader struct {
	QuestIDRaw     [4]byte  // 0-3: Quest ID (lower 16 bits) + 2 padding
	GivenNPCRaw    [4]byte  // 4-7: Given NPC ID (lower 16 bits) + 2 padding
	TargetNPCBlock [24]byte // 8-31: Target NPC (first 2 bytes = ID) + 22 bytes
	MinLevel       uint8    // 32
	MinLevelPad    [3]byte  // 33-35
	MaxLevel       uint8    // 36
	MaxLevelPad    [3]byte  // 37-39
	QuestFlags     uint32   // 40-43
	RewardSlot1    [4]byte  // 44-47: item code (UInt16) + 2 padding
	RewardSlot2    [4]byte  // 48-51
	RewardSlot3    [4]byte  // 52-55
	RewardSlot4Pad [4]byte  // 56-59: padding (not used)
	RewardAreaPad  [8]byte  // 60-67
	Count1         uint8    // 68
	Count1Pad      [3]byte  // 69-71
	Count2         uint8    // 72
	Count2Pad      [3]byte  // 73-75
	Count3         uint8    // 76
	Count3Pad      [3]byte  // 77-79
	EXP            uint32   // 80-83
	Woonz          uint32   // 84-87
	Lore           uint32   // 88-91
	HeaderTail     [4]byte  // 92-95
}

// Objective is one of exactly 7 objectives: a 96-byte block plus optional name bytes.
type Objective struct {
	Block [96]byte // fixed block; NameLength at offset 92
	Name  []byte   // exactly NameLength bytes after block (only for DROP/FIND when > 0)
}

// QuestFile is the in-memory representation of an A3 quest file.
type QuestFile struct {
	Header       QuestHeader
	Objectives   [NumObjectives]Objective
	Continuation [3]uint32 // 0xFFFFFFFF = unused
}

// Read reads a complete quest file from r.
// Returns io.ErrUnexpectedEOF on truncation, ErrInvalidObjectiveType for invalid type,
// ErrNameLengthForType when name length is non-zero for KILL/QUESTITEM/BRINGNPC,
// and ErrTrailingBytes if extra data follows the continuation section.
func Read(r io.Reader) (QuestFile, error) {
	var q QuestFile

	// Header: 96 bytes
	if err := binary.Read(r, binary.LittleEndian, &q.Header); err != nil {
		if err == io.EOF {
			return QuestFile{}, io.ErrUnexpectedEOF
		}

		return QuestFile{}, err
	}

	// Exactly 7 objectives
	for i := range q.Objectives {
		if _, err := io.ReadFull(r, q.Objectives[i].Block[:]); err != nil {
			if err == io.EOF {
				return QuestFile{}, io.ErrUnexpectedEOF
			}

			return QuestFile{}, err
		}

		objType := q.Objectives[i].Block[0]
		nameLen := q.Objectives[i].Block[92]

		if objType > TypeFIND {
			return QuestFile{}, ErrInvalidObjectiveType
		}

		if objType <= TypeBRINGNPC && nameLen != 0 {
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

	// Continuation: 12 bytes (3 x uint32)
	for i := range q.Continuation {
		if err := binary.Read(r, binary.LittleEndian, &q.Continuation[i]); err != nil {
			if err == io.EOF {
				return QuestFile{}, io.ErrUnexpectedEOF
			}

			return QuestFile{}, err
		}
	}

	// No trailing bytes allowed
	var one [1]byte
	if n, err := r.Read(one[:]); n > 0 || (err != nil && err != io.EOF) {
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

// QuestID returns the quest ID (lower 16 bits of first header field).
func (h *QuestHeader) QuestID() uint16 {
	return binary.LittleEndian.Uint16(h.QuestIDRaw[:2])
}

// SetQuestID sets the quest ID while preserving the upper 16 bits (padding).
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

// NameLength returns the name length byte at offset 92 in the block.
func (o *Objective) NameLength() uint8 {
	return o.Block[92]
}
