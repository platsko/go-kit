// Copyright © 2020-2021 The EVEN Solutions Developers Team

package drive

import (
	"os"
)

const (
	// DefaultFilePerm contents the default Unix permission bits for file.
	DefaultFilePerm = 0o644

	// DefaultFileMode controls the default permissions on any file.
	DefaultFileMode = os.FileMode(DefaultFilePerm)
)

// MakeFile creates or truncates the file.
func MakeFile(name string) (*os.File, error) {
	return os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, DefaultFileMode)
}
