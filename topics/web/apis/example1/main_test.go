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

	for _, v := range []string{"v1", "v2"} {
		res, err := http.Get(ts.URL + "/api/" + v + "/users")
		if err != nil {
			t.Fatal(err)
		}
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
		act := string(b)
		if act != v {
			t.Fatalf("expected %s, got %s", v, act)
		}
	}
}
