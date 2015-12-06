// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/PS9sqY6kSp

// Declare an interface named speaker with a method named speak. Declare a struct
// named english that represents a person who speaks english and declare a struct named
// chinese for someone who speaks chinese. Implement the speaker interface for each
// struct using a value receiver and these literal strings "Hello World" and "你好世界".
// Declare a variable of type speaker and assign the address of a value of type english
// and call the method. Do it again for a value of type chinese.
//
// Add a new function named sayHello that accepts a value of type speaker.
// Implement that function to call the speak method on the interface value. Then create
// new values of each type and use the function.
package main

import "fmt"

// speaker implements the voice of anyone.
type speaker interface {
	speak()
}

// english represents an english speaking person.
type english struct{}

// speak implements the speaker interface.
func (english) speak() {
	fmt.Println("Hello World")
}

// chinese represents a chinese speaking person.
type chinese struct{}

// speak implements the speaker interface.
func (chinese) speak() {
	fmt.Println("你好世界")
}

// main is the entry point for the application.
func main() {
	// Declare a variable of the interface type.
	var sp speaker

	// Assign a value to the interface type variable and
	// call the interface method.
	var e english
	sp = e
	sp.speak()

	// Assign a different value to the interface type
	// variable and call the interface method.
	var c chinese
	sp = c
	sp.speak()

	// Create new values and call the function.
	sayHello(new(english))
	sayHello(&chinese{})

	// The use of new() and the empty literal is there
	// as a talking point about these options.
}

// sayHello abstracts speaking functionality.
func sayHello(sp speaker) {
	sp.speak()
}
