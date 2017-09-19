# Gonum BLAS [![GoDoc](https://godoc.org/gonum.org/v1/gonum/blas?status.svg)](https://godoc.org/gonum.org/v1/gonum/blas)

A collection of packages to provide BLAS functionality for the [Go programming
language](http://golang.org)

## Installation
```sh
  go get gonum.org/v1/gonum/blas/...
```

### BLAS C-bindings

If you want to use OpenBLAS, install it in any directory:
```sh
  git clone https://github.com/xianyi/OpenBLAS
  cd OpenBLAS
  make
```

The blas/cgo package provides bindings to C-backed BLAS packages. blas/cgo needs the `CGO_LDFLAGS`
environment variable to point to the blas installation. More information can be found in the
[cgo command documentation](http://golang.org/cmd/cgo/).

Then install the blas/cgo package:
```sh
  CGO_LDFLAGS="-L/path/to/OpenBLAS -lopenblas" go install gonum.org/v1/netlib/blas
```

For Windows you can download binary packages for OpenBLAS at
[SourceForge](http://sourceforge.net/projects/openblas/files/).

If you want to use a different BLAS package such as the Intel MKL you can
adjust the `CGO_LDFLAGS` variable:
```sh
  CGO_LDFLAGS="-lmkl_rt" go install gonum.org/v1/netlib/blas
```

On OS X the easiest solution is to use the libraries provided by the system:
```sh
  CGO_LDFLAGS="-framework Accelerate" go install gonum.org/v1/netlib/blas
```

## Packages

### blas

Defines [BLAS API](http://www.netlib.org/blas/blast-forum/cinterface.pdf) split in several
interfaces.

### blas/gonum

Go implementation of the BLAS API (incomplete, implements the `float32` and `float64` API).

### blas/blas64 and blas/blas32

Wrappers for an implementation of the double (i.e., `float64`) and single (`float32`)
precision real parts of the BLAS API.

```Go
package main

import (
	"fmt"

	"gonum.org/v1/gonum/blas/blas64"
)

func main() {
	v := blas64.Vector{Inc: 1, Data: []float64{1, 1, 1}}
	fmt.Println("v has length:", blas64.Nrm2(len(v.Data), v))
}
```

### blas/cblas128 and blas/cblas64

Wrappers for an implementation of the double (i.e., `complex128`) and single (`complex64`) 
precision complex parts of the blas API.

Currently blas/cblas64 and blas/cblas128 require gonum.org/v1/netlib/blas.
