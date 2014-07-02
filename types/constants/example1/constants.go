package main

import "fmt"

func main() {
	// Constant live within the compiler. This is why we can have a paralell type system.
	// Compiler can perform implicit conversions of constants.

	// Untyped Constants with Implicit Conversions.
	const ui = 12345    // kind: integer
	const uf = 3.141592 // kind: floating-point

	fmt.Printf("const ui = 12345 \t %T [%v]\n", ui, ui)
	fmt.Printf("const uf = 3.141592 \t %T [%v]\n", uf, uf)

	// Type Constants.
	const ti int = 12345        // kind: integer
	const tf float64 = 3.141592 // kind: floating-point

	fmt.Printf("const ti int = 12345 \t %T [%v]\n", ti, ti)
	fmt.Printf("const tf float64 = 3.141592 \t %T [%v]\n", tf, tf)

	// ./constants.go:14: constant 1000 overflows uint8
	// const myUint8 uint8 = 1000

	// Kind Promotion
	var answer = 3 * 0.333

	fmt.Printf("var answer = 3 * 0.333 \t %T [%v]\n", answer, answer)

	const third = 1 / 3.0
	const zero = 1 / 3

	fmt.Printf("const third = 1 / 3.0 \t %T [%v]\n", third, third)
	fmt.Printf("const zero = 1 / 3 \t %T [%v]\n", zero, zero)

	const one int8 = 1
	const two = 2 * one

	fmt.Printf("const two = 2 * one \t %T [%v]\n", two, two)
}
