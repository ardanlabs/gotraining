// Copyright ©2016 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mathext

import (
	"math"
)

// Digamma returns the logorithmic derivative of the gamma function at x.
//  ψ(x) = d/dx (Ln (Γ(x)).
// Note that if x is a negative integer in [-7, 0] this function will return
// negative Inf.
func Digamma(x float64) float64 {
	// This is adapted from
	// http://web.science.mq.edu.au/~mjohnson/code/digamma.c
	var result float64
	for ; x < 7.0; x++ {
		result -= 1 / x
	}
	x -= 1.0 / 2.0
	xx := 1.0 / x
	xx2 := xx * xx
	xx4 := xx2 * xx2
	result += math.Log(x) + (1./24.)*xx2 - (7.0/960.0)*xx4 + (31.0/8064.0)*xx4*xx2 - (127.0/30720.0)*xx4*xx4
	return result
}
