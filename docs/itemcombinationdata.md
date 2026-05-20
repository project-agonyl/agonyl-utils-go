# ItemCombinationData Package

Documentation for the `github.com/project-agonyl/agonyl-utils-go/itemcombinationdata` package: read and write A3 `ItemCombinationData` crafting files.

---

## Overview

The `itemcombinationdata` package handles the binary crafting formula file used by A3/Agonyl server tooling.

- **Read** - reads a raw `ItemCombinationData` file from an `io.Reader`.
- **Write** - writes crafting formulas to an `io.Writer` in the same binary format.
- **CraftFormula** - one 32-byte crafting formula with 10 ingredient item IDs, success rate, outcome item ID, and four preserved unknown fields.

The file is a contiguous sequence of fixed-size records. There is no header and no count prefix.

---

## API Reference

### Constants

```go
const FormulaSize = 32
```

`FormulaSize` is the encoded byte size of one `CraftFormula`.

### Type: `CraftFormula`

```go
type CraftFormula struct {
    Item1       uint16
    Item2       uint16
    Item3       uint16
    Item4       uint16
    Item5       uint16
    Item6       uint16
    Item7       uint16
    Item8       uint16
    Item9       uint16
    Item10      uint16
    SuccessRate uint16
    Outcome     uint16
    Unknown1    uint16
    Unknown2    uint16
    Unknown3    uint16
    Unknown4    uint16
}
```

`Item1` through `Item10` are ingredient item IDs. Use `0` for empty item slots when that is how your game data represents unused inputs. `SuccessRate` and `Outcome` are stored exactly as provided. `Unknown1` through `Unknown4` are preserved raw fields.

### Type: `ItemCombinationData`

```go
type ItemCombinationData []CraftFormula
```

A slice of crafting formulas as stored in the file.

### Function: `Read`

```go
func Read(r io.Reader) (ItemCombinationData, error)
```

Reads all bytes from `r` and decodes them as little-endian `CraftFormula` records. Empty input returns a non-nil empty slice. If the stream length is not a multiple of 32 bytes, `Read` returns `io.ErrUnexpectedEOF`.

### Function: `Write`

```go
func Write(w io.Writer, data ItemCombinationData) error
```

Writes `data` to `w` in `ItemCombinationData` binary format. Nil and empty slices write zero bytes.

---

## Binary Format

Each formula is 32 bytes: 16 little-endian `uint16` fields.

| Offset | Field       | Size |
|--------|-------------|------|
| 0      | Item1       | 2    |
| 2      | Item2       | 2    |
| 4      | Item3       | 2    |
| 6      | Item4       | 2    |
| 8      | Item5       | 2    |
| 10     | Item6       | 2    |
| 12     | Item7       | 2    |
| 14     | Item8       | 2    |
| 16     | Item9       | 2    |
| 18     | Item10      | 2    |
| 20     | SuccessRate | 2    |
| 22     | Outcome     | 2    |
| 24     | Unknown1    | 2    |
| 26     | Unknown2    | 2    |
| 28     | Unknown3    | 2    |
| 30     | Unknown4    | 2    |

---

## Usage

### Read Crafting Data

```go
f, err := os.Open("ItemCombinationData")
if err != nil {
    log.Fatal(err)
}
defer f.Close()

data, err := itemcombinationdata.Read(f)
if err != nil {
    log.Fatal(err)
}

for i := range data {
    log.Printf("formula outcome=%d success=%d", data[i].Outcome, data[i].SuccessRate)
}
```

### Write Crafting Data

```go
data := itemcombinationdata.ItemCombinationData{
    {
        Item1:       100,
        Item2:       200,
        SuccessRate: 75,
        Outcome:     500,
    },
}

f, err := os.Create("ItemCombinationData")
if err != nil {
    log.Fatal(err)
}
defer f.Close()

if err := itemcombinationdata.Write(f, data); err != nil {
    log.Fatal(err)
}
```

---

## Testing

Run:

```bash
go test ./itemcombinationdata/...
```

Recommended coverage:

- Empty, single-formula, and multi-formula reads.
- Binary byte layout for all 16 fields.
- Truncated stream errors when the byte length is not a multiple of 32.
- Write then read round trips, including `Unknown1` through `Unknown4`.
