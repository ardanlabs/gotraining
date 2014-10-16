// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/0v_VHlYF7f

// Sample program to show how only types that can have
// equality defined on them can be a map key.
package main

import (
	"fmt"
)

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// users define a set of users.
type users []user

// main is the entry point for the application.
func main() {
	// Declare and make a map uses a slice
	// of users as the key.
	u := make(map[users]int)

	// ./example3.go:24: invalid map key type users

	// Iterate over the map.
	for key, value := range u {
		fmt.Println(key, value)
	}
}
