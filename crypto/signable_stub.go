// Copyright © 2020-2021 The EVEN Solutions Developers Team

package crypto

import (
	"github.com/evenlab/go-kit/errors"
)

type (
	// SignableStub implements Signable interface.
	SignableStub struct {
		Blob  []byte
		Sign  Signature
		PbKey PublicKey
	}
)

// NewSignable constructs Signable interface over given bytes.
func NewSignable(blob []byte) Signable {
	return &SignableStub{Blob: blob}
}

// GetSignature implements Signable.GetSignature method of interface.
func (c *SignableStub) GetSignature() Signature {
	return c.Sign
}

// Hash implements Hasher.Hash method of interface.
func (c *SignableStub) Hash() (Hash256, error) {
	if c.Blob == nil {
		return Hash256{}, errors.ErrNilPointerValue()
	}

	return NewHash256(c.Blob), nil
}

// SetPublicKey implements Signable.SetPublicKey method of interface.
func (c *SignableStub) SetPublicKey(pbKey PublicKey) {
	c.PbKey = pbKey
}

// SetSignature implements Signable.SetSignature method of interface.
func (c *SignableStub) SetSignature(sign Signature) {
	c.Sign = sign
}
