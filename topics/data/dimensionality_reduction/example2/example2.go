// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example2

// Sample program to visualize the impact of dimensionality reduction.
package main

import (
	"bytes"
	"image/color"
	"log"

	"github.com/gonum/floats"
	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/plotutil"
	"github.com/gonum/plot/vg"
	"github.com/gonum/stat"
	"github.com/kniren/gota/dataframe"
	"github.com/pachyderm/pachyderm/src/client"
)

func main() {

	// Connect to Pachyderm on our localhost.  By default
	// Pachyderm will be exposed on port 30650.
	c, err := client.NewFromAddress("0.0.0.0:30650")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	// Get the Iris dataset from Pachyderm's data
	// versioning at the latest commit.
	var b bytes.Buffer
	if err := c.GetFile("iris", "master", "iris.csv", 0, 0, "", false, nil, &b); err != nil {
		log.Fatal()
	}

	// Parse the CSV file into a dataframe.
	irisDF := dataframe.ReadCSV(bytes.NewReader(b.Bytes()))

	// Form the matrix.
	mat := irisDF.Select([]string{"sepal_length", "sepal_width", "petal_length", "petal_width"}).Matrix()

	// Calculate the principal component direction vectors
	// and variances.
	_, vars, ok := stat.PrincipalComponents(mat, nil)
	if !ok {
		log.Fatal("Could not calculate principal components")
	}

	// Sum the eigenvalues (variances).
	total := floats.Sum(vars)

	// Calculate cumulative variance percentages for each sorted value.
	cumVar := make(plotter.Values, 4)
	var cumSum float64
	for idx, variance := range vars {
		cumSum += (variance / total) * 100.0
		cumVar[idx] = cumSum
	}

	// Create a bar plot to visualize the variance percentages.
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	p.X.Label.Text = "Principal components"
	p.Y.Label.Text = "Percent of variance captured"
	p.Y.Max = 110.0
	p.X.Max = 3.1
	p.X.Min = -0.1
	w := vg.Points(20)

	// Create the bars for the percent values.
	bars, err := plotter.NewBarChart(cumVar, w)
	if err != nil {
		log.Fatal(err)
	}
	bars.LineStyle.Width = vg.Length(0)
	bars.Color = plotutil.Color(0)

	// Format the bars.
	p.Add(bars)
	p.NominalX("One", "Two", "Three", "Four")

	// Plot a line at 100% for easy inspection.
	hundred := plotter.NewFunction(func(x float64) float64 { return 100.0 })
	hundred.Color = color.RGBA{B: 255, A: 255}
	hundred.Dashes = []vg.Length{vg.Points(2), vg.Points(2)}
	hundred.Width = vg.Points(2)
	p.Add(hundred)

	// Save the graph.
	if err := p.Save(4*vg.Inch, 5*vg.Inch, "barchart.png"); err != nil {
		log.Fatal(err)
	}
}
