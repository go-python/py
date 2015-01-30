// Copyright 2015 The go-python Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package py

import (
	"reflect"
)

// Object represents any Python object and codifies the Python object protocol.
type Object interface {
	// Increment the reference count.
	// The Object may be nil, in which case the function has no effect.
	IncRef()

	// Decrement the reference count.
	// If the Object is nil, nothing happens.
	DecRef()

	// HasAttr returns whether this object has an attribute with name 'n'
	HasAttr(n string) bool

	// GetAttr returns the attribute with name 'n'
	GetAttr(n string) (Object, error)

	// SetAttr sets the value of the attribute named 'n' of the object to the value 'v'
	SetAttr(n string, v Object) error

	// DelAttr deletes the attribute named 'n' of the object
	DelAttr(n string) error

	String() string

	// Call calls a callable Object with args as arguments
	Call(args ...Object) (Object, error)

	// Hash computes and returns the hash value of an Object
	// On failure, Hash returns -1.
	Hash() int64

	// IsTrue returns whether an Object is considered to be true.
	//IsTrue() bool

	// Not returns whether an Object is not considered to be true.
	//Not() bool

	// Type returns the Type of this Object.
	Type() Object
}

// New creates a new py.Object value from a Go value
func New(v interface{}) Object {
	rt := reflect.TypeOf(v)
	switch rt.Kind() {
	case reflect.Int:
		return nil
	}
	return nil
}
