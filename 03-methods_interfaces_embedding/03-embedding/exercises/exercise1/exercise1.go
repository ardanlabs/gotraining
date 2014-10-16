// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/YtdNsTwAN7

// Declare a struct type named animal with two fields name and age. Declare a struct
// type named dog with the field bark. Embed the animal type into the dog type. Declare
// and initalize a value of type dog. Display the value of the variable.
//
// Declare a method named yelp to the animal type using a pointer reciever which displays the
// literal string "Not Implemented". Call the method from the value of type dog.
//
// Declare an interface named speaker with a single method called yelp. Declare a value of
// type speaker and assign the address of the value of type dog. Call the method yelp.
//
// Implement the speaker interface for the dog type. Be creative with the
// bark field. Call the method yelp again from the value of type speaker.
package main

import (
	"fmt"
)

// speaker represents talking animals.
type speaker interface {
	yelp()
}

// animal represents all animals.
type animal struct {
	name string
	age  int
}

// yelp represents how an animal can speak.
func (a *animal) yelp() {
	fmt.Println("Not Implemented")
}

// dog represents a dog.
type dog struct {
	animal
	bark int
}

// yelp represents how an animal can speak.
func (d *dog) yelp() {
	for bark := 0; bark < d.bark; bark++ {
		fmt.Print("BARK ")
	}
	fmt.Println()
}

// main is the entry point for the application.
func main() {
	// Create a value of type dog.
	d := dog{
		animal: animal{
			name: "Chole",
			age:  1,
		},
		bark: 10,
	}

	// Display the value.
	fmt.Println(d)

	// Use the interface.
	var s speaker
	s = &d
	s.yelp()
}
