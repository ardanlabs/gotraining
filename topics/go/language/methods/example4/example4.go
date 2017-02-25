// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to declare and use function types.
package main

import "fmt"

// event displays global events.
func event(message string) {
	fmt.Println(message)
}

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

// fireEvent1 uses an anonymous function type.
func fireEvent1(f func(string)) {
	f("anonymous")
}

// handler represents a function for handling events.
type handler func(string)

// fireEvent2 uses a function type.
func fireEvent2(h handler) {
	h("handler")
}

// =============================================================================

func main() {

	// Declare a variable of type data.
	d := data{
		name: "Bill",
	}

	// Use the fireEvent1 handler that accepts any
	// function or method with the right signature.
	fireEvent1(event)
	fireEvent1(d.event)

	// Use the fireEvent2 handler that accepts any
	// function or method of type `handler` or any
	// literal function or method with the right signature.
	fireEvent2(event)
	fireEvent2(d.event)

	// Declare a variable of type handler for the
	// global and method based event functions.
	h1 := handler(event)
	h2 := handler(d.event)

	// User the fireEvent2 handler that accepts
	// values of type handler.
	fireEvent2(h1)
	fireEvent2(h2)

	// User the fireEvent1 handler that accepts
	// any function or method with the right signature.
	fireEvent1(h1)
	fireEvent1(h2)
}
