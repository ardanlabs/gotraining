// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Tests for the sample program that sends and receives JSON.
package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGet(t *testing.T) {

	// Start a server to handle these requests.
	ts := httptest.NewServer(App())
	defer ts.Close()

	// Request the root URL `/`.
	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	if got, want := res.Header.Get("Content-Type"), "application/json; charset=utf-8"; got != want {
		t.Error("Response content type did not match")
		t.Log("Wanted:", want)
		t.Log("Got   :", got)
	}

	// Read in the response from the api call.
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	got := strings.TrimSpace(string(b))
	want := `{"load":1,"messages":42}`
	if got != want {
		t.Log("Wanted:", want)
		t.Log("Got   :", got)
		t.Fatal("Mismatch")
	}
}

func TestPost(t *testing.T) {

	// Start a server to handle these requests.
	ts := httptest.NewServer(App())
	defer ts.Close()

	event := `{"host": "localhost", "level": 5, "message": "Hello, world!"}`

	// Send this event to the system
	res, err := http.Post(ts.URL+"/customers", "application/json", strings.NewReader(event))
	if err != nil {
		t.Fatal(err)
	}

	got := res.StatusCode
	want := 204
	if got != want {
		t.Log("Wanted:", want)
		t.Log("Got   :", got)
		t.Fatal("Mismatch")
	}
}
