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
			"method": req.Method,
		})
	})
}

func Test_App(t *testing.T) {
	ts := httptest.NewServer(App())
	defer ts.Close()

	client := http.Client{}
	req, err := http.NewRequest("PUT", ts.URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	res, err := client.Do(req)

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	act := strings.TrimSpace(string(b))
	exp := `{"method":"PUT"}`
	if act != exp {
		t.Fatalf("Expected %s to equal %s", act, exp)
	}
}
