// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mat64

import (
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"testing"

	"github.com/gonum/blas"
	"github.com/gonum/blas/blas64"
	"github.com/gonum/floats"
	"github.com/gonum/matrix"
)

func TestNewSymmetric(t *testing.T) {
	for i, test := range []struct {
		data []float64
		n    int
		mat  *SymDense
	}{
		{
			data: []float64{
				1, 2, 3,
				4, 5, 6,
				7, 8, 9,
			},
			n: 3,
			mat: &SymDense{
				mat: blas64.Symmetric{
					N:      3,
					Stride: 3,
					Uplo:   blas.Upper,
					Data:   []float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
				},
				cap: 3,
			},
		},
	} {
		sym := NewSymDense(test.n, test.data)
		rows, cols := sym.Dims()

		if rows != test.n {
			t.Errorf("unexpected number of rows for test %d: got: %d want: %d", i, rows, test.n)
		}
		if cols != test.n {
			t.Errorf("unexpected number of cols for test %d: got: %d want: %d", i, cols, test.n)
		}
		if !reflect.DeepEqual(sym, test.mat) {
			t.Errorf("unexpected data slice for test %d: got: %v want: %v", i, sym, test.mat)
		}

		m := NewDense(test.n, test.n, test.data)
		if !reflect.DeepEqual(sym.mat.Data, m.mat.Data) {
			t.Errorf("unexpected data slice mismatch for test %d: got: %v want: %v", i, sym.mat.Data, m.mat.Data)
		}
	}

	panicked, message := panics(func() { NewSymDense(3, []float64{1, 2}) })
	if !panicked || message != matrix.ErrShape.Error() {
		t.Error("expected panic for invalid data slice length")
	}
}

