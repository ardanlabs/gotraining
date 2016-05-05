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

func inspect(n *notifier, u *user) {
	word := uintptr(unsafe.Pointer(n)) + uintptr(unsafe.Sizeof(&u))
	value := (**user)(unsafe.Pointer(word))
	fmt.Printf("Addr User: %p  Word Value: %p  Ptr Value: %v\n", u, *value, **value)
}

func main() {

	// Create a notifier interface and concrete type value.
	var n1 notifier
	u := user{"bill"}

	// Store a copy of the user value inside the notifier
	// interface value.
	n1 = u

	// We see the interface has its own copy.
	// Addr User: 0x1040a120  Word Value: 0x10427f70  Ptr Value: {bill}
	inspect(&n1, &u)

	// Make a copy of the interface value.
	n2 := n1

	// We see the interface is sharing the same value stored in
	// the n1 interface value.
	// Addr User: 0x1040a120  Word Value: 0x10427f70  Ptr Value: {bill}
	inspect(&n2, &u)

	// Store a copy of the user address value inside the
	// notifier interface value.
	n1 = &u

	// We see the interface is sharing the u variables value
	// directly. There is no copy.
	// Addr User: 0x1040a120  Word Value: 0x1040a120  Ptr Value: {bill}
	inspect(&n1, &u)
}
