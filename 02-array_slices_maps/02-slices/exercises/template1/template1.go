// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/ASdQF3Tgv2

// Declare a nil slice of integers. Create a loop that increments a counter variable
// by 10 five times and appends these values to the slice. Iterate over the slice and
// display each value.
//
// Declare a slice of five strings and initialize the slice with string literal
// values. Display all the elements. Take a slice of the second and third elements (index 1 and 2)
// and display the index position and value of each element in the new slice.
package main

// main is the entry point for the application.
func main() {
	// Declare a nil slice of integers.
	var slice_name []type

	// Appens numbers to the slice.
	for variable_name := 0; variable_name < N; variable_name++ {
		slice_name = append(slice_name, variable_name*10)
	}

	// Display each value.
	for _, variable_name := range slice_name {
		fmt.Println(variable_name)
	}

	// Declare a slice of strings.
	slice_name := []type{Intialize values here}

	// Display each index position and name.
	for variable_name, variable_name := range slice_name {
		fmt.Printf("Index: %d  Name: %s\n", variable_name, variable_name)
	}

	// Take a slice of the second and third elements.
	slice_name2 := slice_name[I:J]

	// Display the value of the new slice.
	for variable_name, variable_name := range slice_name2 {
		fmt.Printf("Index: %d  Name: %s\n", variable_name, variable_name)
	}
}
