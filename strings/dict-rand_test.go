// Copyright © 2020-2021 The EVEN Solutions Developers Team

package strings_test

import (
	"testing"

	. "github.com/platsko/go-kit/strings"
)

var (
	defaultDictRand = GetDictRand()
)

func Benchmark_GetDictRand(tb *testing.B) {
	for i := 0; i < tb.N; i++ {
		_ = GetDictRand()
	}
}

func Benchmark_SetDictRand(tb *testing.B) {
	dict := GetDictRand()
	tb.ResetTimer()
	for i := 0; i < tb.N; i++ {
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
			if got := GetDictRand(); got != test.dict {
				t.Errorf("GetDictRand() got: %v | want: %v", got, test.dict)
			}
			SetDictRand(dict) // restore previous dictionary
		})
	}
}
