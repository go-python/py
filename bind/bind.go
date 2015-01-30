// Copyright 2015 The go-python Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bind

import (
	"bytes"
	"fmt"
	"go/token"
	"io"
	"os"

	"golang.org/x/tools/go/types"
)

// ErrorList is a list of errors
type ErrorList []error

func (list ErrorList) Error() string {
	buf := new(bytes.Buffer)
	for i, err := range list {
		if i > 0 {
			buf.WriteRune('\n')
		}
		io.WriteString(buf, err.Error())
	}
	return buf.String()
}

// GenCPython generates a (C)Python package from a Go package
func GenCPython(w io.Writer, fset *token.FileSet, pkg *types.Package) error {
	buf := new(bytes.Buffer)
	gen := &cpyGen{
		printer: &printer{buf: buf, indentEach: []byte("\t")},
		fset:    fset,
		pkg:     pkg,
	}
	err := gen.gen()
	if err != nil {
		return err
	}

	_, err = io.Copy(w, gen.buf)

	return err
}

// GenGo generates a cgo package from a Go package
func GenGo(w io.Writer, fset *token.FileSet, pkg *types.Package) error {
	buf := new(bytes.Buffer)
	gen := &goGen{
		printer: &printer{buf: buf, indentEach: []byte("\t")},
		fset:    fset,
		pkg:     pkg,
	}
	err := gen.gen()
	if err != nil {
		return err
	}

	_, err = io.Copy(w, gen.buf)

	return err
}

const (
	doDebug = true
)

func debugf(format string, args ...interface{}) (int, error) {
	if doDebug {
		return fmt.Fprintf(os.Stderr, format, args...)
	}
	return 0, nil
}
