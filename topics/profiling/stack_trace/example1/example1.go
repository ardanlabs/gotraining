// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to read a stack trace.
package main

func main() {
	slice := make([]string, 2, 4)
	example(slice, "hello", 10)
}

func example(slice []string, str string, i int) {
	panic("Want stack trace")
}

/*
panic: Want stack trace

goroutine 1 [running]:
panic(0x56a60, 0xc82000a110)
	/usr/local/go/src/runtime/panic.go:464 +0x3e6
main.example(0xc82003bf08, 0x2, 0x4, 0x708a8, 0x5, 0xa)
	/Users/bill/.../stack_trace/example1/example1.go:13 +0x65
main.main()
	/Users/bill/.../stack_trace/example1/example1.go:9 +0xa5
*/
