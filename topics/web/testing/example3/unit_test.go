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

//////////////////////////////////////////////////////////////////////
// Application Code                                                 //
//////////////////////////////////////////////////////////////////////

// HelloHandler is one of the handlers in our application.
func HelloHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello World!")
}

// GoodbyeHandler is the application handler we want to test.
func GoodbyeHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Goodbye Cruel World!")
}

// App bootstraps the application and returns a handler that can be
// used to serve requests.
func App() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("/goodbye", GoodbyeHandler)
	m.HandleFunc("/hello", HelloHandler)

	return m
}

//////////////////////////////////////////////////////////////////////
// Test Code                                                        //
//////////////////////////////////////////////////////////////////////

func Test_GoodbyeMux(t *testing.T) {

	// Create a new request for the /goodbye path.
	req := httptest.NewRequest("GET", "http://example.com/goodbye", nil)

	// Create a ResponseRecorder which implements http.ResponseWriter
	res := httptest.NewRecorder()

	// Get a handler for the app to test
	m := App()

	// Execute the handler through the mux. This will let
	// us also test the routes are valid.
	m.ServeHTTP(res, req)

	// Validate we received the expected response.
	got := res.Body.String()
	want := "Goodbye Cruel World!"
	if got != want {
		t.Log("Wanted:", want)
		t.Log("Got   :", got)
		t.Fatal("Mismatch")
	}
}
