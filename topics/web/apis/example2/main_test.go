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
		req, err := http.NewRequest("GET", ts.URL+"/api/users", nil)
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("x-version", v)

		c := http.Client{}
		res, err := c.Do(req)
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
