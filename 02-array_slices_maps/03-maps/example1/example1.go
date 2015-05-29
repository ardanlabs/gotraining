// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/ZgPw0LU_Nv

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
	// of type user with a key of type string.
	users := make(map[string]user)

	// Add key/value pairs to the map.
	users["Rob"] = user{"Roy", "Rob"}
	users["Ford"] = user{"Henry", "Ford"}
	users["Mouse"] = user{"Mickey", "Mouse"}
	users["Jackson"] = user{"Michael", "Jackson"}

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
