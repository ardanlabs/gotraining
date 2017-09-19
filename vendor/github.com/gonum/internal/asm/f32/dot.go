// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package f32

// DotUnitary is
//  for i, v := range x {
//  	sum += y[i] * v
//  }
//  return sum
func DotUnitary(x, y []float32) (sum float32) {
	for i, v := range x {
		sum += y[i] * v
	}
	return sum
}

// DotInc is
//  for i := 0; i < int(n); i++ {
//  	sum += y[iy] * x[ix]
//  	ix += incX
//  	iy += incY
//  }
//  return sum
func DotInc(x, y []float32, n, incX, incY, ix, iy uintptr) (sum float32) {
	for i := 0; i < int(n); i++ {
		sum += y[iy] * x[ix]
		ix += incX
		iy += incY
	}
	return sum
}
