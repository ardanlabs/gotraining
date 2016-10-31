package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApp(t *testing.T) {
	ts := httptest.NewServer(App())
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	want := "mary-jane"
	got := string(b)
	if got != want {
		t.Log("Wanted:", want)
		t.Log("Got   :", got)
		t.Fatal("Mismatch")
	}
}
