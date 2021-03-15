// Copyright © 2020-2021 The EVEN Solutions Developers Team

package base58_test

import (
	"testing"

	"github.com/platsko/go-kit/base58"
	"github.com/platsko/go-kit/errors"
)

func TestBase58Check(t *testing.T) {
	t.Parallel()

	tests := [11]struct {
		version byte
		in      string
		out     string
	}{
		{
			version: 20,
			in:      "",
			out:     "3MNQE1X",
		},
		{
			version: 20,
			in:      " ",
			out:     "B2Kr6dBE",
		},
		{
			version: 20,
			in:      "-",
			out:     "B3jv1Aft",
		},
		{
			version: 20,
			in:      "0",
			out:     "B482yuaX",
		},
		{
			version: 20,
			in:      "1",
			out:     "B4CmeGAC",
		},
		{
			version: 20,
			in:      "-1",
			out:     "mM7eUf6kB",
		},
		{
			version: 20,
			in:      "11",
			out:     "mP7BMTDVH",
		},
		{
			version: 20,
			in:      "abc",
			out:     "4QiVtDjUdeq",
		},
		{
			version: 20,
			in:      "1234598760",
			out:     "ZmNb8uQn5zvnUohNCEPP",
		},
		{
			version: 20,
			in:      "abcdefghijklmnopqrstuvwxyz",
			out:     "K2RYDcKfupxwXdWhSAxQPCeiULntKm63UXyx5MvEH2",
		},
		{
			version: 20,
			in:      "00000000000000000000000000000000000000000000000000000000000000",
			out:     "bi1EWXwJay2udZVxLJozuTb8Meg4W9c6xnmJaRDjg6pri5MBAxb9XwrpQXbtnqEoRV5U2pixnFfwyXC8tRAVC8XxnjK",
		},
	}

	for x, test := range tests {
		// test encoding
		if res := base58.CheckEncode([]byte(test.in), test.version); res != test.out {
			t.Errorf("CheckEncode test #%d got: %v | want: %v", x, res, test.out)
		}

		// test decoding
		res, version, err := base58.CheckDecode([]byte(test.out))
		switch {
		case err != nil:
			t.Errorf("CheckDecode() test #%d failed: %v", x, err)
		case version != test.version:
			t.Errorf("CheckDecode() test #%d version got: %v | want: %v", x, version, test.version)
		case string(res) != test.in:
			t.Errorf("CheckDecode() test #%d got: %v | want: %v", x, res, test.in)
		}
	}

	// test the two decoding failure cases
	// case 1: checksum error
	_, _, err := base58.CheckDecode([]byte("3MNQE1Y"))
	if !errors.Is(err, base58.ErrChecksumMismatch()) {
		t.Error("CheckDecode() test failed, expected ErrChecksumMismatch")
	}

	// case 2: invalid formats
	// string lengths below 5 mean the version byte and/or the checksum bytes are missing
	b := []byte("")
	for l := 0; l < 4; l++ {
		_, _, err = base58.CheckDecode(b) // make a string of length `len`
		if !errors.Is(err, base58.ErrInvalidFormat()) {
			t.Error("CheckDecode() test failed, expected ErrInvalidFormat")
		}
	}
}
