// Copyright © 2020-2021 The EVEN Solutions Developers Team

package crypto

import (
	"github.com/platsko/go-kit/errors"
)

const (
	ErrVectorsNotSameSizeMsg = "vectors must be the same size"
	ErrVectorZeroSizeMsg     = "vectors cannot be zero size"
)

var (
	errVectorsNotSameSize = errors.New(ErrVectorsNotSameSizeMsg)
	errVectorZeroSize     = errors.New(ErrVectorZeroSizeMsg)
)

func ErrVectorsNotSameSize() error {
	return errVectorsNotSameSize
}

func ErrVectorZeroSize() error {
	return errVectorZeroSize
}
