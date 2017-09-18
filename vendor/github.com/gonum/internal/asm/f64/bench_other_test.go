// Copyright ©2016 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package f64

import (
	"math"
	"testing"
)

func benchL1Norm(f func(x []float64) float64, sz int, t *testing.B) {
	dst := y[:sz]
	for i := 0; i < t.N; i++ {
		f(dst)
	}
}

var naiveL1Norm = func(x []float64) (sum float64) {
	for _, v := range x {
		sum += math.Abs(v)
	}
	return sum
}

func BenchmarkL1Norm1(t *testing.B)      { benchL1Norm(L1Norm, 1, t) }
func BenchmarkL1Norm2(t *testing.B)      { benchL1Norm(L1Norm, 2, t) }
func BenchmarkL1Norm3(t *testing.B)      { benchL1Norm(L1Norm, 3, t) }
func BenchmarkL1Norm4(t *testing.B)      { benchL1Norm(L1Norm, 4, t) }
func BenchmarkL1Norm5(t *testing.B)      { benchL1Norm(L1Norm, 5, t) }
func BenchmarkL1Norm10(t *testing.B)     { benchL1Norm(L1Norm, 10, t) }
func BenchmarkL1Norm100(t *testing.B)    { benchL1Norm(L1Norm, 100, t) }
func BenchmarkL1Norm1000(t *testing.B)   { benchL1Norm(L1Norm, 1000, t) }
func BenchmarkL1Norm10000(t *testing.B)  { benchL1Norm(L1Norm, 10000, t) }
func BenchmarkL1Norm100000(t *testing.B) { benchL1Norm(L1Norm, 100000, t) }
func BenchmarkL1Norm500000(t *testing.B) { benchL1Norm(L1Norm, 500000, t) }

func BenchmarkLL1Norm1(t *testing.B)      { benchL1Norm(naiveL1Norm, 1, t) }
func BenchmarkLL1Norm2(t *testing.B)      { benchL1Norm(naiveL1Norm, 2, t) }
func BenchmarkLL1Norm3(t *testing.B)      { benchL1Norm(naiveL1Norm, 3, t) }
func BenchmarkLL1Norm4(t *testing.B)      { benchL1Norm(naiveL1Norm, 4, t) }
func BenchmarkLL1Norm5(t *testing.B)      { benchL1Norm(naiveL1Norm, 5, t) }
func BenchmarkLL1Norm10(t *testing.B)     { benchL1Norm(naiveL1Norm, 10, t) }
func BenchmarkLL1Norm100(t *testing.B)    { benchL1Norm(naiveL1Norm, 100, t) }
func BenchmarkLL1Norm1000(t *testing.B)   { benchL1Norm(naiveL1Norm, 1000, t) }
func BenchmarkLL1Norm10000(t *testing.B)  { benchL1Norm(naiveL1Norm, 10000, t) }
func BenchmarkLL1Norm100000(t *testing.B) { benchL1Norm(naiveL1Norm, 100000, t) }
func BenchmarkLL1Norm500000(t *testing.B) { benchL1Norm(naiveL1Norm, 500000, t) }

func benchL1NormInc(t *testing.B, ln, inc int, f func(x []float64, n, incX int) float64) {
	for i := 0; i < t.N; i++ {
		f(x, ln, inc)
	}
}

var naiveL1NormInc = func(x []float64, n, incX int) (sum float64) {
	for i := 0; i < n*incX; i += incX {
		sum += math.Abs(x[i])
	}
	return sum
}

func BenchmarkF64L1NormIncN1Inc1(b *testing.B) { benchL1NormInc(b, 1, 1, L1NormInc) }

func BenchmarkF64L1NormIncN2Inc1(b *testing.B)  { benchL1NormInc(b, 2, 1, L1NormInc) }
func BenchmarkF64L1NormIncN2Inc2(b *testing.B)  { benchL1NormInc(b, 2, 2, L1NormInc) }
func BenchmarkF64L1NormIncN2Inc4(b *testing.B)  { benchL1NormInc(b, 2, 4, L1NormInc) }
func BenchmarkF64L1NormIncN2Inc10(b *testing.B) { benchL1NormInc(b, 2, 10, L1NormInc) }

