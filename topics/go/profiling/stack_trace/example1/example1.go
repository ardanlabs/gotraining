// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to read a stack trace.
package main

func main() {
	example(make([]string, 2, 4), "hello", 10)
}

//go:noinline
func example(slice []string, str string, i int) error {
	panic("Want stack trace")
}

/*
	panic: Want stack trace

	goroutine 1 [running]:
	main.example(0xc000042748, 0x2, 0x4, 0x106abae, 0x5, 0xa, 0x0, 0xc000054778)
		stack_trace/example1/example1.go:13 +0x39
	main.main()
		stack_trace/example1/example1.go:8 +0x85

	--------------------------------------------------------------------------------

	// Declaration
	main.example(slice []string, str string, i int) error

    // Call
    main.example(make([]string, 2, 4), "hello", 10)

    // Values (0xc000042748, 0x2, 0x4, 0x106abae, 0x5, 0xa, 0x0, 0xc000054778)
    Slice Value:      0xc000042748, 0x2, 0x4
    String Value:     0x106abae, 0x5
    Integer Value:    0xa
    Return Arguments: 0x0, 0xc000054778
*/

// Note: https://go-review.googlesource.com/c/go/+/109918
