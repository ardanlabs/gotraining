// Copyright 2013 The Gonum Authors. All rights reserved.
// Use of this code is governed by a BSD-style
// license that can be found in the LICENSE file

package floats

import (
	"math"
	"math/rand"
	"strconv"
	"testing"
	"testing/quick"
)

const (
	EqTolerance = 1E-14
	Small       = 10
	Medium      = 1000
	Large       = 100000
	Huge        = 10000000
)

func AreSlicesEqual(t *testing.T, truth, comp []float64, str string) {
	if !EqualApprox(comp, truth, EqTolerance) {
		t.Errorf(str+". Expected %v, returned %v", truth, comp)
	}
}

func Panics(fun func()) (b bool) {
	defer func() {
		err := recover()
		if err != nil {
			b = true
		}
	}()
	fun()
	return
}

func TestAdd(t *testing.T) {
	a := []float64{1, 2, 3}
	b := []float64{4, 5, 6}
	c := []float64{7, 8, 9}
	truth := []float64{12, 15, 18}
	n := make([]float64, len(a))

	Add(n, a)
	Add(n, b)
	Add(n, c)
	AreSlicesEqual(t, truth, n, "Wrong addition of slices new receiver")
	Add(a, b)
	Add(a, c)
	AreSlicesEqual(t, truth, n, "Wrong addition of slices for no new receiver")

	// Test that it panics
	if !Panics(func() { Add(make([]float64, 2), make([]float64, 3)) }) {
		t.Errorf("Did not panic with length mismatch")
	}
}

func TestAddTo(t *testing.T) {
	a := []float64{1, 2, 3}
	b := []float64{4, 5, 6}
	truth := []float64{5, 7, 9}
	n1 := make([]float64, len(a))

	n2 := AddTo(n1, a, b)
	AreSlicesEqual(t, truth, n1, "Bad addition from mutator")
	AreSlicesEqual(t, truth, n2, "Bad addition from returned slice")

	// Test that it panics
	if !Panics(func() { AddTo(make([]float64, 2), make([]float64, 3), make([]float64, 3)) }) {
		t.Errorf("Did not panic with length mismatch")
	}
	if !Panics(func() { AddTo(make([]float64, 3), make([]float64, 3), make([]float64, 2)) }) {
		t.Errorf("Did not panic with length mismatch")
	}

}

func TestAddConst(t *testing.T) {
	s := []float64{3, 4, 1, 7, 5}
	c := 6.0
	truth := []float64{9, 10, 7, 13, 11}
	AddConst(c, s)
	AreSlicesEqual(t, truth, s, "Wrong addition of constant")
}

func TestAddScaled(t *testing.T) {
	s := []float64{3, 4, 1, 7, 5}
	alpha := 6.0
	dst := []float64{1, 2, 3, 4, 5}
	ans := []float64{19, 26, 9, 46, 35}
	AddScaled(dst, alpha, s)
	if !EqualApprox(dst, ans, EqTolerance) {
		t.Errorf("Adding scaled did not match")
	}
	short := []float64{1}
	if !Panics(func() { AddScaled(dst, alpha, short) }) {
		t.Errorf("Doesn't panic if s is smaller than dst")
	}
	if !Panics(func() { AddScaled(short, alpha, s) }) {
		t.Errorf("Doesn't panic if dst is smaller than s")
	}
}

func TestAddScaledTo(t *testing.T) {
	s := []float64{3, 4, 1, 7, 5}
	alpha := 6.0
	y := []float64{1, 2, 3, 4, 5}
	dst1 := make([]float64, 5)
	ans := []float64{19, 26, 9, 46, 35}
	dst2 := AddScaledTo(dst1, y, alpha, s)
	if !EqualApprox(dst1, ans, EqTolerance) {
		t.Errorf("AddScaledTo did not match for mutator")
	}
	if !EqualApprox(dst2, ans, EqTolerance) {
		t.Errorf("AddScaledTo did not match for returned slice")
	}
	AddScaledTo(dst1, y, alpha, s)
	if !EqualApprox(dst1, ans, EqTolerance) {
		t.Errorf("Reusing dst did not match")
	}
	short := []float64{1}
	if !Panics(func() { AddScaledTo(dst1, y, alpha, short) }) {
		t.Errorf("Doesn't panic if s is smaller than dst")
	}
	if !Panics(func() { AddScaledTo(short, y, alpha, s) }) {
		t.Errorf("Doesn't panic if dst is smaller than s")
	}
	if !Panics(func() { AddScaledTo(dst1, short, alpha, s) }) {
		t.Errorf("Doesn't panic if y is smaller than dst")
	}
}

func TestArgsort(t *testing.T) {
	s := []float64{3, 4, 1, 7, 5}
	inds := make([]int, len(s))

	Argsort(s, inds)

	sortedS := []float64{1, 3, 4, 5, 7}
	trueInds := []int{2, 0, 1, 4, 3}

	if !Equal(s, sortedS) {
		t.Error("elements not sorted correctly")
	}
	for i := range trueInds {
		if trueInds[i] != inds[i] {
			t.Error("inds not correct")
		}
	}

	inds = []int{1, 2}
	if !Panics(func() { Argsort(s, inds) }) {
		t.Error("does not panic if lengths do not match")
	}
}

func TestCount(t *testing.T) {
	s := []float64{3, 4, 1, 7, 5}
	f := func(v float64) bool { return v > 3.5 }
	truth := 3
	n := Count(f, s)
	if n != truth {
		t.Errorf("Wrong number of elements counted")
	}
}

func TestCumProd(t *testing.T) {
	s := []float64{3, 4, 1, 7, 5}
	receiver := make([]float64, len(s))
	result := CumProd(receiver, s)
	truth := []float64{3, 12, 12, 84, 420}
	AreSlicesEqual(t, truth, receiver, "Wrong cumprod mutated with new receiver")
	AreSlicesEqual(t, truth, result, "Wrong cumprod result with new receiver")
	CumProd(receiver, s)
	AreSlicesEqual(t, truth, receiver, "Wrong cumprod returned with reused receiver")

	// Test that it panics
	if !Panics(func() { CumProd(make([]float64, 2), make([]float64, 3)) }) {
		t.Errorf("Did not panic with length mismatch")
	}

	// Test empty CumProd
	emptyReceiver := make([]float64, 0)
	truth = []float64{}
	CumProd(emptyReceiver, emptyReceiver)
	AreSlicesEqual(t, truth, emptyReceiver, "Wrong cumprod returned with emtpy receiver")

}

