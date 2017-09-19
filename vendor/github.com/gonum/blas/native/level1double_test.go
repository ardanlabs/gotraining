// Copyright ©2014 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package native

import (
	"testing"

	"github.com/gonum/blas/testblas"
)

var impl Implementation

func TestDasum(t *testing.T) {
	testblas.DasumTest(t, impl)
}

func TestDaxpy(t *testing.T) {
	testblas.DaxpyTest(t, impl)
}

func TestDdot(t *testing.T) {
	testblas.DdotTest(t, impl)
}

func TestDnrm2(t *testing.T) {
	testblas.Dnrm2Test(t, impl)
}

func TestIdamax(t *testing.T) {
	testblas.IdamaxTest(t, impl)
}

func TestDswap(t *testing.T) {
	testblas.DswapTest(t, impl)
}

func TestDcopy(t *testing.T) {
	testblas.DcopyTest(t, impl)
}

func TestDrotg(t *testing.T) {
	testblas.DrotgTest(t, impl)
}

func TestDrotmg(t *testing.T) {
	testblas.DrotmgTest(t, impl)
}

func TestDrot(t *testing.T) {
	testblas.DrotTest(t, impl)
}

func TestDrotm(t *testing.T) {
	testblas.DrotmTest(t, impl)
}

func TestDscal(t *testing.T) {
	testblas.DscalTest(t, impl)
}
