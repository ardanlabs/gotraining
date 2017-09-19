// Copyright ©2017 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter

import (
	"log"
	"testing"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/internal/cmpimg"
	"gonum.org/v1/plot/palette/moreland"
)

func ExampleColorBar_horizontal() {
	p, err := plot.New()
	if err != nil {
		log.Panic(err)
	}
	l := &ColorBar{ColorMap: moreland.ExtendedBlackBody()}
	l.ColorMap.SetMin(0.5)
	l.ColorMap.SetMax(1.5)
	p.Add(l)
	p.HideY()
	p.X.Padding = 0
	p.Title.Text = "Title"

	if err = p.Save(300, 48, "testdata/colorBarHorizontal.png"); err != nil {
		log.Panic(err)
	}
}

func TestColorBar_horizontal(t *testing.T) {
	cmpimg.CheckPlot(ExampleColorBar_horizontal, t, "colorBarHorizontal.png")
}

func ExampleColorBar_vertical() {
	p, err := plot.New()
	if err != nil {
		log.Panic(err)
	}
	l := &ColorBar{ColorMap: moreland.ExtendedBlackBody()}
	l.ColorMap.SetMin(0.5)
	l.ColorMap.SetMax(1.5)
	l.Vertical = true
	p.Add(l)
	p.HideX()
	p.Y.Padding = 0
	p.Title.Text = "Title"

	if err = p.Save(40, 300, "testdata/colorBarVertical.png"); err != nil {
		log.Panic(err)
	}
}

func TestColorBar_vertical(t *testing.T) {
	cmpimg.CheckPlot(ExampleColorBar_vertical, t, "colorBarVertical.png")
}
