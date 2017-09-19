// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./exercise2

// Sample program to generate a histogram of diabetes bmi values.
package main

import (
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {

	// Open the diabetes dataset file.
	f, err := os.Open("../../data/diabetes.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a dataframe from the CSV file.
	// The types of the columns will be inferred.
	diabetesDF := dataframe.ReadCSV(f)

	// Create a plotter.Values value and fill it with the
	// values from the respective column of the dataframe.
	v := make(plotter.Values, diabetesDF.Nrow())
	for i, val := range diabetesDF.Col("bmi").Float() {
		v[i] = val
	}

	// Make a plot and set its title.
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	p.Title.Text = "Histogram of a BMI"

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

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "bmi_hist.png"); err != nil {
		log.Fatal(err)
	}

}
