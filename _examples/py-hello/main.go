// Copyright 2015 The go-python Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/go-python/py"
)

func init() {
	err := py.Initialize()
	if err != nil {
		panic(err)
	}
}

func main() {
	gostr := "foo"
	pystr := py.NewString(gostr)
	fmt.Printf("hello [%v]\n", pystr)
}
