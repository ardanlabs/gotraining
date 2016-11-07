// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example3

// Sample program to register of CSV file as an in-memory
// SQL database and execute SQL queries on the CSV.
package main

import (
	"fmt"
	"log"

	"github.com/go-hep/csvutil"
)

func main() {

	// Register the CSV file as a table.
	fileName := "../data/iris.csv"
	tbl, err := csvutil.Open(fileName)
	if err != nil {
		log.Fatalf("could not open %s: %v\n", fileName, err)
	}
	defer tbl.Close()

	// Specify the delimiter and comment character.
	tbl.Reader.Comma = ','
	tbl.Reader.Comment = '#'

	// Read in the first 10 non-header rows.
	rows, err := tbl.ReadRows(1, 11)
	if err != nil {
		log.Fatalf("could read rows [1, 11): %v\n", err)
	}
	defer rows.Close()

	// Scan the rows and read each row into a struct. Output
	// only the Petal measure and Species to standard out.
	var irow int
	for rows.Next() {
		data := struct {
			SepalLength float64
			SepalWidth  float64
			PetalLength float64
			PetalWidth  float64
			Species     string
		}{}
		if err = rows.Scan(&data); err != nil {
			log.Fatalf("error reading row %d: %v\n", irow, err)
		}
		fmt.Printf("petal length: %.2f, petal width: %.2f, species: %s\n",
			data.PetalLength, data.PetalWidth, data.Species)
	}
	if err = rows.Err(); err != nil {
		log.Fatalf("error: %v\n", err)
	}
}
