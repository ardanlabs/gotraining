// Copyright ©2017 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plot

import (
	"math"
	"os"
	"testing"

	"gonum.org/v1/plot/internal/cmpimg"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
)

func ExampleAlign() {
	const rows, cols = 4, 3
	plots := make([][]*Plot, rows)
	for j := 0; j < rows; j++ {
		plots[j] = make([]*Plot, cols)
		for i := 0; i < cols; i++ {
			if i == 0 && j == 2 {
				// This shows what happens when there are nil plots.
				continue
			}

			p, err := New()
			if err != nil {
				panic(err)
			}

			if j == 0 && i == 2 {
				// This shows what happens when the axis padding
				// is different among plots.
				p.X.Padding, p.Y.Padding = 0, 0
			}

			if j == 1 && i == 1 {
				// To test the Align function, we make the axis labels
				// on one of the plots stick out.
				p.Y.Max = 1e9
				p.X.Max = 1e9
				p.X.Tick.Label.Rotation = math.Pi / 2
				p.X.Tick.Label.XAlign = draw.XRight
				p.X.Tick.Label.YAlign = draw.YCenter
				p.X.Tick.Label.Font.Size = 8
				p.Y.Tick.Label.Font.Size = 8
			} else {
				p.Y.Max = 1e9
				p.X.Max = 1e9
				p.X.Tick.Label.Font.Size = 1
				p.Y.Tick.Label.Font.Size = 1
			}

			plots[j][i] = p
		}
	}

	img := vgimg.New(vg.Points(150), vg.Points(175))
	dc := draw.New(img)

	t := draw.Tiles{
		Rows: rows,
		Cols: cols,
	}

	canvases := Align(plots, t, dc)
	for j := 0; j < rows; j++ {
		for i := 0; i < cols; i++ {
			if plots[j][i] != nil {
				plots[j][i].Draw(canvases[j][i])
			}
		}
	}

	w, err := os.Create("testdata/align.png")
	if err != nil {
		panic(err)
	}

	png := vgimg.PngCanvas{Canvas: img}
	if _, err := png.WriteTo(w); err != nil {
		panic(err)
	}
}

func TestAlign(t *testing.T) {
	cmpimg.CheckPlot(ExampleAlign, t, "align.png")
}
