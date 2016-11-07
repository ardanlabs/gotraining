// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Tests for the sample program to show how to handle forms
// using a single handler for the GET and POST calls.
package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestApp(t *testing.T) {

	// Startup a server to handle processing these routes.
	ts := httptest.NewServer(App())
	defer ts.Close()

	// Create a sub-test for each verb.
	t.Run("GET", testGet(ts))
	t.Run("POST", testPost(ts))
}

// testGet validates the GET verb.
func testGet(ts *httptest.Server) func(*testing.T) {

	// Test function for execution as a sub-test.
	tf := func(t *testing.T) {

		// Perform a GET call against the url.
		res, err := http.Get(ts.URL)
		if err != nil {
			t.Fatal(err)
		}

		// Read in the response from the call.
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}

		// Validate we received the correct document.
		got := string(b)
		want := `
<form action="/" method="POST">
<p>
	<input type="text" name="FirstName" placeholder="First Name" />
</p>
<p>
	<input type="text" name="LastName" placeholder="Last Name" />
</p>
<p>
	<input type="submit" value="CLICK ME!!" />
</p>
</form>`
		if got != want {
			t.Log("Wanted:", want)
			t.Log("Got   :", got)
			t.Fatal("Mismatch")
		}
	}

	return tf
}

// testPost validates the POST verb.
func testPost(ts *httptest.Server) func(*testing.T) {

	// Test function for execution as a sub-test.
	tf := func(t *testing.T) {

		// Add form variables with expected values.
		form := url.Values{
			"FirstName": []string{"Mary"},
			"LastName":  []string{"Jane"},
		}

		// Perform a POST call against the url.
		res, err := http.PostForm(ts.URL, form)
		if err != nil {
			t.Fatal(err)
		}

		// Read in the response from the call.
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}

		// Validate we received the correct document.
		got := string(b)
		want := "First Name: Mary\nLast Name: Jane"
		if got != want {
			t.Log("Wanted:", want)
			t.Log("Got   :", got)
			t.Fatal("Mismatch")
		}
	}

	return tf
}
