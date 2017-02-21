// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Tests for the sample program to show how to have a single route
// for the api but have access to either through configuration.
package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApp(t *testing.T) {

	// Startup a server to handle processing these routes.
	ts := httptest.NewServer(App())
	defer ts.Close()

	// Define the different versions for this test.
	tests := []struct {
		version string
	}{
		{"v1"}, {"v2"},
	}

	for _, tt := range tests {

		// Setup a request for user api without thinking about version.
		req, err := http.NewRequest("GET", ts.URL+"/api/users", nil)
		if err != nil {
			t.Fatal(err)
		}

		// Set the version header into the request.
		req.Header.Set("x-version", tt.version)

		// Create a client and perform the request.
		var c http.Client
		res, err := c.Do(req)
		if err != nil {
			t.Fatal(err)
		}

		// Read in the response from the api call.
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}

		// Validate we received the correct version.
		got := string(b)
		want := tt.version
		if got != want {
			t.Log("Wanted:", want)
			t.Log("Got   :", got)
			t.Fatal("Mismatch")
		}
	}
}
