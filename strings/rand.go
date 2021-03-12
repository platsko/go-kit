// Copyright © 2020-2021 The EVEN Solutions Developers Team

package strings

import (
	"strings"

	"github.com/platsko/go-kit/bytes"
)

// RandString returns random generated string with given size.
func RandString(size int) string {
	blob, builder := bytes.RandBytes(size), strings.Builder{}
	for i := 0; i < size; i++ {
		blob[i] = dictRandChar[int(blob[i])%dictRandSize]
	}

	_, _ = builder.Write(blob)

	return builder.String()
}
