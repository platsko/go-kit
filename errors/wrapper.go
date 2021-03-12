// Copyright © 2020-2021 The EVEN Solutions Developers Team

package errors

type (
	// Wrapper describes error wrapper interface.
	Wrapper interface {
		// Error returns error as string.
		Error() string

		// Unwrap returns wrapped error.
		Unwrap() error
	}

	// wrapper implements Wrapper interface.
	wrapper struct {
		err error
		txt string
	}
)

// Error implements error Wrapper interface.
func (e *wrapper) Error() string {
	return e.txt
}

// Unwrap implements error Wrapper interface.
func (e *wrapper) Unwrap() error {
	return e.err
}
