// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/lTMxc-oExx

// Sample program to show how you can't always get the address of a value.
package main

import "fmt"

// duration is a named type with a base type of int.
type duration int

// format pretty-prints the duration value.
func (d *duration) pretty() {
	fmt.Println("Duration:", *d)
}

// main is the entry point for the application.
func main() {
	duration(42).pretty()

	// ./example3.go:21: cannot call pointer method on duration(42)
	// ./example3.go:21: cannot take the address of duration(42)
}
