// Copyright © 2020-2021 The EVEN Solutions Developers Team

package crypto_test

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"reflect"
	"testing"

	json "github.com/json-iterator/go"
	cc "github.com/libp2p/go-libp2p-core/crypto"
	"google.golang.org/protobuf/proto"

	. "github.com/evenlab/go-kit/crypto"
	"github.com/evenlab/go-kit/crypto/proto/pb"
)

func Benchmark_NewPublicKey(b *testing.B) {
	_, ki := mockCryptoKeyPair(Ed25519)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewPublicKey(ki)
	}
}

func Benchmark_DecodePublicKey(b *testing.B) {
	_, pbKey := mockGenerateKeyPair(Ed25519)
	pbuf, err := pbKey.Encode()
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = DecodePublicKey(pbuf)
	}
}

func Benchmark_publicKey_Algo(b *testing.B) {
	_, pbKey := mockGenerateKeyPair(Ed25519)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = pbKey.Algo()
	}
}

func Benchmark_publicKey_Base64(b *testing.B) {
	_, pbKey := mockGenerateKeyPair(Ed25519)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := pbKey.Base64(); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_publicKey_Decode(b *testing.B) {
	_, pbKey := mockGenerateKeyPair(Ed25519)
	pbuf, err := pbKey.Encode()
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pbKey = NewPublicKey(nil)
		if err := pbKey.Decode(pbuf); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_publicKey_Encode(b *testing.B) {
	_, pbKey := mockGenerateKeyPair(Ed25519)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := pbKey.Encode(); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_publicKey_Equals(b *testing.B) {
	_, pbKey := mockGenerateKeyPair(Ed25519)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = pbKey.Equals(pbKey)
	}
}

func Benchmark_publicKey_Hash224(b *testing.B) {
	_, pbKey := mockGenerateKeyPair(Ed25519)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := pbKey.Hash224(); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_publicKey_Marshal(b *testing.B) {
	_, pbKey := mockGenerateKeyPair(Ed25519)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := pbKey.Marshal(); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_publicKey_MarshalJSON(b *testing.B) {
	_, pbKey := mockGenerateKeyPair(Ed25519)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := pbKey.MarshalJSON(); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_publicKey_Raw(b *testing.B) {
	_, pbKey := mockGenerateKeyPair(Ed25519)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := pbKey.Raw(); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_publicKey_String(b *testing.B) {
	_, pbKey := mockGenerateKeyPair(Ed25519)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = pbKey.String()
	}
}

func Benchmark_publicKey_Unmarshal(b *testing.B) {
	_, pbKey := mockGenerateKeyPair(Ed25519)
	blob, err := pbKey.Marshal()
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := pbKey.Unmarshal(blob); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_publicKey_UnmarshalJSON(b *testing.B) {
	_, pbKey := mockGenerateKeyPair(Ed25519)
	blob, err := pbKey.MarshalJSON()
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := pbKey.UnmarshalJSON(blob); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_publicKey_Verify(b *testing.B) {
	signable, prKey := mockSignable(Ed25519, 1024)
	pbKey := prKey.PublicKey()
	if _, err := prKey.Sign(signable); err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if ok, _ := pbKey.Verify(signable); !ok {
			b.Fatalf("Verify() got: %v | want: true", ok)
		}
	}
}

func Test_NewPublicKey(t *testing.T) {
	t.Parallel()

	type (
		testCase struct {
			name string
			ki   cc.PubKey
			want PublicKey
		}
		testList []testCase
	)

	algos := GetAlgos()
	tests := make(testList, 0, algos.Len()+1)
	for name, algo := range algos {
		_, ki := mockCryptoKeyPair(algo)
		tests = append(tests, testCase{
			name: name + "_OK",
			ki:   ki,
			want: NewPublicKey(ki),
		})
	}

	tests = append(tests, testCase{
		name: "nil_OK",
		ki:   nil,
		want: NewPublicKey(nil),
	})

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := NewPublicKey(test.ki); !reflect.DeepEqual(got, test.want) {
				t.Errorf("NewPrivateKey() got: %v | want: %v", got, test.want)
			}
		})
	}
}

func Test_DecodePublicKey(t *testing.T) {
	t.Parallel()

	type (
		testCase struct {
			name    string
			pbuf    *pb.PublicKey
			want    PublicKey
			wantErr bool
		}
		testList []testCase
	)

	algos := GetAlgos()
	tests := make(testList, 0, algos.Len()*2)
	for name, algo := range algos {
		_, ki := mockCryptoKeyPair(algo)
		blob, _ := cc.MarshalPublicKey(ki)
		pbuf := pb.PublicKey{Blob: blob}
		tests = append(tests, testCase{
			name: name + "_OK",
			pbuf: &pbuf,
			want: NewPublicKey(ki),
		}, testCase{
			name:    name + "_ERR",
			pbuf:    &pb.PublicKey{},
			want:    NewPublicKey(nil),
			wantErr: true,
		})
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, err := DecodePublicKey(test.pbuf)
			if (err != nil) != test.wantErr {
				t.Errorf("DecodePublicKey() error: %v | want: %v", err, test.wantErr)
				return
			}
			if !test.want.Equals(got) {
				t.Errorf("DecodePublicKey() got: %v | want: %v", got, test.want)
			}
		})
	}
}

func Test_publicKey_Algo(t *testing.T) {
	t.Parallel()

	type (
		testCase struct {
			name  string
			pbKey PublicKey
			want  Algo
		}
		testList []testCase
	)

	algos := GetAlgos()
	tests := make(testList, 0, algos.Len()+1)
	for name, algo := range algos {
		_, pbKey := mockGenerateKeyPair(algo)
		tests = append(tests, testCase{
			name:  name + "_OK",
			pbKey: pbKey,
			want:  algo,
		})
	}

	tests = append(tests, testCase{
		name:  "UNKNOWN_OK",
		pbKey: NewPublicKey(nil),
		want:  UNKNOWN,
	})

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := test.pbKey.Algo(); got != test.want {
				t.Errorf("Algo() got: %v | want: %v", got, test.want)
			}
		})
	}
}

