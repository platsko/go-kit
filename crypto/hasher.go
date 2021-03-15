// Copyright © 2020-2021 The EVEN Solutions Developers Team

package crypto

type (
	// Hasher represents interface for types that can compute hash of themselves.
	Hasher interface {
		// Hash returns calculates hash checksum.
		Hash() (Hash256, error)
	}
)
