// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Tests for the sample program to show how to apply basic
// authentication to your web request.
package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestIndexHandler(t *testing.T) {
	tests := []struct {
		Username   string
		Password   string
		Want       string
		StatusCode int
	}{
		{"username", "password", "Welcome Authorized User!", http.StatusOK},
		{"badusername", "badpassword", "Not authorized", http.StatusUnauthorized},
		{"", "", "Not authorized", http.StatusUnauthorized},
	}

	// Startup a server to handle processing these routes.
	ts := httptest.NewServer(App())
	defer ts.Close()

	for _, tt := range tests {

		// Create a new request for the GET call.
		req, err := http.NewRequest("GET", ts.URL, nil)
		if err != nil {
			t.Fatal(err)
		}

		// Only apply the credentials if we have them.
		if tt.Username != "" {

			// Set the username and password into the request.
			req.SetBasicAuth(tt.Username, tt.Password)
		}

		// Create a Client and perform the GET call.
		var c http.Client
		res, err := c.Do(req)
		if err != nil {
			t.Fatal(err)
		}

		if res.StatusCode != tt.StatusCode {
			t.Log("Wanted:", tt.StatusCode)
			t.Log("Got   :", res.StatusCode)
			t.Fatal("Mismatch")
		}

		// Read in the response from the api call.
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}

		// Validate we received the expected response.
		got := strings.TrimSpace(string(b))
		want := tt.Want
		if got != want {
			t.Log("Wanted:", want)
			t.Log("Got   :", got)
			t.Fatal("Mismatch")
		}
	}
}
