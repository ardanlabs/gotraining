// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how unexported fields from an exported struct
// type can't be accessed directly.
package main

import (
	"fmt"

	"github.com/ardanlabs/gotraining/04-packaging_exporting/example4/animals"
)

// main is the entry point for the application.
func main() {
	// Create a value of type Dog from the animals package.
	dog := animals.Dog{
		Name:         "Chole",
		BarkStrength: 10,
		age:          5,
	}

	// ./example4.go:20: unknown animals.Dog field 'age' in struct literal

	fmt.Printf("Dog: %#v\n", dog)
}
