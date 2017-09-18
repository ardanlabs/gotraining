// Copyright ©2013 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mat64

import (
	"math"

	"github.com/gonum/blas"
	"github.com/gonum/blas/blas64"
	"github.com/gonum/lapack/lapack64"
	"github.com/gonum/matrix"
)

// Add adds a and b element-wise, placing the result in the receiver. Add
// will panic if the two matrices do not have the same shape.
func (m *Dense) Add(a, b Matrix) {
	ar, ac := a.Dims()
	br, bc := b.Dims()
	if ar != br || ac != bc {
		panic(matrix.ErrShape)
	}

	aU, _ := untranspose(a)
	bU, _ := untranspose(b)
	m.reuseAs(ar, ac)

	if arm, ok := a.(RawMatrixer); ok {
		if brm, ok := b.(RawMatrixer); ok {
			amat, bmat := arm.RawMatrix(), brm.RawMatrix()
			if m != aU {
				m.checkOverlap(amat)
			}
			if m != bU {
				m.checkOverlap(bmat)
			}
			for ja, jb, jm := 0, 0, 0; ja < ar*amat.Stride; ja, jb, jm = ja+amat.Stride, jb+bmat.Stride, jm+m.mat.Stride {
				for i, v := range amat.Data[ja : ja+ac] {
					m.mat.Data[i+jm] = v + bmat.Data[i+jb]
				}
			}
			return
		}
	}

	var restore func()
	if m == aU {
		m, restore = m.isolatedWorkspace(aU)
		defer restore()
	} else if m == bU {
		m, restore = m.isolatedWorkspace(bU)
		defer restore()
	}

	for r := 0; r < ar; r++ {
		for c := 0; c < ac; c++ {
			m.set(r, c, a.At(r, c)+b.At(r, c))
		}
	}
}

// Sub subtracts the matrix b from a, placing the result in the receiver. Sub
// will panic if the two matrices do not have the same shape.
func (m *Dense) Sub(a, b Matrix) {
	ar, ac := a.Dims()
	br, bc := b.Dims()
	if ar != br || ac != bc {
		panic(matrix.ErrShape)
	}

	aU, _ := untranspose(a)
	bU, _ := untranspose(b)
	m.reuseAs(ar, ac)

	if arm, ok := a.(RawMatrixer); ok {
		if brm, ok := b.(RawMatrixer); ok {
			amat, bmat := arm.RawMatrix(), brm.RawMatrix()
			if m != aU {
				m.checkOverlap(amat)
			}
			if m != bU {
				m.checkOverlap(bmat)
			}
			for ja, jb, jm := 0, 0, 0; ja < ar*amat.Stride; ja, jb, jm = ja+amat.Stride, jb+bmat.Stride, jm+m.mat.Stride {
				for i, v := range amat.Data[ja : ja+ac] {
					m.mat.Data[i+jm] = v - bmat.Data[i+jb]
				}
			}
			return
		}
	}

	var restore func()
	if m == aU {
		m, restore = m.isolatedWorkspace(aU)
		defer restore()
	} else if m == bU {
		m, restore = m.isolatedWorkspace(bU)
		defer restore()
	}

	for r := 0; r < ar; r++ {
		for c := 0; c < ac; c++ {
			m.set(r, c, a.At(r, c)-b.At(r, c))
		}
	}
}

// MulElem performs element-wise multiplication of a and b, placing the result
// in the receiver. MulElem will panic if the two matrices do not have the same
// shape.
func (m *Dense) MulElem(a, b Matrix) {
	ar, ac := a.Dims()
	br, bc := b.Dims()
	if ar != br || ac != bc {
		panic(matrix.ErrShape)
	}

	aU, _ := untranspose(a)
	bU, _ := untranspose(b)
	m.reuseAs(ar, ac)

	if arm, ok := a.(RawMatrixer); ok {
		if brm, ok := b.(RawMatrixer); ok {
			amat, bmat := arm.RawMatrix(), brm.RawMatrix()
			if m != aU {
				m.checkOverlap(amat)
			}
			if m != bU {
				m.checkOverlap(bmat)
			}
			for ja, jb, jm := 0, 0, 0; ja < ar*amat.Stride; ja, jb, jm = ja+amat.Stride, jb+bmat.Stride, jm+m.mat.Stride {
				for i, v := range amat.Data[ja : ja+ac] {
					m.mat.Data[i+jm] = v * bmat.Data[i+jb]
				}
			}
			return
		}
	}

	var restore func()
	if m == aU {
		m, restore = m.isolatedWorkspace(aU)
		defer restore()
	} else if m == bU {
		m, restore = m.isolatedWorkspace(bU)
		defer restore()
	}

	for r := 0; r < ar; r++ {
		for c := 0; c < ac; c++ {
			m.set(r, c, a.At(r, c)*b.At(r, c))
		}
	}
}

