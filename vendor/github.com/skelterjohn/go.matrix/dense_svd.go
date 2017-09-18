// Copyright 2009 The GoMatrix Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

import "math"

/*
Returns U, Σ, V st Σ is diagonal (or block diagonal) and UΣV'=Arg
*/
func (Arg *DenseMatrix) SVD() (theU, Σ, theV *DenseMatrix, err error) {
	//copied from Jama
	// Derived from LINPACK code.
	// Initialize.
	A := Arg.Copy().Arrays()
	m := Arg.rows
	n := Arg.cols

	/* Apparently the failing cases are only a proper subset of (m<n),
		 so let's not throw error.  Correct fix to come later?
	      if (m<n) {
		  throw new IllegalArgumentException("Jama SVD only works for m >= n"); }
	*/

	if m < n {
		err = ErrorDimensionMismatch
		return
	}

	nu := minInt(m, n)
	s := make([]float64, minInt(m+1, n))

	U := make([][]float64, m)
	for i := 0; i < m; i++ {
		U[i] = make([]float64, nu)
	}
	V := make([][]float64, n)
	for i := 0; i < n; i++ {
		V[i] = make([]float64, n)
	}

	e := make([]float64, n)
	work := make([]float64, m)
	wantu := true
	wantv := true

	// Reduce A to bidiagonal form, storing the diagonal elements
	// in s and the super-diagonal elements in e.

	nct := minInt(m-1, n)
	nrt := maxInt(0, minInt(n-2, m))
	for k := 0; k < maxInt(nct, nrt); k++ {
		if k < nct {

			// Compute the transformation for the k-th column and
			// place the k-th diagonal in s[k].
			// Compute 2-norm of k-th column without under/overflow.
			s[k] = 0
			for i := k; i < m; i++ {
				s[k] = math.Hypot(s[k], A[i][k])
			}
			if s[k] != 0.0 {
				if A[k][k] < 0.0 {
					s[k] = -s[k]
				}
				for i := k; i < m; i++ {
					A[i][k] /= s[k]
				}
				A[k][k] += 1.0
			}
			s[k] = -s[k]
		}
		for j := k + 1; j < n; j++ {
			if (k < nct) && (s[k] != 0.0) {

				// Apply the transformation.

				t := float64(0)
				for i := k; i < m; i++ {
					t += A[i][k] * A[i][j]
				}
				t = -t / A[k][k]
				for i := k; i < m; i++ {
					A[i][j] += t * A[i][k]
				}
			}

			// Place the k-th row of A into e for the
			// subsequent calculation of the row transformation.

			e[j] = A[k][j]
		}
		if wantu && (k < nct) {

			// Place the transformation in U for subsequent back
			// multiplication.

			for i := k; i < m; i++ {
				U[i][k] = A[i][k]
			}
		}
		if k < nrt {

			// Compute the k-th row transformation and place the
			// k-th super-diagonal in e[k].
			// Compute 2-norm without under/overflow.
			e[k] = 0
			for i := k + 1; i < n; i++ {
				e[k] = math.Hypot(e[k], e[i])
			}
			if e[k] != 0.0 {
				if e[k+1] < 0.0 {
					e[k] = -e[k]
				}
				for i := k + 1; i < n; i++ {
					e[i] /= e[k]
				}
				e[k+1] += 1.0
			}
			e[k] = -e[k]
			if (k+1 < m) && (e[k] != 0.0) {

				// Apply the transformation.

				for i := k + 1; i < m; i++ {
					work[i] = 0.0
				}
				for j := k + 1; j < n; j++ {
					for i := k + 1; i < m; i++ {
						work[i] += e[j] * A[i][j]
					}
				}
				for j := k + 1; j < n; j++ {
					t := -e[j] / e[k+1]
					for i := k + 1; i < m; i++ {
						A[i][j] += t * work[i]
					}
				}
			}
			if wantv {

				// Place the transformation in V for subsequent
				// back multiplication.

				for i := k + 1; i < n; i++ {
					V[i][k] = e[i]
				}
			}
		}
	}

	// Set up the final bidiagonal matrix or order p.

	p := minInt(n, m+1)
	if nct < n {
		s[nct] = A[nct][nct]
	}
	if m < p {
		s[p-1] = 0.0
	}
	if nrt+1 < p {
		e[nrt] = A[nrt][p-1]
	}
	e[p-1] = 0.0

	// If required, generate U.

	if wantu {
		for j := nct; j < nu; j++ {
			for i := 0; i < m; i++ {
				U[i][j] = 0.0
			}
			U[j][j] = 1.0
		}
		for k := nct - 1; k >= 0; k-- {
			if s[k] != 0.0 {
				for j := k + 1; j < nu; j++ {
					t := float64(0)
					for i := k; i < m; i++ {
						t += U[i][k] * U[i][j]
					}
					t = -t / U[k][k]
					for i := k; i < m; i++ {
						U[i][j] += t * U[i][k]
					}
				}
				for i := k; i < m; i++ {
					U[i][k] = -U[i][k]
				}
				U[k][k] = 1.0 + U[k][k]
				for i := 0; i < k-1; i++ {
					U[i][k] = 0.0
				}
			} else {
				for i := 0; i < m; i++ {
					U[i][k] = 0.0
				}
				U[k][k] = 1.0
			}
		}
	}

	// If required, generate V.

	if wantv {
		for k := n - 1; k >= 0; k-- {
			if (k < nrt) && (e[k] != 0.0) {
				for j := k + 1; j < nu; j++ {
					t := float64(0)
					for i := k + 1; i < n; i++ {
						t += V[i][k] * V[i][j]
					}
					t = -t / V[k+1][k]
					for i := k + 1; i < n; i++ {
						V[i][j] += t * V[i][k]
					}
				}
			}
			for i := 0; i < n; i++ {
				V[i][k] = 0.0
			}
			V[k][k] = 1.0
		}
	}

	// Main iteration loop for the singular values.

	pp := p - 1
	iter := 0
	eps := math.Pow(2.0, -52.0)
	tiny := math.Pow(2.0, -966.0)
	for p > 0 {
		var k, kase int

		// Here is where a test for too many iterations would go.

		// This section of the program inspects for
		// negligible elements in the s and e arrays.  On
		// completion the variables kase and k are set as follows.

		// kase = 1     if s(p) and e[k-1] are negligible and k<p
		// kase = 2     if s(k) is negligible and k<p
		// kase = 3     if e[k-1] is negligible, k<p, and
		//              s(k), ..., s(p) are not negligible (qr step).
		// kase = 4     if e(p-1) is negligible (convergence).

		for k = p - 2; k >= -1; k-- {
			if k == -1 {
				break
			}
			if math.Abs(e[k]) <=
				tiny+eps*(math.Abs(s[k])+math.Abs(s[k+1])) {
				e[k] = 0.0
				break
			}
		}
		if k == p-2 {
			kase = 4
		} else {
			var ks int
			for ks = p - 1; ks >= k; ks-- {
				if ks == k {
					break
				}
				t := float64(0)
				if ks != p {
					t = math.Abs(e[ks])
				}
				if ks != k+1 {
					t += math.Abs(e[ks-1])
				}
				//double t = (ks != p ? Math.abs(e[ks]) : 0.) +
				//           (ks != k+1 ? Math.abs(e[ks-1]) : 0.);
				if math.Abs(s[ks]) <= tiny+eps*t {
					s[ks] = 0.0
					break
				}
			}
			if ks == k {
				kase = 3
			} else if ks == p-1 {
				kase = 1
			} else {
				kase = 2
				k = ks
			}
		}
		k++

		// Perform the task indicated by kase.
		//fmt.Printf("kase = %d\n", kase);
		switch kase {

		// Deflate negligible s(p).

		case 1:
			{
				f := e[p-2]
				e[p-2] = 0.0
				for j := p - 2; j >= k; j-- {
					t := math.Hypot(s[j], f)
					cs := s[j] / t
					sn := f / t
					s[j] = t
					if j != k {
						f = -sn * e[j-1]
						e[j-1] = cs * e[j-1]
					}
					if wantv {
						for i := 0; i < n; i++ {
							t = cs*V[i][j] + sn*V[i][p-1]
							V[i][p-1] = -sn*V[i][j] + cs*V[i][p-1]
							V[i][j] = t
						}
					}
				}
			}
			break

		// Split at negligible s(k).

		case 2:
			{
				f := e[k-1]
				e[k-1] = 0.0
				for j := k; j < p; j++ {
					t := math.Hypot(s[j], f)
					cs := s[j] / t
					sn := f / t
					s[j] = t
					f = -sn * e[j]
					e[j] = cs * e[j]
					if wantu {
						for i := 0; i < m; i++ {
							t = cs*U[i][j] + sn*U[i][k-1]
							U[i][k-1] = -sn*U[i][j] + cs*U[i][k-1]
							U[i][j] = t
						}
					}
				}
			}
			break

		// Perform one qr step.

		case 3:
			{

				// Calculate the shift.

				scale := max(max(max(max(
					math.Abs(s[p-1]), math.Abs(s[p-2])),
					math.Abs(e[p-2])),
					math.Abs(s[k])),
					math.Abs(e[k]))
				sp := s[p-1] / scale
				spm1 := s[p-2] / scale
				epm1 := e[p-2] / scale
				sk := s[k] / scale
				ek := e[k] / scale
				b := ((spm1+sp)*(spm1-sp) + epm1*epm1) / 2.0
				c := (sp * epm1) * (sp * epm1)
				shift := float64(0)
				if (b != 0.0) || (c != 0.0) {
					shift = math.Sqrt(b*b + c)
					if b < 0.0 {
						shift = -shift
					}
					shift = c / (b + shift)
				}
				f := (sk+sp)*(sk-sp) + shift
				g := sk * ek

				// Chase zeros.

				for j := k; j < p-1; j++ {
					t := math.Hypot(f, g)
					cs := f / t
					sn := g / t
					if j != k {
						e[j-1] = t
					}
					f = cs*s[j] + sn*e[j]
					e[j] = cs*e[j] - sn*s[j]
					g = sn * s[j+1]
					s[j+1] = cs * s[j+1]
					if wantv {
						for i := 0; i < n; i++ {
							t = cs*V[i][j] + sn*V[i][j+1]
							V[i][j+1] = -sn*V[i][j] + cs*V[i][j+1]
							V[i][j] = t
						}
					}
					t = math.Hypot(f, g)
					cs = f / t
					sn = g / t
					s[j] = t
					f = cs*e[j] + sn*s[j+1]
					s[j+1] = -sn*e[j] + cs*s[j+1]
					g = sn * e[j+1]
					e[j+1] = cs * e[j+1]
					if wantu && (j < m-1) {
						for i := 0; i < m; i++ {
							t = cs*U[i][j] + sn*U[i][j+1]
							U[i][j+1] = -sn*U[i][j] + cs*U[i][j+1]
							U[i][j] = t
						}
					}
				}
				e[p-2] = f
				iter = iter + 1
			}
			break

		// Convergence.

		case 4:
			{

				// Make the singular values positive.

				if s[k] <= 0.0 {
					if s[k] < 0.0 {
						s[k] = -s[k]
					} else {
						s[k] = 0
					}
					if wantv {
						for i := 0; i <= pp; i++ {
							V[i][k] = -V[i][k]
						}
					}
				}

				// Order the singular values.

				for k < pp {
					if s[k] >= s[k+1] {
						break
					}
					t := s[k]
					s[k] = s[k+1]
					s[k+1] = t
					if wantv && (k < n-1) {
						for i := 0; i < n; i++ {
							t = V[i][k+1]
							V[i][k+1] = V[i][k]
							V[i][k] = t
						}
					}
					if wantu && (k < m-1) {
						for i := 0; i < m; i++ {
							t = U[i][k+1]
							U[i][k+1] = U[i][k]
							U[i][k] = t
						}
					}
					k++
				}
				iter = 0
				p--
			}
			break
		}
	}
	//fmt.Printf("testing\n%v\n%v\n%v\n%v\n%v\n", A, V, U, e, s);

	theU = MakeDenseMatrixStacked(U).GetMatrix(0, 0, m, minInt(m+1, n))
	Σ = Diagonal(s)
	theV = MakeDenseMatrixStacked(V)

	return
}
