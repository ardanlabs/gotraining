package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func App() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello World!"))
	})
}

type customTransporter struct {
	transporter http.RoundTripper
	logs        []string
}

func (c *customTransporter) RoundTrip(req *http.Request) (*http.Response, error) {
	c.logs = append(c.logs, "started request")
	res, err := c.transporter.RoundTrip(req)
	c.logs = append(c.logs, "completed request")
	return res, err
}

func Test_App(t *testing.T) {
	ts := httptest.NewServer(App())
	defer ts.Close()

	trans := &customTransporter{
		transporter: http.DefaultTransport,
		logs:        []string{},
	}

	client := http.Client{
		Transport: trans,
	}

	req, err := http.NewRequest("GET", ts.URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	act := strings.TrimSpace(string(b))
	exp := `Hello World!`
	if act != exp {
		t.Fatalf("Expected %s to equal %s", act, exp)
	}

	elogs := []string{"started request", "completed request"}
	if strings.Join(elogs, " ") != strings.Join(trans.logs, " ") {
		t.Fatalf("expected %s to equal %s", trans.logs, elogs)
	}
}
