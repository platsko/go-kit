// Copyright © 2020-2021 The EVEN Solutions Developers Team

package errors

import (
	"sync"
)

var (
	delimiter = ": " // nolint: gochecknoglobals

	rwDelimMutex = sync.RWMutex{} // nolint: gochecknoglobals
)

// GetDelimiter returns current use delimiter chars for wrap errors.
func GetDelimiter() string {
	rwDelimMutex.Lock()
	delim := delimiter
	rwDelimMutex.Unlock()

	return delim
}

// SetDelimiter sets new delimiter chars to use for wrap errors.
func SetDelimiter(s string) {
	rwDelimMutex.RLock()
	delimiter = s
	rwDelimMutex.RUnlock()
}
