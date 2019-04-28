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

// userSVC is a service for dealing with users.
type userSVC struct {
	host string
}

// find implements the finder interface using pointer semantics.
func (*userSVC) find(id int) (*user, error) {
	return &user{id: id, name: "Anna Walker"}, nil
}

// mockSVC defines a mock service we will access.
type mockSVC struct{}

// find implements the finder interface using pointer semantics.
func (*mockSVC) find(id int) (*user, error) {
	return &user{id: id, name: "Jacob Walker"}, nil
}

func main() {
	var svc mockSVC

	if err := run(&svc); err != nil {
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
	// type *userSVC, then "ok" will be true and "svc" will be a copy of the
	// pointer stored inside the interface.
	if svc, ok := f.(*userSVC); ok {
		log.Println("queried", svc.host)
	}

	return nil
}
