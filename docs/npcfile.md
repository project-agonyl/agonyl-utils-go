# Npcfile Package

Documentation for the `github.com/project-agonyl/agonyl-utils-go/npcfile` package: read and write the NPC file binary format (a single fixed-size little-endian record containing NPC stats, attacks, and name).

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

The `npcfile` package provides:

- **Read** — reads a single NPC record from an `io.Reader` in little-endian binary format. Returns `NPCFileData` or an error if the stream is truncated or invalid.
- **Write** — writes one `NPCFileData` to an `io.Writer` in the same format.
- **NPCFileData** — a single NPC record with name (0x14 bytes), ID, respawn/attack/defense stats, up to three **NPCAttack** slots, movement speed, level, HP, attack defenses, and related fields.
- **NPCAttack** — one attack slot (range, area, damage, additional damage).
- **GetName** — method on `NPCFileData` that returns the NPC display name as a string (trimmed of null padding).

Typical use cases include loading or saving NPC definition files used by the A3/Agonyl client (e.g. from game data or tooling).

---

## Installation

```bash
go get github.com/project-agonyl/agonyl-utils-go
```

Import in your code:

```go
import "github.com/project-agonyl/agonyl-utils-go/npcfile"
```

---

## API Reference

### Type: `NPCFileData`

```go
type NPCFileData struct {
    Name                [0x14]byte
    Id                  uint16
    RespawnRate         uint16
    AttackTypeInfo      byte
    TargetSelectionInfo byte
    Defense             byte
    AdditionalDefense   byte
    Attacks             [0x3]NPCAttack
    AttackSpeedLow      uint16
    AttackSpeedHigh     uint16
    MovementSpeed       uint32
    Level               byte
    PlayerExp           uint16
    Appearance          byte
    HP                  uint32
    BlueAttackDefense   uint16
    RedAttackDefense    uint16
    GreyAttackDefense   uint16
    MercenaryExp        uint16
    Unknown             uint16
}
```

A single NPC record as stored in the NPC file.

- **Name** — fixed 0x14 (20) bytes, null-padded; use **GetName()** for a trimmed string.
- **Id** — NPC identifier.
- **RespawnRate** — respawn timing (format-specific).
- **AttackTypeInfo**, **TargetSelectionInfo** — attack/target behavior flags.
- **Defense**, **AdditionalDefense** — defense values.
- **Attacks** — up to 3 attack definitions (range, area, damage).
- **AttackSpeedLow**, **AttackSpeedHigh** — attack speed bounds.
- **MovementSpeed** — movement speed (uint32).
- **Level** — NPC level.
- **PlayerExp** — experience granted to player.
- **Appearance** — appearance/model ID.
- **HP** — hit points.
- **BlueAttackDefense**, **RedAttackDefense**, **GreyAttackDefense** — attack-type defenses.
- **MercenaryExp** — experience for mercenaries.
- **Unknown** — reserved field.

---

### Type: `NPCAttack`

```go
type NPCAttack struct {
    Range            uint16
    Area             uint16
    Damage           uint16
    AdditionalDamage uint16
}
```

One attack slot for an NPC (range, area of effect, damage, and additional damage).

---

### Function: `Read`

```go
func Read(r io.Reader) (NPCFileData, error)
```

Reads a single NPC record from **r** in little-endian binary format.

- **r** — source of binary data (e.g. file, buffer).
- **Returns** — decoded **NPCFileData** and **nil** on success; zero value and a non-nil **error** if the stream is truncated or a read fails.

---

### Function: `Write`

```go
func Write(w io.Writer, data NPCFileData) error
```

Writes **data** to **w** in NPC file format (one fixed-size little-endian record).

- **w** — destination for binary data.
- **data** — NPC record to write.
- **Returns** — **nil** on success; non-nil **error** if a write fails.

---

### Method: `NPCFileData.GetName`

```go
func (n *NPCFileData) GetName() string
```

Returns the NPC display name as a string. The fixed **Name** field (0x14 bytes) is interpreted as a null-padded string and trimmed to the first null or end of buffer.

---

## Binary Format

The file contains **one** fixed-size record (no entry count). All multi-byte values are little-endian.

| Part     | Type         | Description                                      |
|----------|--------------|--------------------------------------------------|
| Name     | [0x14]byte   | Null-padded display name (20 bytes).             |
| Id       | uint16       | NPC ID.                                          |
| RespawnRate | uint16     | Respawn rate.                                    |
| AttackTypeInfo | byte    | Attack type flag.                                |
| TargetSelectionInfo | byte | Target selection flag.                    |
| Defense  | byte         | Defense value.                                   |
| AdditionalDefense | byte   | Additional defense.                             |
| Attacks  | [3]NPCAttack | Three attack slots (each 8 bytes).               |
| …        | …            | Remaining fields (attack speed, movement, level, HP, defenses, etc.). |

Each **NPCAttack** is 8 bytes: Range (uint16), Area (uint16), Damage (uint16), AdditionalDamage (uint16).

Total **NPCFileData** size is 78 bytes (20 + 2×2 + 4×1 + 3×8 + 2×2 + 4 + 1 + 2 + 1 + 4 + 5×2).

---

## Usage

### Read an NPC file

```go
f, err := os.Open("npc.dat")
if err != nil {
    log.Fatal(err)
}
defer f.Close()

data, err := npcfile.Read(f)
if err != nil {
    log.Fatal(err)
}

name := data.GetName()
log.Printf("NPC ID=%d Name=%q Level=%d HP=%d", data.Id, name, data.Level, data.HP)
```

### Write an NPC file

```go
data := npcfile.NPCFileData{
    Id:           42,
    Level:        5,
    HP:           1000,
    RespawnRate:  30,
}
copy(data.Name[:], "Guard")
data.Attacks[0] = npcfile.NPCAttack{Range: 50, Area: 10, Damage: 100, AdditionalDamage: 20}

f, _ := os.Create("npc.dat")
defer f.Close()

if err := npcfile.Write(f, data); err != nil {
    log.Fatal(err)
}
```

### Round-trip (read, modify, write)

```go
buf := &bytes.Buffer{}
if err := npcfile.Write(buf, data); err != nil {
    log.Fatal(err)
}

readBack, err := npcfile.Read(buf)
if err != nil {
    log.Fatal(err)
}
// readBack equals data
```

### Building names with copy

```go
npc := npcfile.NPCFileData{Id: 1, Level: 10}
copy(npc.Name[:], "Blacksmith")
```

---

## Testing

The package can be tested with the standard library. Run:

```bash
go test ./npcfile/...
```

Recommended coverage:

- **Read** with a valid full record returns the same data as written.
- **Read** with truncated input, empty stream, or failing reader returns an error.
- **Write** with a failing writer returns an error.
- **Write** then **Read** round-trips to the same **NPCFileData** (including all fields and attacks).
- **GetName** returns the name trimmed at the first null; empty name returns `""`; full 0x14-byte name returns 20 characters.
