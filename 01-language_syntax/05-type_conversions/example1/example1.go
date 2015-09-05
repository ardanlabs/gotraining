// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// http://play.golang.org/p/mhKlxSyuxr

// Sample program to show how to declare and use a named type.
package main

import "fmt"

// Duration is a named type that represents a duration
// of time in Nanosecond.
type Duration int64

func main() {
	// Declare a variable of type Duration.
	var duration Duration
	fmt.Println(duration)

	// Declare a variable of type int64 and assign a value.
	nanosecond := int64(10)

	// Attempted to assign a variable of type int64 (base type of Duration) to
	// a variable of type Duration.
	// duration = nanosecond

	// ./example1.go:22: cannot use nanosecond (type int64) as type Duration in assignment

	// Convert a value of type int64 to type Duration.
	duration = Duration(nanosecond)
	fmt.Println(duration)
}