func TestCumSum(t *testing.T) {
	s := []float64{3, 4, 1, 7, 5}
	receiver := make([]float64, len(s))
	result := CumSum(receiver, s)
	truth := []float64{3, 7, 8, 15, 20}
	AreSlicesEqual(t, truth, receiver, "Wrong cumsum mutated with new receiver")
	AreSlicesEqual(t, truth, result, "Wrong cumsum returned with new receiver")
	CumSum(receiver, s)
	AreSlicesEqual(t, truth, receiver, "Wrong cumsum returned with reused receiver")

	// Test that it panics
	if !Panics(func() { CumSum(make([]float64, 2), make([]float64, 3)) }) {
		t.Errorf("Did not panic with length mismatch")
	}

	// Test empty CumSum
	emptyReceiver := make([]float64, 0)
	truth = []float64{}
	CumSum(emptyReceiver, emptyReceiver)
	AreSlicesEqual(t, truth, emptyReceiver, "Wrong cumsum returned with emtpy receiver")

}

func TestDistance(t *testing.T) {
	norms := []float64{1, 2, 4, math.Inf(1)}
	slices := []struct {
		s []float64
		t []float64
	}{
		{
			nil,
			nil,
		},
		{
			[]float64{8, 9, 10, -12},
			[]float64{8, 9, 10, -12},
		},
		{
			[]float64{1, 2, 3, -4, -5, 8},
			[]float64{-9.2, -6.8, 9, -3, -2, 1},
		},
	}

	for j, test := range slices {
		tmp := make([]float64, len(test.s))
		for i, L := range norms {
			dist := Distance(test.s, test.t, L)
			copy(tmp, test.s)
			Sub(tmp, test.t)
			norm := Norm(tmp, L)
			if dist != norm { // Use equality because they should be identical
				t.Errorf("Distance does not match norm for case %v, %v. Expected %v, Found %v.", i, j, norm, dist)
			}
		}
	}

	if !Panics(func() { Distance([]float64{}, norms, 1) }) {
		t.Errorf("Did not panic with unequal lengths")
	}

}

func TestDiv(t *testing.T) {
	s1 := []float64{5, 12, 27}
	s2 := []float64{1, 2, 3}
	ans := []float64{5, 6, 9}
	Div(s1, s2)
	if !EqualApprox(s1, ans, EqTolerance) {
		t.Errorf("Mul doesn't give correct answer")
	}
	s1short := []float64{1}
	if !Panics(func() { Div(s1short, s2) }) {
		t.Errorf("Did not panic with unequal lengths")
	}
	s2short := []float64{1}
	if !Panics(func() { Div(s1, s2short) }) {
		t.Errorf("Did not panic with unequal lengths")
	}
}

func TestDivTo(t *testing.T) {
	s1 := []float64{5, 12, 27}
	s1orig := []float64{5, 12, 27}
	s2 := []float64{1, 2, 3}
	s2orig := []float64{1, 2, 3}
	dst1 := make([]float64, 3)
	ans := []float64{5, 6, 9}
	dst2 := DivTo(dst1, s1, s2)
	if !EqualApprox(dst1, ans, EqTolerance) {
		t.Errorf("DivTo doesn't give correct answer in mutated slice")
	}
	if !EqualApprox(dst2, ans, EqTolerance) {
		t.Errorf("DivTo doesn't give correct answer in returned slice")
	}
	if !EqualApprox(s1, s1orig, EqTolerance) {
		t.Errorf("S1 changes during multo")
	}
	if !EqualApprox(s2, s2orig, EqTolerance) {
		t.Errorf("s2 changes during multo")
	}
	DivTo(dst1, s1, s2)
	if !EqualApprox(dst1, ans, EqTolerance) {
		t.Errorf("DivTo doesn't give correct answer reusing dst")
	}
	dstShort := []float64{1}
	if !Panics(func() { DivTo(dstShort, s1, s2) }) {
		t.Errorf("Did not panic with s1 wrong length")
	}
	s1short := []float64{1}
	if !Panics(func() { DivTo(dst1, s1short, s2) }) {
		t.Errorf("Did not panic with s1 wrong length")
	}
	s2short := []float64{1}
	if !Panics(func() { DivTo(dst1, s1, s2short) }) {
		t.Errorf("Did not panic with s2 wrong length")
	}
}

func TestDot(t *testing.T) {
	s1 := []float64{1, 2, 3, 4}
	s2 := []float64{-3, 4, 5, -6}
	truth := -4.0
	ans := Dot(s1, s2)
	if ans != truth {
		t.Errorf("Dot product computed incorrectly")
	}

	// Test that it panics
	if !Panics(func() { Dot(make([]float64, 2), make([]float64, 3)) }) {
		t.Errorf("Did not panic with length mismatch")
	}
}

func TestEquals(t *testing.T) {
	s1 := []float64{1, 2, 3, 4}
	s2 := []float64{1, 2, 3, 4}
	if !Equal(s1, s2) {
		t.Errorf("Equal slices returned as unequal")
	}
	s2 = []float64{1, 2, 3, 4 + 1e-14}
	if Equal(s1, s2) {
		t.Errorf("Unequal slices returned as equal")
	}
	if Equal(s1, []float64{}) {
		t.Errorf("Unequal slice lengths returned as equal")
	}
}

func TestEqualApprox(t *testing.T) {
	s1 := []float64{1, 2, 3, 4}
	s2 := []float64{1, 2, 3, 4 + 1e-10}
	if EqualApprox(s1, s2, 1e-13) {
		t.Errorf("Unequal slices returned as equal for absolute")
	}
	if !EqualApprox(s1, s2, 1e-5) {
		t.Errorf("Equal slices returned as unequal for absolute")
	}
	s1 = []float64{1, 2, 3, 1000}
	s2 = []float64{1, 2, 3, 1000 * (1 + 1e-7)}
	if EqualApprox(s1, s2, 1e-8) {
		t.Errorf("Unequal slices returned as equal for relative")
	}
	if !EqualApprox(s1, s2, 1e-5) {
		t.Errorf("Equal slices returned as unequal for relative")
	}
	if EqualApprox(s1, []float64{}, 1e-5) {
		t.Errorf("Unequal slice lengths returned as equal")
	}
}

func TestEqualFunc(t *testing.T) {
	s1 := []float64{1, 2, 3, 4}
	s2 := []float64{1, 2, 3, 4}
	eq := func(x, y float64) bool { return x == y }
	if !EqualFunc(s1, s2, eq) {
		t.Errorf("Equal slices returned as unequal")
	}
	s2 = []float64{1, 2, 3, 4 + 1e-14}
	if EqualFunc(s1, s2, eq) {
		t.Errorf("Unequal slices returned as equal")
	}
	if EqualFunc(s1, []float64{}, eq) {
		t.Errorf("Unequal slice lengths returned as equal")
	}
}

