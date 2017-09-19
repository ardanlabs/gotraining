// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter

import (
	"image/color"
	"log"
	"math"
	"math/rand"
	"testing"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/internal/cmpimg"
	"gonum.org/v1/plot/vg"
)

// An example of making a histogram.
func ExampleHistogram() {
	rnd := rand.New(rand.NewSource(1))

	// stdNorm returns the probability of drawing a
	// value from a standard normal distribution.
	stdNorm := func(x float64) float64 {
		const sigma = 1.0
		const mu = 0.0
		const root2π = 2.50662827459517818309
		return 1.0 / (sigma * root2π) * math.Exp(-((x-mu)*(x-mu))/(2*sigma*sigma))
	}

	n := 10000
	vals := make(Values, n)
	for i := 0; i < n; i++ {
		vals[i] = rnd.NormFloat64()
	}

	p, err := plot.New()
	if err != nil {
		log.Panic(err)
	}
	p.Title.Text = "Histogram"
	h, err := NewHist(vals, 16)
	if err != nil {
		log.Panic(err)
	}
	h.Normalize(1)
	p.Add(h)

	// The normal distribution function
	norm := NewFunction(stdNorm)
	norm.Color = color.RGBA{R: 255, A: 255}
	norm.Width = vg.Points(2)
	p.Add(norm)

	err = p.Save(200, 200, "testdata/histogram.png")
	if err != nil {
		log.Panic(err)
	}
}

func TestHistogram(t *testing.T) {
	cmpimg.CheckPlot(ExampleHistogram, t, "histogram.png")
}
