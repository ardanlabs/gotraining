// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to declare and use a named type.
package main

import "fmt"

// Duration is a named type that represents a duration
// of time in Nanosecond.
type duration int64

func main() {

	// Declare a variable of type Duration.
	var d duration
	fmt.Println(d)

	// Declare a variable of type int64 and assign a value.
	nanosecond := int64(10)

	// Attempted to assign a variable of type int64 (base type of duration) to
	// a variable of type duration.
	d = nanosecond

	// ./example1.go:24: cannot use nanosecond (type int64) as type duration in assignment

	// Convert a value of type int64 to type Duration.
	d = duration(nanosecond)
	fmt.Println(d)
}