func TestEqualsRelative(t *testing.T) {
	var equalityTests = []struct {
		a, b  float64
		tol   float64
		equal bool
	}{
		{1000000, 1000001, 0, true},
		{1000001, 1000000, 0, true},
		{10000, 10001, 0, false},
		{10001, 10000, 0, false},
		{-1000000, -1000001, 0, true},
		{-1000001, -1000000, 0, true},
		{-10000, -10001, 0, false},
		{-10001, -10000, 0, false},
		{1.0000001, 1.0000002, 0, true},
		{1.0000002, 1.0000001, 0, true},
		{1.0002, 1.0001, 0, false},
		{1.0001, 1.0002, 0, false},
		{-1.000001, -1.000002, 0, true},
		{-1.000002, -1.000001, 0, true},
		{-1.0001, -1.0002, 0, false},
		{-1.0002, -1.0001, 0, false},
		{0.000000001000001, 0.000000001000002, 0, true},
		{0.000000001000002, 0.000000001000001, 0, true},
		{0.000000000001002, 0.000000000001001, 0, false},
		{0.000000000001001, 0.000000000001002, 0, false},
		{-0.000000001000001, -0.000000001000002, 0, true},
		{-0.000000001000002, -0.000000001000001, 0, true},
		{-0.000000000001002, -0.000000000001001, 0, false},
		{-0.000000000001001, -0.000000000001002, 0, false},
		{0, 0, 0, true},
		{0, -0, 0, true},
		{-0, -0, 0, true},
		{0.00000001, 0, 0, false},
		{0, 0.00000001, 0, false},
		{-0.00000001, 0, 0, false},
		{0, -0.00000001, 0, false},
		{0, 1e-310, 0.01, true},
		{1e-310, 0, 0.01, true},
		{1e-310, 0, 0.000001, false},
		{0, 1e-310, 0.000001, false},
		{0, -1e-310, 0.1, true},
		{-1e-310, 0, 0.1, true},
		{-1e-310, 0, 0.00000001, false},
		{0, -1e-310, 0.00000001, false},
		{math.Inf(1), math.Inf(1), 0, true},
		{math.Inf(-1), math.Inf(-1), 0, true},
		{math.Inf(-1), math.Inf(1), 0, false},
		{math.Inf(1), math.MaxFloat64, 0, false},
		{math.Inf(-1), -math.MaxFloat64, 0, false},
		{math.NaN(), math.NaN(), 0, false},
		{math.NaN(), 0, 0, false},
		{-0, math.NaN(), 0, false},
		{math.NaN(), -0, 0, false},
		{0, math.NaN(), 0, false},
		{math.NaN(), math.Inf(1), 0, false},
		{math.Inf(1), math.NaN(), 0, false},
		{math.NaN(), math.Inf(-1), 0, false},
		{math.Inf(-1), math.NaN(), 0, false},
		{math.NaN(), math.MaxFloat64, 0, false},
		{math.MaxFloat64, math.NaN(), 0, false},
		{math.NaN(), -math.MaxFloat64, 0, false},
		{-math.MaxFloat64, math.NaN(), 0, false},
		{math.NaN(), math.SmallestNonzeroFloat64, 0, false},
		{math.SmallestNonzeroFloat64, math.NaN(), 0, false},
		{math.NaN(), -math.SmallestNonzeroFloat64, 0, false},
		{-math.SmallestNonzeroFloat64, math.NaN(), 0, false},
		{1.000000001, -1.0, 0, false},
		{-1.0, 1.000000001, 0, false},
		{-1.000000001, 1.0, 0, false},
		{1.0, -1.000000001, 0, false},
		{10 * math.SmallestNonzeroFloat64, 10 * -math.SmallestNonzeroFloat64, 0, true},
		{1e11 * math.SmallestNonzeroFloat64, 1e11 * -math.SmallestNonzeroFloat64, 0, false},
		{math.SmallestNonzeroFloat64, -math.SmallestNonzeroFloat64, 0, true},
		{-math.SmallestNonzeroFloat64, math.SmallestNonzeroFloat64, 0, true},
		{math.SmallestNonzeroFloat64, 0, 0, true},
		{0, math.SmallestNonzeroFloat64, 0, true},
		{-math.SmallestNonzeroFloat64, 0, 0, true},
		{0, -math.SmallestNonzeroFloat64, 0, true},
		{0.000000001, -math.SmallestNonzeroFloat64, 0, false},
		{0.000000001, math.SmallestNonzeroFloat64, 0, false},
		{math.SmallestNonzeroFloat64, 0.000000001, 0, false},
		{-math.SmallestNonzeroFloat64, 0.000000001, 0, false},
	}
	for _, ts := range equalityTests {
		if ts.tol == 0 {
			ts.tol = 1e-5
		}
		if equal := EqualWithinRel(ts.a, ts.b, ts.tol); equal != ts.equal {
			t.Errorf("Relative equality of %g and %g with tolerance %g returned: %v. Expected: %v",
				ts.a, ts.b, ts.tol, equal, ts.equal)
		}
	}
}

func nextAfterN(x, y float64, n int) float64 {
	for i := 0; i < n; i++ {
		x = math.Nextafter(x, y)
	}
	return x
}

func TestEqualsULP(t *testing.T) {
	if f := 67329.242; !EqualWithinULP(f, nextAfterN(f, math.Inf(1), 10), 10) {
		t.Errorf("Equal values returned as unequal")
	}
	if f := 67329.242; EqualWithinULP(f, nextAfterN(f, math.Inf(1), 5), 1) {
		t.Errorf("Unequal values returned as equal")
	}
	if f := 67329.242; EqualWithinULP(nextAfterN(f, math.Inf(1), 5), f, 1) {
		t.Errorf("Unequal values returned as equal")
	}
	if f := nextAfterN(0, math.Inf(1), 2); !EqualWithinULP(f, nextAfterN(f, math.Inf(-1), 5), 10) {
		t.Errorf("Equal values returned as unequal")
	}
	if !EqualWithinULP(67329.242, 67329.242, 10) {
		t.Errorf("Equal float64s not returned as equal")
	}
	if EqualWithinULP(1, math.NaN(), 10) {
		t.Errorf("NaN returned as equal")
	}

}

func TestEqualLengths(t *testing.T) {
	s1 := []float64{1, 2, 3, 4}
	s2 := []float64{1, 2, 3, 4}
	s3 := []float64{1, 2, 3}
	if !EqualLengths(s1, s2) {
		t.Errorf("Equal lengths returned as unequal")
	}
	if EqualLengths(s1, s3) {
		t.Errorf("Unequal lengths returned as equal")
	}
	if !EqualLengths(s1) {
		t.Errorf("Single slice returned as unequal")
	}
	if !EqualLengths() {
		t.Errorf("No slices returned as unequal")
	}
}

