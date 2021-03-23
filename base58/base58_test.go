// Copyright © 2020-2021 The EVEN Solutions Developers Team

package base58_test

import (
	"encoding/hex"
	"reflect"
	"testing"

	. "github.com/platsko/go-kit/base58"
	"github.com/platsko/go-kit/bytes"
)

const (
	strBase58 = "TTe8GAjHDwbcnY1MYsBjNkBanp9GgyzPK8PxePH7zayyp"
)

func Benchmark_Decode(b *testing.B) {
	base58 := []byte(strBase58)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := Decode(base58); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_DecodeString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := DecodeString(strBase58); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Encode(b *testing.B) {
	blob := bytes.RandBytes(28)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Encode(blob)
	}
}

func Benchmark_EncodeToString(b *testing.B) {
	blob := bytes.RandBytes(28)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = EncodeToString(blob)
	}
}

func Test_Decode(t *testing.T) {
	t.Parallel()

	type (
		testCase struct {
			name    string
			base    []byte
			want    []byte
			wantErr bool
		}
		testList []testCase
	)

	hexCases := mockTestCaseDecode()
	errCases := mockTestCaseDecodeErr()
	tests := make(testList, 0, len(hexCases)+len(errCases))

	for _, c := range hexCases {
		want, err := hex.DecodeString(c.want)
		if err != nil {
			t.Errorf("hex.DecodeString() error: %v | want: %v", err, nil)
			continue
		}
		tests = append(tests, testCase{
			name: c.base + "_OK",
			base: []byte(c.base),
			want: want,
		})
	}

	for _, c := range errCases {
		tests = append(tests, testCase{
			name:    c.base + "_ERR",
			base:    []byte(c.base),
			wantErr: true,
		})
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, err := Decode(test.base)
			if (err != nil) != test.wantErr {
				t.Errorf("Decode() error: %v | want: %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Decode() got = %v | want: %v", got, test.want)
			}
		})
	}
}

func Test_DecodeString(t *testing.T) {
	t.Parallel()

	type (
		testCase struct {
			name    string
			base    string
			want    []byte
			wantErr bool
		}
		testList []testCase
	)

	hexCases := mockTestCaseDecode()
	errCases := mockTestCaseDecodeErr()
	tests := make(testList, 0, len(hexCases)+len(errCases))

	for _, c := range hexCases {
		want, err := hex.DecodeString(c.want)
		if err != nil {
			t.Errorf("hex.DecodeString() error: %v | want: %v", err, nil)
			continue
		}
		tests = append(tests, testCase{
			name: c.base + "_OK",
			base: c.base,
			want: want,
		})
	}

	for _, c := range errCases {
		tests = append(tests, testCase{
			name:    c.base + "_ERR",
			base:    c.base,
			wantErr: true,
		})
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, err := DecodeString(test.base)
			if (err != nil) != test.wantErr {
				t.Errorf("DecodeString() error: %v | want: %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("DecodeString() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}

func Test_Encode(t *testing.T) {
	t.Parallel()

	type (
		testCase struct {
			name string
			blob []byte
			want []byte
		}
		testList []testCase
	)

	cases := mockTestCaseEncode()
	tests := make(testList, 0, len(cases))
	for _, c := range cases {
		tests = append(tests, testCase{
			name: c.base + "_OK",
			blob: []byte(c.base),
			want: []byte(c.want),
		})
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := Encode(test.blob); !reflect.DeepEqual(got, test.want) {
				t.Errorf("Encode() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}

func Test_EncodeToString(t *testing.T) {
	t.Parallel()

	type (
		testCase struct {
			name string
			blob []byte
			want string
		}
		testList []testCase
	)

	cases := mockTestCaseEncode()
	tests := make(testList, 0, len(cases))
	for _, c := range cases {
		tests = append(tests, testCase{
			name: c.base + "_OK",
			blob: []byte(c.base),
			want: c.want,
		})
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := EncodeToString(test.blob); got != test.want {
				t.Errorf("EncodeToString() pass: %#v got: %v | want: %v", test.blob, got, test.want)
			}
		})
	}
}