// DivElem performs element-wise division of a by b, placing the result
// in the receiver. DivElem will panic if the two matrices do not have the same
// shape.
func (m *Dense) DivElem(a, b Matrix) {
	ar, ac := a.Dims()
	br, bc := b.Dims()
	if ar != br || ac != bc {
		panic(matrix.ErrShape)
	}

	aU, _ := untranspose(a)
	bU, _ := untranspose(b)
	m.reuseAs(ar, ac)

	if arm, ok := a.(RawMatrixer); ok {
		if brm, ok := b.(RawMatrixer); ok {
			amat, bmat := arm.RawMatrix(), brm.RawMatrix()
			if m != aU {
				m.checkOverlap(amat)
			}
			if m != bU {
				m.checkOverlap(bmat)
			}
			for ja, jb, jm := 0, 0, 0; ja < ar*amat.Stride; ja, jb, jm = ja+amat.Stride, jb+bmat.Stride, jm+m.mat.Stride {
				for i, v := range amat.Data[ja : ja+ac] {
					m.mat.Data[i+jm] = v / bmat.Data[i+jb]
				}
			}
			return
		}
	}

	var restore func()
	if m == aU {
		m, restore = m.isolatedWorkspace(aU)
		defer restore()
	} else if m == bU {
		m, restore = m.isolatedWorkspace(bU)
		defer restore()
	}

	for r := 0; r < ar; r++ {
		for c := 0; c < ac; c++ {
			m.set(r, c, a.At(r, c)/b.At(r, c))
		}
	}
}

// Inverse computes the inverse of the matrix a, storing the result into the
// receiver. If a is ill-conditioned, a Condition error will be returned.
// Note that matrix inversion is numerically unstable, and should generally
// be avoided where possible, for example by using the Solve routines.
func (m *Dense) Inverse(a Matrix) error {
	// TODO(btracey): Special case for RawTriangular, etc.
	r, c := a.Dims()
	if r != c {
		panic(matrix.ErrSquare)
	}
	m.reuseAs(a.Dims())
	aU, aTrans := untranspose(a)
	switch rm := aU.(type) {
	case RawMatrixer:
		if m != aU || aTrans {
			if m == aU || m.checkOverlap(rm.RawMatrix()) {
				tmp := getWorkspace(r, c, false)
				tmp.Copy(a)
				m.Copy(tmp)
				putWorkspace(tmp)
				break
			}
			m.Copy(a)
		}
	default:
		m.Copy(a)
	}
	ipiv := make([]int, r)
	ok := lapack64.Getrf(m.mat, ipiv)
	if !ok {
		return matrix.Condition(math.Inf(1))
	}
	work := make([]float64, 1, 4*r) // must be at least 4*r for cond.
	lapack64.Getri(m.mat, ipiv, work, -1)
	if int(work[0]) > 4*r {
		work = make([]float64, int(work[0]))
	} else {
		work = work[:4*r]
	}
	lapack64.Getri(m.mat, ipiv, work, len(work))
	norm := lapack64.Lange(matrix.CondNorm, m.mat, work)
	rcond := lapack64.Gecon(matrix.CondNorm, m.mat, norm, work, ipiv) // reuse ipiv
	if rcond == 0 {
		return matrix.Condition(math.Inf(1))
	}
	cond := 1 / rcond
	if cond > matrix.ConditionTolerance {
		return matrix.Condition(cond)
	}
	return nil
}

