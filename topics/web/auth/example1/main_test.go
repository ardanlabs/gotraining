package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexHandlerAuthorized(t *testing.T) {
	ts := httptest.NewServer(App())
	defer ts.Close()

	req, err := http.NewRequest("GET", ts.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	req.SetBasicAuth("username", "password")

	c := http.Client{}
	res, err := c.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	want := "Welcome Authorized User!"
	got := string(b)

	if want != got {
		t.Logf("Wanted: %s", want)
		t.Logf("Got   : %s", got)
		t.Fail()
	}
}

func TestIndexHandlerInvalidCredentials(t *testing.T) {
	ts := httptest.NewServer(App())
	defer ts.Close()

	req, err := http.NewRequest("GET", ts.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	req.SetBasicAuth("badusername", "badpassword")

	c := http.Client{}
	res, err := c.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	want := http.StatusUnauthorized
	got := res.StatusCode

	if want != got {
		t.Logf("Wanted: %s", want)
		t.Logf("Got   : %s", got)
		t.Fail()
	}
}

func TestIndexHandlerNoCredentials(t *testing.T) {
	ts := httptest.NewServer(App())
	defer ts.Close()

	req, err := http.NewRequest("GET", ts.URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	c := http.Client{}
	res, err := c.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	want := http.StatusUnauthorized
	got := res.StatusCode

	if want != got {
		t.Logf("Wanted: %s", want)
		t.Logf("Got   : %s", got)
		t.Fail()
	}
}
