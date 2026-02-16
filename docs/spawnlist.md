# Spawnlist Package

Documentation for the `github.com/project-agonyl/agonyl-utils-go/spawnlist` package: read and write the spawn list binary format (a contiguous sequence of little-endian spawn entries: position, orientation, and related fields).

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

The `spawnlist` package provides:

- **Read** — reads a spawn list from an `io.Reader`: the entire stream is decoded as a contiguous sequence of **SpawnListItem** values until EOF. Returns a **SpawnList** slice or an error if the stream is truncated (e.g. byte length not a multiple of item size) or a read fails.
- **Write** — writes a **SpawnList** to an `io.Writer` in the same format (items only; no count prefix).
- **SpawnListItem** — a single spawn entry with ID, X/Y coordinates, reserved field, orientation, and spawn step.
- **SpawnList** — a slice of **SpawnListItem**, used as the in-memory representation and the argument/return type for **Read** and **Write**.

Typical use cases include loading or saving spawn list files used by the A3/Agonyl client (e.g. map spawn definitions for NPCs or objects).

---

## Installation

```bash
go get github.com/project-agonyl/agonyl-utils-go
```

Import in your code:

```go
import "github.com/project-agonyl/agonyl-utils-go/spawnlist"
```

---

## API Reference

### Type: `SpawnListItem`

```go
type SpawnListItem struct {
    Id          uint16 // Spawn/npc identifier
    X           byte   // X coordinate
    Y           byte   // Y coordinate
    Unknown1    uint16 // Reserved
    Orientation byte   // Facing direction
    SpwanStep   byte   // Spawn step
}
```

A single spawn entry as stored in the spawn list file.

- **Id** — spawn or NPC identifier (uint16).
- **X**, **Y** — position coordinates (bytes).
- **Unknown1** — reserved uint16; layout and meaning are format-specific.
- **Orientation** — facing direction (byte).
- **SpwanStep** — spawn step value (byte). (Note: field name is spelled as in the format.)

---

### Type: `SpawnList`

```go
type SpawnList []SpawnListItem
```

A slice of spawn entries as stored in the spawn list file. Used as both the in-memory representation and the argument/return type for **Read** and **Write**.

---

### Function: `Read`

```go
func Read(r io.Reader) (SpawnList, error)
```

Reads a spawn list from **r**. The entire stream is consumed and decoded as a contiguous sequence of **SpawnListItem** values. The number of items is implied by the stream length (must be a multiple of the fixed item size, 8 bytes).

- **r** — source of binary data (e.g. file, buffer).
- **Returns** — decoded **SpawnList** and **nil** on success; **nil** and a non-nil **error** (e.g. **io.ErrUnexpectedEOF** if the byte count is not a multiple of 8) if the stream is truncated or a read fails.

---

### Function: `Write`

```go
func Write(w io.Writer, data SpawnList) error
```

Writes **data** to **w** in spawn list format: each **SpawnListItem** in order, with no leading count. Accepts **nil** or empty slice (writes zero bytes).

- **w** — destination for binary data.
- **data** — slice of spawn entries to write.
- **Returns** — **nil** on success; non-nil **error** if a write fails.

---

## Binary Format

| Part     | Type   | Description                                      |
|----------|--------|--------------------------------------------------|
| Item 0   | struct | One **SpawnListItem** (Id, X, Y, Unknown1, Orientation, SpwanStep). |
| Item 1   | struct | Same.                                            |
| …        | …      | Repeated until end of stream.                    |

There is **no entry count**; the file is a raw sequence of fixed-size records. Each **SpawnListItem** is 8 bytes (little-endian): Id (2), X (1), Y (1), Unknown1 (2), Orientation (1), SpwanStep (1).

---

## Usage

### Read a spawn list from a file

```go
f, err := os.Open("spawnlist.bin")
if err != nil {
    log.Fatal(err)
}
defer f.Close()

data, err := spawnlist.Read(f)
if err != nil {
    log.Fatal(err)
}

for i := range data {
    item := &data[i]
    log.Printf("Spawn %d: ID=%d at (%d,%d) orientation=%d",
        i, item.Id, item.X, item.Y, item.Orientation)
}
```

### Write a spawn list

```go
data := spawnlist.SpawnList{
    {Id: 1, X: 10, Y: 20, Orientation: 0, SpwanStep: 1},
    {Id: 2, X: 30, Y: 40, Orientation: 2, SpwanStep: 0},
}

f, _ := os.Create("spawnlist.bin")
defer f.Close()

if err := spawnlist.Write(f, data); err != nil {
    log.Fatal(err)
}
```

### Round-trip (read, modify, write)

```go
buf := &bytes.Buffer{}
if err := spawnlist.Write(buf, data); err != nil {
    log.Fatal(err)
}

readBack, err := spawnlist.Read(buf)
if err != nil {
    log.Fatal(err)
}
// readBack has the same count and items as data
```

### Empty list

```go
// Writing nil or empty slice produces zero bytes
var list spawnlist.SpawnList
_ = spawnlist.Write(w, list)

// Reading an empty stream returns a non-nil empty slice and no error
data, err := spawnlist.Read(bytes.NewReader(nil))
// err == nil, len(data) == 0
```

---

## Testing

The package can be tested with the standard library. Run:

```bash
go test ./spawnlist/...
```

Recommended coverage:

- **Read** with empty stream returns an empty list and no error.
- **Read** with valid one or more items returns the correct **SpawnList**.
- **Read** with truncated stream (byte length not a multiple of 8) returns an error (e.g. **io.ErrUnexpectedEOF**).
- **Read** with failing reader returns an error.
- **Write** with nil or empty slice produces zero bytes.
- **Write** with one or more items produces the expected byte layout (8 bytes per item, little-endian).
- **Write** with failing writer returns an error.
- **Write** then **Read** round-trips to the same **SpawnList** (all fields preserved).
- Decoding known bytes yields the expected **SpawnListItem** field values.
- **SpawnListItem** has fixed size 8 bytes for format compatibility.
