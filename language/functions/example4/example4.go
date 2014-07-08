// Example shows how to create and use variadic functions.
package main

import "fmt"

// user is a struct type that declares user information.
type user struct {
	Id   int
	Name string
}

func main() {
	Log("TRACE", "main", "Started")

	// Declare and initalize a value of type user.
	u := user{
		Id:   1432,
		Name: "Betty",
	}

	// Log each individual field from the user value.
	Log("TRACE", "main", "Id[%d] Name[%s]", u.Id, u.Name)

	Log("TRACE", "main", "Completed")
}

// Log provides simple formatting for logging the program.
func Log(title string, functionName string, format string, a ...interface{}) {
	fmt.Printf("%s : %s : %s\n", title, functionName, fmt.Sprintf(format, a...))
}
