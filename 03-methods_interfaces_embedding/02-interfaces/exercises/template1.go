// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/sgCkOodyYf

// Declare an interface named speaker with a method named sayHello. Declare a struct
// named English that represents a person who speaks english and declare a struct named
// Chinese for someone who speaks chinese. Implement the speaker interface for each
// struct using a pointer receiver and these literal strings "Hello World" and "你好世界".
// Declare a variable of type speaker and assign the _address of_ a value of type English
// and call the method. Do it again for a value of type Chinese.
//
// From exercise 1, add a new function named sayHello that accepts a value of type speaker.
// Implement that function to call the sayHello method on the interface value. Then create
// new values of each type and use the function.
package main

// speaker implements the voice of anyone.
/*
	type interface_type_name interface {
		method_name()
	}
*/

// english represents an english speaking person.
/*
	type english_type_name struct{}
*/

// sayHello implements the speaker interface.
/*
	func (receiver_name english_type_name) method_name() {
		fmt.Println("Hello World")
	}
*/

// chinese represents a chinese speaking person.
/*
	type chinese_type_name struct{}
*/

// sayHello implements the speaker interface.
/*
	func (receiver_name chinese_type_name) method_name() {
		fmt.Println("你好世界")
	}
*/

// main is the entry point for the application.
func main() {
	// Declare a variable of the interfafe type.
	/*
		var variable_name interface_type_name
	*/

	// Assign a value to the interface type and
	// call the interface method.
	/*
		var variable_name english_type_name
		sp = variable_name
		sp.method_name()
	*/

	// Assign a different value to the interface type and
	// call the interface method.
	/*
		var variable_name chinese_type_name
		sp = variable_name
		sp.method_name()
	*/

	// Create new values and call the function.
	/*
		function_name(new(english_type_name))
		function_name(new(chinese_type_name))
	*/
}

// SatHello abstracts speaking functionality.
/*
	func function_name(variable_name interface_type_name) {
		variable_name.method_name()
	}
*/
