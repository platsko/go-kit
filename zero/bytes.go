// Copyright © 2020-2021 The EVEN Solutions Developers Team

package zero

// Bytea28 clears the 28-byte (224 bits) array by filling it with the zero value.
// This is used to explicitly clear private key material from memory.
func Bytea28(b *[28]byte) {
	*b = [28]byte{}
}

// Bytea32 clears the 32-byte (256 bits) array by filling it with the zero value.
// This is used to explicitly clear private key material from memory.
func Bytea32(b *[32]byte) {
	*b = [32]byte{}
}

// Bytea64 clears the 64-byte (512 bits) array by filling it with the zero value.
// This is used to explicitly clear sensitive material from memory.
func Bytea64(b *[64]byte) {
	*b = [64]byte{}
}

// Bytes sets all bytes in the passed slice to zero.
// This is used to explicitly clear data from memory.
func Bytes(b []byte) {
	for idx := range b {
		b[idx] = 0
	}
}
