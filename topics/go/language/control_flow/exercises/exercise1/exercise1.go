// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Write a program that inspects a user's name and greets them in a certain way
// if they are on a list or in a different way if they are not. Also look at
// the user's age and tell them some special secret if they are old enough to
// know it.
package main

import "fmt"

func main() {

	// Change these values and rerun your program.
	name := "Carter"
	age := 6

	// If the user's name is on a special list then give them a secret greeting.
	switch name {
	case "Anna", "Jacob", "Kell", "Carter", "Rory":
		fmt.Println("What's up, Walker family?")

	case "Seth", "Julia", "Tanner", "Kenton", "Britten":
		fmt.Println("Welcome, my friend!")

	default:
		fmt.Println("It's nice to meet you!")
	}

	// If the user is old enough then tell them a secret.
	if age > 10 {
		fmt.Println("The tooth fairy is just your mom or dad.")
	} else {
		fmt.Println("What do you think unicorns dream about?")
	}
}
