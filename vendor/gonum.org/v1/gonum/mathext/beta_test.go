// Copyright ©2016 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mathext_test

import (
	"math"
	"testing"

	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mathext"
)

var betaTests = []struct {
	p, q float64
	want float64
}{
	{
		p:    1,
		q:    2,
		want: 0.5, // obtained from scipy.special.beta(1,2) (version=0.18.0)
	},
	{
		p:    10,
		q:    20,
		want: 4.9925087406346778e-09, // obtained from scipy.special.beta(10,20) (version=0.18.0)
	},
	{
		p:    +0,
		q:    10,
		want: math.Inf(+1),
	},
	{
		p:    -0,
		q:    10,
		want: math.Inf(+1),
	},
	{
		p:    0,
		q:    0,
		want: math.NaN(),
	},
	{
		p:    0,
		q:    math.Inf(-1),
		want: math.NaN(),
	},
	{
		p:    10,
		q:    math.Inf(-1),
		want: math.NaN(),
	},
	{
		p:    0,
		q:    math.Inf(+1),
		want: math.NaN(),
	},
	{
		p:    10,
		q:    math.Inf(+1),
		want: math.NaN(),
	},
	{
		p:    math.NaN(),
		q:    10,
		want: math.NaN(),
	},
	{
		p:    math.NaN(),
		q:    0,
		want: math.NaN(),
	},
	{
		p:    -1,
		q:    0,
		want: math.NaN(),
	},
	{
		p:    -1,
		q:    +1,
		want: math.NaN(),
	},
}

func TestBeta(t *testing.T) {
	for i, test := range betaTests {
		v := mathext.Beta(test.p, test.q)
		testOK := func(x float64) bool {
			return floats.EqualWithinAbsOrRel(x, test.want, 1e-15, 1e-15) || (math.IsNaN(test.want) && math.IsNaN(x))
		}
		if !testOK(v) {
			t.Errorf("test #%d: Beta(%v, %v)=%v. want=%v\n",
				i, test.p, test.q, v, test.want,
			)
		}

		u := mathext.Beta(test.q, test.p)
		if !testOK(u) {
			t.Errorf("test #%[1]d: Beta(%[2]v, %[3]v)=%[4]v != Beta(%[3]v, %[2]v)=%[5]v)\n",
				i, test.p, test.q, v, u,
			)
		}

		if math.IsInf(v, +1) || math.IsNaN(v) {
			continue
		}

		vv := mathext.Beta(test.p, test.q+1)
		uu := mathext.Beta(test.p+1, test.q)
		if !floats.EqualWithinAbsOrRel(v, vv+uu, 1e-15, 1e-15) {
			t.Errorf(
				"test #%[1]d: Beta(%[2]v, %[3]v)=%[4]v != Beta(%[2]v+1, %[3]v) + Beta(%[2]v, %[3]v+1) (=%[5]v + %[6]v = %[7]v)\n",
				i, test.p, test.q, v, uu, vv, uu+vv,
			)
		}

		vbeta2 := beta2(test.p, test.q)
		if !floats.EqualWithinAbsOrRel(v, vbeta2, 1e-15, 1e-15) {
			t.Errorf(
				"test #%[1]d: Beta(%[2]v, %[3]v) != Γ(p)Γ(q) / Γ(p+q) (v=%[4]v u=%[5]v)\n",
				i, test.p, test.q, v, vbeta2,
			)
		}
	}
}

func beta2(x, y float64) float64 {
	return math.Gamma(x) * math.Gamma(y) / math.Gamma(x+y)
}

func BenchmarkBeta(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = mathext.Beta(10, 20)
	}
}

func BenchmarkBeta2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = math.Gamma(10) * math.Gamma(20) / math.Gamma(10+20)
	}
}

func TestLbeta(t *testing.T) {
	for i, test := range betaTests {
		want := math.Log(test.want)
		v := mathext.Lbeta(test.p, test.q)

		testOK := func(x float64) bool {
			return floats.EqualWithinAbsOrRel(x, want, 1e-15, 1e-15) || (math.IsNaN(want) && math.IsNaN(x))
		}
		if !testOK(v) {
			t.Errorf("test #%d: Lbeta(%v, %v)=%v. want=%v\n",
				i, test.p, test.q, v, want,
			)
		}

		u := mathext.Lbeta(test.q, test.p)
		if !testOK(u) {
			t.Errorf("test #%[1]d: Lbeta(%[2]v, %[3]v)=%[4]v != Lbeta(%[3]v, %[2]v)=%[5]v)\n",
				i, test.p, test.q, v, u,
			)
		}

		if math.IsInf(v, +1) || math.IsNaN(v) {
			continue
		}

		vbeta2 := math.Log(beta2(test.p, test.q))
		if !floats.EqualWithinAbsOrRel(v, vbeta2, 1e-15, 1e-15) {
			t.Errorf(
				"test #%[1]d: Lbeta(%[2]v, %[3]v) != Log(Γ(p)Γ(q) / Γ(p+q)) (v=%[4]v u=%[5]v)\n",
				i, test.p, test.q, v, vbeta2,
			)
		}
	}
}
