// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://docs.google.com/document/d/1CxgUBPlx9iJzkz9JWkb6tIpTe5q32QDmz8l0BouG0Cw/view

// http://play.golang.org/p/_uK8EYlsd0

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
./example1.go:22: can inline stayOnStack
./example1.go:17: inlining call to stayOnStack
./example1.go:36: "Heap Addr:" escapes to heap
./example1.go:36: &x escapes to heap
./example1.go:33: moved to heap: x
./example1.go:36: &x escapes to heap
./example1.go:36: escapeToHeap ... argument does not escape
./example1.go:17: main &x does not escape
./example1.go:27: stayOnStack &x does not escape
*/
