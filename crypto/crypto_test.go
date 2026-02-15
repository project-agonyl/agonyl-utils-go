package crypto

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewCrypto562(t *testing.T) {
	c := NewCrypto562(12345)
	require.NotNil(t, c, "NewCrypto562 should return non-nil Crypto")
	assert.Implements(t, (*Crypto)(nil), c, "NewCrypto562 should return a Crypto implementation")
}

func TestEncryptDecryptRoundTrip(t *testing.T) {
	c := NewCrypto562(0x1234)
	// Need at least offset + 4 bytes for one full block (offset is 0x0C = 12, so 16+ bytes)
	original := []byte{
		0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
		0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F,
		0x10, 0x11, 0x12, 0x13,
	}
	plain := make([]byte, len(original))
	copy(plain, original)

	c.EncryptInPlace(plain)
	assert.NotEqual(t, original, plain, "EncryptInPlace should mutate data")

	c.DecryptInPlace(plain)
	assert.Equal(t, original, plain, "Decrypt after Encrypt should restore original data")
}

func TestDecryptEncryptRoundTrip(t *testing.T) {
	c := NewCrypto562(999)
	original := []byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0xAA, 0xBB, 0xCC, 0xDD,
	}
	data := make([]byte, len(original))
	copy(data, original)

	c.DecryptInPlace(data)
	assert.NotEqual(t, original, data, "DecryptInPlace should mutate data")

	c.EncryptInPlace(data)
	assert.Equal(t, original, data, "Encrypt after Decrypt should restore original data")
}

func TestEncryptInPlace_ModifiesInPlace(t *testing.T) {
	c := NewCrypto562(42)
	data := make([]byte, 20)
	for i := range data {
		data[i] = byte(i)
	}
	origLen := len(data)

	c.EncryptInPlace(data)
	assert.Len(t, data, origLen, "EncryptInPlace must not change slice length")
	assert.NotEqual(t, bytes.Repeat([]byte{0}, 20), data, "EncryptInPlace should modify payload bytes")
}

func TestEncryptInPlace_ShortDataUnchanged(t *testing.T) {
	c := NewCrypto562(100)
	// Length 12 = exactly offset, no full 4-byte block; nothing is processed
	short := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C}
	backup := make([]byte, len(short))
	copy(backup, short)

	c.EncryptInPlace(short)
	assert.Equal(t, backup, short, "Data shorter than offset+4 should be left unchanged")
}

func TestDifferentDynamicKeys_DifferentCiphertext(t *testing.T) {
	plain := []byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x11, 0x22, 0x33, 0x44,
	}
	c1 := NewCrypto562(1)
	c2 := NewCrypto562(2)

	p1 := make([]byte, len(plain))
	p2 := make([]byte, len(plain))
	copy(p1, plain)
	copy(p2, plain)

	c1.EncryptInPlace(p1)
	c2.EncryptInPlace(p2)
	assert.NotEqual(t, p1, p2, "Different dynamic keys should produce different ciphertext")
}

func TestEncryptInPlace_EmptySlice(t *testing.T) {
	c := NewCrypto562(0)
	data := []byte{}
	require.NotPanics(t, func() { c.EncryptInPlace(data) })
	assert.Empty(t, data)
}

func TestDecryptInPlace_EmptySlice(t *testing.T) {
	c := NewCrypto562(0)
	data := []byte{}
	require.NotPanics(t, func() { c.DecryptInPlace(data) })
	assert.Empty(t, data)
}

func TestEncryptInPlace_MultipleBlocks(t *testing.T) {
	c := NewCrypto562(0xDEAD)
	// 12 (offset) + 3*4 = 24 bytes of payload = 3 blocks
	original := make([]byte, 24)
	for i := range original {
		original[i] = byte(i * 7)
	}
	data := make([]byte, len(original))
	copy(data, original)

	c.EncryptInPlace(data)
	assert.NotEqual(t, original, data)

	c.DecryptInPlace(data)
	assert.Equal(t, original, data, "Round-trip should hold for multiple blocks")
}