func eqIntSlice(one, two []int) string {
	if len(one) != len(two) {
		return "Length mismatch"
	}
	for i, val := range one {
		if val != two[i] {
			return "Index " + strconv.Itoa(i) + " mismatch"
		}
	}
	return ""
}

func TestFind(t *testing.T) {
	s := []float64{3, 4, 1, 7, 5}
	f := func(v float64) bool { return v > 3.5 }
	allTrueInds := []int{1, 3, 4}

	// Test finding first two elements
	inds, err := Find(nil, f, s, 2)
	if err != nil {
		t.Errorf("Find first two: Improper error return")
	}
	trueInds := allTrueInds[:2]
	str := eqIntSlice(inds, trueInds)
	if str != "" {
		t.Errorf("Find first two: " + str)
	}

	// Test finding no elements with non nil slice
	inds = []int{1, 2, 3, 4, 5, 6}
	inds, err = Find(inds, f, s, 0)
	if err != nil {
		t.Errorf("Find no elements: Improper error return")
	}
	str = eqIntSlice(inds, []int{})
	if str != "" {
		t.Errorf("Find no non-nil: " + str)
	}

	// Test finding first two elements with non nil slice
	inds = []int{1, 2, 3, 4, 5, 6}
	inds, err = Find(inds, f, s, 2)
	if err != nil {
		t.Errorf("Find first two non-nil: Improper error return")
	}
	str = eqIntSlice(inds, trueInds)
	if str != "" {
		t.Errorf("Find first two non-nil: " + str)
	}

	// Test finding too many elements
	inds, err = Find(inds, f, s, 4)
	if err == nil {
		t.Errorf("Request too many: No error returned")
	}
	str = eqIntSlice(inds, allTrueInds)
	if str != "" {
		t.Errorf("Request too many: Does not match all of the inds: " + str)
	}

	// Test finding all elements
	inds, err = Find(nil, f, s, -1)
	if err != nil {
		t.Errorf("Find all: Improper error returned")
	}
	str = eqIntSlice(inds, allTrueInds)
	if str != "" {
		t.Errorf("Find all: Does not match all of the inds: " + str)
	}
}

func TestHasNaN(t *testing.T) {
	for i, test := range []struct {
		s   []float64
		ans bool
	}{
		{},
		{
			s: []float64{1, 2, 3, 4},
		},
		{
			s:   []float64{1, math.NaN(), 3, 4},
			ans: true,
		},
		{
			s:   []float64{1, 2, 3, math.NaN()},
			ans: true,
		},
	} {
		b := HasNaN(test.s)
		if b != test.ans {
			t.Errorf("HasNaN mismatch case %d. Expected %v, Found %v", i, test.ans, b)
		}
	}
}

func TestLogSpan(t *testing.T) {
	receiver1 := make([]float64, 6)
	truth := []float64{0.001, 0.01, 0.1, 1, 10, 100}
	receiver2 := LogSpan(receiver1, 0.001, 100)
	tst := make([]float64, 6)
	for i := range truth {
		tst[i] = receiver1[i] / truth[i]
	}
	comp := make([]float64, 6)
	for i := range comp {
		comp[i] = 1
	}
	AreSlicesEqual(t, comp, tst, "Improper logspace from mutator")

	for i := range truth {
		tst[i] = receiver2[i] / truth[i]
	}
	AreSlicesEqual(t, comp, tst, "Improper logspace from returned slice")

	if !Panics(func() { LogSpan(nil, 1, 5) }) {
		t.Errorf("Span accepts nil argument")
	}
	if !Panics(func() { LogSpan(make([]float64, 1), 1, 5) }) {
		t.Errorf("Span accepts argument of len = 1")
	}
}

func TestLogSumExp(t *testing.T) {
	s := []float64{1, 2, 3, 4, 5}
	val := LogSumExp(s)
	// http://www.wolframalpha.com/input/?i=log%28exp%281%29+%2B+exp%282%29+%2B+exp%283%29+%2B+exp%284%29+%2B+exp%285%29%29
	truth := 5.4519143959375933331957225109748087179338972737576824
	if math.Abs(val-truth) > EqTolerance {
		t.Errorf("Wrong logsumexp for many values")
	}
	s = []float64{1, 2}
	// http://www.wolframalpha.com/input/?i=log%28exp%281%29+%2B+exp%282%29%29
	truth = 2.3132616875182228340489954949678556419152800856703483
	val = LogSumExp(s)
	if math.Abs(val-truth) > EqTolerance {
		t.Errorf("Wrong logsumexp for two values. %v expected, %v found", truth, val)
	}
	// This case would normally underflow
	s = []float64{-1001, -1002, -1003, -1004, -1005}
	// http://www.wolframalpha.com/input/?i=log%28exp%28-1001%29%2Bexp%28-1002%29%2Bexp%28-1003%29%2Bexp%28-1004%29%2Bexp%28-1005%29%29
	truth = -1000.54808560406240666680427748902519128206610272624
	val = LogSumExp(s)
	if math.Abs(val-truth) > EqTolerance {
		t.Errorf("Doesn't match for underflow case. %v expected, %v found", truth, val)
	}
	// positive infinite case
	s = []float64{1, 2, 3, 4, 5, math.Inf(1)}
	val = LogSumExp(s)
	truth = math.Inf(1)
	if val != truth {
		t.Errorf("Doesn't match for pos Infinity case. %v expected, %v found", truth, val)
	}
	// negative infinite case
	s = []float64{1, 2, 3, 4, 5, math.Inf(-1)}
	val = LogSumExp(s)
	truth = 5.4519143959375933331957225109748087179338972737576824 // same as first case
	if math.Abs(val-truth) > EqTolerance {
		t.Errorf("Wrong logsumexp for values with negative infinity")
	}

}

func TestMaxAndIdx(t *testing.T) {
	s := []float64{3, 4, 1, 7, 5}
	ind := MaxIdx(s)
	val := Max(s)
	if val != 7 {
		t.Errorf("Wrong value returned")
	}
	if ind != 3 {
		t.Errorf("Wrong index returned")
	}
}

func TestMinAndIdx(t *testing.T) {
	s := []float64{3, 4, 1, 7, 5}
	ind := MinIdx(s)
	val := Min(s)
	if val != 1 {
		t.Errorf("Wrong value returned")
	}
	if ind != 2 {
		t.Errorf("Wrong index returned")
	}
}

func TestMul(t *testing.T) {
	s1 := []float64{1, 2, 3}
	s2 := []float64{1, 2, 3}
	ans := []float64{1, 4, 9}
	Mul(s1, s2)
	if !EqualApprox(s1, ans, EqTolerance) {
		t.Errorf("Mul doesn't give correct answer")
	}
	s1short := []float64{1}
	if !Panics(func() { Mul(s1short, s2) }) {
		t.Errorf("Did not panic with unequal lengths")
	}
	s2short := []float64{1}
	if !Panics(func() { Mul(s1, s2short) }) {
		t.Errorf("Did not panic with unequal lengths")
	}
}

