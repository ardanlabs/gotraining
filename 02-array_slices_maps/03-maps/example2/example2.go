// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/htR56-yyqC

// Sample program to show how to declare and initalize a map
// using a composite literal.
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
	// Declare and initalize the map with
	// values.
	users := map[int]user{
		1: user{"Roy", "Rob"},
		2: user{"Henry", "Ford"},
		3: user{"Mickey", "Mouse"},
		4: user{"Michael", "Jackson"},
	}

	// Iterate over the map.
	for key, value := range users {
		fmt.Println(key, value)
	}
}
