// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// https://play.golang.org/p/yaUWLZjidB

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
