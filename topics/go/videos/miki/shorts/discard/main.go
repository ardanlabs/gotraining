package main

import (
	"flag"
	"io"
	"log"
)

func main() {
	var config struct {
		Verbose bool
	}
	flag.BoolVar(&config.Verbose, "verbose", false, "be more verbose")
	flag.Parse()

	if !config.Verbose {
		log.SetOutput(io.Discard)
	}
	log.Printf("Please reinstall universe and reboot")
}
