// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example2

// Sample program to visualize the impact of dimensionality reduction.
package main

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"log"

	"github.com/gonum/floats"
	"github.com/gonum/matrix/mat64"
	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/plotutil"
	"github.com/gonum/plot/vg"
	"github.com/gonum/stat"
	"github.com/kniren/gota/data-frame"
)

func main() {

	// Pull in the CSV data.
	irisData, err := ioutil.ReadFile("../data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Create a dataframe from the CSV string.
	// The types of the columns will be inferred.
	irisDF := df.ReadCSV(string(irisData))

	// Sequentially move the columns into a slice of floats.
	floatData := make([]float64, 4*irisDF.Nrow())
	var dataIndex int
	for colIndex, colName := range irisDF.Names() {

		// If the column is one of the float columns, move it
		// into the slice of floats.
		if colIndex < 4 {

			// Extract the columns as a slice of floats.
			floatCol, ok := irisDF.Col(colName).Elements.(df.FloatElements)
			if !ok {
				log.Fatal(fmt.Errorf("Could not parse float column."))
			}

			// Append the float values to floatData.
			for _, floatVal := range floatCol {
				floatData[dataIndex] = *floatVal.Float()
				dataIndex++
			}
		}
	}

	// Form the matrix.
	mat := mat64.NewDense(irisDF.Nrow(), 4, floatData)

	// Calculate the principal component direction vectors
	// and variances.
	_, vars, ok := stat.PrincipalComponents(mat, nil)
	if !ok {
		log.Fatal(fmt.Errorf("Could not calculate prinicple components"))
	}

	// Sum the eignvalues (variances).
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
	p.X.Label.Text = "Principle components"
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
