// Copyright ©2017 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package draw_test

import (
	"fmt"

	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

// SplitHorizontal returns the left and right portions of c after splitting it
// along a vertical line distance x from the left of c.
func SplitHorizontal(c draw.Canvas, x vg.Length) (left, right draw.Canvas) {
	return draw.Crop(c, 0, c.Min.X-c.Max.X+x, 0, 0), draw.Crop(c, x, 0, 0, 0)
}

func ExampleCrop_splitHorizontal() {
	var c draw.Canvas
	// Split c along a vertical line centered on the canvas.
	left, right := SplitHorizontal(c, c.Size().X/2)
	fmt.Println(left.Rectangle.Size(), right.Rectangle.Size())
}
