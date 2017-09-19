// Copyright ©2016 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter

import (
	"image/color"
	"log"
	"math"
	"testing"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/internal/cmpimg"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

// Example_rotation gives some examples of rotating text.
func Example_rotation() {
	n := 100
	xmax := 2 * math.Pi

	// Sin creates a sine curve.
	sin := func(n int, xmax float64) XYs {
		xy := make(XYs, n)
		for i := 0; i < n; i++ {
			xy[i].X = xmax / float64(n) * float64(i)
			xy[i].Y = math.Sin(xy[i].X) * 100
		}
		return xy
	}

	// These points will make up our sine curve.
	linePoints := sin(n, xmax)

	// These points are our label locations.
	labelPoints := sin(8, xmax)

	p, err := plot.New()
	if err != nil {
		log.Panic(err)
	}
	p.Title.Text = "Rotation Example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "100 × Sine X"

	l, err := NewLine(linePoints)
	if err != nil {
		log.Panic(err)
	}
	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Color = color.RGBA{B: 255, A: 255}

	labelData := XYLabels{
		XYs:    labelPoints,
		Labels: []string{"0", "pi/4", "pi/2", "3pi/4", "pi", "5pi/4", "3pi/2", "7pi/4", "2pi"},
	}

	labels, err := NewLabels(labelData)
	if err != nil {
		log.Panic(err)
	}

	for i := range labels.TextStyle {
		x := labels.XYs[i].X

		// Set the label rotation to the slope of the line, so the label is
		// parallel with the line.
		labels.TextStyle[i].Rotation = math.Atan(math.Cos(x))
		labels.TextStyle[i].XAlign = draw.XCenter
		labels.TextStyle[i].YAlign = draw.YCenter
		// Move the labels away from the line so they're more easily readable.
		if x >= math.Pi {
			labels.TextStyle[i].YAlign = draw.YTop
		} else {
			labels.TextStyle[i].YAlign = draw.YBottom
		}
	}

	p.Add(l, labels)

	// Add boundary boxes for debugging.
	p.Add(NewGlyphBoxes())

	p.NominalX("0", "The number 1", "Number 2", "The number 3", "Number 4",
		"The number 5", "Number 6")

	// Change the rotation of the X tick labels to make them fit better.
	p.X.Tick.Label.Rotation = math.Pi / 5
	p.X.Tick.Label.YAlign = draw.YCenter
	p.X.Tick.Label.XAlign = draw.XRight

	// Also change the rotation of the Y tick labels.
	p.Y.Tick.Label.Rotation = math.Pi / 2
	p.Y.Tick.Label.XAlign = draw.XCenter
	p.Y.Tick.Label.YAlign = draw.YBottom

	err = p.Save(200, 150, "testdata/rotation.png")
	if err != nil {
		log.Panic(err)
	}
}

func TestRotation(t *testing.T) {
	cmpimg.CheckPlot(Example_rotation, t, "rotation.png")
}
