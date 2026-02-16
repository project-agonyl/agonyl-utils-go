# Mapbin Package

Documentation for the `github.com/project-agonyl/agonyl-utils-go/mapbin` package: read and write the A3 client map bin binary format (little-endian entry count followed by fixed-size map records).

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

The `mapbin` package provides:

- **Read** — reads a map bin from an `io.Reader`: a uint32 entry count then each fixed-size map item. Returns a `MapBin` slice or an error if the stream is truncated or invalid.
- **Write** — writes a `MapBin` to an `io.Writer` in the same format (count then items).
- **MapBinItem** — a single map record with ID, five reserved uint32 fields (Unknown1–Unknown5), and name (0x20 bytes).
- **GetName** — method on `MapBinItem` that returns the map name as a string (trimmed of null padding).

Typical use cases include loading or saving map definition files used by the A3/Agonyl client (e.g. from game data or tooling).

---

## Installation

```bash
go get github.com/project-agonyl/agonyl-utils-go
```

Import in your code:

```go
import "github.com/project-agonyl/agonyl-utils-go/mapbin"
```

---

## API Reference

### Type: `MapBinItem`

```go
type MapBinItem struct {
    ID       uint32
    Unknown1 uint32
    Unknown2 uint32
    Unknown3 uint32
    Unknown4 uint32
    Unknown5 uint32
    Name     [0x20]byte
}
```

A single map record.

- **ID** — map identifier (uint32).
- **Unknown1** through **Unknown5** — reserved uint32 fields; layout and meaning are format-specific.
- **Name** — fixed 0x20 (32) bytes; use **GetName()** for a trimmed string.

---

### Type: `MapBin`

```go
type MapBin []MapBinItem
```

A slice of map entries as stored in the bin file. Used as both the in-memory representation and the argument/return type for **Read** and **Write**.

---

### Function: `Read`

```go
func Read(r io.Reader) (MapBin, error)
```

Reads a map bin from **r**: first a little-endian uint32 entry count, then that many **MapBinItem** records in sequence.

- **r** — source of binary data (e.g. file, buffer).
- **Returns** — decoded **MapBin** and **nil** on success; **nil** and a non-nil **error** if the stream is truncated or a read fails.

---

### Function: `Write`

```go
func Write(w io.Writer, data MapBin) error
```

Writes **data** to **w** in map bin format: a little-endian uint32 count equal to `len(data)`, then each **MapBinItem** in order.

- **w** — destination for binary data.
- **data** — slice of map entries to write.
- **Returns** — **nil** on success; non-nil **error** if a write fails.

---

### Method: `MapBinItem.GetName`

```go
func (m *MapBinItem) GetName() string
```

Returns the map name as a string. The fixed **Name** field (0x20 bytes) is interpreted as a null-padded string and trimmed to the first null or end of buffer.

---

## Binary Format

| Part        | Type   | Description                          |
|------------|--------|--------------------------------------|
| Entry count| uint32 | Little-endian; number of items.      |
| Item 0     | struct | One **MapBinItem** (ID + Unknown1–5 + Name). |
| Item 1     | struct | Same.                                |
| …          | …      | Repeated for **entry count** items.  |

Each **MapBinItem** is a fixed 4 + (5×4) + 0x20 = 56 bytes (ID + five uint32s + 32-byte name).

---

## Usage

### Read a map bin from a file

```go
f, err := os.Open("map.bin")
if err != nil {
    log.Fatal(err)
}
defer f.Close()

data, err := mapbin.Read(f)
if err != nil {
    log.Fatal(err)
}

for i := range data {
    name := data[i].GetName()
    log.Printf("Map %d: ID=%d Name=%q", i, data[i].ID, name)
}
```

### Write a map bin

```go
data := mapbin.MapBin{
    {ID: 1, Unknown1: 0, Name: [0x20]byte{'F', 'o', 'r', 'e', 's', 't'}}, // rest zero
    {ID: 2, Name: [0x20]byte{'D', 'u', 'n', 'g', 'e', 'o', 'n'}},
}

f, _ := os.Create("map.bin")
defer f.Close()

if err := mapbin.Write(f, data); err != nil {
    log.Fatal(err)
}
```

### Round-trip (read, modify, write)

```go
buf := &bytes.Buffer{}
if err := mapbin.Write(buf, data); err != nil {
    log.Fatal(err)
}

readBack, err := mapbin.Read(buf)
if err != nil {
    log.Fatal(err)
}
// readBack has the same count and items as data
```

### Building items with copy (recommended for names)

```go
item := mapbin.MapBinItem{ID: 100, Unknown1: 1, Unknown2: 2}
copy(item.Name[:], "Town Center")

data := mapbin.MapBin{item}
```

---

## Testing

The package can be tested with the standard library. Run:

```bash
go test ./mapbin/...
```

Recommended coverage:

- **Read** with valid input returns the correct count and items.
- **Read** with truncated input or failing reader returns an error.
- **Write** then **Read** round-trips to the same **MapBin** (including Unknown1–5 and names).
- **GetName** returns the name trimmed at the first null and handles empty or full names.
