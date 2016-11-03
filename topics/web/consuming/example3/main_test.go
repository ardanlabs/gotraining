// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Tests to show how to consume a web API using the default http
// support in the standard library. This shows a PUT call.
package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// App returns a handler that can be used to mock the PUT call.
func App() http.Handler {

	// Handler function will be used for mocking. It returns
	// a document with the Method string.
	h := func(res http.ResponseWriter, req *http.Request) {
		json.NewEncoder(res).Encode(map[string]string{
			"method": req.Method,
		})
	}

	// Return the handler function.
	return http.HandlerFunc(h)
}

func TestApp(t *testing.T) {

	// Startup a server to handle processing these routes.
	ts := httptest.NewServer(App())
	defer ts.Close()

	// Create a new request for the PUT call.
	req, err := http.NewRequest("PUT", ts.URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a Client and perform the PUT call.
	var client http.Client
	res, err := client.Do(req)
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
	want := `{"method":"PUT"}`
	if got != want {
		t.Log("Wanted:", want)
		t.Log("Got   :", got)
		t.Fatal("Mismatch")
	}
}
