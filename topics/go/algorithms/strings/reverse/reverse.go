package main

import "fmt"

func main() {
	// Run ReverseString
	fmt.Println(ReverseString("Hello World"))
}

func ReverseString(str string) string {
	// Create rune.
	chars := []rune(str)

	// Reverse String.
	for start, end := 0, len(chars)-1; start < end; start, end = start+1, end-1 {
		// Swap values.
		chars[start], chars[end] = chars[end], chars[start]
	}

	// Return a new string in reverse order.
	return string(chars)
}
