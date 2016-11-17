// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example2

// Sample program to visualize the impact of dimensionality reduction.
package main

import (
	"encoding/csv"
	"fmt"
	"image/color"
	"log"
	"os"
	"strconv"

	"github.com/gonum/floats"
	"github.com/gonum/matrix/mat64"
	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/plotutil"
	"github.com/gonum/plot/vg"
	"github.com/gonum/stat"
)

func main() {

	// Open the iris dataset file.
	f, err := os.Open("../data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(f)
	reader.FieldsPerRecord = 5

	// Read in all of the CSV records
	rawCSVData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// floatData will hold all the float values that will eventually be
	// used to form out matrix.
	floatData := make([]float64, 4*len(rawCSVData))

	// dataIndex will track the current index of the matrix values.
	var dataIndex int

	// Sequentially move the rows into a slice of floats.
	for _, record := range rawCSVData {

		// Loop over the float columns.
		for i := 0; i < 4; i++ {

			// Convert the value to a float.
			val, err := strconv.ParseFloat(record[i], 64)
			if err != nil {
				log.Fatal(fmt.Errorf("Could not parse float value"))
			}

			// Add the float value to the slice of floats.
			floatData[dataIndex] = val
			dataIndex++
		}
	}

	// Form the matrix.
	mat := mat64.NewDense(len(rawCSVData), 4, floatData)

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
