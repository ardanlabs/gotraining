// Copyright ©2014 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gonum

import (
	"math"
)

type general64 struct {
	data       []float64
	rows, cols int
	stride     int
}

func (g general64) clone() general64 {
	data := make([]float64, len(g.data))
	copy(data, g.data)
	return general64{
		data:   data,
		rows:   g.rows,
		cols:   g.cols,
		stride: g.stride,
	}
}

func (g general64) equal(a general64) bool {
	if g.rows != a.rows || g.cols != a.cols || g.stride != a.stride {
		return false
	}
	for i, v := range g.data {
		if a.data[i] != v {
			return false
		}
	}
	return true
}

func (g general64) equalWithinAbs(a general64, tol float64) bool {
	if g.rows != a.rows || g.cols != a.cols || g.stride != a.stride {
		return false
	}
	for i, v := range g.data {
		if math.Abs(a.data[i]-v) > tol {
			return false
		}
	}
	return true
}
