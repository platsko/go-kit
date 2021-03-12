// Copyright © 2020-2021 The EVEN Solutions Developers Team

package errors_test

import (
	"testing"

	. "github.com/platsko/go-kit/errors"
)

func Benchmark_GetDelimiter(tb *testing.B) {
	for i := 0; i < tb.N; i++ {
		_ = GetDelimiter()
	}
}

func Benchmark_SetDelimiter(tb *testing.B) {
	for i := 0; i < tb.N; i++ {
		SetDelimiter(errDelimiter)
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
			want: ": ",
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
		want  string
	}{
		{
			name:  "OK",
			delim: " | ",
			want:  " | ",
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			SetDelimiter(test.delim)
			if got := GetDelimiter(); got != test.want {
				t.Errorf("SetDelimiter() got: %v | want: %v", got, test.want)
			}
			SetDelimiter(": ") // restore default
		})
	}
}
