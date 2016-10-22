package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func Test_App(t *testing.T) {
	ts := httptest.NewServer(App())
	defer ts.Close()

	t.Run("GET", test_Get(ts))
	t.Run("POST", test_Post(ts))
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
		form := url.Values{
			"FirstName": []string{"Mary"},
			"LastName":  []string{"Jane"},
		}
		res, err := http.PostForm(ts.URL, form)
		if err != nil {
			t.Fatal(err)
		}

		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}

		act := string(b)

		expectations := []string{
			"First Name: Mary",
			"Last Name: Jane",
		}

		for _, exp := range expectations {
			if !strings.Contains(act, exp) {
				t.Fatalf("expected %s to contain %s", act, exp)
			}
		}
	}
}
