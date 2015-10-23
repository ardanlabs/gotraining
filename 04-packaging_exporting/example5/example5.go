// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to create values from exported types with
// embedded unexported types.
package main

import (
	"fmt"

	"github.com/ardanlabs/gotraining/04-packaging_exporting/example5/animals"
)

// main is the entry point for the application.
func main() {
	/// Create a value of type Dog from the animals package.
	dog := animals.Dog{
		BarkStrength: 10,
	}

	// Set the exported fields from the unexported animal inner type.
	dog.Name = "Chole"
	dog.Age = 1

	fmt.Printf("Dog: %#v\n", dog)
}
