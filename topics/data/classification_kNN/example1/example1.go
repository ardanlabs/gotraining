// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example1

// Sample program to profile our data set.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {

	// Open the iris dataset file.
	f, err := os.Open("../data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Parse the file into a Gota dataframe.
	irisDF := dataframe.ReadCSV(f)

	// Create a histogram for each of the float columns in the dataset and
	// output summary statistics.
	for _, colName := range irisDF.Names() {

		if colName != "species" {

			// Create a plotter.Values value and fill it with the
			// values from the respective column of the dataframe.
			plotVals := make(plotter.Values, irisDF.Nrow())
			summaryVals := make([]float64, irisDF.Nrow())
			for i, floatVal := range irisDF.Col(colName).Float() {
				plotVals[i] = floatVal
				summaryVals[i] = floatVal
			}

			// Make a plot and set its title.
			p, err := plot.New()
			if err != nil {
				log.Fatal(err)
			}
			p.Title.Text = fmt.Sprintf("Histogram of a %s", colName)

			// Create a histogram of our values drawn
			// from the standard normal.
			h, err := plotter.NewHist(plotVals, 16)
			if err != nil {
				log.Fatal(err)
			}

			// Normalize the histogram.
			h.Normalize(1)

			// Add the histogram to the plot.
			p.Add(h)

			// Save the plot to a PNG file.
			if err := p.Save(4*vg.Inch, 4*vg.Inch, colName+"_hist.png"); err != nil {
				log.Fatal(err)
			}

			// Calculate the summary statistics.
			meanVal := stat.Mean(summaryVals, nil)
			maxVal := floats.Max(summaryVals)
			minVal := floats.Min(summaryVals)
			stdDevVal := stat.StdDev(summaryVals, nil)

			// Output the summary statistics.
			fmt.Printf("\n%s\n", colName)
			fmt.Printf("Mean: %0.2f\n", meanVal)
			fmt.Printf("Min: %0.2f\n", minVal)
			fmt.Printf("Max: %0.2f\n", maxVal)
			fmt.Printf("StdDev: %0.2f\n\n", stdDevVal)
		}
	}
}
