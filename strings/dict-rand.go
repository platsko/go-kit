// Copyright © 2020-2021 The EVEN Solutions Developers Team

package strings

import (
	"sync"
)

var (
	dictRandChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" // nolint: gochecknoglobals

	rwDictMutex = sync.RWMutex{} // nolint: gochecknoglobals
)

// GetDictRand returns current use dictionary chars for rand.
func GetDictRand() string {
	rwDictMutex.RLock()
	dict := dictRandChars
	rwDictMutex.RUnlock()

	return dict
}

// SetDictRand sets new dictionary chars to use for rand.
func SetDictRand(s string) {
	rwDictMutex.Lock()
	dictRandChars = s
	rwDictMutex.Unlock()
}
