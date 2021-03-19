// Copyright © 2020-2021 The EVEN Solutions Developers Team

package bytes_test

import (
	"testing"

	. "github.com/platsko/go-kit/bytes"
)

func Benchmark_RandBytes(tb *testing.B) {
	const size = 1024
	tb.ResetTimer()
	for i := 0; i < tb.N; i++ {
		_ = RandBytes(size)
	}
}

func Test_RandBytes(t *testing.T) {
	t.Parallel()

	tests := [1]struct {
		name string
		size int
	}{
		{
			name: "OK",
			size: 1024,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := RandBytes(test.size); len(got) != test.size {
				t.Errorf("RandBytes() size: %v | want: %v", got, test.size)
			}
		})
	}
}
