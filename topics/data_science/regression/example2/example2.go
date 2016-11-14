// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example2

// Sample program to investigate correlations between our target and our features.
package main

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"log"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"
	"github.com/kniren/gota/data-frame"
)

func main() {

	// Pull in the CSV data.
	diabetesData, err := ioutil.ReadFile("../data/diabetes.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Create a dataframe from the CSV string.
	// The types of the columns will be inferred.
	diabetesDF := df.ReadCSV(string(diabetesData))

	// Extract the target column.
	yCol, ok := diabetesDF.Col("y").Elements.(df.FloatElements)
	if !ok {
		log.Fatal(fmt.Errorf("Could not parse y values"))
	}

	// Convert the target column to a slice of floats.
	yVals := make([]float64, len(yCol))
	for i, yVal := range yCol {
		yVals[i] = *yVal.Float()
	}

	// Create a scatter plot for each of the features in the dataset.
	for _, colName := range diabetesDF.Names() {

		// Extract the columns as a slice of floats.
		floatCol, ok := diabetesDF.Col(colName).Elements.(df.FloatElements)
		if !ok {
			log.Fatal(fmt.Errorf("Could not parse float column."))
		}

		// pts will hold the values for plotting
		pts := make(plotter.XYs, len(floatCol))

		// Fill pts with data.
		for i, floatVal := range floatCol {
			pts[i].X = *floatVal.Float()
			pts[i].Y = yVals[i]
		}

		// Create the plot.
		p, err := plot.New()
		if err != nil {
			log.Fatal(err)
		}
		p.X.Label.Text = colName
		p.Y.Label.Text = "y"
		p.Add(plotter.NewGrid())

		s, err := plotter.NewScatter(pts)
		if err != nil {
			log.Fatal(err)
		}
		s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
		s.GlyphStyle.Radius = vg.Points(3)

		// Save the plot to a PNG file.
		p.Add(s)
		if err := p.Save(4*vg.Inch, 4*vg.Inch, colName+"_scatter.png"); err != nil {
			log.Fatal(err)
		}
	}
}
