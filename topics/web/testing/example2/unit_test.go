// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample test to show how to write a unit test that declares
// a type that implement the http Handler interface.
package unit

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// App provides our application context.
type App struct{}

// ServeHTTP implements the http Handler interface so it can
// provide support for mocking the GET call.
func (a App) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello World!")
}

func Test_App(t *testing.T) {

	// Create a new request.
	req, err := http.NewRequest("GET", "http://example.com/foo", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new ResponseRecorder which implements
	// the ResponseWriter interface.
	res := httptest.NewRecorder()

	// Create the application value.
	var a App

	// Execute the handler from the application value.
	a.ServeHTTP(res, req)

	got := res.Body.String()
	want := "Hello World!"
	if got != want {
		t.Log("Wanted:", want)
		t.Log("Got   :", got)
		t.Fatal("Mismatch")
	}
}
