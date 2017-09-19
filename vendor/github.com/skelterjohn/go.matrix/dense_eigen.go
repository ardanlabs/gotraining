// Copyright 2009 The GoMatrix Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

import "math"

/*
Returns V,D st V*D*inv(V) = A and D is diagonal (or block diagonal).
*/
func (A *DenseMatrix) Eigen() (V, D *DenseMatrix, err error) {
	//code translated/ripped off from Jama
	if A.cols != A.rows {
		err = ErrorDimensionMismatch
		return
	}
	n := A.cols
	Va := A.Copy().Arrays()
	d := make([]float64, n)
	e := make([]float64, n)
	if A.Symmetric() {

		tred2(Va[0:n], d[0:n], e[0:n]) //pass slices so they're references

		tql2(Va[0:n], d[0:n], e[0:n])
	} else {
		H := A.GetMatrix(0, 0, n, n).Copy().Arrays()
		ort := make([]float64, n)

		// Reduce to Hessenberg form.
		orthes(Va[0:n], d[0:n], e[0:n], H[0:n], ort[0:n])

		// Reduce Hessenberg to real Schur form.
		hqr2(Va[0:n], d[0:n], e[0:n], H[0:n], ort[0:n])
	}
	V, D = MakeDenseMatrixStacked(Va), makeD(d, e)
	return
}

func makeD(d []float64, e []float64) *DenseMatrix {
	n := len(d)
	X := Zeros(n, n)
	D := X.Arrays()
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			D[i][j] = 0.0
		}
		D[i][i] = d[i]
		if e[i] > 0 {
			D[i][i+1] = e[i]
		} else if e[i] < 0 {
			D[i][i-1] = e[i]
		}
	}
	return X
}

func tred2(V [][]float64, d []float64, e []float64) {
	n := len(V)

	//  This is derived from the Algol procedures tred2 by
	//  Bowdler, Martin, Reinsch, and Wilkinson, Handbook for
	//  Auto. Comp., Vol.ii-Linear Algebra, and the corresponding
	//  Fortran subroutine in EISPACK.

	for j := 0; j < n; j++ {
		d[j] = V[n-1][j]
	}

	// Householder reduction to tridiagonal form.

	for i := n - 1; i > 0; i-- {

		// Scale to avoid under/overflow.

		scale := float64(0)
		h := float64(0)
		for k := 0; k < i; k++ {
			scale = scale + math.Abs(d[k])
		}
		if scale == 0.0 {
			e[i] = d[i-1]
			for j := 0; j < i; j++ {
				d[j] = V[i-1][j]
				V[i][j] = 0.0
				V[j][i] = 0.0
			}
		} else {

			// Generate Householder vector.

			for k := 0; k < i; k++ {
				d[k] /= scale
				h += d[k] * d[k]
			}
			f := d[i-1]
			g := math.Sqrt(h)
			if f > 0 {
				g = -g
			}
			e[i] = scale * g
			h = h - f*g
			d[i-1] = f - g
			for j := 0; j < i; j++ {
				e[j] = 0.0
			}
			// Apply similarity transformation to remaining columns.

			for j := 0; j < i; j++ {
				f = d[j]
				V[j][i] = f
				g = e[j] + V[j][j]*f
				for k := j + 1; k <= i-1; k++ {
					g += V[k][j] * d[k]
					e[k] += V[k][j] * f
				}
				e[j] = g
			}

			f = 0.0
			for j := 0; j < i; j++ {
				e[j] /= h
				f += e[j] * d[j]
			}
			hh := f / (h + h)
			for j := 0; j < i; j++ {
				e[j] -= hh * d[j]
			}

			for j := 0; j < i; j++ {
				f = d[j]
				g = e[j]
				for k := j; k <= i-1; k++ {
					V[k][j] -= (f*e[k] + g*d[k])
				}
				d[j] = V[i-1][j]
				V[i][j] = 0.0
			}
		}
		d[i] = h
	}

	// Accumulate transformations.

	for i := 0; i < n-1; i++ {
		V[n-1][i] = V[i][i]
		V[i][i] = 1.0
		h := d[i+1]
		if h != 0.0 {
			for k := 0; k <= i; k++ {
				d[k] = V[k][i+1] / h
			}
			for j := 0; j <= i; j++ {
				g := float64(0)
				for k := 0; k <= i; k++ {
					g += V[k][i+1] * V[k][j]
				}
				for k := 0; k <= i; k++ {
					V[k][j] -= g * d[k]
				}
			}
		}
		for k := 0; k <= i; k++ {
			V[k][i+1] = 0.0
		}
	}
	for j := 0; j < n; j++ {
		d[j] = V[n-1][j]
		V[n-1][j] = 0.0
	}
	V[n-1][n-1] = 1.0
	e[0] = 0.0
}

