// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mat64

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"testing"

	"github.com/gonum/blas"
	"github.com/gonum/blas/blas64"
	"github.com/gonum/floats"
	"github.com/gonum/matrix"
)

// legalSizeSameRectangular returns whether the two matrices have the same rectangular shape.
func legalSizeSameRectangular(ar, ac, br, bc int) bool {
	if ar != br {
		return false
	}
	if ac != bc {
		return false
	}
	return true
}

// legalSizeSameSquare returns whether the two matrices have the same square shape.
func legalSizeSameSquare(ar, ac, br, bc int) bool {
	if ar != br {
		return false
	}
	if ac != bc {
		return false
	}
	if ar != ac {
		return false
	}
	return true
}

// legalSizeSameHeight returns whether the two matrices have the same number of rows.
func legalSizeSameHeight(ar, _, br, _ int) bool {
	return ar == br
}

// legalSizeSameWidth returns whether the two matrices have the same number of columns.
func legalSizeSameWidth(_, ac, _, bc int) bool {
	return ac == bc
}

// legalSizeSolve returns whether the two matrices can be used in a linear solve.
func legalSizeSolve(ar, ac, br, bc int) bool {
	return ar == br
}

// legalSizeSameVec returns whether the two matrices are column vectors of the
// same dimension.
func legalSizeSameVec(ar, ac, br, bc int) bool {
	return ac == 1 && bc == 1 && ar == br
}

// isAnySize returns true for all matrix sizes.
func isAnySize(ar, ac int) bool {
	return true
}

// isAnySize2 returns true for all matrix sizes.
func isAnySize2(ar, ac, br, bc int) bool {
	return true
}

// isAnyVector returns true for any column vector sizes.
func isAnyVector(ar, ac int) bool {
	return ac == 1
}

// isSquare returns whether the input matrix is square.
func isSquare(r, c int) bool {
	return r == c
}

// sameAnswerFloat returns whether the two inputs are both NaN or are equal.
func sameAnswerFloat(a, b interface{}) bool {
	if math.IsNaN(a.(float64)) {
		return math.IsNaN(b.(float64))
	}
	return a.(float64) == b.(float64)
}

// sameAnswerFloatApproxTol returns a function that determines whether its two
// inputs are both NaN or within tol of each other.
func sameAnswerFloatApproxTol(tol float64) func(a, b interface{}) bool {
	return func(a, b interface{}) bool {
		if math.IsNaN(a.(float64)) {
			return math.IsNaN(b.(float64))
		}
		return floats.EqualWithinAbsOrRel(a.(float64), b.(float64), tol, tol)
	}
}

func sameAnswerF64SliceOfSlice(a, b interface{}) bool {
	for i, v := range a.([][]float64) {
		if same := floats.Same(v, b.([][]float64)[i]); !same {
			return false
		}
	}
	return true
}

// sameAnswerBool returns whether the two inputs have the same value.
func sameAnswerBool(a, b interface{}) bool {
	return a.(bool) == b.(bool)
}

// isAnyType returns true for all Matrix types.
func isAnyType(Matrix) bool {
	return true
}

// legalTypesAll returns true for all Matrix types.
func legalTypesAll(a, b Matrix) bool {
	return true
}

// legalTypeSym returns whether a is a Symmetric.
func legalTypeSym(a Matrix) bool {
	_, ok := a.(Symmetric)
	return ok
}

// legalTypesSym returns whether both input arguments are Symmetric.
func legalTypesSym(a, b Matrix) bool {
	if _, ok := a.(Symmetric); !ok {
		return false
	}
	if _, ok := b.(Symmetric); !ok {
		return false
	}
	return true
}

// legalTypeVec returns whether v is a *Vector.
func legalTypeVec(v Matrix) bool {
	_, ok := v.(*Vector)
	return ok
}

// legalTypesVecVec returns whether both inputs are *Vector.
func legalTypesVecVec(a, b Matrix) bool {
	if _, ok := a.(*Vector); !ok {
		return false
	}
	if _, ok := b.(*Vector); !ok {
		return false
	}
	return true
}

// legalTypesNotVecVec returns whether the first input is an arbitrary Matrix
// and the second input is a *Vector.
func legalTypesNotVecVec(a, b Matrix) bool {
	_, ok := b.(*Vector)
	return ok
}

