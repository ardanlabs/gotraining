// Copyright ©2013 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mat

import (
	"gonum.org/v1/gonum/blas"
	"gonum.org/v1/gonum/blas/blas64"
	"gonum.org/v1/gonum/internal/asm/f64"
)

var (
	vector *VecDense

	_ Matrix  = vector
	_ Vector  = vector
	_ Reseter = vector
)

// Vector is a column vector.
type Vector interface {
	Matrix
	Len() int
}

// VecDense represents a column vector.
type VecDense struct {
	mat blas64.Vector
	n   int
	// A BLAS vector can have a negative increment, but allowing this
	// in the mat type complicates a lot of code, and doesn't gain anything.
	// VecDense must have positive increment in this package.
}

// NewVecDense creates a new VecDense of length n. If data == nil,
// a new slice is allocated for the backing slice. If len(data) == n, data is
// used as the backing slice, and changes to the elements of the returned VecDense
// will be reflected in data. If neither of these is true, NewVecDense will panic.
func NewVecDense(n int, data []float64) *VecDense {
	if len(data) != n && data != nil {
		panic(ErrShape)
	}
	if data == nil {
		data = make([]float64, n)
	}
	return &VecDense{
		mat: blas64.Vector{
			Inc:  1,
			Data: data,
		},
		n: n,
	}
}

// SliceVec returns a new VecDense that shares backing data with the receiver.
// The returned matrix starts at i of the receiver and extends k-i elements.
// SliceVec panics with ErrIndexOutOfRange if the slice is outside the capacity
// of the receiver.
func (v *VecDense) SliceVec(i, k int) *VecDense {
	if i < 0 || k <= i || v.Cap() < k {
		panic(ErrIndexOutOfRange)
	}
	return &VecDense{
		n: k - i,
		mat: blas64.Vector{
			Inc:  v.mat.Inc,
			Data: v.mat.Data[i*v.mat.Inc : (k-1)*v.mat.Inc+1],
		},
	}
}

// Dims returns the number of rows and columns in the matrix. Columns is always 1
// for a non-Reset vector.
func (v *VecDense) Dims() (r, c int) {
	if v.IsZero() {
		return 0, 0
	}
	return v.n, 1
}

// Caps returns the number of rows and columns in the backing matrix. Columns is always 1
// for a non-Reset vector.
func (v *VecDense) Caps() (r, c int) {
	if v.IsZero() {
		return 0, 0
	}
	return v.Cap(), 1
}

// Len returns the length of the vector.
func (v *VecDense) Len() int {
	return v.n
}

// Cap returns the capacity of the vector.
func (v *VecDense) Cap() int {
	if v.IsZero() {
		return 0
	}
	return (cap(v.mat.Data)-1)/v.mat.Inc + 1
}

// T performs an implicit transpose by returning the receiver inside a Transpose.
func (v *VecDense) T() Matrix {
	return Transpose{v}
}

// Reset zeros the length of the vector so that it can be reused as the
// receiver of a dimensionally restricted operation.
//
// See the Reseter interface for more information.
func (v *VecDense) Reset() {
	// No change of Inc or n to 0 may be
	// made unless both are set to 0.
	v.mat.Inc = 0
	v.n = 0
	v.mat.Data = v.mat.Data[:0]
}

// CloneVec makes a copy of a into the receiver, overwriting the previous value
// of the receiver.
func (v *VecDense) CloneVec(a *VecDense) {
	if v == a {
		return
	}
	v.n = a.n
	v.mat = blas64.Vector{
		Inc:  1,
		Data: use(v.mat.Data, v.n),
	}
	blas64.Copy(v.n, a.mat, v.mat)
}

func (v *VecDense) RawVector() blas64.Vector {
	return v.mat
}

// CopyVec makes a copy of elements of a into the receiver. It is similar to the
// built-in copy; it copies as much as the overlap between the two vectors and
// returns the number of elements it copied.
func (v *VecDense) CopyVec(a *VecDense) int {
	n := min(v.Len(), a.Len())
	if v != a {
		blas64.Copy(n, a.mat, v.mat)
	}
	return n
}

