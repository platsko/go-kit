// Copyright © 2020-2021 The EVEN Solutions Developers Team

package merkle_test

import (
	"log"
	"reflect"
	"testing"

	"github.com/evenlab/go-kit/bytes"
	"github.com/evenlab/go-kit/crypto"
	"github.com/evenlab/go-kit/errors"
	. "github.com/evenlab/go-kit/merkle"
)

func Benchmark_NewIterator(b *testing.B) {
	size := 1000
	data := make([][]byte, size)
	for idx := range data {
		data[idx] = make([]byte, 1024)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewIterator(data...)
	}
}

func Benchmark_Iterator_HasherNext(b *testing.B) {
	size := 1000000
	data := make([][]byte, size)
	for idx := range data {
		data[idx] = make([]byte, 1024)
	}
	iter := NewIterator(data...)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = iter.HasherNext()
	}
}

func Benchmark_Iterator_HasherNext_Hash(b *testing.B) {
	size := 1000000
	data := make([][]byte, size)
	blob := bytes.RandBytes(1024)
	for idx := range data {
		data[idx] = make([]byte, 1024)
		copy(data[idx], blob)
	}
	iter := NewIterator(data...)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if item := iter.HasherNext(); item != nil {
			if _, err := item.Hash(); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func Benchmark_Iterator_HasNext(b *testing.B) {
	iter := NewIterator(make([]byte, 0))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = iter.HasNext()
	}
}

func Benchmark_Len(b *testing.B) {
	iter := NewIterator(make([]byte, 0))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = iter.Len()
	}
}

func Benchmark_Rewind(b *testing.B) {
	iter := NewIterator()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = iter.Rewind()
	}
}

func Test_NewIterator(t *testing.T) {
	t.Parallel()

	size := 1024
	data, list := make([][]byte, size), make([]IterItem, size)
	for idx := range data {
		data[idx] = bytes.RandBytes(size)
		list[idx] = make(IterItem, size)
		copy(list[idx], data[idx])
	}

	tests := [1]struct {
		name string
		data [][]byte
		want Iterator
	}{
		{
			name: "OK",
			data: data,
			want: &IterStub{List: list},
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := NewIterator(test.data...); !reflect.DeepEqual(got, test.want) {
				t.Errorf("NewIterator() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}

func Test_Iterator_HasherNext(t *testing.T) {
	t.Parallel()

	size := 1024
	blob := bytes.RandBytes(size)
	item := IterItem(blob)

	tests := [2]struct {
		name string
		iter Iterator
		want crypto.Hasher
	}{
		{
			name: "OK",
			iter: NewIterator(blob),
			want: &item,
		},
		{
			name: "nil_OK",
			iter: NewIterator(),
			want: nil,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := test.iter.HasherNext(); !reflect.DeepEqual(got, test.want) {
				t.Errorf("HasherNext() got: %v | want: %v", got, test.want)
			}
		})
	}
}

func Test_Iterator_HasherNext_Hash(t *testing.T) {
	t.Parallel()

	size := 1024
	blob := bytes.RandBytes(size)
	tests := [2]struct {
		name    string
		iter    Iterator
		want    crypto.Hash256
		wantErr error
	}{
		{
			name: "OK",
			iter: NewIterator(blob),
			want: crypto.NewHash256(blob),
		},
		{
			name:    errors.ErrZeroSizeValueMsg + "_ERR",
			iter:    NewIterator(make([]byte, 0)),
			wantErr: errors.ErrZeroSizeValue(),
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, err := test.iter.HasherNext().Hash()
			if !errors.Is(err, test.wantErr) {
				t.Errorf("Hash() error: %v | want: %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Hash() got = %v | want: %v", got, test.want)
			}
		})
	}
}

func Test_Iterator_HasNext(t *testing.T) {
	t.Parallel()

	tests := [2]struct {
		name string
		iter Iterator
		want bool
	}{
		{
			name: "TRUE",
			iter: NewIterator(make([]byte, 0)),
			want: true,
		},
		{
			name: "FALSE",
			iter: NewIterator(),
			want: false,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := test.iter.HasNext(); got != test.want {
				t.Errorf("HasNext() got: %v | want: %v", got, test.want)
			}
		})
	}
}

func Test_Iterator_Len(t *testing.T) {
	t.Parallel()

	tests := [2]struct {
		name string
		iter Iterator
		want int
	}{
		{
			name: "OK",
			iter: NewIterator(make([]byte, 0)),
			want: 1,
		},
		{
			name: "zero_OK",
			iter: NewIterator(),
			want: 0,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := test.iter.Len(); got != test.want {
				t.Errorf("Len() got: %v | want: %v", got, test.want)
			}
		})
	}
}

func Test_Iterator_Rewind(t *testing.T) {
	t.Parallel()

	size := 1024
	blob := bytes.RandBytes(size)
	iter := NewIterator(blob)

	for iter.HasNext() { // enforce fast forward of iter list
		_ = iter.HasherNext()
	}

	tests := []struct {
		name string
		iter Iterator
		want Iterator
	}{
		{
			name: "OK",
			iter: iter,
			want: NewIterator(blob),
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if reflect.DeepEqual(test.iter, test.want) {
				t.Errorf("Rewind() got: %v | unwant: %v", test.iter, test.want)
			}
			if got := test.iter.Rewind(); !reflect.DeepEqual(got, test.want) {
				t.Errorf("Rewind() got: %v | want: %v", got, test.want)
			}
		})
	}
}
