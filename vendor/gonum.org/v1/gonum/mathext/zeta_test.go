// Copyright ©2016 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mathext

import (
	"math"
	"testing"
)

func TestZeta(t *testing.T) {
	for i, test := range []struct {
		x, q, want float64
	}{
		// Results computed using scipy.special.zeta
		{1, 1, math.Inf(1)},
		{1.00001, 0.5, 100001.96352290553},
		{1.0001, 25, 9996.8017690244506},
		{1.001, 1, 1000.5772884760117},
		{1.01, 10, 97.773405639173305},
		{1.5, 2, 1.6123753486854886},
		{1.5, 20, 0.45287361712938717},
		{2, -0.7, 14.28618087263834},
		{2.5, 0.5, 6.2471106345688137},
		{5, 2.5, 0.013073166646113805},
		{7.5, 5, 7.9463377443314306e-06},
		{10, -0.5, 2048.0174503557578},
		{10, 0.5, 1024.0174503557578},
		{10, 7.5, 2.5578265694201971e-9},
		{12, 2.5, 1.7089167198843551e-5},
		{17, 0.5, 131072.00101513157},
		{20, -2.5, 2097152.0006014798},
		{20, 0.75, 315.3368689825316},
		{25, 0.25, 1125899906842624.0},
		{30, 1, 1.0000000009313275},
	} {
		if got := Zeta(test.x, test.q); math.Abs(got-test.want) > 1e-10 {
			t.Errorf("test %d Zeta(%g, %g) failed: got %g want %g", i, test.x, test.q, got, test.want)
		}
	}
}
