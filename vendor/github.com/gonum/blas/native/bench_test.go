// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package native

import (
	"github.com/gonum/blas"
	"github.com/gonum/blas/testblas"
)

const (
	Sm  = testblas.SmallMat
	Med = testblas.MediumMat
	Lg  = testblas.LargeMat
	Hg  = testblas.HugeMat
)

const (
	T  = blas.Trans
	NT = blas.NoTrans
)
