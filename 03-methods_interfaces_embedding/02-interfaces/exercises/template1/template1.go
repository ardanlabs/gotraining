// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/oijJdRW3cD

// Declare an interface named speaker with a method named sayHello. Declare a struct
// named English that represents a person who speaks english and declare a struct named
// Chinese for someone who speaks chinese. Implement the speaker interface for each
// struct using a value receiver and these literal strings "Hello World" and "你好世界".
// Declare a variable of type speaker and assign the _address of_ a value of type English
// and call the method. Do it again for a value of type Chinese.
//
// Add a new function named sayHello that accepts a value of type speaker.
// Implement that function to call the sayHello method on the interface value. Then create
// new values of each type and use the function.
package main

// Add imports.

// Declare the speaker interface with a single method called sayHello.

// Declare an empty struct type named english.

// Declare a method named sayHello for the english type
// using a value receiver. "Hello World"

// Declare an empty struct type named chinese.

// Declare a method named sayHello for the chinese type
// using a value receiver. "你好世界"

// sayHello accepts values of the interface type.
func sayHello( /* Declare parameter */ ) {
	// Call the sayHello() method from the interface parameter.
}

// main is the entry point for the application.
func main() {
	// Declare a variable of the interface type set to its zero value.

	// Declare a variable of type english and assign it to
	// the interface variable.
	// Call the sayHello() method from the interface variable.

	// Declare a variable of type chinese and assign it to
	// the interface variable.
	// Call the sayHello() method from the interface variable.

	// Call the sayHello function passing each concrete type.
}
