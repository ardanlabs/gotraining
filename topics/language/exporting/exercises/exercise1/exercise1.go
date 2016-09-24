// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Create a package named toy with a single exported struct type named Toy. Add
// the exported fields Name and Weight. Then add two unexported fields named
// onHand and sold. Declare a factory function called New to create values of
// type toy and accept parameters for the exported fields. Then declare methods
// that return and update values for the unexported fields.
//
// Create a program that imports the toy package. Use the New function to create a
// value of type toy. Then use the methods to set the counts and display the
// field values of that toy value.
package main

import (
	"fmt"

	"github.com/ardanlabs/gotraining/topics/language/exporting/exercises/exercise1/toy"
)

func main() {

	// Create a value of type toy.
	t := toy.New("Bat", 28)

	// Update the counts.
	t.UpdateOnHand(100)
	t.UpdateSold(2)

	// Display each field separately.
	fmt.Println("Name", t.Name)
	fmt.Println("Weight", t.Weight)
	fmt.Println("OnHand", t.OnHand())
	fmt.Println("Sold", t.Sold())
}
