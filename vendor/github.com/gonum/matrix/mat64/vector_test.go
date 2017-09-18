package mat64

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/gonum/blas/blas64"
	"github.com/gonum/matrix"
)

func TestNewVector(t *testing.T) {
	for i, test := range []struct {
		n      int
		data   []float64
		vector *Vector
	}{
		{
			n:    3,
			data: []float64{4, 5, 6},
			vector: &Vector{
				mat: blas64.Vector{
					Data: []float64{4, 5, 6},
					Inc:  1,
				},
				n: 3,
			},
		},
		{
			n:    3,
			data: nil,
			vector: &Vector{
				mat: blas64.Vector{
					Data: []float64{0, 0, 0},
					Inc:  1,
				},
				n: 3,
			},
		},
	} {
		v := NewVector(test.n, test.data)
		rows, cols := v.Dims()
		if rows != test.n {
			t.Errorf("unexpected number of rows for test %d: got: %d want: %d", i, rows, test.n)
		}
		if cols != 1 {
			t.Errorf("unexpected number of cols for test %d: got: %d want: 1", i, cols)
		}
		if !reflect.DeepEqual(v, test.vector) {
			t.Errorf("unexpected data slice for test %d: got: %v want: %v", i, v, test.vector)
		}
	}
}

func TestVectorAtSet(t *testing.T) {
	for i, test := range []struct {
		vector *Vector
	}{
		{
			vector: &Vector{
				mat: blas64.Vector{
					Data: []float64{0, 1, 2},
					Inc:  1,
				},
				n: 3,
			},
		},
		{
			vector: &Vector{
				mat: blas64.Vector{
					Data: []float64{0, 10, 10, 1, 10, 10, 2},
					Inc:  3,
				},
				n: 3,
			},
		},
	} {
		v := test.vector
		n := test.vector.n

		for _, row := range []int{-1, n} {
			panicked, message := panics(func() { v.At(row, 0) })
			if !panicked || message != matrix.ErrRowAccess.Error() {
				t.Errorf("expected panic for invalid row access for test %d n=%d r=%d", i, n, row)
			}
		}
		for _, col := range []int{-1, 1} {
			panicked, message := panics(func() { v.At(0, col) })
			if !panicked || message != matrix.ErrColAccess.Error() {
				t.Errorf("expected panic for invalid column access for test %d n=%d c=%d", i, n, col)
			}
		}

		for _, row := range []int{0, 1, n - 1} {
			if e := v.At(row, 0); e != float64(row) {
				t.Errorf("unexpected value for At(%d, 0) for test %d : got: %v want: %v", row, i, e, float64(row))
			}
		}

		for _, row := range []int{-1, n} {
			panicked, message := panics(func() { v.SetVec(row, 100) })
			if !panicked || message != matrix.ErrVectorAccess.Error() {
				t.Errorf("expected panic for invalid row access for test %d n=%d r=%d", i, n, row)
			}
		}

		for inc, row := range []int{0, 2} {
			v.SetVec(row, 100+float64(inc))
			if e := v.At(row, 0); e != 100+float64(inc) {
				t.Errorf("unexpected value for At(%d, 0) after SetVec(%[1]d, %v) for test %d: got: %v want: %[2]v", row, 100+float64(inc), i, e)
			}
		}
	}
}

func TestVectorMul(t *testing.T) {
	method := func(receiver, a, b Matrix) {
		type mulVecer interface {
			MulVec(a Matrix, b *Vector)
		}
		rd := receiver.(mulVecer)
		rd.MulVec(a, b.(*Vector))
	}
	denseComparison := func(receiver, a, b *Dense) {
		receiver.Mul(a, b)
	}
	legalSizeMulVec := func(ar, ac, br, bc int) bool {
		var legal bool
		if bc != 1 {
			legal = false
		} else {
			legal = ac == br
		}
		return legal
	}
	testTwoInput(t, "MulVec", &Vector{}, method, denseComparison, legalTypesNotVecVec, legalSizeMulVec, 1e-14)
}

