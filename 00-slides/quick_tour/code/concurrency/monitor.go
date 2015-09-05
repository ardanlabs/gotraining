// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/GDaMvunNMZ

// Basic command line program that accepts arguments.
package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()

	// flag.Args contains all non-flag arguments
	fmt.Println(flag.Args())
}
