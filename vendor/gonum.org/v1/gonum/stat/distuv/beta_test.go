// Copyright ©2016 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package distuv

import (
	"math"
	"math/rand"
	"sort"
	"testing"

	"gonum.org/v1/gonum/floats"
)

func TestBetaProb(t *testing.T) {
	// Values a comparison with scipy
	for _, test := range []struct {
		x, alpha, beta, want float64
	}{
		{0.1, 2, 0.5, 0.079056941504209499},
		{0.5, 1, 5.1, 0.29740426605235754},
		{0.1, 0.5, 0.5, 1.0610329539459691},
		{1, 0.5, 0.5, math.Inf(1)},
		{-1, 0.5, 0.5, 0},
	} {
		pdf := Beta{Alpha: test.alpha, Beta: test.beta}.Prob(test.x)
		if !floats.EqualWithinAbsOrRel(pdf, test.want, 1e-10, 1e-10) {
			t.Errorf("Pdf mismatch. Got %v, want %v", pdf, test.want)
		}
	}
}

func TestBetaRand(t *testing.T) {
	src := rand.New(rand.NewSource(1))
	for i, b := range []Beta{
		{Alpha: 0.5, Beta: 0.5, Source: src},
		{Alpha: 5, Beta: 1, Source: src},
		{Alpha: 2, Beta: 2, Source: src},
		{Alpha: 2, Beta: 5, Source: src},
	} {
		testBeta(t, b, i)
	}
}

func testBeta(t *testing.T, b Beta, i int) {
	tol := 1e-2
	const n = 5e4
	const bins = 10
	x := make([]float64, n)
	generateSamples(x, b)
	sort.Float64s(x)

	testRandLogProbContinuous(t, i, 0, x, b, tol, bins)
	checkMean(t, i, x, b, tol)
	checkVarAndStd(t, i, x, b, tol)
	checkExKurtosis(t, i, x, b, 5e-2)
	checkProbContinuous(t, i, x, b, 1e-3)
	checkQuantileCDFSurvival(t, i, x, b, tol)
	checkProbQuantContinuous(t, i, x, b, tol)
}
