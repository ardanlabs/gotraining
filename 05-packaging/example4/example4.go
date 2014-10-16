// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// Sample program to show how unexported fields from an exported struct
// type can't be accessed directly.
package main

import (
	"fmt"

	"github.com/ArdanStudios/gotraining/05-packaging/example4/animals"
)

// main is the entry point for the application.
func main() {
	// Create a value of type Dog from the animals package.
	// This will NOT compile.

	// ./example4.go:20: unknown animals.Dog field 'age' in struct literal
	dog := animals.Dog{
		Name:         "Chole",
		BarkStrength: 10,
		age:          5,
	}

	fmt.Printf("Dog: %#v\n", dog)
}
