// Copyright ©2016 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mathext

import (
	"math"
	"testing"
)

func TestMvLgamma(t *testing.T) {
	// Values compared with scipy
	for i, test := range []struct {
		v   float64
		dim int
		ans float64
	}{
		{10, 5, 58.893841851237397},
		{3, 1, 0.69314718055994529},
	} {
		ans := MvLgamma(test.v, test.dim)
		if math.Abs(test.ans-ans) > 1e-14 {
			t.Errorf("Case %v. got=%v want=%v.", i, ans, test.ans)
		}
	}
}
