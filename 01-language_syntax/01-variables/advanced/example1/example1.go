// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/a0gQyj6TBb

// Sample program to show varaibles stay on or escape from the stack.
package main

import "fmt"

// main is the entry point for the application.
func main() {
	stayOnStack()
	escapeToHeap()
}

// stayOnStack shows how the variable does not escape.
func stayOnStack() {
	// Declare a variable of type integer.
	var x int

	// Display the address of the variable.
	println("Stack Addr:", &x)
}

// escapeToHeap shows how the variable does escape.
func escapeToHeap() {
	// Declare a variable of type integer.
	var x int

	// Display the address of the variable.
	fmt.Println("Heap Addr:", &x)
}
