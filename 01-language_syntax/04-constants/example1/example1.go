// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/OLuzwK1oHT

// Sample program to show how to declare constants and their
// implementation in Go.
package main

import "fmt"

// main is the entry point for the application.
func main() {
	// Constants live within the compiler.
	// They have a paralell type system.
	// Compiler can perform implicit conversions of untyped constants.

	// Untyped Constants.
	const ui = 12345    // kind: integer
	const uf = 3.141592 // kind: floating-point

	// Constants in this context requires a converstion to a built-in type.
	fmt.Printf("const ui = 12345 \t %T [%v]\n", ui, ui)
	fmt.Printf("const uf = 3.141592 \t %T [%v]\n", uf, uf)

	// Typed Constants still use the constant type system but their precision
	// is restricted.
	const ti int = 12345        // type: int64
	const tf float64 = 3.141592 // type: float64

	// ./constants.go:14: constant 1000 overflows uint8
	// const myUint8 uint8 = 1000

	// Constants in this context requires a converstion to a built-in type.
	fmt.Printf("const ti int = 12345 \t %T [%v]\n", ti, ti)
	fmt.Printf("const tf float64 = 3.141592 \t %T [%v]\n", tf, tf)

	// Constant arithmetic supports different kinds.
	// Kind Promotion is used to determine kind in these scenarios.
	// Variable answer will be implicitly converted to type floating point.
	var answer = 3 * 0.333 // KindInt(3) / KindFloat(3.0)

	fmt.Printf("var answer = 3 * 0.333 \t %T [%v]\n", answer, answer)

	// Variable third will be of kind floating point.
	const third = 1 / 3.0 // KindInt(1) / KindFloat(3.0)

	// Variable zero will be of kind integer.
	const zero = 1 / 3 // KindInt(1) / KindInt(3)

	// Constants in this context requires a converstion to a built-in type.
	fmt.Printf("const third = 1 / 3.0 \t %T [%v]\n", third, third)
	fmt.Printf("const zero = 1 / 3 \t %T [%v]\n", zero, zero)

	// This is an example of constant arithmetic between typed and
	// untyped constants. Must have like types to perform math.
	const one int8 = 1
	const two = 2 * one // int8(2) * int8(1)

	// Constants in this context requires a converstion to a built-in type.
	fmt.Printf("const two = 2 * one \t %T [%v]\n", two, two)
}
