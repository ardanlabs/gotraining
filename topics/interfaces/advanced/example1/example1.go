// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program that explores how interface assignments work when
// values are stored inside the interface.
package main

import (
	"fmt"
	"unsafe"
)

// notifier provides support for notifying events.
type notifier interface {
	notify()
}

// user represents a user in the system.
type user struct {
	name string
}

// notify implements the notifier interface.
func (u user) notify() {
	fmt.Println("Alert", u.name)
}

func main() {

	// Capture the size of a word in this arch.
	size := unsafe.Sizeof(10)

	// Create a notifier interface and concrete type value.
	var n1 notifier
	u := user{"bill"}

	// The assignment stores a copy of the user value inside
	// the notifier interface value.
	n1 = u

	// Get a pointer to the second word of the interface value.
	// We want to inspect the value that the interface is pointing to.
	word := uintptr(unsafe.Pointer(&n1)) + uintptr(size)
	value := (**user)(unsafe.Pointer(word))
	fmt.Printf("N1: Addr User: %p  Word Value: %p  Ptr Value: %v\n", &u, *value, **value)

	// Create a second interface value and assign the orginal interface
	// value to the new interface value.
	var n2 notifier
	n2 = n1

	// Get a pointer to the second word of the interface value.
	// We want to inspect the value that the interface is pointing to.
	word = uintptr(unsafe.Pointer(&n2)) + uintptr(size)
	value = (**user)(unsafe.Pointer(word))
	fmt.Printf("N2: Addr User: %p  Word Value: %p  Ptr Value: %v\n", &u, *value, **value)

	// Mutate the value that the interface is pointing to.
	(**value).name = "lisa"

	// We see the change in both interface values. What this means is
	// that when we make an assignment between different interface values,
	// like on line 42, we are copying just the interface value. We are not making
	// a second copy of the value that was stored. Both b and b1 are pointing to
	// the same value we originally stored on line 26.
	fmt.Println("N1:", n1.(user), "N2:", n2.(user))

	// The type assertion is returning a copy of the value stored. This is
	// whether the value was an address or not.
}