// Mul takes the matrix product of a and b, placing the result in the receiver.
// If the number of columns in a does not equal the number of rows in b, Mul will panic.
func (m *Dense) Mul(a, b Matrix) {
	ar, ac := a.Dims()
	br, bc := b.Dims()

	if ac != br {
		panic(matrix.ErrShape)
	}

	aU, aTrans := untranspose(a)
	bU, bTrans := untranspose(b)
	m.reuseAs(ar, bc)
	var restore func()
	if m == aU {
		m, restore = m.isolatedWorkspace(aU)
		defer restore()
	} else if m == bU {
		m, restore = m.isolatedWorkspace(bU)
		defer restore()
	}
	aT := blas.NoTrans
	if aTrans {
		aT = blas.Trans
	}
	bT := blas.NoTrans
	if bTrans {
		bT = blas.Trans
	}

	// Some of the cases do not have a transpose option, so create
	// temporary memory.
	// C = A^T * B = (B^T * A)^T
	// C^T = B^T * A.
	if aUrm, ok := aU.(RawMatrixer); ok {
		amat := aUrm.RawMatrix()
		if restore == nil {
			m.checkOverlap(amat)
		}
		if bUrm, ok := bU.(RawMatrixer); ok {
			bmat := bUrm.RawMatrix()
			if restore == nil {
				m.checkOverlap(bmat)
			}
			blas64.Gemm(aT, bT, 1, amat, bmat, 0, m.mat)
			return
		}
		if bU, ok := bU.(RawSymmetricer); ok {
			bmat := bU.RawSymmetric()
			if aTrans {
				c := getWorkspace(ac, ar, false)
				blas64.Symm(blas.Left, 1, bmat, amat, 0, c.mat)
				strictCopy(m, c.T())
				putWorkspace(c)
				return
			}
			blas64.Symm(blas.Right, 1, bmat, amat, 0, m.mat)
			return
		}
		if bU, ok := bU.(RawTriangular); ok {
			// Trmm updates in place, so copy aU first.
			bmat := bU.RawTriangular()
			if aTrans {
				c := getWorkspace(ac, ar, false)
				var tmp Dense
				tmp.SetRawMatrix(amat)
				c.Copy(&tmp)
				bT := blas.Trans
				if bTrans {
					bT = blas.NoTrans
				}
				blas64.Trmm(blas.Left, bT, 1, bmat, c.mat)
				strictCopy(m, c.T())
				putWorkspace(c)
				return
			}
			m.Copy(a)
			blas64.Trmm(blas.Right, bT, 1, bmat, m.mat)
			return
		}
		if bU, ok := bU.(*Vector); ok {
			m.checkOverlap(bU.asGeneral())
			bvec := bU.RawVector()
			if bTrans {
				// {ar,1} x {1,bc}, which is not a vector.
				// Instead, construct B as a General.
				bmat := blas64.General{
					Rows:   bc,
					Cols:   1,
					Stride: bvec.Inc,
					Data:   bvec.Data,
				}
				blas64.Gemm(aT, bT, 1, amat, bmat, 0, m.mat)
				return
			}
			cvec := blas64.Vector{
				Inc:  m.mat.Stride,
				Data: m.mat.Data,
			}
			blas64.Gemv(aT, 1, amat, bvec, 0, cvec)
			return
		}
	}
	if bUrm, ok := bU.(RawMatrixer); ok {
		bmat := bUrm.RawMatrix()
		if restore == nil {
			m.checkOverlap(bmat)
		}
		if aU, ok := aU.(RawSymmetricer); ok {
			amat := aU.RawSymmetric()
			if bTrans {
				c := getWorkspace(bc, br, false)
				blas64.Symm(blas.Right, 1, amat, bmat, 0, c.mat)
				strictCopy(m, c.T())
				putWorkspace(c)
				return
			}
			blas64.Symm(blas.Left, 1, amat, bmat, 0, m.mat)
			return
		}
		if aU, ok := aU.(RawTriangular); ok {
			// Trmm updates in place, so copy bU first.
			amat := aU.RawTriangular()
			if bTrans {
				c := getWorkspace(bc, br, false)
				var tmp Dense
				tmp.SetRawMatrix(bmat)
				c.Copy(&tmp)
				aT := blas.Trans
				if aTrans {
					aT = blas.NoTrans
				}
				blas64.Trmm(blas.Right, aT, 1, amat, c.mat)
				strictCopy(m, c.T())
				putWorkspace(c)
				return
			}
			m.Copy(b)
			blas64.Trmm(blas.Left, aT, 1, amat, m.mat)
			return
		}
		if aU, ok := aU.(*Vector); ok {
			m.checkOverlap(aU.asGeneral())
			avec := aU.RawVector()
			if aTrans {
				// {1,ac} x {ac, bc}
				// Transpose B so that the vector is on the right.
				cvec := blas64.Vector{
					Inc:  1,
					Data: m.mat.Data,
				}
				bT := blas.Trans
				if bTrans {
					bT = blas.NoTrans
				}
				blas64.Gemv(bT, 1, bmat, avec, 0, cvec)
				return
			}
			// {ar,1} x {1,bc} which is not a vector result.
			// Instead, construct A as a General.
			amat := blas64.General{
				Rows:   ar,
				Cols:   1,
				Stride: avec.Inc,
				Data:   avec.Data,
			}
			blas64.Gemm(aT, bT, 1, amat, bmat, 0, m.mat)
			return
		}
	}

	row := make([]float64, ac)
	for r := 0; r < ar; r++ {
		for i := range row {
			row[i] = a.At(r, i)
		}
		for c := 0; c < bc; c++ {
			var v float64
			for i, e := range row {
				v += e * b.At(i, c)
			}
			m.mat.Data[r*m.mat.Stride+c] = v
		}
	}
}