func BenchmarkF64L1NormIncN3Inc1(b *testing.B)  { benchL1NormInc(b, 3, 1, L1NormInc) }
func BenchmarkF64L1NormIncN3Inc2(b *testing.B)  { benchL1NormInc(b, 3, 2, L1NormInc) }
func BenchmarkF64L1NormIncN3Inc4(b *testing.B)  { benchL1NormInc(b, 3, 4, L1NormInc) }
func BenchmarkF64L1NormIncN3Inc10(b *testing.B) { benchL1NormInc(b, 3, 10, L1NormInc) }

func BenchmarkF64L1NormIncN4Inc1(b *testing.B)  { benchL1NormInc(b, 4, 1, L1NormInc) }
func BenchmarkF64L1NormIncN4Inc2(b *testing.B)  { benchL1NormInc(b, 4, 2, L1NormInc) }
func BenchmarkF64L1NormIncN4Inc4(b *testing.B)  { benchL1NormInc(b, 4, 4, L1NormInc) }
func BenchmarkF64L1NormIncN4Inc10(b *testing.B) { benchL1NormInc(b, 4, 10, L1NormInc) }

func BenchmarkF64L1NormIncN10Inc1(b *testing.B)  { benchL1NormInc(b, 10, 1, L1NormInc) }
func BenchmarkF64L1NormIncN10Inc2(b *testing.B)  { benchL1NormInc(b, 10, 2, L1NormInc) }
func BenchmarkF64L1NormIncN10Inc4(b *testing.B)  { benchL1NormInc(b, 10, 4, L1NormInc) }
func BenchmarkF64L1NormIncN10Inc10(b *testing.B) { benchL1NormInc(b, 10, 10, L1NormInc) }

func BenchmarkF64L1NormIncN1000Inc1(b *testing.B)  { benchL1NormInc(b, 1000, 1, L1NormInc) }
func BenchmarkF64L1NormIncN1000Inc2(b *testing.B)  { benchL1NormInc(b, 1000, 2, L1NormInc) }
func BenchmarkF64L1NormIncN1000Inc4(b *testing.B)  { benchL1NormInc(b, 1000, 4, L1NormInc) }
func BenchmarkF64L1NormIncN1000Inc10(b *testing.B) { benchL1NormInc(b, 1000, 10, L1NormInc) }

func BenchmarkF64L1NormIncN100000Inc1(b *testing.B)  { benchL1NormInc(b, 100000, 1, L1NormInc) }
func BenchmarkF64L1NormIncN100000Inc2(b *testing.B)  { benchL1NormInc(b, 100000, 2, L1NormInc) }
func BenchmarkF64L1NormIncN100000Inc4(b *testing.B)  { benchL1NormInc(b, 100000, 4, L1NormInc) }
func BenchmarkF64L1NormIncN100000Inc10(b *testing.B) { benchL1NormInc(b, 100000, 10, L1NormInc) }

func BenchmarkLF64L1NormIncN1Inc1(b *testing.B) { benchL1NormInc(b, 1, 1, naiveL1NormInc) }

func BenchmarkLF64L1NormIncN2Inc1(b *testing.B)  { benchL1NormInc(b, 2, 1, naiveL1NormInc) }
func BenchmarkLF64L1NormIncN2Inc2(b *testing.B)  { benchL1NormInc(b, 2, 2, naiveL1NormInc) }
func BenchmarkLF64L1NormIncN2Inc4(b *testing.B)  { benchL1NormInc(b, 2, 4, naiveL1NormInc) }
func BenchmarkLF64L1NormIncN2Inc10(b *testing.B) { benchL1NormInc(b, 2, 10, naiveL1NormInc) }

func BenchmarkLF64L1NormIncN3Inc1(b *testing.B)  { benchL1NormInc(b, 3, 1, naiveL1NormInc) }
func BenchmarkLF64L1NormIncN3Inc2(b *testing.B)  { benchL1NormInc(b, 3, 2, naiveL1NormInc) }
func BenchmarkLF64L1NormIncN3Inc4(b *testing.B)  { benchL1NormInc(b, 3, 4, naiveL1NormInc) }
func BenchmarkLF64L1NormIncN3Inc10(b *testing.B) { benchL1NormInc(b, 3, 10, naiveL1NormInc) }

