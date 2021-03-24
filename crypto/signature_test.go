// Copyright © 2020-2021 The EVEN Solutions Developers Team

package crypto_test

import (
	"encoding/hex"
	"reflect"
	"testing"

	json "github.com/json-iterator/go"
	"google.golang.org/protobuf/proto"

	. "github.com/evenlab/go-kit/crypto"
	"github.com/evenlab/go-kit/crypto/proto/pb"
)

func Benchmark_NewSignature(b *testing.B) {
	sign, _ := mockSignature(Ed25519)
	blob, err := sign.Raw()
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewSignature(blob)
	}
}

func Benchmark_DecodeSignature(b *testing.B) {
	sign, _ := mockSignature(Ed25519)
	pbuf := sign.Encode()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = DecodeSignature(pbuf)
	}
}

func Benchmark_signature_Decode(b *testing.B) {
	sign, _ := mockSignature(Ed25519)
	pbuf := sign.Encode()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewSignature(nil).Decode(pbuf)
	}
}

func Benchmark_signature_Encode(b *testing.B) {
	sign, _ := mockSignature(Ed25519)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sign.Encode()
	}
}

func Benchmark_signature_Equals(b *testing.B) {
	sign, _ := mockSignature(Ed25519)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sign.Equals(sign)
	}
}

