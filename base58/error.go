// Copyright © 2020-2021 The EVEN Solutions Developers Team

package base58

import (
	"github.com/platsko/go-kit/errors"
)

const (
	ErrChecksumMismatchMsg = "checksum mismatch"
	ErrInvalidFormatMsg    = "invalid format"
	ErrUnknownFormatMsg    = "unknown format"
)

var (
	errChecksumMismatch = errors.New(ErrChecksumMismatchMsg)
	errInvalidFormat    = errors.New(ErrInvalidFormatMsg)
	errUnknownFormat    = errors.New(ErrUnknownFormatMsg)
)

func ErrChecksumMismatch() error {
	return errChecksumMismatch
}

func ErrInvalidFormat() error {
	return errInvalidFormat
}

func ErrUnknownFormat() error {
	return errUnknownFormat
}
