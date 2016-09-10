// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./template1

// Sample program to retrieve results from a database.
package main

import (
	"database/sql"
	"log"

	// go-sqlite3 is the libary that allows us to connect
	// to sqlite with databases/sql.
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// Open a database value.  Specify the sqlite3 driver
	// for databases/sql.
	db, err := sql.Open("sqlite3", "../../data/iris.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Query the database.

	// Iterate over the rows, sending the results to
	// standard out.

	// Check for errors after we are done iterating over rows.
}
