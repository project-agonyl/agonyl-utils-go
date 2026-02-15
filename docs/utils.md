# Utils Package

Documentation for the `github.com/project-agonyl/agonyl-utils-go/utils` package: display-name helpers for character classes and nations, and ULL (A3 client data file) encode/decode used in the Agonyl protocol.

---

## Table of Contents

- [Overview](#overview)
- [Installation](#installation)
- [API Reference](#api-reference)
- [Usage](#usage)
- [Testing](#testing)

---

## Overview

The `utils` package provides:

- **GetClassName** — maps a character class ID (byte) to its display name (e.g. Holy Knight, Mage, Archer, Warrior).
- **GetNationName** — maps a nation ID (byte) to its display name (Quanato or Temoz).
- **EncodeULL** / **DecodeULL** — in-place XOR encode/decode for ULL (A3 client data file) byte buffers using a fixed lookup table.

The display-name helpers are intended for logging, UI labels, or debugging when working with protocol or game data that uses numeric class and nation identifiers. ULL encode/decode is used when reading or writing ULL-formatted data (e.g. client data files) in the Agonyl/A3 context.

---

## Installation

```bash
go get github.com/project-agonyl/agonyl-utils-go
```

Import in your code:

```go
import "github.com/project-agonyl/agonyl-utils-go/utils"
```

---

## API Reference

### GetClassName

```go
func GetClassName(class byte) string
```

Returns the display name for the given character class ID.

| Class (byte) | Name        |
|--------------|-------------|
| 1            | Holy Knight |
| 2            | Mage        |
| 3            | Archer      |
| 0, 4–255     | Warrior     |

---

### GetNationName

```go
func GetNationName(nation byte) string
```

Returns the display name for the given nation ID.

| Nation (byte) | Name    |
|---------------|---------|
| 1             | Quanato |
| 0, 2–255      | Temoz   |

---

### EncodeULL / DecodeULL

In-place XOR transformation for ULL (A3 client data file) buffers. The two functions are inverses: `Decode(Encode(buf))` and `Encode(Decode(buf))` restore the buffer.

```go
func DecodeULL(buffer []byte, size int)
func EncodeULL(buffer []byte, size int)
```

- **buffer** — slice to transform; only the first `size` bytes are read/written. Modified in place.
- **size** — number of bytes to process; must be in `[0, len(buffer)]`.

Decode processes bytes from high index to low (right to left); Encode processes low to high (left to right) so each step uses the already-encoded value at the previous index.

---

## Usage

### Display class and nation in logs or UI

```go
classID := byte(2)
nationID := byte(1)

className := utils.GetClassName(classID)   // "Mage"
nationName := utils.GetNationName(nationID) // "Quanato"

log.Printf("Character: %s from %s", className, nationName)
```

### Default for unknown IDs

Unknown or reserved IDs map to a sensible default (Warrior for class, Temoz for nation), so you can safely pass any byte:

```go
name := utils.GetClassName(255)  // "Warrior"
nation := utils.GetNationName(0) // "Temoz"
```

### ULL encode/decode

Decode received ULL data, or encode before sending:

```go
data := []byte{0x01, 0x02, 0x03, 0x04, 0x05}
size := len(data)

// Decode (e.g. after reading from file/network)
utils.DecodeULL(data, size)

// ... use decoded data ...

// Encode again (e.g. before writing)
utils.EncodeULL(data, size)
```

Round-trip preserves content: `Decode(Encode(buf))` and `Encode(Decode(buf))` leave the buffer unchanged.

---

## Testing

The package is tested with the standard library and [testify](https://github.com/stretchr/testify). Run tests with:

```bash
go test ./utils/...
```

Covered behavior includes:

- **GetClassName:** All defined classes (1–3) return the correct names; 0 and unknown values return "Warrior".
- **GetNationName:** Nation 1 returns "Quanato"; 0 and unknown values return "Temoz".
- **EncodeULL / DecodeULL:** Round-trip tests: `Decode(Encode(plain)) == plain` and `Encode(Decode(encoded)) == encoded` for various buffer sizes.

See `utils/character_test.go`, `utils/nation_test.go`, and `utils/ull_test.go` for the test cases.
