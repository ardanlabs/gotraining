// Copyright ©2014 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package native

import (
	"math/rand"
	"testing"

	"github.com/gonum/blas"
)

func TestDgemmParallel(t *testing.T) {
	for i, test := range []struct {
		m     int
		n     int
		k     int
		alpha float64
		tA    blas.Transpose
		tB    blas.Transpose
	}{
		{
			m:     3,
			n:     4,
			k:     2,
			alpha: 2.5,
			tA:    blas.NoTrans,
			tB:    blas.NoTrans,
		},
		{
			m:     blockSize*2 + 5,
			n:     3,
			k:     2,
			alpha: 2.5,
			tA:    blas.NoTrans,
			tB:    blas.NoTrans,
		},
		{
			m:     3,
			n:     blockSize * 2,
			k:     2,
			alpha: 2.5,
			tA:    blas.NoTrans,
			tB:    blas.NoTrans,
		},
		{
			m:     2,
			n:     3,
			k:     blockSize*3 - 2,
			alpha: 2.5,
			tA:    blas.NoTrans,
			tB:    blas.NoTrans,
		},
		{
			m:     blockSize * minParBlock,
			n:     3,
			k:     2,
			alpha: 2.5,
			tA:    blas.NoTrans,
			tB:    blas.NoTrans,
		},
		{
			m:     3,
			n:     blockSize * minParBlock,
			k:     2,
			alpha: 2.5,
			tA:    blas.NoTrans,
			tB:    blas.NoTrans,
		},
		{
			m:     2,
			n:     3,
			k:     blockSize * minParBlock,
			alpha: 2.5,
			tA:    blas.NoTrans,
			tB:    blas.NoTrans,
		},
		{
			m:     blockSize*minParBlock + 1,
			n:     blockSize * minParBlock,
			k:     3,
			alpha: 2.5,
			tA:    blas.NoTrans,
			tB:    blas.NoTrans,
		},
		{
			m:     3,
			n:     blockSize*minParBlock + 2,
			k:     blockSize * 3,
			alpha: 2.5,
			tA:    blas.NoTrans,
			tB:    blas.NoTrans,
		},
		{
			m:     blockSize * minParBlock,
			n:     3,
			k:     blockSize * minParBlock,
			alpha: 2.5,
			tA:    blas.NoTrans,
			tB:    blas.NoTrans,
		},
		{
			m:     blockSize * minParBlock,
			n:     blockSize * minParBlock,
			k:     blockSize * 3,
			alpha: 2.5,
			tA:    blas.NoTrans,
			tB:    blas.NoTrans,
		},
		{
			m:     blockSize + blockSize/2,
			n:     blockSize + blockSize/2,
			k:     blockSize + blockSize/2,
			alpha: 2.5,
			tA:    blas.NoTrans,
			tB:    blas.NoTrans,
		},
	} {
		testMatchParallelSerial(t, i, blas.NoTrans, blas.NoTrans, test.m, test.n, test.k, test.alpha)
		testMatchParallelSerial(t, i, blas.Trans, blas.NoTrans, test.m, test.n, test.k, test.alpha)
		testMatchParallelSerial(t, i, blas.NoTrans, blas.Trans, test.m, test.n, test.k, test.alpha)
		testMatchParallelSerial(t, i, blas.Trans, blas.Trans, test.m, test.n, test.k, test.alpha)
	}
}

func testMatchParallelSerial(t *testing.T, i int, tA, tB blas.Transpose, m, n, k int, alpha float64) {
	var (
		rowA, colA int
		rowB, colB int
	)
	if tA == blas.NoTrans {
		rowA = m
		colA = k
	} else {
		rowA = k
		colA = m
	}
	if tB == blas.NoTrans {
		rowB = k
		colB = n
	} else {
		rowB = n
		colB = k
	}
	a := randmat(rowA, colA, colA)
	b := randmat(rowB, colB, colB)
	c := randmat(m, n, n)

	aClone := a.clone()
	bClone := b.clone()
	cClone := c.clone()

	lda := colA
	ldb := colB
	ldc := n
	dgemmSerial(tA == blas.Trans, tB == blas.Trans, m, n, k, a.data, lda, b.data, ldb, cClone.data, ldc, alpha)
	dgemmParallel(tA == blas.Trans, tB == blas.Trans, m, n, k, a.data, lda, b.data, ldb, c.data, ldc, alpha)
	if !a.equal(aClone) {
		t.Errorf("Case %v: a changed during call to dgemmParallel", i)
	}
	if !b.equal(bClone) {
		t.Errorf("Case %v: b changed during call to dgemmParallel", i)
	}
	if !c.equalWithinAbs(cClone, 1e-12) {
		t.Errorf("Case %v: answer not equal parallel and serial", i)
	}
}

func randmat(r, c, stride int) general64 {
	data := make([]float64, r*stride+c)
	for i := range data {
		data[i] = rand.Float64()
	}
	return general64{
		data:   data,
		rows:   r,
		cols:   c,
		stride: stride,
	}
}
