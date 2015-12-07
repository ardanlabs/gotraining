// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// http://play.golang.org/p/2kfVP_SGA4

// Sample program demonstrating that type assertions are a runtime and
// not compile time construct.
package main

import (
	"fmt"
	"math/rand"
)

// car represents something that can move.
type car struct{}

// Move implements the Mover interface.
func (car) String() string {
	return "Moving Car"
}

// cloud represents something that can move.
type cloud struct{}

// Move implements the Mover interface.
func (cloud) String() string {
	return "Moving Cloud"
}

// =============================================================================

// main is the entry point for the application.
func main() {
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