// legalDims returns whether {m,n} is a valid dimension of the given matrix type.
func legalDims(a Matrix, m, n int) bool {
	switch t := a.(type) {
	default:
		panic("legal dims type not coded")
	case Untransposer:
		return legalDims(t.Untranspose(), n, m)
	case *Dense, *basicMatrix:
		if m < 0 || n < 0 {
			return false
		}
		return true
	case *SymDense, *TriDense, *basicSymmetric, *basicTriangular:
		if m < 0 || n < 0 || m != n {
			return false
		}
		return true
	case *Vector:
		if m < 0 || n < 0 {
			return false
		}
		return n == 1
	}
}

// returnAs returns the matrix a with the type of t. Used for making a concrete
// type and changing to the basic form.
func returnAs(a, t Matrix) Matrix {
	switch mat := a.(type) {
	default:
		panic("unknown type for a")
	case *Dense:
		switch t.(type) {
		default:
			panic("bad type")
		case *Dense:
			return mat
		case *basicMatrix:
			return asBasicMatrix(mat)
		}
	case *SymDense:
		switch t.(type) {
		default:
			panic("bad type")
		case *SymDense:
			return mat
		case *basicSymmetric:
			return asBasicSymmetric(mat)
		}
	case *TriDense:
		switch t.(type) {
		default:
			panic("bad type")
		case *TriDense:
			return mat
		case *basicTriangular:
			return asBasicTriangular(mat)
		}
	}
}

// retranspose returns the matrix m inside an Untransposer of the type
// of a.
func retranspose(a, m Matrix) Matrix {
	switch a.(type) {
	case TransposeTri:
		return TransposeTri{m.(Triangular)}
	case Transpose:
		return Transpose{m}
	case Untransposer:
		panic("unknown transposer type")
	default:
		panic("a is not an untransposer")
	}
}

// makeRandOf returns a new randomly filled m×n matrix of the underlying matrix type.
func makeRandOf(a Matrix, m, n int) Matrix {
	var rMatrix Matrix
	switch t := a.(type) {
	default:
		panic("unknown type for make rand of")
	case Untransposer:
		rMatrix = retranspose(a, makeRandOf(t.Untranspose(), n, m))
	case *Dense, *basicMatrix:
		mat := NewDense(m, n, nil)
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				mat.Set(i, j, rand.NormFloat64())
			}
		}
		rMatrix = returnAs(mat, t)
	case *Vector:
		if m == 0 && n == 0 {
			return &Vector{}
		}
		if n != 1 {
			panic(fmt.Sprintf("bad vector size: m = %v, n = %v", m, n))
		}
		length := m
		inc := 1
		if t.mat.Inc != 0 {
			inc = t.mat.Inc
		}
		mat := &Vector{
			mat: blas64.Vector{
				Inc:  inc,
				Data: make([]float64, inc*(length-1)+1),
			},
			n: length,
		}
		for i := 0; i < length; i++ {
			mat.SetVec(i, rand.NormFloat64())
		}
		return mat
	case *SymDense, *basicSymmetric:
		if m != n {
			panic("bad size")
		}
		mat := NewSymDense(n, nil)
		for i := 0; i < m; i++ {
			for j := i; j < n; j++ {
				mat.SetSym(i, j, rand.NormFloat64())
			}
		}
		rMatrix = returnAs(mat, t)
	case *TriDense, *basicTriangular:
		if m != n {
			panic("bad size")
		}

		// This is necessary because we are making
		// a triangle from the zero value, which
		// always returns upper as true.
		var triKind matrix.TriKind
		switch t := t.(type) {
		case *TriDense:
			triKind = t.triKind()
		case *basicTriangular:
			triKind = (*TriDense)(t).triKind()
		}

		mat := NewTriDense(n, triKind, nil)
		if triKind == matrix.Upper {
			for i := 0; i < m; i++ {
				for j := i; j < n; j++ {
					mat.SetTri(i, j, rand.NormFloat64())
				}
			}
		} else {
			for i := 0; i < m; i++ {
				for j := 0; j <= i; j++ {
					mat.SetTri(i, j, rand.NormFloat64())
				}
			}
		}
		rMatrix = returnAs(mat, t)
	}
	if mr, mc := rMatrix.Dims(); mr != m || mc != n {
		panic(fmt.Sprintf("makeRandOf for %T returns wrong size: %d×%d != %d×%d", a, m, n, mr, mc))
	}
	return rMatrix
}

