// Copyright © 2020-2021 The EVEN Solutions Developers Team

package crypto

import (
	cc "github.com/libp2p/go-libp2p-core/crypto"
)

// GenerateKeyPair returns generated private and public keys.
func GenerateKeyPair(algo Algo) (PrivateKey, PublicKey, error) {
	bits := -1
	if algo == RSA {
		bits = 2048
	}

	prKey, pbKey, err := cc.GenerateKeyPair(int(algo), bits)
	if err != nil {
		return nil, nil, err
	}

	return NewPrivateKey(prKey), NewPublicKey(pbKey), nil
}
