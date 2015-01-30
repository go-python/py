// Copyright 2015 The go-python Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package py

// Dict represents the python dict-like protocol
type Dict interface {
	// Clear empties an existing dictionay of all key-value pairs.
	Clear()

	// Contains determines whether a dictionary contains the key k
	Contains(k Object) bool

	// SetItem inserts the value v into the dictionary with a key k
	SetItem(k, v Object) error

	// DelItem removes the entry with key k from the dictionary
	DelItem(k Object) error

	// GetItem returns the object with key from the dictionary
	GetItem(k Object) (Object, error)

	// Items returns all the items from the dictionary
	Items() [][2]Object

	// Keys returns all the keys from the dictionary
	Keys() []Object

	// Values returns all the values from the dictionary
	Values() []Object

	// Size returns the number of items in the dictionary
	Size() int
}