// makeCopyOf returns a copy of the matrix.
func makeCopyOf(a Matrix) Matrix {
	switch t := a.(type) {
	default:
		panic("unknown type in makeCopyOf")
	case Untransposer:
		return retranspose(a, makeCopyOf(t.Untranspose()))
	case *Dense, *basicMatrix:
		var m Dense
		m.Clone(a)
		return returnAs(&m, t)
	case *SymDense, *basicSymmetric:
		n := t.(Symmetric).Symmetric()
		m := NewSymDense(n, nil)
		m.CopySym(t.(Symmetric))
		return returnAs(m, t)
	case *TriDense, *basicTriangular:
		n, upper := t.(Triangular).Triangle()
		m := NewTriDense(n, upper, nil)
		if upper {
			for i := 0; i < n; i++ {
				for j := i; j < n; j++ {
					m.SetTri(i, j, t.At(i, j))
				}
			}
		} else {
			for i := 0; i < n; i++ {
				for j := 0; j <= i; j++ {
					m.SetTri(i, j, t.At(i, j))
				}
			}
		}
		return returnAs(m, t)
	case *Vector:
		m := &Vector{
			mat: blas64.Vector{
				Inc:  t.mat.Inc,
				Data: make([]float64, t.mat.Inc*(t.n-1)+1),
			},
			n: t.n,
		}
		copy(m.mat.Data, t.mat.Data)
		return m
	}
}

// sameType returns true if a and b have the same underlying type.
func sameType(a, b Matrix) bool {
	return reflect.ValueOf(a).Type() == reflect.ValueOf(b).Type()
}

// maybeSame returns true if the two matrices could be represented by the same
// pointer.
func maybeSame(receiver, a Matrix) bool {
	rr, rc := receiver.Dims()
	u, trans := a.(Untransposer)
	if trans {
		a = u.Untranspose()
	}
	if !sameType(receiver, a) {
		return false
	}
	ar, ac := a.Dims()
	if rr != ar || rc != ac {
		return false
	}
	if _, ok := a.(Triangular); ok {
		// They are both triangular types. The TriType needs to match
		_, aKind := a.(Triangular).Triangle()
		_, rKind := receiver.(Triangular).Triangle()
		if aKind != rKind {
			return false
		}
	}
	return true
}

// equalApprox returns whether the elements of a and b are the same to within
// the tolerance. If ignoreNaN is true the test is relaxed such that NaN == NaN.
func equalApprox(a, b Matrix, tol float64, ignoreNaN bool) bool {
	ar, ac := a.Dims()
	br, bc := b.Dims()
	if ar != br {
		return false
	}
	if ac != bc {
		return false
	}
	for i := 0; i < ar; i++ {
		for j := 0; j < ac; j++ {
			if !floats.EqualWithinAbsOrRel(a.At(i, j), b.At(i, j), tol, tol) {
				if ignoreNaN && math.IsNaN(a.At(i, j)) && math.IsNaN(b.At(i, j)) {
					continue
				}
				return false
			}
		}
	}
	return true
}

// equal returns true if the matrices have equal entries.
func equal(a, b Matrix) bool {
	ar, ac := a.Dims()
	br, bc := b.Dims()
	if ar != br {
		return false
	}
	if ac != bc {
		return false
	}
	for i := 0; i < ar; i++ {
		for j := 0; j < ac; j++ {
			if a.At(i, j) != b.At(i, j) {
				return false
			}
		}
	}
	return true
}

// isDiagonal returns whether a is a diagonal matrix.
func isDiagonal(a Matrix) bool {
	r, c := a.Dims()
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if a.At(i, j) != 0 && i != j {
				return false
			}
		}
	}
	return true
}

// equalDiagonal returns whether a and b are equal on the diagonal.
func equalDiagonal(a, b Matrix) bool {
	ar, ac := a.Dims()
	br, bc := a.Dims()
	if min(ar, ac) != min(br, bc) {
		return false
	}
	for i := 0; i < min(ar, ac); i++ {
		if a.At(i, i) != b.At(i, i) {
			return false
		}
	}
	return true
}

// underlyingData extracts the underlying data of the matrix a.
func underlyingData(a Matrix) []float64 {
	switch t := a.(type) {
	default:
		panic("matrix type not implemented for extracting underlying data")
	case Untransposer:
		return underlyingData(t.Untranspose())
	case *Dense:
		return t.mat.Data
	case *SymDense:
		return t.mat.Data
	case *TriDense:
		return t.mat.Data
	case *Vector:
		return t.mat.Data
	}
}

