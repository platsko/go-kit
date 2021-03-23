// Copyright © 2020-2021 The EVEN Solutions Developers Team

package errors

const (
	ErrNilPointerValueMsg = "nil pointer value"
	ErrZeroSizeValueMsg   = "zero size value"
)

var (
	errNilPointerValue = New(ErrNilPointerValueMsg)
	errZeroSizeValue   = New(ErrZeroSizeValueMsg)
)

func ErrNilPointerValue() error {
	return errNilPointerValue
}

func ErrZeroSizeValue() error {
	return errZeroSizeValue
}
