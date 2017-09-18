// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package f32

// DdotUnitary is
//   for i, v := range x {
//   	sum += float64(y[i]) * float64(v)
//   }
//   return
func DdotUnitary(x, y []float32) (sum float64) {
	for i, v := range x {
		sum += float64(y[i]) * float64(v)
	}
	return
}

// DdotInc is
//  for i := 0; i < int(n); i++ {
//  	sum += float64(y[iy]) * float64(x[ix])
//  	ix += incX
//  	iy += incY
//  }
//  return
func DdotInc(x, y []float32, n, incX, incY, ix, iy uintptr) (sum float64) {
	for i := 0; i < int(n); i++ {
		sum += float64(y[iy]) * float64(x[ix])
		ix += incX
		iy += incY
	}
	return
}
