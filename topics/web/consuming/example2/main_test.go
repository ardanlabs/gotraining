// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Tests to show how to consume a web API using the default http
// support in the standard library. This shows a POST call.
package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// App returns a handler that can be used to mock the POST call.
func App() http.Handler {

	// Handler function will be used for mocking. It just
	// returns the request back as the repsonse.
	h := func(res http.ResponseWriter, req *http.Request) {
		io.Copy(res, req.Body)
	}

	// Return the handler function.
	return http.HandlerFunc(h)
}

func TestApp(t *testing.T) {

	// Startup a server to handle processing these routes.
	ts := httptest.NewServer(App())
	defer ts.Close()

	// Generate a JSON document for this map.
	b, err := json.Marshal(map[string]string{
		"first_name": "Mary",
		"last_name":  "Jane",
	})
	if err != nil {
		t.Fatal(err)
	}

	// Perform the POST request to our mock handler.
	res, err := http.Post(ts.URL, "application/json", bytes.NewReader(b))
	if err != nil {
		t.Fatal(err)
	}

	// Read in the response from the api call.
	if b, err = ioutil.ReadAll(res.Body); err != nil {
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
