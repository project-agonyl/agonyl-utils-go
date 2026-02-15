// Package crypto provides in-place encryption and decryption utilities
// for a custom stream cipher (562 variant).
package crypto

// Crypto is the interface for in-place encryption and decryption.
// Implementations mutate the given byte slice; callers should pass a copy
// if the original data must be preserved.
type Crypto interface {
	EncryptInPlace(data []byte)
	DecryptInPlace(data []byte)
}

// crypto562 holds constant and dynamic keys for the A3 client v562 cipher.
type crypto562 struct {
	constKey1   int
	constKey2   int
	dynamicKey  int
	dynamicKey1 byte
	dynamicKey2 byte
	constKeyEn  uint32
	constKeyDe  uint32
}

// offset is the starting point for the 562 cipher.
const offset = 0x0C

// NewCrypto562 returns a Crypto implementation using the 562 cipher with
// the given dynamic key. The dynamic key is typically derived from
// session or packet context and must match between encrypt and decrypt.
func NewCrypto562(dynamicKey int) Crypto {
	return &crypto562{
		constKey1:   0x241AE7,
		constKey2:   0x15DCB2,
		dynamicKey:  dynamicKey,
		dynamicKey1: 0x02,
		dynamicKey2: 0x01,
		constKeyEn:  0xA7F0753B,
		constKeyDe:  0xAAF29BF3,
	}
}

// DecryptInPlace decrypts data in place using the 562 cipher.
// Only bytes from offset onward are modified, in 4-byte blocks.
// Data is modified in place; the slice length is unchanged.
func (c *crypto562) DecryptInPlace(data []byte) {
	bufferLen := len(data)
	for i := offset; i+4 <= bufferLen; i += 4 {
		DynamicKey := c.dynamicKey
		for j := i; j < i+4; j++ {
			pSrc := data[j]
			data[j] = pSrc ^ byte(DynamicKey>>8)
			DynamicKey = (int(pSrc)+DynamicKey)*c.constKey1 + c.constKey2
		}
	}
}

// EncryptInPlace encrypts data in place using the 562 cipher.
// Only bytes from offset onward are modified, in 4-byte blocks.
// Data is modified in place; the slice length is unchanged.
func (c *crypto562) EncryptInPlace(data []byte) {
	bufferLen := len(data)
	for i := offset; i+4 <= bufferLen; i += 4 {
		DynamicKey := c.dynamicKey
		for j := i; j < i+4; j++ {
			data[j] = data[j] ^ byte(DynamicKey>>8)
			DynamicKey = (int(data[j])+DynamicKey)*c.constKey1 + c.constKey2
		}
	}
}
