// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// Sample program to show how to write tests for a practical
// program that makes database requests to MongoDB.
package main

import (
	"log"
	"os"

	"github.com/ArdanStudios/gotraining/06-testing/example1/buoy"
)

// init is called before main.
func init() {
	// Change the output device from the default
	// stderr to stdout.
	log.SetOutput(os.Stdout)

	// Set the prefix string for each log line.
	log.SetPrefix("TRACE: ")

	// Set the flag to show the code file and date/time
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// main is the entry point for the application.
func main() {
	// Retrieve a document for station 42002.
	stationID := "42002"
	station, err := buoy.FindStation(stationID)
	if err != nil {
		log.Printf("main : ERROR : %s\n", err)
		os.Exit(1)
	}

	buoy.Print(station)

	// Retrieve a slice of documents for the
	// Gulf Of Mexico region.
	region := "Gulf Of Mexico"
	stations, err := buoy.FindRegion(region, 5)
	if err != nil {
		log.Printf("main : ERROR : %s\n", err)
		os.Exit(1)
	}

	for _, station := range stations {
		buoy.Print(&station)
	}
}
