// Copyright ©2014 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package distuv

import (
	"math"
	"testing"
)

func TestExponentialProb(t *testing.T) {
	pts := []univariateProbPoint{
		{
			loc:     0,
			prob:    1,
			cumProb: 0,
			logProb: 0,
		},
		{
			loc:     -1,
			prob:    0,
			cumProb: 0,
			logProb: math.Inf(-1),
		},
		{
			loc:     1,
			prob:    1 / (math.E),
			cumProb: 0.6321205588285576784044762298385391325541888689682321654921631983025385042551001966428527256540803563,
			logProb: -1,
		},
		{
			loc:     20,
			prob:    math.Exp(-20),
			cumProb: 0.999999997938846377561442172034059619844179023624192724400896307027755338370835976215440646720089072,
			logProb: -20,
		},
	}
	testDistributionProbs(t, Exponential{Rate: 1}, "Exponential", pts)
}

func TestExponentialFitPrior(t *testing.T) {
	testConjugateUpdate(t, func() ConjugateUpdater { return &Exponential{Rate: 13.7} })
}

func TestExponentialScore(t *testing.T) {
	for _, test := range []*Exponential{
		{
			Rate: 1,
		},
		{
			Rate: 0.35,
		},
		{
			Rate: 4.6,
		},
	} {
		testDerivParam(t, test)
	}
}

func TestExponentialFitPanic(t *testing.T) {
	e := Exponential{Rate: 2}
	defer func() {
		r := recover()
		if r != nil {
			t.Errorf("unexpected panic for Fit call: %v", r)
		}
	}()
	e.Fit(make([]float64, 10), nil)
}
