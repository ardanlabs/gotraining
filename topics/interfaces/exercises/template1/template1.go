// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/8XNVqdVYV2

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

// Add imports.

// Declare the speaker interface with a single method called speak.

// Declare an empty struct type named english.

// Declare a method named speak for the english type
// using a value receiver. "Hello World"

// Declare an empty struct type named chinese.

// Declare a method named speak for the chinese type
// using a value receiver. "你好世界"

// sayHello accepts values of the speaker type.
func sayHello( /* Declare parameter */ ) {
	// Call the speak method from the speaker parameter.
}

// main is the entry point for the application.
func main() {
	// Declare a variable of the speaker type set to its zero value.

	// Declare a variable of type english and assign it to
	// the speaker variable.
	// Call the speak method from the speaker variable.

	// Declare a variable of type chinese and assign it to
	// the speaker variable.
	// Call the speak method from the speaker variable.

	// Call the sayHello function with new values of each concrete type.
}
