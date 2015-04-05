// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// Sample program to show how the program can't access an
// unexported identifier from another package.
package main

import (
	"fmt"

	"github.com/ArdanStudios/gotraining/04-packaging_exporting/example2/counters"
)

// main is the entry point for the application.
func main() {
	// Create a variable of the unexported type and initialize the value to 10.
	counter := counters.alertCounter(10)

	// ./example2.go:17: cannot refer to unexported name counters.alertCounter
	// ./example2.go:17: undefined: counters.alertCounter

	fmt.Printf("Counter: %d\n", counter)
}
