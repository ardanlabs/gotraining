// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how iota works.
package main

import "fmt"

func main() {

	// Constant block with iota using increment.
	const (
		A int         = iota //  0 : Start at 0
		B                    //  1 : Increment by 1
		C                    //  2 : Increment by 1
		D = 10               // 10 : Set to 10
		E                    // 10 : Set to 10
		F = iota             //  5 : This is the 6th constant so set to 5.
		G = 20 + iota        // 26 : Add 20 + 6
		H                    // 27 : Increment by 1
		I                    // 28 : Increment by 1
		J = iota             //  9 : This is the 10th constant so set to 9.
		K                    // 10 : Increment by 1
		L                    // 11 : Increment by 1
	)

	fmt.Println("A-L:", A, B, C, D, E, "-", F, G, H, I, "-", J, K, L)

	// New constant block with iota using bitwise operation.
	const (
		Ldate         = 1 << iota //  1 : Shift 1 to the left 0.  0000 0001
		Ltime                     //  2 : Shift 1 to the left 1.  0000 0010
		Lmicroseconds             //  4 : Shift 1 to the left 2.  0000 0100
		Llongfile                 //  8 : Shift 1 to the left 3.  0000 1000
		Lshortfile                // 16 : Shift 1 to the left 4.  0001 0000
		LUTC                      // 32 : Shift 1 to the left 5.  0010 0000
	)

	fmt.Println("Log:", Ldate, Ltime, Lmicroseconds, Llongfile, Lshortfile, LUTC)
}
