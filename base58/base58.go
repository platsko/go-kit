// Copyright © 2020-2021 The EVEN Solutions Developers Team

package base58

import (
	"unicode/utf8"
)

const (
	// Alphabet is the modified base58 alphabet.
	Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

	alphabetIdx0 = '1'
	alphabetSize = len(Alphabet)

	// i255 is a magic number represented symbol that does not exist in Alphabet.
	i255 = 255
)

// Decode decodes base58 encoded bytes.
func Decode(blob []byte) ([]byte, error) {
	// nolint: gofmt, goimports
	decodeTable := [256]byte{
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255,   0,   1,   2,   3,   4,   5,   6,   7,   8, 255, 255, 255, 255, 255, 255,
		255,   9,  10,  11,  12,  13,  14,  15,  16, 255,  17,  18,  19,  20,  21, 255,
		 22,  23,  24,  25,  26,  27,  28,  29,  30,  31,  32, 255, 255, 255, 255, 255,
		255,  33,  34,  35,  36,  37,  38,  39,  40,  41,  42,  43, 255,  44,  45,  46,
		 47,  48,  49,  50,  51,  52,  53,  54,  55,  56,  57, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	}

	// nolint: gomnd
	capacity := utf8.RuneCount(blob)*733/1000 + 1 // log(58) / log(256)
	output := make([]byte, capacity)
	outputSize := capacity - 1
	skipZeros, leadZeros := false, 0

	for _, b := range blob {
		// collect prefix zeros
		if !skipZeros {
			if b == alphabetIdx0 {
				leadZeros++
				continue // nolint: nlreturn
			} else {
				skipZeros = true
			}
		}

		carry := int(decodeTable[b])
		if carry == i255 {
			return nil, ErrUnknownFormat()
		}

		idx := capacity - 1
		for ; idx > outputSize || carry != 0; idx-- {
			carry += alphabetSize * int(output[idx])
			output[idx] = byte(rune(carry) % 256) // nolint: gomnd
			carry /= 256
		}

		outputSize = idx
	}

	res := make([]byte, leadZeros+(capacity-1-outputSize))
	copy(res[leadZeros:], output[outputSize+1:])

	return res, nil
}

// DecodeString decodes base58 encoded string to bytes.
func DecodeString(s string) ([]byte, error) {
	return Decode([]byte(s))
}

// Encode encodes given bytes to base58 encoded bytes.
func Encode(blob []byte) []byte {
	leadZeros, size := 0, len(blob)
	for leadZeros < size && blob[leadZeros] == 0 {
		leadZeros++
	}

	// nolint: gomnd
	capacity := (size-leadZeros)*138/100 + 1 // log256 / log58
	outputSize := capacity - 1

	output := make([]byte, capacity)
	for _, b := range blob[leadZeros:] {
		idx := capacity - 1
		for carry := int(b); idx > outputSize || carry != 0; idx-- {
			carry += int(output[idx]) << 8 // nolint: gomnd
			output[idx] = byte(carry % alphabetSize)
			carry /= alphabetSize
		}
		outputSize = idx
	}

	encodeTable := []byte(Alphabet)
	res := make([]byte, leadZeros+(capacity-1-outputSize))
	for i := 0; i < leadZeros; i++ {
		res[i] = alphabetIdx0
	}

	for i, n := range output[outputSize+1:] {
		res[leadZeros+i] = encodeTable[n]
	}

	return res
}

// EncodeToString encodes given bytes to base58 encoded string.
func EncodeToString(b []byte) string {
	return string(Encode(b))
}
