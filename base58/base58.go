// Copyright © 2020-2021 The EVEN Solutions Developers Team

package base58

import (
	"unicode/utf8"
)

const (
	// alphabet is the modified base58 alphabet.
	alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

	alphabetIdx0 = '1'
	alphabetSize = len(alphabet)

	// i255 is a magic number represented symbol that does not exist in alphabet.
	i255 = 255
)

// Decode decodes base58 encoded bytes.
func Decode(b []byte) ([]byte, error) {
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
	capacity := utf8.RuneCount(b)*733/1000 + 1 // log(58) / log(256)
	output := make([]byte, capacity)
	outputReverseEnd := capacity - 1

	skipZeros, prefixZeroes := false, 0
	for _, target := range b {
		// collect prefix zeros
		if !skipZeros {
			if target == alphabetIdx0 {
				prefixZeroes++
				continue // nolint: nlreturn
			} else {
				skipZeros = true
			}
		}

		carry := int(decodeTable[target])
		if carry == i255 {
			return nil, ErrUnknownFormat()
		}

		outputIdx := capacity - 1
		for ; outputIdx > outputReverseEnd || carry != 0; outputIdx-- {
			carry += alphabetSize * int(output[outputIdx])
			output[outputIdx] = byte(rune(carry) % 256) // nolint: gomnd
			carry /= 256
		}

		outputReverseEnd = outputIdx
	}

	retBytes := make([]byte, prefixZeroes+(capacity-1-outputReverseEnd))
	copy(retBytes[prefixZeroes:], output[outputReverseEnd+1:])

	return retBytes, nil
}

// DecodeString decodes base58 encoded string to bytes.
func DecodeString(s string) ([]byte, error) {
	return Decode([]byte(s))
}

// Encode encodes given bytes to base58 encoded bytes.
func Encode(b []byte) []byte {
	prefixZeroes, inputLength := 0, len(b)
	for prefixZeroes < inputLength && b[prefixZeroes] == 0 {
		prefixZeroes++
	}

	// nolint: gomnd
	capacity := (inputLength-prefixZeroes)*138/100 + 1 // log256 / log58
	output := make([]byte, capacity)
	outputReverseEnd := capacity - 1

	for _, inputByte := range b[prefixZeroes:] {
		outputIdx := capacity - 1
		for carry := int(inputByte); outputIdx > outputReverseEnd || carry != 0; outputIdx-- {
			carry += int(output[outputIdx]) << 8 // nolint: gomnd
			output[outputIdx] = byte(carry % alphabetSize)
			carry /= alphabetSize
		}

		outputReverseEnd = outputIdx
	}

	blob, encodeTable := make([]byte, prefixZeroes+(capacity-1-outputReverseEnd)), []byte(alphabet)
	for i := 0; i < prefixZeroes; i++ {
		blob[i] = encodeTable[0]
	}

	for i, n := range output[outputReverseEnd+1:] {
		blob[prefixZeroes+i] = encodeTable[n]
	}

	return blob
}

// EncodeToString encodes given bytes to base58 encoded string.
func EncodeToString(b []byte) string {
	return string(Encode(b))
}
