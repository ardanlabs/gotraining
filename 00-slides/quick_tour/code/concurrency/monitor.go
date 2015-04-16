// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

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
