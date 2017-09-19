// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter

import (
	"fmt"
	"log"
	"os"
	"testing"

	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/internal/cmpimg"
	"gonum.org/v1/plot/palette"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
)

type offsetUnitGrid struct {
	XOffset, YOffset float64

	Data mat.Matrix
}

func (g offsetUnitGrid) Dims() (c, r int)   { r, c = g.Data.Dims(); return c, r }
func (g offsetUnitGrid) Z(c, r int) float64 { return g.Data.At(r, c) }
func (g offsetUnitGrid) X(c int) float64 {
	_, n := g.Data.Dims()
	if c < 0 || c >= n {
		panic("index out of range")
	}
	return float64(c) + g.XOffset
}
func (g offsetUnitGrid) Y(r int) float64 {
	m, _ := g.Data.Dims()
	if r < 0 || r >= m {
		panic("index out of range")
	}
	return float64(r) + g.YOffset
}

func ExampleHeatMap() {
	m := offsetUnitGrid{
		XOffset: -2,
		YOffset: -1,
		Data: mat.NewDense(3, 4, []float64{
			1, 2, 3, 4,
			5, 6, 7, 8,
			9, 10, 11, 12,
		})}
	pal := palette.Heat(12, 1)
	h := NewHeatMap(m, pal)

	p, err := plot.New()
	if err != nil {
		log.Panic(err)
	}
	p.Title.Text = "Heat map"

	p.Add(h)

	// Create a legend.
	thumbs := PaletteThumbnailers(pal)
	for i := len(thumbs) - 1; i >= 0; i-- {
		t := thumbs[i]
		if i != 0 && i != len(thumbs)-1 {
			p.Legend.Add("", t)
			continue
		}
		var val float64
		switch i {
		case 0:
			val = h.Min
		case len(thumbs) - 1:
			val = h.Max
		}
		p.Legend.Add(fmt.Sprintf("%.2g", val), t)
	}
	// This is the width of the legend, experimentally determined.
	const legendWidth = 1.25 * vg.Centimeter
	// Slide the legend over so it doesn't overlap the HeatMap.
	p.Legend.XOffs = legendWidth

	p.X.Padding = 0
	p.Y.Padding = 0
	p.X.Max = 1.5
	p.Y.Max = 1.5

	img := vgimg.New(250, 175)
	dc := draw.New(img)
	dc = draw.Crop(dc, 0, -legendWidth, 0, 0) // Make space for the legend.
	p.Draw(dc)
	w, err := os.Create("testdata/heatMap.png")
	if err != nil {
		log.Panic(err)
	}
	png := vgimg.PngCanvas{Canvas: img}
	if _, err = png.WriteTo(w); err != nil {
		log.Panic(err)
	}
}

func TestHeatMap(t *testing.T) {
	cmpimg.CheckPlot(ExampleHeatMap, t, "heatMap.png")
}
