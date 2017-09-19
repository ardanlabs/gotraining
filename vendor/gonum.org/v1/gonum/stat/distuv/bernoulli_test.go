// Copyright ©2016 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package distuv

import "testing"

func TestBernoulli(t *testing.T) {
	for i, dist := range []Bernoulli{
		{
			P: 0.5,
		},
		{
			P: 0.9,
		},
		{
			P: 0.2,
		},
	} {
		testFullDist(t, dist, i, false)
	}
}
