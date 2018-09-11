// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to use if statements.
package main

import "fmt"

func main() {

	// num is the value we are inspecting. Change this and see what happens.
	num := 42

	fmt.Printf("Input is %d.\n", num)

	// Familiar C-like patterns are supported.
	if num < 50 {
		fmt.Println("That is less than 50.")
	} else if num == 50 {
		fmt.Println("That is exactly 50.")
	} else {
		fmt.Println("That is greater 50.")
	}

	// Though Gophers tend to avoid using else.
	if num < 100 {
		fmt.Println("That is also less than 100.")
	}

	// An optional expression may be ran before the comparison. Variables
	// declared in that expression are scoped to the if block.
	if dbl := num * 2; dbl < 100 {
		fmt.Println("Double that number is also less than 100.")
	}

	// Will not compile.
	// example1.go:38:27: undefined: dbl
	//fmt.Println("double is", dbl)

	// Also will not compile. An explicit boolean expression is required.
	// example1.go:42:2: non-bool num (type int) used as if condition
	//if num {
	//fmt.Println("Num is true?")
	//}
}
