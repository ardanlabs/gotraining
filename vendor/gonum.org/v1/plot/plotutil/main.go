// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
)

var examples = []struct {
	name   string
	mkplot func() *plot.Plot
}{
	{"example_errpoints", Example_errpoints},
	{"example_stackedAreaChart", Example_stackedAreaChart},
}

func main() {
	for _, ex := range examples {
		drawEps(ex.name, ex.mkplot)
		drawSvg(ex.name, ex.mkplot)
		drawPng(ex.name, ex.mkplot)
		drawTiff(ex.name, ex.mkplot)
		drawJpg(ex.name, ex.mkplot)
		drawPdf(ex.name, ex.mkplot)
	}
}

func drawEps(name string, mkplot func() *plot.Plot) {
	if err := mkplot().Save(4, 4, name+".eps"); err != nil {
		panic(err)
	}
}

func drawPdf(name string, mkplot func() *plot.Plot) {
	if err := mkplot().Save(4, 4, name+".pdf"); err != nil {
		panic(err)
	}
}

func drawSvg(name string, mkplot func() *plot.Plot) {
	if err := mkplot().Save(4, 4, name+".svg"); err != nil {
		panic(err)
	}
}

func drawPng(name string, mkplot func() *plot.Plot) {
	if err := mkplot().Save(4, 4, name+".png"); err != nil {
		panic(err)
	}
}

func drawTiff(name string, mkplot func() *plot.Plot) {
	if err := mkplot().Save(4, 4, name+".tiff"); err != nil {
		panic(err)
	}
}

func drawJpg(name string, mkplot func() *plot.Plot) {
	if err := mkplot().Save(4, 4, name+".jpg"); err != nil {
		panic(err)
	}
}

// Example_errpoints draws some error points.
func Example_errpoints() *plot.Plot {
	rnd := rand.New(rand.NewSource(1))

	// Get some random data.
	n, m := 5, 10
	pts := make([]plotter.XYer, n)
	for i := range pts {
		xys := make(plotter.XYs, m)
		pts[i] = xys
		center := float64(i)
		for j := range xys {
			xys[j].X = center + (rnd.Float64() - 0.5)
			xys[j].Y = center + (rnd.Float64() - 0.5)
		}
	}

	plt, err := plot.New()
	if err != nil {
		panic(err)
	}

	mean95, err := plotutil.NewErrorPoints(plotutil.MeanAndConf95, pts...)
	if err != nil {
		panic(err)
	}
	medMinMax, err := plotutil.NewErrorPoints(plotutil.MedianAndMinMax, pts...)
	if err != nil {
		panic(err)
	}
	plotutil.AddLinePoints(plt,
		"mean and 95% confidence", mean95,
		"median and minimum and maximum", medMinMax)
	if err := plotutil.AddErrorBars(plt, mean95, medMinMax); err != nil {
		panic(err)
	}
	if err := plotutil.AddScatters(plt, pts[0], pts[1], pts[2], pts[3], pts[4]); err != nil {
		panic(err)
	}

	return plt
}

type stackValues struct{ vs []plotter.Values }

func (n stackValues) Len() int { return n.vs[0].Len() }
func (n stackValues) Value(i int) float64 {
	sum := 0.0
	for _, v := range n.vs {
		sum += v.Value(i)
	}
	return sum
}

// An example of making a stacked area chart.
func Example_stackedAreaChart() *plot.Plot {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Example: Software Version Comparison"
	p.X.Label.Text = "Date"
	p.Y.Label.Text = "Users (in thousands)"

	p.Legend.Top = true
	p.Legend.Left = true

	vals := []plotter.Values{
		plotter.Values{0.02, 0.015, 0, 0, 0, 0, 0},
		plotter.Values{0, 0.48, 0.36, 0.34, 0.32, 0.32, 0.28},
		plotter.Values{0, 0, 0.87, 1.4, 0.64, 0.32, 0.28},
		plotter.Values{0, 0, 0, 1.26, 0.34, 0.12, 0.09},
		plotter.Values{0, 0, 0, 0, 2.48, 2.68, 2.13},
		plotter.Values{0, 0, 0, 0, 0, 1.32, 0.54},
		plotter.Values{0, 0, 0, 0, 0, 0.68, 5.67},
	}

	err = plotutil.AddStackedAreaPlots(p, plotter.Values{2007, 2008, 2009, 2010, 2011, 2012, 2013},
		"Version 3.0",
		stackValues{vs: vals[0:7]},
		"Version 2.1",
		stackValues{vs: vals[0:6]},
		"Version 2.0.1",
		stackValues{vs: vals[0:5]},
		"Version 2.0",
		stackValues{vs: vals[0:4]},
		"Version 1.1",
		stackValues{vs: vals[0:3]},
		"Version 1.0",
		stackValues{vs: vals[0:2]},
		"Beta",
		stackValues{vs: vals[0:1]},
	)

	if err != nil {
		panic(err)
	}

	return p
}
