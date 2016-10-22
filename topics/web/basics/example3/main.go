// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to implement your own Handler.
package main

import (
	"fmt"
	"log"
	"net/http"
)

// App provides application level context for our handler.
type App struct{}

// ServeHTTP implements the http.Handler interface.
func (a App) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello World!")
}

func main() {

	// Start the http server to handle the requests using
	// our new Handler.
	log.Panic(http.ListenAndServe(":3000", App{}))
}
