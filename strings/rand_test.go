// Copyright © 2020-2021 The EVEN Solutions Developers Team

package strings_test

import (
	"strings"
	"testing"

	. "github.com/platsko/go-kit/strings"
)

func Benchmark_RandString(tb *testing.B) {
	const size = 1024
	tb.ResetTimer()
	for i := 0; i < tb.N; i++ {
		_ = RandString(size)
	}
}

func Test_RandString(t *testing.T) {
	t.Parallel()

	tests := [1]struct {
		name string
		size int
		dict string
	}{
		{
			name: "OK",
			size: 1024,
			dict: defaultDictRand,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			dict := GetDictRand()        // save current dictionary
			SetDictRand(test.dict)       // set test dictionary
			got := RandString(test.size) // rand by test dictionary
			SetDictRand(dict)            // restore previous dictionary

			if len(got) != test.size { // check expected size
				t.Errorf("RandString() size: %v | want: %v", got, test.size)
			}
			if out := strings.Trim(got, test.dict); out != "" {
				t.Errorf("RandString() has chars: %v | allowed: %v", out, test.dict)
			}
		})
	}
}
