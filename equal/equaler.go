// Copyright © 2020-2021 The EVEN Solutions Developers Team

package equal

import (
	"crypto/subtle"

	"github.com/platsko/go-kit/bytes"
	"github.com/platsko/go-kit/errors"
)

type (
	// Equaler represents equaler interface.
	Equaler interface {
		// Equals checks whether two equalers are the same.
		Equals(Equaler) bool

		// Raw returns the copy of raw bytes data.
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
	b1, err := e1.Raw()
	if err != nil {
		return false
	}

	b2, err := e2.Raw()
	if err != nil {
		return false
	}

	return subtle.ConstantTimeCompare(b1, b2) == 1 || bytes.Equal(b1, b2)
}

// NewEqualer constructs Equaler interface with provided raw bytes data.
func NewEqualer(blob []byte) Equaler {
	return &equaler{blob: blob}
}

// Equals implements Equaler.Equals method of interface.
func (e *equaler) Equals(compare Equaler) bool {
	return BasicEqual(e, compare)
}

// Raw implements Equaler.Raw method of interface.
func (e *equaler) Raw() ([]byte, error) {
	if e.blob == nil {
		return nil, errors.ErrNilPointerValue()
	}

	blob := make([]byte, len(e.blob))
	copy(blob, e.blob)

	return blob, nil
}
