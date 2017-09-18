// Copyright ©2013 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mat64

import (
	"fmt"
	"math"
	"reflect"
	"testing"

	"github.com/gonum/blas"
	"github.com/gonum/blas/blas64"
	"github.com/gonum/floats"
	"github.com/gonum/matrix"
)

func panics(fn func()) (panicked bool, message string) {
	defer func() {
		r := recover()
		panicked = r != nil
		message = fmt.Sprint(r)
	}()
	fn()
	return
}

func flatten(f [][]float64) (r, c int, d []float64) {
	r = len(f)
	if r == 0 {
		panic("bad test: no row")
	}
	c = len(f[0])
	d = make([]float64, 0, r*c)
	for _, row := range f {
		if len(row) != c {
			panic("bad test: ragged input")
		}
		d = append(d, row...)
	}
	return r, c, d
}

func unflatten(r, c int, d []float64) [][]float64 {
	m := make([][]float64, r)
	for i := 0; i < r; i++ {
		m[i] = d[i*c : (i+1)*c]
	}
	return m
}

// eye returns a new identity matrix of size n×n.
func eye(n int) *Dense {
	d := make([]float64, n*n)
	for i := 0; i < n*n; i += n + 1 {
		d[i] = 1
	}
	return NewDense(n, n, d)
}

func TestCol(t *testing.T) {
	for id, af := range [][][]float64{
		{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
			{10, 11, 12},
		},
		{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
		},
	} {
		a := NewDense(flatten(af))
		col := make([]float64, a.mat.Rows)
		for j := range af[0] {
			for i := range col {
				col[i] = float64(i*a.mat.Cols + j + 1)
			}

			if got := Col(nil, j, a); !reflect.DeepEqual(got, col) {
				t.Errorf("test %d: unexpected values returned for dense col %d: got: %v want: %v",
					id, j, got, col)
			}

			got := make([]float64, a.mat.Rows)
			if Col(got, j, a); !reflect.DeepEqual(got, col) {
				t.Errorf("test %d: unexpected values filled for dense col %d: got: %v want: %v",
					id, j, got, col)
			}
		}
	}

	denseComparison := func(a *Dense) interface{} {
		r, c := a.Dims()
		ans := make([][]float64, c)
		for j := range ans {
			ans[j] = make([]float64, r)
			for i := range ans[j] {
				ans[j][i] = a.At(i, j)
			}
		}
		return ans
	}

	f := func(a Matrix) interface{} {
		_, c := a.Dims()
		ans := make([][]float64, c)
		for j := range ans {
			ans[j] = Col(nil, j, a)
		}
		return ans
	}
	testOneInputFunc(t, "Col", f, denseComparison, sameAnswerF64SliceOfSlice, isAnyType, isAnySize)

	f = func(a Matrix) interface{} {
		r, c := a.Dims()
		ans := make([][]float64, c)
		for j := range ans {
			ans[j] = make([]float64, r)
			Col(ans[j], j, a)
		}
		return ans
	}
	testOneInputFunc(t, "Col", f, denseComparison, sameAnswerF64SliceOfSlice, isAnyType, isAnySize)
}

func TestRow(t *testing.T) {
	for id, af := range [][][]float64{
		{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
			{10, 11, 12},
		},
		{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
		},
	} {
		a := NewDense(flatten(af))
		for i, row := range af {
			if got := Row(nil, i, a); !reflect.DeepEqual(got, row) {
				t.Errorf("test %d: unexpected values returned for dense row %d: got: %v want: %v",
					id, i, got, row)
			}

			got := make([]float64, len(row))
			if Row(got, i, a); !reflect.DeepEqual(got, row) {
				t.Errorf("test %d: unexpected values filled for dense row %d: got: %v want: %v",
					id, i, got, row)
			}
		}
	}

	denseComparison := func(a *Dense) interface{} {
		r, c := a.Dims()
		ans := make([][]float64, r)
		for i := range ans {
			ans[i] = make([]float64, c)
			for j := range ans[i] {
				ans[i][j] = a.At(i, j)
			}
		}
		return ans
	}

	f := func(a Matrix) interface{} {
		r, _ := a.Dims()
		ans := make([][]float64, r)
		for i := range ans {
			ans[i] = Row(nil, i, a)
		}
		return ans
	}
	testOneInputFunc(t, "Row", f, denseComparison, sameAnswerF64SliceOfSlice, isAnyType, isAnySize)

	f = func(a Matrix) interface{} {
		r, c := a.Dims()
		ans := make([][]float64, r)
		for i := range ans {
			ans[i] = make([]float64, c)
			Row(ans[i], i, a)
		}
		return ans
	}
	testOneInputFunc(t, "Row", f, denseComparison, sameAnswerF64SliceOfSlice, isAnyType, isAnySize)
}

