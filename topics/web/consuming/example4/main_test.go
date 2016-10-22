package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func App() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		time.Sleep(2 * time.Second)
	})
}

func Test_App(t *testing.T) {
	ts := httptest.NewServer(App())
	defer ts.Close()

	client := http.Client{
		Timeout: 500 * time.Millisecond,
	}
	req, err := http.NewRequest("GET", ts.URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.Do(req)

	if err == nil {
		t.Fatal("request was supposed to timeout")
	}

}