func TestSymAtSet(t *testing.T) {
	sym := &SymDense{
		mat: blas64.Symmetric{
			N:      3,
			Stride: 3,
			Uplo:   blas.Upper,
			Data:   []float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		cap: 3,
	}
	rows, cols := sym.Dims()

	// Check At out of bounds
	for _, row := range []int{-1, rows, rows + 1} {
		panicked, message := panics(func() { sym.At(row, 0) })
		if !panicked || message != matrix.ErrRowAccess.Error() {
			t.Errorf("expected panic for invalid row access N=%d r=%d", rows, row)
		}
	}
	for _, col := range []int{-1, cols, cols + 1} {
		panicked, message := panics(func() { sym.At(0, col) })
		if !panicked || message != matrix.ErrColAccess.Error() {
			t.Errorf("expected panic for invalid column access N=%d c=%d", cols, col)
		}
	}

	// Check Set out of bounds
	for _, row := range []int{-1, rows, rows + 1} {
		panicked, message := panics(func() { sym.SetSym(row, 0, 1.2) })
		if !panicked || message != matrix.ErrRowAccess.Error() {
			t.Errorf("expected panic for invalid row access N=%d r=%d", rows, row)
		}
	}
	for _, col := range []int{-1, cols, cols + 1} {
		panicked, message := panics(func() { sym.SetSym(0, col, 1.2) })
		if !panicked || message != matrix.ErrColAccess.Error() {
			t.Errorf("expected panic for invalid column access N=%d c=%d", cols, col)
		}
	}

	for _, st := range []struct {
		row, col  int
		orig, new float64
	}{
		{row: 1, col: 2, orig: 6, new: 15},
		{row: 2, col: 1, orig: 15, new: 12},
	} {
		if e := sym.At(st.row, st.col); e != st.orig {
			t.Errorf("unexpected value for At(%d, %d): got: %v want: %v", st.row, st.col, e, st.orig)
		}
		if e := sym.At(st.col, st.row); e != st.orig {
			t.Errorf("unexpected value for At(%d, %d): got: %v want: %v", st.col, st.row, e, st.orig)
		}
		sym.SetSym(st.row, st.col, st.new)
		if e := sym.At(st.row, st.col); e != st.new {
			t.Errorf("unexpected value for At(%d, %d) after SetSym(%[1]d, %[2]d, %[4]v): got: %[3]v want: %v", st.row, st.col, e, st.new)
		}
		if e := sym.At(st.col, st.row); e != st.new {
			t.Errorf("unexpected value for At(%d, %d) after SetSym(%[2]d, %[1]d, %[4]v): got: %[3]v want: %v", st.col, st.row, e, st.new)
		}
	}
}

func TestSymAdd(t *testing.T) {
	for _, test := range []struct {
		n int
	}{
		{n: 1},
		{n: 2},
		{n: 3},
		{n: 4},
		{n: 5},
		{n: 10},
	} {
		n := test.n
		a := NewSymDense(n, nil)
		for i := range a.mat.Data {
			a.mat.Data[i] = rand.Float64()
		}
		b := NewSymDense(n, nil)
		for i := range a.mat.Data {
			b.mat.Data[i] = rand.Float64()
		}
		var m Dense
		m.Add(a, b)

		// Check with new receiver
		var s SymDense
		s.AddSym(a, b)
		for i := 0; i < n; i++ {
			for j := i; j < n; j++ {
				want := m.At(i, j)
				if got := s.At(i, j); got != want {
					t.Errorf("unexpected value for At(%d, %d): got: %v want: %v", i, j, got, want)
				}
			}
		}

		// Check with equal receiver
		s.CopySym(a)
		s.AddSym(&s, b)
		for i := 0; i < n; i++ {
			for j := i; j < n; j++ {
				want := m.At(i, j)
				if got := s.At(i, j); got != want {
					t.Errorf("unexpected value for At(%d, %d): got: %v want: %v", i, j, got, want)
				}
			}
		}
	}

	method := func(receiver, a, b Matrix) {
		type addSymer interface {
			AddSym(a, b Symmetric)
		}
		rd := receiver.(addSymer)
		rd.AddSym(a.(Symmetric), b.(Symmetric))
	}
	denseComparison := func(receiver, a, b *Dense) {
		receiver.Add(a, b)
	}
	testTwoInput(t, "AddSym", &SymDense{}, method, denseComparison, legalTypesSym, legalSizeSameSquare, 1e-14)
}

func TestCopy(t *testing.T) {
	for _, test := range []struct {
		n int
	}{
		{n: 1},
		{n: 2},
		{n: 3},
		{n: 4},
		{n: 5},
		{n: 10},
	} {
		n := test.n
		a := NewSymDense(n, nil)
		for i := range a.mat.Data {
			a.mat.Data[i] = rand.Float64()
		}
		s := NewSymDense(n, nil)
		s.CopySym(a)
		for i := 0; i < n; i++ {
			for j := i; j < n; j++ {
				want := a.At(i, j)
				if got := s.At(i, j); got != want {
					t.Errorf("unexpected value for At(%d, %d): got: %v want: %v", i, j, got, want)
				}
			}
		}
	}
}

// TODO(kortschak) Roll this into testOneInput when it exists.
// https://github.com/gonum/matrix/issues/171
func TestSymCopyPanic(t *testing.T) {
	var (
		a SymDense
		n int
	)
	m := NewSymDense(1, nil)
	panicked, message := panics(func() { n = m.CopySym(&a) })
	if panicked {
		t.Errorf("unexpected panic: %v", message)
	}
	if n != 0 {
		t.Errorf("unexpected n: got: %d want: 0", n)
	}
}

func TestSymRankOne(t *testing.T) {
	for _, test := range []struct {
		n int
	}{
		{n: 1},
		{n: 2},
		{n: 3},
		{n: 4},
		{n: 5},
		{n: 10},
	} {
		n := test.n
		alpha := 2.0
		a := NewSymDense(n, nil)
		for i := range a.mat.Data {
			a.mat.Data[i] = rand.Float64()
		}
		x := make([]float64, n)
		for i := range x {
			x[i] = rand.Float64()
		}

		xMat := NewDense(n, 1, x)
		var m Dense
		m.Mul(xMat, xMat.T())
		m.Scale(alpha, &m)
		m.Add(&m, a)

		// Check with new receiver
		s := NewSymDense(n, nil)
		s.SymRankOne(a, alpha, NewVector(len(x), x))
		for i := 0; i < n; i++ {
			for j := i; j < n; j++ {
				want := m.At(i, j)
				if got := s.At(i, j); got != want {
					t.Errorf("unexpected value for At(%d, %d): got: %v want: %v", i, j, got, want)
				}
			}
		}

		// Check with reused receiver
		copy(s.mat.Data, a.mat.Data)
		s.SymRankOne(s, alpha, NewVector(len(x), x))
		for i := 0; i < n; i++ {
			for j := i; j < n; j++ {
				want := m.At(i, j)
				if got := s.At(i, j); got != want {
					t.Errorf("unexpected value for At(%d, %d): got: %v want: %v", i, j, got, want)
				}
			}
		}
	}

	alpha := 3.0
	method := func(receiver, a, b Matrix) {
		type SymRankOner interface {
			SymRankOne(a Symmetric, alpha float64, x *Vector)
		}
		rd := receiver.(SymRankOner)
		rd.SymRankOne(a.(Symmetric), alpha, b.(*Vector))
	}
	denseComparison := func(receiver, a, b *Dense) {
		var tmp Dense
		tmp.Mul(b, b.T())
		tmp.Scale(alpha, &tmp)
		receiver.Add(a, &tmp)
	}
	legalTypes := func(a, b Matrix) bool {
		_, ok := a.(Symmetric)
		if !ok {
			return false
		}
		_, ok = b.(*Vector)
		return ok
	}
	legalSize := func(ar, ac, br, bc int) bool {
		if ar != ac {
			return false
		}
		return br == ar
	}
	testTwoInput(t, "SymRankOne", &SymDense{}, method, denseComparison, legalTypes, legalSize, 1e-14)
}

func TestIssue250SymRankOne(t *testing.T) {
	x := NewVector(5, []float64{1, 2, 3, 4, 5})
	var s1, s2 SymDense
	s1.SymRankOne(NewSymDense(5, nil), 1, x)
	s2.SymRankOne(NewSymDense(5, nil), 1, x)
	s2.SymRankOne(NewSymDense(5, nil), 1, x)
	if !Equal(&s1, &s2) {
		t.Error("unexpected result from repeat")
	}
}

func TestRankTwo(t *testing.T) {
	for _, test := range []struct {
		n int
	}{
		{n: 1},
		{n: 2},
		{n: 3},
		{n: 4},
		{n: 5},
		{n: 10},
	} {
		n := test.n
		alpha := 2.0
		a := NewSymDense(n, nil)
		for i := range a.mat.Data {
			a.mat.Data[i] = rand.Float64()
		}
		x := make([]float64, n)
		y := make([]float64, n)
		for i := range x {
			x[i] = rand.Float64()
			y[i] = rand.Float64()
		}

		xMat := NewDense(n, 1, x)
		yMat := NewDense(n, 1, y)
		var m Dense
		m.Mul(xMat, yMat.T())
		var tmp Dense
		tmp.Mul(yMat, xMat.T())
		m.Add(&m, &tmp)
		m.Scale(alpha, &m)
		m.Add(&m, a)

		// Check with new receiver
		s := NewSymDense(n, nil)
		s.RankTwo(a, alpha, NewVector(len(x), x), NewVector(len(y), y))
		for i := 0; i < n; i++ {
			for j := i; j < n; j++ {
				if !floats.EqualWithinAbsOrRel(s.At(i, j), m.At(i, j), 1e-14, 1e-14) {
					t.Errorf("unexpected element value at (%d,%d): got: %f want: %f", i, j, m.At(i, j), s.At(i, j))
				}
			}
		}

		// Check with reused receiver
		copy(s.mat.Data, a.mat.Data)
		s.RankTwo(s, alpha, NewVector(len(x), x), NewVector(len(y), y))
		for i := 0; i < n; i++ {
			for j := i; j < n; j++ {
				if !floats.EqualWithinAbsOrRel(s.At(i, j), m.At(i, j), 1e-14, 1e-14) {
					t.Errorf("unexpected element value at (%d,%d): got: %f want: %f", i, j, m.At(i, j), s.At(i, j))
				}
			}
		}
	}
}

func TestSymRankK(t *testing.T) {
	alpha := 3.0
	method := func(receiver, a, b Matrix) {
		type SymRankKer interface {
			SymRankK(a Symmetric, alpha float64, x Matrix)
		}
		rd := receiver.(SymRankKer)
		rd.SymRankK(a.(Symmetric), alpha, b)
	}
	denseComparison := func(receiver, a, b *Dense) {
		var tmp Dense
		tmp.Mul(b, b.T())
		tmp.Scale(alpha, &tmp)
		receiver.Add(a, &tmp)
	}
	legalTypes := func(a, b Matrix) bool {
		_, ok := a.(Symmetric)
		return ok
	}
	legalSize := func(ar, ac, br, bc int) bool {
		if ar != ac {
			return false
		}
		return br == ar
	}
	testTwoInput(t, "SymRankK", &SymDense{}, method, denseComparison, legalTypes, legalSize, 1e-14)
}

func TestSymOuterK(t *testing.T) {
	for _, f := range []float64{0.5, 1, 3} {
		method := func(receiver, x Matrix) {
			type SymOuterKer interface {
				SymOuterK(alpha float64, x Matrix)
			}
			rd := receiver.(SymOuterKer)
			rd.SymOuterK(f, x)
		}
		denseComparison := func(receiver, x *Dense) {
			receiver.Mul(x, x.T())
			receiver.Scale(f, receiver)
		}
		testOneInput(t, "SymOuterK", &SymDense{}, method, denseComparison, isAnyType, isAnySize, 1e-14)
	}
}

func TestIssue250SymOuterK(t *testing.T) {
	x := NewVector(5, []float64{1, 2, 3, 4, 5})
	var s1, s2 SymDense
	s1.SymOuterK(1, x)
	s2.SymOuterK(1, x)
	s2.SymOuterK(1, x)
	if !Equal(&s1, &s2) {
		t.Error("unexpected result from repeat")
	}
}

func TestScaleSym(t *testing.T) {
	for _, f := range []float64{0.5, 1, 3} {
		method := func(receiver, a Matrix) {
			type ScaleSymer interface {
				ScaleSym(f float64, a Symmetric)
			}
			rd := receiver.(ScaleSymer)
			rd.ScaleSym(f, a.(Symmetric))
		}
		denseComparison := func(receiver, a *Dense) {
			receiver.Scale(f, a)
		}
		testOneInput(t, "ScaleSym", &SymDense{}, method, denseComparison, legalTypeSym, isSquare, 1e-14)
	}
}

func TestSubsetSym(t *testing.T) {
	for _, test := range []struct {
		a    *SymDense
		dims []int
		ans  *SymDense
	}{
		{
			a: NewSymDense(3, []float64{
				1, 2, 3,
				0, 4, 5,
				0, 0, 6,
			}),
			dims: []int{0, 2},
			ans: NewSymDense(2, []float64{
				1, 3,
				0, 6,
			}),
		},
		{
			a: NewSymDense(3, []float64{
				1, 2, 3,
				0, 4, 5,
				0, 0, 6,
			}),
			dims: []int{2, 0},
			ans: NewSymDense(2, []float64{
				6, 3,
				0, 1,
			}),
		},
		{
			a: NewSymDense(3, []float64{
				1, 2, 3,
				0, 4, 5,
				0, 0, 6,
			}),
			dims: []int{1, 1, 1},
			ans: NewSymDense(3, []float64{
				4, 4, 4,
				0, 4, 4,
				0, 0, 4,
			}),
		},
	} {
		var s SymDense
		s.SubsetSym(test.a, test.dims)
		if !Equal(&s, test.ans) {
			t.Errorf("SubsetSym mismatch dims %v\nGot:\n% v\nWant:\n% v\n", test.dims, s, test.ans)
		}
	}

	dims := []int{0, 2}
	maxDim := dims[0]
	for _, v := range dims {
		if maxDim < v {
			maxDim = v
		}
	}
	method := func(receiver, a Matrix) {
		type SubsetSymer interface {
			SubsetSym(a Symmetric, set []int)
		}
		rd := receiver.(SubsetSymer)
		rd.SubsetSym(a.(Symmetric), dims)
	}
	denseComparison := func(receiver, a *Dense) {
		*receiver = *NewDense(len(dims), len(dims), nil)
		sz := len(dims)
		for i := 0; i < sz; i++ {
			for j := 0; j < sz; j++ {
				receiver.Set(i, j, a.At(dims[i], dims[j]))
			}
		}
	}
	legalSize := func(ar, ac int) bool {
		return ar == ac && ar > maxDim
	}

	testOneInput(t, "SubsetSym", &SymDense{}, method, denseComparison, legalTypeSym, legalSize, 0)
}

func TestViewGrowSquare(t *testing.T) {
	// n is the size of the original SymDense.
	// The first view uses start1, span1. The second view uses start2, span2 on
	// the first view.
	for _, test := range []struct {
		n, start1, span1, start2, span2 int
	}{
		{10, 0, 10, 0, 10},
		{10, 0, 8, 0, 8},
		{10, 2, 8, 0, 6},
		{10, 2, 7, 4, 2},
		{10, 2, 6, 0, 5},
	} {
		n := test.n
		s := NewSymDense(n, nil)
		for i := 0; i < n; i++ {
			for j := i; j < n; j++ {
				s.SetSym(i, j, float64((i+1)*n+j+1))
			}
		}

		// Take a subset and check the view matches.
		start1 := test.start1
		span1 := test.span1
		v := s.SliceSquare(start1, start1+span1).(*SymDense)
		for i := 0; i < span1; i++ {
			for j := i; j < span1; j++ {
				if v.At(i, j) != s.At(start1+i, start1+j) {
					t.Errorf("View mismatch")
				}
			}
		}

		start2 := test.start2
		span2 := test.span2
		v2 := v.SliceSquare(start2, start2+span2).(*SymDense)

		for i := 0; i < span2; i++ {
			for j := i; j < span2; j++ {
				if v2.At(i, j) != s.At(start1+start2+i, start1+start2+j) {
					t.Errorf("Second view mismatch")
				}
			}
		}

		// Check that a write to the view is reflected in the original.
		v2.SetSym(0, 0, 1.2)
		if s.At(start1+start2, start1+start2) != 1.2 {
			t.Errorf("Write to view not reflected in original")
		}

		// Grow the matrix back to the original view
		gn := n - start1 - start2
		g := v2.GrowSquare(gn - v2.Symmetric()).(*SymDense)
		g.SetSym(1, 1, 2.2)

		for i := 0; i < gn; i++ {
			for j := 0; j < gn; j++ {
				if g.At(i, j) != s.At(start1+start2+i, start1+start2+j) {
					t.Errorf("Grow mismatch")

					fmt.Printf("g=\n% v\n", Formatted(g))
					fmt.Printf("s=\n% v\n", Formatted(s))
					os.Exit(1)
				}
			}
		}

		// View g, then grow it and make sure all the elements were copied.
		gv := g.SliceSquare(0, gn-1).(*SymDense)

		gg := gv.GrowSquare(2)
		for i := 0; i < gn; i++ {
			for j := 0; j < gn; j++ {
				if g.At(i, j) != gg.At(i, j) {
					t.Errorf("Expand mismatch")
				}
			}
		}
	}
}

func TestPowPSD(t *testing.T) {
	for cas, test := range []struct {
		a   *SymDense
		pow float64
		ans *SymDense
	}{
		// Comparison with Matlab.
		{
			a:   NewSymDense(2, []float64{10, 5, 5, 12}),
			pow: 0.5,
			ans: NewSymDense(2, []float64{3.065533767740645, 0.776210486171016, 0.776210486171016, 3.376017962209052}),
		},
		{
			a:   NewSymDense(2, []float64{11, -1, -1, 8}),
			pow: 0.5,
			ans: NewSymDense(2, []float64{3.312618742210524, -0.162963396980939, -0.162963396980939, 2.823728551267709}),
		},
		{
			a:   NewSymDense(2, []float64{10, 5, 5, 12}),
			pow: -0.5,
			ans: NewSymDense(2, []float64{0.346372134547712, -0.079637515547296, -0.079637515547296, 0.314517128328794}),
		},
		{
			a:   NewSymDense(3, []float64{15, -1, -3, -1, 8, 6, -3, 6, 14}),
			pow: 0.6,
			ans: NewSymDense(3, []float64{
				5.051214323034288, -0.163162161893975, -0.612153996497505,
				-0.163162161893976, 3.283474884617009, 1.432842761381493,
				-0.612153996497505, 1.432842761381494, 4.695873060862573,
			}),
		},
	} {
		var s SymDense
		err := s.PowPSD(test.a, test.pow)
		if err != nil {
			panic("bad test")
		}
		if !EqualApprox(&s, test.ans, 1e-10) {
			t.Errorf("Case %d, pow mismatch", cas)
			fmt.Println(Formatted(&s))
			fmt.Println(Formatted(test.ans))
		}
	}

	// Compare with Dense.Pow
	rnd := rand.New(rand.NewSource(1))
	for dim := 2; dim < 10; dim++ {
		for pow := 2; pow < 6; pow++ {
			a := NewDense(dim, dim, nil)
			for i := 0; i < dim; i++ {
				for j := 0; j < dim; j++ {
					a.Set(i, j, rnd.Float64())
				}
			}
			var mat SymDense
			mat.SymOuterK(1, a)

			var sym SymDense
			sym.PowPSD(&mat, float64(pow))

			var dense Dense
			dense.Pow(&mat, pow)

			if !EqualApprox(&sym, &dense, 1e-10) {
				t.Errorf("Dim %d: pow mismatch")
			}
		}
	}
}