// testMatrices is a list of matrix types to test.
// The TriDense types have actual sizes because the return from Triangular is
// only valid when n == 0.
var testMatrices = []Matrix{
	&Dense{},
	&SymDense{},
	NewTriDense(3, true, nil),
	NewTriDense(3, false, nil),
	NewVector(0, nil),
	&Vector{mat: blas64.Vector{Inc: 10}},
	&basicMatrix{},
	&basicSymmetric{},
	&basicTriangular{cap: 3, mat: blas64.Triangular{N: 3, Stride: 3, Uplo: blas.Upper}},
	&basicTriangular{cap: 3, mat: blas64.Triangular{N: 3, Stride: 3, Uplo: blas.Lower}},

	Transpose{&Dense{}},
	Transpose{NewTriDense(3, true, nil)},
	TransposeTri{NewTriDense(3, true, nil)},
	Transpose{NewTriDense(3, false, nil)},
	TransposeTri{NewTriDense(3, false, nil)},
	Transpose{NewVector(0, nil)},
	Transpose{&Vector{mat: blas64.Vector{Inc: 10}}},
	Transpose{&basicMatrix{}},
	Transpose{&basicSymmetric{}},
	Transpose{&basicTriangular{cap: 3, mat: blas64.Triangular{N: 3, Stride: 3, Uplo: blas.Upper}}},
	Transpose{&basicTriangular{cap: 3, mat: blas64.Triangular{N: 3, Stride: 3, Uplo: blas.Lower}}},
}

var sizes = []struct {
	ar, ac int
}{
	{1, 1},
	{1, 3},
	{3, 1},

	{6, 6},
	{6, 11},
	{11, 6},
}

func testOneInputFunc(t *testing.T,
	// name is the name of the function being tested.
	name string,

	// f is the function being tested.
	f func(a Matrix) interface{},

	// denseComparison performs the same operation, but using Dense matrices for
	// comparison.
	denseComparison func(a *Dense) interface{},

	// sameAnswer compares the result from two different evaluations of the function
	// and returns true if they are the same. The specific function being tested
	// determines the definition of "same". It may mean identical or it may mean
	// approximately equal.
	sameAnswer func(a, b interface{}) bool,

	// legalType returns true if the type of the input is a legal type for the
	// input of the function.
	legalType func(a Matrix) bool,

	// legalSize returns true if the size is valid for the function.
	legalSize func(r, c int) bool,
) {
	for _, aMat := range testMatrices {
		for _, test := range sizes {
			// Skip the test if the argument would not be assignable to the
			// method's corresponding input parameter or it is not possible
			// to construct an argument of the requested size.
			if !legalType(aMat) {
				continue
			}
			if !legalDims(aMat, test.ar, test.ac) {
				continue
			}
			a := makeRandOf(aMat, test.ar, test.ac)

			// Compute the true answer if the sizes are legal.
			dimsOK := legalSize(test.ar, test.ac)
			var want interface{}
			if dimsOK {
				var aDense Dense
				aDense.Clone(a)
				want = denseComparison(&aDense)
			}
			aCopy := makeCopyOf(a)
			// Test the method for a zero-value of the receiver.
			aType, aTrans := untranspose(a)
			errStr := fmt.Sprintf("%v(%T), size: %#v, atrans %t", name, aType, test, aTrans)
			var got interface{}
			panicked, err := panics(func() { got = f(a) })
			if !dimsOK && !panicked {
				t.Errorf("Did not panic with illegal size: %s", errStr)
				continue
			}
			if dimsOK && panicked {
				t.Errorf("Panicked with legal size: %s: %v", errStr, err)
				continue
			}
			if !equal(a, aCopy) {
				t.Errorf("First input argument changed in call: %s", errStr)
			}
			if !dimsOK {
				continue
			}
			if !sameAnswer(want, got) {
				t.Errorf("Answer mismatch: %s", errStr)
			}
		}
	}
}

