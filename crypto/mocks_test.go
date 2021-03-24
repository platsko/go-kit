// Copyright © 2020-2021 The EVEN Solutions Developers Team

package crypto_test

import (
	cc "github.com/libp2p/go-libp2p-core/crypto"

	"github.com/evenlab/go-kit/bytes"
	. "github.com/evenlab/go-kit/crypto"
)

func mockCryptoKeyPair(algo Algo) (cc.PrivKey, cc.PubKey) {
	bits := -1
	if algo == RSA {
		bits = 2048
	}

	prKey, pbKey, err := cc.GenerateKeyPair(int(algo), bits)
	if err != nil {
		panic(err)
	}

	return prKey, pbKey
}

func mockGenerateKeyPair(algo Algo) (PrivateKey, PublicKey) {
	prKey, pbKey, err := GenerateKeyPair(algo)
	if err != nil {
		panic(err)
	}

	return prKey, pbKey
}

func mockSignable(algo Algo, size int) (*SignableStub, PrivateKey) {
	prKi, pbKi := mockCryptoKeyPair(algo)
	signable := SignableStub{Blob: bytes.RandBytes(size)}

	h256, err := signable.Hash()
	if err != nil {
		panic(err)
	}

	sign, err := prKi.Sign(h256[:])
	if err != nil {
		panic(err)
	}

	signable.Sign = NewSignature(sign)
	signable.PbKey = NewPublicKey(pbKi)

	return &signable, NewPrivateKey(prKi)
}

func mockSignature(algo Algo) (Signature, PrivateKey) {
	signable, prKey := mockSignable(algo, 1024)

	return signable.Sign, prKey
}
