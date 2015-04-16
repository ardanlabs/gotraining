// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

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
