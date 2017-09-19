Gonum LAPACK [![GoDoc](https://godoc.org/gonum.org/v1/gonum/lapack?status.svg)](https://godoc.org/gonum.org/v1/gonum/lapack)
======

A collection of packages to provide LAPACK functionality for the Go programming
language (http://golang.org). This provides a partial implementation in native go
and a wrapper using cgo to a c-based implementation.

## Installation

```
  go get gonum.org/v1/gonum/lapack/...
```


Install OpenBLAS:
```
  git clone https://github.com/xianyi/OpenBLAS
  cd OpenBLAS
  make
```

Then install the lapack/cgo package:
```sh
  CGO_LDFLAGS="-L/path/to/OpenBLAS -lopenblas" go install gonum.org/v1/netlib/lapack
```

For Windows you can download binary packages for OpenBLAS at
http://sourceforge.net/projects/openblas/files/

If you want to use a different BLAS package such as the Intel MKL you can
adjust the `CGO_LDFLAGS` variable:
```sh
  CGO_LDFLAGS="-lmkl_rt" go install gonum.org/v1/netlib/lapack
```

## Packages

### lapack

Defines the LAPACK API based on http://www.netlib.org/lapack/lapacke.html

### lapack/gonum

Go implementation of the LAPACK API (incomplete, implements the `float64` API).

### lapack/lapack64

Wrappers for an implementation of the double (i.e., `float64`) precision real parts of
the LAPACK API.