func tql2(V [][]float64, d []float64, e []float64) {

	//  This is derived from the Algol procedures tql2, by
	//  Bowdler, Martin, Reinsch, and Wilkinson, Handbook for
	//  Auto. Comp., Vol.ii-Linear Algebra, and the corresponding
	//  Fortran subroutine in EISPACK.

	n := len(V)

	for i := 1; i < n; i++ {
		e[i-1] = e[i]
	}
	e[n-1] = 0.0

	f := float64(0)
	tst1 := float64(0)
	eps := math.Pow(2.0, -52.0)
	for l := 0; l < n; l++ {

		// Find small subdiagonal element

		tst1 = max(tst1, math.Abs(d[l])+math.Abs(e[l]))
		m := l
		for m < n {
			if math.Abs(e[m]) <= eps*tst1 {
				break
			}
			m++
		}

		// If m == l, d[l] is an eigenvalue,
		// otherwise, iterate.

		if m > l {
			iter := 0
			for true {
				iter = iter + 1 // (Could check iteration count here.)

				// Compute implicit shift

				g := d[l]
				p := (d[l+1] - g) / (2.0 * e[l])
				r := math.Sqrt(p*p + 1.0)
				if p < 0 {
					r = -r
				}
				d[l] = e[l] / (p + r)
				d[l+1] = e[l] * (p + r)
				dl1 := d[l+1]
				h := g - d[l]
				for i := l + 2; i < n; i++ {
					d[i] -= h
				}
				f = f + h

				// Implicit QL transformation.

				p = d[m]
				c := float64(1)
				c2 := c
				c3 := c
				el1 := e[l+1]
				s := float64(0)
				s2 := float64(0)
				for i := m - 1; i >= l; i-- {
					c3 = c2
					c2 = c
					s2 = s
					g = c * e[i]
					h = c * p
					r = math.Sqrt(p*p + e[i]*e[i])
					e[i+1] = s * r
					s = e[i] / r
					c = p / r
					p = c*d[i] - s*g
					d[i+1] = h + s*(c*g+s*d[i])

					// Accumulate transformation.

					for k := 0; k < n; k++ {
						h = V[k][i+1]
						V[k][i+1] = s*V[k][i] + c*h
						V[k][i] = c*V[k][i] - s*h
					}
				}
				p = -s * s2 * c3 * el1 * e[l] / dl1
				e[l] = s * p
				d[l] = c * p

				// Check for convergence.
				if !(math.Abs(e[l]) > eps*tst1) {
					break
				}
			}
		}
		d[l] = d[l] + f
		e[l] = 0.0
	}

	// Sort eigenvalues and corresponding vectors.

	for i := 0; i < n-1; i++ {
		k := i
		p := d[i]
		for j := i + 1; j < n; j++ {
			if d[j] < p {
				k = j
				p = d[j]
			}
		}
		if k != i {
			d[k] = d[i]
			d[i] = p
			for j := 0; j < n; j++ {
				p = V[j][i]
				V[j][i] = V[j][k]
				V[j][k] = p
			}
		}
	}
}

func orthes(V [][]float64, d []float64, e []float64, H [][]float64, ort []float64) {

	//  This is derived from the Algol procedures orthes and ortran,
	//  by Martin and Wilkinson, Handbook for Auto. Comp.,
	//  Vol.ii-Linear Algebra, and the corresponding
	//  Fortran subroutines in EISPACK.

	n := len(V)

	low := 0
	high := n - 1

	for m := low + 1; m <= high-1; m++ {

		// Scale column.

		scale := float64(0)
		for i := m; i <= high; i++ {
			scale = scale + math.Abs(H[i][m-1])
		}
		if scale != 0.0 {

			// Compute Householder transformation.

			h := float64(0)
			for i := high; i >= m; i-- {
				ort[i] = H[i][m-1] / scale
				h += ort[i] * ort[i]
			}
			g := math.Sqrt(h)
			if ort[m] > 0 {
				g = -g
			}
			h = h - ort[m]*g
			ort[m] = ort[m] - g

			// Apply Householder similarity transformation
			// H = (I-u*u'/h)*H*(I-u*u')/h)

			for j := m; j < n; j++ {
				f := float64(0)
				for i := high; i >= m; i-- {
					f += ort[i] * H[i][j]
				}
				f = f / h
				for i := m; i <= high; i++ {
					H[i][j] -= f * ort[i]
				}
			}

			for i := 0; i <= high; i++ {
				f := float64(0)
				for j := high; j >= m; j-- {
					f += ort[j] * H[i][j]
				}
				f = f / h
				for j := m; j <= high; j++ {
					H[i][j] -= f * ort[j]
				}
			}
			ort[m] = scale * ort[m]
			H[m][m-1] = scale * g
		}
	}

	// Accumulate transformations (Algol's ortran).

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				V[i][j] = 1
			} else {
				V[i][j] = 0
			}
		}
	}

	for m := high - 1; m >= low+1; m-- {
		if H[m][m-1] != 0.0 {
			for i := m + 1; i <= high; i++ {
				ort[i] = H[i][m-1]
			}
			for j := m; j <= high; j++ {
				g := float64(0)
				for i := m; i <= high; i++ {
					g += ort[i] * V[i][j]
				}
				// Double division avoids possible underflow
				g = (g / ort[m]) / H[m][m-1]
				for i := m; i <= high; i++ {
					V[i][j] += g * ort[i]
				}
			}
		}
	}
}

