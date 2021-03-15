// Copyright © 2020 The Evenlab Team

package crypto

import (
	"crypto/sha256"
	"encoding/hex"
	"math/bits"

	"github.com/platsko/go-kit/base58"
)

const (
	// Hash224Size defines size in bytes.
	Hash224Size = sha256.Size224
)

type (
	// Hash224 represents hashed blob with Hash224Size bytes length.
	Hash224 [Hash224Size]byte
)

// NewHash224 makes initialized Hash224.
// Initial hash calculates SHA224 checksum over SHA256 checksum over specified bytes.
func NewHash224(data ...[]byte) (h224 Hash224) {
	size := 0
	for _, d := range data {
		size += len(d)
	}

	buf := make([]byte, 0, size)
	for _, blob := range data {
		buf = append(buf, blob...)
	}

	h256 := sha256.Sum256(buf)
	h224 = sha256.Sum224(h256[:])

	return h224
}

// StrToHash224 makes initialized Hash224.
// Initial hash calculates SHA224 checksum SHA256 checksum over specified strings.
func StrToHash224(str ...string) (h224 Hash224) {
	size := 0
	for _, s := range str {
		size += len(s)
	}

	idx, buf := 0, make([]byte, size)
	for _, s := range str {
		for i, l := 0, len(s); i < l; i, idx = i+1, idx+1 {
			buf[idx] = s[i]
		}
	}

	h256 := sha256.Sum256(buf)
	h224 = sha256.Sum224(h256[:])

	return h224
}

// Base58 returns Base58 encoded string over hashed bytes.
func (c Hash224) Base58() string {
	return base58.EncodeToString(c[:])
}

// Empty returns true if the hash is zeroed.
func (c Hash224) Empty() bool {
	for _, b := range c {
		if b != 0 {
			return false
		}
	}

	return true
}

// Encode returns hex-encoded string.
func (c Hash224) Encode() string {
	return hex.EncodeToString(c[:])
}

// Hamming returns distance between
// Hash224 bytes and specified vector the same size.
func (c Hash224) Hamming(v224 [Hash224Size]byte) (dist int) {
	for i := 0; i < Hash224Size; i++ {
		dist += bits.OnesCount8(c[i] ^ v224[i])
	}

	return dist
}

// String implements stringer interface.
func (c Hash224) String() string {
	return string(c[:])
}
