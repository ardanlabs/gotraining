// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/ZuB82kgz2K

// Alignment is about placing types on boundaries that make the
// CPU access the fastest.

// Sample program to show how struct types align on boundaries.
package main

import (
	"fmt"
	"unsafe"
)

// example represents a type with different fields.
type example struct {
	flag    bool
	counter int16
	pi      float32
}

// main is the entry point for the application.
func main() {
	// Declare variable of type example and init using
	// a composite literal.
	e := example{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	// Declare variable of type example and init using
	// a composite literal.
	eNext := example{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	// By placing both value one after the other on the stack
	// we can see the actual padding and where it is located.

	alignmentBoundary := unsafe.Alignof(e)

	sizeBool := unsafe.Sizeof(e.flag)
	offsetBool := unsafe.Offsetof(e.flag)

	sizeInt := unsafe.Sizeof(e.counter)
	offsetInt := unsafe.Offsetof(e.counter)

	sizeFloat := unsafe.Sizeof(e.pi)
	offsetFloat := unsafe.Offsetof(e.pi)

	sizeBoolNext := unsafe.Sizeof(eNext.flag)
	offsetBoolNext := unsafe.Offsetof(eNext.flag)

	fmt.Printf("Alignment Boundary: %d\n", alignmentBoundary)

	fmt.Printf("flag = Size: %d Offset: %d Addr: %v\n",
		sizeBool, offsetBool, &e.flag)

	fmt.Printf("counter = Size: %d Offset: %d Addr: %v\n",
		sizeInt, offsetInt, &e.counter)

	fmt.Printf("pi = Size: %d Offset: %d Addr: %v\n",
		sizeFloat, offsetFloat, &e.pi)

	fmt.Printf("Next = Size: %d Offset: %d Addr: %v\n",
		sizeBoolNext, offsetBoolNext, &eNext.flag)
}

/*
Alignment Boundary: 4
flag = Size: 1 Offset: 0 Addr: 0x20817a170
counter = Size: 2 Offset: 2 Addr: 0x20817a172
pi = Size: 4 Offset: 4 Addr: 0x20817a174
Next = Size: 1 Offset: 0 Addr: 0x20817a178
*/
