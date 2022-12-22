package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	version = "dev"
)

func main() {
	var showVersion bool
	flag.BoolVar(&showVersion, "version", false, "show version and exit")
	flag.Parse()

	if showVersion {
		fmt.Printf("inject version %s\n", version)
		os.Exit(0)
	}

	fmt.Println("Go!")
}
