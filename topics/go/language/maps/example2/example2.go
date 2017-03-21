// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to declare and initialize a map
// using a map literal and delete a key.
package main

import "fmt"

// user defines a user in the program.
type user struct {
	name    string
	surname string
}

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

	// Delete the Roy key.
	delete(users, "Roy")

	fmt.Println("=================")

	// Find the Roy key.
	u, found := users["Roy"]

	// Display the value and found flag.
	fmt.Println("Roy", found, u)
}
