// Copyright © 2020-2021 The EVEN Solutions Developers Team

package errors

const (
	ErrNilPointerValueMsg = "nil pointer value"
)

var (
	errNilPointerValue = New(ErrNilPointerValueMsg)
)

func ErrNilPointerValue() error {
	return errNilPointerValue
}
