package hmac

import (
	"hash"
)

// Hmac calculate HMAC with given hash algorithm
func Hmac(h hash.Hash, key []byte, keyLen int, data []byte) []byte {
	key = InitializeKey(h, key, keyLen)
	ipad := Ipad(key, h.BlockSize())
	opad := Opad(key, h.BlockSize())

	h.Reset()
	h.Write(append(ipad, data...))
	hmac := h.Sum(nil)
	h.Reset()
	h.Write(append(opad, hmac...))
	hmac = h.Sum(nil)
	return hmac
}

// InitializeKey initializes secret key to be able to use it in HMAC
func InitializeKey(h hash.Hash, key []byte, keyLen int) []byte {
	if keyLen > h.BlockSize() { // make sure that Hash is initial state
		h.Reset()
		key = h.Sum(key)
	}
	// We don't need to check whether key length is less than block size
	// because when we create new slice with make(), it'll be initialized with zero-value
	initializedKey := make([]byte, h.BlockSize())
	copy(initializedKey, key)
	return initializedKey
}

// Ipad XOR key with 0x36
func Ipad(key []byte, blockSize int) []byte {
	ipad := make([]byte, blockSize)
	copy(ipad, key)
	for i := 0; i < blockSize; i++ {
		ipad[i] ^= 0x36
	}
	return ipad
}

// Opad XOR key with 0x5c
func Opad(key []byte, blockSize int) []byte {
	opad := make([]byte, blockSize)
	copy(opad, key)
	for i := 0; i < blockSize; i++ {
		opad[i] ^= 0x5c
	}
	return opad
}
