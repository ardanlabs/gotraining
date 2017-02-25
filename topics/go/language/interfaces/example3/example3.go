// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how you can't always get the address of a value.
package main

import "fmt"

// duration is a named type with a base type of int.
type duration int

// notify implements the notifier interface.
func (d *duration) notify() {
	fmt.Println("Sending Notification in", *d)
}

func main() {
	duration(42).notify()

	// ./example3.go:18: cannot call pointer method on duration(42)
	// ./example3.go:18: cannot take the address of duration(42)
}
