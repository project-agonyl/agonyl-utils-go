# Monsterbin Package

Documentation for the `github.com/project-agonyl/agonyl-utils-go/monsterbin` package: read and write the A3 client monster bin binary format (little-endian entry count followed by fixed-size monster records).

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

The `monsterbin` package provides:

- **Read** — reads a monster bin from an `io.Reader`: a uint32 entry count then each fixed-size monster item. Returns a `MonsterBin` slice or an error if the stream is truncated or invalid.
- **Write** — writes a `MonsterBin` to an `io.Writer` in the same format (count then items).
- **MonsterBinItem** — a single monster record with ID, name (0x1F bytes), and reserved bytes (0x3D).
- **GetName** — method on `MonsterBinItem` that returns the monster name as a string (trimmed of null padding).

Typical use cases include loading or saving monster definition files used by the A3/Agonyl client (e.g. from game data or tooling).

---

## Installation

```bash
go get github.com/project-agonyl/agonyl-utils-go
```

Import in your code:

```go
import "github.com/project-agonyl/agonyl-utils-go/monsterbin"
```

---

## API Reference

### Type: `MonsterBinItem`

```go
type MonsterBinItem struct {
    ID      uint32
    Name    [0x1F]byte
    Unknown [0x3D]byte
}
```

A single monster record.

- **ID** — monster identifier (uint32).
- **Name** — fixed 0x1F (31) bytes; use **GetName()** for a trimmed string.
- **Unknown** — 0x3D (61) bytes of reserved/padding data; layout is format-specific.

---

### Type: `MonsterBin`

```go
type MonsterBin []MonsterBinItem
```

A slice of monster entries as stored in the bin file. Used as both the in-memory representation and the argument/return type for **Read** and **Write**.

---

### Function: `Read`

```go
func Read(r io.Reader) (MonsterBin, error)
```

Reads a monster bin from **r**: first a little-endian uint32 entry count, then that many **MonsterBinItem** records in sequence.

- **r** — source of binary data (e.g. file, buffer).
- **Returns** — decoded **MonsterBin** and **nil** on success; **nil** and a non-nil **error** if the stream is truncated or a read fails.

---

### Function: `Write`

```go
func Write(w io.Writer, data MonsterBin) error
```

Writes **data** to **w** in monster bin format: a little-endian uint32 count equal to `len(data)`, then each **MonsterBinItem** in order.

- **w** — destination for binary data.
- **data** — slice of monster entries to write.
- **Returns** — **nil** on success; non-nil **error** if a write fails.

---

### Method: `MonsterBinItem.GetName`

```go
func (m *MonsterBinItem) GetName() string
```

Returns the monster name as a string. The fixed **Name** field (0x1F bytes) is interpreted as a null-padded string and trimmed to the first null or end of buffer.

---

## Binary Format

| Part        | Type   | Description                          |
|------------|--------|--------------------------------------|
| Entry count| uint32 | Little-endian; number of items.      |
| Item 0     | struct | One **MonsterBinItem** (ID + Name + Unknown). |
| Item 1     | struct | Same.                                |
| …          | …      | Repeated for **entry count** items.  |

Each **MonsterBinItem** is a fixed 4 + 0x1F + 0x3D = 100 bytes (4 + 31 + 61).

---

## Usage

### Read a monster bin from a file

```go
f, err := os.Open("monster.bin")
if err != nil {
    log.Fatal(err)
}
defer f.Close()

data, err := monsterbin.Read(f)
if err != nil {
    log.Fatal(err)
}

for i := range data {
    name := data[i].GetName()
    log.Printf("Monster %d: ID=%d Name=%q", i, data[i].ID, name)
}
```

### Write a monster bin

```go
data := monsterbin.MonsterBin{
    {ID: 1, Name: [0x1F]byte{'G', 'o', 'b', 'l', 'i', 'n'}}, // rest zero
    {ID: 2, Name: [0x1F]byte{'O', 'r', 'c'}},
}

f, _ := os.Create("monster.bin")
defer f.Close()

if err := monsterbin.Write(f, data); err != nil {
    log.Fatal(err)
}
```

### Round-trip (read, modify, write)

```go
buf := &bytes.Buffer{}
if _, err := monsterbin.Write(buf, data); err != nil {
    log.Fatal(err)
}

readBack, err := monsterbin.Read(buf)
if err != nil {
    log.Fatal(err)
}
// readBack has the same count and items as data
```

---

## Testing

The package can be tested with the standard library. Run:

```bash
go test ./monsterbin/...
```

Recommended coverage:

- **Read** with valid input returns the correct count and items.
- **Read** with truncated input or empty reader returns an error.
- **Write** then **Read** round-trips to the same **MonsterBin**.
- **GetName** returns the name trimmed at the first null and handles empty or full names.
