// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// https://docs.google.com/document/d/1CxgUBPlx9iJzkz9JWkb6tIpTe5q32QDmz8l0BouG0Cw/view

// http://play.golang.org/p/931Cw6uzsn

// go build -gcflags -m

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

/*
./example1.go:20: can inline stayOnStack
./example1.go:15: inlining call to stayOnStack
./example1.go:31: moved to heap: x
./example1.go:34: &x escapes to heap
./example1.go:34: escapeToHeap ... argument does not escape
./example1.go:15: stayOnStack &x does not escape
./example1.go:25: stayOnStack &x does not escape
*/
