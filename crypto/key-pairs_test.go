// Copyright © 2020-2021 The EVEN Solutions Developers Team

package crypto_test

import (
	"testing"

	. "github.com/platsko/go-kit/crypto"
)

func Benchmark_GenerateKeyPair_Ed25519(tb *testing.B) {
	for i := 0; i < tb.N; i++ {
		if _, _, err := GenerateKeyPair(Ed25519); err != nil {
			tb.Fatal(err)
		}
	}
}

func Benchmark_GenerateKeyPair_ECDSA(tb *testing.B) {
	for i := 0; i < tb.N; i++ {
		if _, _, err := GenerateKeyPair(ECDSA); err != nil {
			tb.Fatal(err)
		}
	}
}

func Benchmark_GenerateKeyPair_RSA(tb *testing.B) {
	for i := 0; i < tb.N; i++ {
		if _, _, err := GenerateKeyPair(RSA); err != nil {
			tb.Fatal(err)
		}
	}
}

func Benchmark_GenerateKeyPair_Secp256k1(tb *testing.B) {
	for i := 0; i < tb.N; i++ {
		if _, _, err := GenerateKeyPair(Secp256k1); err != nil {
			tb.Fatal(err)
		}
	}
}

func Test_GenerateKeyPair(t *testing.T) {
	t.Parallel()

	type (
		testCase struct {
			name      string
			algo      Algo
			wantPrKey PrivateKey
			wantPbKey PublicKey
			wantErr   bool
		}
		testList []testCase
	)

	algos := GetAlgos()
	tests := make(testList, 0, algos.Len()+1)
	for name, algo := range algos {
		tests = append(tests, testCase{
			name: name + "_OK",
			algo: algo,
		})
	}

	tests = append(tests, testCase{
		name:    "ERR",
		algo:    Algo(-1), // nonexistent Algo
		wantErr: true,
	})

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			prKey, pbKey, err := GenerateKeyPair(test.algo)
			if (err != nil) != test.wantErr {
				t.Errorf("GenerateKeyPair() error: %v | want: %v", err, test.wantErr)
				return
			}
			if prKey != nil && prKey.Algo() != test.algo {
				t.Errorf("GenerateKeyPair() got: %#v | want: %#v", prKey, test.wantPrKey)
			}
			if pbKey != nil && pbKey.Algo() != test.algo {
				t.Errorf("GenerateKeyPair() got: %#v | want: %#v", pbKey, test.wantPbKey)
			}
		})
	}
}