func TestMulTo(t *testing.T) {
	s1 := []float64{1, 2, 3}
	s1orig := []float64{1, 2, 3}
	s2 := []float64{1, 2, 3}
	s2orig := []float64{1, 2, 3}
	dst1 := make([]float64, 3)
	ans := []float64{1, 4, 9}
	dst2 := MulTo(dst1, s1, s2)
	if !EqualApprox(dst1, ans, EqTolerance) {
		t.Errorf("MulTo doesn't give correct answer in mutated slice")
	}
	if !EqualApprox(dst2, ans, EqTolerance) {
		t.Errorf("MulTo doesn't give correct answer in returned slice")
	}
	if !EqualApprox(s1, s1orig, EqTolerance) {
		t.Errorf("S1 changes during multo")
	}
	if !EqualApprox(s2, s2orig, EqTolerance) {
		t.Errorf("s2 changes during multo")
	}
	MulTo(dst1, s1, s2)
	if !EqualApprox(dst1, ans, EqTolerance) {
		t.Errorf("MulTo doesn't give correct answer reusing dst")
	}
	dstShort := []float64{1}
	if !Panics(func() { MulTo(dstShort, s1, s2) }) {
		t.Errorf("Did not panic with s1 wrong length")
	}
	s1short := []float64{1}
	if !Panics(func() { MulTo(dst1, s1short, s2) }) {
		t.Errorf("Did not panic with s1 wrong length")
	}
	s2short := []float64{1}
	if !Panics(func() { MulTo(dst1, s1, s2short) }) {
		t.Errorf("Did not panic with s2 wrong length")
	}
}

func TestNearest(t *testing.T) {
	s := []float64{6.2, 3, 5, 6.2, 8}
	ind := Nearest(s, 2.0)
	if ind != 1 {
		t.Errorf("Wrong index returned when value is less than all of elements")
	}
	ind = Nearest(s, 9.0)
	if ind != 4 {
		t.Errorf("Wrong index returned when value is greater than all of elements")
	}
	ind = Nearest(s, 3.1)
	if ind != 1 {
		t.Errorf("Wrong index returned when value is greater than closest element")
	}
	ind = Nearest(s, 3.1)
	if ind != 1 {
		t.Errorf("Wrong index returned when value is greater than closest element")
	}
	ind = Nearest(s, 2.9)
	if ind != 1 {
		t.Errorf("Wrong index returned when value is less than closest element")
	}
	ind = Nearest(s, 3)
	if ind != 1 {
		t.Errorf("Wrong index returned when value is equal to element")
	}
	ind = Nearest(s, 6.2)
	if ind != 0 {
		t.Errorf("Wrong index returned when value is equal to several elements")
	}
	ind = Nearest(s, 4)
	if ind != 1 {
		t.Errorf("Wrong index returned when value is exactly between two closest elements")
	}
}

func TestNearestWithinSpan(t *testing.T) {
	if !Panics(func() { NearestWithinSpan(10, 8, 2, 4.5) }) {
		t.Errorf("Did not panic when upper bound is lower than greater bound")
	}
	for i, test := range []struct {
		length int
		lower  float64
		upper  float64
		value  float64
		idx    int
	}{
		{
			length: 13,
			lower:  7,
			upper:  8.2,
			value:  6,
			idx:    -1,
		},
		{
			length: 13,
			lower:  7,
			upper:  8.2,
			value:  10,
			idx:    -1,
		},
		{
			length: 13,
			lower:  7,
			upper:  8.2,
			value:  7.19,
			idx:    2,
		},
		{
			length: 13,
			lower:  7,
			upper:  8.2,
			value:  7.21,
			idx:    2,
		},
		{
			length: 13,
			lower:  7,
			upper:  8.2,
			value:  7.2,
			idx:    2,
		},
		{
			length: 13,
			lower:  7,
			upper:  8.2,
			value:  7.151,
			idx:    2,
		},
		{
			length: 13,
			lower:  7,
			upper:  8.2,
			value:  7.249,
			idx:    2,
		},
	} {
		if idx := NearestWithinSpan(test.length, test.lower, test.upper, test.value); test.idx != idx {
			t.Errorf("Case %v mismatch: Want: %v, Got: %v", i, test.idx, idx)
		}
	}
}

func TestNorm(t *testing.T) {
	s := []float64{-1, -3.4, 5, -6}
	val := Norm(s, math.Inf(1))
	truth := 6.0
	if math.Abs(val-truth) > EqTolerance {
		t.Errorf("Doesn't match for inf norm. %v expected, %v found", truth, val)
	}
	// http://www.wolframalpha.com/input/?i=%28%28-1%29%5E2+%2B++%28-3.4%29%5E2+%2B+5%5E2%2B++6%5E2%29%5E%281%2F2%29
	val = Norm(s, 2)
	truth = 8.5767126569566267590651614132751986658027271236078592
	if math.Abs(val-truth) > EqTolerance {
		t.Errorf("Doesn't match for inf norm. %v expected, %v found", truth, val)
	}
	// http://www.wolframalpha.com/input/?i=%28%28%7C-1%7C%29%5E3+%2B++%28%7C-3.4%7C%29%5E3+%2B+%7C5%7C%5E3%2B++%7C6%7C%5E3%29%5E%281%2F3%29
	val = Norm(s, 3)
	truth = 7.2514321388020228478109121239004816430071237369356233
	if math.Abs(val-truth) > EqTolerance {
		t.Errorf("Doesn't match for inf norm. %v expected, %v found", truth, val)
	}

	//http://www.wolframalpha.com/input/?i=%7C-1%7C+%2B+%7C-3.4%7C+%2B+%7C5%7C%2B++%7C6%7C
	val = Norm(s, 1)
	truth = 15.4
	if math.Abs(val-truth) > EqTolerance {
		t.Errorf("Doesn't match for inf norm. %v expected, %v found", truth, val)
	}
}

func TestProd(t *testing.T) {
	s := []float64{}
	val := Prod(s)
	if val != 1 {
		t.Errorf("Val not returned as default when slice length is zero")
	}
	s = []float64{3, 4, 1, 7, 5}
	val = Prod(s)
	if val != 420 {
		t.Errorf("Wrong prod returned. Expected %v returned %v", 420, val)
	}
}

