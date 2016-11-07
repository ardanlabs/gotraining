// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Tests for the sample program to show how to apply basic authentication
// with the goth package for your web request.
package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/faux"
)

func TestIndex(t *testing.T) {

	// Startup a server to handle processing these routes.
	ts := httptest.NewServer(App())
	defer ts.Close()

	// Perform a GET request for the index page.
	res, err := http.Get(ts.URL)
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
	want := `<p><a href="/auth/github">Log in with Github</a></p>`
	if got != want {
		t.Log("Wanted:", want)
		t.Log("Got   :", got)
		t.Fatal("Mismatch")
	}
}

func TestCallback(t *testing.T) {

	// A physical location where sessions will be saved.
	gothic.Store = sessions.NewFilesystemStore(os.TempDir(), []byte("goth-test"))

	// Create a new goth faux provider.
	goth.UseProviders(&faux.Provider{})

	// Create a ResponseRecorder for our mock request.
	res := httptest.NewRecorder()

	// Create a new GET request for the user page.
	req, err := http.NewRequest("GET", "/auth/faux/callback", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a faux session with the connection details.
	sess := faux.Session{
		Name:  "Mary Jane",
		Email: "mary@example.com",
	}

	// Get the cached session.
	session, err := gothic.Store.Get(req, gothic.SessionName)
	if err != nil {
		t.Fatal(err)
	}

	// Marshal the fuz session for this test.
	session.Values[gothic.SessionName] = sess.Marshal()

	// Save this session.
	if err := session.Save(req, res); err != nil {
		t.Fatal(err)
	}

	// Process the request and get the response.
	App().ServeHTTP(res, req)

	// Read in the response from the api call.
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Validate we received the expected response.
	want := "Mary Jane"
	got := strings.TrimSpace(string(b))
	if !strings.Contains(got, want) {
		t.Logf("Wanted: %s", want)
		t.Logf("Got   : %s", got)
		t.Fail()
	}
}