func hqr2(V [][]float64, d []float64, e []float64, H [][]float64, ort []float64) {

	//  This is derived from the Algol procedure hqr2,
	//  by Martin and Wilkinson, Handbook for Auto. Comp.,
	//  Vol.ii-Linear Algebra, and the corresponding
	//  Fortran subroutine in EISPACK.

	// Initialize

	n := len(V)

	nn := n
	n = nn - 1
	low := 0
	high := nn - 1
	eps := math.Pow(2.0, -52.0)
	exshift := float64(0)
	p := float64(0)
	q := float64(0)
	r := float64(0)
	s := float64(0)
	z := float64(0)
	var t, w, x, y float64

	// Store roots isolated by balanc and compute matrix norm

	norm := float64(0)
	for i := 0; i < nn; i++ {
		if i < low || i > high {
			d[i] = H[i][i]
			e[i] = 0.0
		}
		for j := int(max(float64(i)-1, 0)); j < nn; j++ {
			norm = norm + math.Abs(H[i][j])
		}
	}

	// Outer loop over eigenvalue index

	iter := 0
	for n >= low {

		// Look for single small sub-diagonal element

		l := n
		for l > low {
			s = math.Abs(H[l-1][l-1]) + math.Abs(H[l][l])
			if s == 0.0 {
				s = norm
			}
			if math.Abs(H[l][l-1]) < eps*s {
				break
			}
			l--
		}

		// Check for convergence
		// One root found

		if l == n {
			H[n][n] = H[n][n] + exshift
			d[n] = H[n][n]
			e[n] = 0.0
			n--
			iter = 0

			// Two roots found

		} else if l == n-1 {
			w = H[n][n-1] * H[n-1][n]
			p = (H[n-1][n-1] - H[n][n]) / 2.0
			q = p*p + w
			z = math.Sqrt(math.Abs(q))
			H[n][n] = H[n][n] + exshift
			H[n-1][n-1] = H[n-1][n-1] + exshift
			x = H[n][n]

			// Real pair

			if q >= 0 {
				if p >= 0 {
					z = p + z
				} else {
					z = p - z
				}
				d[n-1] = x + z
				d[n] = d[n-1]
				if z != 0.0 {
					d[n] = x - w/z
				}
				e[n-1] = 0.0
				e[n] = 0.0
				x = H[n][n-1]
				s = math.Abs(x) + math.Abs(z)
				p = x / s
				q = z / s
				r = math.Sqrt(p*p + q*q)
				p = p / r
				q = q / r

				// Row modification

				for j := n - 1; j < nn; j++ {
					z = H[n-1][j]
					H[n-1][j] = q*z + p*H[n][j]
					H[n][j] = q*H[n][j] - p*z
				}

				// Column modification

				for i := 0; i <= n; i++ {
					z = H[i][n-1]
					H[i][n-1] = q*z + p*H[i][n]
					H[i][n] = q*H[i][n] - p*z
				}

				// Accumulate transformations

				for i := low; i <= high; i++ {
					z = V[i][n-1]
					V[i][n-1] = q*z + p*V[i][n]
					V[i][n] = q*V[i][n] - p*z
				}

				// Complex pair

			} else {
				d[n-1] = x + p
				d[n] = x + p
				e[n-1] = z
				e[n] = -z
			}
			n = n - 2
			iter = 0

			// No convergence yet

		} else {

			// Form shift

			x = H[n][n]
			y = 0.0
			w = 0.0
			if l < n {
				y = H[n-1][n-1]
				w = H[n][n-1] * H[n-1][n]
			}

			// Wilkinson's original ad hoc shift

			if iter == 10 {
				exshift += x
				for i := low; i <= n; i++ {
					H[i][i] -= x
				}
				s = math.Abs(H[n][n-1]) + math.Abs(H[n-1][n-2])
				y = 0.75 * s
				x = y
				w = -0.4375 * s * s
			}

			// MATLAB's new ad hoc shift

			if iter == 30 {
				s = (y - x) / 2.0
				s = s*s + w
				if s > 0 {
					s = math.Sqrt(s)
					if y < x {
						s = -s
					}
					s = x - w/((y-x)/2.0+s)
					for i := low; i <= n; i++ {
						H[i][i] -= s
					}
					exshift += s
					w = 0.964
					y = w
					x = y
				}
			}

			iter = iter + 1 // (Could check iteration count here.)

			// Look for two consecutive small sub-diagonal elements

			m := n - 2
			for m >= l {
				z = H[m][m]
				r = x - z
				s = y - z
				p = (r*s-w)/H[m+1][m] + H[m][m+1]
				q = H[m+1][m+1] - z - r - s
				r = H[m+2][m+1]
				s = math.Abs(p) + math.Abs(q) + math.Abs(r)
				p = p / s
				q = q / s
				r = r / s
				if m == l {
					break
				}
				if math.Abs(H[m][m-1])*(math.Abs(q)+math.Abs(r)) <
					eps*(math.Abs(p)*(math.Abs(H[m-1][m-1])+math.Abs(z)+
						math.Abs(H[m+1][m+1]))) {
					break
				}
				m--
			}

			for i := m + 2; i <= n; i++ {
				H[i][i-2] = 0.0
				if i > m+2 {
					H[i][i-3] = 0.0
				}
			}

			// Double QR step involving rows l:n and columns m:n

			for k := m; k <= n-1; k++ {
				notlast := (k != n-1)
				if k != m {
					p = H[k][k-1]
					q = H[k+1][k-1]
					if notlast {
						r = H[k+2][k-1]
					} else {
						r = 0
					}

					x = math.Abs(p) + math.Abs(q) + math.Abs(r)
					if x != 0.0 {
						p = p / x
						q = q / x
						r = r / x
					}
				}
				if x == 0.0 {
					break
				}
				s = math.Sqrt(p*p + q*q + r*r)
				if p < 0 {
					s = -s
				}
				if s != 0 {
					if k != m {
						H[k][k-1] = -s * x
					} else if l != m {
						H[k][k-1] = -H[k][k-1]
					}
					p = p + s
					x = p / s
					y = q / s
					z = r / s
					q = q / p
					r = r / p

					// Row modification

					for j := k; j < nn; j++ {
						p = H[k][j] + q*H[k+1][j]
						if notlast {
							p = p + r*H[k+2][j]
							H[k+2][j] = H[k+2][j] - p*z
						}
						H[k][j] = H[k][j] - p*x
						H[k+1][j] = H[k+1][j] - p*y
					}

					// Column modification

					for i := 0; i <= int(min(float64(n), float64(k)+3)); i++ {
						p = x*H[i][k] + y*H[i][k+1]
						if notlast {
							p = p + z*H[i][k+2]
							H[i][k+2] = H[i][k+2] - p*r
						}
						H[i][k] = H[i][k] - p
						H[i][k+1] = H[i][k+1] - p*q
					}

					// Accumulate transformations

					for i := low; i <= high; i++ {
						p = x*V[i][k] + y*V[i][k+1]
						if notlast {
							p = p + z*V[i][k+2]
							V[i][k+2] = V[i][k+2] - p*r
						}
						V[i][k] = V[i][k] - p
						V[i][k+1] = V[i][k+1] - p*q
					}
				} // (s != 0)
			} // k loop
		} // check convergence
	} // while (n >= low)

	// Backsubstitute to find vectors of upper triangular form

	if norm == 0.0 {
		return
	}

	for n = nn - 1; n >= 0; n-- {
		p = d[n]
		q = e[n]

		// Real vector

		if q == 0 {
			l := n
			H[n][n] = 1.0
			for i := n - 1; i >= 0; i-- {
				w = H[i][i] - p
				r = 0.0
				for j := l; j <= n; j++ {
					r = r + H[i][j]*H[j][n]
				}
				if e[i] < 0.0 {
					z = w
					s = r
				} else {
					l = i
					if e[i] == 0.0 {
						if w != 0.0 {
							H[i][n] = -r / w
						} else {
							H[i][n] = -r / (eps * norm)
						}

						// Solve real equations

					} else {
						x = H[i][i+1]
						y = H[i+1][i]
						q = (d[i]-p)*(d[i]-p) + e[i]*e[i]
						t = (x*s - z*r) / q
						H[i][n] = t
						if math.Abs(x) > math.Abs(z) {
							H[i+1][n] = (-r - w*t) / x
						} else {
							H[i+1][n] = (-s - y*t) / z
						}
					}

					// Overflow control

					t = math.Abs(H[i][n])
					if (eps*t)*t > 1 {
						for j := i; j <= n; j++ {
							H[j][n] = H[j][n] / t
						}
					}
				}
			}

			// Complex vector

		} else if q < 0 {
			l := n - 1

			// Last vector component imaginary so matrix is triangular

			if math.Abs(H[n][n-1]) > math.Abs(H[n-1][n]) {
				H[n-1][n-1] = q / H[n][n-1]
				H[n-1][n] = -(H[n][n] - p) / H[n][n-1]
			} else {
				cdivr, cdivi := cdiv(0.0, -H[n-1][n], H[n-1][n-1]-p, q)
				H[n-1][n-1] = cdivr
				H[n-1][n] = cdivi
			}
			H[n][n-1] = 0.0
			H[n][n] = 1.0
			for i := n - 2; i >= 0; i-- {
				var ra, sa, vr, vi float64
				ra = 0.0
				sa = 0.0
				for j := l; j <= n; j++ {
					ra = ra + H[i][j]*H[j][n-1]
					sa = sa + H[i][j]*H[j][n]
				}
				w = H[i][i] - p

				if e[i] < 0.0 {
					z = w
					r = ra
					s = sa
				} else {
					l = i
					if e[i] == 0 {
						cdivr, cdivi := cdiv(-ra, -sa, w, q)
						H[i][n-1] = cdivr
						H[i][n] = cdivi
					} else {

						// Solve complex equations

						x = H[i][i+1]
						y = H[i+1][i]
						vr = (d[i]-p)*(d[i]-p) + e[i]*e[i] - q*q
						vi = (d[i] - p) * 2.0 * q
						if vr == 0.0 && vi == 0.0 {
							vr = eps * norm * (math.Abs(w) + math.Abs(q) +
								math.Abs(x) + math.Abs(y) + math.Abs(z))
						}
						cdivr, cdivi := cdiv(x*r-z*ra+q*sa, x*s-z*sa-q*ra, vr, vi)
						H[i][n-1] = cdivr
						H[i][n] = cdivi
						if math.Abs(x) > (math.Abs(z) + math.Abs(q)) {
							H[i+1][n-1] = (-ra - w*H[i][n-1] + q*H[i][n]) / x
							H[i+1][n] = (-sa - w*H[i][n] - q*H[i][n-1]) / x
						} else {
							cdiv(-r-y*H[i][n-1], -s-y*H[i][n], z, q)
							H[i+1][n-1] = cdivr
							H[i+1][n] = cdivi
						}
					}

					// Overflow control

					t = max(math.Abs(H[i][n-1]), math.Abs(H[i][n]))
					if (eps*t)*t > 1 {
						for j := i; j <= n; j++ {
							H[j][n-1] = H[j][n-1] / t
							H[j][n] = H[j][n] / t
						}
					}
				}
			}
		}
	}

	// Vectors of isolated roots

	for i := 0; i < nn; i++ {
		if i < low || i > high {
			for j := i; j < nn; j++ {
				V[i][j] = H[i][j]
			}
		}
	}

	// Back transformation to get eigenvectors of original matrix

	for j := nn - 1; j >= low; j-- {
		for i := low; i <= high; i++ {
			z = 0.0
			for k := low; k <= int(min(float64(j), float64(high))); k++ {
				z = z + V[i][k]*H[k][j]
			}
			V[i][j] = z
		}
	}
}

func cdiv(xr float64, xi float64, yr float64, yi float64) (cdivr float64, cdivi float64) {
	var r, d float64
	if math.Abs(yr) > math.Abs(yi) {
		r = yi / yr
		d = yr + r*yi
		cdivr = (xr + r*xi) / d
		cdivi = (xi - r*xr) / d
	} else {
		r = yr / yi
		d = yi + r*yr
		cdivr = (r*xr + xi) / d
		cdivi = (r*xi - xr) / d
	}
	return
}
