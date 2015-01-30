py
==

`py` is a high-level API wrapping the low-level `CPython` C-API, for `go`.

## Installation

```sh
$ go get github.com/go-python/py
```


## Documentation

Documentation is available on [godoc](https://godoc.org):

 [github.com/go-python/py](https://godoc.org/github.com/go-python/py)


## Examples

```go
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
```

```sh
$ go run ./main.go
hello [foo]
```
