// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/9zQuoGSuTe

// Sample program to show how to declare and initialize a map
// using a map literal.
package main

import "fmt"

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// main is the entry point for the application.
func main() {
	// Declare and initialize the map with values.
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
