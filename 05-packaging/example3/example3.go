// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// Sample program to show how the program can access a value
// of an unexported identifier from another package.
package main

import (
	"fmt"

	"github.com/ArdanStudios/gotraining/05-packaging/example3/counters"
)

// main is the entry point for the application.
func main() {
	// Create a variable of the unexported type using the
	// exported New function from the package counters.
	counter := counters.New(10)

	fmt.Printf("Counter: %d\n", counter)
}
