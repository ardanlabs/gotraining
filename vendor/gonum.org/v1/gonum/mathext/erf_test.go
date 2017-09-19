// Copyright ©2017 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mathext

import (
	"testing"

	"gonum.org/v1/gonum/floats"
)

func TestNormalQuantile(t *testing.T) {
	// Values from https://www.johndcook.com/blog/normal_cdf_inverse/
	p := []float64{
		0.0000001,
		0.00001,
		0.001,
		0.05,
		0.15,
		0.25,
		0.35,
		0.45,
		0.55,
		0.65,
		0.75,
		0.85,
		0.95,
		0.999,
		0.99999,
		0.9999999,
	}
	ans := []float64{
		-5.199337582187471,
		-4.264890793922602,
		-3.090232306167813,
		-1.6448536269514729,
		-1.0364333894937896,
		-0.6744897501960817,
		-0.38532046640756773,
		-0.12566134685507402,
		0.12566134685507402,
		0.38532046640756773,
		0.6744897501960817,
		1.0364333894937896,
		1.6448536269514729,
		3.090232306167813,
		4.264890793922602,
		5.199337582187471,
	}
	for i, v := range p {
		got := NormalQuantile(v)
		if !floats.EqualWithinAbsOrRel(got, ans[i], 1e-10, 1e-10) {
			t.Errorf("Quantile mismatch. Case %d, want: %v, got: %v", i, ans[i], got)
		}
	}
}

var nqtmp float64

func BenchmarkNormalQuantile(b *testing.B) {
	ps := make([]float64, 1000) // ensure there are small values
	floats.Span(ps, 0, 1)
	for i := 0; i < b.N; i++ {
		for _, v := range ps {
			nqtmp = NormalQuantile(v)
		}
	}
}
