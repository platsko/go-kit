// Copyright © 2020-2021 The EVEN Solutions Developers Team

package errors_test

import (
	"testing"

	. "github.com/platsko/go-kit/errors"
)

var (
	defaultDelimiter = GetDelimiter()
)

func Benchmark_GetDelimiter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GetDelimiter()
	}
}

func Benchmark_SetDelimiter(b *testing.B) {
	delim := GetDelimiter()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
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
			got := GetDelimiter()    // get test delimiter
			SetDelimiter(delim)      // restore previous delimiter

			if got != test.delim { // check test delimiter
				t.Errorf("SetDelimiter() got: %v | want: %v", got, test.delim)
			}
		})
	}
}
