// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample test to show how to write a unit test.
package unit

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// MyHandler is provide support for mocking the GET call.
func MyHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello World!")
}

func Test_MyHandler(t *testing.T) {

	// Create a new request.
	req, err := http.NewRequest("GET", "http://example.com/foo", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new ResponseRecorder which implements
	// the ResponseWriter interface.
	res := httptest.NewRecorder()

	// Execute the handler function direction.
	MyHandler(res, req)

	// Validate we received all the known customers.
	got := res.Body.String()
	want := "Hello World!"
	if got != want {
		t.Log("Wanted:", want)
		t.Log("Got   :", got)
		t.Fatal("Mismatch")
	}
}
