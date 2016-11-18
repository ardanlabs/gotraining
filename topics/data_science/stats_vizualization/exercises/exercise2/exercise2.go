// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./exercise2

// Sample program to generate a histogram of diabetes bmi values.
package main

import (
	"io/ioutil"
	"log"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"
	"github.com/kniren/gota/data-frame"
)

func main() {

	// Pull in the CSV data.
	diabetesData, err := ioutil.ReadFile("../../data/diabetes.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Create a dataframe from the CSV string.
	// The types of the columns will be inferred.
	diabetesDF := df.ReadCSV(string(diabetesData))

	// Extract the bmi col as a slice of floats.
	bmi, err := diabetesDF.Col("bmi").Float()
	if err != nil {
		log.Fatal(err)
	}

	// Create a plotter.Values value and fill it with the
	// values from the respective column of the dataframe.
	v := make(plotter.Values, len(bmi))
	for i, val := range bmi {
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
