// Copyright © 2020-2021 The EVEN Solutions Developers Team

package errors

var (
	delimiter = ": " // nolint: gochecknoglobals
)

func GetDelimiter() string {
	return delimiter
}

func SetDelimiter(s string) {
	delimiter = s
}
