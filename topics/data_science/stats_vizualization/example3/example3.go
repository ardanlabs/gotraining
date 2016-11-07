// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example3

// Sample program to generate a box plot of example distributions.
package main

import (
	"log"
	"math/rand"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"
)

func main() {

	// Get some random data from three different distributions
	// to display in our plot.
	rand.Seed(int64(0))
	n := 100
	uniform := make(plotter.Values, n)
	normal := make(plotter.Values, n)
	expon := make(plotter.Values, n)
	for i := 0; i < n; i++ {
		uniform[i] = rand.Float64()
		normal[i] = rand.NormFloat64()
		expon[i] = rand.ExpFloat64()
	}

	// Create the plot and set its title and axis label.
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "Box plots"
	p.Y.Label.Text = "Values"

	// Make boxes for our data and add them to the plot.
	w := vg.Points(50)
	b0, err := plotter.NewBoxPlot(w, 0, uniform)
	if err != nil {
		log.Fatal(err)
	}
	b1, err := plotter.NewBoxPlot(w, 1, normal)
	if err != nil {
		log.Fatal(err)
	}
	b2, err := plotter.NewBoxPlot(w, 2, expon)
	if err != nil {
		log.Fatal(err)
	}
	p.Add(b0, b1, b2)

	// Set the X axis of the plot to nominal with
	// the given names for x=0, x=1 and x=2.
	p.NominalX("Uniform\nDistribution", "Normal\nDistribution",
		"Exponential\nDistribution")

	if err := p.Save(6*vg.Inch, 8*vg.Inch, "boxplot.png"); err != nil {
		log.Fatal(err)
	}
}
