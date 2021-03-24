// Copyright © 2020-2021 The EVEN Solutions Developers Team

package drive

import (
	"os"

	"github.com/evenlab/go-kit/errors"
)

const (
	// DefaultDirPerm contents the default Unix permission bits for dir.
	DefaultDirPerm = 0o755

	// DefaultDirMode controls the default permissions on any dir.
	DefaultDirMode = os.FileMode(DefaultDirPerm)
)

// MakeDirs makes dirs that the full paths you specified.
func MakeDirs(dirs ...string) error {
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, DefaultDirMode); err != nil {
			return errors.WrapErr("make dirs", err)
		}
	}

	return nil
}
