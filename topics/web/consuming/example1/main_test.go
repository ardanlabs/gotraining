package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func App() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		json.NewEncoder(res).Encode(map[string]string{
			"first_name": "Mary",
			"last_name":  "Jane",
		})
	})
}

func Test_App(t *testing.T) {
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

	act := strings.TrimSpace(string(b))
	exp := `{"first_name":"Mary","last_name":"Jane"}`
	if act != exp {
		t.Fatalf("Expected %s to equal %s", act, exp)
	}
}
