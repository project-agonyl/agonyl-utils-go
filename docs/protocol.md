# Protocol

This document describes how to convert between protocol messages and byte slices using the `protocol` package helpers.

## Helper functions

- **GetBytesFromMsg** — serialize a message (or any encodable value) to a byte slice.
- **ReadMsgFromBytes** — deserialize a byte slice into a message (or any decodable value).

Encoding and decoding use **little-endian** binary format via `encoding/binary`. Use these helpers with fixed-size structs and types that `binary.Write` / `binary.Read` support (e.g. fixed-size arrays, numeric types, structs composed of such fields). Slices, maps, and strings are not supported by the binary package.

---

## Message to bytes (GetBytesFromMsg)

Use `GetBytesFromMsg` when you have a protocol message (or struct) and need its byte representation for sending over the wire or further processing.

**Signature:**

```go
func GetBytesFromMsg(v any) ([]byte, error)
```

**Example:**

```go
import "your-module/protocol"

msg := protocol.MsgHead{
    MsgHeadNoProtocol: protocol.MsgHeadNoProtocol{
        Size: 12,
        PcId: 1,
        Ctrl: 0x03,
        Cmd:  0xFF,
    },
    Protocol: protocol.C2SSay,
}

data, err := protocol.GetBytesFromMsg(msg)
if err != nil {
    // handle error (e.g. unsupported type)
    return err
}
// use data (e.g. send over network, write to buffer)
```

Pass the **value** (e.g. struct or value type). The function returns the encoded bytes and an error if encoding fails.

---

## Bytes to message (ReadMsgFromBytes)

Use `ReadMsgFromBytes` when you have a byte slice (e.g. from the network or a buffer) and want to decode it into a protocol message.

**Signature:**

```go
func ReadMsgFromBytes(data []byte, v any) error
```

**Important:** `v` must be a **pointer** to the type you want to decode into (e.g. `*MsgHead`), so that `binary.Read` can write the decoded fields into it.

**Example:**

```go
import "your-module/protocol"

data := []byte{...} // e.g. received from network

var msg protocol.MsgHead
err := protocol.ReadMsgFromBytes(data, &msg)
if err != nil {
    // handle error (e.g. not enough bytes, type mismatch)
    return err
}
// use msg
```

For a message type that includes a size field, ensure `data` has at least as many bytes as the message expects; otherwise `ReadMsgFromBytes` may return an error or fill only part of the struct.

---

## Round-trip example

```go
// Encode: message → bytes
msg := protocol.MsgHead{ /* ... */ }
data, err := protocol.GetBytesFromMsg(msg)
if err != nil {
    return err
}

// Decode: bytes → message
var decoded protocol.MsgHead
err = protocol.ReadMsgFromBytes(data, &decoded)
if err != nil {
    return err
}
// decoded now holds the same values as msg (for the bytes that were written)
```
