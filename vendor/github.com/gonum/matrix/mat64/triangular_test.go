package mat64

import (
	"math"
	"math/rand"
	"reflect"
	"testing"

	"github.com/gonum/blas"
	"github.com/gonum/blas/blas64"
	"github.com/gonum/matrix"
)

func TestNewTriangular(t *testing.T) {
	for i, test := range []struct {
		data []float64
		n    int
		kind matrix.TriKind
		mat  *TriDense
	}{
		{
			data: []float64{
				1, 2, 3,
				4, 5, 6,
				7, 8, 9,
			},
			n:    3,
			kind: matrix.Upper,
			mat: &TriDense{
				mat: blas64.Triangular{
					N:      3,
					Stride: 3,
					Uplo:   blas.Upper,
					Data:   []float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
					Diag:   blas.NonUnit,
				},
				cap: 3,
			},
		},
	} {
		tri := NewTriDense(test.n, test.kind, test.data)
		rows, cols := tri.Dims()

		if rows != test.n {
			t.Errorf("unexpected number of rows for test %d: got: %d want: %d", i, rows, test.n)
		}
		if cols != test.n {
			t.Errorf("unexpected number of cols for test %d: got: %d want: %d", i, cols, test.n)
		}
		if !reflect.DeepEqual(tri, test.mat) {
			t.Errorf("unexpected data slice for test %d: got: %v want: %v", i, tri, test.mat)
		}
	}

	for _, kind := range []matrix.TriKind{matrix.Lower, matrix.Upper} {
		panicked, message := panics(func() { NewTriDense(3, kind, []float64{1, 2}) })
		if !panicked || message != matrix.ErrShape.Error() {
			t.Errorf("expected panic for invalid data slice length for upper=%t", kind)
		}
	}
}

func TestTriAtSet(t *testing.T) {
	tri := &TriDense{
		mat: blas64.Triangular{
			N:      3,
			Stride: 3,
			Uplo:   blas.Upper,
			Data:   []float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
			Diag:   blas.NonUnit,
		},
		cap: 3,
	}

	rows, cols := tri.Dims()

	// Check At out of bounds
	for _, row := range []int{-1, rows, rows + 1} {
		panicked, message := panics(func() { tri.At(row, 0) })
		if !panicked || message != matrix.ErrRowAccess.Error() {
			t.Errorf("expected panic for invalid row access N=%d r=%d", rows, row)
		}
	}
	for _, col := range []int{-1, cols, cols + 1} {
		panicked, message := panics(func() { tri.At(0, col) })
		if !panicked || message != matrix.ErrColAccess.Error() {
			t.Errorf("expected panic for invalid column access N=%d c=%d", cols, col)
		}
	}

	// Check Set out of bounds
	for _, row := range []int{-1, rows, rows + 1} {
		panicked, message := panics(func() { tri.SetTri(row, 0, 1.2) })
		if !panicked || message != matrix.ErrRowAccess.Error() {
			t.Errorf("expected panic for invalid row access N=%d r=%d", rows, row)
		}
	}
	for _, col := range []int{-1, cols, cols + 1} {
		panicked, message := panics(func() { tri.SetTri(0, col, 1.2) })
		if !panicked || message != matrix.ErrColAccess.Error() {
			t.Errorf("expected panic for invalid column access N=%d c=%d", cols, col)
		}
	}

	for _, st := range []struct {
		row, col int
		uplo     blas.Uplo
	}{
		{row: 2, col: 1, uplo: blas.Upper},
		{row: 1, col: 2, uplo: blas.Lower},
	} {
		tri.mat.Uplo = st.uplo
		panicked, message := panics(func() { tri.SetTri(st.row, st.col, 1.2) })
		if !panicked || message != matrix.ErrTriangleSet.Error() {
			t.Errorf("expected panic for %+v", st)
		}
	}

	for _, st := range []struct {
		row, col  int
		uplo      blas.Uplo
		orig, new float64
	}{
		{row: 2, col: 1, uplo: blas.Lower, orig: 8, new: 15},
		{row: 1, col: 2, uplo: blas.Upper, orig: 6, new: 15},
	} {
		tri.mat.Uplo = st.uplo
		if e := tri.At(st.row, st.col); e != st.orig {
			t.Errorf("unexpected value for At(%d, %d): got: %v want: %v", st.row, st.col, e, st.orig)
		}
		tri.SetTri(st.row, st.col, st.new)
		if e := tri.At(st.row, st.col); e != st.new {
			t.Errorf("unexpected value for At(%d, %d) after SetTri(%[1]d, %d, %v): got: %v want: %[3]v", st.row, st.col, st.new, e)
		}
	}
}

