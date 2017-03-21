// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how the range will make a copy of the supplied
// data structure when the second value is requested during iteration.
package main

import "fmt"

func main() {

	// In this case the range is using the `five` array directly.
	five := [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	fmt.Printf("Bfr[%s] : ", five[1])

	for i := range five {
		five[1] = "Jack"

		if i == 1 {
			fmt.Printf("Aft[%s]\n", five[1])
		}
	}

	// In this case the range makes a copy of the `five` array. The v
	// variable is based on the copy.
	five = [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	fmt.Printf("Bfr[%s] : ", five[1])

	for i, v := range five {
		five[1] = "Jack"

		if i == 1 {
			fmt.Printf("v[%s]\n", v)
		}
	}

	// In this case the range makes a copy of the `five` array's address.
	// The v variable is based on the five array directly.
	five = [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	fmt.Printf("Bfr[%s] : ", five[1])

	for i, v := range &five {
		five[1] = "Jack"

		if i == 1 {
			fmt.Printf("v[%s]\n", v)
		}
	}
}