func TestCond(t *testing.T) {
	for i, test := range []struct {
		a       *Dense
		condOne float64
		condTwo float64
		condInf float64
	}{
		{
			a: NewDense(3, 3, []float64{
				8, 1, 6,
				3, 5, 7,
				4, 9, 2,
			}),
			condOne: 16.0 / 3.0,
			condTwo: 4.330127018922192,
			condInf: 16.0 / 3.0,
		},
		{
			a: NewDense(4, 4, []float64{
				2, 9, 3, 2,
				10, 9, 9, 3,
				1, 1, 5, 2,
				8, 4, 10, 2,
			}),
			condOne: 1 / 0.024740155174938,
			condTwo: 34.521576567075087,
			condInf: 1 / 0.012034465570035,
		},
		{
			a: NewDense(3, 3, []float64{
				5, 6, 7,
				8, -2, 1,
				7, 7, 7}),
			condOne: 30.769230769230749,
			condTwo: 21.662689498448440,
			condInf: 31.153846153846136,
		},
	} {
		condOne := Cond(test.a, 1)
		if !floats.EqualWithinAbsOrRel(test.condOne, condOne, 1e-13, 1e-13) {
			t.Errorf("Case %d: inf norm mismatch. Want %v, got %v", i, test.condOne, condOne)
		}
		condTwo := Cond(test.a, 2)
		if !floats.EqualWithinAbsOrRel(test.condTwo, condTwo, 1e-13, 1e-13) {
			t.Errorf("Case %d: inf norm mismatch. Want %v, got %v", i, test.condTwo, condTwo)
		}
		condInf := Cond(test.a, math.Inf(1))
		if !floats.EqualWithinAbsOrRel(test.condInf, condInf, 1e-13, 1e-13) {
			t.Errorf("Case %d: inf norm mismatch. Want %v, got %v", i, test.condInf, condInf)
		}
	}

	for _, test := range []struct {
		name string
		norm float64
	}{
		{
			name: "CondOne",
			norm: 1,
		},
		{
			name: "CondTwo",
			norm: 2,
		},
		{
			name: "CondInf",
			norm: math.Inf(1),
		},
	} {
		f := func(a Matrix) interface{} {
			return Cond(a, test.norm)
		}
		denseComparison := func(a *Dense) interface{} {
			return Cond(a, test.norm)
		}
		testOneInputFunc(t, test.name, f, denseComparison, sameAnswerFloatApproxTol(1e-12), isAnyType, isAnySize)
	}
}

func TestDet(t *testing.T) {
	for c, test := range []struct {
		a   *Dense
		ans float64
	}{
		{
			a:   NewDense(2, 2, []float64{1, 0, 0, 1}),
			ans: 1,
		},
		{
			a:   NewDense(2, 2, []float64{1, 0, 0, -1}),
			ans: -1,
		},
		{
			a: NewDense(3, 3, []float64{
				1, 2, 0,
				0, 1, 2,
				0, 2, 1,
			}),
			ans: -3,
		},
		{
			a: NewDense(3, 3, []float64{
				1, 2, 3,
				5, 7, 9,
				6, 9, 12,
			}),
			ans: 0,
		},
	} {
		a := DenseCopyOf(test.a)
		det := Det(a)
		if !Equal(a, test.a) {
			t.Errorf("Input matrix changed during Det. Case %d.", c)
		}
		if !floats.EqualWithinAbsOrRel(det, test.ans, 1e-14, 1e-14) {
			t.Errorf("Det mismatch case %d. Got %v, want %v", c, det, test.ans)
		}
	}
	// Perform the normal list test to ensure it works for all types.
	f := func(a Matrix) interface{} {
		return Det(a)
	}
	denseComparison := func(a *Dense) interface{} {
		return Det(a)
	}
	testOneInputFunc(t, "Det", f, denseComparison, sameAnswerFloatApproxTol(1e-12), isAnyType, isSquare)

	// Check that it gives approximately the same answer as Cholesky
	// Ensure the input matrices are wider than tall so they are full rank
	isWide := func(ar, ac int) bool {
		return ar <= ac
	}
	f = func(a Matrix) interface{} {
		ar, ac := a.Dims()
		if !isWide(ar, ac) {
			panic(matrix.ErrShape)
		}
		var tmp Dense
		tmp.Mul(a, a.T())
		return Det(&tmp)
	}
	denseComparison = func(a *Dense) interface{} {
		ar, ac := a.Dims()
		if !isWide(ar, ac) {
			panic(matrix.ErrShape)
		}
		var tmp SymDense
		tmp.SymOuterK(1, a)
		var chol Cholesky
		ok := chol.Factorize(&tmp)
		if !ok {
			panic("bad chol test")
		}
		return chol.Det()
	}
	testOneInputFunc(t, "DetVsChol", f, denseComparison, sameAnswerFloatApproxTol(1e-10), isAnyType, isWide)
}

