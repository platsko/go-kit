// Copyright © 2020-2021 The EVEN Solutions Developers Team

package strings

import (
	"strings"

	"github.com/evenlab/go-kit/bytes"
)

// RandString returns random generated string with given size.
func RandString(size int) string {
	s := GetDictRand()
	blob, builder, dictSize := bytes.RandBytes(size), strings.Builder{}, len(s)
	for i := 0; i < size; i++ {
		blob[i] = s[int(blob[i])%dictSize]
	}

	_, _ = builder.Write(blob)

	return builder.String()
}
