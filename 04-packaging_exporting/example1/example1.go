// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to access an exported identifier.
package main

import (
	"fmt"

	"github.com/ardanlabs/gotraining/04-packaging_exporting/example1/counters"
)

// main is the entry point for the application.
func main() {
	// Create a variable of the exported type and initialize the value to 10.
	counter := counters.AlertCounter(10)

	fmt.Printf("Counter: %d\n", counter)
}