// strictCopy copies a into m panicking if the shape of a and m differ.
func strictCopy(m *Dense, a Matrix) {
	r, c := m.Copy(a)
	if r != m.mat.Rows || c != m.mat.Cols {
		// Panic with a string since this
		// is not a user-facing panic.
		panic(matrix.ErrShape.Error())
	}
}

// Exp calculates the exponential of the matrix a, e^a, placing the result
// in the receiver. Exp will panic with matrix.ErrShape if a is not square.
//
// Exp uses the scaling and squaring method described in section 3 of
// http://www.cs.cornell.edu/cv/researchpdf/19ways+.pdf.
func (m *Dense) Exp(a Matrix) {
	r, c := a.Dims()
	if r != c {
		panic(matrix.ErrShape)
	}

	var w *Dense
	if m.isZero() {
		m.reuseAsZeroed(r, r)
		w = m
	} else {
		w = getWorkspace(r, r, true)
	}
	for i := 0; i < r*r; i += r + 1 {
		w.mat.Data[i] = 1
	}

	const (
		terms   = 10
		scaling = 4
	)

	small := getWorkspace(r, r, false)
	small.Scale(math.Pow(2, -scaling), a)
	power := getWorkspace(r, r, false)
	power.Copy(small)

	var (
		tmp   = getWorkspace(r, r, false)
		factI = 1.
	)
	for i := 1.; i < terms; i++ {
		factI *= i

		// This is OK to do because power and tmp are
		// new Dense values so all rows are contiguous.
		// TODO(kortschak) Make this explicit in the NewDense doc comment.
		for j, v := range power.mat.Data {
			tmp.mat.Data[j] = v / factI
		}

		w.Add(w, tmp)
		if i < terms-1 {
			tmp.Mul(power, small)
			tmp, power = power, tmp
		}
	}
	putWorkspace(small)
	putWorkspace(power)
	for i := 0; i < scaling; i++ {
		tmp.Mul(w, w)
		tmp, w = w, tmp
	}
	putWorkspace(tmp)

	if w != m {
		m.Copy(w)
		putWorkspace(w)
	}
}

// Pow calculates the integral power of the matrix a to n, placing the result
// in the receiver. Pow will panic if n is negative or if a is not square.
func (m *Dense) Pow(a Matrix, n int) {
	if n < 0 {
		panic("matrix: illegal power")
	}
	r, c := a.Dims()
	if r != c {
		panic(matrix.ErrShape)
	}

	m.reuseAs(r, c)

	// Take possible fast paths.
	switch n {
	case 0:
		for i := 0; i < r; i++ {
			zero(m.mat.Data[i*m.mat.Stride : i*m.mat.Stride+c])
			m.mat.Data[i*m.mat.Stride+i] = 1
		}
		return
	case 1:
		m.Copy(a)
		return
	case 2:
		m.Mul(a, a)
		return
	}

	// Perform iterative exponentiation by squaring in work space.
	w := getWorkspace(r, r, false)
	w.Copy(a)
	s := getWorkspace(r, r, false)
	s.Copy(a)
	x := getWorkspace(r, r, false)
	for n--; n > 0; n >>= 1 {
		if n&1 != 0 {
			x.Mul(w, s)
			w, x = x, w
		}
		if n != 1 {
			x.Mul(s, s)
			s, x = x, s
		}
	}
	m.Copy(w)
	putWorkspace(w)
	putWorkspace(s)
	putWorkspace(x)
}

