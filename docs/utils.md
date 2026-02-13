# Utils Package

Documentation for the `github.com/project-agonyl/agonyl-utils-go/utils` package: display-name helpers for character classes and nations used in the Agonyl protocol.

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

These helpers are intended for logging, UI labels, or debugging when working with protocol or game data that uses numeric class and nation identifiers.

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

---

## Testing

The package is tested with the standard library and [testify](https://github.com/stretchr/testify). Run tests with:

```bash
go test ./utils/...
```

Covered behavior includes:

- **GetClassName:** All defined classes (1–3) return the correct names; 0 and unknown values return "Warrior".
- **GetNationName:** Nation 1 returns "Quanato"; 0 and unknown values return "Temoz".

See `utils/character_test.go` and `utils/nation_test.go` for the test cases.
