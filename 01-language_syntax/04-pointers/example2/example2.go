// http://play.golang.org/p/JEVXsBtt8Z

// Sample program to show the basic concept of using a pointer
// to share data.
package main

import "fmt"

// main is the entry point for the application.
func main() {
	// Declare variable of type int with a value of 10.
	count := 10

	// Display the "value of" and "address of" count.
	fmt.Println("Before:", count, &count)

	// Pass the "address of" the variable count.
	inc(&count)

	fmt.Println("After:", count, &count)
}

// inc declares count as a pointer variable whose value is
// always an address and points to values of type int.
func inc(count *int) {
	// Increment the value that the "pointer points to". (de-referencing)
	*count++
	fmt.Println("Inc:", &count, count, *count)
}
