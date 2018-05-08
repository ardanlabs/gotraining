// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare three variables that are initialized to their zero value and three
// declared with a literal value. Declare variables of type string, int and
// bool. Display the values of those variables.
//
// Declare a new variable of type float32 and initialize the variable by
// converting the literal value of Pi (3.14).
package main

import "fmt"

func main() {

	// Declare variables that are set to their zero value.
	var age int
	var name string
	var legal bool

	// Display the value of those variables.
	fmt.Printf(
		"age[%d] name[%s] legal[%t]\n",
		age,
		name,
		legal,
	)

	// Declare variables and initialize.
	// Using the short variable declaration operator.
	month := 10
	dayOfWeek := "Tuesday"
	happy := true

	// Display the value of those variables.
	fmt.Printf(
		"month[%d] dayOfWeek[%s] happy[%t]\n",
		month,
		dayOfWeek,
		happy,
	)

	// Perform a type conversion.
	pi := float32(3.14)

	// Display the value of that variable.
	fmt.Printf("Pi type[%T] value[%v]\n", pi, pi)
}
