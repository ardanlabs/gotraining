// Sample program to show how to create and use variadic functions.
package main

import "fmt"

// user is a struct type that declares user information.
type user struct {
	ID   int
	Name string
}

// main is the entry point for the application.
func main() {
	log("TRACE", "main", "Started")

	// Declare and initalize a value of type user.
	u := user{
		ID:   1432,
		Name: "Betty",
	}

	// Log each individual field from the user value.
	log("TRACE", "main", "ID[%d] Name[%s]", u.ID, u.Name)

	log("TRACE", "main", "Completed")
}

// log provides simple formatting for logging the program.
func log(title string, functionName string, format string, a ...interface{}) {
	fmt.Printf("%s : %s : %s\n", title, functionName, fmt.Sprintf(format, a...))
}
