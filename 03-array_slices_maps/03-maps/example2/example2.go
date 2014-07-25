// http://play.golang.org/p/Qy_nYK9zmb

// Sample program to show how to declare and initalize a map
// using a composite literal.
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
	// Declare and initalize the map with
	// values.
	users := map[int]User{
		1: User{"Roy", "Rob"},
		2: User{"Henry", "Ford"},
		3: User{"Mickey", "Mouse"},
		4: User{"Michael", "Jackson"},
	}

	// Iterate over the map.
	for key, value := range users {
		fmt.Println(key, value)
	}
}
