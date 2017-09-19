// Copyright ©2017 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mat64

import (
	"math/rand"
	"testing"
)

func TestHOGSVD(t *testing.T) {
	const tol = 1e-10
	rnd := rand.New(rand.NewSource(1))
	for cas, test := range []struct {
		r, c int
	}{
		{5, 3},
		{5, 5},
		{150, 150},
		{200, 150},

		// Calculating A_i*A_j^T and A_j*A_i^T fails for wide matrices.
		{3, 5},
	} {
		r := test.r
		c := test.c
		for n := 3; n < 6; n++ {
			data := make([]Matrix, n)
			dataCopy := make([]*Dense, n)
			for trial := 0; trial < 10; trial++ {
				for i := range data {
					d := NewDense(r, c, nil)
					for j := range d.mat.Data {
						d.mat.Data[j] = rnd.Float64()
					}
					data[i] = d
					dataCopy[i] = DenseCopyOf(d)
				}

				var gsvd HOGSVD
				ok := gsvd.Factorize(data...)
				if r >= c {
					if !ok {
						t.Errorf("HOGSVD factorization failed for %d %d×%d matrices: %v", n, r, c, gsvd.Err())
						continue
					}
				} else {
					if ok {
						t.Errorf("HOGSVD factorization unexpectedly succeeded for for %d %d×%d matrices", n, r, c)
					}
					continue
				}
				for i := range data {
					if !Equal(data[i], dataCopy[i]) {
						t.Errorf("A changed during call to HOGSVD.Factorize")
					}
				}
				u, s, v := extractHOGSVD(&gsvd)
				for i, want := range data {
					var got Dense
					sigma := NewDense(c, c, nil)
					for j := 0; j < c; j++ {
						sigma.Set(j, j, s[i][j])
					}

					got.Product(u[i], sigma, v.T())
					if !EqualApprox(&got, want, tol) {
						t.Errorf("test %d n=%d trial %d: unexpected answer\nU_%[4]d * S_%[4]d * V^T:\n% 0.2f\nD_%d:\n% 0.2f",
							cas, n, trial, i, Formatted(&got, Excerpt(5)), i, Formatted(want, Excerpt(5)))
					}
				}
			}
		}
	}
}

func extractHOGSVD(gsvd *HOGSVD) (u []*Dense, s [][]float64, v *Dense) {
	u = make([]*Dense, gsvd.Len())
	s = make([][]float64, gsvd.Len())
	for i := 0; i < gsvd.Len(); i++ {
		u[i] = &Dense{}
		u[i].UFromHOGSVD(gsvd, i)
		s[i] = gsvd.Values(nil, i)
	}
	v = &Dense{}
	v.VFromHOGSVD(gsvd)
	return u, s, v
}