var sizePairs = []struct {
	ar, ac, br, bc int
}{
	{1, 1, 1, 1},
	{6, 6, 6, 6},
	{7, 7, 7, 7},

	{1, 1, 1, 5},
	{1, 1, 5, 1},
	{1, 5, 1, 1},
	{5, 1, 1, 1},

	{5, 5, 5, 1},
	{5, 5, 1, 5},
	{5, 1, 5, 5},
	{1, 5, 5, 5},

	{6, 6, 6, 11},
	{6, 6, 11, 6},
	{6, 11, 6, 6},
	{11, 6, 6, 6},
	{11, 11, 11, 6},
	{11, 11, 6, 11},
	{11, 6, 11, 11},
	{6, 11, 11, 11},

	{1, 1, 5, 5},
	{1, 5, 1, 5},
	{1, 5, 5, 1},
	{5, 1, 1, 5},
	{5, 1, 5, 1},
	{5, 5, 1, 1},
	{6, 6, 11, 11},
	{6, 11, 6, 11},
	{6, 11, 11, 6},
	{11, 6, 6, 11},
	{11, 6, 11, 6},
	{11, 11, 6, 6},

	{1, 1, 17, 11},
	{1, 1, 11, 17},
	{1, 11, 1, 17},
	{1, 17, 1, 11},
	{1, 11, 17, 1},
	{1, 17, 11, 1},
	{11, 1, 1, 17},
	{17, 1, 1, 11},
	{11, 1, 17, 1},
	{17, 1, 11, 1},
	{11, 17, 1, 1},
	{17, 11, 1, 1},

	{6, 6, 1, 11},
	{6, 6, 11, 1},
	{6, 11, 6, 1},
	{6, 1, 6, 11},
	{6, 11, 1, 6},
	{6, 1, 11, 6},
	{11, 6, 6, 1},
	{1, 6, 6, 11},
	{11, 6, 1, 6},
	{1, 6, 11, 6},
	{11, 1, 6, 6},
	{1, 11, 6, 6},

	{6, 6, 17, 1},
	{6, 6, 1, 17},
	{6, 1, 6, 17},
	{6, 17, 6, 1},
	{6, 1, 17, 6},
	{6, 17, 1, 6},
	{1, 6, 6, 17},
	{17, 6, 6, 1},
	{1, 6, 17, 6},
	{17, 6, 1, 6},
	{1, 17, 6, 6},
	{17, 1, 6, 6},

	{6, 6, 17, 11},
	{6, 6, 11, 17},
	{6, 11, 6, 17},
	{6, 17, 6, 11},
	{6, 11, 17, 6},
	{6, 17, 11, 6},
	{11, 6, 6, 17},
	{17, 6, 6, 11},
	{11, 6, 17, 6},
	{17, 6, 11, 6},
	{11, 17, 6, 6},
	{17, 11, 6, 6},
}

func testTwoInputFunc(t *testing.T,
	// name is the name of the function being tested.
	name string,

	// f is the function being tested.
	f func(a, b Matrix) interface{},

	// denseComparison performs the same operation, but using Dense matrices for
	// comparison.
	denseComparison func(a, b *Dense) interface{},

	// sameAnswer compares the result from two different evaluations of the function
	// and returns true if they are the same. The specific function being tested
	// determines the definition of "same". It may mean identical or it may mean
	// approximately equal.
	sameAnswer func(a, b interface{}) bool,

	// legalType returns true if the types of the inputs are legal for the
	// input of the function.
	legalType func(a, b Matrix) bool,

	// legalSize returns true if the sizes are valid for the function.
	legalSize func(ar, ac, br, bc int) bool,
) {
	for _, aMat := range testMatrices {
		for _, bMat := range testMatrices {
			// Loop over all of the size combinations (bigger, smaller, etc.).
			for _, test := range sizePairs {
				// Skip the test if the argument would not be assignable to the
				// method's corresponding input parameter or it is not possible
				// to construct an argument of the requested size.
				if !legalType(aMat, bMat) {
					continue
				}
				if !legalDims(aMat, test.ar, test.ac) {
					continue
				}
				if !legalDims(bMat, test.br, test.bc) {
					continue
				}
				a := makeRandOf(aMat, test.ar, test.ac)
				b := makeRandOf(bMat, test.br, test.bc)

				// Compute the true answer if the sizes are legal.
				dimsOK := legalSize(test.ar, test.ac, test.br, test.bc)
				var want interface{}
				if dimsOK {
					var aDense, bDense Dense
					aDense.Clone(a)
					bDense.Clone(b)
					want = denseComparison(&aDense, &bDense)
				}
				aCopy := makeCopyOf(a)
				bCopy := makeCopyOf(b)
				// Test the method for a zero-value of the receiver.
				aType, aTrans := untranspose(a)
				bType, bTrans := untranspose(b)
				errStr := fmt.Sprintf("%v(%T, %T), size: %#v, atrans %t, btrans %t", name, aType, bType, test, aTrans, bTrans)
				var got interface{}
				panicked, err := panics(func() { got = f(a, b) })
				if !dimsOK && !panicked {
					t.Errorf("Did not panic with illegal size: %s", errStr)
					continue
				}
				if dimsOK && panicked {
					t.Errorf("Panicked with legal size: %s: %v", errStr, err)
					continue
				}
				if !equal(a, aCopy) {
					t.Errorf("First input argument changed in call: %s", errStr)
				}
				if !equal(b, bCopy) {
					t.Errorf("First input argument changed in call: %s", errStr)
				}
				if !dimsOK {
					continue
				}
				if !sameAnswer(want, got) {
					t.Errorf("Answer mismatch: %s", errStr)
				}
			}
		}
	}
}

