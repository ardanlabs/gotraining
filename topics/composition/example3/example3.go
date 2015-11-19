// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/0wWsHq-nib

// Sample program demonstrating when implicit interface conversions
// are provided by the compiler.
package main

import "fmt"

// =============================================================================

// MoveHider provides support for moving and hiding.
type MoveHider interface {
	Move()
	Hide()
}

// Mover provides support for moving things.
type Mover interface {
	Move()
}

// Locker provides support for locking things.
type Locker interface {
	Lock()
	Unlock()
}

// =============================================================================

// house represents a concrete type for the example.
type house struct{}

// Move implements the Mover interface.
func (house) Move() {
	fmt.Println("Moving the house")
}

// Hide helps to implement the MoveHider interface.
func (house) Hide() {
	fmt.Println("Hiding the house")
}

// Lock implements the Locker interface.
func (house) Lock() {
	fmt.Println("Locking the house")
}

// Unlock implements the Locker interface.
func (house) Unlock() {
	fmt.Println("Unlocking the house")
}

// =============================================================================

func main() {
	// Declare variables of the MoveHider and Mover interfaces set to their
	// zero value.
	var mh MoveHider
	var m Mover

	// Create a value of type house and assign the value to the MoveHider
	// interface value.
	mh = house{}

	// An interface value of type MoveHider can be implicitly convered into
	// a value of type Mover. They both declare a method named move.
	m = mh

	// Declare a variable of the interface type Locker set to its zero value.
	var l Locker

	// Interface type MoveHider does not declare methods named lock and unlock.
	// Therefore, the compiler can't perform an implicit conversion between
	// interface values of type MoveHider and Locker. It is irrelevant that the
	// concrete type value of type house that was assigned to the MoveHider
	// interface value implements the Locker interface.

	// prog.go:83: cannot use mh (type MoveHider) as type Locker in assignment:
	//	   MoveHider does not implement Locker (missing Lock method)
	l = mh

	// We can perform a type assertion at runtime to support the assignment.

	// Perform a type assertion against the MoveHider interface value to access
	// the concrete type value of type house that was stored inside of it. Then
	// assign the concrete type to the Locker interface.
	l = mh.(house)
}
