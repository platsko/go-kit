// Copyright © 2020-2021 The EVEN Solutions Developers Team

package crypto

type (
	// Signable represents interface for signable types.
	// It provides ability to calculate a hash func and set the signature as property.
	Signable interface {
		// Embedded Hasher interface.
		Hasher

		// GetSignature returns the signature.
		GetSignature() Signature

		// SetSignature sets the sign to the particular property.
		SetSignature(Signature)

		// SetPublicKey sets the public key to the particular property.
		SetPublicKey(PublicKey)
	}

	// Signer represents interface for signing signable objects.
	Signer interface {
		// Sign signs signable object.
		Sign(Signable) (Signature, error)
	}
)