func BenchmarkLF64L1NormIncN4Inc1(b *testing.B)  { benchL1NormInc(b, 4, 1, naiveL1NormInc) }
func BenchmarkLF64L1NormIncN4Inc2(b *testing.B)  { benchL1NormInc(b, 4, 2, naiveL1NormInc) }
func BenchmarkLF64L1NormIncN4Inc4(b *testing.B)  { benchL1NormInc(b, 4, 4, naiveL1NormInc) }
func BenchmarkLF64L1NormIncN4Inc10(b *testing.B) { benchL1NormInc(b, 4, 10, naiveL1NormInc) }

func BenchmarkLF64L1NormIncN10Inc1(b *testing.B)  { benchL1NormInc(b, 10, 1, naiveL1NormInc) }
func BenchmarkLF64L1NormIncN10Inc2(b *testing.B)  { benchL1NormInc(b, 10, 2, naiveL1NormInc) }
func BenchmarkLF64L1NormIncN10Inc4(b *testing.B)  { benchL1NormInc(b, 10, 4, naiveL1NormInc) }
func BenchmarkLF64L1NormIncN10Inc10(b *testing.B) { benchL1NormInc(b, 10, 10, naiveL1NormInc) }

func BenchmarkLF64L1NormIncN1000Inc1(b *testing.B)  { benchL1NormInc(b, 1000, 1, naiveL1NormInc) }
func BenchmarkLF64L1NormIncN1000Inc2(b *testing.B)  { benchL1NormInc(b, 1000, 2, naiveL1NormInc) }
func BenchmarkLF64L1NormIncN1000Inc4(b *testing.B)  { benchL1NormInc(b, 1000, 4, naiveL1NormInc) }
func BenchmarkLF64L1NormIncN1000Inc10(b *testing.B) { benchL1NormInc(b, 1000, 10, naiveL1NormInc) }

func BenchmarkLF64L1NormIncN100000Inc1(b *testing.B)  { benchL1NormInc(b, 100000, 1, naiveL1NormInc) }
func BenchmarkLF64L1NormIncN100000Inc2(b *testing.B)  { benchL1NormInc(b, 100000, 2, naiveL1NormInc) }
func BenchmarkLF64L1NormIncN100000Inc4(b *testing.B)  { benchL1NormInc(b, 100000, 4, naiveL1NormInc) }
func BenchmarkLF64L1NormIncN100000Inc10(b *testing.B) { benchL1NormInc(b, 100000, 10, naiveL1NormInc) }

func benchAdd(f func(dst, s []float64), sz int, t *testing.B) {
	dst, s := y[:sz], x[:sz]
	for i := 0; i < t.N; i++ {
		f(dst, s)
	}
}

var naiveAdd = func(dst, s []float64) {
	for i, v := range s {
		dst[i] += v
	}
}

func BenchmarkAdd1(t *testing.B)      { benchAdd(Add, 1, t) }
func BenchmarkAdd2(t *testing.B)      { benchAdd(Add, 2, t) }
func BenchmarkAdd3(t *testing.B)      { benchAdd(Add, 3, t) }
func BenchmarkAdd4(t *testing.B)      { benchAdd(Add, 4, t) }
func BenchmarkAdd5(t *testing.B)      { benchAdd(Add, 5, t) }
func BenchmarkAdd10(t *testing.B)     { benchAdd(Add, 10, t) }
func BenchmarkAdd100(t *testing.B)    { benchAdd(Add, 100, t) }
func BenchmarkAdd1000(t *testing.B)   { benchAdd(Add, 1000, t) }
func BenchmarkAdd10000(t *testing.B)  { benchAdd(Add, 10000, t) }
func BenchmarkAdd100000(t *testing.B) { benchAdd(Add, 100000, t) }
func BenchmarkAdd500000(t *testing.B) { benchAdd(Add, 500000, t) }

