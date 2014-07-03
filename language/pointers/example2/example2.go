package main

import "fmt"

func main() {
	count := 10

	// Display the "value of" and "address of" count.
	fmt.Println("Before:", count, &count)

	// Pass the "address of" the variable count.
	Inc(&count)

	fmt.Println("After:", count, &count)
}

// Declaring bark as a pointer variable whose value is
// always an address and points to values of type int.
func Inc(count *int) {
	// Increment the value that the "pointer points to". (de-referencing)
	*count++
	fmt.Println("Inc:", &count, count, *count)
}
