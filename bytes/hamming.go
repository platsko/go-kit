// Copyright © 2020-2021 The EVEN Solutions Developers Team

package bytes

import (
	"math/bits"
)

// Hamming returns the distance between the given vectors.
func Hamming(vx []byte, vy []byte) (int, error) {
	if len(vx) == 0 || len(vy) == 0 {
		return 0, ErrVectorZeroSize()
	}

	if len(vx) != len(vy) {
		return 0, ErrVectorsNotSameSize()
	}

	dist := 0
	for i, l := 0, len(vx); i < l; i++ {
		dist += bits.OnesCount8(vx[i] ^ vy[i])
	}

	return dist, nil
}
