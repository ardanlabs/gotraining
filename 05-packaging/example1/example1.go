// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// Sample program to show how to access an exported identifier.
package main

import (
	"fmt"

	"github.com/ArdanStudios/gotraining/05-packaging/example1/counters"
)

// main is the entry point for the application.
func main() {
	// Create a variable of the exported type and
	// initialize the value to 10.
	counter := counters.AlertCounter(10)

	fmt.Printf("Counter: %d\n", counter)
}
