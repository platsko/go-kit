// Copyright © 2020-2021 The EVEN Solutions Developers Team

package zero

// Bytes sets all bytes in the passed slice to zero.
// This is used to explicitly clear data from memory.
func Bytes(b []byte) {
	for idx := range b {
		b[idx] = 0
	}
}

// Bytea32 clears the 32-byte array by filling it with the zero value.
// This is used to explicitly clear private key material from memory.
func Bytea32(b *[32]byte) {
	*b = [32]byte{}
}

// Bytea64 clears the 64-byte array by filling it with the zero value.
// This is used to explicitly clear sensitive material from memory.
func Bytea64(b *[64]byte) {
	*b = [64]byte{}
}