func BenchmarkLAdd1(t *testing.B)      { benchAdd(naiveAdd, 1, t) }
func BenchmarkLAdd2(t *testing.B)      { benchAdd(naiveAdd, 2, t) }
func BenchmarkLAdd3(t *testing.B)      { benchAdd(naiveAdd, 3, t) }
func BenchmarkLAdd4(t *testing.B)      { benchAdd(naiveAdd, 4, t) }
func BenchmarkLAdd5(t *testing.B)      { benchAdd(naiveAdd, 5, t) }
func BenchmarkLAdd10(t *testing.B)     { benchAdd(naiveAdd, 10, t) }
func BenchmarkLAdd100(t *testing.B)    { benchAdd(naiveAdd, 100, t) }
func BenchmarkLAdd1000(t *testing.B)   { benchAdd(naiveAdd, 1000, t) }
func BenchmarkLAdd10000(t *testing.B)  { benchAdd(naiveAdd, 10000, t) }
func BenchmarkLAdd100000(t *testing.B) { benchAdd(naiveAdd, 100000, t) }
func BenchmarkLAdd500000(t *testing.B) { benchAdd(naiveAdd, 500000, t) }

func benchAddConst(f func(a float64, x []float64), sz int, t *testing.B) {
	a, x := 1., x[:sz]
	for i := 0; i < t.N; i++ {
		f(a, x)
	}
}

var naiveAddConst = func(a float64, x []float64) {
	for i := range x {
		x[i] += a
	}
}

func BenchmarkAddConst1(t *testing.B)      { benchAddConst(AddConst, 1, t) }
func BenchmarkAddConst2(t *testing.B)      { benchAddConst(AddConst, 2, t) }
func BenchmarkAddConst3(t *testing.B)      { benchAddConst(AddConst, 3, t) }
func BenchmarkAddConst4(t *testing.B)      { benchAddConst(AddConst, 4, t) }
func BenchmarkAddConst5(t *testing.B)      { benchAddConst(AddConst, 5, t) }
func BenchmarkAddConst10(t *testing.B)     { benchAddConst(AddConst, 10, t) }
func BenchmarkAddConst100(t *testing.B)    { benchAddConst(AddConst, 100, t) }
func BenchmarkAddConst1000(t *testing.B)   { benchAddConst(AddConst, 1000, t) }
func BenchmarkAddConst10000(t *testing.B)  { benchAddConst(AddConst, 10000, t) }
func BenchmarkAddConst100000(t *testing.B) { benchAddConst(AddConst, 100000, t) }
func BenchmarkAddConst500000(t *testing.B) { benchAddConst(AddConst, 500000, t) }

func BenchmarkLAddConst1(t *testing.B)      { benchAddConst(naiveAddConst, 1, t) }
func BenchmarkLAddConst2(t *testing.B)      { benchAddConst(naiveAddConst, 2, t) }
func BenchmarkLAddConst3(t *testing.B)      { benchAddConst(naiveAddConst, 3, t) }
func BenchmarkLAddConst4(t *testing.B)      { benchAddConst(naiveAddConst, 4, t) }
func BenchmarkLAddConst5(t *testing.B)      { benchAddConst(naiveAddConst, 5, t) }
func BenchmarkLAddConst10(t *testing.B)     { benchAddConst(naiveAddConst, 10, t) }
func BenchmarkLAddConst100(t *testing.B)    { benchAddConst(naiveAddConst, 100, t) }
func BenchmarkLAddConst1000(t *testing.B)   { benchAddConst(naiveAddConst, 1000, t) }
func BenchmarkLAddConst10000(t *testing.B)  { benchAddConst(naiveAddConst, 10000, t) }
func BenchmarkLAddConst100000(t *testing.B) { benchAddConst(naiveAddConst, 100000, t) }
func BenchmarkLAddConst500000(t *testing.B) { benchAddConst(naiveAddConst, 500000, t) }

func benchCumSum(f func(a, b []float64) []float64, sz int, t *testing.B) {
	a, b := x[:sz], y[:sz]
	for i := 0; i < t.N; i++ {
		f(a, b)
	}
}

var naiveCumSum = func(dst, s []float64) []float64 {
	if len(s) == 0 {
		return dst
	}
	dst[0] = s[0]
	for i, v := range s[1:] {
		dst[i+1] = dst[i] + v
	}
	return dst
}

