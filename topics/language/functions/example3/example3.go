// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// From Spec:
// a short variable declaration may redeclare variables provided they
// were originally declared earlier in the same block with the same
// type, and at least one of the non-blank variables is new.

// Sample program to show some of the mechanics behind the
// short variable declaration operator redeclares.
package main

import "fmt"

// user is a struct type that declares user information.
type user struct {
	id   int
	name string
}

func main() {

	// Declare the error variable.
	var err1 error

	// The short variable declaration operator will
	// declare u and redeclare err1.
	u, err1 := getUser()
	if err1 != nil {
		return
	}

	fmt.Println(u)

	// The short variable declaration operator will
	// redeclare u and declare err2.
	u, err2 := getUser()
	if err2 != nil {
		return
	}

	fmt.Println(u)
}

// getUser returns a pointer of type user.
func getUser() (*user, error) {
	return &user{1432, "Betty"}, nil
}
