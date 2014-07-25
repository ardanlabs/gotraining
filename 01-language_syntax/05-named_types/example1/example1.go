// http://play.golang.org/p/rHVQ0tGgiT

// Sample program to show how to declare and use a named type.
package main

import "fmt"

// Duration is a named type that represents a duration
// of time in Nanosecond.
type Duration int64

func main() {
	// Declare a variable of type Duration
	var duration Duration
	fmt.Println(duration)

	// Declare a variable of type int and assign a value.
	nanosecond := 10

	// Attemped to assign a variable of type int (base type of Duration) to
	// a variable of type Duration.
	// duration = nanosecond

	// ./example1.go:20: cannot use value (type int) as type Duration in assignment

	// Convert a value of type int to type Duration.
	duration = Duration(nanosecond)
	fmt.Println(duration)
}
