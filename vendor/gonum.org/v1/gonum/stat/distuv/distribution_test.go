// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package distuv

import (
	"math"
	"sort"
	"testing"

	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/integrate/quad"
	"gonum.org/v1/gonum/stat"
)

type meaner interface {
	Mean() float64
}

type quantiler interface {
	Quantile(float64) float64
}

type medianer interface {
	quantiler
	Median() float64
}

type varStder interface {
	StdDev() float64
	Variance() float64
}

type entropyer interface {
	LogProber
	Entropy() float64
}

type exKurtosiser interface {
	ExKurtosis() float64
	Mean() float64
}

type skewnesser interface {
	StdDev() float64
	Mean() float64
	Skewness() float64
}

type cumulanter interface {
	Quantiler
	CDF(x float64) float64
	Survival(x float64) float64
}

func generateSamples(x []float64, r Rander) {
	for i := range x {
		x[i] = r.Rand()
	}
}

type probLogprober interface {
	Prob(x float64) float64
	LogProb(x float64) float64
}

type cumulantProber interface {
	cumulanter
	probLogprober
}

func checkMean(t *testing.T, i int, x []float64, m meaner, tol float64) {
	mean := stat.Mean(x, nil)
	if !floats.EqualWithinAbsOrRel(mean, m.Mean(), tol, tol) {
		t.Errorf("Mean mismatch case %v: want: %v, got: %v", i, mean, m.Mean())
	}
}

func checkMedian(t *testing.T, i int, x []float64, m medianer, tol float64) {
	median := stat.Quantile(0.5, stat.Empirical, x, nil)
	if !floats.EqualWithinAbsOrRel(median, m.Median(), tol, tol) {
		t.Errorf("Median mismatch case %v: want: %v, got: %v", i, median, m.Median())
	}
}

func checkVarAndStd(t *testing.T, i int, x []float64, v varStder, tol float64) {
	variance := stat.Variance(x, nil)
	if !floats.EqualWithinAbsOrRel(variance, v.Variance(), tol, tol) {
		t.Errorf("Variance mismatch case %v: want: %v, got: %v", i, variance, v.Variance())
	}
	std := math.Sqrt(variance)
	if !floats.EqualWithinAbsOrRel(std, v.StdDev(), tol, tol) {
		t.Errorf("StdDev mismatch case %v: want: %v, got: %v", i, std, v.StdDev())
	}
}

func checkEntropy(t *testing.T, i int, x []float64, e entropyer, tol float64) {
	tmp := make([]float64, len(x))
	for i, v := range x {
		tmp[i] = -e.LogProb(v)
	}
	entropy := stat.Mean(tmp, nil)
	if !floats.EqualWithinAbsOrRel(entropy, e.Entropy(), tol, tol) {
		t.Errorf("Entropy mismatch case %v: want: %v, got: %v", i, entropy, e.Entropy())
	}
}

func checkExKurtosis(t *testing.T, i int, x []float64, e exKurtosiser, tol float64) {
	mean := e.Mean()
	tmp := make([]float64, len(x))
	for i, x := range x {
		tmp[i] = math.Pow(x-mean, 4)
	}
	variance := stat.Variance(x, nil)
	mu4 := stat.Mean(tmp, nil)
	kurtosis := mu4/(variance*variance) - 3
	if !floats.EqualWithinAbsOrRel(kurtosis, e.ExKurtosis(), tol, tol) {
		t.Errorf("ExKurtosis mismatch case %v: want: %v, got: %v", i, kurtosis, e.ExKurtosis())
	}
}

func checkSkewness(t *testing.T, i int, x []float64, s skewnesser, tol float64) {
	mean := s.Mean()
	std := s.StdDev()
	tmp := make([]float64, len(x))
	for i, v := range x {
		tmp[i] = math.Pow(v-mean, 3)
	}
	mu3 := stat.Mean(tmp, nil)
	skewness := mu3 / math.Pow(std, 3)
	if !floats.EqualWithinAbsOrRel(skewness, s.Skewness(), tol, tol) {
		t.Errorf("Skewness mismatch case %v: want: %v, got: %v", i, skewness, s.Skewness())
	}
}

func checkQuantileCDFSurvival(t *testing.T, i int, xs []float64, c cumulanter, tol float64) {
	// Quantile, CDF, and survival check.
	for i, p := range []float64{0.1, 0.25, 0.5, 0.75, 0.9} {
		x := c.Quantile(p)
		cdf := c.CDF(x)
		estCDF := stat.CDF(x, stat.Empirical, xs, nil)
		if !floats.EqualWithinAbsOrRel(cdf, estCDF, tol, tol) {
			t.Errorf("CDF mismatch case %v: want: %v, got: %v", i, estCDF, cdf)
		}
		if !floats.EqualWithinAbsOrRel(cdf, p, tol, tol) {
			t.Errorf("Quantile/CDF mismatch case %v: want: %v, got: %v", i, p, cdf)
		}
		if math.Abs(1-cdf-c.Survival(x)) > 1e-14 {
			t.Errorf("Survival/CDF mismatch case %v: want: %v, got: %v", i, 1-cdf, c.Survival(x))
		}
	}
}