// ScaleVec scales the vector a by alpha, placing the result in the receiver.
func (v *VecDense) ScaleVec(alpha float64, a *VecDense) {
	n := a.Len()
	if v != a {
		v.reuseAs(n)
		if v.mat.Inc == 1 && a.mat.Inc == 1 {
			f64.ScalUnitaryTo(v.mat.Data, alpha, a.mat.Data)
			return
		}
		f64.ScalIncTo(v.mat.Data, uintptr(v.mat.Inc),
			alpha, a.mat.Data, uintptr(n), uintptr(a.mat.Inc))
		return
	}
	if v.mat.Inc == 1 {
		f64.ScalUnitary(alpha, v.mat.Data)
		return
	}
	f64.ScalInc(alpha, v.mat.Data, uintptr(n), uintptr(v.mat.Inc))
}

// AddScaledVec adds the vectors a and alpha*b, placing the result in the receiver.
func (v *VecDense) AddScaledVec(a *VecDense, alpha float64, b *VecDense) {
	if alpha == 1 {
		v.AddVec(a, b)
		return
	}
	if alpha == -1 {
		v.SubVec(a, b)
		return
	}

	ar := a.Len()
	br := b.Len()

	if ar != br {
		panic(ErrShape)
	}

	if v != a {
		v.checkOverlap(a.mat)
	}
	if v != b {
		v.checkOverlap(b.mat)
	}

	v.reuseAs(ar)

	switch {
	case alpha == 0: // v <- a
		v.CopyVec(a)
	case v == a && v == b: // v <- v + alpha * v = (alpha + 1) * v
		blas64.Scal(ar, alpha+1, v.mat)
	case v == a && v != b: // v <- v + alpha * b
		if v.mat.Inc == 1 && b.mat.Inc == 1 {
			// Fast path for a common case.
			f64.AxpyUnitaryTo(v.mat.Data, alpha, b.mat.Data, a.mat.Data)
		} else {
			f64.AxpyInc(alpha, b.mat.Data, v.mat.Data,
				uintptr(ar), uintptr(b.mat.Inc), uintptr(v.mat.Inc), 0, 0)
		}
	default: // v <- a + alpha * b or v <- a + alpha * v
		if v.mat.Inc == 1 && a.mat.Inc == 1 && b.mat.Inc == 1 {
			// Fast path for a common case.
			f64.AxpyUnitaryTo(v.mat.Data, alpha, b.mat.Data, a.mat.Data)
		} else {
			f64.AxpyIncTo(v.mat.Data, uintptr(v.mat.Inc), 0,
				alpha, b.mat.Data, a.mat.Data,
				uintptr(ar), uintptr(b.mat.Inc), uintptr(a.mat.Inc), 0, 0)
		}
	}
}

// AddVec adds the vectors a and b, placing the result in the receiver.
func (v *VecDense) AddVec(a, b *VecDense) {
	ar := a.Len()
	br := b.Len()

	if ar != br {
		panic(ErrShape)
	}

	if v != a {
		v.checkOverlap(a.mat)
	}
	if v != b {
		v.checkOverlap(b.mat)
	}

	v.reuseAs(ar)

	if v.mat.Inc == 1 && a.mat.Inc == 1 && b.mat.Inc == 1 {
		// Fast path for a common case.
		f64.AxpyUnitaryTo(v.mat.Data, 1, b.mat.Data, a.mat.Data)
		return
	}
	f64.AxpyIncTo(v.mat.Data, uintptr(v.mat.Inc), 0,
		1, b.mat.Data, a.mat.Data,
		uintptr(ar), uintptr(b.mat.Inc), uintptr(a.mat.Inc), 0, 0)
}

