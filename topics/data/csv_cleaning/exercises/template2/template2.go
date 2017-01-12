// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./template2

// Sample program to register of CSV file as an in-memory
// SQL database, sum float columns, and output a process CSV.
package main

import (
	"log"

	"github.com/go-hep/csvutil/csvdriver"
)

func main() {

	// Open the CSV file as a database table.
	db, err := csvdriver.Conn{
		File:    "../../data/iris.csv",
		Comment: '#',
		Comma:   ',',
	}.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Start a database transaction.
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Commit()

	// Query the CSV via a SQL statement.

	// Register the output file as a table.

	// Output the results of the query to a processed CSV file.
}
