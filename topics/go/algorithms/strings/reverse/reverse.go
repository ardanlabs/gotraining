package main

import (
	"fmt"
)

func main() {

	fmt.Println(ReverseString("Hello World"))
}

func ReverseString(str string) string {

	// Here we create a slice of runes.
	str_table := []rune(str)

	// Here we create int that will be a pointer to the front of the str_table.
	frontIndex := 0

	// Here  we create int that will be a pointer to the end of the str_table.
	BackIndex := len(str_table) - 1

	// While frontIndex is less than BackIndex.
	for frontIndex < BackIndex {

		// Here we swap the values of the slice at the str_table[frontIndex] and str_table[BackIndex].
		str_table[frontIndex], str_table[BackIndex] = str_table[BackIndex], str_table[frontIndex]

		// Here we increase the value of frontIndex by 1.
		frontIndex = frontIndex + 1

		// Here we decrease the value of BackIndex by 1.
		BackIndex = BackIndex - 1
	}

	// Here we return a new string in reverse order.
	return string(str_table)
}
