// Copyright ©2016 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mathext

import (
	"math"
	"testing"
)

func TestDigamma(t *testing.T) {
	for i, test := range []struct {
		x, want float64
	}{
		// Results computed using WolframAlpha.
		{-100.5, 4.615124601338064117341315601525112558522917517910505881343},
		{.5, -1.96351002602142347944097633299875556719315960466043},
		{10, 2.251752589066721107647456163885851537211808918028330369448},
		{math.Pow10(20), 46.05170185988091368035482909368728415202202143924212618733},
	} {

		if got := Digamma(test.x); math.Abs(got-test.want) > 1e-10 {
			t.Errorf("test %d Digamma(%g) failed: got %g want %g", i, test.x, got, test.want)
		}
	}
}
