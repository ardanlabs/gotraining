// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// Create a package named toy with a single unexported struct type named bat. Add
// the exported fields Height and Weight to the bat type. Then create an exported
// factory method called NewBat that returns pointers of type bat that are initialized
// to their zero value.
//
// Create a program that imports the toy package. Use the NewBat function to create a
// value of bat and populate the values of Height and Width. Then display the value of
// the bat variable.
package main

import (
	"fmt"

	"github.com/ArdanStudios/gotraining/05-packaging/exercises/exercise1/toy"
)

// main is the entry point for the application.
func main() {
	// Create a value of type bat.
	bat := toy.NewBat()
	bat.Height = 28
	bat.Weight = 16

	// Display the value.
	fmt.Println(bat)
}
