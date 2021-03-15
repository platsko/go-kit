// Copyright © 2020-2021 The EVEN Solutions Developers Team

package crypto_test

import (
	"crypto/sha256"
	"encoding/hex"
	"reflect"
	"testing"

	"github.com/platsko/go-kit/base58"
	"github.com/platsko/go-kit/bytes"
	. "github.com/platsko/go-kit/crypto"
	"github.com/platsko/go-kit/strings"
)

var (
	h224 = NewHash224(
		bytes.RandBytes(128),
		bytes.RandBytes(256),
		bytes.RandBytes(512),
	)
)

func Benchmark_NewHash224(tb *testing.B) {
	b1 := bytes.RandBytes(128)
	b2 := bytes.RandBytes(256)
	b3 := bytes.RandBytes(512)
	tb.ResetTimer()
	for i := 0; i < tb.N; i++ {
		_ = NewHash224(b1, b2, b3)
	}
}

func Benchmark_StrToHash224(tb *testing.B) {
	s1 := strings.RandString(16)
	s2 := strings.RandString(32)
	s3 := strings.RandString(64)
	tb.ResetTimer()
	for i := 0; i < tb.N; i++ {
		_ = StrToHash224(s1, s2, s3)
	}
}

func Benchmark_Hash224_Base58(tb *testing.B) {
	for i := 0; i < tb.N; i++ {
		_ = h224.Base58()
	}
}

func Benchmark_Hash224_Encode(tb *testing.B) {
	for i := 0; i < tb.N; i++ {
		_ = h224.Encode()
	}
}

func Benchmark_Hash224_Hamming(tb *testing.B) {
	v224 := [Hash224Size]byte{}
	copy(v224[:], bytes.RandBytes(160))
	tb.ResetTimer()
	for i := 0; i < tb.N; i++ {
		_ = h224.Hamming(v224)
	}
}

func Benchmark_Hash224_String(tb *testing.B) {
	for i := 0; i < tb.N; i++ {
		_ = h224.String()
	}
}

func Test_NewHash224(t *testing.T) {
	t.Parallel()

	blob := make([]byte, 0)
	args := [][]byte{
		bytes.RandBytes(128),
		bytes.RandBytes(256),
		bytes.RandBytes(512),
	}

	for _, b := range args {
		blob = append(blob, b...)
	}
	h256 := sha256.Sum256(blob)
	h224 := sha256.Sum224(h256[:])

	tests := [1]struct {
		name string
		args [][]byte
		want Hash224
	}{
		{
			name: "OK",
			args: args,
			want: h224,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := NewHash224(test.args...); !reflect.DeepEqual(got, test.want) {
				t.Errorf("NewHash224() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}

func Test_StrToHash224(t *testing.T) {
	t.Parallel()

	args := []string{
		strings.RandString(16),
		strings.RandString(32),
		strings.RandString(64),
	}

	str := ""
	for _, s := range args {
		str += s
	}
	h256 := sha256.Sum256([]byte(str))
	h224 := sha256.Sum224(h256[:])

	tests := [1]struct {
		name string
		args []string
		want Hash224
	}{
		{
			name: "OK",
			args: args,
			want: h224,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := StrToHash224(test.args...); !reflect.DeepEqual(got, test.want) {
				t.Errorf("StrToHash224() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}

func Test_Hash224_Base58(t *testing.T) {
	t.Parallel()

	tests := [1]struct {
		name string
		h224 Hash224
		want string
	}{
		{
			name: "OK",
			h224: h224,
			want: base58.EncodeToString(h224[:]),
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := test.h224.Base58(); got != test.want {
				t.Errorf("Base58() got: %v | want: %v", got, test.want)
			}
		})
	}
}

func Test_Hash224_Encode(t *testing.T) {
	t.Parallel()

	tests := [1]struct {
		name string
		h224 Hash224
		want string
	}{
		{
			name: "OK",
			h224: h224,
			want: hex.EncodeToString(h224[:]),
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := test.h224.Encode(); got != test.want {
				t.Errorf("Encode() got: %v | want: %v", got, test.want)
			}
		})
	}
}

func Test_Hash224_Hamming(t *testing.T) {
	t.Parallel()

	tests := [1]struct {
		name string
		h224 Hash224
		v160 [Hash224Size]byte
		want int
	}{
		{
			name: "OK",
			h224: Hash224{1},
			v160: [Hash224Size]byte{15},
			want: 3,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := test.h224.Hamming(test.v160); got != test.want {
				t.Errorf("Hamming() got: %v | want: %v", got, test.want)
			}
		})
	}
}

func Test_Hash224_String(t *testing.T) {
	t.Parallel()

	tests := [1]struct {
		name string
		h224 Hash224
		want string
	}{
		{
			name: "OK",
			h224: h224,
			want: string(h224[:]),
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := test.h224.String(); got != test.want {
				t.Errorf("String() got: %v | want: %v", got, test.want)
			}
		})
	}
}
