// Copyright © 2020-2021 The EVEN Solutions Developers Team

package crypto

import (
	cc "github.com/libp2p/go-libp2p-core/crypto"

	"github.com/platsko/go-kit/errors"
)

type (
	// PrivateKey represents private key interface.
	PrivateKey interface {
		// Embedded Signer interface.
		Signer

		// Algo returns the private key Algo.
		Algo() Algo

		// PublicKey returns the public key paired with this private key.
		PublicKey() PublicKey
	}

	// privateKey implements PrivateKey interface.
	privateKey struct {
		ki cc.PrivKey
	}
)

var (
	// Make sure privateKey implements PrivateKey interface.
	_ PrivateKey = (*privateKey)(nil)
)

// NewPrivateKey returns PrivateKey interface.
func NewPrivateKey(ki cc.PrivKey) PrivateKey {
	return &privateKey{ki: ki}
}

// Algo implements PrivateKey.Algo method of interface.
func (c *privateKey) Algo() Algo {
	if c.ki == nil {
		return UNKNOWN
	}

	algo := c.ki.Type()

	return Algo(algo)
}

// PublicKey implements PrivateKey.PublicKey method of interface.
func (c *privateKey) PublicKey() PublicKey {
	pbKey := publicKey{ki: nil}
	if c.ki != nil {
		pbKey.ki = c.ki.GetPublic()
	}

	return &pbKey
}

// Sign implements Signer.Sign method of interface.
func (c *privateKey) Sign(signable Signable) (Signature, error) {
	if signable == nil {
		return nil, ErrSignableCannotBeNil()
	}

	if c.ki == nil {
		return nil, errors.ErrNilPointerValue()
	}

	pbKey := c.PublicKey()
	signable.SetPublicKey(pbKey)

	h256, err := signable.Hash()
	if err != nil {
		return nil, err
	}

	blob, err := c.ki.Sign(h256[:])
	if err != nil {
		return nil, err
	}

	sign := NewSignature(blob)
	signable.SetSignature(sign)

	return sign, nil
}
