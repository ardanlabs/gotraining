// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how only types that can have
// equality defined on them can be a map key.
package main

import "fmt"

// user defines a user in the program.
type user struct {
	name    string
	surname string
}

// users define a set of users.
type users []user

func main() {

	// Declare and make a map uses a slice of users as the key.
	u := make(map[users]int)

	// ./example3.go:22: invalid map key type users

	// Iterate over the map.
	for key, value := range u {
		fmt.Println(key, value)
	}
}
