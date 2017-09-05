// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./exercise1.go

// Sample program to visualize the accuracy of models with various
// decision tree pruning parameters.
package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/sjwhitworth/golearn/base"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
)

func main() {

	// Open the iris dataset file.
	f, err := os.Open("../../data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create golearn instances from the iris data.
	irisData, err := CreateInstancesFromReader(f)
	if err != nil {
		log.Fatal(err)
	}

	// Define the parameters we will test.

	// Loop over the parameter choices.

	// Create a new plot.
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}

	// Label the graph.
	p.X.Label.Text = "Pruning parameter"
	p.Y.Label.Text = "Accuracy"

	// Create the XYs value.

	// Add our data to the plot.

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "accuracy_vs_pruning.png"); err != nil {
		log.Fatal(err)
	}
}

// CreateInstancesFromReader creates golearn instances from
// an io.Reader.
func CreateInstancesFromReader(r io.Reader) (*base.DenseInstances, error) {

	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(r)

	// Read in all of the CSV records
	reader.FieldsPerRecord = 5
	tmpData, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// Create the output file.
	f, err := os.Create("/tmp/iris.csv")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Create a CSV writer.
	w := csv.NewWriter(f)

	// Write all the records out to the file. Note, this can
	// also we done record by record, with a final call to
	// Flush().
	w.WriteAll(tmpData)
	if err := w.Error(); err != nil {
		return nil, err
	}

	// Read in the iris data set into golearn "instances".
	instances, err := base.ParseCSVToInstances("/tmp/iris.csv", true)
	if err != nil {
		return nil, err
	}

	return instances, nil
}
