// Sample program to show how to create values from exported types with
// embedded types.
package main

import (
	"fmt"

	"github.com/ArdanStudios/gotraining/05-packaging/example5/animals"
)

// main is the entry point for the application.
func main() {
	// Create a value of type Dog from the animals package.
	dog := animals.Dog{
		Animal: animals.Animal{
			Name: "Chole",
			Age:  1,
		},
		BarkStrength: 10,
	}

	fmt.Printf("Dog: %#v\n", dog)
}
