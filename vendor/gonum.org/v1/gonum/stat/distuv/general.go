// Copyright ©2014 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package distuv

import "math"

// Parameter represents a parameter of a probability distribution
type Parameter struct {
	Name  string
	Value float64
}

var (
	badPercentile = "distuv: percentile out of bounds"
	badLength     = "distuv: slice length mismatch"
	badSuffStat   = "distuv: wrong suffStat length"
	badNoSamples  = "distuv: must have at least one sample"
)

var (
	expNegOneHalf = math.Exp(-0.5)
)
