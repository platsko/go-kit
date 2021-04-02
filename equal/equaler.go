// Copyright © 2020-2021 The EVEN Solutions Developers Team

package equal

import (
	"crypto/subtle"

	"github.com/platsko/go-kit/errors"
)

type (
	// Equaler represents equaler interface.
	Equaler interface {
		// Raw returns a copy of the raw bytes of the implementation.
		Raw() ([]byte, error)
	}

	// equaler implements Equaler interface.
	equaler struct {
		blob []byte
	}
)

// BasicEqual reports whether e1 and e2 are the same length
// and contain the same bytes, nil argument is equivalent to an empty slice.
func BasicEqual(e1, e2 Equaler) bool {
	if e1 == nil { // wrap nil pointer into Equaler interface to avoid panic
		e1 = NewEqualer(nil)
	}

	if e2 == nil { // wrap nil pointer into Equaler interface to avoid panic
		e2 = NewEqualer(nil)
	}

	b1, err := e1.Raw()
	if err != nil && !errors.Is(err, errors.ErrNilPointerValue()) { // ignore nil pointer error
		return false
	}

	b2, err := e2.Raw()
	if err != nil && !errors.Is(err, errors.ErrNilPointerValue()) { // ignore nil pointer error
		return false
	}

	return subtle.ConstantTimeCompare(b1, b2) == 1
}

// NewEqualer constructs Equaler interface with provided raw bytes data.
func NewEqualer(blob []byte) Equaler {
	return &equaler{blob: blob}
}

// Raw implements Equaler.Raw method of interface.
func (e *equaler) Raw() ([]byte, error) {
	if e.blob == nil {
		return nil, errors.ErrNilPointerValue()
	}

	if len(e.blob) == 0 { // enforce this error for unit testing
		return nil, errors.ErrZeroSizeValue()
	}

	blob := make([]byte, len(e.blob))
	copy(blob, e.blob)

	return blob, nil
}
