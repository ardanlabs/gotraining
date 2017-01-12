// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./exercise1

// Sample program to generate a box plot of diabetes bmi values.
package main

import (
	"bytes"
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

	// Get the diabetes dataset from Pachyderm's data
	// versioning at the latest commit.
	var b bytes.Buffer
	if err := c.GetFile("diabetes", "master", "diabetes.csv", 0, 0, "", false, nil, &b); err != nil {
		log.Fatal()
	}

	// Create a dataframe from the CSV file.
	// The types of the columns will be inferred.
	diabetesDF := dataframe.ReadCSV(bytes.NewReader(b.Bytes()))

	// Create the plot and set its title and axis label.
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	p.Y.Label.Text = "Values"

	// Create the box for our data.
	w := vg.Points(50)

	// Create a plotter.Values value and fill it with the
	// values from the respective column of the dataframe.
	v := make(plotter.Values, diabetesDF.Nrow())
	for i, val := range diabetesDF.Col("bmi").Float() {
		v[i] = val
	}

	// Add the data to the plot.
	box, err := plotter.NewBoxPlot(w, 0, v)
	if err != nil {
		log.Fatal(err)
	}
	p.Add(box)

	// Set the X axis of the plot to nominal with
	// the given names for x=0, x=1, etc.
	p.NominalX("bmi")

	if err := p.Save(4*vg.Inch, 8*vg.Inch, "boxplots.png"); err != nil {
		log.Fatal(err)
	}
}
