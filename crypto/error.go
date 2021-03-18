// Copyright © 2020-2021 The EVEN Solutions Developers Team

package crypto

import (
	"github.com/platsko/go-kit/errors"
)

const (
	ErrPublicKeyCannotBeNilMsg = "public key cannot be nil"
	ErrSignableCannotBeNilMsg  = "signable cannot be nil"
	ErrSignatureCannotBeNilMsg = "signature cannot be nil"
)

var (
	errPublicKeyCannotBeNil = errors.New(ErrPublicKeyCannotBeNilMsg)
	errSignableCannotBeNil  = errors.New(ErrSignableCannotBeNilMsg)
	errSignatureCannotBeNil = errors.New(ErrSignatureCannotBeNilMsg)
)

func ErrPublicKeyCannotBeNil() error {
	return errPublicKeyCannotBeNil
}

func ErrSignableCannotBeNil() error {
	return errSignableCannotBeNil
}

func ErrSignatureCannotBeNil() error {
	return errSignatureCannotBeNil
}
