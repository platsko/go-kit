// Copyright © 2020-2021 The EVEN Solutions Developers Team

package base58

import (
	"crypto/sha256"
)

const (
	// checksumSize is a size of checksum in bytes.
	checksumSize = 4

	// versionSize is a size of version in bytes.
	versionSize = 1
)

// CheckEncode prepends a version byte and appends a four byte checksum.
func CheckEncode(b []byte, ver byte) string {
	blob := make([]byte, 0, versionSize+len(b)+checksumSize)
	blob = append(blob, ver)
	blob = append(blob, b...)

	checksum := Checksum(blob)
	blob = append(blob, checksum[:]...)

	return EncodeToString(blob)
}

// CheckDecode decodes a string that was encoded with CheckEncode and verifies the checksum.
func CheckDecode(b []byte) ([]byte, byte, error) {
	decoded, err := Decode(b)
	if err != nil {
		return nil, 0, err
	}

	if len(decoded) < versionSize+checksumSize {
		return nil, 0, ErrInvalidFormat()
	}

	size := len(decoded) - checksumSize
	checksum := [checksumSize]byte{}
	copy(checksum[:], decoded[size:])

	if Checksum(decoded[:size]) != checksum {
		return nil, 0, ErrChecksumMismatch()
	}

	payload := decoded[versionSize:size]
	result := make([]byte, 0, len(payload))
	result = append(result, payload...)

	version := decoded[0]

	return result, version, nil
}

// Checksum returns first of checksumSize bytes
// of SHA256 double-hash over provided bytes.
func Checksum(b []byte) (checksum [checksumSize]byte) {
	h := sha256.Sum256(b)
	h = sha256.Sum256(h[:])

	copy(checksum[:], h[:checksumSize])

	return checksum
}
