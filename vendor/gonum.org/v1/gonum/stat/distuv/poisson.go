// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package distuv

import (
	"math"

	"golang.org/x/exp/rand"

	"gonum.org/v1/gonum/mathext"
)

// Poisson implements the Poisson distribution, a discrete probability distribution
// that expresses the probability of a given number of events occurring in a fixed
// interval.
// The poisson distribution has density function:
//  f(k) = λ^k / k! e^(-λ)
// For more information, see https://en.wikipedia.org/wiki/Poisson_distribution.
type Poisson struct {
	// Lambda is the average number of events in an interval.
	// Lambda must be greater than 0.
	Lambda float64

	Src rand.Source
}

// CDF computes the value of the cumulative distribution function at x.
func (p Poisson) CDF(x float64) float64 {
	if x < 0 {
		return 0
	}
	return mathext.GammaIncComp(math.Floor(x+1), p.Lambda)
}

// ExKurtosis returns the excess kurtosis of the distribution.
func (p Poisson) ExKurtosis() float64 {
	return 1 / p.Lambda
}

// LogProb computes the natural logarithm of the value of the probability
// density function at x.
func (p Poisson) LogProb(x float64) float64 {
	if x < 0 || math.Floor(x) != x {
		return math.Inf(-1)
	}
	lg, _ := math.Lgamma(math.Floor(x) + 1)
	return x*math.Log(p.Lambda) - p.Lambda - lg
}

// Mean returns the mean of the probability distribution.
func (p Poisson) Mean() float64 {
	return p.Lambda
}

// NumParameters returns the number of parameters in the distribution.
func (Poisson) NumParameters() int {
	return 1
}

// Prob computes the value of the probability density function at x.
func (p Poisson) Prob(x float64) float64 {
	return math.Exp(p.LogProb(x))
}

// Rand returns a random sample drawn from the distribution.
func (p Poisson) Rand() float64 {
	// NUMERICAL RECIPES IN C: THE ART OF SCIENTIFIC COMPUTING (ISBN 0-521-43108-5)
	// p. 294
	// <http://www.aip.de/groups/soe/local/numres/bookcpdf/c7-3.pdf>

	rnd := rand.ExpFloat64
	var rng *rand.Rand
	if p.Src != nil {
		rng = rand.New(p.Src)
		rnd = rng.ExpFloat64
	}

	if p.Lambda < 10.0 {
		// Use direct method.
		var em float64
		t := 0.0
		for {
			t += rnd()
			if t >= p.Lambda {
				break
			}
			em++
		}
		return em
	}
	// Use rejection method.
	rnd = rand.Float64
	if rng != nil {
		rnd = rng.Float64
	}
	sq := math.Sqrt(2.0 * p.Lambda)
	alxm := math.Log(p.Lambda)
	lg, _ := math.Lgamma(p.Lambda + 1)
	g := p.Lambda*alxm - lg
	for {
		var em, y float64
		for {
			y = math.Tan(math.Pi * rnd())
			em = sq*y + p.Lambda
			if em >= 0 {
				break
			}
		}
		em = math.Floor(em)
		lg, _ = math.Lgamma(em + 1)
		t := 0.9 * (1.0 + y*y) * math.Exp(em*alxm-lg-g)
		if rnd() <= t {
			return em
		}
	}
}

// Skewness returns the skewness of the distribution.
func (p Poisson) Skewness() float64 {
	return 1 / math.Sqrt(p.Lambda)
}

// StdDev returns the standard deviation of the probability distribution.
func (p Poisson) StdDev() float64 {
	return math.Sqrt(p.Variance())
}

// Survival returns the survival function (complementary CDF) at x.
func (p Poisson) Survival(x float64) float64 {
	return 1 - p.CDF(x)
}

// Variance returns the variance of the probability distribution.
func (p Poisson) Variance() float64 {
	return p.Lambda
}
