# Protobuf wrapperspb helpers

[![pkg-img]][pkg-url]
[![coverage-img]][coverage-url]

Provides helpers functions converting [wrapperspb](https://pkg.go.dev/google.golang.org/protobuf/types/known/wrapperspb)

## Installation
Go version 1.18+
```bash
go get github.com/itcomusic/protowrap
```

## Usage
```go
package main

import (
	"fmt"

	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/itcomusic/protowrap"
)

func main() {
	i := protowrap.IntF32V[int](&wrapperspb.Int32Value{Value: 1})
	fmt.Println(*i) // 1

	ints := protowrap.Ints[int]([]int32{1, 2, 3})
	fmt.Println(ints) // [1, 2, 3]

	wi32 := protowrap.Int32Value(i)
	fmt.Println(wi32) // value:1
}
```

[pkg-img]: https://pkg.go.dev/badge/github.com/itcomusic/protowrap.svg
[pkg-url]: https://pkg.go.dev/github.com/itcomusic/protowrap
[coverage-img]: https://codecov.io/gh/itcomusic/protowrap/branch/main/graph/badge.svg
[coverage-url]: https://codecov.io/gh/itcomusic/protowrap