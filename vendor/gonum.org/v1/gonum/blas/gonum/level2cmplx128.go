// Copyright ©2017 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gonum

import (
	"math/cmplx"

	"gonum.org/v1/gonum/blas"
	"gonum.org/v1/gonum/internal/asm/c128"
)

// Zgemv performs one of the matrix-vector operations
//  y = alpha * A * x + beta * y    if trans = blas.NoTrans
//  y = alpha * A^T * x + beta * y  if trans = blas.Trans
//  y = alpha * A^H * x + beta * y  if trans = blas.ConjTrans
// where alpha and beta are scalars, x and y are vectors, and A is an m×n dense matrix.
func (Implementation) Zgemv(trans blas.Transpose, m, n int, alpha complex128, a []complex128, lda int, x []complex128, incX int, beta complex128, y []complex128, incY int) {
	checkZMatrix('A', m, n, a, lda)
	switch trans {
	default:
		panic(badTranspose)
	case blas.NoTrans:
		checkZVector('x', n, x, incX)
		checkZVector('y', m, y, incY)
	case blas.Trans, blas.ConjTrans:
		checkZVector('x', m, x, incX)
		checkZVector('y', n, y, incY)
	}

	if m == 0 || n == 0 || (alpha == 0 && beta == 1) {
		return
	}

	var lenX, lenY int
	if trans == blas.NoTrans {
		lenX = n
		lenY = m
	} else {
		lenX = m
		lenY = n
	}
	var kx int
	if incX < 0 {
		kx = (1 - lenX) * incX
	}
	var ky int
	if incY < 0 {
		ky = (1 - lenY) * incY
	}

	// Form y := beta*y.
	if beta != 1 {
		if incY == 1 {
			if beta == 0 {
				for i := range y {
					y[i] = 0
				}
			} else {
				c128.ScalUnitary(beta, y[:lenY])
			}
		} else {
			iy := ky
			if beta == 0 {
				for i := 0; i < lenY; i++ {
					y[iy] = 0
					iy += incY
				}
			} else {
				if incY > 0 {
					c128.ScalInc(beta, y, uintptr(lenY), uintptr(incY))
				} else {
					c128.ScalInc(beta, y, uintptr(lenY), uintptr(-incY))
				}
			}
		}
	}

	if alpha == 0 {
		return
	}

	switch trans {
	default:
		// Form y := alpha*A*x + y.
		iy := ky
		if incX == 1 {
			for i := 0; i < m; i++ {
				y[iy] += alpha * c128.DotuUnitary(a[i*lda:i*lda+n], x[:n])
				iy += incY
			}
			return
		}
		for i := 0; i < m; i++ {
			y[iy] += alpha * c128.DotuInc(a[i*lda:i*lda+n], x, uintptr(n), 1, uintptr(incX), 0, uintptr(kx))
			iy += incY
		}
		return

	case blas.Trans:
		// Form y := alpha*A^T*x + y.
		ix := kx
		if incY == 1 {
			for i := 0; i < m; i++ {
				c128.AxpyUnitary(alpha*x[ix], a[i*lda:i*lda+n], y[:n])
				ix += incX
			}
			return
		}
		for i := 0; i < m; i++ {
			c128.AxpyInc(alpha*x[ix], a[i*lda:i*lda+n], y, uintptr(n), 1, uintptr(incY), 0, uintptr(ky))
			ix += incX
		}
		return

	case blas.ConjTrans:
		// Form y := alpha*A^H*x + y.
		ix := kx
		if incY == 1 {
			for i := 0; i < m; i++ {
				tmp := alpha * x[ix]
				for j := 0; j < n; j++ {
					y[j] += tmp * cmplx.Conj(a[i*lda+j])
				}
				ix += incX
			}
			return
		}
		for i := 0; i < m; i++ {
			tmp := alpha * x[ix]
			jy := ky
			for j := 0; j < n; j++ {
				y[jy] += tmp * cmplx.Conj(a[i*lda+j])
				jy += incY
			}
			ix += incX
		}
		return
	}
}

// Zgerc performs the rank-one operation
//  A += alpha * x * y^H
// where A is an m×n dense matrix, alpha is a scalar, x is an m element vector,
// and y is an n element vector.
func (Implementation) Zgerc(m, n int, alpha complex128, x []complex128, incX int, y []complex128, incY int, a []complex128, lda int) {
	checkZMatrix('A', m, n, a, lda)
	checkZVector('x', m, x, incX)
	checkZVector('y', n, y, incY)

	if m == 0 || n == 0 || alpha == 0 {
		return
	}

	var kx, jy int
	if incX < 0 {
		kx = (1 - m) * incX
	}
	if incY < 0 {
		jy = (1 - n) * incY
	}
	for j := 0; j < n; j++ {
		if y[jy] != 0 {
			tmp := alpha * cmplx.Conj(y[jy])
			c128.AxpyInc(tmp, x, a[j:], uintptr(m), uintptr(incX), uintptr(lda), uintptr(kx), 0)
		}
		jy += incY
	}
}

