// Copyright © 2020-2021 The EVEN Solutions Developers Team

package zero_test

import (
	"testing"

	. "github.com/platsko/go-kit/zero"
)

func TestBytea32(t *testing.T) {
	t.Parallel()

	data := [32]byte{}
	for i := 0; i < len(data); i++ {
		data[i] = byte(i)
	}

	tests := [1]struct {
		name string
		data [32]byte
	}{
		{
			name: "TestBytea32_OK",
			data: data,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			Bytea32(&test.data)
			for i := range test.data {
				if test.data[i] != 0 {
					t.Errorf("Bytea32() got: %v | want: [32]byte zeros filled", test.data)
					return
				}
			}
		})
	}
}

func TestBytea64(t *testing.T) {
	t.Parallel()

	data := [64]byte{}
	for i := 0; i < len(data); i++ {
		data[i] = byte(i)
	}

	tests := [1]struct {
		name string
		data [64]byte
	}{
		{
			name: "TestBytea64_OK",
			data: data,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			Bytea64(&test.data)
			for i := range test.data {
				if test.data[i] != 0 {
					t.Errorf("Bytea64() got: %v | want: [64]byte zeros filled", test.data)
					return
				}
			}
		})
	}
}

func TestBytes(t *testing.T) {
	t.Parallel()

	data := make([]byte, 512)
	for i := 0; i < len(data); i++ {
		data[i] = byte(i)
	}

	tests := [1]struct {
		name string
		data []byte
	}{
		{
			name: "TestBytes_OK",
			data: data,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			Bytes(test.data)
			for i := range test.data {
				if test.data[i] != 0 {
					t.Errorf("Bytes() got: %v | want: []byte zeros filled", test.data)
					return
				}
			}
		})
	}
}
