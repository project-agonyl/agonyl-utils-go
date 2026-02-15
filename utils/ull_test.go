package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeULLRoundTrip_EncodeThenDecode(t *testing.T) {
	// Decode(Encode(plain)) must equal plain.
	plain := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}
	size := len(plain)
	buf := make([]byte, size)
	copy(buf, plain)

	EncodeULL(buf, size)
	DecodeULL(buf, size)

	assert.Equal(t, plain, buf, "Decode(Encode(plain)) should equal plain")
}

func TestDecodeULLRoundTrip_DecodeThenEncode(t *testing.T) {
	// Encode(Decode(encoded)) must equal encoded.
	encoded := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}
	size := len(encoded)
	buf := make([]byte, size)
	copy(buf, encoded)

	DecodeULL(buf, size)
	EncodeULL(buf, size)

	assert.Equal(t, encoded, buf, "Encode(Decode(encoded)) should equal encoded")
}

func TestDecodeULL_EncodeThenDecode_VariousSizes(t *testing.T) {
	sizes := []int{2, 3, 5, 10, 32, 64}
	for _, size := range sizes {
		plain := make([]byte, size)
		for i := range plain {
			plain[i] = byte(i * 7)
		}
		buf := make([]byte, size)
		copy(buf, plain)
		EncodeULL(buf, size)
		DecodeULL(buf, size)
		assert.Equal(t, plain, buf, "size=%d: Decode(Encode(plain)) should equal plain", size)
	}
}

func TestDecodeULL_DecodeThenEncode_VariousSizes(t *testing.T) {
	sizes := []int{2, 3, 5, 10, 32, 64}
	for _, size := range sizes {
		encoded := make([]byte, size)
		for i := range encoded {
			encoded[i] = byte(i * 11)
		}
		buf := make([]byte, size)
		copy(buf, encoded)
		DecodeULL(buf, size)
		EncodeULL(buf, size)
		assert.Equal(t, encoded, buf, "size=%d: Encode(Decode(encoded)) should equal encoded", size)
	}
}
