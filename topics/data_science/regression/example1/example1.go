// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example1

// Sample program to profile our data set.
package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gonum/floats"
	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"
	"github.com/gonum/stat"
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

	// Create a histogram for each of the columns in the dataset and
	// output summary statistics.
	for _, colName := range diabetesDF.Names() {

		// Extract the columns as a slice of floats.
		floatCol, err := diabetesDF.Col(colName).Float()
		if err != nil {
			log.Fatal(err)
		}

		// Create a plotter.Values value and fill it with the
		// values from the respective column of the dataframe.
		plotVals := make(plotter.Values, len(floatCol))
		summaryVals := make([]float64, len(floatCol))
		for i, floatVal := range floatCol {
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
