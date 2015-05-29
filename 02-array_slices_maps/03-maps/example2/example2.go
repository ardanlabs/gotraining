// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/hC1o26x7Q5

// Sample program to show how to declare and initalize a map
// using a map literal.
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
	// Declare and initalize the map with values.
	users := map[string]user{
		"Rob":     user{"Roy", "Rob"},
		"Ford":    user{"Henry", "Ford"},
		"Mouse":   user{"Mickey", "Mouse"},
		"Jackson": user{"Michael", "Jackson"},
	}

	// Iterate over the map.
	for key, value := range users {
		fmt.Println(key, value)
	}
}
