package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Args[0] is the program name
	fmt.Printf("Arguments: %q\n", os.Args)

	// os.Getenv returns the value part of name=value
	fmt.Printf("Parent Shell: %q\n", os.Getenv("SHELL"))

	fmt.Println("Environment:")

	// Environ returns a slice of "name=value" strings
	for _, pair := range os.Environ() {
		fmt.Printf("\t%q\n", pair)
	}
}