func TestDot(t *testing.T) {
	f := func(a, b Matrix) interface{} {
		return Dot(a.(*Vector), b.(*Vector))
	}
	denseComparison := func(a, b *Dense) interface{} {
		ra, ca := a.Dims()
		rb, cb := b.Dims()
		if ra != rb || ca != cb {
			panic(matrix.ErrShape)
		}
		var sum float64
		for i := 0; i < ra; i++ {
			for j := 0; j < ca; j++ {
				sum += a.At(i, j) * b.At(i, j)
			}
		}
		return sum
	}
	testTwoInputFunc(t, "Dot", f, denseComparison, sameAnswerFloatApproxTol(1e-12), legalTypesVecVec, legalSizeSameVec)
}

func TestEqual(t *testing.T) {
	f := func(a, b Matrix) interface{} {
		return Equal(a, b)
	}
	denseComparison := func(a, b *Dense) interface{} {
		return Equal(a, b)
	}
	testTwoInputFunc(t, "Equal", f, denseComparison, sameAnswerBool, legalTypesAll, isAnySize2)
}

func TestMax(t *testing.T) {
	// A direct test of Max with *Dense arguments is in TestNewDense.
	f := func(a Matrix) interface{} {
		return Max(a)
	}
	denseComparison := func(a *Dense) interface{} {
		return Max(a)
	}
	testOneInputFunc(t, "Max", f, denseComparison, sameAnswerFloat, isAnyType, isAnySize)
}

func TestMin(t *testing.T) {
	// A direct test of Min with *Dense arguments is in TestNewDense.
	f := func(a Matrix) interface{} {
		return Min(a)
	}
	denseComparison := func(a *Dense) interface{} {
		return Min(a)
	}
	testOneInputFunc(t, "Min", f, denseComparison, sameAnswerFloat, isAnyType, isAnySize)
}

func TestNorm(t *testing.T) {
	for i, test := range []struct {
		a    [][]float64
		ord  float64
		norm float64
	}{
		{
			a:    [][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}},
			ord:  1,
			norm: 30,
		},
		{
			a:    [][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}},
			ord:  2,
			norm: 25.495097567963924,
		},
		{
			a:    [][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}},
			ord:  math.Inf(1),
			norm: 33,
		},
		{
			a:    [][]float64{{1, -2, -2}, {-4, 5, 6}},
			ord:  1,
			norm: 8,
		},
		{
			a:    [][]float64{{1, -2, -2}, {-4, 5, 6}},
			ord:  math.Inf(1),
			norm: 15,
		},
	} {
		a := NewDense(flatten(test.a))
		if math.Abs(Norm(a, test.ord)-test.norm) > 1e-14 {
			t.Errorf("Mismatch test %d: %v norm = %f", i, test.a, test.norm)
		}
	}

	for _, test := range []struct {
		name string
		norm float64
	}{
		{"NormOne", 1},
		{"NormTwo", 2},
		{"NormInf", math.Inf(1)},
	} {
		f := func(a Matrix) interface{} {
			return Norm(a, test.norm)
		}
		denseComparison := func(a *Dense) interface{} {
			return Norm(a, test.norm)
		}
		testOneInputFunc(t, test.name, f, denseComparison, sameAnswerFloatApproxTol(1e-12), isAnyType, isAnySize)
	}
}

func TestNormZero(t *testing.T) {
	for _, a := range []Matrix{
		&Dense{},
		&SymDense{},
		&SymDense{mat: blas64.Symmetric{Uplo: blas.Upper}},
		&TriDense{},
		&TriDense{mat: blas64.Triangular{Uplo: blas.Upper, Diag: blas.NonUnit}},
		&Vector{},
	} {
		for _, norm := range []float64{1, 2, math.Inf(1)} {
			panicked, message := panics(func() { Norm(a, norm) })
			if !panicked {
				t.Errorf("expected panic for Norm(&%T{}, %v)", a, norm)
			}
			if message != matrix.ErrShape.Error() {
				t.Errorf("unexpected panic string for Norm(&%T{}, %v): got:%s want:%s",
					a, norm, message, matrix.ErrShape.Error())
			}
		}
	}
}

func TestSum(t *testing.T) {
	f := func(a Matrix) interface{} {
		return Sum(a)
	}
	denseComparison := func(a *Dense) interface{} {
		return Sum(a)
	}
	testOneInputFunc(t, "Sum", f, denseComparison, sameAnswerFloatApproxTol(1e-12), isAnyType, isAnySize)
}

func TestTrace(t *testing.T) {
	for _, test := range []struct {
		a     *Dense
		trace float64
	}{
		{
			a:     NewDense(3, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}),
			trace: 15,
		},
	} {
		trace := Trace(test.a)
		if trace != test.trace {
			t.Errorf("Trace mismatch. Want %v, got %v", test.trace, trace)
		}
	}
	f := func(a Matrix) interface{} {
		return Trace(a)
	}
	denseComparison := func(a *Dense) interface{} {
		return Trace(a)
	}
	testOneInputFunc(t, "Trace", f, denseComparison, sameAnswerFloat, isAnyType, isSquare)
}
