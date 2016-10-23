// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Tests for the sample program to show how to create
// a simple web api with different versions.
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

	// Range over the different versions of the API we have.
	for _, tt := range tests {

		// Call the user api for the specified version.
		res, err := http.Get(ts.URL + "/api/" + tt.version + "/users")
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
