// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example1

// Sample program to generate a histogram of a normal distribution.
package main

import (
	"image/color"
	"log"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"
	"github.com/gonum/stat/distuv"
)

func main() {

	// Create a normal distribution.
	normDist := distuv.Normal{
		Mu:    0,
		Sigma: 1,
	}

	// Let's draw 100 random points from the normal distribution.
	var samples [100]float64
	for i := 0; i < 100; i++ {
		samples[i] = normDist.Rand()
	}

	// Put these samples into t a plotter.Values variable for plotting.
	v := make(plotter.Values, 100)
	for i := range v {
		v[i] = samples[i]
	}

	// Make a plot and set its title.
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	p.Title.Text = "Histogram of a Normal Distribution"

	// Create a histogram of our values drawn
	// from the standard normal.
	h, err := plotter.NewHist(v, 16)
	if err != nil {
		log.Fatal(err)
	}

	// Normalize the histogram.
	h.Normalize(1)

	// Add the histogram to the plot.
	p.Add(h)

	// Overlay The normal distribution function for comparison.
	norm := plotter.NewFunction(normDist.Prob)
	norm.Color = color.RGBA{R: 255, A: 255}
	norm.Width = vg.Points(2)
	p.Add(norm)

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "normal_hist.png"); err != nil {
		log.Fatal(err)
	}
}
