// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/A8yp-bzj1H

// Sample program to show how to read a stack trace.
package main

func main() {
	slice := make([]string, 2, 4)
	Example(slice, "hello", 10)
}

func Example(slice []string, str string, i int) {
	panic("Want stack trace")
}

/*
panic: Want stack trace

goroutine 1 [running]:
main.Example(0x820123f08, 0x2, 0x4, 0x6d4f8, 0x5, 0xa)
	/Users/bill/code/go/src/github.com/ardanstudios/gotraining/10-testing/06-stack_trace/example1/example1.go:15 +0x65
main.main()
	/Users/bill/code/go/src/github.com/ardanstudios/gotraining/10-testing/06-stack_trace/example1/example1.go:11 +0xa5
*/
