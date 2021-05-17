// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to read a stack trace when it packs values.
package main

func main() {
	example(true, false, true, 25)
}

//go:noinline
func example(b1, b2, b3 bool, i uint8) error {
	panic("Want stack trace")
}

/*
	panic: Want stack trace

	goroutine 1 [running]:
	main.example(0x19010001, 0x1064ee0, 0xc000076058)
		stack_trace/example2/example2.go:13 +0x39
	main.main()
		stack_trace/example2/example2.go:8 +0x29

--------------------------------------------------------------------------------

	// Declaration
	main.example(b1, b2, b3 bool, i uint8)

	// Call
	main.example(true, false, true, 25)

	// Stack trace
	main.example(0x19010001, 0x1064ee0, 0xc000076058)

	// Word value (0xc019010001)
	Bits    Binary      Hex   Value
	00-07   0000 0001   01    true
	08-15   0000 0000   00    false
	16-23   0000 0001   01    true
	24-31   0001 1001   19    25

	Return Arguments: 0x1064ee0, 0xc00007605

	Use `go build -gcflags -S` to map the PC offset values, +0x39 and +0x29 for
	each function call.
*/

// Note: https://go-review.googlesource.com/c/go/+/109918
