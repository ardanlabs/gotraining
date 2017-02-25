// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program that implements a simple web service.
package main

import (
	"log"
	"net/http"

	"github.com/ardanlabs/gotraining/topics/testing/tests/example4/handlers"
)

func main() {
	handlers.Routes()

	log.Println("listener : Started : Listening on: http://localhost:4000")
	http.ListenAndServe(":4000", nil)
}
