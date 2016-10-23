// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Tests for the sample program to show how to bundle assets, static files,
// etc into web application and access these bundled resources.
package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestApp(t *testing.T) {

	// Startup a server to handle processing these routes.
	ts := httptest.NewServer(App())
	defer ts.Close()

	// Add a set of sub-tests that can run in parallel
	// or on demand from the command line.
	t.Run("js", testJS(ts))
	t.Run("css", testCSS(ts))
	t.Run("html", testHTML(ts))
}

// testJS tests the delivery of the javascript file.
func testJS(ts *httptest.Server) func(*testing.T) {
	return func(t *testing.T) {

		// Perform a call to get the app.js file.
		res, err := http.Get(ts.URL + "/assets/app.js")
		if err != nil {
			t.Fatal(err)
		}

		// Read in the response which should be
		// the file.
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}

		// Validate we received a response that contains
		// the string we want.
		got := string(b)
		want := "getElementsByTagName"
		if !strings.Contains(got, want) {
			t.Logf("Wanted: %v", want)
			t.Logf("Got   : %v", got)
			t.Fatal("Mismatch")
		}
	}
}

// testCSS tests the delivery of the style sheet file.
func testCSS(ts *httptest.Server) func(*testing.T) {
	return func(t *testing.T) {

		// Perform a call to get the styles.css file.
		res, err := http.Get(ts.URL + "/assets/styles.css")
		if err != nil {
			t.Fatal(err)
		}

		// Read in the response which should be
		// the file.
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}

		// Validate we received a response that contains
		// the string we want.
		got := string(b)
		want := "color: blue"
		if !strings.Contains(got, want) {
			t.Logf("Wanted: %v", want)
			t.Logf("Got   : %v", got)
			t.Fatal("Mismatch")
		}
	}
}

// testHTML tests the delivery of the index file.
func testHTML(ts *httptest.Server) func(*testing.T) {
	return func(t *testing.T) {

		// Perform a call to get the index.html file.
		res, err := http.Get(ts.URL + "/assets/index.html")
		if err != nil {
			t.Fatal(err)
		}

		// Read in the response which should be
		// the file.
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}

		// Validate we received a response that contains
		// the string we want.
		got := string(b)
		want := "<title>Ultimate Web</title>"
		if !strings.Contains(got, want) {
			t.Logf("Wanted: %v", want)
			t.Logf("Got   : %v", got)
			t.Fatal("Mismatch")
		}
	}
}
