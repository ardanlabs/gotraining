// Copyright ©2014 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package distuv

import (
	"fmt"
	"math"
	"testing"

	"gonum.org/v1/gonum/diff/fd"
	"gonum.org/v1/gonum/floats"
)

type univariateProbPoint struct {
	loc     float64
	logProb float64
	cumProb float64
	prob    float64
}

type UniProbDist interface {
	Prob(float64) float64
	CDF(float64) float64
	LogProb(float64) float64
	Quantile(float64) float64
	Survival(float64) float64
}

func absEq(a, b float64) bool {
	// This is expressed as the inverse to catch the
	// case a = Inf and b = Inf of the same sign.
	return !(math.Abs(a-b) > 1e-14)
}

// TODO: Implement a better test for Quantile
func testDistributionProbs(t *testing.T, dist UniProbDist, name string, pts []univariateProbPoint) {
	for _, pt := range pts {
		logProb := dist.LogProb(pt.loc)
		if !absEq(logProb, pt.logProb) {
			t.Errorf("Log probability doesnt match for "+name+". Expected %v. Found %v", pt.logProb, logProb)
		}
		prob := dist.Prob(pt.loc)
		if !absEq(prob, pt.prob) {
			t.Errorf("Probability doesn't match for "+name+". Expected %v. Found %v", pt.prob, prob)
		}
		cumProb := dist.CDF(pt.loc)
		if !absEq(cumProb, pt.cumProb) {
			t.Errorf("Cumulative Probability doesn't match for "+name+". Expected %v. Found %v", pt.cumProb, cumProb)
		}
		if !absEq(dist.Survival(pt.loc), 1-pt.cumProb) {
			t.Errorf("Survival doesn't match for %v. Expected %v, Found %v", name, 1-pt.cumProb, dist.Survival(pt.loc))
		}
		if pt.prob != 0 {
			if math.Abs(dist.Quantile(pt.cumProb)-pt.loc) > 1e-4 {
				fmt.Println("true =", pt.loc)
				fmt.Println("calculated=", dist.Quantile(pt.cumProb))
				t.Errorf("Quantile doesn't match for "+name+", loc =  %v", pt.loc)
			}
		}
	}
}

type ConjugateUpdater interface {
	NumParameters() int
	parameters([]Parameter) []Parameter

	NumSuffStat() int
	SuffStat([]float64, []float64, []float64) float64
	ConjugateUpdate([]float64, float64, []float64)

	Rand() float64
}

func testConjugateUpdate(t *testing.T, newFittable func() ConjugateUpdater) {
	for i, test := range []struct {
		samps   []float64
		weights []float64
	}{
		{
			samps:   randn(newFittable(), 10),
			weights: nil,
		},
		{
			samps:   randn(newFittable(), 10),
			weights: ones(10),
		},
		{
			samps:   randn(newFittable(), 10),
			weights: randn(&Exponential{Rate: 1}, 10),
		},
	} {
		// ensure that conjugate produces the same result both incrementally and all at once
		incDist := newFittable()
		stats := make([]float64, incDist.NumSuffStat())
		prior := make([]float64, incDist.NumParameters())
		for j := range test.samps {
			var incWeights, allWeights []float64
			if test.weights != nil {
				incWeights = test.weights[j : j+1]
				allWeights = test.weights[0 : j+1]
			}
			nsInc := incDist.SuffStat(stats, test.samps[j:j+1], incWeights)
			incDist.ConjugateUpdate(stats, nsInc, prior)

			allDist := newFittable()
			nsAll := allDist.SuffStat(stats, test.samps[0:j+1], allWeights)
			allDist.ConjugateUpdate(stats, nsAll, make([]float64, allDist.NumParameters()))
			if !parametersEqual(incDist.parameters(nil), allDist.parameters(nil), 1e-12) {
				t.Errorf("prior doesn't match after incremental update for (%d, %d). Incremental is %v, all at once is %v", i, j, incDist, allDist)
			}

			if test.weights == nil {
				onesDist := newFittable()
				nsOnes := onesDist.SuffStat(stats, test.samps[0:j+1], ones(j+1))
				onesDist.ConjugateUpdate(stats, nsOnes, make([]float64, onesDist.NumParameters()))
				if !parametersEqual(onesDist.parameters(nil), incDist.parameters(nil), 1e-14) {
					t.Errorf("nil and uniform weighted prior doesn't match for incremental update for (%d, %d). Uniform weighted is %v, nil is %v", i, j, onesDist, incDist)
				}
				if !parametersEqual(onesDist.parameters(nil), allDist.parameters(nil), 1e-14) {
					t.Errorf("nil and uniform weighted prior doesn't match for all at once update for (%d, %d). Uniform weighted is %v, nil is %v", i, j, onesDist, incDist)
				}
			}
		}
	}
}

// randn generates a specified number of random samples
func randn(dist Rander, n int) []float64 {
	x := make([]float64, n)
	for i := range x {
		x[i] = dist.Rand()
	}
	return x
}

func ones(n int) []float64 {
	x := make([]float64, n)
	for i := range x {
		x[i] = 1
	}
	return x
}

func parametersEqual(p1, p2 []Parameter, tol float64) bool {
	for i, p := range p1 {
		if p.Name != p2[i].Name {
			return false
		}
		if math.Abs(p.Value-p2[i].Value) > tol {
			return false
		}
	}
	return true
}

type derivParamTester interface {
	LogProb(x float64) float64
	Score(deriv []float64, x float64) []float64
	Quantile(p float64) float64
	NumParameters() int
	parameters([]Parameter) []Parameter
	setParameters([]Parameter)
}

func testDerivParam(t *testing.T, d derivParamTester) {
	// Tests that the derivative matches for a number of different quantiles
	// along the distribution.
	nTest := 10
	quantiles := make([]float64, nTest)
	floats.Span(quantiles, 0.1, 0.9)

	deriv := make([]float64, d.NumParameters())
	fdDeriv := make([]float64, d.NumParameters())

	initParams := d.parameters(nil)
	init := make([]float64, d.NumParameters())
	for i, v := range initParams {
		init[i] = v.Value
	}
	for _, v := range quantiles {
		d.setParameters(initParams)
		x := d.Quantile(v)
		d.Score(deriv, x)
		f := func(p []float64) float64 {
			params := d.parameters(nil)
			for i, v := range p {
				params[i].Value = v
			}
			d.setParameters(params)
			return d.LogProb(x)
		}
		fd.Gradient(fdDeriv, f, init, nil)
		if !floats.EqualApprox(deriv, fdDeriv, 1e-6) {
			t.Fatal("Derivative mismatch. Want", fdDeriv, ", got", deriv, ".")
		}
		d.setParameters(initParams)
		d2 := d.Score(nil, x)
		if !floats.EqualApprox(d2, deriv, 1e-14) {
			t.Errorf("Derivative mismatch when input nil Want %v, got %v", d2, deriv)
		}
	}
}