// Zgeru performs the rank-one operation
//  A += alpha * x * y^T
// where A is an m×n dense matrix, alpha is a scalar, x is an m element vector,
// and y is an n element vector.
func (Implementation) Zgeru(m, n int, alpha complex128, x []complex128, incX int, y []complex128, incY int, a []complex128, lda int) {
	checkZMatrix('A', m, n, a, lda)
	checkZVector('x', m, x, incX)
	checkZVector('y', n, y, incY)

	if m == 0 || n == 0 || alpha == 0 {
		return
	}

	var kx int
	if incX < 0 {
		kx = (1 - m) * incX
	}
	if incY == 1 {
		for i := 0; i < m; i++ {
			if x[kx] != 0 {
				tmp := alpha * x[kx]
				c128.AxpyUnitary(tmp, y[:n], a[i*lda:i*lda+n])
			}
			kx += incX
		}
		return
	}
	var jy int
	if incY < 0 {
		jy = (1 - n) * incY
	}
	for i := 0; i < m; i++ {
		if x[kx] != 0 {
			tmp := alpha * x[kx]
			c128.AxpyInc(tmp, y, a[i*lda:i*lda+n], uintptr(n), uintptr(incY), 1, uintptr(jy), 0)
		}
		kx += incX
	}
}

// Zher performs the Hermitian rank-one operation
//  A += alpha * x * x^H
// where A is an n×n Hermitian matrix, alpha is a real scalar, and x is an n
// element vector. On entry, the imaginary parts of the diagonal elements of A
// are ignored and assumed to be zero, on return they will be set to zero.
func (Implementation) Zher(uplo blas.Uplo, n int, alpha float64, x []complex128, incX int, a []complex128, lda int) {
	if uplo != blas.Upper && uplo != blas.Lower {
		panic(badUplo)
	}
	checkZMatrix('A', n, n, a, lda)
	checkZVector('x', n, x, incX)

	if n == 0 || alpha == 0 {
		return
	}

	var kx int
	if incX < 0 {
		kx = (1 - n) * incX
	}
	if uplo == blas.Upper {
		if incX == 1 {
			for i := 0; i < n; i++ {
				if x[i] != 0 {
					tmp := complex(alpha*real(x[i]), alpha*imag(x[i]))
					aii := real(a[i*lda+i])
					xtmp := real(tmp * cmplx.Conj(x[i]))
					a[i*lda+i] = complex(aii+xtmp, 0)
					for j := i + 1; j < n; j++ {
						a[i*lda+j] += tmp * cmplx.Conj(x[j])
					}
				} else {
					aii := real(a[i*lda+i])
					a[i*lda+i] = complex(aii, 0)
				}
			}
			return
		}

		ix := kx
		for i := 0; i < n; i++ {
			if x[ix] != 0 {
				tmp := complex(alpha*real(x[ix]), alpha*imag(x[ix]))
				aii := real(a[i*lda+i])
				xtmp := real(tmp * cmplx.Conj(x[ix]))
				a[i*lda+i] = complex(aii+xtmp, 0)
				jx := ix + incX
				for j := i + 1; j < n; j++ {
					a[i*lda+j] += tmp * cmplx.Conj(x[jx])
					jx += incX
				}
			} else {
				aii := real(a[i*lda+i])
				a[i*lda+i] = complex(aii, 0)
			}
			ix += incX
		}
		return
	}

	if incX == 1 {
		for i := 0; i < n; i++ {
			if x[i] != 0 {
				tmp := complex(alpha*real(x[i]), alpha*imag(x[i]))
				for j := 0; j < i; j++ {
					a[i*lda+j] += tmp * cmplx.Conj(x[j])
				}
				aii := real(a[i*lda+i])
				xtmp := real(tmp * cmplx.Conj(x[i]))
				a[i*lda+i] = complex(aii+xtmp, 0)
			} else {
				aii := real(a[i*lda+i])
				a[i*lda+i] = complex(aii, 0)
			}
		}
		return
	}

	ix := kx
	for i := 0; i < n; i++ {
		if x[ix] != 0 {
			tmp := complex(alpha*real(x[ix]), alpha*imag(x[ix]))
			jx := kx
			for j := 0; j < i; j++ {
				a[i*lda+j] += tmp * cmplx.Conj(x[jx])
				jx += incX
			}
			aii := real(a[i*lda+i])
			xtmp := real(tmp * cmplx.Conj(x[ix]))
			a[i*lda+i] = complex(aii+xtmp, 0)

		} else {
			aii := real(a[i*lda+i])
			a[i*lda+i] = complex(aii, 0)
		}
		ix += incX
	}
}

