// Copyright © 2020-2021 The EVEN Solutions Developers Team

package merkle_test

import (
	"reflect"
	"testing"

	"github.com/evenlab/go-kit/crypto"
	"github.com/evenlab/go-kit/errors"
	. "github.com/evenlab/go-kit/merkle"
)

func Benchmark_BuildTreeStore(b *testing.B) {
	size := 1000
	list := mockIterable(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := BuildTreeStore(list); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_TreeStore_Root(b *testing.B) {
	tree, _ := mockTreeStoreCase5()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := tree.Root(); err != nil {
			b.Fatal(err)
		}
	}
}

func Test_BuildTreeStore(t *testing.T) {
	t.Parallel()

	empty := make([]byte, 0)
	tree, iter := mockTreeStoreCase5()
	tests := [4]struct {
		name    string
		iter    Iterator
		want    TreeStore
		wantErr error
	}{
		{
			name: "OK",
			iter: iter,
			want: tree,
		},
		{
			name:    errors.ErrNilPointerValueMsg + "_ERR",
			iter:    nil,
			wantErr: errors.ErrNilPointerValue(),
		},
		{
			name:    errors.ErrZeroSizeValueMsg + "_iter_ERR",
			iter:    NewIterator(),
			wantErr: errors.ErrZeroSizeValue(),
		},
		{
			name:    errors.ErrZeroSizeValueMsg + "_item_ERR",
			iter:    NewIterator(empty, empty, empty),
			wantErr: errors.ErrZeroSizeValue(),
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, err := BuildTreeStore(test.iter)
			if !errors.Is(err, test.wantErr) {
				t.Errorf("BuildTreeStore() error: %v | want: %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("BuildTreeStore() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}

func Test_TreeStore_Root(t *testing.T) {
	t.Parallel()

	treeCase1, _ := mockTreeStoreCase1()
	treeCase5, _ := mockTreeStoreCase5()

	tests := [3]struct {
		name    string
		tree    TreeStore
		want    crypto.Hash256
		wantErr error
	}{
		{
			name: "case_1_OK",
			tree: treeCase1,
			want: treeCase1[len(treeCase1)-1], // root is the last element in the the list
		},
		{
			name: "case_5_OK",
			tree: treeCase5,
			want: treeCase5[len(treeCase5)-1], // root is the last element in the the list
		},
		{
			name: ErrMerkleTreeBuiltImproperlyMsg + "_ERR",
			tree: TreeStore{ // mock tree store with size under MinTreeStoreSize
				crypto.Hash256{},
				crypto.Hash256{},
			},
			wantErr: ErrMerkleTreeBuiltImproperly(),
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, err := test.tree.Root()
			if !errors.Is(err, test.wantErr) {
				t.Errorf("Root() error: %v | want: %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Root() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}
