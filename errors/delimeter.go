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
	rwDelimMutex.RLock()
	delim := delimiter
	rwDelimMutex.RUnlock()

	return delim
}

// SetDelimiter sets new delimiter chars to use for wrap errors.
func SetDelimiter(s string) {
	rwDelimMutex.Lock()
	delimiter = s
	rwDelimMutex.Unlock()
}
