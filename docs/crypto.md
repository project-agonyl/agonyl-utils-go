# Crypto Package

Comprehensive documentation for the `github.com/project-agonyl/agonyl-utils-go/crypto` package: a small, in-place encryption/decryption library for the **562 variant** stream cipher.

---

## Table of Contents

- [Overview](#overview)
- [Installation](#installation)
- [Cipher Details](#cipher-details)
- [API Reference](#api-reference)
- [Usage](#usage)
- [Data Layout and Constraints](#data-layout-and-constraints)
- [Edge Cases and Behavior](#edge-cases-and-behavior)
- [Security and Key Management](#security-and-key-management)
- [Examples](#examples)
- [Testing](#testing)

---

## Overview

The `crypto` package provides:

- **In-place** encryption and decryption: the given byte slice is modified directly; no new buffer is allocated.
- A **stream-cipher-style** algorithm (562 variant) that operates on 4-byte blocks starting at a fixed offset.
- A single constructor, **NewCrypto562**, which takes a **dynamic key** used to seed the cipher state. The same key must be used for both encrypt and decrypt to get a correct round-trip.

Typical use cases include protocol payloads or packet bodies where a 12-byte header is left in the clear and only the remainder is encrypted (e.g. game or legacy protocol compatibility).

---

## Installation

```bash
go get github.com/project-agonyl/agonyl-utils-go
```

Import in your code:

```go
import "github.com/project-agonyl/agonyl-utils-go/crypto"
```

---

## Cipher Details

- **Name / variant:** 562 (A3 client v562-style).
- **Operation:** XOR with a byte derived from an internal key, then key update using linear congruential-style constants. Process is symmetric for encrypt vs decrypt when the same dynamic key is used.
- **Offset:** Only bytes at index **0x0C (12)** and beyond are touched. Bytes `data[0:12]` are never modified.
- **Block size:** 4 bytes. The cipher advances in 4-byte blocks from the offset. If the length after the offset is not a multiple of 4, the last incomplete block is left unchanged.
- **Minimum length for any transformation:** `offset + 4` = **16 bytes**. Shorter slices are left unchanged.

---

## API Reference

### Type: `Crypto`

```go
type Crypto interface {
    EncryptInPlace(data []byte)
    DecryptInPlace(data []byte)
}
```

- **EncryptInPlace(data)** – Encrypts `data` in place. Only bytes from index `0x0C` onward are modified, in 4-byte blocks. Slice length is unchanged.
- **DecryptInPlace(data)** – Decrypts `data` in place under the same rules.

Both methods may be called with the same or different slices; they do not keep internal per-call state beyond what is implied by the current `data` and the instance’s dynamic key.

### Constructor: `NewCrypto562`

```go
func NewCrypto562(dynamicKey int) Crypto
```

- **dynamicKey:** Integer used to seed the cipher. Must be the same for encryption and decryption of the same data.
- **Returns:** A non-nil `Crypto` implementation (562 cipher). Safe for concurrent use from multiple goroutines if each goroutine uses its own instance or access is synchronized.

---

## Usage

### Basic encrypt / decrypt round-trip

1. Create a `Crypto` with a chosen dynamic key.
2. Copy your plaintext into a byte slice (or use the slice you want to overwrite).
3. Call `EncryptInPlace` on that slice.
4. To decrypt, call `DecryptInPlace` on the same (or an identical) slice with the **same** `Crypto` instance (same dynamic key).

```go
c := crypto.NewCrypto562(0x1234)
payload := []byte{/* at least 16 bytes for one block */}
// ... fill payload ...

c.EncryptInPlace(payload)  // payload is now encrypted in place
c.DecryptInPlace(payload)  // payload is back to plaintext
```

### Preserving the original buffer

Because both functions **modify the slice in place**, pass a copy if you need to keep the original:

```go
original := []byte{0x00, 0x01, /* ... */}
working := make([]byte, len(original))
copy(working, original)
c.EncryptInPlace(working)  // only working is changed; original is unchanged
```

### Using the same key for many packets

Use one `Crypto` instance per logical stream/session so the same dynamic key is applied consistently:

```go
sessionKey := deriveKeyFromSession(sessionID)
c := crypto.NewCrypto562(sessionKey)
for _, packet := range packets {
    c.DecryptInPlace(packet.Payload)  // or EncryptInPlace for sending
}
```

---

## Data Layout and Constraints

| Region              | Index range     | Modified by cipher? |
|---------------------|-----------------|----------------------|
| Header / untouched  | `0` to `0x0B`   | No                   |
| Payload             | `0x0C` onward   | Yes, in 4-byte blocks |

- **Length &lt; 16:** No bytes are modified (no full block after offset).
- **Length 16:** Exactly one 4-byte block is encrypted/decrypted (indices 12–15).
- **Length 20:** Two blocks (12–15, 16–19).
- **Length 24:** Three blocks, and so on. Any trailing bytes that do not form a full 4-byte block are left unchanged.

Empty slices are valid: no modification and no panic.

---

## Edge Cases and Behavior

| Input              | Behavior |
|--------------------|----------|
| `nil` slice        | Allowed; no-op (length 0). |
| Empty slice `[]byte{}` | No-op. |
| Length 1–11        | No bytes modified. |
| Length 12–15       | No full block; no bytes modified. |
| Length ≥ 16        | From index 12, full 4-byte blocks are processed; remainder bytes unchanged. |
| Same slice twice  | Encrypt then decrypt (or vice versa) with the same key restores original. |
| Different keys     | Encrypt with key A and decrypt with key B does not restore plaintext. |

The implementation does not change the **length** of the slice, only the contents from offset `0x0C` onward.

---

## Security and Key Management

- **Key agreement:** The dynamic key must be shared or derived the same way on both sides (e.g. from session ID, handshake, or protocol-specific rules). The package does not define how to derive or exchange this key.
- **Key type:** The key is an `int`; typically only the low 32 bits (or fewer) affect the cipher output. Use a consistent width if you exchange keys across systems.
- **Algorithm:** This is a custom 562-variant stream cipher for compatibility, not a modern authenticated cipher. Do not rely on it for new security-sensitive designs without a separate integrity/authentication mechanism if required.
- **Concurrency:** A single `Crypto` instance is safe for concurrent use only if callers do not pass the same slice to multiple goroutines at once. Different goroutines can use the same instance with different slices.

---

## Examples

### Encrypt a packet payload (preserve header)

```go
packet := make([]byte, 12+bodyLen) // 12-byte header + body
copy(packet[12:], body)
c := crypto.NewCrypto562(packetKey)
c.EncryptInPlace(packet)  // only packet[12:] is modified
send(packet)
```

### Decrypt and then process

```go
raw := receive()
c := crypto.NewCrypto562(sessionKey)
c.DecryptInPlace(raw)
processPayload(raw[0x0C:])
```

### Round-trip with copy (keep original)

```go
plain := []byte("sensitive data that must be 16+ bytes!!")
buf := make([]byte, len(plain))
copy(buf, plain)
c := crypto.NewCrypto562(0xCAFE)
c.EncryptInPlace(buf)
// ... send or store buf ...
c.DecryptInPlace(buf)
// buf now equals plain again; plain was never modified
```

---

## Testing

The package is tested with the standard library and [testify](https://github.com/stretchr/testify). Run tests with:

```bash
go test ./crypto/...
```

Covered behavior includes:

- **NewCrypto562** returns a non-nil implementation of `Crypto`.
- **Encrypt then decrypt** (and decrypt then encrypt) round-trips to the original bytes.
- **In-place:** Slice length is unchanged; only content from the offset is modified.
- **Short data:** Slices shorter than 16 bytes are left unchanged.
- **Empty slice:** No panic.
- **Different dynamic keys** produce different ciphertext for the same plaintext.
- **Multiple 4-byte blocks** round-trip correctly.

See `crypto/crypto_test.go` for the exact test cases and usage patterns.
