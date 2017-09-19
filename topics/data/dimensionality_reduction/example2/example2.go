// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example2

// Sample program to visualize the impact of dimensionality reduction.
package main

import (
	"image/color"
	"log"
	"os"

	// These use the deprecated import because of a dependency on mat64 in gota
	"github.com/gonum/matrix/mat64"
	"github.com/gonum/stat"
	"github.com/kniren/gota/dataframe"

	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {

	// Open the iris dataset file.
	f, err := os.Open("../data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Parse the CSV file into a dataframe.
	irisDF := dataframe.ReadCSV(f)

	// Form the matrix.
	mat := irisDF.Select([]string{"sepal_length", "sepal_width", "petal_length", "petal_width"}).Matrix()

	// Calculate the principal component direction vectors
	// and variances.
	var pc stat.PC
	ok := pc.PrincipalComponents(mat, nil)
	if !ok {
		log.Fatal("Could not calculate principal components")
	}

	// Get the prinipal components and the corresponding vectors.
	var vars []float64
	var vecs *mat64.Dense
	vars = pc.Vars(vars)
	vecs = pc.Vectors(vecs)

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
