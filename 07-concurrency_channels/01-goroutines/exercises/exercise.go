// Program uses a function type and closures to create a recursive call
// that maintains its own state to calculate Fibonacci numbers.
package main

import "fmt"

// fib is a function type that takes no parameters
// and returns an integer.
type fib func() int

// main is the entry point for all Go programs.
func main() {
	calculate(5, fibonacci())
}

// fib returns a function that returns successive
// Fibonacci numbers. This function uses closure
// to maintain the state for the next call.
func fibonacci() fib {
	first, second := 0, 1
	return func() int {
		first, second = second, first+second
		return first
	}
}

// calculate recurses the call to the fibonacci function.
func calculate(toIndex int, f fib) {
	if toIndex > 0 {
		fmt.Println(f())
		calculate(toIndex-1, f)
	}
}
