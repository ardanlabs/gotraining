// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// This sample program demonstrates how to write idiomatic Go. The sample
// uses many features from the language to provide a well rounded
// view Go's capabilities and how to apply them.
package main

import (
	"log"
	"os"

	_ "github.com/ArdanStudios/gotraining/feed_app/sample/matchers"
	"github.com/ArdanStudios/gotraining/feed_app/sample/search"
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
