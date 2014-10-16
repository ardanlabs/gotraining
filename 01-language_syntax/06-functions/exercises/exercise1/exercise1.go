// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/hJ532ON6Xb

// Declare a struct type to maintain information about a user. Declare a function
// that creates value of and returns pointers of this type and an error value. Call
// this function from main and display the value.
//
// Make a second call to your function but this time ignore the value and just test
// the error value.
package main

import "fmt"

// user represents a user in the system.
type user struct {
	name  string
	email string
}

// main is the entry point for the application.
func main() {
	// Create a value of user.
	u, err := newUser()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Display the value.
	fmt.Println(u)

	// Create a value of user and ignore the
	// second value being returned.
	u2, _ := newUser()
	fmt.Println(u2)
}

// newUser creates user values.
func newUser() (*user, error) {
	return &user{"Bill", "bill@ardanstudios.com"}, nil
}