// SubVec subtracts the vector b from a, placing the result in the receiver.
func (v *VecDense) SubVec(a, b *VecDense) {
	ar := a.Len()
	br := b.Len()

	if ar != br {
		panic(ErrShape)
	}

	if v != a {
		v.checkOverlap(a.mat)
	}
	if v != b {
		v.checkOverlap(b.mat)
	}

	v.reuseAs(ar)

	if v.mat.Inc == 1 && a.mat.Inc == 1 && b.mat.Inc == 1 {
		// Fast path for a common case.
		f64.AxpyUnitaryTo(v.mat.Data, -1, b.mat.Data, a.mat.Data)
		return
	}
	f64.AxpyIncTo(v.mat.Data, uintptr(v.mat.Inc), 0,
		-1, b.mat.Data, a.mat.Data,
		uintptr(ar), uintptr(b.mat.Inc), uintptr(a.mat.Inc), 0, 0)
}

// MulElemVec performs element-wise multiplication of a and b, placing the result
// in the receiver.
func (v *VecDense) MulElemVec(a, b *VecDense) {
	ar := a.Len()
	br := b.Len()

	if ar != br {
		panic(ErrShape)
	}

	if v != a {
		v.checkOverlap(a.mat)
	}
	if v != b {
		v.checkOverlap(b.mat)
	}

	v.reuseAs(ar)

	amat, bmat := a.RawVector(), b.RawVector()
	for i := 0; i < v.n; i++ {
		v.mat.Data[i*v.mat.Inc] = amat.Data[i*amat.Inc] * bmat.Data[i*bmat.Inc]
	}
}

// DivElemVec performs element-wise division of a by b, placing the result
// in the receiver.
func (v *VecDense) DivElemVec(a, b *VecDense) {
	ar := a.Len()
	br := b.Len()

	if ar != br {
		panic(ErrShape)
	}

	if v != a {
		v.checkOverlap(a.mat)
	}
	if v != b {
		v.checkOverlap(b.mat)
	}

	v.reuseAs(ar)

	amat, bmat := a.RawVector(), b.RawVector()
	for i := 0; i < v.n; i++ {
		v.mat.Data[i*v.mat.Inc] = amat.Data[i*amat.Inc] / bmat.Data[i*bmat.Inc]
	}
}

// MulVec computes a * b. The result is stored into the receiver.
// MulVec panics if the number of columns in a does not equal the number of rows in b.
func (v *VecDense) MulVec(a Matrix, b *VecDense) {
	r, c := a.Dims()
	br := b.Len()
	if c != br {
		panic(ErrShape)
	}

	if v != b {
		v.checkOverlap(b.mat)
	}

	a, trans := untranspose(a)
	ar, ac := a.Dims()
	v.reuseAs(r)
	var restore func()
	if v == a {
		v, restore = v.isolatedWorkspace(a.(*VecDense))
		defer restore()
	} else if v == b {
		v, restore = v.isolatedWorkspace(b)
		defer restore()
	}

	switch a := a.(type) {
	case *VecDense:
		if v != a {
			v.checkOverlap(a.mat)
		}

		if a.Len() == 1 {
			// {1,1} x {1,n}
			av := a.At(0, 0)
			for i := 0; i < b.Len(); i++ {
				v.mat.Data[i*v.mat.Inc] = av * b.mat.Data[i*b.mat.Inc]
			}
			return
		}
		if b.Len() == 1 {
			// {1,n} x {1,1}
			bv := b.At(0, 0)
			for i := 0; i < a.Len(); i++ {
				v.mat.Data[i*v.mat.Inc] = bv * a.mat.Data[i*a.mat.Inc]
			}
			return
		}
		// {n,1} x {1,n}
		var sum float64
		for i := 0; i < c; i++ {
			sum += a.At(i, 0) * b.At(i, 0)
		}
		v.SetVec(0, sum)
		return
	case RawSymmetricer:
		amat := a.RawSymmetric()
		blas64.Symv(1, amat, b.mat, 0, v.mat)
	case RawTriangular:
		v.CopyVec(b)
		amat := a.RawTriangular()
		ta := blas.NoTrans
		if trans {
			ta = blas.Trans
		}
		blas64.Trmv(ta, amat, v.mat)
	case RawMatrixer:
		amat := a.RawMatrix()
		// We don't know that a is a *Dense, so make
		// a temporary Dense to check overlap.
		(&Dense{mat: amat}).checkOverlap(v.asGeneral())
		t := blas.NoTrans
		if trans {
			t = blas.Trans
		}
		blas64.Gemv(t, 1, amat, b.mat, 0, v.mat)
	default:
		if trans {
			col := make([]float64, ar)
			for c := 0; c < ac; c++ {
				for i := range col {
					col[i] = a.At(i, c)
				}
				var f float64
				for i, e := range col {
					f += e * b.mat.Data[i*b.mat.Inc]
				}
				v.mat.Data[c*v.mat.Inc] = f
			}
		} else {
			row := make([]float64, ac)
			for r := 0; r < ar; r++ {
				for i := range row {
					row[i] = a.At(r, i)
				}
				var f float64
				for i, e := range row {
					f += e * b.mat.Data[i*b.mat.Inc]
				}
				v.mat.Data[r*v.mat.Inc] = f
			}
		}
	}
}

