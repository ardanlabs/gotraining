// Copyright ©2017 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package draw_test

import (
	"fmt"

	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

// SplitVertical returns the lower and upper portions of c after
// splitting it along a horizontal line distance y from the
// bottom of c.
func SplitVertical(c draw.Canvas, y vg.Length) (lower, upper draw.Canvas) {
	return draw.Crop(c, 0, 0, 0, c.Min.Y-c.Max.Y+y), draw.Crop(c, 0, 0, y, 0)
}

func ExampleCrop_splitVertical() {
	var c draw.Canvas
	// Split c along a horizontal line centered on the canvas.
	bottom, top := SplitHorizontal(c, c.Size().Y/2)
	fmt.Println(bottom.Rectangle.Size(), top.Rectangle.Size())
}
