// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Tests for the sample program to show how to use a
// cookie in your web app.
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
	t.Run("POST/GET", testPostGet(ts))
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
<html>
    <form action="/save" method="POST">
        <label>What is your name?</label><br>
        <input type="text" name="myName" placeholder="Name goes here">
        <input type="submit" value="Submit">
    </form>
</html>`
		if got != want {
			t.Log("Wanted:", want)
			t.Log("Got   :", got)
			t.Fatal("Mismatch")
		}
	}

	return tf
}

// testPost validates the POST verb.
func testPostGet(ts *httptest.Server) func(*testing.T) {

	// Test function for execution as a sub-test.
	tf := func(t *testing.T) {

		// Add form variables with expected values.
		form := url.Values{
			"myName": []string{"Mary"},
		}

		// Perform a POST call against the url.
		res, err := http.PostForm(ts.URL+"/save", form)
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
<html>
    <h1>Hello Mary!</h1>
</html>`
		if got != want {
			t.Log("Wanted:", want)
			t.Log("Got   :", got)
			t.Fatal("Mismatch")
		}

		// Set up a second GET request.
		req, err := http.NewRequest("GET", ts.URL, nil)
		if err != nil {
			t.Fatal(err)
		}

		// Copy the cookies over into the new request
		// for the call.
		for _, c := range res.Cookies() {
			req.AddCookie(c)
		}

		// Create a client and perform the GET request.
		var c http.Client
		res, err = c.Do(req)
		if err != nil {
			t.Fatal(err)
		}

		// Read in the response from the call.
		b, err = ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}

		// Validate we received the correct document.
		got = string(b)
		want = `
<html>
    <h1>Hello Mary!</h1>
</html>`
		if got != want {
			t.Log("Wanted:", want)
			t.Log("Got   :", got)
			t.Fatal("Mismatch")
		}
	}

	return tf
}
