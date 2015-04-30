// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// Sample program to show how to test routes without running a server.
package main_test

import (
	"log"
	"os"

	ex3 "github.com/ArdanStudios/gotraining/09-testing/01-testing/example3"
)

// ExampleLogResponse provides a basic example test example.
func ExampleLogResponse() {
	log.SetFlags(0)
	log.SetOutput(os.Stdout)

	u := struct {
		Name  string
		Email string
	}{
		Name:  "Bill",
		Email: "bill@ardanstudios.com",
	}

	ex3.LogResponse(&u)
	// Output:
	// {
	//     "Name": "Bill",
	//     "Email": "bill@ardanstudios.com"
	// }
}
