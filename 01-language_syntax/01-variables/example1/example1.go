// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/m4PJ0FpSwX

// Sample program to show how to declare variables.
package main

import "fmt"

// main is the entry point for the application.
func main() {
	// Declare variables that are set to their zero value.
	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var b string \t %T [%v]\n", b, b)
	fmt.Printf("var c float64 \t %T [%v]\n", c, c)
	fmt.Printf("var d bool \t %T [%v]\n\n", d, d)

	// Declare variables and initalize.
	// Using the short variable declaration operator.
	aa := 10
	bb := "hello"
	cc := 3.1459
	dd := true

	fmt.Printf("aa := 10 \t %T [%v]\n", aa, aa)
	fmt.Printf("bb := \"hello\" \t %T [%v]\n", bb, bb)
	fmt.Printf("cc := 3.1459 \t %T [%v]\n", cc, cc)
	fmt.Printf("dd := true \t %T [%v]\n\n", dd, dd)

	// Specify type and perform a conversion.
	aaa := int32(10)

	fmt.Printf("aaa := int32(10) %T [%v]\n", aaa, aaa)
}

/*
	Zero Values:

	Type Initialized Value
	Boolean false
	Integer 0
	Floating Point 0
	Complex 0i
	String "" (empty string)
	Pointer nil
*/
