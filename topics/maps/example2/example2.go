// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// http://play.golang.org/p/EtPZoqHgaL

// Sample program to show how to declare and initialize a map
// using a map literal.
package main

import "fmt"

// user defines a user in the program.
type user struct {
	name    string
	surname string
}

// main is the entry point for the application.
func main() {
	// Declare and initialize the map with values.
	users := map[string]user{
		"Roy":     {"Rob", "Roy"},
		"Ford":    {"Henry", "Ford"},
		"Mouse":   {"Mickey", "Mouse"},
		"Jackson": {"Michael", "Jackson"},
	}

	// Iterate over the map.
	for key, value := range users {
		fmt.Println(key, value)
	}
}