// Scale multiplies the elements of a by f, placing the result in the receiver.
//
// See the Scaler interface for more information.
func (m *Dense) Scale(f float64, a Matrix) {
	ar, ac := a.Dims()

	m.reuseAs(ar, ac)

	aU, aTrans := untranspose(a)
	if rm, ok := aU.(RawMatrixer); ok {
		amat := rm.RawMatrix()
		if m == aU || m.checkOverlap(amat) {
			var restore func()
			m, restore = m.isolatedWorkspace(a)
			defer restore()
		}
		if !aTrans {
			for ja, jm := 0, 0; ja < ar*amat.Stride; ja, jm = ja+amat.Stride, jm+m.mat.Stride {
				for i, v := range amat.Data[ja : ja+ac] {
					m.mat.Data[i+jm] = v * f
				}
			}
		} else {
			for ja, jm := 0, 0; ja < ac*amat.Stride; ja, jm = ja+amat.Stride, jm+1 {
				for i, v := range amat.Data[ja : ja+ar] {
					m.mat.Data[i*m.mat.Stride+jm] = v * f
				}
			}
		}
		return
	}

	for r := 0; r < ar; r++ {
		for c := 0; c < ac; c++ {
			m.set(r, c, f*a.At(r, c))
		}
	}
}

// Apply applies the function fn to each of the elements of a, placing the
// resulting matrix in the receiver. The function fn takes a row/column
// index and element value and returns some function of that tuple.
func (m *Dense) Apply(fn func(i, j int, v float64) float64, a Matrix) {
	ar, ac := a.Dims()

	m.reuseAs(ar, ac)

	aU, aTrans := untranspose(a)
	if rm, ok := aU.(RawMatrixer); ok {
		amat := rm.RawMatrix()
		if m == aU || m.checkOverlap(amat) {
			var restore func()
			m, restore = m.isolatedWorkspace(a)
			defer restore()
		}
		if !aTrans {
			for j, ja, jm := 0, 0, 0; ja < ar*amat.Stride; j, ja, jm = j+1, ja+amat.Stride, jm+m.mat.Stride {
				for i, v := range amat.Data[ja : ja+ac] {
					m.mat.Data[i+jm] = fn(j, i, v)
				}
			}
		} else {
			for j, ja, jm := 0, 0, 0; ja < ac*amat.Stride; j, ja, jm = j+1, ja+amat.Stride, jm+1 {
				for i, v := range amat.Data[ja : ja+ar] {
					m.mat.Data[i*m.mat.Stride+jm] = fn(i, j, v)
				}
			}
		}
		return
	}

	for r := 0; r < ar; r++ {
		for c := 0; c < ac; c++ {
			m.set(r, c, fn(r, c, a.At(r, c)))
		}
	}
}

// RankOne performs a rank-one update to the matrix a and stores the result
// in the receiver. If a is zero, see Outer.
//  m = a + alpha * x * y'
func (m *Dense) RankOne(a Matrix, alpha float64, x, y *Vector) {
	ar, ac := a.Dims()
	if x.Len() != ar {
		panic(matrix.ErrShape)
	}
	if y.Len() != ac {
		panic(matrix.ErrShape)
	}

	m.checkOverlap(x.asGeneral())
	m.checkOverlap(y.asGeneral())

	var w Dense
	if m == a {
		w = *m
	}
	w.reuseAs(ar, ac)

	// Copy over to the new memory if necessary
	if m != a {
		w.Copy(a)
	}
	blas64.Ger(alpha, x.mat, y.mat, w.mat)
	*m = w
}

// Outer calculates the outer product of x and y, and stores the result
// in the receiver.
//  m = alpha * x * y'
// In order to update an existing matrix, see RankOne.
func (m *Dense) Outer(alpha float64, x, y *Vector) {
	r := x.Len()
	c := y.Len()

	// Copied from reuseAs with use replaced by useZeroed
	// and a final zero of the matrix elements if we pass
	// the shape checks.
	// TODO(kortschak): Factor out into reuseZeroedAs if
	// we find another case that needs it.
	if m.mat.Rows > m.capRows || m.mat.Cols > m.capCols {
		// Panic as a string, not a mat64.Error.
		panic("mat64: caps not correctly set")
	}
	if m.isZero() {
		m.mat = blas64.General{
			Rows:   r,
			Cols:   c,
			Stride: c,
			Data:   useZeroed(m.mat.Data, r*c),
		}
		m.capRows = r
		m.capCols = c
	} else if r != m.mat.Rows || c != m.mat.Cols {
		panic(matrix.ErrShape)
	} else {
		m.checkOverlap(x.asGeneral())
		m.checkOverlap(y.asGeneral())
		for i := 0; i < r; i++ {
			zero(m.mat.Data[i*m.mat.Stride : i*m.mat.Stride+c])
		}
	}

	blas64.Ger(alpha, x.mat, y.mat, m.mat)
}
