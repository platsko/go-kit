// Copyright © 2020-2021 The EVEN Solutions Developers Team

package strings_test

import (
	"testing"

	. "github.com/platsko/go-kit/strings"
)

var (
	defaultDictRand = GetDictRand()
)

func Benchmark_GetDictRand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GetDictRand()
	}
}

func Benchmark_SetDictRand(b *testing.B) {
	dict := GetDictRand()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SetDictRand(dict)
	}
}

func Test_GetDictRand(t *testing.T) {
	t.Parallel()

	tests := [1]struct {
		name string
		want string
	}{
		{
			name: "Default",
			want: defaultDictRand,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := GetDictRand(); got != test.want {
				t.Errorf("GetDictRand() got: %v | want: %v", got, test.want)
			}
		})
	}
}

func Test_SetDictRand(t *testing.T) {
	t.Parallel()

	tests := [1]struct {
		name string
		dict string
	}{
		{
			name: "OK",
			dict: "01234567889",
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			dict := GetDictRand()  // save current dictionary
			SetDictRand(test.dict) // set test dictionary
			got := GetDictRand()   // get test dictionary
			SetDictRand(dict)      // restore previous dictionary

			if got != test.dict { // check test dictionary
				t.Errorf("SetDictRand() got: %v | want: %v", got, test.dict)
			}
		})
	}
}
