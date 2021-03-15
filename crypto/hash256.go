// Copyright © 2020-2021 The EVEN Solutions Developers Team

package crypto

import (
	"crypto/sha256"
	"encoding/hex"
	"math/bits"

	"github.com/platsko/go-kit/base58"
)

const (
	// Hash256Size defines size in bytes.
	Hash256Size = sha256.Size
)

type (
	// Hash256 represents hashed blob with Hash256Size bytes length.
	Hash256 [Hash256Size]byte
)

// NewHash256 makes initialized Hash256.
// Initial hash calculates SHA256 checksum over specified bytes.
func NewHash256(data ...[]byte) (h256 Hash256) {
	size := 0
	for _, d := range data {
		size += len(d)
	}

	buf := make([]byte, 0, size)
	for _, blob := range data {
		buf = append(buf, blob...)
	}
	h256 = sha256.Sum256(buf)

	return h256
}

// StrToHash256 makes initialized Hash256.
// Initial hash calculates SHA256 checksum over specified strings.
func StrToHash256(str ...string) (h256 Hash256) {
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
	h256 = sha256.Sum256(buf)

	return h256
}

// Base58 returns Base58 encoded string over hashed bytes.
func (c Hash256) Base58() string {
	return base58.EncodeToString(c[:])
}

// Empty returns true if the hash is zeroed.
func (c Hash256) Empty() bool {
	for _, b := range c {
		if b != 0 {
			return false
		}
	}

	return true
}

// Encode returns hex-encoded string.
func (c Hash256) Encode() string {
	return hex.EncodeToString(c[:])
}

// Hamming returns distance between
// Hash256 bytes and specified vector the same size.
func (c Hash256) Hamming(v256 [Hash256Size]byte) (dist int) {
	for i := 0; i < Hash256Size; i++ {
		dist += bits.OnesCount8(c[i] ^ v256[i])
	}

	return dist
}

// String implements stringer interface.
func (c Hash256) String() string {
	return string(c[:])
}
