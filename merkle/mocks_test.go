// Copyright © 2020-2021 The EVEN Solutions Developers Team

package merkle_test

import (
	"log"

	"github.com/evenlab/go-kit/bytes"
	"github.com/evenlab/go-kit/crypto"
	. "github.com/evenlab/go-kit/merkle"
)

func mockTreeStoreCase1() (TreeStore, Iterator) {
	const size = 1
	var err error

	iter, h256 := mockIterable(size), crypto.Hash256{}

	h := make([]crypto.Hash256, size)
	h[0], err = iter.HasherNext().Hash()
	if err != nil {
		log.Fatal(err)
	}

	root := crypto.NewHash256(h[0][:], h[0][:]) // double one use h[0] because we've no h[1]

	tree := TreeStore{
		h[0], h256,
		root,
	}

	// rewind the internal cursor of the iter
	// before return it for futures use
	return tree, iter.Rewind()
}

func mockTreeStoreCase5() (TreeStore, Iterator) {
	const size = 5
	var err error

	iter, h256 := mockIterable(size), crypto.Hash256{}

	h := make([]crypto.Hash256, size)
	for idx := 0; iter.HasNext(); idx++ {
		h[idx], err = iter.HasherNext().Hash()
		if err != nil {
			log.Fatal(err)
		}
	}

	h0h1 := crypto.NewHash256(h[0][:], h[1][:])
	h2h3 := crypto.NewHash256(h[2][:], h[3][:])
	h4h5 := crypto.NewHash256(h[4][:], h[4][:]) // double one use h[4] because we've no h[5]

	h0h1h2h3 := crypto.NewHash256(h0h1[:], h2h3[:])
	h4h5h6h7 := crypto.NewHash256(h4h5[:], h4h5[:]) // double one use h4h5 because we've no h6h7

	root := crypto.NewHash256(h0h1h2h3[:], h4h5h6h7[:])

	tree := TreeStore{
		h[0], h[1], h[2], h[3], h[4], h256, h256, h256,
		h0h1, h2h3, h4h5, h256,
		h0h1h2h3, h4h5h6h7,
		root,
	}

	// rewind the internal cursor of the iter
	// before return it for futures use
	return tree, iter.Rewind()
}

func mockIterable(size int) Iterator {
	items := make([][]byte, size)
	for idx := range items {
		items[idx] = bytes.RandBytes(1)
	}

	return NewIterator(items...)
}
