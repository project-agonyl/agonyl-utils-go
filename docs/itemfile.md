# Itemfile Package

Documentation for the `github.com/project-agonyl/agonyl-utils-go/itemfile` package: read, write, and parse A3 item binary files.

---

## Overview

The `itemfile` package provides helpers for the item file families used by the A3/Agonyl client and server tooling:

- **Raw read/write** helpers for IT0, IT0Ex, IT1, IT2, and IT3 files.
- **Parsed item** helpers that convert raw records into upstream-style `Item` values.
- **Name helpers** on raw records with fixed 32-byte names.

Raw item files are stored as contiguous little-endian records with no count prefix.

---

## API Reference

### Raw Types

```go
type IT0File []IT0Raw
type IT0ExFile []IT0ExRaw
type IT1File []IT1Raw
type IT2File []IT2Raw
type IT3File []IT3Raw
```

Each raw file type is a slice of fixed-size binary records.

### Parsed Types

```go
type Item struct {
    ItemCode    uint32
    SlotIndex   byte
    ItemName    string
    Itemtype    byte
    NPCPrice    uint32
    IT0Property *IT0Property
    IT1Property *IT1Property
    IT2Property *IT2Property
}
```

`Item` is the parsed representation shared by IT0, IT1, IT2, and IT3 records.

### Raw Read/Write

```go
func ReadIT0(r io.Reader) (IT0File, error)
func WriteIT0(w io.Writer, data IT0File) error

func ReadIT0Ex(r io.Reader) (IT0ExFile, error)
func WriteIT0Ex(w io.Writer, data IT0ExFile) error

func ReadIT1(r io.Reader) (IT1File, error)
func WriteIT1(w io.Writer, data IT1File) error

func ReadIT2(r io.Reader) (IT2File, error)
func WriteIT2(w io.Writer, data IT2File) error

func ReadIT3(r io.Reader) (IT3File, error)
func WriteIT3(w io.Writer, data IT3File) error
```

Read helpers consume the full stream and return `io.ErrUnexpectedEOF` if the byte length is not a multiple of the record size. Write helpers write the records exactly as provided.

### Parsed Helpers

```go
func ParseIT0Items(it0 IT0File, it0Ex IT0ExFile) ([]Item, error)
func ParseIT1Items(it1 IT1File) ([]Item, error)
func ParseIT2Items(it2 IT2File) ([]Item, error)
func ParseIT3Items(it3 IT3File) ([]Item, error)

func ReadIT0Items(it0 io.Reader, it0Ex io.Reader) ([]Item, error)
func ReadIT1Items(r io.Reader) ([]Item, error)
func ReadIT2Items(r io.Reader) ([]Item, error)
func ReadIT3Items(r io.Reader) ([]Item, error)
```

Parsed helpers compute item codes, slot indexes, names, NPC prices, and per-type property structures. They return an error when a row index is outside the parsed item slice or when an IT0Ex row references a missing IT0 base item.

---

## Binary Format

| File  | Record Type | Record Size | Count Prefix |
|-------|-------------|-------------|--------------|
| IT0   | `IT0Raw`    | 242 bytes   | No           |
| IT0Ex | `IT0ExRaw`  | 92 bytes    | No           |
| IT1   | `IT1Raw`    | 52 bytes    | No           |
| IT2   | `IT2Raw`    | 48 bytes    | No           |
| IT3   | `IT3Raw`    | 48 bytes    | No           |

All integer fields are little-endian.

---

## Usage

### Read Raw IT1 Data

```go
f, err := os.Open("it1.bin")
if err != nil {
    log.Fatal(err)
}
defer f.Close()

raw, err := itemfile.ReadIT1(f)
if err != nil {
    log.Fatal(err)
}

for i := range raw {
    log.Printf("raw item %d: %s", i, raw[i].GetName())
}
```

### Read Parsed IT0 Items

```go
it0, err := os.Open("it0.bin")
if err != nil {
    log.Fatal(err)
}
defer it0.Close()

it0Ex, err := os.Open("it0ex.bin")
if err != nil {
    log.Fatal(err)
}
defer it0Ex.Close()

items, err := itemfile.ReadIT0Items(it0, it0Ex)
if err != nil {
    log.Fatal(err)
}

for i := range items {
    log.Printf("item code=%d name=%q", items[i].ItemCode, items[i].ItemName)
}
```

### Write Raw IT3 Data

```go
data := itemfile.IT3File{
    {Type: 7, Row: 0, NPCPrice: 100},
}
copy(data[0].Name[:], "Quest Item")

f, err := os.Create("it3.bin")
if err != nil {
    log.Fatal(err)
}
defer f.Close()

if err := itemfile.WriteIT3(f, data); err != nil {
    log.Fatal(err)
}
```

---

## Testing

Run:

```bash
go test ./itemfile/...
```

Recommended coverage:

- Raw read/write round trips for every item file type.
- Truncated stream errors when byte length is not a record multiple.
- Parsed item code, name, type, NPC price, slot index, and property mappings.
- Bad row indexes return errors instead of panics.
