// Copyright © 2020-2021 The EVEN Solutions Developers Team

package errors_test

import (
	"testing"

	. "github.com/platsko/go-kit/errors"
)

var (
	defaultDelimiter = GetDelimiter()
)

func Benchmark_GetDelimiter(tb *testing.B) {
	for i := 0; i < tb.N; i++ {
		_ = GetDelimiter()
	}
}

func Benchmark_SetDelimiter(tb *testing.B) {
	delim := GetDelimiter()
	tb.ResetTimer()
	for i := 0; i < tb.N; i++ {
		SetDelimiter(delim)
	}
}

func Test_GetDelimiter(t *testing.T) {
	t.Parallel()

	tests := [1]struct {
		name string
		want string
	}{
		{
			name: "Default",
			want: defaultDelimiter,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := GetDelimiter(); got != test.want {
				t.Errorf("GetDelimiter() got: %v | want: %v", got, test.want)
			}
		})
	}
}

func Test_SetDelimiter(t *testing.T) {
	t.Parallel()

	tests := [1]struct {
		name  string
		delim string
	}{
		{
			name:  "OK",
			delim: " | ",
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			delim := GetDelimiter()  // save current delimiter
			SetDelimiter(test.delim) // set test delimiter
			if got := GetDelimiter(); got != test.delim {
				t.Errorf("GetDictRand() got: %v | want: %v", got, test.delim)
			}
			SetDelimiter(delim) // restore previous delimiter
		})
	}
}
