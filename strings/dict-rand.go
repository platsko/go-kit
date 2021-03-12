// Copyright © 2020-2021 The EVEN Solutions Developers Team

package strings

var (
	dictRandChar = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" // nolint: gochecknoglobals
	dictRandSize = len(dictRandChar)                                      // nolint: gochecknoglobals
)

// GetDictRand returns current use dictionary chars for rand.
func GetDictRand() string {
	return dictRandChar
}

// SetDictRand sets new dictionary chars to use for rand.
func SetDictRand(s string) {
	dictRandChar = s
	dictRandSize = len(dictRandChar)
}
