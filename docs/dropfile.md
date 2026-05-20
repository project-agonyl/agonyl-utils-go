# Dropfile Package

Documentation for the `github.com/project-agonyl/agonyl-utils-go/dropfile` package: read and write A3 monster `.itm` drop files.

---

## Overview

The `dropfile` package handles the binary monster drop files used by A3/Agonyl server tooling.

- **Read** - reads a raw monster `.itm` drop file from an `io.Reader`.
- **Write** - writes monster drops to an `io.Writer` in the same binary format.
- **Drop** - one 6-byte drop record with item ID, drop rate, and drop group.

The file is a contiguous sequence of fixed-size records. There is no header and no count prefix.

Item names are not stored in the file. AgonylDropEditor resolves names from item text files by using `ItemID & 0x3FFF`, and treats `0xFFFF` as an empty drop slot.

---

## API Reference

### Constants

```go
const DropSize = 6
const EmptyItemID uint16 = 0xFFFF
const ItemIDMask uint16 = 0x3FFF
```

`DropSize` is the encoded byte size of one `Drop`.

`EmptyItemID` is the item ID used for an empty drop slot.

`ItemIDMask` strips item flag bits before item-name lookup.

### Type: `Drop`

```go
type Drop struct {
    ItemID    uint16
    DropRate  uint16
    DropGroup uint16
}
```

`ItemID` is stored exactly as provided, including flag bits and `0xFFFF`. Use `ItemID & ItemIDMask` when you need the base item ID for lookup. `DropRate` and `DropGroup` are stored exactly as provided.

### Type: `DropFile`

```go
type DropFile []Drop
```

A slice of drop records as stored in the file.

### Function: `Read`

```go
func Read(r io.Reader) (DropFile, error)
```

Reads all bytes from `r` and decodes them as little-endian `Drop` records. Empty input returns a non-nil empty slice. If the stream length is not a multiple of 6 bytes, `Read` returns `io.ErrUnexpectedEOF`.

### Function: `Write`

```go
func Write(w io.Writer, data DropFile) error
```

Writes `data` to `w` in monster `.itm` drop file format. Nil and empty slices write zero bytes.

---

## Binary Format

Each drop is 6 bytes: 3 little-endian `uint16` fields.

| Offset | Field     | Size |
|--------|-----------|------|
| 0      | ItemID    | 2    |
| 2      | DropRate  | 2    |
| 4      | DropGroup | 2    |

---

## Usage

### Read Drop Data

```go
f, err := os.Open("100.itm")
if err != nil {
    log.Fatal(err)
}
defer f.Close()

drops, err := dropfile.Read(f)
if err != nil {
    log.Fatal(err)
}

for i := range drops {
    itemID := drops[i].ItemID
    if itemID == dropfile.EmptyItemID {
        log.Printf("drop slot %d is empty", i)
        continue
    }

    baseItemID := itemID & dropfile.ItemIDMask
    log.Printf("drop item=%d rate=%d group=%d",
        baseItemID, drops[i].DropRate, drops[i].DropGroup)
}
```

### Write Drop Data

```go
drops := dropfile.DropFile{
    {ItemID: 100, DropRate: 75, DropGroup: 1},
    {ItemID: dropfile.EmptyItemID},
}

f, err := os.Create("100.itm")
if err != nil {
    log.Fatal(err)
}
defer f.Close()

if err := dropfile.Write(f, drops); err != nil {
    log.Fatal(err)
}
```

---

## Testing

Run:

```bash
go test ./dropfile/...
```

Recommended coverage:

- Empty, single-drop, and multi-drop reads.
- Binary byte layout for all 3 fields.
- Truncated stream errors when the byte length is not a multiple of 6.
- Write then read round trips, including flagged item IDs and `0xFFFF`.
- Constants for empty slots and item lookup masking.
