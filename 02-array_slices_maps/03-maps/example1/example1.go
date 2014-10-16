// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/voXAyiydFf

// Sample program to show how to declare, initalize and iterate
// over a map. Shows how iterating over a map is random.
package main

import (
	"fmt"
)

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// main is the entry point for the application.
func main() {
	// Declare and make a map that stores values
	// of type user with a key of type integer.
	users := make(map[int]user)

	// Add key/value pairs to the map.
	users[1] = user{"Roy", "Rob"}
	users[2] = user{"Henry", "Ford"}
	users[3] = user{"Mickey", "Mouse"}
	users[4] = user{"Michael", "Jackson"}

	// Iterate over the map.
	for key, value := range users {
		fmt.Println(key, value)
	}

	fmt.Println()

	// Iterate over the map and notice the
	// results are different.
	for key := range users {
		fmt.Println(key)
	}
}
