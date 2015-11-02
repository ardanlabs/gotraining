// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// http://play.golang.org/p/vP5cZsU6uU

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
main.example(0x820123f08, 0x2, 0x4, 0x6d4f8, 0x5, 0xa)
	/Users/bill/.../example1/example1.go:15 +0x65
main.main()
	/Users/bill/.../example1/example1.go:11 +0xa5
*/
