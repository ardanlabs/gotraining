// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example4

// Sample program to save records to a CSV file.
package main

import (
	"encoding/csv"
	"log"
	"os"
)

// Define the data we want to save.
var data = [][]string{
	{"1.2", "1.3", "0.3", "0.12", "Iris-setosa"},
	{"1.0", "2.1", "0.4", "0.8", "Iris-setosa"},
	{"2.1", "8.2", "0.7", "0.2", "Iris-setosa"},
	{"3.2", "1.8", "0.2", "0.15", "Iris-versicolor"},
	{"2.5", "2.7", "0.5", "0.1", "Iris-versicolor"},
	{"1.7", "3.5", "1.0", "0.7", "Iris-virginica"},
	{"1.7", "3.1", "0.5", "0.2", "Iris-virginica"},
	{"1.1", "3.0", "0.2", "0.1", "Iris-virginica"},
}

func main() {

	// Create the output file.
	f, err := os.Create("output.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a CSV writer.
	w := csv.NewWriter(f)

	// Write all the records out to the file. Note, this can
	// also we done record by record, with a final call to
	// Flush().
	w.WriteAll(data)
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
