// Copyright ©2016 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stat

import (
	"testing"

	"github.com/gonum/floats"
)

// Test cases where calculated manually.
func TestROC(t *testing.T) {
	cases := []struct {
		y       []float64
		c       []bool
		w       []float64
		n       int
		wantTPR []float64
		wantFPR []float64
	}{
		{
			y:       []float64{0, 3, 5, 6, 7.5, 8},
			c:       []bool{true, false, true, false, false, false},
			wantTPR: []float64{0, 0.5, 0.5, 1, 1, 1, 1},
			wantFPR: []float64{0, 0, 0.25, 0.25, 0.5, 0.75, 1},
		},
		{
			y:       []float64{0, 3, 5, 6, 7.5, 8},
			c:       []bool{true, false, true, false, false, false},
			w:       []float64{4, 1, 6, 3, 2, 2},
			wantTPR: []float64{0, 0.4, 0.4, 1, 1, 1, 1},
			wantFPR: []float64{0, 0, 0.125, 0.125, 0.5, 0.75, 1},
		},
		{
			y:       []float64{0, 3, 5, 6, 7.5, 8},
			c:       []bool{true, false, true, false, false, false},
			n:       int(5),
			wantTPR: []float64{0, 0.5, 0.5, 1, 1},
			wantFPR: []float64{0, 0, 0.25, 0.5, 1},
		},
		{
			y:       []float64{0, 3, 5, 6, 7.5, 8},
			c:       []bool{true, false, true, false, false, false},
			n:       int(9),
			wantTPR: []float64{0, 0.5, 0.5, 0.5, 0.5, 1, 1, 1, 1},
			wantFPR: []float64{0, 0, 0, 0.25, 0.25, 0.25, 0.5, 0.5, 1},
		},
		{
			y:       []float64{0, 3, 5, 6, 7.5, 8},
			c:       []bool{true, false, true, false, false, false},
			w:       []float64{4, 1, 6, 3, 2, 2},
			n:       int(5),
			wantTPR: []float64{0, 0.4, 0.4, 1, 1},
			wantFPR: []float64{0, 0, 0.125, 0.5, 1},
		},
		{
			y:       []float64{0, 3, 5, 6, 7.5, 8},
			c:       []bool{true, false, true, false, false, false},
			w:       []float64{4, 1, 6, 3, 2, 2},
			n:       int(9),
			wantTPR: []float64{0, 0.4, 0.4, 0.4, 0.4, 1, 1, 1, 1},
			wantFPR: []float64{0, 0, 0, 0.125, 0.125, 0.125, 0.5, 0.5, 1},
		},
		{
			y:       []float64{0, 3, 6, 6, 6, 8},
			c:       []bool{true, false, true, false, false, false},
			wantTPR: []float64{0, 0.5, 0.5, 1, 1},
			wantFPR: []float64{0, 0, 0.25, 0.75, 1},
		},
		{
			y:       []float64{0, 3, 6, 6, 6, 8},
			c:       []bool{true, false, true, false, false, false},
			w:       []float64{4, 1, 6, 3, 2, 2},
			wantTPR: []float64{0, 0.4, 0.4, 1, 1},
			wantFPR: []float64{0, 0, 0.125, 0.75, 1},
		},
		{
			y:       []float64{0, 3, 6, 6, 6, 8},
			c:       []bool{true, false, true, false, false, false},
			n:       int(5),
			wantTPR: []float64{0, 0.5, 0.5, 1, 1},
			wantFPR: []float64{0, 0, 0.25, 0.75, 1},
		},
		{
			y:       []float64{0, 3, 6, 6, 6, 8},
			c:       []bool{true, false, true, false, false, false},
			n:       int(9),
			wantTPR: []float64{0, 0.5, 0.5, 0.5, 0.5, 0.5, 1, 1, 1},
			wantFPR: []float64{0, 0, 0, 0.25, 0.25, 0.25, 0.75, 0.75, 1},
		},
		{
			y:       []float64{0, 3, 6, 6, 6, 8},
			c:       []bool{true, false, true, false, false, false},
			w:       []float64{4, 1, 6, 3, 2, 2},
			n:       int(5),
			wantTPR: []float64{0, 0.4, 0.4, 1, 1},
			wantFPR: []float64{0, 0, 0.125, 0.75, 1},
		},
		{
			y:       []float64{0, 3, 6, 6, 6, 8},
			c:       []bool{true, false, true, false, false, false},
			w:       []float64{4, 1, 6, 3, 2, 2},
			n:       int(9),
			wantTPR: []float64{0, 0.4, 0.4, 0.4, 0.4, 0.4, 1, 1, 1},
			wantFPR: []float64{0, 0, 0, 0.125, 0.125, 0.125, 0.75, 0.75, 1},
		},
		{
			y:       []float64{1, 2},
			c:       []bool{true, true},
			wantTPR: []float64{0, 0.5, 1},
			wantFPR: []float64{0, 0, 1},
		},
		{
			y:       []float64{1, 2},
			c:       []bool{true, true},
			n:       int(2),
			wantTPR: []float64{0, 1},
			wantFPR: []float64{0, 1},
		},
		{
			y:       []float64{1, 2},
			c:       []bool{true, true},
			n:       int(7),
			wantTPR: []float64{0, 0.5, 0.5, 0.5, 0.5, 0.5, 1},
			wantFPR: []float64{0, 0, 0, 0, 0, 0, 1},
		},
		{
			y:       []float64{1},
			c:       []bool{true},
			wantTPR: []float64{0, 1},
			wantFPR: []float64{0, 1},
		},
		{
			y:       []float64{1},
			c:       []bool{true},
			n:       int(2),
			wantTPR: []float64{0, 1},
			wantFPR: []float64{0, 1},
		},
		{
			y:       []float64{1},
			c:       []bool{false},
			wantTPR: []float64{0, 1},
			wantFPR: []float64{0, 1},
		},
		{
			y:       []float64{0.01, 0.02, 0.03, 0.04, 0.05, 0.06, 10},
			c:       []bool{true, false, true, true, false, false, true},
			n:       int(5),
			wantTPR: []float64{0, 0.75, 0.75, 0.75, 1},
			wantFPR: []float64{0, 1, 1, 1, 1},
		},
		{
			y:       []float64{},
			c:       []bool{},
			wantTPR: nil,
			wantFPR: nil,
		},
		{
			y:       []float64{},
			c:       []bool{},
			n:       int(5),
			wantTPR: nil,
			wantFPR: nil,
		},
	}
	for i, test := range cases {
		gotTPR, gotFPR := ROC(test.n, test.y, test.c, test.w)
		if !floats.Same(gotTPR, test.wantTPR) {
			t.Errorf("%d: unexpected TPR got:%v want:%v", i, gotTPR, test.wantTPR)
		}
		if !floats.Same(gotFPR, test.wantFPR) {
			t.Errorf("%d: unexpected FPR got:%v want:%v", i, gotFPR, test.wantFPR)
		}
	}
}
