// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/POH6kq8KLL

// Sample program to show the basic concept of pass by value.
package main

import "fmt"

// main is the entry point for the application.
func main() {
	// Declare variable of type int with a value of 10.
	count := 10

	// Display the "value of" and "address of" count.
	fmt.Println("Before:", count, &count)

	// Pass the "value of" the variable count.
	increment(count)

	fmt.Println("After:", count, &count)
}

// increment declares count as a variable whose value is
// always an integer.
func increment(count int) {
	// Increment the "value of" count.
	count++
	fmt.Println("Inc:", count, &count)
}