func TestTriDenseCopy(t *testing.T) {
	for i := 0; i < 100; i++ {
		size := rand.Intn(100)
		r, err := randDense(size, 0.9, rand.NormFloat64)
		if size == 0 {
			if err != matrix.ErrZeroLength {
				t.Fatalf("expected error %v: got: %v", matrix.ErrZeroLength, err)
			}
			continue
		}
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		u := NewTriDense(size, true, nil)
		l := NewTriDense(size, false, nil)

		for _, typ := range []Matrix{r, (*basicMatrix)(r)} {
			for j := range u.mat.Data {
				u.mat.Data[j] = math.NaN()
				l.mat.Data[j] = math.NaN()
			}
			u.Copy(typ)
			l.Copy(typ)
			for m := 0; m < size; m++ {
				for n := 0; n < size; n++ {
					want := typ.At(m, n)
					switch {
					case m < n: // Upper triangular matrix.
						if got := u.At(m, n); got != want {
							t.Errorf("unexpected upper value for At(%d, %d) for test %d: got: %v want: %v", m, n, i, got, want)
						}
					case m == n: // Diagonal matrix.
						if got := u.At(m, n); got != want {
							t.Errorf("unexpected upper value for At(%d, %d) for test %d: got: %v want: %v", m, n, i, got, want)
						}
						if got := l.At(m, n); got != want {
							t.Errorf("unexpected diagonal value for At(%d, %d) for test %d: got: %v want: %v", m, n, i, got, want)
						}
					case m < n: // Lower triangular matrix.
						if got := l.At(m, n); got != want {
							t.Errorf("unexpected lower value for At(%d, %d) for test %d: got: %v want: %v", m, n, i, got, want)
						}
					}
				}
			}
		}
	}
}

func TestTriTriDenseCopy(t *testing.T) {
	for i := 0; i < 100; i++ {
		size := rand.Intn(100)
		r, err := randDense(size, 1, rand.NormFloat64)
		if size == 0 {
			if err != matrix.ErrZeroLength {
				t.Fatalf("expected error %v: got: %v", matrix.ErrZeroLength, err)
			}
			continue
		}
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		ur := NewTriDense(size, true, nil)
		lr := NewTriDense(size, false, nil)

		ur.Copy(r)
		lr.Copy(r)

		u := NewTriDense(size, true, nil)
		u.Copy(ur)
		if !equal(u, ur) {
			t.Fatal("unexpected result for U triangle copy of U triangle: not equal")
		}

		l := NewTriDense(size, false, nil)
		l.Copy(lr)
		if !equal(l, lr) {
			t.Fatal("unexpected result for L triangle copy of L triangle: not equal")
		}

		zero(u.mat.Data)
		u.Copy(lr)
		if !isDiagonal(u) {
			t.Fatal("unexpected result for U triangle copy of L triangle: off diagonal non-zero element")
		}
		if !equalDiagonal(u, lr) {
			t.Fatal("unexpected result for U triangle copy of L triangle: diagonal not equal")
		}

		zero(l.mat.Data)
		l.Copy(ur)
		if !isDiagonal(l) {
			t.Fatal("unexpected result for L triangle copy of U triangle: off diagonal non-zero element")
		}
		if !equalDiagonal(l, ur) {
			t.Fatal("unexpected result for L triangle copy of U triangle: diagonal not equal")
		}
	}
}

func TestTriInverse(t *testing.T) {
	for _, kind := range []matrix.TriKind{matrix.Upper, matrix.Lower} {
		for _, n := range []int{1, 3, 5, 9} {
			data := make([]float64, n*n)
			for i := range data {
				data[i] = rand.NormFloat64()
			}
			a := NewTriDense(n, kind, data)
			var tr TriDense
			err := tr.InverseTri(a)
			if err != nil {
				t.Errorf("Bad test: %s", err)
			}
			var d Dense
			d.Mul(a, &tr)
			if !equalApprox(eye(n), &d, 1e-8, false) {
				var diff Dense
				diff.Sub(eye(n), &d)
				t.Errorf("Tri times inverse is not identity. Norm of difference: %v", Norm(&diff, 2))
			}
		}
	}
}

func TestTriMul(t *testing.T) {
	method := func(receiver, a, b Matrix) {
		type MulTrier interface {
			MulTri(a, b Triangular)
		}
		receiver.(MulTrier).MulTri(a.(Triangular), b.(Triangular))
	}
	denseComparison := func(receiver, a, b *Dense) {
		receiver.Mul(a, b)
	}
	legalSizeTriMul := func(ar, ac, br, bc int) bool {
		// Need both to be square and the sizes to be the same
		return ar == ac && br == bc && ar == br
	}

	// The legal types are triangles with the same TriKind.
	// legalTypesTri returns whether both input arguments are Triangular.
	legalTypes := func(a, b Matrix) bool {
		at, ok := a.(Triangular)
		if !ok {
			return false
		}
		bt, ok := b.(Triangular)
		if !ok {
			return false
		}
		_, ak := at.Triangle()
		_, bk := bt.Triangle()
		return ak == bk
	}
	legalTypesLower := func(a, b Matrix) bool {
		legal := legalTypes(a, b)
		if !legal {
			return false
		}
		_, kind := a.(Triangular).Triangle()
		r := kind == matrix.Lower
		return r
	}
	receiver := NewTriDense(3, matrix.Lower, nil)
	testTwoInput(t, "TriMul", receiver, method, denseComparison, legalTypesLower, legalSizeTriMul, 1e-14)

	legalTypesUpper := func(a, b Matrix) bool {
		legal := legalTypes(a, b)
		if !legal {
			return false
		}
		_, kind := a.(Triangular).Triangle()
		r := kind == matrix.Upper
		return r
	}
	receiver = NewTriDense(3, matrix.Upper, nil)
	testTwoInput(t, "TriMul", receiver, method, denseComparison, legalTypesUpper, legalSizeTriMul, 1e-14)
}
