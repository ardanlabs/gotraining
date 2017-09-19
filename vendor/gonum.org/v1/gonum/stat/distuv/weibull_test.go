// Copyright ©2014 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package distuv

import (
	"math"
	"testing"
)

func TestHalfKStandardWeibullProb(t *testing.T) {
	pts := []univariateProbPoint{
		{
			loc:     0,
			prob:    math.Inf(1),
			cumProb: 0,
			logProb: math.Inf(1),
		},
		{
			loc:     -1,
			prob:    0,
			cumProb: 0,
			logProb: 0,
		},
		{
			loc:     1,
			prob:    0.183939720585721,
			cumProb: 0.632120558828558,
			logProb: -1.693147180559950,
		},
		{
			loc:     20,
			prob:    0.001277118038048,
			cumProb: 0.988577109006533,
			logProb: -6.663149272336520,
		},
	}
	testDistributionProbs(t, Weibull{K: 0.5, Lambda: 1}, "0.5K Standard Weibull", pts)
}

func TestExponentialStandardWeibullProb(t *testing.T) {
	pts := []univariateProbPoint{
		{
			loc:     0,
			prob:    1,
			cumProb: 0,
			logProb: math.Inf(1),
		},
		{
			loc:     -1,
			prob:    0,
			cumProb: 0,
			logProb: 0,
		},
		{
			loc:     1,
			prob:    0.367879441171442,
			cumProb: 0.632120558828558,
			logProb: -1.0,
		},
		{
			loc:     20,
			prob:    0.000000002061154,
			cumProb: 0.999999997938846,
			logProb: -20.0,
		},
	}
	testDistributionProbs(t, Weibull{K: 1, Lambda: 1}, "1K (Exponential) Standard Weibull", pts)
}

func TestRayleighStandardWeibullProb(t *testing.T) {
	pts := []univariateProbPoint{
		{
			loc:     0,
			prob:    0,
			cumProb: 0,
			logProb: math.Inf(-1),
		},
		{
			loc:     -1,
			prob:    0,
			cumProb: 0,
			logProb: 0,
		},
		{
			loc:     1,
			prob:    0.735758882342885,
			cumProb: 0.632120558828558,
			logProb: -0.306852819440055,
		},
		{
			loc:     20,
			prob:    0,
			cumProb: 1,
			logProb: -396.31112054588607,
		},
	}
	testDistributionProbs(t, Weibull{K: 2, Lambda: 1}, "2K (Rayleigh) Standard Weibull", pts)
}

func TestFiveKStandardWeibullProb(t *testing.T) {
	pts := []univariateProbPoint{
		{
			loc:     0,
			prob:    0,
			cumProb: 0,
			logProb: math.Inf(-1),
		},
		{
			loc:     -1,
			prob:    0,
			cumProb: 0,
			logProb: 0,
		},
		{
			loc:     1,
			prob:    1.839397205857210,
			cumProb: 0.632120558828558,
			logProb: 0.609437912434100,
		},
		{
			loc:     20,
			prob:    0,
			cumProb: 1,
			logProb: -3199986.4076329935,
		},
	}
	testDistributionProbs(t, Weibull{K: 5, Lambda: 1}, "5K Standard Weibull", pts)
}

func TestScaledUpHalfKStandardWeibullProb(t *testing.T) {
	pts := []univariateProbPoint{
		{
			loc:     0,
			prob:    math.Inf(1),
			cumProb: 0,
			logProb: math.Inf(1),
		},
		{
			loc:     -1,
			prob:    0,
			cumProb: 0,
			logProb: 0,
		},
		{
			loc:     1,
			prob:    0.180436508682207,
			cumProb: 0.558022622759326,
			logProb: -1.712376315541750,
		},
		{
			loc:     20,
			prob:    0.002369136850928,
			cumProb: 0.974047406098605,
			logProb: -6.045229588092130,
		},
	}
	testDistributionProbs(t, Weibull{K: 0.5, Lambda: 1.5}, "0.5K 1.5λ Weibull", pts)
}

func TestScaledDownHalfKStandardWeibullProb(t *testing.T) {
	pts := []univariateProbPoint{
		{
			loc:     0,
			prob:    math.Inf(1),
			cumProb: 0,
			logProb: math.Inf(1),
		},
		{
			loc:     -1,
			prob:    0,
			cumProb: 0,
			logProb: 0,
		},
		{
			loc:     1,
			prob:    0.171909491538362,
			cumProb: 0.756883265565786,
			logProb: -1.760787152653070,
		},
		{
			loc:     20,
			prob:    0.000283302579100,
			cumProb: 0.998208237166091,
			logProb: -8.168995047393730,
		},
	}
	testDistributionProbs(t, Weibull{K: 0.5, Lambda: 0.5}, "0.5K 0.5λ Weibull", pts)
}

func TestWeibullScore(t *testing.T) {
	for _, test := range []*Weibull{
		{
			K:      1,
			Lambda: 1,
		},
		{
			K:      2,
			Lambda: 3.6,
		},
		{
			K:      3.4,
			Lambda: 8,
		},
	} {
		testDerivParam(t, test)
	}
}
