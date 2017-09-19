// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math32

import (
	"math"
	"testing"
	"testing/quick"

	"github.com/gonum/floats"
)

const tol = 1e-7

func TestAbs(t *testing.T) {
	f := func(x float32) bool {
		y := Abs(x)
		return y == float32(math.Abs(float64(x)))
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestCopySign(t *testing.T) {
	f := func(x struct{ X, Y float32 }) bool {
		y := Copysign(x.X, x.Y)
		return y == float32(math.Copysign(float64(x.X), float64(x.Y)))
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestHypot(t *testing.T) {
	f := func(x struct{ X, Y float32 }) bool {
		y := Hypot(x.X, x.Y)
		if math.Hypot(float64(x.X), float64(x.Y)) > math.MaxFloat32 {
			return true
		}
		return floats.EqualWithinRel(float64(y), math.Hypot(float64(x.X), float64(x.Y)), tol)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInf(t *testing.T) {
	if float64(Inf(1)) != math.Inf(1) || float64(Inf(-1)) != math.Inf(-1) {
		t.Error("float32(inf) not infinite")
	}
}

func TestIsInf(t *testing.T) {
	posInf := float32(math.Inf(1))
	negInf := float32(math.Inf(-1))
	if !IsInf(posInf, 0) || !IsInf(negInf, 0) || !IsInf(posInf, 1) || !IsInf(negInf, -1) || IsInf(posInf, -1) || IsInf(negInf, 1) {
		t.Error("unexpected isInf value")
	}
	f := func(x struct {
		F    float32
		Sign int
	}) bool {
		y := IsInf(x.F, x.Sign)
		return y == math.IsInf(float64(x.F), x.Sign)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestIsNaN(t *testing.T) {
	f := func(x float32) bool {
		y := IsNaN(x)
		return y == math.IsNaN(float64(x))
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestNaN(t *testing.T) {
	if !math.IsNaN(float64(NaN())) {
		t.Errorf("float32(nan) is a number: %f", NaN())
	}
}

func TestSqrt(t *testing.T) {
	f := func(x float32) bool {
		y := Sqrt(x)
		if IsNaN(y) && IsNaN(sqrt(x)) {
			return true
		}
		return floats.EqualWithinRel(float64(y), float64(sqrt(x)), tol)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The original C code and the long comment below are
// from FreeBSD's /usr/src/lib/msun/src/e_sqrt.c and
// came with this notice.  The go code is a simplified
// version of the original C.
//
// ====================================================
// Copyright (C) 1993 by Sun Microsystems, Inc. All rights reserved.
//
// Developed at SunPro, a Sun Microsystems, Inc. business.
// Permission to use, copy, modify, and distribute this
// software is freely granted, provided that this notice
// is preserved.
// ====================================================
//
// __ieee754_sqrt(x)
// Return correctly rounded sqrt.
//           -----------------------------------------
//           | Use the hardware sqrt if you have one |
//           -----------------------------------------
// Method:
//   Bit by bit method using integer arithmetic. (Slow, but portable)
//   1. Normalization
//      Scale x to y in [1,4) with even powers of 2:
//      find an integer k such that  1 <= (y=x*2**(2k)) < 4, then
//              sqrt(x) = 2**k * sqrt(y)
//   2. Bit by bit computation
//      Let q  = sqrt(y) truncated to i bit after binary point (q = 1),
//           i                                                   0
//                                     i+1         2
//          s  = 2*q , and      y  =  2   * ( y - q  ).          (1)
//           i      i            i                 i
//
//      To compute q    from q , one checks whether
//                  i+1       i
//
//                            -(i+1) 2
//                      (q + 2      )  <= y.                     (2)
//                        i
//                                                            -(i+1)
//      If (2) is false, then q   = q ; otherwise q   = q  + 2      .
//                             i+1   i             i+1   i
//
//      With some algebraic manipulation, it is not difficult to see
//      that (2) is equivalent to
//                             -(i+1)
//                      s  +  2       <= y                       (3)
//                       i                i
//
//      The advantage of (3) is that s  and y  can be computed by
//                                    i      i
//      the following recurrence formula:
//          if (3) is false
//
//          s     =  s  ,       y    = y   ;                     (4)
//           i+1      i          i+1    i
//
//      otherwise,
//                         -i                      -(i+1)
//          s     =  s  + 2  ,  y    = y  -  s  - 2              (5)
//           i+1      i          i+1    i     i
//
//      One may easily use induction to prove (4) and (5).
//      Note. Since the left hand side of (3) contain only i+2 bits,
//            it does not necessary to do a full (53-bit) comparison
//            in (3).
//   3. Final rounding
//      After generating the 53 bits result, we compute one more bit.
//      Together with the remainder, we can decide whether the
//      result is exact, bigger than 1/2ulp, or less than 1/2ulp
//      (it will never equal to 1/2ulp).
//      The rounding mode can be detected by checking whether
//      huge + tiny is equal to huge, and whether huge - tiny is
//      equal to huge for some floating point number "huge" and "tiny".
//
func sqrt(x float32) float32 {
	// special cases
	switch {
	case x == 0 || IsNaN(x) || IsInf(x, 1):
		return x
	case x < 0:
		return NaN()
	}
	ix := math.Float32bits(x)
	// normalize x
	exp := int((ix >> shift) & mask)
	if exp == 0 { // subnormal x
		for ix&1<<shift == 0 {
			ix <<= 1
			exp--
		}
		exp++
	}
	exp -= bias // unbias exponent
	ix &^= mask << shift
	ix |= 1 << shift
	if exp&1 == 1 { // odd exp, double x to make it even
		ix <<= 1
	}
	exp >>= 1 // exp = exp/2, exponent of square root
	// generate sqrt(x) bit by bit
	ix <<= 1
	var q, s uint32               // q = sqrt(x)
	r := uint32(1 << (shift + 1)) // r = moving bit from MSB to LSB
	for r != 0 {
		t := s + r
		if t <= ix {
			s = t + r
			ix -= t
			q += r
		}
		ix <<= 1
		r >>= 1
	}
	// final rounding
	if ix != 0 { // remainder, result not exact
		q += q & 1 // round according to extra bit
	}
	ix = q>>1 + uint32(exp-1+bias)<<shift // significand + biased exponent
	return math.Float32frombits(ix)
}
