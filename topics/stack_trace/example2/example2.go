// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

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
panic(0x569e0, 0xc82000a110)
	/usr/local/go/src/runtime/panic.go:464 +0x3e6
main.example(0xc819010001)
	/Users/bill/.../stack_trace/example2/example2.go:12 +0x65
main.main()
	/Users/bill/.../stack_trace/example2/example2.go:8 +0x2b

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
