// Copyright © 2020-2021 The EVEN Solutions Developers Team

package base58_test

import (
	"strings"

	. "github.com/platsko/go-kit/base58"
)

type (
	testMock struct {
		base string
		want string
	}
)

func mockTestCaseCheckBase58() [11]testMock {
	return [11]testMock{
		{
			base: "",
			want: "3MNQE1X",
		},
		{
			base: " ",
			want: "B2Kr6dBE",
		},
		{
			base: "-",
			want: "B3jv1Aft",
		},
		{
			base: "0",
			want: "B482yuaX",
		},
		{
			base: "1",
			want: "B4CmeGAC",
		},
		{
			base: "-1",
			want: "mM7eUf6kB",
		},
		{
			base: "11",
			want: "mP7BMTDVH",
		},
		{
			base: "abc",
			want: "4QiVtDjUdeq",
		},
		{
			base: "1234598760",
			want: "ZmNb8uQn5zvnUohNCEPP",
		},
		{
			base: "abcdefghijklmnopqrstuvwxyz",
			want: "K2RYDcKfupxwXdWhSAxQPCeiULntKm63UXyx5MvEH2",
		},
		{
			base: "00000000000000000000000000000000000000000000000000000000000000",
			want: "bi1EWXwJay2udZVxLJozuTb8Meg4W9c6xnmJaRDjg6pri5MBAxb9XwrpQXbtnqEoRV5U2pixnFfwyXC8tRAVC8XxnjK",
		},
	}
}

func mockTestCaseDecode() [11]testMock {
	return [11]testMock{
		{
			base: "2g",
			want: "61", // hex to string encoded
		},
		{
			base: "a3gV",
			want: "626262", // hex to string encoded
		},
		{
			base: "aPEr",
			want: "636363", // hex to string encoded
		},
		{
			base: "2cFupjhnEsSn59qHXstmK2ffpLv2",
			want: "73696d706c792061206c6f6e6720737472696e67", // hex to string encoded
		},
		{
			base: "1NS17iag9jJgTHD1VXjvLCEnZuQ3rJDE9L",
			want: "00eb15231dfceb60925886b67d065299925915aeb172c06647", // hex to string encoded
		},
		{
			base: "ABnLTmg",
			want: "516b6fcd0f", // hex to string encoded
		},
		{
			base: "3SEo3LWLoPntC",
			want: "bf4f89001e670274dd", // hex to string encoded
		},
		{
			base: "3EFU7m",
			want: "572e4794", // hex to string encoded
		},
		{
			base: "EJDM8drfXA6uyA",
			want: "ecac89cad93923c02321", // hex to string encoded
		},
		{
			base: "Rt5zm",
			want: "10c8511e", // hex to string encoded
		},
		{
			base: "1111111111",
			want: "00000000000000000000", // hex to string encoded
		},
	}
}

func mockTestCaseDecodeErr() []testMock {
	cases := make([]testMock, 0, 256-len(Alphabet))
	for i := 0; i < 256; i++ {
		s := string(rune(i))
		if strings.Contains(Alphabet, s) {
			continue
		}
		cases = append(cases, testMock{base: s})
	}

	return cases
}

func mockTestCaseEncode() [11]testMock {
	return [11]testMock{
		{
			base: "",
			want: "",
		},
		{
			base: " ",
			want: "Z",
		},
		{
			base: "0",
			want: "q",
		},
		{
			base: "1",
			want: "r",
		},
		{
			base: string(byte(0)),
			want: "1",
		},
		{
			base: "-1",
			want: "4SU",
		},
		{
			base: "11",
			want: "4k8",
		},
		{
			base: "abc",
			want: "ZiCa",
		},
		{
			base: "1234598760",
			want: "3mJr7AoUXx2Wqd",
		},
		{
			base: "abcdefghijklmnopqrstuvwxyz",
			want: "3yxU3u1igY8WkgtjK92fbJQCd4BZiiT1v25f",
		},
		{
			base: "00000000000000000000000000000000000000000000000000000000000000",
			want: "3sN2THZeE9Eh9eYrwkvZqNstbHGvrxSAM7gXUXvyFQP8XvQLUqNCS27icwUeDT7ckHm4FUHM2mTVh1vbLmk7y",
		},
	}
}
