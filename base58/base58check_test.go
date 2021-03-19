// Copyright © 2020-2021 The EVEN Solutions Developers Team

package base58_test

import (
	"crypto/sha256"
	"reflect"
	"testing"

	. "github.com/platsko/go-kit/base58"
	"github.com/platsko/go-kit/bytes"
	"github.com/platsko/go-kit/errors"
)

func Benchmark_CheckDecode(tb *testing.B) {
	base58 := []byte(strBase58)
	tb.ResetTimer()
	for i := 0; i < tb.N; i++ {
		if _, _, err := CheckDecode(base58); err != nil {
			tb.Fatal(err)
		}
	}
}

func Benchmark_CheckEncode(tb *testing.B) {
	blob, ver, err := CheckDecode([]byte(strBase58))
	if err != nil {
		tb.Fatal(err)
	}
	tb.ResetTimer()
	for i := 0; i < tb.N; i++ {
		_ = CheckEncode(blob, ver)
	}
}

func Benchmark_Checksum(tb *testing.B) {
	base58 := []byte(strBase58)
	tb.ResetTimer()
	for i := 0; i < tb.N; i++ {
		_ = Checksum(base58)
	}
}

func Test_CheckDecode(t *testing.T) {
	t.Parallel()

	type (
		testCase struct {
			name    string
			blob    []byte
			want    []byte
			wantVer byte
			wantErr error
		}
		testList []testCase
	)

	cases := mockTestCaseCheckBase58()
	tests := make(testList, 0, len(cases)+7)
	for _, c := range cases {
		tests = append(tests, testCase{
			name:    c.want + "_OK",
			blob:    []byte(c.want),
			want:    []byte(c.base),
			wantVer: byte(20),
		})
	}

	// append test unknown format error
	tests = append(tests, testCase{
		name:    ErrUnknownFormatMsg + "_ERR",
		blob:    []byte("3MNQE10"),
		wantErr: ErrUnknownFormat(),
	})

	// append test checksum error
	tests = append(tests, testCase{
		name:    ErrChecksumMismatchMsg + "_ERR",
		blob:    []byte("3MNQE1Y"),
		wantErr: ErrChecksumMismatch(),
	})

	// append tests invalid formats error - string with size below 5
	// mean the version byte or the checksum bytes are missing
	for size := 0; size < 5; size++ {
		blob := make([]byte, size)
		for idx := range blob {
			blob[idx] = '1'
		}
		tests = append(tests, testCase{
			name:    ErrInvalidFormatMsg + "_ERR",
			blob:    blob,
			wantErr: ErrInvalidFormat(),
		})
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, ver, err := CheckDecode(test.blob)
			if err != nil && !errors.Is(err, test.wantErr) {
				t.Errorf("CheckDecode() error: %v | want: %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("CheckDecode() got: %#v | want: %#v", got, test.want)
			}
			if ver != test.wantVer {
				t.Errorf("CheckDecode() ver: %#v | want: %v", ver, test.wantVer)
			}
		})
	}
}

func Test_CheckEncode(t *testing.T) {
	t.Parallel()

	type (
		testCase struct {
			name string
			base []byte
			ver  byte
			want string
		}
		testList []testCase
	)

	cases := mockTestCaseCheckBase58()
	tests := make(testList, 0, len(cases))
	for _, c := range cases {
		tests = append(tests, testCase{
			name: c.base + "_OK",
			base: []byte(c.base),
			ver:  byte(20),
			want: c.want,
		})
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := CheckEncode(test.base, test.ver); got != test.want {
				t.Errorf("CheckEncode() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}

func Test_Checksum(t *testing.T) {
	t.Parallel()

	blob, want := bytes.RandBytes(1024), [ChecksumSize]byte{}
	h256 := sha256.Sum256(blob)
	h256 = sha256.Sum256(h256[:])
	copy(want[:], h256[:ChecksumSize])

	tests := [1]struct {
		name string
		blob []byte
		want [ChecksumSize]byte
	}{
		{
			name: "OK",
			blob: blob,
			want: want,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := Checksum(test.blob); !reflect.DeepEqual(got, test.want) {
				t.Errorf("Checksum() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}
