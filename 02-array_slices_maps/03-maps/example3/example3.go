// http://play.golang.org/p/1eZogI1d_o

// Sample program to show how only types that can have
// equality defined on them can be a map key.
package main

import (
	"fmt"
)

// User defines a user in the program.
type User struct {
	Name  string
	Email string
}

// Users define a set of users.
type Users []User

// main is the entry point for the application.
func main() {
	// Declare and make a map uses a slice
	// of users as the key.
	users := make(map[Users]int)

	// ./example3.go:22: invalid map key type Users

	// Iterate over the map.
	for key, value := range users {
		fmt.Println(key, value)
	}
}
