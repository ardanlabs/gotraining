// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example2

// Sample program to load the iris dataset into a database.
package main

import (
	"database/sql"
	"io"
	"log"
	"os"

	"github.com/go-hep/csvutil"
	"github.com/lib/pq"
)

func main() {

	// Get my ElephantSQL postgres URL. I have it stored in
	// an environmental variable.
	pgURL := os.Getenv("PGURL")
	if pgURL == "" {
		log.Fatal("PGURL empty")
	}

	// Open a database value.  Specify the postgres driver
	// for databases/sql.
	db, err := sql.Open("postgres", pgURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the table for the iris dataset.
	createStmt := `
		CREATE TABLE iris (
			sepal_length NUMERIC,
			sepal_width NUMERIC,
			petal_length NUMERIC,
			petal_width NUMERIC,
			species TEXT
		)`
	if _, err := db.Exec(createStmt); err != nil {
		log.Fatal(err)
	}

	// Register the CSV file, holding the data we want to load, as a table.
	tbl, err := csvutil.Open("../data/iris.csv")
	if err != nil {
		log.Fatalf("could not open %s: %v\n", "iris.csv", err)
	}
	defer tbl.Close()

	// Specify the delimiter and comment character.
	tbl.Reader.Comma = ','
	tbl.Reader.Comment = '#'

	// Read in all the non-header rows in the CSV.
	rows, err := tbl.ReadRows(0, -1)
	if err != nil {
		log.Fatalf("could read csv rows: %v\n", err)
	}
	defer rows.Close()

	// Create a transaction to load the data.
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// Prepare an insert statement.
	insertStmt, err := tx.Prepare(pq.CopyIn("iris", "sepal_length", "sepal_width", "petal_length", "petal_width", "species"))
	if err != nil {
		log.Fatal(err)
	}

	// Add the insert statements to the transaction.
	for rows.Next() {

		// Define a struct that specifies the types of the columns.
		data := struct {
			SepalLength float64
			SepalWidth  float64
			PetalLength float64
			PetalWidth  float64
			Species     string
		}{}

		// Scan the row for the struct fields.
		if err = rows.Scan(&data); err != nil {
			log.Fatalf("error reading row: %v\n", err)
		}

		// Add the struct fields to the prepared insert statement.
		if _, err := insertStmt.Exec(data.SepalLength, data.SepalWidth, data.PetalLength, data.PetalWidth, data.Species); err != nil {
			log.Fatal(err)
		}
	}

	// Handle any errors from rows.Next().
	if err = rows.Err(); err != nil && err != io.EOF {
		log.Fatal(err)
	}

	// We have to close
	if err = insertStmt.Close(); err != nil {
		log.Fatal(err)
	}

	// Commit the transaction.
	tx.Commit()
}