// reuseAs resizes an empty vector to a r×1 vector,
// or checks that a non-empty matrix is r×1.
func (v *VecDense) reuseAs(r int) {
	if v.IsZero() {
		v.mat = blas64.Vector{
			Inc:  1,
			Data: use(v.mat.Data, r),
		}
		v.n = r
		return
	}
	if r != v.n {
		panic(ErrShape)
	}
}

// IsZero returns whether the receiver is zero-sized. Zero-sized vectors can be the
// receiver for size-restricted operations. VecDenses can be zeroed using Reset.
func (v *VecDense) IsZero() bool {
	// It must be the case that v.Dims() returns
	// zeros in this case. See comment in Reset().
	return v.mat.Inc == 0
}

func (v *VecDense) isolatedWorkspace(a *VecDense) (n *VecDense, restore func()) {
	l := a.Len()
	n = getWorkspaceVec(l, false)
	return n, func() {
		v.CopyVec(n)
		putWorkspaceVec(n)
	}
}

// asDense returns a Dense representation of the receiver with the same
// underlying data.
func (v *VecDense) asDense() *Dense {
	return &Dense{
		mat:     v.asGeneral(),
		capRows: v.n,
		capCols: 1,
	}
}

// asGeneral returns a blas64.General representation of the receiver with the
// same underlying data.
func (v *VecDense) asGeneral() blas64.General {
	return blas64.General{
		Rows:   v.n,
		Cols:   1,
		Stride: v.mat.Inc,
		Data:   v.mat.Data,
	}
}

// ColViewOf reflects the column j of the RawMatrixer m, into the receiver
// backed by the same underlying data. The length of the receiver must either be
// zero or match the number of rows in m.
func (v *VecDense) ColViewOf(m RawMatrixer, j int) {
	rm := m.RawMatrix()

	if j >= rm.Cols || j < 0 {
		panic(ErrColAccess)
	}
	if !v.IsZero() && v.n != rm.Rows {
		panic(ErrShape)
	}

	v.mat.Inc = rm.Stride
	v.mat.Data = rm.Data[j : (rm.Rows-1)*rm.Stride+j+1]
	v.n = rm.Rows
}

// RowViewOf reflects the row i of the RawMatrixer m, into the receiver
// backed by the same underlying data. The length of the receiver must either be
// zero or match the number of columns in m.
func (v *VecDense) RowViewOf(m RawMatrixer, i int) {
	rm := m.RawMatrix()

	if i >= rm.Rows || i < 0 {
		panic(ErrRowAccess)
	}
	if !v.IsZero() && v.n != rm.Cols {
		panic(ErrShape)
	}

	v.mat.Inc = 1
	v.mat.Data = rm.Data[i*rm.Stride : i*rm.Stride+rm.Cols]
	v.n = rm.Cols
}
