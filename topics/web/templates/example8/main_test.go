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

	t.Run("js", testJS(ts))
	t.Run("css", testCSS(ts))
	t.Run("html", testHTML(ts))
}

func testJS(ts *httptest.Server) func(*testing.T) {
	return func(t *testing.T) {
		res, err := http.Get(ts.URL + "/assets/app.js")
		if err != nil {
			t.Fatal(err)
		}

		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}

		act := string(b)
		exp := "h1.innerHTML"

		if !strings.Contains(act, exp) {
			t.Fatalf("expected %s to contain %s", act, exp)
		}
	}
}

func testCSS(ts *httptest.Server) func(*testing.T) {
	return func(t *testing.T) {
		res, err := http.Get(ts.URL + "/assets/styles.css")
		if err != nil {
			t.Fatal(err)
		}

		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}

		act := string(b)
		exp := "color: blue"

		if !strings.Contains(act, exp) {
			t.Fatalf("expected %s to contain %s", act, exp)
		}
	}
}

func testHTML(ts *httptest.Server) func(*testing.T) {
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
		exp := "<title>Ultimate Web</title>"

		if !strings.Contains(act, exp) {
			t.Fatalf("expected %s to contain %s", act, exp)
		}
	}
}
