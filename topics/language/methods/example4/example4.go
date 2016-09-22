// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to declare and use function types.
package main

import "fmt"

// handler represents a function for handling events.
type handler func(string)

// =============================================================================

// data is a struct to bind methods to.
type data struct {
	name string
	age  int
}

// event displays events for this data.
func (d *data) event(message string) {
	fmt.Println(d.name, message)
}

// =============================================================================

// event displays global events.
func event(message string) {
	fmt.Println(message)
}

// =============================================================================

func main() {

	// Declare a variable of type data.
	d := data{
		name: "Bill",
	}

	// Use the fireEvent handler that accepts any
	// function or method with the right signature.
	fireEvent1(d.event)

	// Declare a variable of type handler for the global event function.
	h := handler(event)

	// User the fireEvent handler that accepts
	// values of type handler.
	fireEvent2(h)
}

// fireEvent1 uses an anonymous function type.
func fireEvent1(f func(string)) {
	f("message 1")
}

// fireEvent2 uses a function type.
func fireEvent2(h handler) {
	h("message 2")
}
