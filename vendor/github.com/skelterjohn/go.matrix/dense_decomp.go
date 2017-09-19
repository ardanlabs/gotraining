// Copyright 2009 The GoMatrix Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

import "math"

/*
Returns the cholesky decomposition C of A, st CC'=A.
*/
func (A *DenseMatrix) Cholesky() (L *DenseMatrix, err error) {
	n := A.Rows()
	L = Zeros(n, n)
	isspd := A.Cols() == n

	for j := 0; j < n; j++ {
		Lrowj := L.RowCopy(j)
		d := float64(0)
		for k := 0; k < j; k++ {
			Lrowk := L.RowCopy(k)
			s := float64(0)
			for i := 0; i < k; i++ {
				s += Lrowk[i] * Lrowj[i]
			}
			s = (A.Get(j, k) - s) / Lrowk[k]
			Lrowj[k] = s
			L.Set(j, k, s)
			d += s * s
			isspd = isspd && (A.Get(k, j) == A.Get(j, k))
		}
		d = A.Get(j, j) - d
		isspd = isspd && (d > 0.0)
		L.Set(j, j, math.Sqrt(max(d, float64(0))))
		for k := j + 1; k < n; k++ {
			L.Set(j, k, 0)
		}
	}

	if !isspd {
		err = ExceptionNotSPD
	}

	return
}

/*
return L,U,P, st PLU=A.
*/
func (A *DenseMatrix) LU() (L, U *DenseMatrix, P *PivotMatrix) {
	m := A.Rows()
	n := A.Cols()
	C := A.Copy()

	P = C.LUInPlace()

	L = C.L()
	for i := 0; i < m && i < n; i++ {
		L.Set(i, i, 1)
	}
	U = C.U()

	return
}

/*
Overwrites A with [L\U] and returns P, st PLU=A. L is considered to
have 1s in the diagonal.
*/
func (A *DenseMatrix) LUInPlace() (P *PivotMatrix) {
	m := A.Rows()
	n := A.Cols()
	LUcolj := make([]float64, m)
	LUrowi := make([]float64, n)
	piv := make([]int, m)
	for i := 0; i < m; i++ {
		piv[i] = i
	}
	pivsign := float64(1.0)

	for j := 0; j < n; j++ {
		A.BufferCol(j, LUcolj)
		for i := 0; i < m; i++ {
			A.BufferRow(i, LUrowi)
			kmax := i
			if j < i {
				kmax = j
			}
			s := float64(0)
			for k := 0; k < kmax; k++ {
				s += LUrowi[k] * LUcolj[k]
			}
			LUcolj[i] -= s
			LUrowi[j] = LUcolj[i]
			A.Set(i, j, LUrowi[j])
		}

		p := j
		for i := j + 1; i < m; i++ {
			if math.Abs(LUcolj[i]) > math.Abs(LUcolj[p]) {
				p = i
			}
		}
		if p != j {
			A.SwapRows(p, j)
			k := piv[p]
			piv[p] = piv[j]
			piv[j] = k
			pivsign = -pivsign
		}

		if j < m && A.Get(j, j) != 0 {
			for i := j + 1; i < m; i++ {
				A.Set(i, j, A.Get(i, j)/A.Get(j, j))
			}
		}
	}

	P = MakePivotMatrix(piv, pivsign)

	return
}

func (A *DenseMatrix) QR() (Q, R *DenseMatrix) {
	m := A.Rows()
	n := A.Cols()
	QR := A.Copy()
	Q = Zeros(m, n)
	R = Zeros(m, n)
	i, j, k := 0, 0, 0
	norm := float64(0.0)
	s := float64(0.0)

	for k = 0; k < n; k++ {
		norm = 0
		for i = k; i < m; i++ {
			norm = math.Hypot(norm, QR.Get(i, k))
		}

		if norm != 0.0 {
			if QR.Get(k, k) < 0 {
				norm = -norm
			}

			for i = k; i < m; i++ {
				QR.Set(i, k, QR.Get(i, k)/norm)
			}
			QR.Set(k, k, QR.Get(k, k)+1.0)

			for j = k + 1; j < n; j++ {
				s = 0.0
				for i = k; i < m; i++ {
					s += QR.Get(i, k) * QR.Get(i, j)
				}
				s = -s / QR.Get(k, k)
				for i = k; i < m; i++ {
					QR.Set(i, j, QR.Get(i, j)+s*QR.Get(i, k))

					if i < j {
						R.Set(i, j, QR.Get(i, j))
					}
				}

			}
		}

		R.Set(k, k, -norm)

	}

	//Q Matrix:
	i, j, k = 0, 0, 0

	for k = n - 1; k >= 0; k-- {
		Q.Set(k, k, 1.0)
		for j = k; j < n; j++ {
			if QR.Get(k, k) != 0 {
				s = 0.0
				for i = k; i < m; i++ {
					s += QR.Get(i, k) * Q.Get(i, j)
				}
				s = -s / QR.Get(k, k)
				for i = k; i < m; i++ {
					Q.Set(i, j, Q.Get(i, j)+s*QR.Get(i, k))
				}
			}
		}
	}

	return
}
