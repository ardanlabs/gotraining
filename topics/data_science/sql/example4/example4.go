// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example4

// Sample program to modify data in a database.
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
	db, err := sql.Open("sqlite3", "../data/iris.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Update some values.
	res, err := db.Exec("UPDATE iris SET species = 'setosa' WHERE species = 'Iris-setosa'")
	if err != nil {
		log.Fatal(err)
	}

	// See how many rows where updated.
	rowCount, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	// Output the number of rows to standard out.
	log.Printf("affected = %d\n", rowCount)
}
