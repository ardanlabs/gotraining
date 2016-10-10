// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This is an example of using type hierarchies with a OOP pattern.
// This is not something we want to do in Go. Go does not have the
// concept of sub-typing. All types are their own and the concepts of
// base and derived types do not exist in Go. This pattern does not
// provide a good design principle in a Go program.
package main

import "fmt"

// Animal contains all the base fields for animals.
type Animal struct {
	Name     string
	IsMammal bool
}

// Speak provides generic behavior for all animals and
// how they speak.
// SMELL - This can't apply to all animals.
func (a Animal) Speak() {
	fmt.Println("UGH!",
		"My name is", a.Name,
		", it is", a.IsMammal,
		"I am a mammal")
}

// Dog contains everything an Animal is but specific
// attributes that only a Dog has.
type Dog struct {
	Animal
	PackFactor int
}

// Speak knows how to speak like a dog.
func (d Dog) Speak() {
	fmt.Println("Woof!",
		"My name is", d.Name,
		", it is", d.IsMammal,
		"I am a mammal with a pack factor of", d.PackFactor)
}

// Cat contains everything an Animal is but specific
// attributes that only a Cat has.
type Cat struct {
	Animal
	ClimbFactor int
}

// Speak knows how to speak like a cat.
func (c Cat) Speak() {
	fmt.Println("Meow!",
		"My name is", c.Name,
		", it is", c.IsMammal,
		"I am a mammal with a climb factor of", c.ClimbFactor)
}

func main() {

	// SMELL - I can't create a list of Cats and Dogs using
	// the Animal type. Can't create a list based on a
	// common set of state.

	// DOES NOT COMPILE

	// Create a list of Animals that know how to speak.
	animals := []Animal{

		// Create a Dog by initializing its Animal parts
		// and then its specific Dog attributes.
		Dog{
			Animal: Animal{
				Name:     "Fido",
				IsMammal: true,
			},
			PackFactor: 5,
		},

		// Create a Cat by initializing its Animal parts
		// and then its specific Cat attributes.
		Cat{
			Animal: Animal{
				Name:     "Milo",
				IsMammal: true,
			},
			ClimbFactor: 4,
		},
	}

	// Have the Animals speak.
	for _, animal := range animals {
		animal.Speak()
	}
}

// =============================================================================

// NOTES:

// Smells:
// 	* The Animal type is providing an abstraction layer of reusable state.
// 	* The program never needs to create or solely use a value of type Animal.
// 	* The implementation of the Speak method for the Animal type is a generalization.
// 	* The Speak method for the Animal type is never going to be called.

// Here are some guidelines around declaring types:
// 	* Declare types that represent something new or unique.
// 	* Validate that a value of any type is created or used on its own.
// 	* Embed types to reuse existing behaviors you need to satisfy.
// 	* Question types that are an alias or abstraction for an existing type.
// 	* Question types whose sole purpose is to share common state.
