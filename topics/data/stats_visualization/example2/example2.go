// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example2

// Sample program to generate a histogram of the iris data variables.
package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"
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

	// Create a dataframe from the CSV file.
	// The types of the columns will be inferred.
	irisDF := dataframe.ReadCSV(bytes.NewReader(b.Bytes()))

	// Create a histogram for each of the feature columns in the dataset.
	for _, colName := range irisDF.Names() {

		// If the column is one of the feature columns, let's create
		// a histogram of the values.
		if colName != "species" {

			// Create a plotter.Values value and fill it with the
			// values from the respective column of the dataframe.
			v := make(plotter.Values, irisDF.Nrow())
			for i, floatVal := range irisDF.Col(colName).Float() {
				v[i] = floatVal
			}

			// Make a plot and set its title.
			p, err := plot.New()
			if err != nil {
				log.Fatal(err)
			}
			p.Title.Text = fmt.Sprintf("Histogram of a %s", colName)

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
			if err := p.Save(4*vg.Inch, 4*vg.Inch, colName+"_hist.png"); err != nil {
				log.Fatal(err)
			}
		}
	}
}
