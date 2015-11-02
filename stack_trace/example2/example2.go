// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/NdhLzZJf_X

// Sample program to show how to read a stack trace when it packs values.
package main

func main() {
	example(true, false, true, 25)
}

func example(b1, b2, b3 bool, i uint8) {
	panic("Want stack trace")
}

/*
panic: Want stack trace

goroutine 1 [running]:
main.example(0x819010001)
	/Users/bill/.../example2/example2.go:14 +0x65
main.main()
	/Users/bill/.../example2/example2.go:10 +0x2b

--------------------------------------------------------------------------------

// Parameter values
true, false, true, 25

// Word value (0x819010001)
Bits    Binary      Hex   Value
00-07   0000 0001   01    true
08-15   0000 0000   00    false
16-23   0000 0001   01    true
24-31   0001 1001   19    25

// Declaration
main.Example(b1, b2, b3 bool, i uint8)

// Stack trace
main.Example(0x819010001)
*/