// testOneInput tests a method that has one matrix input argument
func testOneInput(t *testing.T,
	// name is the name of the method being tested.
	name string,

	// receiver is a value of the receiver type.
	receiver Matrix,

	// method is the generalized receiver.Method(a).
	method func(receiver, a Matrix),

	// denseComparison performs the same operation as method, but with dense
	// matrices for comparison with the result.
	denseComparison func(receiver, a *Dense),

	// legalTypes returns whether the concrete types in Matrix are valid for
	// the method.
	legalType func(a Matrix) bool,

	// legalSize returns whether the matrix sizes are valid for the method.
	legalSize func(ar, ac int) bool,

	// tol is the tolerance for equality when comparing method results.
	tol float64,
) {
	for _, aMat := range testMatrices {
		for _, test := range sizes {
			// Skip the test if the argument would not be assignable to the
			// method's corresponding input parameter or it is not possible
			// to construct an argument of the requested size.
			if !legalType(aMat) {
				continue
			}
			if !legalDims(aMat, test.ar, test.ac) {
				continue
			}
			a := makeRandOf(aMat, test.ar, test.ac)

			// Compute the true answer if the sizes are legal.
			dimsOK := legalSize(test.ar, test.ac)
			var want Dense
			if dimsOK {
				var aDense Dense
				aDense.Clone(a)
				denseComparison(&want, &aDense)
			}
			aCopy := makeCopyOf(a)

			// Test the method for a zero-value of the receiver.
			aType, aTrans := untranspose(a)
			errStr := fmt.Sprintf("%T.%s(%T), size: %#v, atrans %v", receiver, name, aType, test, aTrans)
			zero := makeRandOf(receiver, 0, 0)
			panicked, err := panics(func() { method(zero, a) })
			if !dimsOK && !panicked {
				t.Errorf("Did not panic with illegal size: %s", errStr)
				continue
			}
			if dimsOK && panicked {
				t.Errorf("Panicked with legal size: %s: %v", errStr, err)
				continue
			}
			if !equal(a, aCopy) {
				t.Errorf("First input argument changed in call: %s", errStr)
			}
			if !dimsOK {
				continue
			}
			if !equalApprox(zero, &want, tol, false) {
				t.Errorf("Answer mismatch with zero receiver: %s.\nGot:\n% v\nWant:\n% v\n", errStr, Formatted(zero), Formatted(&want))
				continue
			}

			// Test the method with a non-zero-value of the receiver.
			// The receiver has been overwritten in place so use its size
			// to construct a new random matrix.
			rr, rc := zero.Dims()
			neverZero := makeRandOf(receiver, rr, rc)
			panicked, _ = panics(func() { method(neverZero, a) })
			if panicked {
				t.Errorf("Panicked with non-zero receiver: %s", errStr)
			}
			if !equalApprox(neverZero, &want, tol, false) {
				t.Errorf("Answer mismatch non-zero receiver: %s", errStr)
			}

			// Test with an incorrectly sized matrix.
			switch receiver.(type) {
			default:
				panic("matrix type not coded for incorrect receiver size")
			case *Dense:
				wrongSize := makeRandOf(receiver, rr+1, rc)
				panicked, _ = panics(func() { method(wrongSize, a) })
				if !panicked {
					t.Errorf("Did not panic with wrong number of rows: %s", errStr)
				}
				wrongSize = makeRandOf(receiver, rr, rc+1)
				panicked, _ = panics(func() { method(wrongSize, a) })
				if !panicked {
					t.Errorf("Did not panic with wrong number of columns: %s", errStr)
				}
			case *TriDense, *SymDense:
				// Add to the square size.
				wrongSize := makeRandOf(receiver, rr+1, rc+1)
				panicked, _ = panics(func() { method(wrongSize, a) })
				if !panicked {
					t.Errorf("Did not panic with wrong size: %s", errStr)
				}
			case *Vector:
				// Add to the column length.
				wrongSize := makeRandOf(receiver, rr+1, rc)
				panicked, _ = panics(func() { method(wrongSize, a) })
				if !panicked {
					t.Errorf("Did not panic with wrong number of rows: %s", errStr)
				}
			}

			// The receiver and the input may share a matrix pointer
			// if the type and size of the receiver and one of the
			// arguments match. Test the method works properly
			// when this is the case.
			aMaybeSame := maybeSame(neverZero, a)
			if aMaybeSame {
				aSame := makeCopyOf(a)
				receiver = aSame
				u, ok := aSame.(Untransposer)
				if ok {
					receiver = u.Untranspose()
				}
				preData := underlyingData(receiver)
				panicked, err = panics(func() { method(receiver, aSame) })
				if panicked {
					t.Errorf("Panics when a maybeSame: %s: %v", errStr, err)
				} else {
					if !equalApprox(receiver, &want, tol, false) {
						t.Errorf("Wrong answer when a maybeSame: %s", errStr)
					}
					postData := underlyingData(receiver)
					if !floats.Equal(preData, postData) {
						t.Errorf("Original data slice not modified when a maybeSame: %s", errStr)
					}
				}
			}
		}
	}
}

