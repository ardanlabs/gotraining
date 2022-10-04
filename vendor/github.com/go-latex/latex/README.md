# latex

[![go.dev reference](https://pkg.go.dev/badge/github.com/go-latex/latex)](https://pkg.go.dev/github.com/go-latex/latex)
[![GitHub release](https://img.shields.io/github/release/go-latex/latex.svg)](https://github.com/go-latex/latex/releases)
[![CI](https://github.com/go-latex/latex/workflows/CI/badge.svg)](https://github.com/go-latex/latex/actions)
[![codecov](https://codecov.io/gh/go-latex/latex/branch/master/graph/badge.svg)](https://codecov.io/gh/go-latex/latex)
[![GoDoc](https://godoc.org/github.com/go-latex/latex?status.svg)](https://godoc.org/github.com/go-latex/latex)
[![License](https://img.shields.io/badge/License-BSD--3-blue.svg)](https://github.com/go-latex/latex/raw/master/LICENSE)

`latex` is a package holding Go tools for [LaTeX](https://www.latex-project.org/).

`latex` is supposed to provide features akin to `MathJax` or `matplotlib`'s `TeX` capabilities.
_ie:_ it is supposed to be able to draw mathematical equations, in pure-Go.

`latex` is *NOT SUPPOSED* to be a complete typesetting system like `LaTeX` or `TeX`.

For this, please take look at:

- [Star-TeX](https://star-tex.org)
- [Star-TeX (git repo)](https://git.sr.ht/~sbinet/star-tex)

Eventually, `go-latex/latex` might just use `star-tex.org/...` to provide the `MathJax`-like capabilities.
(once `star-tex.org/...` is ready and exports a nice Go API.)

## Installation

```
$> go get github.com/go-latex/latex/...
```

## Documentation

Documentation is served by [godoc](https://godoc.org), here:

- [godoc.org/github.com/go-latex/latex](https://godoc.org/github.com/go-latex/latex)

The main use case for `go-latex/latex` is to draw a mathematical equation.
This is typically achieved via the `latex/mtex.Render` function that knows how to render mathematical `TeX` equations to a renderer interface.

### Example

```go
package main

import (
	"os"

	"github.com/go-latex/latex/drawtex/drawimg"
	"github.com/go-latex/latex/mtex"
)

func main() {
	f, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	dst := drawimg.NewRenderer(f)
	err = mtex.Render(dst, `$f(x) = \frac{\sqrt{x +20}}{2\pi} +\hbar \sum y\partial y$`, 12, 72, nil)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
}
```

## LICENSE

BSD-3.
