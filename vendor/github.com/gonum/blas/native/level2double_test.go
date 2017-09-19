// Copyright ©2014 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package native

import (
	"testing"

	"github.com/gonum/blas/testblas"
)

func TestDgemv(t *testing.T) {
	testblas.DgemvTest(t, impl)
}

func TestDger(t *testing.T) {
	testblas.DgerTest(t, impl)
}

func TestDtxmv(t *testing.T) {
	testblas.DtxmvTest(t, impl)
}

func TestDgbmv(t *testing.T) {
	testblas.DgbmvTest(t, impl)
}

func TestDtbsv(t *testing.T) {
	testblas.DtbsvTest(t, impl)
}

func TestDsbmv(t *testing.T) {
	testblas.DsbmvTest(t, impl)
}

func TestDtbmv(t *testing.T) {
	testblas.DtbmvTest(t, impl)
}

func TestDtrsv(t *testing.T) {
	testblas.DtrsvTest(t, impl)
}

func TestDtrmv(t *testing.T) {
	testblas.DtrmvTest(t, impl)
}

func TestDsymv(t *testing.T) {
	testblas.DsymvTest(t, impl)
}

func TestDsyr(t *testing.T) {
	testblas.DsyrTest(t, impl)
}

func TestDsyr2(t *testing.T) {
	testblas.Dsyr2Test(t, impl)
}

func TestDspr2(t *testing.T) {
	testblas.Dspr2Test(t, impl)
}

func TestDspr(t *testing.T) {
	testblas.DsprTest(t, impl)
}

func TestDspmv(t *testing.T) {
	testblas.DspmvTest(t, impl)
}

func TestDtpsv(t *testing.T) {
	testblas.DtpsvTest(t, impl)
}

func TestDtpmv(t *testing.T) {
	testblas.DtpmvTest(t, impl)
}
