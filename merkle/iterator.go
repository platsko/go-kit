// Copyright © 2020-2021 The EVEN Solutions Developers Team

package merkle

import (
	"github.com/evenlab/go-kit/crypto"
	"github.com/evenlab/go-kit/errors"
)

type (
	// Iterable represents interface for types that implements Iterator interface.
	Iterable interface {
		// Iterator returns Iterator interface for types which can iterate by themselves.
		MerkleIter() Iterator
	}

	// Iterator represents interface for types which can iterate by themselves.
	Iterator interface {
		// HasherNext returns Hasher interface for the current item of list,
		// and then move the internal cursor the the next list's item.
		// NOTICE: returns nil if there is no current element.
		HasherNext() crypto.Hasher

		// HasNext returns true if next element is not nil and false if is.
		HasNext() bool

		// Len returns length of all elements in Iterator list.
		Len() int

		// Rewind rewinds the internal cursor to the first item.
		Rewind() Iterator
	}

	// IterItem implements crypto.Hasher interface.
	IterItem []byte

	// IterStub implements Iterator interface.
	IterStub struct {
		Curr int
		List []IterItem
	}
)

var (
	// Make sure IterStub implements Iterator interface.
	_ Iterator = (*IterStub)(nil)
)

// NewIterator constructs Iterator interface over given data.
func NewIterator(data ...[]byte) Iterator {
	items := make([]IterItem, len(data))
	for idx, blob := range data {
		items[idx] = make(IterItem, len(blob))
		copy(items[idx], blob)
	}

	return &IterStub{List: items}
}

// Hash implements crypto.Hasher interface.
func (m IterItem) Hash() (crypto.Hash256, error) {
	if len(m) == 0 {
		return crypto.Hash256{}, errors.ErrZeroSizeValue()
	}

	return crypto.NewHash256(m), nil
}

// HasherNext implements Iterator.HasherNext method of interface.
func (m *IterStub) HasherNext() crypto.Hasher {
	if !m.HasNext() {
		return nil
	}

	item := m.List[m.Curr]
	m.Curr++

	return &item
}

// HasNext implements Iterator.HasNext of interface.
func (m *IterStub) HasNext() bool {
	return m.List != nil && m.Curr < len(m.List)
}

// Len implements Iterator.Len of interface.
func (m *IterStub) Len() int {
	return len(m.List)
}

// Rewind implements Iterator.Rewind method of interface.
func (m *IterStub) Rewind() Iterator {
	m.Curr = 0

	return m
}