func TestVectorScale(t *testing.T) {
	for i, test := range []struct {
		a     *Vector
		alpha float64
		want  *Vector
	}{
		{
			a:     NewVector(3, []float64{0, 1, 2}),
			alpha: 0,
			want:  NewVector(3, []float64{0, 0, 0}),
		},
		{
			a:     NewVector(3, []float64{0, 1, 2}),
			alpha: 1,
			want:  NewVector(3, []float64{0, 1, 2}),
		},
		{
			a:     NewVector(3, []float64{0, 1, 2}),
			alpha: -2,
			want:  NewVector(3, []float64{0, -2, -4}),
		},
		{
			a:     NewDense(3, 1, []float64{0, 1, 2}).ColView(0),
			alpha: 0,
			want:  NewVector(3, []float64{0, 0, 0}),
		},
		{
			a:     NewDense(3, 1, []float64{0, 1, 2}).ColView(0),
			alpha: 1,
			want:  NewVector(3, []float64{0, 1, 2}),
		},
		{
			a:     NewDense(3, 1, []float64{0, 1, 2}).ColView(0),
			alpha: -2,
			want:  NewVector(3, []float64{0, -2, -4}),
		},
		{
			a: NewDense(3, 3, []float64{
				0, 1, 2,
				3, 4, 5,
				6, 7, 8,
			}).ColView(1),
			alpha: -2,
			want:  NewVector(3, []float64{-2, -8, -14}),
		},
	} {
		var v Vector
		v.ScaleVec(test.alpha, test.a)
		if !reflect.DeepEqual(v.RawVector(), test.want.RawVector()) {
			t.Errorf("test %d: unexpected result for v = alpha * a: got: %v want: %v", i, v.RawVector(), test.want.RawVector())
		}

		v.CopyVec(test.a)
		v.ScaleVec(test.alpha, &v)
		if !reflect.DeepEqual(v.RawVector(), test.want.RawVector()) {
			t.Errorf("test %d: unexpected result for v = alpha * v: got: %v want: %v", i, v.RawVector(), test.want.RawVector())
		}
	}

	for _, alpha := range []float64{0, 1, -1, 2.3, -2.3} {
		method := func(receiver, a Matrix) {
			type scaleVecer interface {
				ScaleVec(float64, *Vector)
			}
			v := receiver.(scaleVecer)
			v.ScaleVec(alpha, a.(*Vector))
		}
		denseComparison := func(receiver, a *Dense) {
			receiver.Scale(alpha, a)
		}
		testOneInput(t, "ScaleVec", &Vector{}, method, denseComparison, legalTypeVec, isAnyVector, 0)
	}
}

func TestVectorAddScaled(t *testing.T) {
	for _, alpha := range []float64{0, 1, -1, 2.3, -2.3} {
		method := func(receiver, a, b Matrix) {
			type addScaledVecer interface {
				AddScaledVec(*Vector, float64, *Vector)
			}
			v := receiver.(addScaledVecer)
			v.AddScaledVec(a.(*Vector), alpha, b.(*Vector))
		}
		denseComparison := func(receiver, a, b *Dense) {
			var sb Dense
			sb.Scale(alpha, b)
			receiver.Add(a, &sb)
		}
		testTwoInput(t, "AddScaledVec", &Vector{}, method, denseComparison, legalTypesVecVec, legalSizeSameVec, 1e-14)
	}
}

func TestVectorAdd(t *testing.T) {
	for i, test := range []struct {
		a, b *Vector
		want *Vector
	}{
		{
			a:    NewVector(3, []float64{0, 1, 2}),
			b:    NewVector(3, []float64{0, 2, 3}),
			want: NewVector(3, []float64{0, 3, 5}),
		},
		{
			a:    NewVector(3, []float64{0, 1, 2}),
			b:    NewDense(3, 1, []float64{0, 2, 3}).ColView(0),
			want: NewVector(3, []float64{0, 3, 5}),
		},
		{
			a:    NewDense(3, 1, []float64{0, 1, 2}).ColView(0),
			b:    NewDense(3, 1, []float64{0, 2, 3}).ColView(0),
			want: NewVector(3, []float64{0, 3, 5}),
		},
	} {
		var v Vector
		v.AddVec(test.a, test.b)
		if !reflect.DeepEqual(v.RawVector(), test.want.RawVector()) {
			t.Errorf("unexpected result for test %d: got: %v want: %v", i, v.RawVector(), test.want.RawVector())
		}
	}
}

