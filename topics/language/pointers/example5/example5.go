// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how stacks grow/change.
package main

// Number of elements to grow each stack frame.
// Run with 10 and then with 1024
const size = 10

// main is the entry point for the application.
func main() {
	s := "HELLO"
	stackCopy(&s, 0, [size]int{})
}

// stackCopy recursively runs increasing the size
// of the stack.
func stackCopy(s *string, c int, a [size]int) {
	println(c, s, *s)

	c++
	if c == 10 {
		return
	}

	stackCopy(s, c, a)
}
