// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show the syntax of type assertions.
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

func main() {
	db := userDB{
		host: "localhost:3434",
	}

	if err := run(&db); err != nil {
		log.Fatal(err)
	}
}

// run performs the find operation against the concrete data that
// is passed into the call.
func run(f finder) error {
	u, err := f.find(1234)
	if err != nil {
		return err
	}
	fmt.Printf("Found user %+v\n", u)

	// Ideally the finder abstraction would encompass all of
	// the behavior you care about. But what if, for some reason,
	// you really need to get to the concrete value stored inside
	// the interface?

	// Can you access the "host" field from the concrete userDB type pointer
	// that is stored inside this interface variable? No, not directly.
	// All you know is the data has a method named "find".
	// ./example5.go:61:26: f.host undefined (type finder has no field or method host)
	log.Println("queried", f.host)

	// You can use a type assertion to get a copy of the userDB pointer
	// that is stored inside the interface.
	db := f.(*userDB)
	log.Println("queried", db.host)

	return nil
}
