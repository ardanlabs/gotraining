package main

import "fmt"

func main() {
	count := 10

	// Display the "value of" and "address of" count.
	fmt.Println("Before:", count, &count)

	// Pass the "value of" the variable count.
	Inc(count)

	fmt.Println("After:", count, &count)
}

// Declaring count as a variable whose value is
// always an integer.
func Inc(count int) {
	// Increment the "value of" count.
	count++
	fmt.Println("Inc:", count, &count)
}
