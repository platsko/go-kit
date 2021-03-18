// Copyright © 2020-2021 The EVEN Solutions Developers Team

package equal

import (
	"github.com/platsko/go-kit/errors"
)

const (
	ErrZeroSizeValueMsg = "zero size value"
)

var (
	errZeroSizeValue = errors.New(ErrZeroSizeValueMsg)
)

func ErrZeroSizeValue() error {
	return errZeroSizeValue
}
