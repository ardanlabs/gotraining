// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example1

// Sample program to connect to and ping a database connection.
package main

import (
	"database/sql"
	"fmt"
	"log"

	// go-sqlite3 is the libary that allows us to connect
	// to sqlite with databases/sql.
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// Open a database value.  Specify the sqlite3 driver
	// for databases/sql.
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// sql.Open() does not establish any connections to the
	// database.  It just prepares the database connection value
	// for later use.  To make sure the database is available and
	// accessible, we will use db.Ping().
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Able to connect to the sqlite database!")
}
