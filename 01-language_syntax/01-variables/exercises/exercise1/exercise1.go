// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/Kr7CaO6LdF

// Declare three variables that are initalized to their zero value and three
// declared with a literal value. Declare variables of type string, int and
// bool. Display the values of those variables.
//
// Declare a new variable of type float32 and initalize the variable by
// converting the literal value of Pi (3.14).
package main

import "fmt"

// main is the entry point for the application.
func main() {
	// Declare variables that are set to their zero value.
	var age int
	var name string
	var legal bool

	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(legal)

	// Declare variables and initalize.
	// Using the short variable declaration operator.
	month := 10
	dayOfWeek := "Tuesday"
	happy := true

	fmt.Println(month)
	fmt.Println(dayOfWeek)
	fmt.Println(happy)

	// Specify type and perform a conversion.
	pi := float32(3.14)

	fmt.Printf("%T [%v]\n", pi, pi)
}
