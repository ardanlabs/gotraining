// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package distuv

type LogProber interface {
	LogProb(float64) float64
}

type Rander interface {
	Rand() float64
}

type RandLogProber interface {
	Rander
	LogProber
}

type Quantiler interface {
	Quantile(p float64) float64
}
