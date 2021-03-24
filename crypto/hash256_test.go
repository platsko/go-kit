// Copyright © 2020-2021 The EVEN Solutions Developers Team

package crypto_test

import (
	"crypto/sha256"
	"encoding/hex"
	"reflect"
	"testing"

	"github.com/evenlab/go-kit/base58"
	"github.com/evenlab/go-kit/bytes"
	. "github.com/evenlab/go-kit/crypto"
	"github.com/evenlab/go-kit/strings"
)

var (
	h256 = NewHash256(
		bytes.RandBytes(128),
		bytes.RandBytes(256),
		bytes.RandBytes(512),
	)
)

func Benchmark_NewHash256(b *testing.B) {
	b1 := bytes.RandBytes(128)
	b2 := bytes.RandBytes(256)
	b3 := bytes.RandBytes(512)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewHash256(b1, b2, b3)
	}
}

func Benchmark_StrToHash256(b *testing.B) {
	s1 := strings.RandString(16)
	s2 := strings.RandString(32)
	s3 := strings.RandString(64)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = StrToHash256(s1, s2, s3)
	}
}

func Benchmark_Hash256_Base58(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = h256.Base58()
	}
}

func Benchmark_Hash256_Empty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Hash256{}.Empty()
	}
}

func Benchmark_Hash256_Encode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = h256.Encode()
	}
}

func Benchmark_Hash256_Hamming(b *testing.B) {
	v256 := [Hash256Size]byte{}
	copy(v256[:], bytes.RandBytes(256))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = h256.Hamming(v256)
	}
}

func Benchmark_Hash256_String(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = h256.String()
	}
}

func Test_NewHash256(t *testing.T) {
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
	h256, b256 := Hash256{}, sha256.Sum256(blob)
	copy(h256[:], b256[:])

	tests := [1]struct {
		name string
		args [][]byte
		want Hash256
	}{
		{
			name: "Test_NewHash256_OK",
			args: args,
			want: h256,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := NewHash256(test.args...); !reflect.DeepEqual(got, test.want) {
				t.Errorf("crypto.NewHash256() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}

func Test_StrToHash256(t *testing.T) {
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
	blob := []byte(str)
	h256, b256 := Hash256{}, sha256.Sum256(blob)
	copy(h256[:], b256[:])

	tests := [1]struct {
		name string
		args []string
		want Hash256
	}{
		{
			name: "Test_StrToHash256_OK",
			args: args,
			want: h256,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := StrToHash256(test.args...); !reflect.DeepEqual(got, test.want) {
				t.Errorf("StrToHash256() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}

func Test_Hash256_Base58(t *testing.T) {
	t.Parallel()

	tests := [1]struct {
		name string
		h256 Hash256
		want string
	}{
		{
			name: "Test_Hash256_Base58_OK",
			h256: h256,
			want: base58.EncodeToString(h256[:]),
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := test.h256.Base58(); got != test.want {
				t.Errorf("Base58() got: %v | want: %v", got, test.want)
			}
		})
	}
}

func Test_Hash256_Empty(t *testing.T) {
	t.Parallel()

	tests := [2]struct {
		name string
		h256 Hash256
		want bool
	}{
		{
			name: "TRUE",
			h256: Hash256{},
			want: true,
		},
		{
			name: "FALSE",
			h256: h256,
			want: false,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := test.h256.Empty(); got != test.want {
				t.Errorf("Encode() got: %v | want: %v", got, test.want)
			}
		})
	}
}

func Test_Hash256_Encode(t *testing.T) {
	t.Parallel()

	tests := [1]struct {
		name string
		h256 Hash256
		want string
	}{
		{
			name: "Test_Hash256_Encode_OK",
			h256: h256,
			want: hex.EncodeToString(h256[:]),
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := test.h256.Encode(); got != test.want {
				t.Errorf("EncodeToString() got: %v | want: %v", got, test.want)
			}
		})
	}
}

func Test_Hash256_Hamming(t *testing.T) {
	t.Parallel()

	tests := [1]struct {
		name string
		h256 Hash256
		v256 [Hash256Size]byte
		want int
	}{
		{
			name: "Test_Hash256_Hamming_OK",
			h256: Hash256{1},
			v256: [Hash256Size]byte{15},
			want: 3,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := test.h256.Hamming(test.v256); got != test.want {
				t.Errorf("Hamming() got: %v | want: %v", got, test.want)
			}
		})
	}
}

func Test_Hash256_String(t *testing.T) {
	t.Parallel()

	tests := [1]struct {
		name string
		h256 Hash256
		want string
	}{
		{
			name: "Test_Hash256_String_OK",
			h256: h256,
			want: string(h256[:]),
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := test.h256.String(); got != test.want {
				t.Errorf("String() got: %v | want: %v", got, test.want)
			}
		})
	}
}