func TestVectorSub(t *testing.T) {
	for i, test := range []struct {
		a, b *Vector
		want *Vector
	}{
		{
			a:    NewVector(3, []float64{0, 1, 2}),
			b:    NewVector(3, []float64{0, 0.5, 1}),
			want: NewVector(3, []float64{0, 0.5, 1}),
		},
		{
			a:    NewVector(3, []float64{0, 1, 2}),
			b:    NewDense(3, 1, []float64{0, 0.5, 1}).ColView(0),
			want: NewVector(3, []float64{0, 0.5, 1}),
		},
		{
			a:    NewDense(3, 1, []float64{0, 1, 2}).ColView(0),
			b:    NewDense(3, 1, []float64{0, 0.5, 1}).ColView(0),
			want: NewVector(3, []float64{0, 0.5, 1}),
		},
	} {
		var v Vector
		v.SubVec(test.a, test.b)
		if !reflect.DeepEqual(v.RawVector(), test.want.RawVector()) {
			t.Errorf("unexpected result for test %d: got: %v want: %v", i, v.RawVector(), test.want.RawVector())
		}
	}
}

func TestVectorMulElem(t *testing.T) {
	for i, test := range []struct {
		a, b *Vector
		want *Vector
	}{
		{
			a:    NewVector(3, []float64{0, 1, 2}),
			b:    NewVector(3, []float64{0, 2, 3}),
			want: NewVector(3, []float64{0, 2, 6}),
		},
		{
			a:    NewVector(3, []float64{0, 1, 2}),
			b:    NewDense(3, 1, []float64{0, 2, 3}).ColView(0),
			want: NewVector(3, []float64{0, 2, 6}),
		},
		{
			a:    NewDense(3, 1, []float64{0, 1, 2}).ColView(0),
			b:    NewDense(3, 1, []float64{0, 2, 3}).ColView(0),
			want: NewVector(3, []float64{0, 2, 6}),
		},
	} {
		var v Vector
		v.MulElemVec(test.a, test.b)
		if !reflect.DeepEqual(v.RawVector(), test.want.RawVector()) {
			t.Errorf("unexpected result for test %d: got: %v want: %v", i, v.RawVector(), test.want.RawVector())
		}
	}
}

func TestVectorDivElem(t *testing.T) {
	for i, test := range []struct {
		a, b *Vector
		want *Vector
	}{
		{
			a:    NewVector(3, []float64{0.5, 1, 2}),
			b:    NewVector(3, []float64{0.5, 0.5, 1}),
			want: NewVector(3, []float64{1, 2, 2}),
		},
		{
			a:    NewVector(3, []float64{0.5, 1, 2}),
			b:    NewDense(3, 1, []float64{0.5, 0.5, 1}).ColView(0),
			want: NewVector(3, []float64{1, 2, 2}),
		},
		{
			a:    NewDense(3, 1, []float64{0.5, 1, 2}).ColView(0),
			b:    NewDense(3, 1, []float64{0.5, 0.5, 1}).ColView(0),
			want: NewVector(3, []float64{1, 2, 2}),
		},
	} {
		var v Vector
		v.DivElemVec(test.a, test.b)
		if !reflect.DeepEqual(v.RawVector(), test.want.RawVector()) {
			t.Errorf("unexpected result for test %d: got: %v want: %v", i, v.RawVector(), test.want.RawVector())
		}
	}
}

