// Copyright © 2020-2021 The EVEN Solutions Developers Team

package bytes

// Equal determines that vectors are equals.
func Equal(vx, vy []byte) bool {
	if (vx == nil) != (vy == nil) || len(vx) != len(vy) {
		return false
	}

	for i, b := range vx {
		if b != vy[i] {
			return false
		}
	}

	return true
}
