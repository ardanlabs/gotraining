// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This is an example of using composition and interfaces. This is
// something we want to do in Go. We will group common types by
// their behavior and not by their state. This pattern does
// provide a good design principle in a Go program.

// =============================================================================
// NOTES:

// Now I can create a list of different Animals because they
// share a common behavior that is Speak(). The interface
// provides that common type we need. The contract.

// A little copy/paste can go a long way. Treat each type as
// its own reusable independent type. Avoid creating dependency
// trees when it is reasonable and practical to do so. Each type
// can declare Name and IsMammal and should because they are
// their own type.

// Use embedding to compose for behavior not state. If you are using
// composition with state in mind question it and validate it makes
// sense for the given situation. Embed the behavior you need to
// satisfy existing interfaces.

// Go is pushing us to think about common behavior and contracts.
// We can group and work with a set of types through the common
// behavior they exhibit. Group concrete types by their behavior
// not their state.

// These facts help to flush out the code is good.
// 1) We can now create a list of different animal and work with
//    the list through the common behavior.
// 2) We have no type pollution by declaring types that are
//    never created or used directly by our code.
// 3) We have no types just providing a state based abstraction
//    layer that adds no real value to the program we are writing.

package main

import "fmt"

// Speaker provide a common behavior for all concrete types
// to follow if they want to be a part of this group. This
// is a contract for these concrete types to follow.
type Speaker interface {
	Speak()
}

// Dog contains everything an Animal is but specific
// attributes that only a Dog has.
type Dog struct {
	Name     string
	IsMammal bool
	Bark     int
}

// Speak knows how to speak like a dog.
// This makes a Dog now part of a group of concrete
// types that know how to speak.
func (d Dog) Speak() {
	fmt.Println("Woof!", d.Name, d.Bark, d.IsMammal)
}

// Cat contains everything an Animal is but specific
// attributes that only a Cat has.
type Cat struct {
	Name     string
	IsMammal bool
	Meow     int
}

// Speak knows how to speak like a cat.
// This makes a Cat now part of a group of concrete
// types that know how to speak.
func (c Cat) Speak() {
	fmt.Println("Meow!", c.Name, c.Meow, c.IsMammal)
}

func main() {

	// Create a list of Animals that know how to speak.
	speakers := []Speaker{

		// Create a Dog by initializing its Animal parts
		// and then its specific Dog attributes.
		Dog{
			Name:     "Fido",
			IsMammal: true,
			Bark:     5,
		},

		// Create a Cat by initializing its Animal parts
		// and then its specific Cat attributes.
		Cat{
			Name:     "Milo",
			IsMammal: true,
			Meow:     4,
		},
	}

	// Have the Animals speak.
	for _, spkr := range speakers {
		spkr.Speak()
	}
}
