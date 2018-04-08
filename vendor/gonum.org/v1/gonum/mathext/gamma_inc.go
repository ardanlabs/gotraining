// Copyright ©2016 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mathext

import (
	"gonum.org/v1/gonum/mathext/internal/cephes"
)

// GammaInc computes the incomplete Gamma integral.
//  GammaInc(a,x) = (1/ Γ(a)) \int_0^x e^{-t} t^{a-1} dt
// The input argument a must be positive and x must be non-negative or GammaInc
// will panic.
//
// See http://mathworld.wolfram.com/IncompleteGammaFunction.html
// or https://en.wikipedia.org/wiki/Incomplete_gamma_function for more detailed
// information.
func GammaInc(a, x float64) float64 {
	return cephes.Igam(a, x)
}

// GammaIncComp computes the complemented incomplete Gamma integral.
//  GammaIncComp(a,x) = 1 - GammaInc(a,x)
//                    = (1/ Γ(a)) \int_0^\infty e^{-t} t^{a-1} dt
// The input argument a must be positive and x must be non-negative or
// GammaIncComp will panic.
func GammaIncComp(a, x float64) float64 {
	return cephes.IgamC(a, x)
}

// GammaIncInv computes the inverse of the incomplete Gamma integral. That is,
// it returns the x such that:
//  GammaInc(a, x) = y
// The input argument a must be positive and y must be between 0 and 1
// inclusive or GammaIncInv will panic. GammaIncInv should return a positive
// number, but can return NaN if there is a failure to converge.
func GammaIncInv(a, y float64) float64 {
	return gammaIncInv(a, y)
}

// GammaIncCompInv computes the inverse of the complemented incomplete Gamma
// integral. That is, it returns the x such that:
//  GammaIncComp(a, x) = y
// The input argument a must be positive and y must be between 0 and 1
// inclusive or GammaIncCompInv will panic. GammaIncCompInv should return a
// positive number, but can return 0 even with non-zero y due to underflow.
func GammaIncCompInv(a, y float64) float64 {
	return cephes.IgamI(a, y)
}
