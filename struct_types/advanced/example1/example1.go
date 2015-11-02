// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// http://play.golang.org/p/IiElaanvbY

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
	flag2   bool
	pi      float32
}

// main is the entry point for the application.
func main() {
	var e example

	alignmentBoundary := unsafe.Alignof(e)
	sizeE := unsafe.Sizeof(e)

	sizeBool := unsafe.Sizeof(e.flag)
	offsetBool := unsafe.Offsetof(e.flag)

	sizeInt := unsafe.Sizeof(e.counter)
	offsetInt := unsafe.Offsetof(e.counter)

	sizeBool2 := unsafe.Sizeof(e.flag2)
	offsetBool2 := unsafe.Offsetof(e.flag2)

	sizeFloat := unsafe.Sizeof(e.pi)
	offsetFloat := unsafe.Offsetof(e.pi)

	fmt.Printf("Alignment(%d):\tSize: %d\n", alignmentBoundary, sizeE)

	fmt.Printf("flag:\t\tSize: %d\t\tOffset: %d\tAddr: %v\n",
		sizeBool, offsetBool, &e.flag)

	fmt.Printf("counter:\tSize: %d\t\tOffset: %d\tAddr: %v\n",
		sizeInt, offsetInt, &e.counter)

	fmt.Printf("flag1:\t\tSize: %d\t\tOffset: %d\tAddr: %v\n",
		sizeBool2, offsetBool2, &e.flag2)

	fmt.Printf("pi\t\tSize: %d\t\tOffset: %d\tAddr: %v\n",
		sizeFloat, offsetFloat, &e.pi)
}

/*
Alignment(4):	Size: 12
flag:		Size: 1		Offset: 0	Addr: 0x104382e0
counter:	Size: 2		Offset: 2	Addr: 0x104382e2
flag1:		Size: 1		Offset: 4	Addr: 0x104382e4
pi		Size: 4		Offset: 8	Addr: 0x104382e8
*/
