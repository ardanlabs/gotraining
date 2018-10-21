// Copyright ©2016 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package distuv

import (
	"math"

	"golang.org/x/exp/rand"

	"gonum.org/v1/gonum/mathext"
)

// Gamma implements the Gamma distribution, a two-parameter continuous distribution
// with support over the positive real numbers.
//
// The gamma distribution has density function
//  β^α / Γ(α) x^(α-1)e^(-βx)
//
// For more information, see https://en.wikipedia.org/wiki/Gamma_distribution
type Gamma struct {
	// Alpha is the shape parameter of the distribution. Alpha must be greater
	// than 0. If Alpha == 1, this is equivalent to an exponential distribution.
	Alpha float64
	// Beta is the rate parameter of the distribution. Beta must be greater than 0.
	// If Beta == 2, this is equivalent to a Chi-Squared distribution.
	Beta float64

	Src rand.Source
}

// CDF computes the value of the cumulative distribution function at x.
func (g Gamma) CDF(x float64) float64 {
	if x < 0 {
		return 0
	}
	return mathext.GammaInc(g.Alpha, g.Beta*x)
}

// ExKurtosis returns the excess kurtosis of the distribution.
func (g Gamma) ExKurtosis() float64 {
	return 6 / g.Alpha
}

// LogProb computes the natural logarithm of the value of the probability
// density function at x.
func (g Gamma) LogProb(x float64) float64 {
	if x <= 0 {
		return math.Inf(-1)
	}
	a := g.Alpha
	b := g.Beta
	lg, _ := math.Lgamma(a)
	return a*math.Log(b) - lg + (a-1)*math.Log(x) - b*x
}

// Mean returns the mean of the probability distribution.
func (g Gamma) Mean() float64 {
	return g.Alpha / g.Beta
}

// Mode returns the mode of the normal distribution.
//
// The mode is NaN in the special case where the Alpha (shape) parameter
// is less than 1.
func (g Gamma) Mode() float64 {
	if g.Alpha < 1 {
		return math.NaN()
	}
	return (g.Alpha - 1) / g.Beta
}

// NumParameters returns the number of parameters in the distribution.
func (Gamma) NumParameters() int {
	return 2
}

// Prob computes the value of the probability density function at x.
func (g Gamma) Prob(x float64) float64 {
	return math.Exp(g.LogProb(x))
}

// Quantile returns the inverse of the cumulative distribution function.
func (g Gamma) Quantile(p float64) float64 {
	if p < 0 || p > 1 {
		panic(badPercentile)
	}
	return mathext.GammaIncInv(g.Alpha, p) / g.Beta
}

