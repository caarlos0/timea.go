# timea.go

[![Build Status](https://img.shields.io/github/workflow/status/caarlos0/timea.go/build?style=for-the-badge)](https://github.com/caarlos0/timea.go/actions?workflow=build)
[![Coverage Status](https://img.shields.io/codecov/c/gh/caarlos0/timea.go.svg?logo=codecov&style=for-the-badge)](https://codecov.io/gh/caarlos0/timea.go)
[![](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=for-the-badge)](https://pkg.go.dev/github.com/caarlos0/timea.go/v6)


`timea.go` (_did you see what I did there?_) is a simple library to print given times in "time ago" manner.

Example:

```sh
go get github.com/caarlos0/timea.go
```

Simple usage:

```go
package main

import (
	"fmt"
	"time"

	timeago "github.com/caarlos0/timea.go"
)

func main() {
	fmt.Println(timeago.Of(time.Now().Add(-5 * time.Second)))
}
```