func BenchmarkCumSum1(t *testing.B)      { benchCumSum(CumSum, 1, t) }
func BenchmarkCumSum2(t *testing.B)      { benchCumSum(CumSum, 2, t) }
func BenchmarkCumSum3(t *testing.B)      { benchCumSum(CumSum, 3, t) }
func BenchmarkCumSum4(t *testing.B)      { benchCumSum(CumSum, 4, t) }
func BenchmarkCumSum5(t *testing.B)      { benchCumSum(CumSum, 5, t) }
func BenchmarkCumSum10(t *testing.B)     { benchCumSum(CumSum, 10, t) }
func BenchmarkCumSum100(t *testing.B)    { benchCumSum(CumSum, 100, t) }
func BenchmarkCumSum1000(t *testing.B)   { benchCumSum(CumSum, 1000, t) }
func BenchmarkCumSum10000(t *testing.B)  { benchCumSum(CumSum, 10000, t) }
func BenchmarkCumSum100000(t *testing.B) { benchCumSum(CumSum, 100000, t) }
func BenchmarkCumSum500000(t *testing.B) { benchCumSum(CumSum, 500000, t) }

func BenchmarkLCumSum1(t *testing.B)      { benchCumSum(naiveCumSum, 1, t) }
func BenchmarkLCumSum2(t *testing.B)      { benchCumSum(naiveCumSum, 2, t) }
func BenchmarkLCumSum3(t *testing.B)      { benchCumSum(naiveCumSum, 3, t) }
func BenchmarkLCumSum4(t *testing.B)      { benchCumSum(naiveCumSum, 4, t) }
func BenchmarkLCumSum5(t *testing.B)      { benchCumSum(naiveCumSum, 5, t) }
func BenchmarkLCumSum10(t *testing.B)     { benchCumSum(naiveCumSum, 10, t) }
func BenchmarkLCumSum100(t *testing.B)    { benchCumSum(naiveCumSum, 100, t) }
func BenchmarkLCumSum1000(t *testing.B)   { benchCumSum(naiveCumSum, 1000, t) }
func BenchmarkLCumSum10000(t *testing.B)  { benchCumSum(naiveCumSum, 10000, t) }
func BenchmarkLCumSum100000(t *testing.B) { benchCumSum(naiveCumSum, 100000, t) }
func BenchmarkLCumSum500000(t *testing.B) { benchCumSum(naiveCumSum, 500000, t) }

func benchCumProd(f func(a, b []float64) []float64, sz int, t *testing.B) {
	a, b := x[:sz], y[:sz]
	for i := 0; i < t.N; i++ {
		f(a, b)
	}
}

var naiveCumProd = func(dst, s []float64) []float64 {
	if len(s) == 0 {
		return dst
	}
	dst[0] = s[0]
	for i, v := range s[1:] {
		dst[i+1] = dst[i] + v
	}
	return dst
}

func BenchmarkCumProd1(t *testing.B)      { benchCumProd(CumProd, 1, t) }
func BenchmarkCumProd2(t *testing.B)      { benchCumProd(CumProd, 2, t) }
func BenchmarkCumProd3(t *testing.B)      { benchCumProd(CumProd, 3, t) }
func BenchmarkCumProd4(t *testing.B)      { benchCumProd(CumProd, 4, t) }
func BenchmarkCumProd5(t *testing.B)      { benchCumProd(CumProd, 5, t) }
func BenchmarkCumProd10(t *testing.B)     { benchCumProd(CumProd, 10, t) }
func BenchmarkCumProd100(t *testing.B)    { benchCumProd(CumProd, 100, t) }
func BenchmarkCumProd1000(t *testing.B)   { benchCumProd(CumProd, 1000, t) }
func BenchmarkCumProd10000(t *testing.B)  { benchCumProd(CumProd, 10000, t) }
func BenchmarkCumProd100000(t *testing.B) { benchCumProd(CumProd, 100000, t) }
func BenchmarkCumProd500000(t *testing.B) { benchCumProd(CumProd, 500000, t) }

