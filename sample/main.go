package main

import (
	"log"
	"os"

	_ "github.com/ArdanStudios/gotraining/sample/matchers"
	"github.com/ArdanStudios/gotraining/sample/search"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main() {
	log.Println("V1")
	// Perform the search for the specified term.
	search.Run("president")
}
