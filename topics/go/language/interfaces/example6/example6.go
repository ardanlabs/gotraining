// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show type assertions using the comma-ok idiom.
package main

import (
	"fmt"
	"log"
)

// user defines a user in our application.
type user struct {
	id   int
	name string
}

// finder represents the ability to find users.
type finder interface {
	find(id int) (*user, error)
}

// userDB defines a database we will access.
type userDB struct {
	host string
}

// find implements the finder interface using pointer semantics.
func (db *userDB) find(id int) (*user, error) {
	return &user{id: id, name: "Anna Walker"}, nil
}

// mockDB defines a mock database we will access.
type mockDB struct{}

// find implements the finder interface using pointer semantics.
func (db *mockDB) find(id int) (*user, error) {
	return &user{id: id, name: "Jacob Walker"}, nil
}

func main() {
	var db mockDB

	if err := run(&db); err != nil {
		log.Fatal(err)
	}
}

func run(f finder) error {
	u, err := f.find(1234)
	if err != nil {
		return err
	}
	fmt.Printf("Found user %+v\n", u)

	// If the concrete type value stored inside the interface value is of the
	// type *userDB, then "ok" will be true and "db" will be a copy of the
	// pointer stored inside the interface.
	if db, ok := f.(*userDB); ok {
		log.Println("queried", db.host)
	}

	return nil
}
