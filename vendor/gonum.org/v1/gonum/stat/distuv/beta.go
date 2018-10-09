// Copyright ©2016 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package distuv

import (
	"math"

	"golang.org/x/exp/rand"

	"gonum.org/v1/gonum/mathext"
)

// Beta implements the Beta distribution, a two-parameter continuous distribution
// with support between 0 and 1.
//
// The beta distribution has density function
//  x^(α-1) * (1-x)^(β-1) * Γ(α+β) / (Γ(α)*Γ(β))
//
// For more information, see https://en.wikipedia.org/wiki/Beta_distribution
type Beta struct {
	// Alpha is the left shape parameter of the distribution. Alpha must be greater
	// than 0.
	Alpha float64
	// Beta is the right shape parameter of the distribution. Beta must be greater
	// than 0.
	Beta float64

	Src rand.Source
}

// CDF computes the value of the cumulative distribution function at x.
func (b Beta) CDF(x float64) float64 {
	if x <= 0 {
		return 0
	}
	if x >= 1 {
		return 1
	}
	return mathext.RegIncBeta(b.Alpha, b.Beta, x)
}

// Entropy returns the differential entropy of the distribution.
func (b Beta) Entropy() float64 {
	if b.Alpha <= 0 || b.Beta <= 0 {
		panic("beta: negative parameters")
	}
	return mathext.Lbeta(b.Alpha, b.Beta) - (b.Alpha-1)*mathext.Digamma(b.Alpha) -
		(b.Beta-1)*mathext.Digamma(b.Beta) + (b.Alpha+b.Beta-2)*mathext.Digamma(b.Alpha+b.Beta)
}

// ExKurtosis returns the excess kurtosis of the distribution.
func (b Beta) ExKurtosis() float64 {
	num := 6 * ((b.Alpha-b.Beta)*(b.Alpha-b.Beta)*(b.Alpha+b.Beta+1) - b.Alpha*b.Beta*(b.Alpha+b.Beta+2))
	den := b.Alpha * b.Beta * (b.Alpha + b.Beta + 2) * (b.Alpha + b.Beta + 3)
	return num / den
}

// LogProb computes the natural logarithm of the value of the probability
// density function at x.
func (b Beta) LogProb(x float64) float64 {
	if x < 0 || x > 1 {
		return math.Inf(-1)
	}

	if b.Alpha <= 0 || b.Beta <= 0 {
		panic("beta: negative parameters")
	}

	lab, _ := math.Lgamma(b.Alpha + b.Beta)
	la, _ := math.Lgamma(b.Alpha)
	lb, _ := math.Lgamma(b.Beta)
	return lab - la - lb + (b.Alpha-1)*math.Log(x) + (b.Beta-1)*math.Log(1-x)
}

// Mean returns the mean of the probability distribution.
func (b Beta) Mean() float64 {
	return b.Alpha / (b.Alpha + b.Beta)
}

// Mode returns the mode of the distribution.
//
// Mode returns NaN if either parameter is less than or equal to 1 as a special case.
func (b Beta) Mode() float64 {
	if b.Alpha <= 1 || b.Beta <= 1 {
		return math.NaN()
	}
	return (b.Alpha - 1) / (b.Alpha + b.Beta - 2)
}

// NumParameters returns the number of parameters in the distribution.
func (b Beta) NumParameters() int {
	return 2
}

// Prob computes the value of the probability density function at x.
func (b Beta) Prob(x float64) float64 {
	return math.Exp(b.LogProb(x))
}

// Quantile returns the inverse of the cumulative distribution function.
func (b Beta) Quantile(p float64) float64 {
	if p < 0 || p > 1 {
		panic(badPercentile)
	}
	return mathext.InvRegIncBeta(b.Alpha, b.Beta, p)
}

// Rand returns a random sample drawn from the distribution.
func (b Beta) Rand() float64 {
	ga := Gamma{Alpha: b.Alpha, Beta: 1, Src: b.Src}.Rand()
	gb := Gamma{Alpha: b.Beta, Beta: 1, Src: b.Src}.Rand()
	return ga / (ga + gb)
}

// StdDev returns the standard deviation of the probability distribution.
func (b Beta) StdDev() float64 {
	return math.Sqrt(b.Variance())
}

// Survival returns the survival function (complementary CDF) at x.
func (b Beta) Survival(x float64) float64 {
	switch {
	case x <= 0:
		return 1
	case x >= 1:
		return 0
	}
	return mathext.RegIncBeta(b.Beta, b.Alpha, 1-x)
}

// Variance returns the variance of the probability distribution.
func (b Beta) Variance() float64 {
	return b.Alpha * b.Beta / ((b.Alpha + b.Beta) * (b.Alpha + b.Beta) * (b.Alpha + b.Beta + 1))
}
