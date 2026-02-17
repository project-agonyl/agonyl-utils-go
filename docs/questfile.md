# Questfile Package

Documentation for the `github.com/project-agonyl/agonyl-utils-go/questfile` package: read and write the A3 binary quest file format (96-byte header, exactly 7 objective blocks with optional names, and 12-byte continuation section).

---

## Table of Contents

- [Overview](#overview)
- [Installation](#installation)
- [API Reference](#api-reference)
- [Binary Format](#binary-format)
- [Usage](#usage)
- [Testing](#testing)

---

## Overview

The `questfile` package provides:

- **Read** — reads a complete quest file from an `io.Reader`. Returns `QuestFile` or an error if the stream is truncated, has invalid objective type, invalid name length for type, or trailing bytes after the continuation section.
- **Write** — writes a `QuestFile` to an `io.Writer` in A3 quest binary format.
- **QuestFile** — in-memory representation: **QuestHeader** (96 bytes), exactly 7 **Objective** blocks (each 96 bytes + optional name bytes), and **Continuation** (3× uint32).
- **QuestHeader** — quest ID, given NPC, target NPC block (24 bytes), min/max level, reward item slots and counts, EXP/Woonz/Lore, and padding. All padding is preserved for bit-exact round-trip.
- **Objective** — 96-byte block (type, map/location/radius, monster/NPC, kill count, quest item, drop IDs/probabilities, name length at offset 92) plus optional **Name** bytes for DROP/FIND types. Unused slots use type **TypeUnused** (0xFF) with name length 0.
- **QuestID**, **SetQuestID**, **GivenNPCID**, **SetGivenNPCID** — accessors for header IDs (lower 16 bits; padding preserved).
- **ObjectiveType**, **NameLength**, **IsUnused** — accessors on **Objective**; **IsUnused** reports whether the slot is an unused (0xFF) slot.

Typical use cases include loading or saving A3 quest definition files (e.g. from game data or server tooling).

---

## Installation

```bash
go get github.com/project-agonyl/agonyl-utils-go
```

Import in your code:

```go
import "github.com/project-agonyl/agonyl-utils-go/questfile"
```

---

## API Reference

### Constants

- **HeaderSize** = 96  
- **ObjectiveBlockSize** = 96  
- **NumObjectives** = 7  
- **ContinuationSize** = 12  
- **MinFileSize** = 780 (no objective names)  
- **TypeKILL**, **TypeQUESTITEM**, **TypeBRINGNPC**, **TypeDROP**, **TypeFIND** — objective type values (0–4).  
- **TypeUnused** = 0xFF — sentinel for empty/unused objective slots; real quest files always have 7 blocks, and unused slots are filled with 0xFF.  
- **UnusedRewardItemCode** = 0xFFFF  
- **UnusedContinuation** = 0xFFFFFFFF  

### Errors

- **ErrInvalidObjectiveType** — objective type byte is not 0–4 and not **TypeUnused** (0xFF).  
- **ErrNameLengthForType** — name length is non-zero for a type that does not support names: KILL, QUESTITEM, BRINGNPC, or unused (0xFF). Only DROP and FIND may have names.  
- **ErrTrailingBytes** — extra bytes after the 12-byte continuation section.  

Truncation returns **io.ErrUnexpectedEOF** (or an error wrapping it).

### Type: `QuestFile`

```go
type QuestFile struct {
    Header       QuestHeader
    Objectives   [7]Objective
    Continuation [3]uint32  // 0xFFFFFFFF = unused
}
```

### Type: `QuestHeader`

96-byte header with padding preserved. Fields include **QuestIDRaw**, **GivenNPCRaw**, **TargetNPCBlock** (24 bytes), **MinLevel**, **MaxLevel**, **QuestFlags**, reward slots (**RewardSlot1**–**Slot3**, **RewardSlot4Pad**), **RewardAreaPad**, **Count1**–**Count3** (and pads), **EXP**, **Woonz**, **Lore**, **HeaderTail**. Use **QuestID()** / **SetQuestID()** and **GivenNPCID()** / **SetGivenNPCID()** for the logical 16-bit IDs.

### Type: `Objective`

```go
type Objective struct {
    Block [96]byte  // fixed block; type at 0, name length at 92
    Name  []byte    // exactly NameLength bytes after block (for DROP/FIND when > 0)
}
```

Unused slots have **Block[0]** = **TypeUnused** (0xFF) and **NameLength** = 0. Use **IsUnused()** to detect them.

### Function: `Read`

```go
func Read(r io.Reader) (QuestFile, error)
```

Reads a complete quest file from **r**. Returns **QuestFile** and **nil** on success. Returns **io.ErrUnexpectedEOF** on truncation, **ErrInvalidObjectiveType** when the type byte is not 0–4 and not **TypeUnused** (0xFF), **ErrNameLengthForType** when a non-name type (KILL/QUESTITEM/BRINGNPC/unused) has non-zero name length, and **ErrTrailingBytes** if data remains after the continuation.

### Function: `Write`

```go
func Write(w io.Writer, q QuestFile) error
```

Writes **q** to **w** in A3 quest file binary format (little-endian). All padding is written as stored for bit-exact round-trip.

### Method: `Objective.IsUnused`

```go
func (o *Objective) IsUnused() bool
```

Reports whether this objective slot is unused (type byte at offset 0 is **TypeUnused**, 0xFF).

---

## Binary Format

- **Little-endian** throughout.  
- **Header**: 96 bytes (see documentation PDF for offset table). Quest ID and Given NPC use lower 16 bits of 4-byte fields; Target NPC is 24 bytes; reward slots are 4 bytes each (2-byte item code + 2 padding); counts are 1 byte in 4-byte fields; EXP/Woonz/Lore are uint32; tail 4 bytes padding.  
- **Objectives**: Exactly 7. Each is 96 bytes then, if **NameLength** (offset 92) &gt; 0, exactly **NameLength** bytes of name. For types 0 (KILL), 1 (QUESTITEM), 2 (BRINGNPC), and unused (0xFF), **NameLength** must be 0. For 3 (DROP) and 4 (FIND), name is optional. Unused slots use type byte 0xFF.  
- **Continuation**: 12 bytes (3× uint32). **0xFFFFFFFF** means no continuation in that slot.  
- **Trailing**: No bytes may follow the continuation; otherwise **Read** returns **ErrTrailingBytes**.  

Minimum file size: 780 bytes. Maximum: 780 + 7×255 name bytes.

---

## Usage

### Read a quest file

```go
f, err := os.Open("quest.dat")
if err != nil {
    log.Fatal(err)
}
defer f.Close()

q, err := questfile.Read(f)
if err != nil {
    log.Fatal(err)
}

log.Printf("Quest ID=%d", q.Header.QuestID())
for i, obj := range q.Objectives {
    if obj.IsUnused() {
        continue
    }
    
    log.Printf("Objective %d type=%d nameLen=%d", i+1, obj.ObjectiveType(), obj.NameLength())
}
```

### Write a quest file

Build a **QuestFile** with header (e.g. **SetQuestID**, **SetGivenNPCID**, **EXP**, **Woonz**, **Lore**), exactly 7 **Objective**s (each **Block** set, and **Name** only for DROP/FIND when needed), and **Continuation** (use **UnusedContinuation** for empty slots):

```go
var q questfile.QuestFile
q.Header.SetQuestID(100)
q.Header.SetGivenNPCID(200)
q.Header.EXP = 5000
q.Continuation[0] = 2001
q.Continuation[1] = questfile.UnusedContinuation
q.Continuation[2] = questfile.UnusedContinuation
// Set q.Objectives[0..6].Block and optional Name for each

f, _ := os.Create("quest.dat")
defer f.Close()
if err := questfile.Write(f, q); err != nil {
    log.Fatal(err)
}
```

### Round-trip

```go
var buf bytes.Buffer
if err := questfile.Write(&buf, q); err != nil {
    log.Fatal(err)
}
read, err := questfile.Read(&buf)
if err != nil {
    log.Fatal(err)
}
// read equals q for all header, objectives, continuation
```

---

## Testing

Run:

```bash
go test ./questfile/...
```

Benchmarks (minimal and maximal file):

```bash
go test -bench=. ./questfile/...
```

Tests cover: header size and field/padding preservation, objective count and type validation, name length rules, continuation and trailing bytes, minimal/maximal file size, binary and struct round-trip, truncation and corruption, and concurrency.
