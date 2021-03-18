// Copyright © 2020-2021 The EVEN Solutions Developers Team

package crypto

import (
	cc "github.com/libp2p/go-libp2p-core/crypto"
)

const (
	// UNKNOWN is unknown algo type.
	UNKNOWN Algo = -1

	// RSA is a copy for the supported RSA key type
	// from libp2p crypto package to avoid import it in your project.
	RSA Algo = cc.RSA

	// Ed25519 is a copy for the supported Ed25519 key type
	// from libp2p crypto package to avoid import it in your project.
	Ed25519 Algo = cc.Ed25519

	// Secp256k1 is a copy for the supported Secp256k1 key type
	// from libp2p crypto package to avoid import it in your project.
	Secp256k1 Algo = cc.Secp256k1

	// ECDSA is a copy for the supported ECDSA key type
	// from libp2p crypto package to avoid import it in your project.
	ECDSA Algo = cc.ECDSA
)

type (
	// Algo supported algo enum type.
	Algo int

	// Algos describes list of supported algos
	// key is algo name string and val is algo enum type.
	Algos map[string]Algo
)

// GetAlgos returns a list of all supported algos.
func GetAlgos() Algos {
	return Algos{
		"RSA":       RSA,
		"Ed25519":   Ed25519,
		"Secp256k1": Secp256k1,
		"ECDSA":     ECDSA,
	}
}

// Type returns int representation of the algo type.
func (c Algo) Type() int {
	return int(c)
}

// Copy returns a copy of the algos list.
func (c Algos) Copy() Algos {
	list := make(Algos, c.Len())
	for key, val := range c {
		list[key] = val
	}

	return list
}

// Len returns len of the algos list.
func (c Algos) Len() int {
	return len(c)
}
