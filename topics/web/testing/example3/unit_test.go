// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample test to show how to write a unit test that also
// tests the routes inside the mux.
package unit

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// HelloHandler is provide support for mocking the GET call.
func HelloHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello World!")
}

// GoodbyeHandler is provide support for mocking the GET call.
func GoodbyeHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Goodbye Cruel World!")
}

func Test_MyHandler(t *testing.T) {

	// Create a new request.
	req, err := http.NewRequest("GET", "http://example.com/goodbye", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new ResponseRecorder which implements
	// the ResponseWriter interface.
	res := httptest.NewRecorder()

	// Create a mux instead of using the default. Bind the
	// handler inside the mux.
	m := http.NewServeMux()
	m.HandleFunc("/goodbye", GoodbyeHandler)
	m.HandleFunc("/hello", HelloHandler)

	// Execute the handler through the mux. This will let
	// us also test the routes are valid.
	m.ServeHTTP(res, req)

	// Validate we received all the known customers.
	got := res.Body.String()
	want := "Goodbye Cruel World!"
	if got != want {
		t.Log("Wanted:", want)
		t.Log("Got   :", got)
		t.Fatal("Mismatch")
	}
}