// Rand returns a random sample drawn from the distribution.
//
// Rand panics if either alpha or beta is <= 0.
func (g Gamma) Rand() float64 {
	if g.Beta <= 0 {
		panic("gamma: beta <= 0")
	}

	unifrnd := rand.Float64
	exprnd := rand.ExpFloat64
	normrnd := rand.NormFloat64
	if g.Src != nil {
		rnd := rand.New(g.Src)
		unifrnd = rnd.Float64
		exprnd = rnd.ExpFloat64
		normrnd = rnd.NormFloat64
	}

	a := g.Alpha
	b := g.Beta
	switch {
	case a <= 0:
		panic("gamma: alpha < 0")
	case a == 1:
		// Generate from exponential
		return exprnd() / b
	case a < 0.3:
		// Generate using
		//  Liu, Chuanhai, Martin, Ryan and Syring, Nick. "Simulating from a
		//  gamma distribution with small shape parameter"
		//  https://arxiv.org/abs/1302.1884
		//   use this reference: http://link.springer.com/article/10.1007/s00180-016-0692-0

		// Algorithm adjusted to work in log space as much as possible.
		lambda := 1/a - 1
		lw := math.Log(a) - 1 - math.Log(1-a)
		lr := -math.Log(1 + math.Exp(lw))
		lc, _ := math.Lgamma(a + 1)
		for {
			e := exprnd()
			var z float64
			if e >= -lr {
				z = e + lr
			} else {
				z = -exprnd() / lambda
			}
			lh := lc - z - math.Exp(-z/a)
			var lEta float64
			if z >= 0 {
				lEta = lc - z
			} else {
				lEta = lc + lw + math.Log(lambda) + lambda*z
			}
			if lh-lEta > -exprnd() {
				return math.Exp(-z/a) / b
			}
		}
	case a >= 0.3 && a < 1:
		// Generate using:
		//  Kundu, Debasis, and Rameshwar D. Gupta. "A convenient way of generating
		//  gamma random variables using generalized exponential distribution."
		//  Computational Statistics & Data Analysis 51.6 (2007): 2796-2802.

		// TODO(btracey): Change to using Algorithm 3 if we can find the bug in
		// the implementation below.

		// Algorithm 2.
		alpha := g.Alpha
		a := math.Pow(1-expNegOneHalf, alpha) / (math.Pow(1-expNegOneHalf, alpha) + alpha*math.Exp(-1)/math.Pow(2, alpha))
		b := math.Pow(1-expNegOneHalf, alpha) + alpha/math.E/math.Pow(2, alpha)
		var x float64
		for {
			u := unifrnd()
			if u <= a {
				x = -2 * math.Log(1-math.Pow(u*b, 1/alpha))
			} else {
				x = -math.Log(math.Pow(2, alpha) / alpha * b * (1 - u))
			}
			v := unifrnd()
			if x <= 1 {
				if v <= math.Pow(x, alpha-1)*math.Exp(-x/2)/(math.Pow(2, alpha-1)*math.Pow(1-math.Exp(-x/2), alpha-1)) {
					break
				}
			} else {
				if v <= math.Pow(x, alpha-1) {
					break
				}
			}
		}
		return x / g.Beta

		/*
			//  Algorithm 3.
			d := 1.0334 - 0.0766*math.Exp(2.2942*alpha)
			a := math.Pow(2, alpha) * math.Pow(1-math.Exp(-d/2), alpha)
			b := alpha * math.Pow(d, alpha-1) * math.Exp(-d)
			c := a + b
			var x float64
			for {
				u := unifrnd()
				if u <= a/(a+b) {
					x = -2 * math.Log(1-math.Pow(c*u, 1/a)/2)
				} else {
					x = -math.Log(c * (1 - u) / (alpha * math.Pow(d, alpha-1)))
				}
				v := unifrnd()
				if x <= d {
					if v <= (math.Pow(x, alpha-1)*math.Exp(-x/2))/(math.Pow(2, alpha-1)*math.Pow(1-math.Exp(-x/2), alpha-1)) {
						break
					}
				} else {
					if v <= math.Pow(d/x, 1-alpha) {
						break
					}
				}
			}
			return x / g.Beta
		*/
	case a > 1:
		// Generate using:
		//  Marsaglia, George, and Wai Wan Tsang. "A simple method for generating
		//  gamma variables." ACM Transactions on Mathematical Software (TOMS)
		//  26.3 (2000): 363-372.
		d := a - 1.0/3
		c := 1 / (3 * math.Sqrt(d))
		for {
			u := -exprnd()
			x := normrnd()
			v := 1 + x*c
			v = v * v * v
			if u < 0.5*x*x+d*(1-v+math.Log(v)) {
				return d * v / b
			}
		}
	}
	panic("unreachable")
}

// Survival returns the survival function (complementary CDF) at x.
func (g Gamma) Survival(x float64) float64 {
	if x < 0 {
		return 1
	}
	return mathext.GammaIncComp(g.Alpha, g.Beta*x)
}

// StdDev returns the standard deviation of the probability distribution.
func (g Gamma) StdDev() float64 {
	return math.Sqrt(g.Variance())
}

// Variance returns the variance of the probability distribution.
func (g Gamma) Variance() float64 {
	return g.Alpha / g.Beta / g.Beta
}
