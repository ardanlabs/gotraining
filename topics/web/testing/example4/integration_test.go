// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample test to show how to test a handler using a mocking server.
package unit

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// MyHandler is the application handler we want to test.
func MyHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello World!")
}

func Test_IntegrationHandler(t *testing.T) {

	// Start an actual server that serves our handler.
	ts := httptest.NewServer(http.HandlerFunc(MyHandler))
	defer ts.Close()

	// Call the handler through the server's URL. This makes
	// an actual GET request.
	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	// Read in the response from the call.
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Validate we received the expected response.
	got := string(b)
	want := "Hello World!"
	if got != want {
		t.Log("Wanted:", want)
		t.Log("Got   :", got)
		t.Fatal("Mismatch")
	}
}
