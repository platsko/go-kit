// Copyright © 2020-2021 The EVEN Solutions Developers Team

package base58

import (
	"crypto/sha256"
)

const (
	// ChecksumSize is a size of checksum in bytes.
	ChecksumSize = 4

	// VersionSize is a size of version in bytes.
	VersionSize = 1
)

// CheckEncode prepends a version byte and appends a four byte checksum.
func CheckEncode(b []byte, ver byte) string {
	blob := make([]byte, 0, VersionSize+len(b)+ChecksumSize)
	blob = append(blob, ver)
	blob = append(blob, b...)

	checksum := Checksum(blob)
	blob = append(blob, checksum[:]...)

	return EncodeToString(blob)
}

// CheckDecode decodes a string that was encoded with CheckEncode and verifies the checksum.
func CheckDecode(blob []byte) ([]byte, byte, error) {
	decoded, err := Decode(blob)
	if err != nil {
		return nil, 0, err
	}

	if len(decoded) < VersionSize+ChecksumSize {
		return nil, 0, ErrInvalidFormat()
	}

	size := len(decoded) - ChecksumSize
	checksum := [ChecksumSize]byte{}
	copy(checksum[:], decoded[size:])

	if Checksum(decoded[:size]) != checksum {
		return nil, 0, ErrChecksumMismatch()
	}

	payload := decoded[VersionSize:size]
	result := make([]byte, 0, len(payload))
	result = append(result, payload...)

	version := decoded[0]

	return result, version, nil
}

// Checksum returns first of ChecksumSize bytes
// of SHA256 double-hash over provided bytes.
func Checksum(blob []byte) (checksum [ChecksumSize]byte) {
	h256 := sha256.Sum256(blob)
	h256 = sha256.Sum256(h256[:])

	copy(checksum[:], h256[:ChecksumSize])

	return checksum
}
