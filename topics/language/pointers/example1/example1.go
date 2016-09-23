// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show the basic concept of pass by value.
package main

func main() {

	// Declare variable of type int with a value of 10.
	count := 10

	// Display the "value of" and "address of" count.
	println("Before:", count, &count)

	// Pass the "value of" the variable count.
	increment(count)

	println("After: ", count, &count)
}

// increment declares count as a variable whose value is
// always an integer.
func increment(inc int) {

	// Increment the "value of" inc.
	inc++
	println("Inc:   ", inc, &inc)
}
