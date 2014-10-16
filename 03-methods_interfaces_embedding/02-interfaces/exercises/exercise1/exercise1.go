// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/pbcD5WmTX9

// Declare an interface named Speaker with a method named SayHello. Declare a struct
// named English that represents a person who speaks english and declare a struct named
// Chinese for someone who speaks chinese. Implement the Speaker interface for each
// struct using a pointer receiver and these literal strings "Hello World" and "你好世界".
// Declare a variable of type Speaker and assign the _address of_ a value of type English
// and call the method. Do it again for a value of type Chinese.
//
// From exercise 1, add a new function named SayHello that accepts a value of type Speaker.
// Implement that function to call the SayHello method on the interface value. Then create
// new values of each type and use the function.
package main

import (
	"fmt"
)

// speaker implements the voice of anyone.
type speaker interface {
	sayHello()
}

// english represents an english speaking person.
type english struct{}

// sayHello implements the speaker interface.
func (e english) sayHello() {
	fmt.Println("Hello World")
}

// chinese represents a chinese speaking person.
type chinese struct{}

// sayHello implements the speaker interface.
func (c chinese) sayHello() {
	fmt.Println("你好世界")
}

// main is the entry point for the application.
func main() {
	// Declare a variable of the interfafe type.
	var sp speaker

	// Assign a value to the interface type and
	// call the interface method.
	sp = new(english)
	sp.sayHello()

	// Assign a different value to the interface type and
	// call the interface method.
	sp = new(chinese)
	sp.sayHello()

	// Create new values and call the function.
	sayHello(new(english))
	sayHello(new(chinese))
}

// SatHello abstracts speaking functionality.
func sayHello(sp speaker) {
	sp.sayHello()
}
