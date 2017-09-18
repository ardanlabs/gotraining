// Copyright ©2016 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package distuv

import (
	"math/rand"
	"sort"
	"testing"

	"gonum.org/v1/gonum/floats"
)

func TestChiSquaredProb(t *testing.T) {
	for _, test := range []struct {
		x, k, want float64
	}{
		{10, 3, 0.0085003666025203432},
		{2.3, 3, 0.19157345407042367},
		{0.8, 0.2, 0.080363259903912673},
	} {
		pdf := ChiSquared{test.k, nil}.Prob(test.x)
		if !floats.EqualWithinAbsOrRel(pdf, test.want, 1e-10, 1e-10) {
			t.Errorf("Pdf mismatch, x = %v, K = %v. Got %v, want %v", test.x, test.k, pdf, test.want)
		}
	}
}

func TestChiSquaredCDF(t *testing.T) {
	for _, test := range []struct {
		x, k, want float64
	}{
		// Values calculated with scipy.stats.chi2.cdf
		{0, 1, 0},
		{0.01, 5, 5.3002700426865167e-07},
		{0.05, 3, 0.002929332764619924},
		{0.5, 2, 0.22119921692859512},
		{0.95, 3, 0.1866520918701263},
		{0.99, 5, 0.036631697220869196},
		{1, 1, 0.68268949213708596},
		{1.5, 4, 0.17335853270322427},
		{10, 10, 0.55950671493478743},
		{25, 15, 0.95005656637357172},
	} {
		cdf := ChiSquared{test.k, nil}.CDF(test.x)
		if !floats.EqualWithinAbsOrRel(cdf, test.want, 1e-10, 1e-10) {
			t.Errorf("CDF mismatch, x = %v, K = %v. Got %v, want %v", test.x, test.k, cdf, test.want)
		}
	}
}

func TestChiSquared(t *testing.T) {
	src := rand.New(rand.NewSource(1))
	for i, b := range []ChiSquared{
		{3, src},
		{1.5, src},
		{0.9, src},
	} {
		testChiSquared(t, b, i)
	}
}

func testChiSquared(t *testing.T, c ChiSquared, i int) {
	tol := 1e-2
	const n = 1e5
	const bins = 50
	x := make([]float64, n)
	generateSamples(x, c)
	sort.Float64s(x)

	testRandLogProbContinuous(t, i, 0, x, c, tol, bins)
	checkMean(t, i, x, c, tol)
	checkVarAndStd(t, i, x, c, tol)
	checkExKurtosis(t, i, x, c, 7e-2)
	checkProbContinuous(t, i, x, c, 1e-3)
	checkQuantileCDFSurvival(t, i, x, c, 1e-2)
}
