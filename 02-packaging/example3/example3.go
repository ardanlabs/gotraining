// Sample program to show how the program can access a value
// of an unexported identifier from another package.
package main

import (
	"fmt"

	"github.com/ArdanStudios/gotraining/02-packaging/example3/counters"
)

// main is the entry point for the application.
func main() {
	// Create a variable of the unexported type using the
	// exported NewAlertCounter function from the package counters.
	counter := counters.NewAlertCounter(10)

	fmt.Printf("Counter: %d\n", counter)
}
