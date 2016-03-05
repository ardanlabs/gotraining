// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Escape Analysis Flaws:
// https://docs.google.com/document/d/1CxgUBPlx9iJzkz9JWkb6tIpTe5q32QDmz8l0BouG0Cw/view

// https://play.golang.org/p/KGQS9dhSmT

// Sample program to show variables stay on or escape from the stack.
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
// go build -gcflags -m

./example4.go:21: can inline stayOnStack
./example4.go:16: inlining call to stayOnStack
./example4.go:35: "Heap Addr:" escapes to heap
./example4.go:35: &x escapes to heap
./example4.go:32: moved to heap: x
./example4.go:35: &x escapes to heap
./example4.go:35: escapeToHeap ... argument does not escape
./example4.go:16: main &x does not escape
./example4.go:26: stayOnStack &x does not escape


go build -gcflags -S

"".main t=1 size=128 value=0 args=0x0 locals=0x20
	0x0000 00000 (/Users/bill/code/.../example4.go:15)	TEXT	"".main(SB), $32-0
	0x0000 00000 (/Users/bill/code/.../example4.go:15)	MOVQ	(TLS), CX
	0x0009 00009 (/Users/bill/code/.../example4.go:15)	CMPQ	SP, 16(CX)
	0x000d 00013 (/Users/bill/code/.../example4.go:15)	JLS	107
	0x000f 00015 (/Users/bill/code/.../example4.go:15)	SUBQ	$32, SP
*/
