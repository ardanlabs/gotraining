// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program that explores how interface assignments work when
// values are stored inside the interface.
package main

import (
	"fmt"
	"unsafe"
)

// biller is a random interface.
type biller interface {
	bill()
}

// value is a concrete type that will implement the interface.
type value int

// bill implements the biller interface.
func (v value) bill() {
	fmt.Println(v)
}

func main() {

	// Create a biller interface and concrete type value.
	var b biller
	v := value(10)

	// The assignment stores a copy of the value inside
	// the interface value.
	b = v

	// Get a pointer to the second word of the interface value.
	// We want to inspect the value that the interface is pointing to.
	second := uintptr(unsafe.Pointer(&b)) + uintptr(4)
	valPtr := (**value)(unsafe.Pointer(second))
	fmt.Println("Addr V:", &v, "Int Addr:", *valPtr, "Value:", **valPtr)

	// Create a second interface value and assign the orginal interface
	// value to the new interface value.
	var b1 biller
	b1 = b

	// Get a pointer to the second word of the interface value.
	// We want to inspect the value that the interface is pointing to.
	second = uintptr(unsafe.Pointer(&b1)) + uintptr(4)
	valPtr = (**value)(unsafe.Pointer(second))
	fmt.Println("Addr V:", &v, "Int Addr:", *valPtr, "Value:", **valPtr)

	// Change the value that the interface is pointing to.
	**valPtr = 11

	// We see the change in both interface values. What this means is
	// that when we make an assignment between different interface values,
	// like on line 42, we are copying just the interface value. We are not making
	// a second copy of the value that was stored. Both b and b1 are pointing to
	// the same value we originally stored on line 26.
	fmt.Println(b.(value), b1.(value))

	// The type assertion is returning a copy of the value stored. This is
	// whether the value was an address or not.
}
