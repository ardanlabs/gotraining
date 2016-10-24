package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func Test_indexHandler(t *testing.T) {
	ts := httptest.NewServer(App())
	defer ts.Close()

	res, err := http.Get(ts.URL + "/customers")
	if err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	body := string(b)
	for _, n := range []string{"Mary Jane", "Bob Smith"} {
		if !strings.Contains(body, n) {
			t.Fatalf("Expected %s to contain %s", body, n)
		}
	}
}

func Test_showHandler(t *testing.T) {
	ts := httptest.NewServer(App())
	defer ts.Close()

	c1, _ := Customers.Find("1")
	c2, _ := Customers.Find("2")

	res, err := http.Get(ts.URL + "/customers/" + c1.ID)
	if err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	body := string(b)

	if !strings.Contains(body, c1.Name) {
		t.Fatalf("Expected %s to contain %s", body, c1.Name)
	}
	if strings.Contains(body, c2.Name) {
		t.Fatalf("(Echo) Expected %s to not contain %s", body, c2.Name)
	}
}

func Test_createHandler(t *testing.T) {
	ts := httptest.NewServer(App())
	defer ts.Close()

	res, err := http.PostForm(ts.URL+"/customers", url.Values{"name": []string{"Jane Doe"}})
	if err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	body := string(b)

	if !strings.Contains(body, "Jane Doe") {
		t.Fatalf("(echo) Expected %s to contain %s", body, "Jane Doe")
	}

	for _, n := range []string{"Mary Jane", "Bob Smith"} {
		if strings.Contains(body, n) {
			t.Fatalf("Expected %s to not contain %s", body, n)
		}
	}
}
