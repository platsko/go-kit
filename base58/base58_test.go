// Copyright © 2020-2021 The EVEN Solutions Developers Team

package base58_test

import (
	"encoding/hex"
	"testing"

	"github.com/platsko/go-kit/base58"
	"github.com/platsko/go-kit/bytes"
)

type (
	test struct {
		in  string
		out string
	}
)

func BenchmarkDecode(tb *testing.B) {
	in := "1NS17iag9jJgTHD1VXjvLCEnZuQ3rJDE9L"
	tb.ResetTimer()
	for i := 0; i < tb.N; i++ {
		if _, err := base58.DecodeString(in); err != nil {
			tb.Fatal(err)
		}
	}
}

func BenchmarkEncode(tb *testing.B) {
	in := bytes.RandBytes(160)
	tb.ResetTimer()
	for i := 0; i < tb.N; i++ {
		_ = base58.EncodeToString(in)
	}
}

func TestBase58(t *testing.T) {
	t.Parallel()

	for x, test := range makeStringTests() {
		b := []byte(test.in)
		if res := base58.EncodeToString(b); res != test.out {
			t.Errorf("EncodeToString() test #%d failed: got: %v | want: %v", x, res, test.out)
			continue
		}
	}

	for x, test := range makeHexTests() {
		b, err := hex.DecodeString(test.in)
		if err != nil {
			t.Errorf("hex.DecodeString() failed failed #%d: got: %v", x, test.in)
			continue
		}
		if res, _ := base58.DecodeString(test.out); !bytes.Equal(res, b) {
			t.Errorf("DecodeString() test #%d failed: got: %v | want: %v", x, res, test.in)
			continue
		}
	}

	for x, test := range makeInvalidStringTests() {
		if _, err := base58.DecodeString(test.in); err == nil {
			t.Errorf("DecodeString() invalidString test #%d error: %v | want: %v", x, err, base58.ErrUnknownFormat())
			continue
		}
	}
}

func makeStringTests() []test {
	return []test{
		{
			in:  "",
			out: "",
		},
		{
			in:  " ",
			out: "Z",
		},
		{
			in:  "-",
			out: "n",
		},
		{
			in:  "0",
			out: "q",
		},
		{
			in:  "1",
			out: "r",
		},
		{
			in:  "-1",
			out: "4SU",
		},
		{
			in:  "11",
			out: "4k8",
		},
		{
			in:  "abc",
			out: "ZiCa",
		},
		{
			in:  "1234598760",
			out: "3mJr7AoUXx2Wqd",
		},
		{
			in:  "abcdefghijklmnopqrstuvwxyz",
			out: "3yxU3u1igY8WkgtjK92fbJQCd4BZiiT1v25f",
		},
		{
			in:  "00000000000000000000000000000000000000000000000000000000000000",
			out: "3sN2THZeE9Eh9eYrwkvZqNstbHGvrxSAM7gXUXvyFQP8XvQLUqNCS27icwUeDT7ckHm4FUHM2mTVh1vbLmk7y",
		},
	}
}

func makeInvalidStringTests() []test {
	return []test{
		{
			in: "0",
		},
		{
			in: "O",
		},
		{

			in: "I",
		},
		{
			in: "l",
		},
		{
			in: "3mJr0",
		},
		{

			in: "O3yxU",
		},
		{
			in: "3sNI",
		},
		{
			in: "4kl8",
		},
		{
			in: "0OIl",
		},
		{
			in: "!@#$%^&*()-_=+~`",
		},
	}
}

func makeHexTests() []test {
	return []test{
		{
			in:  "61",
			out: "2g",
		},
		{
			in:  "626262",
			out: "a3gV",
		},
		{
			in:  "636363",
			out: "aPEr",
		},
		{
			in:  "73696d706c792061206c6f6e6720737472696e67",
			out: "2cFupjhnEsSn59qHXstmK2ffpLv2",
		},
		{
			in:  "00eb15231dfceb60925886b67d065299925915aeb172c06647",
			out: "1NS17iag9jJgTHD1VXjvLCEnZuQ3rJDE9L",
		},
		{
			in:  "516b6fcd0f",
			out: "ABnLTmg",
		},
		{
			in:  "bf4f89001e670274dd",
			out: "3SEo3LWLoPntC",
		},
		{
			in:  "572e4794",
			out: "3EFU7m",
		},
		{
			in:  "ecac89cad93923c02321",
			out: "EJDM8drfXA6uyA",
		},
		{
			in:  "10c8511e",
			out: "Rt5zm",
		},
		{
			in:  "00000000000000000000",
			out: "1111111111",
		},
	}
}