func checkProbContinuous(t *testing.T, i int, x []float64, p probLogprober, tol float64) {
	// Check that the PDF is consistent (integrates to 1).
	q := quad.Fixed(p.Prob, math.Inf(-1), math.Inf(1), 1000000, nil, 0)
	if math.Abs(q-1) > tol {
		t.Errorf("Probability distribution doesn't integrate to 1. Case %v: Got %v", i, q)
	}

	// Check that PDF and LogPDF are consistent.
	for i, v := range x {
		if math.Abs(math.Log(p.Prob(v))-p.LogProb(v)) > 1e-14 {
			t.Errorf("Prob and LogProb mismatch case %v at %v: want %v, got %v", i, v, math.Log(v), p.LogProb(v))
			break
		}
	}
}

// checkProbQuantContinuous checks that the Prob, Rand, and Quantile are all consistent.
// checkProbContinuous only checks that Prob is a valid distribution (integrates
// to 1 and greater than 0). However, this is also true if the PDF of a different
// distribution is used. This checks that PDF is also consistent with the
// CDF implementation and the random samples.
func checkProbQuantContinuous(t *testing.T, i int, xs []float64, c cumulantProber, tol float64) {
	ps := make([]float64, 101)
	floats.Span(ps, 0, 1)

	var xp, x float64
	for i, p := range ps {
		x = c.Quantile(p)
		if p == 0 {
			xp = x
			if floats.Min(xs) < x {
				t.Errorf("Sample of x less than Quantile(0). Case %v.", i)
				break
			}
			continue
		}
		if p == 1 {
			if floats.Max(xs) > x {
				t.Errorf("Sample of x greater than Quantile(1). Case %v.", i)
				break
			}
		}

		// The integral of the PDF between xp and x should be the difference in
		// the quantiles.
		q := quad.Fixed(c.Prob, xp, x, 1000, nil, 0)
		if math.Abs(q-(p-ps[i-1])) > 1e-5 {
			t.Errorf("Integral of PDF doesn't match quantile. Case %v. Want %v, got %v.", i, p-ps[i-1], q)
			break
		}

		pEst := stat.CDF(x, stat.Empirical, xs, nil)
		if math.Abs(pEst-p) > tol {
			t.Errorf("Empirical CDF doesn't match quantile. Case %v.", i)
		}
		xp = x
	}
}

// checkProbDiscrete confirms that PDF and Rand are consistent for discrete distributions.
func checkProbDiscrete(t *testing.T, i int, xs []float64, p probLogprober, tol float64) {
	// Make a map of all of the unique samples.
	m := make(map[float64]int)
	for _, v := range xs {
		m[v]++
	}
	for x, count := range m {
		prob := float64(count) / float64(len(xs))
		if math.Abs(prob-p.Prob(x)) > tol {
			t.Errorf("PDF mismatch case %v at %v: want %v, got %v", i, x, prob, p.Prob(x))
		}
		if math.Abs(math.Log(p.Prob(x))-p.LogProb(x)) > 1e-14 {
			t.Errorf("Prob and LogProb mismatch case %v at %v: want %v, got %v", i, x, math.Log(x), p.LogProb(x))
		}
	}
}

// dist is a type that implements the standard set of routines.
type fullDist interface {
	CDF(x float64) float64
	Entropy() float64
	ExKurtosis() float64
	LogProb(x float64) float64
	Mean() float64
	Median() float64
	NumParameters() int
	Prob(x float64) float64
	Quantile(p float64) float64
	Rand() float64
	Skewness() float64
	StdDev() float64
	Survival(x float64) float64
	Variance() float64
}

// testFullDist tests all of the functions of a fullDist.
func testFullDist(t *testing.T, f fullDist, i int, continuous bool) {
	tol := 1e-2
	const n = 1e6
	x := make([]float64, n)
	generateSamples(x, f)
	sort.Float64s(x)

	checkMean(t, i, x, f, tol)
	checkVarAndStd(t, i, x, f, tol)
	checkEntropy(t, i, x, f, tol)
	checkExKurtosis(t, i, x, f, tol)
	checkSkewness(t, i, x, f, tol)
	if continuous {
		// In a discrete distribution, the median may not have positive probability.
		checkMedian(t, i, x, f, tol)
		// In a discrete distribution, the CDF and Quantile may not be perfect mappings.
		checkQuantileCDFSurvival(t, i, x, f, tol)
		// Integrate over the PDF
		checkProbContinuous(t, i, x, f, 1e-10)
		checkProbQuantContinuous(t, i, x, f, tol)
	} else {
		// Check against empirical PDF.
		checkProbDiscrete(t, i, x, f, tol)
	}
}

// testRandLogProb tests that LogProb and Rand give consistent results. This
// can be used when the distribution does not implement CDF.
func testRandLogProbContinuous(t *testing.T, i int, min float64, x []float64, f LogProber, tol float64, bins int) {
	for cdf := 1 / float64(bins); cdf <= 1-1/float64(bins); cdf += 1 / float64(bins) {
		// Get the estimated CDF from the samples
		pt := stat.Quantile(cdf, stat.Empirical, x, nil)

		prob := func(x float64) float64 {
			return math.Exp(f.LogProb(x))
		}
		// Integrate the PDF to find the CDF
		estCDF := quad.Fixed(prob, min, pt, 1000, nil, 0)
		if !floats.EqualWithinAbsOrRel(cdf, estCDF, tol, tol) {
			t.Errorf("Mismatch between integral of PDF and empirical CDF. Case %v. Want %v, got %v", i, cdf, estCDF)
		}
	}
}
