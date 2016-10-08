package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_App(t *testing.T) {
	ts := httptest.NewServer(App())
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	foo := res.Header.Get("foo")
	if foo != "bar" {
		t.Fatalf("expected header foo to equal bar got %s", foo)
	}

	b, err := ioutil.ReadAll(res.Body)
	exp := "Hello World"
	act := string(b)
	if act != exp {
		t.Fatalf("expected %s got %s", exp, act)
	}
}
