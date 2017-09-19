// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter

import (
	"log"
	"math/rand"
	"testing"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/internal/cmpimg"
	"gonum.org/v1/plot/vg"
)

func ExampleQuartPlot() {
	rnd := rand.New(rand.NewSource(1))

	// Create the example data.
	n := 100
	uniform := make(Values, n)
	normal := make(Values, n)
	expon := make(Values, n)
	for i := 0; i < n; i++ {
		uniform[i] = rnd.Float64()
		normal[i] = rnd.NormFloat64()
		expon[i] = rnd.ExpFloat64()
	}

	// Create the QuartPlots
	qp1, err := NewQuartPlot(0, uniform)
	if err != nil {
		log.Panic(err)
	}
	qp2, err := NewQuartPlot(1, normal)
	if err != nil {
		log.Panic(err)
	}
	qp3, err := NewQuartPlot(2, expon)
	if err != nil {
		log.Panic(err)
	}

	// Create a vertical plot
	p1, err := plot.New()
	if err != nil {
		log.Panic(err)
	}
	p1.Title.Text = "Quartile Plot"
	p1.Y.Label.Text = "plotter.Values"
	p1.Add(qp1, qp2, qp3)

	// Set the X axis of the plot to nominal with
	// the given names for x=0, x=1 and x=2.
	p1.NominalX("Uniform\nDistribution", "Normal\nDistribution",
		"Exponential\nDistribution")

	err = p1.Save(200, 200, "testdata/verticalQuartPlot.png")
	if err != nil {
		log.Panic(err)
	}

	// Create a horizontal plot
	qp1.Horizontal = true
	qp2.Horizontal = true
	qp3.Horizontal = true

	p2, err := plot.New()
	if err != nil {
		log.Panic(err)
	}
	p2.Title.Text = "Quartile Plot"
	p2.X.Label.Text = "plotter.Values"
	p2.Add(qp1, qp2, qp3)

	// Set the Y axis of the plot to nominal with
	// the given names for y=0, y=1 and y=2.
	p2.NominalY("Uniform\nDistribution", "Normal\nDistribution",
		"Exponential\nDistribution")

	err = p2.Save(200, 200, "testdata/horizontalQuartPlot.png")
	if err != nil {
		log.Panic(err)
	}

	// Now, create a grouped quartile plot.

	p3, err := plot.New()
	if err != nil {
		log.Panic(err)
	}
	p3.Title.Text = "Box Plot"
	p3.Y.Label.Text = "plotter.Values"

	w := vg.Points(10)
	for x := 0.0; x < 3.0; x++ {
		b0, err := NewQuartPlot(x, uniform)
		if err != nil {
			log.Panic(err)
		}
		b0.Offset = -w
		b1, err := NewQuartPlot(x, normal)
		if err != nil {
			log.Panic(err)
		}
		b2, err := NewQuartPlot(x, expon)
		if err != nil {
			log.Panic(err)
		}
		b2.Offset = w
		p3.Add(b0, b1, b2)
	}
	p3.Add(NewGlyphBoxes())

	p3.NominalX("Group 0", "Group 1", "Group 2")

	err = p3.Save(200, 200, "testdata/groupedQuartPlot.png")
	if err != nil {
		log.Panic(err)
	}
}

func TestQuartPlot(t *testing.T) {
	cmpimg.CheckPlot(ExampleQuartPlot, t, "verticalQuartPlot.png",
		"horizontalQuartPlot.png",
		"groupedQuartPlot.png")
}