func TestReverse(t *testing.T) {
	for _, s := range [][]float64{
		{0},
		{1, 0},
		{2, 1, 0},
		{3, 2, 1, 0},
		{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
	} {
		Reverse(s)
		for i, v := range s {
			if v != float64(i) {
				t.Errorf("unexpected values for element %d: got:%v want:%v", i, v, i)
			}
		}
	}
}

func roundFloat(x float64, prec int) float64 {
	f, _ := strconv.ParseFloat(strconv.FormatFloat(x, 'f', prec, 64), 64)
	if f == 0 {
		return math.Abs(f)
	}
	return f
}

func TestRound(t *testing.T) {
	for _, x := range []float64{
		0,
		math.Inf(1),
		math.NaN(),
		func() float64 { var f float64; return -f }(),
		math.MaxFloat64 / 2,
		1 << 64,
		454.4445,
		454.44445,
		0.42499,
		0.42599,
		0.424999999999993,
		0.425000000000001,
		123.42499999999993,

		// FIXME(kortschak): These values fail due to roundFloat rounding
		// 454.45 to ±454.4 when prec=1 and 454.445 to ±454.44 when prec=2.
		// This is a result of fmt's rendering of the value.
		// 454.45,
		// 454.445,
	} {
		for _, sign := range []float64{1, -1} {
			for prec := 0; prec < 10; prec++ {
				got := Round(sign*x, prec)
				want := roundFloat(sign*x, prec)
				if (got != want || math.Signbit(got) != math.Signbit(want)) && !(math.IsNaN(got) && math.IsNaN(want)) {
					t.Errorf("unexpected result for Round(%g, %d): got: %g, want: %g", x, prec, got, want)
				}
			}
		}
	}

	// Special cases.
	for _, test := range []struct {
		x    float64
		prec int
		want float64
	}{
		// Failing cases above.
		{x: 454.45, prec: 0, want: 454},
		{x: 454.45, prec: 1, want: 454.5},
		{x: 454.45, prec: 2, want: 454.45},
		{x: 454.45, prec: 3, want: 454.45},
		{x: 454.445, prec: 0, want: 454},
		{x: 454.445, prec: 1, want: 454.4},
		{x: 454.445, prec: 2, want: 454.45},
		{x: 454.445, prec: 3, want: 454.445},
		{x: 454.445, prec: 4, want: 454.445},

		// Negative precision.
		{x: 454.45, prec: -1, want: 450},
		{x: 454.45, prec: -2, want: 500},
		{x: 500, prec: -3, want: 1000},
		{x: 500, prec: -4, want: 0},
	} {
		for _, sign := range []float64{1, -1} {
			got := Round(sign*test.x, test.prec)
			want := sign * test.want
			if want == 0 {
				want = 0
			}
			if got != want || math.Signbit(got) != math.Signbit(want) {
				t.Errorf("unexpected result for Round(%g, %d): got: %g, want: %g", test.x, test.prec, got, test.want)
			}
		}
	}

	// Test many large numbers. We hit float precision
	// issues below 1e16 so we avoid that domain.
	err := quick.Check(func(x float64, prec int) bool {
		prec %= 20
		if prec < 0 {
			prec = -prec
		}
		for x > 1e16 {
			got := Round(x, prec)
			want := roundFloat(x, prec)
			if (got != want || math.Signbit(got) != math.Signbit(want)) && !(math.IsNaN(got) && math.IsNaN(want)) {
				t.Logf("big numbers failed with prec=%d x=%f got=%f want=%f", prec, x, got, want)
				return false
			}
			x /= 10
		}
		return true
	}, nil)
	if err != nil {
		t.Error(err)
	}

	// Test many small numbers.
	err = quick.Check(func(mant, exp int, prec int) bool {
		prec %= 20
		if prec < 0 {
			prec = -prec
		}
		mant %= 1e10
		exp %= 40
		if exp > 0 {
			exp = -exp
		}
		x := float64(mant) * math.Pow10(exp)
		_, x = math.Modf(x)
		if prec > -exp {
			prec = -exp
		}
		got := strconv.FormatFloat(Round(x, prec), 'g', prec, 64)
		want := strconv.FormatFloat(roundFloat(x, prec), 'g', prec, 64)
		if got != want {
			t.Logf("small numbers failed with prec=%d x=%f got=%s want=%s", prec, x, got, want)
			return false
		}
		return true
	}, nil)
	if err != nil {
		t.Error(err)
	}
}

func TestRoundEven(t *testing.T) {
	for _, test := range []struct {
		x    float64
		prec int
		want float64
	}{
		{x: 0, prec: 1, want: 0},
		{x: math.Inf(1), prec: 1, want: math.Inf(1)},
		{x: math.NaN(), prec: 1, want: math.NaN()},
		{x: func() float64 { var f float64; return -f }(), prec: 1, want: 0},
		{x: math.MaxFloat64 / 2, prec: 1, want: math.MaxFloat64 / 2},
		{x: 1 << 64, prec: 1, want: 1 << 64},
		{x: 454.4445, prec: 3, want: 454.444},
		{x: 454.44445, prec: 4, want: 454.4444},
		{x: 0.42499, prec: 4, want: 0.425},
		{x: 0.42599, prec: 4, want: 0.426},
		{x: 0.424999999999993, prec: 2, want: 0.42},
		{x: 0.425, prec: 2, want: 0.42},
		{x: 0.425000000000001, prec: 2, want: 0.43},
		{x: 123.4244999999999, prec: 3, want: 123.424},
		{x: 123.4245, prec: 3, want: 123.424},
		{x: 123.4245000000001, prec: 3, want: 123.425},

		{x: 454.45, prec: 0, want: 454},
		{x: 454.45, prec: 1, want: 454.4},
		{x: 454.45, prec: 2, want: 454.45},
		{x: 454.45, prec: 3, want: 454.45},
		{x: 454.445, prec: 0, want: 454},
		{x: 454.445, prec: 1, want: 454.4},
		{x: 454.445, prec: 2, want: 454.44},
		{x: 454.445, prec: 3, want: 454.445},
		{x: 454.445, prec: 4, want: 454.445},
		{x: 454.55, prec: 0, want: 455},
		{x: 454.55, prec: 1, want: 454.6},
		{x: 454.55, prec: 2, want: 454.55},
		{x: 454.55, prec: 3, want: 454.55},
		{x: 454.455, prec: 0, want: 454},
		{x: 454.455, prec: 1, want: 454.5},
		{x: 454.455, prec: 2, want: 454.46},
		{x: 454.455, prec: 3, want: 454.455},
		{x: 454.455, prec: 4, want: 454.455},

		// Negative precision.
		{x: 454.45, prec: -1, want: 450},
		{x: 454.45, prec: -2, want: 500},
		{x: 500, prec: -3, want: 0},
		{x: 500, prec: -4, want: 0},
		{x: 1500, prec: -3, want: 2000},
		{x: 1500, prec: -4, want: 0},
	} {
		for _, sign := range []float64{1, -1} {
			got := RoundEven(sign*test.x, test.prec)
			want := sign * test.want
			if want == 0 {
				want = 0
			}
			if (got != want || math.Signbit(got) != math.Signbit(want)) && !(math.IsNaN(got) && math.IsNaN(want)) {
				t.Errorf("unexpected result for RoundEven(%g, %d): got: %g, want: %g", sign*test.x, test.prec, got, want)
			}
		}
	}
}

func TestSame(t *testing.T) {
	s1 := []float64{1, 2, 3, 4}
	s2 := []float64{1, 2, 3, 4}
	if !Same(s1, s2) {
		t.Errorf("Equal slices returned as unequal")
	}
	s2 = []float64{1, 2, 3, 4 + 1e-14}
	if Same(s1, s2) {
		t.Errorf("Unequal slices returned as equal")
	}
	if Same(s1, []float64{}) {
		t.Errorf("Unequal slice lengths returned as equal")
	}
	s1 = []float64{1, 2, math.NaN(), 4}
	s2 = []float64{1, 2, math.NaN(), 4}
	if !Same(s1, s2) {
		t.Errorf("Slices with matching NaN values returned as unequal")
	}
	s1 = []float64{1, 2, math.NaN(), 4}
	s2 = []float64{1, math.NaN(), 3, 4}
	if !Same(s1, s2) {
		t.Errorf("Slices with unmatching NaN values returned as equal")
	}
}

func TestScale(t *testing.T) {
	s := []float64{3, 4, 1, 7, 5}
	c := 5.0
	truth := []float64{15, 20, 5, 35, 25}
	Scale(c, s)
	AreSlicesEqual(t, truth, s, "Bad scaling")
}

func TestSpan(t *testing.T) {
	receiver1 := make([]float64, 5)
	truth := []float64{1, 2, 3, 4, 5}
	receiver2 := Span(receiver1, 1, 5)
	AreSlicesEqual(t, truth, receiver1, "Improper linspace from mutator")
	AreSlicesEqual(t, truth, receiver2, "Improper linspace from returned slice")
	receiver1 = make([]float64, 6)
	truth = []float64{0, 0.2, 0.4, 0.6, 0.8, 1.0}
	Span(receiver1, 0, 1)
	AreSlicesEqual(t, truth, receiver1, "Improper linspace")
	if !Panics(func() { Span(nil, 1, 5) }) {
		t.Errorf("Span accepts nil argument")
	}
	if !Panics(func() { Span(make([]float64, 1), 1, 5) }) {
		t.Errorf("Span accepts argument of len = 1")
	}
}

func TestSub(t *testing.T) {
	s := []float64{3, 4, 1, 7, 5}
	v := []float64{1, 2, 3, 4, 5}
	truth := []float64{2, 2, -2, 3, 0}
	Sub(s, v)
	AreSlicesEqual(t, truth, s, "Bad subtract")
	// Test that it panics
	if !Panics(func() { Sub(make([]float64, 2), make([]float64, 3)) }) {
		t.Errorf("Did not panic with length mismatch")
	}
}

func TestSubTo(t *testing.T) {
	s := []float64{3, 4, 1, 7, 5}
	v := []float64{1, 2, 3, 4, 5}
	truth := []float64{2, 2, -2, 3, 0}
	dst1 := make([]float64, len(s))
	dst2 := SubTo(dst1, s, v)
	AreSlicesEqual(t, truth, dst1, "Bad subtract from mutator")
	AreSlicesEqual(t, truth, dst2, "Bad subtract from returned slice")
	// Test that all mismatch combinations panic
	if !Panics(func() { SubTo(make([]float64, 2), make([]float64, 3), make([]float64, 3)) }) {
		t.Errorf("Did not panic with dst different length")
	}
	if !Panics(func() { SubTo(make([]float64, 3), make([]float64, 2), make([]float64, 3)) }) {
		t.Errorf("Did not panic with subtractor different length")
	}
	if !Panics(func() { SubTo(make([]float64, 3), make([]float64, 3), make([]float64, 2)) }) {
		t.Errorf("Did not panic with subtractee different length")
	}
}

func TestSum(t *testing.T) {
	s := []float64{}
	val := Sum(s)
	if val != 0 {
		t.Errorf("Val not returned as default when slice length is zero")
	}
	s = []float64{3, 4, 1, 7, 5}
	val = Sum(s)
	if val != 20 {
		t.Errorf("Wrong sum returned")
	}
}

func TestWithin(t *testing.T) {
	for i, test := range []struct {
		s      []float64
		v      float64
		idx    int
		panics bool
	}{
		{
			s:   []float64{1, 2, 5, 9},
			v:   1,
			idx: 0,
		},
		{
			s:   []float64{1, 2, 5, 9},
			v:   9,
			idx: -1,
		},
		{
			s:   []float64{1, 2, 5, 9},
			v:   1.5,
			idx: 0,
		},
		{
			s:   []float64{1, 2, 5, 9},
			v:   2,
			idx: 1,
		},
		{
			s:   []float64{1, 2, 5, 9},
			v:   2.5,
			idx: 1,
		},
		{
			s:   []float64{1, 2, 5, 9},
			v:   -3,
			idx: -1,
		},
		{
			s:   []float64{1, 2, 5, 9},
			v:   15,
			idx: -1,
		},
		{
			s:   []float64{1, 2, 5, 9},
			v:   math.NaN(),
			idx: -1,
		},
		{
			s:      []float64{5, 2, 6},
			panics: true,
		},
		{
			panics: true,
		},
		{
			s:      []float64{1},
			panics: true,
		},
	} {
		var idx int
		panics := Panics(func() { idx = Within(test.s, test.v) })
		if panics {
			if !test.panics {
				t.Errorf("Case %v: bad panic", i)
			}
			continue
		}
		if test.panics {
			if !panics {
				t.Errorf("Case %v: did not panic when it should", i)
			}
			continue
		}
		if idx != test.idx {
			t.Errorf("Case %v: Idx mismatch. Want: %v, got: %v", i, test.idx, idx)
		}
	}

}

func randomSlice(l int) []float64 {
	s := make([]float64, l)
	for i := range s {
		s[i] = rand.Float64()
	}
	return s
}

func benchmarkMin(b *testing.B, size int) {
	s := randomSlice(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Min(s)
	}
}
func BenchmarkMinSmall(b *testing.B) { benchmarkMin(b, Small) }
func BenchmarkMinMed(b *testing.B)   { benchmarkMin(b, Medium) }
func BenchmarkMinLarge(b *testing.B) { benchmarkMin(b, Large) }
func BenchmarkMinHuge(b *testing.B)  { benchmarkMin(b, Huge) }

func benchmarkAdd(b *testing.B, size int) {
	s1 := randomSlice(size)
	s2 := randomSlice(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Add(s1, s2)
	}
}
func BenchmarkAddSmall(b *testing.B) { benchmarkAdd(b, Small) }
func BenchmarkAddMed(b *testing.B)   { benchmarkAdd(b, Medium) }
func BenchmarkAddLarge(b *testing.B) { benchmarkAdd(b, Large) }
func BenchmarkAddHuge(b *testing.B)  { benchmarkAdd(b, Huge) }

func benchmarkAddTo(b *testing.B, size int) {
	s1 := randomSlice(size)
	s2 := randomSlice(size)
	dst := randomSlice(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		AddTo(dst, s1, s2)
	}
}
func BenchmarkAddToSmall(b *testing.B) { benchmarkAddTo(b, Small) }
func BenchmarkAddToMed(b *testing.B)   { benchmarkAddTo(b, Medium) }
func BenchmarkAddToLarge(b *testing.B) { benchmarkAddTo(b, Large) }
func BenchmarkAddToHuge(b *testing.B)  { benchmarkAddTo(b, Huge) }

func benchmarkCumProd(b *testing.B, size int) {
	s := randomSlice(size)
	dst := randomSlice(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CumProd(dst, s)
	}
}
func BenchmarkCumProdSmall(b *testing.B) { benchmarkCumProd(b, Small) }
func BenchmarkCumProdMed(b *testing.B)   { benchmarkCumProd(b, Medium) }
func BenchmarkCumProdLarge(b *testing.B) { benchmarkCumProd(b, Large) }
func BenchmarkCumProdHuge(b *testing.B)  { benchmarkCumProd(b, Huge) }

func benchmarkCumSum(b *testing.B, size int) {
	s := randomSlice(size)
	dst := randomSlice(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CumSum(dst, s)
	}
}
func BenchmarkCumSumSmall(b *testing.B) { benchmarkCumSum(b, Small) }
func BenchmarkCumSumMed(b *testing.B)   { benchmarkCumSum(b, Medium) }
func BenchmarkCumSumLarge(b *testing.B) { benchmarkCumSum(b, Large) }
func BenchmarkCumSumHuge(b *testing.B)  { benchmarkCumSum(b, Huge) }

func benchmarkDiv(b *testing.B, size int) {
	s := randomSlice(size)
	dst := randomSlice(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Div(dst, s)
	}
}
func BenchmarkDivSmall(b *testing.B) { benchmarkDiv(b, Small) }
func BenchmarkDivMed(b *testing.B)   { benchmarkDiv(b, Medium) }
func BenchmarkDivLarge(b *testing.B) { benchmarkDiv(b, Large) }
func BenchmarkDivHuge(b *testing.B)  { benchmarkDiv(b, Huge) }

func benchmarkDivTo(b *testing.B, size int) {
	s1 := randomSlice(size)
	s2 := randomSlice(size)
	dst := randomSlice(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DivTo(dst, s1, s2)
	}
}
func BenchmarkDivToSmall(b *testing.B) { benchmarkDivTo(b, Small) }
func BenchmarkDivToMed(b *testing.B)   { benchmarkDivTo(b, Medium) }
func BenchmarkDivToLarge(b *testing.B) { benchmarkDivTo(b, Large) }
func BenchmarkDivToHuge(b *testing.B)  { benchmarkDivTo(b, Huge) }

func benchmarkSub(b *testing.B, size int) {
	s1 := randomSlice(size)
	s2 := randomSlice(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sub(s1, s2)
	}
}
func BenchmarkSubSmall(b *testing.B) { benchmarkSub(b, Small) }
func BenchmarkSubMed(b *testing.B)   { benchmarkSub(b, Medium) }
func BenchmarkSubLarge(b *testing.B) { benchmarkSub(b, Large) }
func BenchmarkSubHuge(b *testing.B)  { benchmarkSub(b, Huge) }

func benchmarkSubTo(b *testing.B, size int) {
	s1 := randomSlice(size)
	s2 := randomSlice(size)
	dst := randomSlice(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SubTo(dst, s1, s2)
	}
}
func BenchmarkSubToSmall(b *testing.B) { benchmarkSubTo(b, Small) }
func BenchmarkSubToMed(b *testing.B)   { benchmarkSubTo(b, Medium) }
func BenchmarkSubToLarge(b *testing.B) { benchmarkSubTo(b, Large) }
func BenchmarkSubToHuge(b *testing.B)  { benchmarkSubTo(b, Huge) }

func benchmarkLogSumExp(b *testing.B, size int) {
	s := randomSlice(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LogSumExp(s)
	}
}
func BenchmarkLogSumExpSmall(b *testing.B) { benchmarkLogSumExp(b, Small) }
func BenchmarkLogSumExpMed(b *testing.B)   { benchmarkLogSumExp(b, Medium) }
func BenchmarkLogSumExpLarge(b *testing.B) { benchmarkLogSumExp(b, Large) }
func BenchmarkLogSumExpHuge(b *testing.B)  { benchmarkLogSumExp(b, Huge) }

func benchmarkDot(b *testing.B, size int) {
	s1 := randomSlice(size)
	s2 := randomSlice(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Dot(s1, s2)
	}
}
func BenchmarkDotSmall(b *testing.B) { benchmarkDot(b, Small) }
func BenchmarkDotMed(b *testing.B)   { benchmarkDot(b, Medium) }
func BenchmarkDotLarge(b *testing.B) { benchmarkDot(b, Large) }
func BenchmarkDotHuge(b *testing.B)  { benchmarkDot(b, Huge) }

func benchmarkAddScaledTo(b *testing.B, size int) {
	dst := randomSlice(size)
	y := randomSlice(size)
	s := randomSlice(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		AddScaledTo(dst, y, 2.3, s)
	}
}
func BenchmarkAddScaledToSmall(b *testing.B)  { benchmarkAddScaledTo(b, Small) }
func BenchmarkAddScaledToMedium(b *testing.B) { benchmarkAddScaledTo(b, Medium) }
func BenchmarkAddScaledToLarge(b *testing.B)  { benchmarkAddScaledTo(b, Large) }
func BenchmarkAddScaledToHuge(b *testing.B)   { benchmarkAddScaledTo(b, Huge) }

func benchmarkScale(b *testing.B, size int) {
	dst := randomSlice(size)
	b.ResetTimer()
	for i := 0; i < b.N; i += 2 {
		Scale(2.0, dst)
		Scale(0.5, dst)
	}
}
func BenchmarkScaleSmall(b *testing.B)  { benchmarkScale(b, Small) }
func BenchmarkScaleMedium(b *testing.B) { benchmarkScale(b, Medium) }
func BenchmarkScaleLarge(b *testing.B)  { benchmarkScale(b, Large) }
func BenchmarkScaleHuge(b *testing.B)   { benchmarkScale(b, Huge) }
