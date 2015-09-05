// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/DPigLqZ5Co

// Sample program to show the basics of using flags.
package main

import (
	"flag"
	"fmt"
)

func main() {
	var cmd string

	flag.StringVar(&cmd, "cmd", cmd, `cmd can be either "hello" or "bye"`)
	flag.Parse()

	switch cmd {
	case "hello":
		fmt.Println("Hello!")
	case "bye":
		fmt.Println("Bye!")
	default:
		flag.Usage()
	}
}
