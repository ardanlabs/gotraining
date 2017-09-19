// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example2

// Sample program to investigate correlations between our target and our features.
package main

import (
	"image/color"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {

	// Open the diabetes dataset file.
	f, err := os.Open("../data/diabetes.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a dataframe from the CSV file.
	// The types of the columns will be inferred.
	diabetesDF := dataframe.ReadCSV(f)

	// Extract the target column.
	yVals := diabetesDF.Col("y").Float()

	// Create a scatter plot for each of the features in the dataset.
	for _, colName := range diabetesDF.Names() {

		// pts will hold the values for plotting
		pts := make(plotter.XYs, diabetesDF.Nrow())

		// Fill pts with data.
		for i, floatVal := range diabetesDF.Col(colName).Float() {
			pts[i].X = floatVal
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
