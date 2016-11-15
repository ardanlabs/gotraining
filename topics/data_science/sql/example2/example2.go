// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example2

// Sample program to load the iris dataset into a database.
package main

import (
	"database/sql"
	"encoding/csv"
	"log"
	"os"

	// go-sqlite3 is the libary that allows us to connect
	// to sqlite with databases/sql.
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// Remove the database if it exists.
	os.Remove("../data/iris.db")

	// Open the iris dataset file.
	csvFile, err := os.Open("../data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = 5

	// Read in all of the CSV records
	rawCSVData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Open a database value.  Specify the sqlite3 driver
	// for databases/sql.
	db, err := sql.Open("sqlite3", "../data/iris.db")
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
			species VARCHAR
		)`
	if _, err := db.Exec(createStmt); err != nil {
		log.Fatal(err)
	}

	// Create a transaction to load the data.
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// Prepare an insert statement.
	insertStmt, err := tx.Prepare("INSERT INTO iris VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer insertStmt.Close()

	// Add the insert statements to the transaction.
	for _, row := range rawCSVData {
		if _, err := insertStmt.Exec(row[0], row[1], row[2], row[3], row[4]); err != nil {
			log.Fatal(err)
		}
	}

	// Commit the transaction.
	tx.Commit()
}