func BenchmarkAddScaledVec10Inc1(b *testing.B)      { addScaledVecBench(b, 10, 1) }
func BenchmarkAddScaledVec100Inc1(b *testing.B)     { addScaledVecBench(b, 100, 1) }
func BenchmarkAddScaledVec1000Inc1(b *testing.B)    { addScaledVecBench(b, 1000, 1) }
func BenchmarkAddScaledVec10000Inc1(b *testing.B)   { addScaledVecBench(b, 10000, 1) }
func BenchmarkAddScaledVec100000Inc1(b *testing.B)  { addScaledVecBench(b, 100000, 1) }
func BenchmarkAddScaledVec10Inc2(b *testing.B)      { addScaledVecBench(b, 10, 2) }
func BenchmarkAddScaledVec100Inc2(b *testing.B)     { addScaledVecBench(b, 100, 2) }
func BenchmarkAddScaledVec1000Inc2(b *testing.B)    { addScaledVecBench(b, 1000, 2) }
func BenchmarkAddScaledVec10000Inc2(b *testing.B)   { addScaledVecBench(b, 10000, 2) }
func BenchmarkAddScaledVec100000Inc2(b *testing.B)  { addScaledVecBench(b, 100000, 2) }
func BenchmarkAddScaledVec10Inc20(b *testing.B)     { addScaledVecBench(b, 10, 20) }
func BenchmarkAddScaledVec100Inc20(b *testing.B)    { addScaledVecBench(b, 100, 20) }
func BenchmarkAddScaledVec1000Inc20(b *testing.B)   { addScaledVecBench(b, 1000, 20) }
func BenchmarkAddScaledVec10000Inc20(b *testing.B)  { addScaledVecBench(b, 10000, 20) }
func BenchmarkAddScaledVec100000Inc20(b *testing.B) { addScaledVecBench(b, 100000, 20) }
func addScaledVecBench(b *testing.B, size, inc int) {
	x := randVector(size, inc, 1, rand.NormFloat64)
	y := randVector(size, inc, 1, rand.NormFloat64)
	b.ResetTimer()
	var v Vector
	for i := 0; i < b.N; i++ {
		v.AddScaledVec(y, 2, x)
	}
}

func BenchmarkScaleVec10Inc1(b *testing.B)      { scaleVecBench(b, 10, 1) }
func BenchmarkScaleVec100Inc1(b *testing.B)     { scaleVecBench(b, 100, 1) }
func BenchmarkScaleVec1000Inc1(b *testing.B)    { scaleVecBench(b, 1000, 1) }
func BenchmarkScaleVec10000Inc1(b *testing.B)   { scaleVecBench(b, 10000, 1) }
func BenchmarkScaleVec100000Inc1(b *testing.B)  { scaleVecBench(b, 100000, 1) }
func BenchmarkScaleVec10Inc2(b *testing.B)      { scaleVecBench(b, 10, 2) }
func BenchmarkScaleVec100Inc2(b *testing.B)     { scaleVecBench(b, 100, 2) }
func BenchmarkScaleVec1000Inc2(b *testing.B)    { scaleVecBench(b, 1000, 2) }
func BenchmarkScaleVec10000Inc2(b *testing.B)   { scaleVecBench(b, 10000, 2) }
func BenchmarkScaleVec100000Inc2(b *testing.B)  { scaleVecBench(b, 100000, 2) }
func BenchmarkScaleVec10Inc20(b *testing.B)     { scaleVecBench(b, 10, 20) }
func BenchmarkScaleVec100Inc20(b *testing.B)    { scaleVecBench(b, 100, 20) }
func BenchmarkScaleVec1000Inc20(b *testing.B)   { scaleVecBench(b, 1000, 20) }
func BenchmarkScaleVec10000Inc20(b *testing.B)  { scaleVecBench(b, 10000, 20) }
func BenchmarkScaleVec100000Inc20(b *testing.B) { scaleVecBench(b, 100000, 20) }
func scaleVecBench(b *testing.B, size, inc int) {
	x := randVector(size, inc, 1, rand.NormFloat64)
	b.ResetTimer()
	var v Vector
	for i := 0; i < b.N; i++ {
		v.ScaleVec(2, x)
	}
}

