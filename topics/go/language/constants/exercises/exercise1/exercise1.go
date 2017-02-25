// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare an untyped and typed constant and display their values.
//
// Multiply two literal constants into a typed variable and display the value.
package main

import "fmt"

const (
	// server is the IP address for connecting.
	server = "124.53.24.123"

	// port is the port to make that connection.
	port int16 = 9000
)

func main() {

	// Display the server information.
	fmt.Println(server)
	fmt.Println(port)

	// Calculate the number of minutes in 5320 seconds.
	minutes := 5320 / 60.0
	fmt.Println(minutes)
}