func Benchmark_signature_Marshal(b *testing.B) {
	sign, _ := mockSignature(Ed25519)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := sign.Marshal(); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_signature_MarshalJSON(b *testing.B) {
	sign, _ := mockSignature(Ed25519)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := sign.MarshalJSON(); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_signature_Raw(b *testing.B) {
	sign, _ := mockSignature(Ed25519)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := sign.Raw(); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_signature_String(b *testing.B) {
	sign, _ := mockSignature(Ed25519)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sign.String()
	}
}

func Benchmark_signature_Unmarshal(b *testing.B) {
	sign, _ := mockSignature(Ed25519)
	blob, _ := sign.Marshal()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := NewSignature(nil).Unmarshal(blob); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_signature_UnmarshalJSON(b *testing.B) {
	sign, _ := mockSignature(Ed25519)
	blob, _ := sign.MarshalJSON()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := NewSignature(nil).UnmarshalJSON(blob); err != nil {
			b.Fatal(err)
		}
	}
}

func Test_DecodeSignature(t *testing.T) {
	t.Parallel()

	type (
		testCase struct {
			name string
			pbuf *pb.Signature
			want Signature
		}
		testList []testCase
	)

	algos := GetAlgos()
	tests := make(testList, 0, algos.Len())
	for name, algo := range algos {
		sign, _ := mockSignature(algo)
		tests = append(tests, testCase{
			name: name + "_OK",
			pbuf: sign.Encode(),
			want: sign,
		})
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := DecodeSignature(test.pbuf); !reflect.DeepEqual(got, test.want) {
				t.Errorf("Decode() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}

func Test_signature_Equals(t *testing.T) {
	t.Parallel()

	type (
		testCase struct {
			name string
			sign Signature
			blob []byte
			want bool
		}
		testList []testCase
	)

	algos := GetAlgos()
	tests := make(testList, 0, algos.Len()*2)
	for name, algo := range algos {
		sign, _ := mockSignature(algo)
		blob, _ := sign.Raw()
		tests = append(tests, testCase{
			name: name + "_TRUE",
			sign: sign,
			blob: blob,
			want: true,
		}, testCase{
			name: name + "_FALSE",
			sign: sign,
			blob: []byte(name),
		})
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := NewSignature(test.blob)
			if got := got.Equals(test.sign); got != test.want {
				t.Errorf("Equals() got: %v | want: %v", got, test.want)
			}
		})
	}
}

func Test_signature_Marshal(t *testing.T) {
	t.Parallel()

	type (
		testCase struct {
			name    string
			blob    []byte
			want    []byte
			wantErr bool
		}
		testList []testCase
	)

	algos := GetAlgos()
	tests := make(testList, 0, algos.Len())
	for name, algo := range algos {
		sign, _ := mockSignature(algo)
		blob, _ := sign.Raw()
		want, _ := proto.Marshal(&pb.Signature{Blob: blob})
		tests = append(tests, testCase{
			name: name + "_OK",
			blob: blob,
			want: want,
		})
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, err := NewSignature(test.blob).Marshal()
			if (err != nil) != test.wantErr {
				t.Errorf("Marshal() error: %v | want: %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Marshal() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}

func Test_signature_MarshalJSON(t *testing.T) {
	t.Parallel()

	type (
		testCase struct {
			name    string
			blob    []byte
			want    []byte
			wantErr bool
		}
		testList []testCase
	)

	algos := GetAlgos()
	tests := make(testList, 0, algos.Len())
	for name, algo := range algos {
		sign, _ := mockSignature(algo)
		blob, _ := sign.Raw()
		want, _ := json.Marshal(&pb.Signature{Blob: blob})
		tests = append(tests, testCase{
			name: name + "_OK",
			blob: blob,
			want: want,
		})
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, err := NewSignature(test.blob).MarshalJSON()
			if (err != nil) != test.wantErr {
				t.Errorf("MarshalJSON() error: %v | want: %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("MarshalJSON() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}

func Test_signature_Decode(t *testing.T) {
	t.Parallel()

	type (
		testCase struct {
			name string
			want Signature
			pbuf *pb.Signature
		}
		testList []testCase
	)

	algos := GetAlgos()
	tests := make(testList, 0, algos.Len())
	for name, algo := range algos {
		sign, _ := mockSignature(algo)
		tests = append(tests, testCase{
			name: name + "_OK",
			pbuf: sign.Encode(),
			want: sign,
		})
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := NewSignature(nil)
			got.Decode(test.pbuf)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Decode() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}

func Test_signature_Encode(t *testing.T) {
	t.Parallel()

	type (
		testCase struct {
			name string
			sign Signature
			want *pb.Signature
		}

		testList []testCase
	)

	algos := GetAlgos()
	tests := make(testList, 0, algos.Len())
	for name, algo := range algos {
		sign, _ := mockSignature(algo)
		blob, _ := sign.Raw()
		tests = append(tests, testCase{
			name: name + "_OK",
			sign: sign,
			want: &pb.Signature{Blob: blob},
		})
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := test.sign.Encode(); !reflect.DeepEqual(got, test.want) {
				t.Errorf("Encode() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}

func Test_signature_Raw(t *testing.T) {
	t.Parallel()

	type (
		testCase struct {
			name    string
			sign    Signature
			want    []byte
			wantErr bool
		}
		testList []testCase
	)

	algos := GetAlgos()
	tests := make(testList, 0, algos.Len()+1)
	for name, algo := range algos {
		signable, _ := mockSignable(algo, 1024)
		tests = append(tests, testCase{
			name: name + "_OK",
			sign: NewSignature(signable.Blob),
			want: signable.Blob,
		})
	}

	tests = append(tests, testCase{
		name:    "ERR",
		sign:    NewSignature(nil),
		wantErr: true,
	})

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, err := test.sign.Raw()
			if (err != nil) != test.wantErr {
				t.Errorf("Raw() error: %v | want: %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Raw() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}

func Test_signature_String(t *testing.T) {
	t.Parallel()

	type (
		testCase struct {
			name string
			sign Signature
			want string
		}
		testList []testCase
	)

	algos := GetAlgos()
	tests := make(testList, 0, algos.Len())
	for name, algo := range algos {
		sign, _ := mockSignature(algo)
		blob, _ := sign.Raw()
		tests = append(tests, testCase{
			name: name + "_OK",
			sign: sign,
			want: hex.EncodeToString(blob),
		}, testCase{
			name: name + "_nil_OK",
			sign: NewSignature(nil),
			want: "<nil>",
		})
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := test.sign.String(); got != test.want {
				t.Errorf("String() got: %v | want: %v", got, test.want)
			}
		})
	}
}

func Test_signature_Unmarshal(t *testing.T) {
	t.Parallel()

	type (
		testCase struct {
			name    string
			blob    []byte
			want    Signature
			wantErr bool
		}
		testList []testCase
	)

	algos := GetAlgos()
	tests := make(testList, 0, algos.Len()+1)
	for name, algo := range algos {
		sign, _ := mockSignature(algo)
		blob, _ := sign.Marshal()
		tests = append(tests, testCase{
			name: name + "_OK",
			blob: blob,
			want: sign,
		})
	}

	tests = append(tests, testCase{
		name:    "ERR",
		blob:    []byte(":"), // invalid data
		wantErr: true,
	})

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := NewSignature(nil)
			if err := got.Unmarshal(test.blob); (err != nil) != test.wantErr {
				t.Errorf("Unmarshal() error: %v | want: %v", err, test.wantErr)
			}
			if !got.Equals(test.want) {
				t.Errorf("Unmarshal() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}

func Test_signature_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	type (
		testCase struct {
			name    string
			blob    []byte
			want    Signature
			wantErr bool
		}
		testList []testCase
	)

	algos := GetAlgos()
	tests := make(testList, 0, algos.Len()+1)
	for name, algo := range algos {
		sign, _ := mockSignature(algo)
		blob, _ := sign.MarshalJSON()
		tests = append(tests, testCase{
			name: name,
			blob: blob,
			want: sign,
		})
	}

	tests = append(tests, testCase{
		name:    "ERR",
		blob:    []byte(":"), // invalid json
		wantErr: true,
	})

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := NewSignature(nil)
			if err := got.UnmarshalJSON(test.blob); (err != nil) != test.wantErr {
				t.Errorf("UnmarshalJSON() error: %v | want: %v", err, test.wantErr)
			}
			if !got.Equals(test.want) {
				t.Errorf("UnmarshalJSON() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}