// testTwoInput tests a method that has two input arguments.
func testTwoInput(t *testing.T,
	// name is the name of the method being tested.
	name string,

	// receiver is a value of the receiver type.
	receiver Matrix,

	// method is the generalized receiver.Method(a, b).
	method func(receiver, a, b Matrix),

	// denseComparison performs the same operation as method, but with dense
	// matrices for comparison with the result.
	denseComparison func(receiver, a, b *Dense),

	// legalTypes returns whether the concrete types in Matrix are valid for
	// the method.
	legalTypes func(a, b Matrix) bool,

	// legalSize returns whether the matrix sizes are valid for the method.
	legalSize func(ar, ac, br, bc int) bool,

	// tol is the tolerance for equality when comparing method results.
	tol float64,
) {
	for _, aMat := range testMatrices {
		for _, bMat := range testMatrices {
			// Loop over all of the size combinations (bigger, smaller, etc.).
			for _, test := range sizePairs {
				// Skip the test if any argument would not be assignable to the
				// method's corresponding input parameter or it is not possible
				// to construct an argument of the requested size.
				if !legalTypes(aMat, bMat) {
					continue
				}
				if !legalDims(aMat, test.ar, test.ac) {
					continue
				}
				if !legalDims(bMat, test.br, test.bc) {
					continue
				}
				a := makeRandOf(aMat, test.ar, test.ac)
				b := makeRandOf(bMat, test.br, test.bc)

				// Compute the true answer if the sizes are legal.
				dimsOK := legalSize(test.ar, test.ac, test.br, test.bc)
				var want Dense
				if dimsOK {
					var aDense, bDense Dense
					aDense.Clone(a)
					bDense.Clone(b)
					denseComparison(&want, &aDense, &bDense)
				}
				aCopy := makeCopyOf(a)
				bCopy := makeCopyOf(b)

				// Test the method for a zero-value of the receiver.
				aType, aTrans := untranspose(a)
				bType, bTrans := untranspose(b)
				errStr := fmt.Sprintf("%T.%s(%T, %T), sizes: %#v, atrans %v, btrans %v", receiver, name, aType, bType, test, aTrans, bTrans)
				zero := makeRandOf(receiver, 0, 0)
				panicked, err := panics(func() { method(zero, a, b) })
				if !dimsOK && !panicked {
					t.Errorf("Did not panic with illegal size: %s", errStr)
					continue
				}
				if dimsOK && panicked {
					t.Errorf("Panicked with legal size: %s: %v", errStr, err)
					continue
				}
				if !equal(a, aCopy) {
					t.Errorf("First input argument changed in call: %s", errStr)
				}
				if !equal(b, bCopy) {
					t.Errorf("Second input argument changed in call: %s", errStr)
				}
				if !dimsOK {
					continue
				}
				wasZero, zero := zero, nil // Nil-out zero so we detect illegal use.
				// NaN equality is allowed because of 0/0 in DivElem test.
				if !equalApprox(wasZero, &want, tol, true) {
					t.Errorf("Answer mismatch with zero receiver: %s", errStr)
					continue
				}

				// Test the method with a non-zero-value of the receiver.
				// The receiver has been overwritten in place so use its size
				// to construct a new random matrix.
				rr, rc := wasZero.Dims()
				neverZero := makeRandOf(receiver, rr, rc)
				panicked, message := panics(func() { method(neverZero, a, b) })
				if panicked {
					t.Errorf("Panicked with non-zero receiver: %s: %s", errStr, message)
				}
				// NaN equality is allowed because of 0/0 in DivElem test.
				if !equalApprox(neverZero, &want, tol, true) {
					t.Errorf("Answer mismatch non-zero receiver: %s", errStr)
				}

				// Test with an incorrectly sized matrix.
				switch receiver.(type) {
				default:
					panic("matrix type not coded for incorrect receiver size")
				case *Dense:
					wrongSize := makeRandOf(receiver, rr+1, rc)
					panicked, _ = panics(func() { method(wrongSize, a, b) })
					if !panicked {
						t.Errorf("Did not panic with wrong number of rows: %s", errStr)
					}
					wrongSize = makeRandOf(receiver, rr, rc+1)
					panicked, _ = panics(func() { method(wrongSize, a, b) })
					if !panicked {
						t.Errorf("Did not panic with wrong number of columns: %s", errStr)
					}
				case *TriDense, *SymDense:
					// Add to the square size.
					wrongSize := makeRandOf(receiver, rr+1, rc+1)
					panicked, _ = panics(func() { method(wrongSize, a, b) })
					if !panicked {
						t.Errorf("Did not panic with wrong size: %s", errStr)
					}
				case *Vector:
					// Add to the column length.
					wrongSize := makeRandOf(receiver, rr+1, rc)
					panicked, _ = panics(func() { method(wrongSize, a, b) })
					if !panicked {
						t.Errorf("Did not panic with wrong number of rows: %s", errStr)
					}
				}

				// The receiver and an input may share a matrix pointer
				// if the type and size of the receiver and one of the
				// arguments match. Test the method works properly
				// when this is the case.
				aMaybeSame := maybeSame(neverZero, a)
				bMaybeSame := maybeSame(neverZero, b)
				if aMaybeSame {
					aSame := makeCopyOf(a)
					receiver = aSame
					u, ok := aSame.(Untransposer)
					if ok {
						receiver = u.Untranspose()
					}
					preData := underlyingData(receiver)
					panicked, err = panics(func() { method(receiver, aSame, b) })
					if panicked {
						t.Errorf("Panics when a maybeSame: %s: %v", errStr, err)
					} else {
						if !equalApprox(receiver, &want, tol, false) {
							t.Errorf("Wrong answer when a maybeSame: %s", errStr)
						}
						postData := underlyingData(receiver)
						if !floats.Equal(preData, postData) {
							t.Errorf("Original data slice not modified when a maybeSame: %s", errStr)
						}
					}
				}
				if bMaybeSame {
					bSame := makeCopyOf(b)
					receiver = bSame
					u, ok := bSame.(Untransposer)
					if ok {
						receiver = u.Untranspose()
					}
					preData := underlyingData(receiver)
					panicked, err = panics(func() { method(receiver, a, bSame) })
					if panicked {
						t.Errorf("Panics when b maybeSame: %s: %v", errStr, err)
					} else {
						if !equalApprox(receiver, &want, tol, false) {
							t.Errorf("Wrong answer when b maybeSame: %s", errStr)
						}
						postData := underlyingData(receiver)
						if !floats.Equal(preData, postData) {
							t.Errorf("Original data slice not modified when b maybeSame: %s", errStr)
						}
					}
				}
				if aMaybeSame && bMaybeSame {
					aSame := makeCopyOf(a)
					receiver = aSame
					u, ok := aSame.(Untransposer)
					if ok {
						receiver = u.Untranspose()
					}
					// Ensure that b is the correct transpose type if applicable.
					// The receiver is always a concrete type so use it.
					bSame := receiver
					u, ok = b.(Untransposer)
					if ok {
						bSame = retranspose(b, receiver)
					}
					// Compute the real answer for this case. It is different
					// from the initial answer since now a and b have the
					// same data.
					zero = makeRandOf(wasZero, 0, 0)
					method(zero, aSame, bSame)
					wasZero, zero = zero, nil // Nil-out zero so we detect illegal use.
					preData := underlyingData(receiver)
					panicked, err = panics(func() { method(receiver, aSame, bSame) })
					if panicked {
						t.Errorf("Panics when both maybeSame: %s: %v", errStr, err)
					} else {
						if !equalApprox(receiver, wasZero, tol, false) {
							t.Errorf("Wrong answer when both maybeSame: %s", errStr)
						}
						postData := underlyingData(receiver)
						if !floats.Equal(preData, postData) {
							t.Errorf("Original data slice not modified when both maybeSame: %s", errStr)
						}
					}
				}
			}
		}
	}
}
