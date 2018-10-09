// Copyright ©2014 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package distuv

import (
	"math"

	"golang.org/x/exp/rand"
)

// UnitUniform is an instantiation of the uniform distribution with Min = 0
// and Max = 1.
var UnitUniform = Uniform{Min: 0, Max: 1}

// Uniform represents a continuous uniform distribution (https://en.wikipedia.org/wiki/Uniform_distribution_%28continuous%29).
type Uniform struct {
	Min float64
	Max float64
	Src rand.Source
}

// CDF computes the value of the cumulative density function at x.
func (u Uniform) CDF(x float64) float64 {
	if x < u.Min {
		return 0
	}
	if x > u.Max {
		return 1
	}
	return (x - u.Min) / (u.Max - u.Min)
}

// Uniform doesn't have any of the DLogProbD? because the derivative is 0 everywhere
// except where it's undefined

// Entropy returns the entropy of the distribution.
func (u Uniform) Entropy() float64 {
	return math.Log(u.Max - u.Min)
}

// ExKurtosis returns the excess kurtosis of the distribution.
func (Uniform) ExKurtosis() float64 {
	return -6.0 / 5.0
}

// Uniform doesn't have Fit because it's a bad idea to fit a uniform from data.

// LogProb computes the natural logarithm of the value of the probability density function at x.
func (u Uniform) LogProb(x float64) float64 {
	if x < u.Min {
		return math.Inf(-1)
	}
	if x > u.Max {
		return math.Inf(-1)
	}
	return -math.Log(u.Max - u.Min)
}

// MarshalParameters implements the ParameterMarshaler interface
func (u Uniform) MarshalParameters(p []Parameter) {
	if len(p) != u.NumParameters() {
		panic("uniform: improper parameter length")
	}
	p[0].Name = "Min"
	p[0].Value = u.Min
	p[1].Name = "Max"
	p[1].Value = u.Max
}

// Mean returns the mean of the probability distribution.
func (u Uniform) Mean() float64 {
	return (u.Max + u.Min) / 2
}

// Median returns the median of the probability distribution.
func (u Uniform) Median() float64 {
	return (u.Max + u.Min) / 2
}

// Uniform doesn't have a mode because it's any value in the distribution

// NumParameters returns the number of parameters in the distribution.
func (Uniform) NumParameters() int {
	return 2
}

// Prob computes the value of the probability density function at x.
func (u Uniform) Prob(x float64) float64 {
	if x < u.Min {
		return 0
	}
	if x > u.Max {
		return 0
	}
	return 1 / (u.Max - u.Min)
}

// Quantile returns the inverse of the cumulative probability distribution.
func (u Uniform) Quantile(p float64) float64 {
	if p < 0 || p > 1 {
		panic(badPercentile)
	}
	return p*(u.Max-u.Min) + u.Min
}

// Rand returns a random sample drawn from the distribution.
func (u Uniform) Rand() float64 {
	var rnd float64
	if u.Src == nil {
		rnd = rand.Float64()
	} else {
		rnd = rand.New(u.Src).Float64()
	}
	return rnd*(u.Max-u.Min) + u.Min
}

// Skewness returns the skewness of the distribution.
func (Uniform) Skewness() float64 {
	return 0
}

// StdDev returns the standard deviation of the probability distribution.
func (u Uniform) StdDev() float64 {
	return math.Sqrt(u.Variance())
}

// Survival returns the survival function (complementary CDF) at x.
func (u Uniform) Survival(x float64) float64 {
	if x < u.Min {
		return 1
	}
	if x > u.Max {
		return 0
	}
	return (u.Max - x) / (u.Max - u.Min)
}

// UnmarshalParameters implements the ParameterMarshaler interface
func (u *Uniform) UnmarshalParameters(p []Parameter) {
	if len(p) != u.NumParameters() {
		panic("uniform: incorrect number of parameters to set")
	}
	if p[0].Name != "Min" {
		panic("uniform: " + panicNameMismatch)
	}
	if p[1].Name != "Max" {
		panic("uniform: " + panicNameMismatch)
	}

	u.Min = p[0].Value
	u.Max = p[1].Value
}

// Variance returns the variance of the probability distribution.
func (u Uniform) Variance() float64 {
	return 1.0 / 12.0 * (u.Max - u.Min) * (u.Max - u.Min)
}
