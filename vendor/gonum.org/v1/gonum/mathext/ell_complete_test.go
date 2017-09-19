// Copyright ©2017 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mathext

import (
	"math"
	"testing"
)

// TestCompleteKE checks if the Legendre's relation for m=0.0001(0.0001)0.9999
// is satisfied with accuracy 1e-14.
func TestCompleteKE(t *testing.T) {
	const tol = 1.0e-14

	for m := 1; m <= 9999; m++ {
		mf := float64(m) / 10000
		mp := 1 - mf
		K, Kp := CompleteK(mf), CompleteK(mp)
		E, Ep := CompleteE(mf), CompleteE(mp)
		legendre := math.Abs(E*Kp + Ep*K - K*Kp - math.Pi/2)
		if legendre > tol {
			t.Fatalf("legendre > tol: m=%v, legendre=%v, tol=%v", mf, legendre, tol)
		}
	}
}
