// "make" has a 3 parameter for initial capacity to avoid redundant memory allocations.

package main

import (
	"fmt"
)

func main() {
	const size = 50
	values := make([]int, 0, size)
	for i := 0; i < size; i++ {
		values = append(values, i*i)
	}
	fmt.Println(values)
}
