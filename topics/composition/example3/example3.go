// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/48bYQqiV9L

// Sample program demonstrating when implicit interface conversions
// are provided by the compiler.
package main

import "fmt"

// =============================================================================

// moveHider provides support for moving and hiding.
type moveHider interface {
	move()
	hide()
}

// mover provides support for moving things.
type mover interface {
	move()
}

// locker provides support for locking things.
type locker interface {
	lock()
}

// =============================================================================

// house represents a concrete type for the example.
type house struct{}

// move implements the mover interface.
func (house) move() {
	fmt.Println("Moving the house")
}

// hide helps to implement the moveHider interface.
func (house) hide() {
	fmt.Println("Hiding the house")
}

// lock implements the locker interface.
func (house) lock() {
	fmt.Println("Locking the house")
}

// =============================================================================

func main() {
	// Declare variables of the moveHider and mover interfaces set to their
	// zero value.
	var mh moveHider
	var m mover

	// Create a value of type house and assign the value to the moveHider
	// interface value.
	mh = house{}

	// An interface value of type moveHider can be implicitly convered into
	// a value of type mover. They both declare a method named move.
	m = mh

	// Declare a variable of the interface type locker set to its zero value.
	var l locker

	// Interface type moveHider does not declare a method named lock. Therefore,
	// the compiler can't perform an implicit conversion between interface values
	// of type moveHider and locker. It is irrelevant that the concrete type
	// value of type house that was assigned to the moveHider interface value
	// implements the lock method and does satisfy the locker interface.

	// prog.go:77: cannot use mh (type moveHider) as type locker in assignment:
	//	   moveHider does not implement locker (missing lock method)
	l = mh

	// We can perform a type assertion at runtime to support the assignment.

	// Perform a type assertion against the moveHider interface value to access
	// the concrete type value of type house that was stored inside of it. Then
	// assign the concrete type to the locker interface.
	l = mh.(house)
}
