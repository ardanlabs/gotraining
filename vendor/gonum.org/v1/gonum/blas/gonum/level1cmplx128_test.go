// Copyright ©2017 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gonum

import (
	"testing"

	"gonum.org/v1/gonum/blas/testblas"
)

func TestDzasum(t *testing.T) {
	testblas.DzasumTest(t, impl)
}

func TestDznrm2(t *testing.T) {
	testblas.Dznrm2Test(t, impl)
}

func TestIzamax(t *testing.T) {
	testblas.IzamaxTest(t, impl)
}

func TestZaxpy(t *testing.T) {
	testblas.ZaxpyTest(t, impl)
}

func TestZcopy(t *testing.T) {
	testblas.ZcopyTest(t, impl)
}

func TestZdotc(t *testing.T) {
	testblas.ZdotcTest(t, impl)
}

func TestZdotu(t *testing.T) {
	testblas.ZdotuTest(t, impl)
}

func TestZdscal(t *testing.T) {
	testblas.ZdscalTest(t, impl)
}

func TestZscal(t *testing.T) {
	testblas.ZscalTest(t, impl)
}

func TestZswap(t *testing.T) {
	testblas.ZswapTest(t, impl)
}
