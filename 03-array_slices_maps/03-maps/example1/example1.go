// Sample program to show how to declare, initalize and iterate
// over a map. Shows how iterating over a map is random.
package main

import (
	"fmt"
)

// User defines a user in the program.
type User struct {
	Name  string
	Email string
}

// main is the entry point for the application.
func main() {
	// Declare and make a map that stores values
	// of type User with a key of type integer.
	users := make(map[int]User)

	// Add key/value pairs to the map.
	users[1] = User{"Roy", "Rob"}
	users[2] = User{"Henry", "Ford"}
	users[3] = User{"Mickey", "Mouse"}
	users[4] = User{"Michael", "Jackson"}

	// Iterate over the map.
	for key, value := range users {
		fmt.Println(key, value)
	}

	fmt.Println()

	// Iterate over the map and notice the
	// results are different.
	for key, value := range users {
		fmt.Println(key, value)
	}
}
