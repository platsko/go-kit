// Copyright © 2020-2021 The EVEN Solutions Developers Team

package errors

var (
	delimiter = ": " // nolint: gochecknoglobals
)

// GetDelimiter returns current use delimiter chars for wrap errors.
func GetDelimiter() string {
	return delimiter
}

// SetDelimiter sets new delimiter chars to use for wrap errors.
func SetDelimiter(s string) {
	delimiter = s
}
