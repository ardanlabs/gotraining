// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program demonstrating that type assertions are a runtime and
// not compile time construct.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// car represents something you drive.
type car struct{}

// String implements the fmt.Stringer interface.
func (car) String() string {
	return "Vroom!"
}

// cloud represents somewhere you store information.
type cloud struct{}

// String implements the fmt.Stringer interface.
func (cloud) String() string {
	return "Big Data!"
}

// =============================================================================

func main() {

	// Seed the number random generator.
	rand.Seed(time.Now().UnixNano())

	// Create a slice of the Stringer interface values.
	mvs := []fmt.Stringer{
		car{},
		cloud{},
	}

	// Let's run this experiment ten times.
	for i := 0; i < 10; i++ {

		// Choose a random number from 0 to 1.
		rn := rand.Intn(2)

		// Perform a type assertion that we have a concrete type
		// of cloud in the interface value we randomly chose.
		if v, ok := mvs[rn].(cloud); ok {
			fmt.Println("Got Lucky:", v)
			continue
		}

		fmt.Println("Got Unlucky")
	}
}