func BenchmarkAddVec10Inc1(b *testing.B)      { addVecBench(b, 10, 1) }
func BenchmarkAddVec100Inc1(b *testing.B)     { addVecBench(b, 100, 1) }
func BenchmarkAddVec1000Inc1(b *testing.B)    { addVecBench(b, 1000, 1) }
func BenchmarkAddVec10000Inc1(b *testing.B)   { addVecBench(b, 10000, 1) }
func BenchmarkAddVec100000Inc1(b *testing.B)  { addVecBench(b, 100000, 1) }
func BenchmarkAddVec10Inc2(b *testing.B)      { addVecBench(b, 10, 2) }
func BenchmarkAddVec100Inc2(b *testing.B)     { addVecBench(b, 100, 2) }
func BenchmarkAddVec1000Inc2(b *testing.B)    { addVecBench(b, 1000, 2) }
func BenchmarkAddVec10000Inc2(b *testing.B)   { addVecBench(b, 10000, 2) }
func BenchmarkAddVec100000Inc2(b *testing.B)  { addVecBench(b, 100000, 2) }
func BenchmarkAddVec10Inc20(b *testing.B)     { addVecBench(b, 10, 20) }
func BenchmarkAddVec100Inc20(b *testing.B)    { addVecBench(b, 100, 20) }
func BenchmarkAddVec1000Inc20(b *testing.B)   { addVecBench(b, 1000, 20) }
func BenchmarkAddVec10000Inc20(b *testing.B)  { addVecBench(b, 10000, 20) }
func BenchmarkAddVec100000Inc20(b *testing.B) { addVecBench(b, 100000, 20) }
func addVecBench(b *testing.B, size, inc int) {
	x := randVector(size, inc, 1, rand.NormFloat64)
	y := randVector(size, inc, 1, rand.NormFloat64)
	b.ResetTimer()
	var v Vector
	for i := 0; i < b.N; i++ {
		v.AddVec(x, y)
	}
}

func BenchmarkSubVec10Inc1(b *testing.B)      { subVecBench(b, 10, 1) }
func BenchmarkSubVec100Inc1(b *testing.B)     { subVecBench(b, 100, 1) }
func BenchmarkSubVec1000Inc1(b *testing.B)    { subVecBench(b, 1000, 1) }
func BenchmarkSubVec10000Inc1(b *testing.B)   { subVecBench(b, 10000, 1) }
func BenchmarkSubVec100000Inc1(b *testing.B)  { subVecBench(b, 100000, 1) }
func BenchmarkSubVec10Inc2(b *testing.B)      { subVecBench(b, 10, 2) }
func BenchmarkSubVec100Inc2(b *testing.B)     { subVecBench(b, 100, 2) }
func BenchmarkSubVec1000Inc2(b *testing.B)    { subVecBench(b, 1000, 2) }
func BenchmarkSubVec10000Inc2(b *testing.B)   { subVecBench(b, 10000, 2) }
func BenchmarkSubVec100000Inc2(b *testing.B)  { subVecBench(b, 100000, 2) }
func BenchmarkSubVec10Inc20(b *testing.B)     { subVecBench(b, 10, 20) }
func BenchmarkSubVec100Inc20(b *testing.B)    { subVecBench(b, 100, 20) }
func BenchmarkSubVec1000Inc20(b *testing.B)   { subVecBench(b, 1000, 20) }
func BenchmarkSubVec10000Inc20(b *testing.B)  { subVecBench(b, 10000, 20) }
func BenchmarkSubVec100000Inc20(b *testing.B) { subVecBench(b, 100000, 20) }
func subVecBench(b *testing.B, size, inc int) {
	x := randVector(size, inc, 1, rand.NormFloat64)
	y := randVector(size, inc, 1, rand.NormFloat64)
	b.ResetTimer()
	var v Vector
	for i := 0; i < b.N; i++ {
		v.SubVec(x, y)
	}
}

func randVector(size, inc int, rho float64, rnd func() float64) *Vector {
	if size <= 0 {
		panic("bad vector size")
	}
	data := make([]float64, size*inc)
	for i := range data {
		if rand.Float64() < rho {
			data[i] = rnd()
		}
	}
	return &Vector{
		mat: blas64.Vector{
			Inc:  inc,
			Data: data,
		},
		n: size,
	}
}