// Zher2 performs the Hermitian rank-two operation
//  A += alpha*x*y^H + conj(alpha)*y*x^H
// where alpha is a scalar, x and y are n element vectors and A is an n×n
// Hermitian matrix. On entry, the imaginary parts of the diagonal elements are
// ignored and assumed to be zero. On return they will be set to zero.
func (Implementation) Zher2(uplo blas.Uplo, n int, alpha complex128, x []complex128, incX int, y []complex128, incY int, a []complex128, lda int) {
	if uplo != blas.Upper && uplo != blas.Lower {
		panic(badUplo)
	}
	checkZMatrix('A', n, n, a, lda)
	checkZVector('x', n, x, incX)
	checkZVector('y', n, y, incY)

	if n == 0 || alpha == 0 {
		return
	}

	var kx, ky int
	var ix, iy int
	if incX != 1 || incY != 1 {
		if incX < 0 {
			kx = (1 - n) * incX
		}
		if incY < 0 {
			ky = (1 - n) * incY
		}
		ix = kx
		iy = ky
	}
	if uplo == blas.Upper {
		if incX == 1 && incY == 1 {
			for i := 0; i < n; i++ {
				if x[i] != 0 || y[i] != 0 {
					tmp1 := alpha * x[i]
					tmp2 := cmplx.Conj(alpha) * y[i]
					aii := real(a[i*lda+i]) + real(tmp1*cmplx.Conj(y[i])) + real(tmp2*cmplx.Conj(x[i]))
					a[i*lda+i] = complex(aii, 0)
					for j := i + 1; j < n; j++ {
						a[i*lda+j] += tmp1*cmplx.Conj(y[j]) + tmp2*cmplx.Conj(x[j])
					}
				} else {
					aii := real(a[i*lda+i])
					a[i*lda+i] = complex(aii, 0)
				}
			}
			return
		}
		for i := 0; i < n; i++ {
			if x[i] != 0 || y[i] != 0 {
				tmp1 := alpha * x[ix]
				tmp2 := cmplx.Conj(alpha) * y[iy]
				aii := real(a[i*lda+i]) + real(tmp1*cmplx.Conj(y[iy])) + real(tmp2*cmplx.Conj(x[ix]))
				a[i*lda+i] = complex(aii, 0)
				jx := ix + incX
				jy := iy + incY
				for j := i + 1; j < n; j++ {
					a[i*lda+j] += tmp1*cmplx.Conj(y[jy]) + tmp2*cmplx.Conj(x[jx])
					jx += incX
					jy += incY
				}
			} else {
				aii := real(a[i*lda+i])
				a[i*lda+i] = complex(aii, 0)
			}
			ix += incX
			iy += incY
		}
		return
	}

	if incX == 1 && incY == 1 {
		for i := 0; i < n; i++ {
			if x[i] != 0 || y[i] != 0 {
				tmp1 := alpha * x[i]
				tmp2 := cmplx.Conj(alpha) * y[i]
				for j := 0; j < i; j++ {
					a[i*lda+j] += tmp1*cmplx.Conj(y[j]) + tmp2*cmplx.Conj(x[j])
				}
				aii := real(a[i*lda+i]) + real(tmp1*cmplx.Conj(y[i])) + real(tmp2*cmplx.Conj(x[i]))
				a[i*lda+i] = complex(aii, 0)
			} else {
				aii := real(a[i*lda+i])
				a[i*lda+i] = complex(aii, 0)
			}
		}
		return
	}
	for i := 0; i < n; i++ {
		if x[i] != 0 || y[i] != 0 {
			tmp1 := alpha * x[ix]
			tmp2 := cmplx.Conj(alpha) * y[iy]
			jx := kx
			jy := ky
			for j := 0; j < i; j++ {
				a[i*lda+j] += tmp1*cmplx.Conj(y[jy]) + tmp2*cmplx.Conj(x[jx])
				jx += incX
				jy += incY
			}
			aii := real(a[i*lda+i]) + real(tmp1*cmplx.Conj(y[iy])) + real(tmp2*cmplx.Conj(x[ix]))
			a[i*lda+i] = complex(aii, 0)
		} else {
			aii := real(a[i*lda+i])
			a[i*lda+i] = complex(aii, 0)
		}
		ix += incX
		iy += incY
	}
}
