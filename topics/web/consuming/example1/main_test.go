// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Tests to show how to consume a web API using the default http
// support in the standard library.
package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// App returns a handler that can be used to mock the GET call.
func App() http.Handler {

	// Handler function will be used for mocking.
	h := func(res http.ResponseWriter, req *http.Request) {
		json.NewEncoder(res).Encode(map[string]string{
			"first_name": "Mary",
			"last_name":  "Jane",
		})
	}

	// Return the handler function.
	return http.HandlerFunc(h)
}

func TestApp(t *testing.T) {

	// Startup a server to handle processing these routes.
	ts := httptest.NewServer(App())
	defer ts.Close()

	// Perform the GET request for our mock handler.
	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	// Read in the response from the api call.
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Validate we received the expected response.
	got := strings.TrimSpace(string(b))
	want := `{"first_name":"Mary","last_name":"Jane"}`
	if got != want {
		t.Log("Wanted:", want)
		t.Log("Got   :", got)
		t.Fatal("Mismatch")
	}
}