func BenchmarkLCumProd1(t *testing.B)      { benchCumProd(naiveCumProd, 1, t) }
func BenchmarkLCumProd2(t *testing.B)      { benchCumProd(naiveCumProd, 2, t) }
func BenchmarkLCumProd3(t *testing.B)      { benchCumProd(naiveCumProd, 3, t) }
func BenchmarkLCumProd4(t *testing.B)      { benchCumProd(naiveCumProd, 4, t) }
func BenchmarkLCumProd5(t *testing.B)      { benchCumProd(naiveCumProd, 5, t) }
func BenchmarkLCumProd10(t *testing.B)     { benchCumProd(naiveCumProd, 10, t) }
func BenchmarkLCumProd100(t *testing.B)    { benchCumProd(naiveCumProd, 100, t) }
func BenchmarkLCumProd1000(t *testing.B)   { benchCumProd(naiveCumProd, 1000, t) }
func BenchmarkLCumProd10000(t *testing.B)  { benchCumProd(naiveCumProd, 10000, t) }
func BenchmarkLCumProd100000(t *testing.B) { benchCumProd(naiveCumProd, 100000, t) }
func BenchmarkLCumProd500000(t *testing.B) { benchCumProd(naiveCumProd, 500000, t) }

func benchDiv(f func(a, b []float64), sz int, t *testing.B) {
	a, b := x[:sz], y[:sz]
	for i := 0; i < t.N; i++ {
		f(a, b)
	}
}

var naiveDiv = func(a, b []float64) {
	for i, v := range b {
		a[i] /= v
	}
}

func BenchmarkDiv1(t *testing.B)      { benchDiv(Div, 1, t) }
func BenchmarkDiv2(t *testing.B)      { benchDiv(Div, 2, t) }
func BenchmarkDiv3(t *testing.B)      { benchDiv(Div, 3, t) }
func BenchmarkDiv4(t *testing.B)      { benchDiv(Div, 4, t) }
func BenchmarkDiv5(t *testing.B)      { benchDiv(Div, 5, t) }
func BenchmarkDiv10(t *testing.B)     { benchDiv(Div, 10, t) }
func BenchmarkDiv100(t *testing.B)    { benchDiv(Div, 100, t) }
func BenchmarkDiv1000(t *testing.B)   { benchDiv(Div, 1000, t) }
func BenchmarkDiv10000(t *testing.B)  { benchDiv(Div, 10000, t) }
func BenchmarkDiv100000(t *testing.B) { benchDiv(Div, 100000, t) }
func BenchmarkDiv500000(t *testing.B) { benchDiv(Div, 500000, t) }

func BenchmarkLDiv1(t *testing.B)      { benchDiv(naiveDiv, 1, t) }
func BenchmarkLDiv2(t *testing.B)      { benchDiv(naiveDiv, 2, t) }
func BenchmarkLDiv3(t *testing.B)      { benchDiv(naiveDiv, 3, t) }
func BenchmarkLDiv4(t *testing.B)      { benchDiv(naiveDiv, 4, t) }
func BenchmarkLDiv5(t *testing.B)      { benchDiv(naiveDiv, 5, t) }
func BenchmarkLDiv10(t *testing.B)     { benchDiv(naiveDiv, 10, t) }
func BenchmarkLDiv100(t *testing.B)    { benchDiv(naiveDiv, 100, t) }
func BenchmarkLDiv1000(t *testing.B)   { benchDiv(naiveDiv, 1000, t) }
func BenchmarkLDiv10000(t *testing.B)  { benchDiv(naiveDiv, 10000, t) }
func BenchmarkLDiv100000(t *testing.B) { benchDiv(naiveDiv, 100000, t) }
func BenchmarkLDiv500000(t *testing.B) { benchDiv(naiveDiv, 500000, t) }

func benchDivTo(f func(dst, a, b []float64) []float64, sz int, t *testing.B) {
	dst, a, b := z[:sz], x[:sz], y[:sz]
	for i := 0; i < t.N; i++ {
		f(dst, a, b)
	}
}

var naiveDivTo = func(dst, s, t []float64) []float64 {
	for i, v := range s {
		dst[i] = v / t[i]
	}
	return dst
}

