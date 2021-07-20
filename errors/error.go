// Copyright © 2020-2021 The EVEN Solutions Developers Team

package errors

import (
	"errors"
)

// As wraps function errors.As
// to avoid import errors package from standard library.
func As(err error, target *error) bool {
	return errors.As(err, target)
}

// Is wraps function errors.Is
// to avoid import errors package from standard library.
func Is(err, target error) bool {
	return errors.Is(err, target)
}

// New returns an error that formats as the given string.
func New(s string) error {
	return &wrapper{txt: s}
}

// Unwrap wraps function errors.Unwrap
// to avoid import errors package from standard library.
func Unwrap(err error) error {
	return errors.Unwrap(err)
}

// WrapErr returns an error that formats the given string and
// wrapped the given error with Delimiter in the middle.
func WrapErr(s string, w error) error {
	return &wrapper{
		err: w,
		txt: s + GetDelimiter() + w.Error(),
	}
}

// WrapStr returns an error that formats the given string and
// wrapped the given string converted to error with Delimiter in the middle.
func WrapStr(s, w string) error {
	return &wrapper{
		err: New(w),
		txt: s + GetDelimiter() + w,
	}
}
