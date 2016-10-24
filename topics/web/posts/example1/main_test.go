package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_App(t *testing.T) {
	ts := httptest.NewServer(App())
	defer ts.Close()
	t.Run("GET", test_Get(ts))
	t.Run("POST", test_Post(ts))
	ts.Close()
}

func test_Get(ts *httptest.Server) func(*testing.T) {
	return func(t *testing.T) {
		res, err := http.Get(ts.URL)
		if err != nil {
			t.Fatal(err)
		}

		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}

		act := string(b)
		exp := "CLICK ME!!"
		if !strings.Contains(act, exp) {
			t.Fatalf("expected %s to contain %s", exp, act)
		}
	}
}

func test_Post(ts *httptest.Server) func(*testing.T) {
	return func(t *testing.T) {
		res, err := http.Post(ts.URL, "text/html", nil)
		if err != nil {
			t.Fatal(err)
		}

		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}

		act := string(b)
		exp := "Thank you"
		if !strings.Contains(act, exp) {
			t.Fatalf("expected %s to contain %s", exp, act)
		}
	}
}
