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

func TestGamma(t *testing.T) {
	// Values a comparison with scipy
	for _, test := range []struct {
		x, alpha, want float64
	}{
		{0.9, 0.1, 0.046986817861555757},
		{0.9, 0.01, 0.0045384353289090401},
		{0.45, 0.01, 0.014137035997241795},
	} {
		pdf := Gamma{Alpha: test.alpha, Beta: 1}.Prob(test.x)
		if !floats.EqualWithinAbsOrRel(pdf, test.want, 1e-10, 1e-10) {
			t.Errorf("Pdf mismatch. Got %v, want %v", pdf, test.want)
		}
	}
	src := rand.New(rand.NewSource(1))
	for i, g := range []Gamma{

		{Alpha: 0.5, Beta: 0.8, Source: src},
		{Alpha: 0.9, Beta: 6, Source: src},
		{Alpha: 0.9, Beta: 500, Source: src},

		{Alpha: 1, Beta: 1, Source: src},

		{Alpha: 1.6, Beta: 0.4, Source: src},
		{Alpha: 2.6, Beta: 1.5, Source: src},
		{Alpha: 5.6, Beta: 0.5, Source: src},
		{Alpha: 30, Beta: 1.7, Source: src},
		{Alpha: 30.2, Beta: 1.7, Source: src},
	} {
		testGamma(t, g, i)
	}
}

func testGamma(t *testing.T, f Gamma, i int) {
	// TODO(btracey): Replace this when Gamma implements FullDist.
	tol := 1e-2
	const n = 1e5
	const bins = 50
	x := make([]float64, n)
	generateSamples(x, f)
	sort.Float64s(x)

	testRandLogProbContinuous(t, i, 0, x, f, tol, bins)
	checkMean(t, i, x, f, tol)
	checkVarAndStd(t, i, x, f, 2e-2)
	checkExKurtosis(t, i, x, f, 2e-1)
	checkProbContinuous(t, i, x, f, 1e-3)
	checkQuantileCDFSurvival(t, i, x, f, 5e-2)
}
