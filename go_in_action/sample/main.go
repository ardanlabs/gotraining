// This sample program demonstrates how to write idiomatic Go. The sample
// uses many features from the language to provide a well rounded
// view Go's capabilities and how to apply them.
package main

import (
	"log"
	"os"

	_ "github.com/ArdanStudios/gotraining/go_in_action/sample/matchers"
	"github.com/ArdanStudios/gotraining/go_in_action/sample/search"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main() {
	// Perform the search for the specified term.
	search.Run("president")
}
