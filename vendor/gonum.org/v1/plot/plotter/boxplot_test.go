// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter

import (
	"fmt"
	"log"
	"math/rand"
	"testing"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/internal/cmpimg"
	"gonum.org/v1/plot/vg"
)

func ExampleBoxPlot() {
	rnd := rand.New(rand.NewSource(1))

	// Create the sample data.
	n := 100
	uniform := make(ValueLabels, n)
	normal := make(ValueLabels, n)
	expon := make(ValueLabels, n)
	for i := 0; i < n; i++ {
		uniform[i].Value = rnd.Float64()
		uniform[i].Label = fmt.Sprintf("%4.4f", uniform[i].Value)
		normal[i].Value = rnd.NormFloat64()
		normal[i].Label = fmt.Sprintf("%4.4f", normal[i].Value)
		expon[i].Value = rnd.ExpFloat64()
		expon[i].Label = fmt.Sprintf("%4.4f", expon[i].Value)
	}

	// Make boxes for our data and add them to the plot.
	uniBox, err := NewBoxPlot(vg.Points(20), 0, uniform)
	if err != nil {
		log.Panic(err)
	}
	normBox, err := NewBoxPlot(vg.Points(20), 1, normal)
	if err != nil {
		log.Panic(err)
	}
	expBox, err := NewBoxPlot(vg.Points(20), 2, expon)
	if err != nil {
		log.Panic(err)
	}

	// Make a vertical box plot.
	uniLabels, err := uniBox.OutsideLabels(uniform)
	if err != nil {
		log.Panic(err)
	}
	normLabels, err := normBox.OutsideLabels(normal)
	if err != nil {
		log.Panic(err)
	}
	expLabels, err := expBox.OutsideLabels(expon)
	if err != nil {
		log.Panic(err)
	}

	p1, err := plot.New()
	if err != nil {
		log.Panic(err)
	}
	p1.Title.Text = "Vertical Box Plot"
	p1.Y.Label.Text = "plotter.Values"
	p1.Add(uniBox, uniLabels, normBox, normLabels, expBox, expLabels)

	// Set the X axis of the plot to nominal with
	// the given names for x=0, x=1 and x=2.
	p1.NominalX("Uniform\nDistribution", "Normal\nDistribution",
		"Exponential\nDistribution")

	err = p1.Save(200, 200, "testdata/verticalBoxPlot.png")
	if err != nil {
		log.Panic(err)
	}

	// Now, make the same plot but horizontal.
	normBox.Horizontal = true
	expBox.Horizontal = true
	uniBox.Horizontal = true
	// We can use the same plotters but the labels need to be recreated.
	uniLabels, err = uniBox.OutsideLabels(uniform)
	if err != nil {
		log.Panic(err)
	}
	normLabels, err = normBox.OutsideLabels(normal)
	if err != nil {
		log.Panic(err)
	}
	expLabels, err = expBox.OutsideLabels(expon)
	if err != nil {
		log.Panic(err)
	}

	p2, err := plot.New()
	if err != nil {
		log.Panic(err)
	}
	p2.Title.Text = "Horizontal Box Plot"
	p2.X.Label.Text = "plotter.Values"

	p2.Add(uniBox, uniLabels, normBox, normLabels, expBox, expLabels)

	// Set the Y axis of the plot to nominal with
	// the given names for y=0, y=1 and y=2.
	p2.NominalY("Uniform\nDistribution", "Normal\nDistribution",
		"Exponential\nDistribution")

	err = p2.Save(200, 200, "testdata/horizontalBoxPlot.png")
	if err != nil {
		log.Panic(err)
	}

	// Now, make a grouped box plot.
	p3, err := plot.New()
	if err != nil {
		log.Panic(err)
	}
	p3.Title.Text = "Box Plot"
	p3.Y.Label.Text = "plotter.Values"

	w := vg.Points(20)
	for x := 0.0; x < 3.0; x++ {
		b0, err := NewBoxPlot(w, x, uniform)
		if err != nil {
			log.Panic(err)
		}
		b0.Offset = -w - vg.Points(3)
		b1, err := NewBoxPlot(w, x, normal)
		if err != nil {
			log.Panic(err)
		}
		b2, err := NewBoxPlot(w, x, expon)
		if err != nil {
			log.Panic(err)
		}
		b2.Offset = w + vg.Points(3)
		p3.Add(b0, b1, b2)
	}
	// Add a GlyphBox plotter for debugging.
	p3.Add(NewGlyphBoxes())

	// Set the X axis of the plot to nominal with
	// the given names for x=0, x=1 and x=2.
	p3.NominalX("Group 0", "Group 1", "Group 2")
	err = p3.Save(300, 300, "testdata/groupedBoxPlot.png")
	if err != nil {
		log.Panic(err)
	}
}

func TestBoxPlot(t *testing.T) {
	cmpimg.CheckPlot(ExampleBoxPlot, t, "verticalBoxPlot.png",
		"horizontalBoxPlot.png", "groupedBoxPlot.png")
}