func BenchmarkDivTo1(t *testing.B)      { benchDivTo(DivTo, 1, t) }
func BenchmarkDivTo2(t *testing.B)      { benchDivTo(DivTo, 2, t) }
func BenchmarkDivTo3(t *testing.B)      { benchDivTo(DivTo, 3, t) }
func BenchmarkDivTo4(t *testing.B)      { benchDivTo(DivTo, 4, t) }
func BenchmarkDivTo5(t *testing.B)      { benchDivTo(DivTo, 5, t) }
func BenchmarkDivTo10(t *testing.B)     { benchDivTo(DivTo, 10, t) }
func BenchmarkDivTo100(t *testing.B)    { benchDivTo(DivTo, 100, t) }
func BenchmarkDivTo1000(t *testing.B)   { benchDivTo(DivTo, 1000, t) }
func BenchmarkDivTo10000(t *testing.B)  { benchDivTo(DivTo, 10000, t) }
func BenchmarkDivTo100000(t *testing.B) { benchDivTo(DivTo, 100000, t) }
func BenchmarkDivTo500000(t *testing.B) { benchDivTo(DivTo, 500000, t) }

func BenchmarkLDivTo1(t *testing.B)      { benchDivTo(naiveDivTo, 1, t) }
func BenchmarkLDivTo2(t *testing.B)      { benchDivTo(naiveDivTo, 2, t) }
func BenchmarkLDivTo3(t *testing.B)      { benchDivTo(naiveDivTo, 3, t) }
func BenchmarkLDivTo4(t *testing.B)      { benchDivTo(naiveDivTo, 4, t) }
func BenchmarkLDivTo5(t *testing.B)      { benchDivTo(naiveDivTo, 5, t) }
func BenchmarkLDivTo10(t *testing.B)     { benchDivTo(naiveDivTo, 10, t) }
func BenchmarkLDivTo100(t *testing.B)    { benchDivTo(naiveDivTo, 100, t) }
func BenchmarkLDivTo1000(t *testing.B)   { benchDivTo(naiveDivTo, 1000, t) }
func BenchmarkLDivTo10000(t *testing.B)  { benchDivTo(naiveDivTo, 10000, t) }
func BenchmarkLDivTo100000(t *testing.B) { benchDivTo(naiveDivTo, 100000, t) }
func BenchmarkLDivTo500000(t *testing.B) { benchDivTo(naiveDivTo, 500000, t) }

func benchL1Dist(f func(a, b []float64) float64, sz int, t *testing.B) {
	a, b := x[:sz], y[:sz]
	for i := 0; i < t.N; i++ {
		f(a, b)
	}
}

var naiveL1Dist = func(s, t []float64) float64 {
	var norm float64
	for i, v := range s {
		norm += math.Abs(t[i] - v)
	}
	return norm
}

func BenchmarkL1Dist1(t *testing.B)      { benchL1Dist(L1Dist, 1, t) }
func BenchmarkL1Dist2(t *testing.B)      { benchL1Dist(L1Dist, 2, t) }
func BenchmarkL1Dist3(t *testing.B)      { benchL1Dist(L1Dist, 3, t) }
func BenchmarkL1Dist4(t *testing.B)      { benchL1Dist(L1Dist, 4, t) }
func BenchmarkL1Dist5(t *testing.B)      { benchL1Dist(L1Dist, 5, t) }
func BenchmarkL1Dist10(t *testing.B)     { benchL1Dist(L1Dist, 10, t) }
func BenchmarkL1Dist100(t *testing.B)    { benchL1Dist(L1Dist, 100, t) }
func BenchmarkL1Dist1000(t *testing.B)   { benchL1Dist(L1Dist, 1000, t) }
func BenchmarkL1Dist10000(t *testing.B)  { benchL1Dist(L1Dist, 10000, t) }
func BenchmarkL1Dist100000(t *testing.B) { benchL1Dist(L1Dist, 100000, t) }
func BenchmarkL1Dist500000(t *testing.B) { benchL1Dist(L1Dist, 500000, t) }

