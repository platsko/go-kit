// Copyright © 2020-2021 The EVEN Solutions Developers Team

package equal_test

import (
	"log"
	"reflect"
	"testing"

	"github.com/platsko/go-kit/bytes"
	. "github.com/platsko/go-kit/equal"
)

func Benchmark_BasicEqual(tb *testing.B) {
	const size = 1024
	blob := bytes.RandBytes(size)
	equaler := NewEqualer(blob)
	tb.ResetTimer()
	for i := 0; i < tb.N; i++ {
		_ = BasicEqual(equaler, equaler)
	}
}

func Benchmark_NewEqualer(tb *testing.B) {
	const size = 1024
	blob := bytes.RandBytes(size)
	tb.ResetTimer()
	for i := 0; i < tb.N; i++ {
		_ = NewEqualer(blob)
	}
}

func Benchmark_equaler_Equals(tb *testing.B) {
	const size = 1024
	blob := bytes.RandBytes(size)
	equ1 := NewEqualer(blob)
	equ2 := NewEqualer(blob)
	tb.ResetTimer()
	for i := 0; i < tb.N; i++ {
		_ = equ1.Equals(equ2)
	}
}

func Benchmark_equaler_Raw(tb *testing.B) {
	const size = 1024
	blob := bytes.RandBytes(size)
	equaler := NewEqualer(blob)
	tb.ResetTimer()
	for i := 0; i < tb.N; i++ {
		if _, err := equaler.Raw(); err != nil {
			log.Fatal(err)
		}
	}
}

func Test_BasicEqual(t *testing.T) {
	t.Parallel()

	const size = 1024
	blob := bytes.RandBytes(size)

	tests := [4]struct {
		name string
		equ1 Equaler
		equ2 Equaler
		want bool
	}{
		{
			name: "TRUE",
			equ1: NewEqualer(blob),
			equ2: NewEqualer(blob),
			want: true,
		},
		{
			name: "FALSE",
			equ1: NewEqualer(blob),
			equ2: NewEqualer(bytes.RandBytes(size)),
		},
		{
			name: "nil_first_Equaler_FALSE",
			equ1: NewEqualer(nil),
			equ2: NewEqualer(blob),
		},
		{
			name: "nil_second_Equaler_FALSE",
			equ1: NewEqualer(blob),
			equ2: NewEqualer(nil),
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := BasicEqual(test.equ1, test.equ2); got != test.want {
				t.Errorf("BasicEqual() got: %v | want: %v", got, test.want)
			}
		})
	}
}

func Test_NewEqualer(t *testing.T) {
	t.Parallel()

	const size = 1024
	blob := bytes.RandBytes(size)

	tests := [1]struct {
		name string
		blob []byte
		want []byte
	}{
		{
			name: "OK",
			blob: blob,
			want: blob,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			equaler := NewEqualer(test.blob)
			if got, _ := equaler.Raw(); !reflect.DeepEqual(got, test.want) {
				t.Errorf("NewEqualer() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}

func Test_equaler_Equals(t *testing.T) {
	t.Parallel()

	const size = 1024
	blob := bytes.RandBytes(size)

	tests := [2]struct {
		name string
		equ1 Equaler
		equ2 Equaler
		want bool
	}{
		{
			name: "TRUE",
			equ1: NewEqualer(blob),
			equ2: NewEqualer(blob),
			want: true,
		},
		{
			name: "FALSE",
			equ1: NewEqualer(blob),
			equ2: NewEqualer(bytes.RandBytes(size)),
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := test.equ1.Equals(test.equ2); got != test.want {
				t.Errorf("Equals() got: %v | want: %v", got, test.want)
			}
		})
	}
}

func Test_equaler_Raw(t *testing.T) {
	t.Parallel()

	const size = 1024
	blob := bytes.RandBytes(size)

	tests := [2]struct {
		name    string
		equaler Equaler
		want    []byte
		wantErr bool
	}{
		{
			name:    "OK",
			equaler: NewEqualer(blob),
			want:    blob,
		},
		{
			name:    "ERR",
			equaler: NewEqualer(nil),
			wantErr: true,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, err := test.equaler.Raw()
			if (err != nil) != test.wantErr {
				t.Errorf("Raw() error: %v | want: %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Raw() got: %v | want: %v", got, test.want)
			}
		})
	}
}
