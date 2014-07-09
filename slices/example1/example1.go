// Example shows that capacity is not available for use.
package main

import (
	"fmt"
)

func main() {
	slice := make([]string, 5, 8)
	slice[0] = "Apple"
	slice[1] = "Orange"
	slice[2] = "Banana"
	slice[3] = "Grape"
	slice[4] = "Plum"

	// You can't access an element of a slice beyond its length.
	slice[5] = "Runtime error"

	fmt.Println(slice)
}

// panic: runtime error: index out of range
