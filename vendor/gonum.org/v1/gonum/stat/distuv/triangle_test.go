// Copyright ©2017 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package distuv

import (
	"math"
	"math/rand"
	"sort"
	"testing"
)

func TestTriangleConstraint(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The constraints were violated, but not caught")
		}
	}()

	// test b < a
	NewTriangle(3, 1, 2)
	// test c > b
	NewTriangle(1, 2, 3)
}

func TestTriangle(t *testing.T) {
	src := rand.New(rand.NewSource(1))
	for i, test := range []struct {
		a, b, c float64
	}{
		{
			a: 0.0,
			b: 1.0,
			c: 0.5,
		},
		{
			a: 0.1,
			b: 0.3,
			c: 0.2,
		},
		{
			a: 1.0,
			b: 2.0,
			c: 1.5,
		},
		{
			a: 0.0,
			b: 1.0,
			c: 0.0,
		},
	} {
		f := NewTriangle(test.a, test.b, test.c)
		f.Source = src
		tol := 1e-2
		const n = 1e5
		x := make([]float64, n)
		generateSamples(x, f)
		sort.Float64s(x)

		checkMean(t, i, x, f, tol)
		checkVarAndStd(t, i, x, f, tol)
		checkEntropy(t, i, x, f, tol)
		checkExKurtosis(t, i, x, f, tol)
		checkSkewness(t, i, x, f, 5e-2)
		checkMedian(t, i, x, f, tol)
		checkQuantileCDFSurvival(t, i, x, f, tol)
		checkProbContinuous(t, i, x, f, 1e-10)
		checkProbQuantContinuous(t, i, x, f, tol)
	}
}

func TestTriangleProb(t *testing.T) {
	pts := []univariateProbPoint{
		{
			loc:     0.5,
			prob:    0,
			cumProb: 0,
			logProb: math.Inf(-1),
		},
		{
			loc:     1,
			prob:    0,
			cumProb: 0,
			logProb: math.Inf(-1),
		},
		{
			loc:     2,
			prob:    1.0,
			cumProb: 0.5,
			logProb: 0,
		},
		{
			loc:     3,
			prob:    0,
			cumProb: 1,
			logProb: math.Inf(-1),
		},
	}
	testDistributionProbs(t, NewTriangle(1, 3, 2), "Standard 1,2,3 Triangle", pts)
}
