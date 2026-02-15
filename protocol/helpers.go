// Package protocol provides message encoding and decoding helpers for the protocol layer.
package protocol

import (
	"bytes"
	"encoding/binary"
)

// GetBytesFromMsg serializes v into a byte slice using little-endian binary encoding.
// It is intended for protocol messages and structs that are safe to encode with encoding/binary.
// Returns the encoded bytes and any error from binary.Write.
func GetBytesFromMsg(v any) ([]byte, error) {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.LittleEndian, v)
	return buf.Bytes(), err
}

// ReadMsgFromBytes decodes data into v using little-endian binary encoding.
// The value v must be a pointer to a type that binary.Read supports (e.g. a struct or fixed-size type).
// Returns any error from binary.Read.
func ReadMsgFromBytes(data []byte, v any) error {
	return binary.Read(bytes.NewReader(data), binary.LittleEndian, v)
}