func BenchmarkLL1Dist1(t *testing.B)      { benchL1Dist(naiveL1Dist, 1, t) }
func BenchmarkLL1Dist2(t *testing.B)      { benchL1Dist(naiveL1Dist, 2, t) }
func BenchmarkLL1Dist3(t *testing.B)      { benchL1Dist(naiveL1Dist, 3, t) }
func BenchmarkLL1Dist4(t *testing.B)      { benchL1Dist(naiveL1Dist, 4, t) }
func BenchmarkLL1Dist5(t *testing.B)      { benchL1Dist(naiveL1Dist, 5, t) }
func BenchmarkLL1Dist10(t *testing.B)     { benchL1Dist(naiveL1Dist, 10, t) }
func BenchmarkLL1Dist100(t *testing.B)    { benchL1Dist(naiveL1Dist, 100, t) }
func BenchmarkLL1Dist1000(t *testing.B)   { benchL1Dist(naiveL1Dist, 1000, t) }
func BenchmarkLL1Dist10000(t *testing.B)  { benchL1Dist(naiveL1Dist, 10000, t) }
func BenchmarkLL1Dist100000(t *testing.B) { benchL1Dist(naiveL1Dist, 100000, t) }
func BenchmarkLL1Dist500000(t *testing.B) { benchL1Dist(naiveL1Dist, 500000, t) }

func benchLinfDist(f func(a, b []float64) float64, sz int, t *testing.B) {
	a, b := x[:sz], y[:sz]
	for i := 0; i < t.N; i++ {
		f(a, b)
	}
}

var naiveLinfDist = func(s, t []float64) float64 {
	var norm float64
	if len(s) == 0 {
		return 0
	}
	norm = math.Abs(t[0] - s[0])
	for i, v := range s[1:] {
		absDiff := math.Abs(t[i+1] - v)
		if absDiff > norm || math.IsNaN(norm) {
			norm = absDiff
		}
	}
	return norm
}

func BenchmarkLinfDist1(t *testing.B)      { benchLinfDist(LinfDist, 1, t) }
func BenchmarkLinfDist2(t *testing.B)      { benchLinfDist(LinfDist, 2, t) }
func BenchmarkLinfDist3(t *testing.B)      { benchLinfDist(LinfDist, 3, t) }
func BenchmarkLinfDist4(t *testing.B)      { benchLinfDist(LinfDist, 4, t) }
func BenchmarkLinfDist5(t *testing.B)      { benchLinfDist(LinfDist, 5, t) }
func BenchmarkLinfDist10(t *testing.B)     { benchLinfDist(LinfDist, 10, t) }
func BenchmarkLinfDist100(t *testing.B)    { benchLinfDist(LinfDist, 100, t) }
func BenchmarkLinfDist1000(t *testing.B)   { benchLinfDist(LinfDist, 1000, t) }
func BenchmarkLinfDist10000(t *testing.B)  { benchLinfDist(LinfDist, 10000, t) }
func BenchmarkLinfDist100000(t *testing.B) { benchLinfDist(LinfDist, 100000, t) }
func BenchmarkLinfDist500000(t *testing.B) { benchLinfDist(LinfDist, 500000, t) }

func BenchmarkLLinfDist1(t *testing.B)      { benchLinfDist(naiveLinfDist, 1, t) }
func BenchmarkLLinfDist2(t *testing.B)      { benchLinfDist(naiveLinfDist, 2, t) }
func BenchmarkLLinfDist3(t *testing.B)      { benchLinfDist(naiveLinfDist, 3, t) }
func BenchmarkLLinfDist4(t *testing.B)      { benchLinfDist(naiveLinfDist, 4, t) }
func BenchmarkLLinfDist5(t *testing.B)      { benchLinfDist(naiveLinfDist, 5, t) }
func BenchmarkLLinfDist10(t *testing.B)     { benchLinfDist(naiveLinfDist, 10, t) }
func BenchmarkLLinfDist100(t *testing.B)    { benchLinfDist(naiveLinfDist, 100, t) }
func BenchmarkLLinfDist1000(t *testing.B)   { benchLinfDist(naiveLinfDist, 1000, t) }
func BenchmarkLLinfDist10000(t *testing.B)  { benchLinfDist(naiveLinfDist, 10000, t) }
func BenchmarkLLinfDist100000(t *testing.B) { benchLinfDist(naiveLinfDist, 100000, t) }
func BenchmarkLLinfDist500000(t *testing.B) { benchLinfDist(naiveLinfDist, 500000, t) }
