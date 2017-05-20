// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to create handlers out of any function
// using http.HandlerFunc.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// fn is just a function.
	fn := func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "Hello from a closure")
	}

	// http.HandlerFunc is a type so the parens () here are doing a type
	// conversion not calling a function. h is a value of type
	// http.HandlerFunc which conveniently implements http.Handler
	h := http.HandlerFunc(fn)

	// Start the http server using our funcy Handler ;)
	log.Print("Listening on localhost:3000")
	log.Fatal(http.ListenAndServe("localhost:3000", h))
}