func Test_publicKey_Base64(t *testing.T) {
	t.Parallel()

	type (
		testCase struct {
			name    string
			pbKey   PublicKey
			want    []byte
			wantErr bool
		}
		testList []testCase
	)

	algos := GetAlgos()
	tests := make(testList, 0, algos.Len()+1)
	for name, algo := range algos {
		_, ki := mockCryptoKeyPair(algo)
		blob, _ := ki.Raw()
		tests = append(tests, testCase{
			name:  name + "_OK",
			pbKey: NewPublicKey(ki),
			want:  blob,
		})
	}

	tests = append(tests, testCase{
		name:    "ERR",
		pbKey:   NewPublicKey(nil),
		wantErr: true,
	})

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, err := test.pbKey.Base64()
			if (err != nil) != test.wantErr {
				t.Errorf("Base64() error: %v | want: %v", err, test.wantErr)
				return
			}
			if test.wantErr {
				return
			}
			if raw, _ := base64.StdEncoding.DecodeString(got); !reflect.DeepEqual(raw, test.want) {
				t.Errorf("Base64() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}

func Test_publicKey_Decode(t *testing.T) {
	t.Parallel()

	type (
		testCase struct {
			name    string
			pbuf    *pb.PublicKey
			want    PublicKey
			wantErr bool
		}
		testList []testCase
	)

	algos := GetAlgos()
	tests := make(testList, 0, algos.Len()*2)
	for name, algo := range algos {
		_, ki := mockCryptoKeyPair(algo)
		blob, _ := cc.MarshalPublicKey(ki)
		pbuf := pb.PublicKey{Blob: blob}
		tests = append(tests, testCase{
			name: name + "_OK",
			pbuf: &pbuf,
			want: NewPublicKey(ki),
		}, testCase{
			name:    name + "_ERR",
			pbuf:    &pb.PublicKey{},
			want:    NewPublicKey(nil),
			wantErr: true,
		})
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := NewPublicKey(nil)
			if err := got.Decode(test.pbuf); (err != nil) != test.wantErr {
				t.Errorf("Decode() error: %v | want: %v", err, test.wantErr)
				return
			}
			if !got.Equals(test.want) {
				t.Errorf("Decode() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}

func Test_publicKey_Encode(t *testing.T) {
	t.Parallel()

	type (
		testCase struct {
			name    string
			pbKey   PublicKey
			want    *pb.PublicKey
			wantErr bool
		}
		testList []testCase
	)

	algos := GetAlgos()
	tests := make(testList, 0, algos.Len()+1)
	for name, algo := range algos {
		_, ki := mockCryptoKeyPair(algo)
		blob, _ := cc.MarshalPublicKey(ki)
		pbuf := pb.PublicKey{Blob: blob}
		tests = append(tests, testCase{
			name:  name + "_OK",
			pbKey: NewPublicKey(ki),
			want:  &pbuf,
		})
	}

	tests = append(tests, testCase{
		name:    "ERR",
		pbKey:   NewPublicKey(nil),
		wantErr: true,
	})

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, err := test.pbKey.Encode()
			if (err != nil) != test.wantErr {
				t.Errorf("Encode() error: %v | want: %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Encode() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}

func Test_publicKey_Equals(t *testing.T) {
	t.Parallel()

	type (
		testCase struct {
			name  string
			pbKey PublicKey
			equal PublicKey
			want  bool
		}
		testList []testCase
	)

	algos := GetAlgos()
	tests := make(testList, 0, algos.Len()*2)
	for name, algo := range algos {
		_, notEq := mockGenerateKeyPair(algo)
		_, pbKey := mockGenerateKeyPair(algo)
		tests = append(tests, testCase{
			name:  name + "_TRUE",
			pbKey: pbKey,
			equal: pbKey,
			want:  true,
		}, testCase{
			name:  name + "_FALSE",
			pbKey: notEq, // there is not equal public key
			equal: pbKey,
		})
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := test.pbKey.Equals(test.equal); got != test.want {
				t.Errorf("Equals() got: %v | want: %v", got, test.want)
			}
		})
	}
}

func Test_publicKey_Hash224(t *testing.T) {
	t.Parallel()

	type (
		testCase struct {
			name    string
			pbKey   PublicKey
			want    Hash224
			wantErr bool
		}
		testList []testCase
	)

	algos := GetAlgos()
	tests := make(testList, 0, algos.Len()+1)
	for name, algo := range algos {
		_, ki := mockCryptoKeyPair(algo)
		blob, _ := ki.Raw()
		h256 := sha256.Sum256(blob)
		h224 := sha256.Sum224(h256[:])
		tests = append(tests, testCase{
			name:  name + "_OK",
			pbKey: NewPublicKey(ki),
			want:  h224,
		})
	}

	tests = append(tests, testCase{
		name:    "ERR",
		pbKey:   NewPublicKey(nil),
		wantErr: true,
	})

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, err := test.pbKey.Hash224()
			if (err != nil) != test.wantErr {
				t.Errorf("Hash224() error: %v | want: %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Hash224() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}

func Test_publicKey_Marshal(t *testing.T) {
	t.Parallel()

	type (
		testCase struct {
			name    string
			pbKey   PublicKey
			want    []byte
			wantErr bool
		}
		testList []testCase
	)

	algos := GetAlgos()
	tests := make(testList, 0, algos.Len()+1)
	for name, algo := range algos {
		_, pbKey := mockGenerateKeyPair(algo)
		pbuf, _ := pbKey.Encode()
		blob, _ := proto.Marshal(pbuf)
		tests = append(tests, testCase{
			name:  name + "_OK",
			pbKey: pbKey,
			want:  blob,
		})
	}

	tests = append(tests, testCase{
		name:    "ERR",
		pbKey:   NewPublicKey(nil),
		wantErr: true,
	})

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, err := test.pbKey.Marshal()
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

func Test_publicKey_MarshalJSON(t *testing.T) {
	t.Parallel()

	type (
		testCase struct {
			name    string
			pbKey   PublicKey
			want    []byte
			wantErr bool
		}
		testList []testCase
	)

	algos := GetAlgos()
	tests := make(testList, 0, algos.Len()+1)
	for name, algo := range algos {
		_, pbKey := mockGenerateKeyPair(algo)
		pbuf, _ := pbKey.Encode()
		blob, _ := json.Marshal(pbuf)
		tests = append(tests, testCase{
			name:  name + "_OK",
			pbKey: pbKey,
			want:  blob,
		})
	}

	tests = append(tests, testCase{
		name:    "ERR",
		pbKey:   NewPublicKey(nil),
		wantErr: true,
	})

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, err := test.pbKey.MarshalJSON()
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

func Test_publicKey_Raw(t *testing.T) {
	t.Parallel()

	type (
		testCase struct {
			name    string
			pbKey   PublicKey
			want    []byte
			wantErr bool
		}
		testList []testCase
	)

	algos := GetAlgos()
	tests := make(testList, 0, algos.Len()+1)
	for name, algo := range algos {
		_, ki := mockCryptoKeyPair(algo)
		blob, _ := ki.Raw()
		tests = append(tests, testCase{
			name:  name + "_OK",
			pbKey: NewPublicKey(ki),
			want:  blob,
		})
	}

	tests = append(tests, testCase{
		name:    "ERR",
		pbKey:   NewPublicKey(nil),
		wantErr: true,
	})

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, err := test.pbKey.Raw()
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

func Test_publicKey_String(t *testing.T) {
	t.Parallel()

	type (
		testCase struct {
			name  string
			pbKey PublicKey
			want  string
		}
		testList []testCase
	)

	algos := GetAlgos()
	tests := make(testList, 0, algos.Len()+1)
	for name, algo := range algos {
		_, ki := mockCryptoKeyPair(algo)
		blob, _ := ki.Raw()
		tests = append(tests, testCase{
			name:  name + "_OK",
			pbKey: NewPublicKey(ki),
			want:  hex.EncodeToString(blob),
		})
	}

	tests = append(tests, testCase{
		name:  "nil_OK",
		pbKey: NewPublicKey(nil),
		want:  "<nil>",
	})

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := test.pbKey.String(); got != test.want {
				t.Errorf("String() got: %v | want: %v", got, test.want)
			}
		})
	}
}

func Test_publicKey_Unmarshal(t *testing.T) {
	t.Parallel()

	type (
		testCase struct {
			name    string
			blob    []byte
			want    PublicKey
			wantErr bool
		}
		testList []testCase
	)

	algos := GetAlgos()
	tests := make(testList, 0, algos.Len()+1)
	for name, algo := range algos {
		_, pbKey := mockGenerateKeyPair(algo)
		blob, _ := pbKey.Marshal()
		tests = append(tests, testCase{
			name: name + "_OK",
			blob: blob,
			want: pbKey,
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

			got := NewPublicKey(nil)
			if err := got.Unmarshal(test.blob); (err != nil) != test.wantErr {
				t.Errorf("Unmarshal() error: %v | want: %v", err, test.wantErr)
				return
			}
			if !got.Equals(test.want) {
				t.Errorf("Unmarshal() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}

func Test_publicKey_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	type (
		testCase struct {
			name    string
			blob    []byte
			want    PublicKey
			wantErr bool
		}
		testList []testCase
	)

	algos := GetAlgos()
	tests := make(testList, 0, algos.Len()+1)
	for name, algo := range algos {
		_, pbKey := mockGenerateKeyPair(algo)
		blob, _ := pbKey.MarshalJSON()
		tests = append(tests, testCase{
			name: name + "_OK",
			blob: blob,
			want: pbKey,
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

			got := NewPublicKey(nil)
			if err := got.UnmarshalJSON(test.blob); (err != nil) != test.wantErr {
				t.Errorf("UnmarshalJSON() error: %v | want: %v", err, test.wantErr)
			}
			if !got.Equals(test.want) {
				t.Errorf("UnmarshalJSON() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}

func Test_publicKey_Verify(t *testing.T) {
	t.Parallel()

	type (
		testCase struct {
			name     string
			pbKey    PublicKey
			signable Signable
			want     bool
			wantErr  bool
		}
		testList []testCase
	)

	algos := GetAlgos()
	tests := make(testList, 0, algos.Len()*2+5)
	for name, algo := range algos {
		_, notEq := mockGenerateKeyPair(algo)
		signable, prKey := mockSignable(algo, 1024)
		tests = append(tests, testCase{
			name:     name + "_TRUE",
			pbKey:    prKey.PublicKey(),
			signable: signable,
			want:     true,
		}, testCase{
			name:     name + "_FALSE",
			pbKey:    notEq,
			signable: signable,
			wantErr:  algo == RSA,
		})
	}

	signable, prKey := mockSignable(Ed25519, 1024)
	pbKey := prKey.PublicKey()
	tests = append(tests, testCase{
		name:     "nil_ki_PublicKey_ERR",
		pbKey:    NewPublicKey(nil),
		signable: nil,
		wantErr:  true,
	}, testCase{
		name:     "nil_Signable_ERR",
		pbKey:    pbKey,
		signable: nil,
		wantErr:  true,
	}, testCase{
		name:     "nil_blob_Signable_ERR",
		pbKey:    pbKey,
		signable: &SignableStub{Sign: signable.GetSignature()},
		wantErr:  true,
	}, testCase{
		name:     "nil_Signature_ERR",
		pbKey:    pbKey,
		signable: &SignableStub{},
		wantErr:  true,
	}, testCase{
		name:     "nil_blob_Signature_ERR",
		pbKey:    pbKey,
		signable: &SignableStub{Sign: NewSignature(nil)},
		wantErr:  true,
	})

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, err := test.pbKey.Verify(test.signable)
			if (err != nil) != test.wantErr {
				t.Errorf("Verify() error: %v | want: %v", err, test.wantErr)
				return
			}
			if got != test.want {
				t.Errorf("Verify() got: %v | want: %v", got, test.want)
			}
		})
	}
}
