// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Tests for the sample program to show how to apply
// middleware using negroni.
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

	// Perform the GET request for root route.
	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	// Validate we have the header key/value.
	foo := res.Header.Get("foo")
	if foo != "bar" {
		t.Fatalf("expected header foo to equal bar got %s", foo)
	}

	// Read in the response from the api call.
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Validate we received the expected response.
	got := string(b)
	want := "Hello World"
	if got != want {
		t.Log("Wanted:", want)
		t.Log("Got   :", got)
		t.Fatal("Mismatch")
	}
}
