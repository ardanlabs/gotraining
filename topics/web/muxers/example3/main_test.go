// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Tests for the sample program to show how to use the echo toolkit.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/ardanlabs/gotraining/topics/web/customer"
)

func TestIndexHandler(t *testing.T) {

	// Startup a server to handle processing these routes.
	ts := httptest.NewServer(App())
	defer ts.Close()

	// Perform the GET request for the list of customers.
	res, err := http.Get(ts.URL + "/customers")
	if err != nil {
		t.Fatal(err)
	}

	// Read in the response from the api call.
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Validate we received all the known customers.
	body := string(b)
	for _, n := range []string{"Mary Jane", "Bob Smith"} {
		if !strings.Contains(body, n) {
			t.Fatalf("Expected %s to contain %s", body, n)
		}
	}
}

func TestShowHandler(t *testing.T) {

	// Startup a server to handle processing these routes.
	ts := httptest.NewServer(App())
	defer ts.Close()

	// Retrieve customer 1, `Mary Jane`.
	c1, err := customer.Find(1)
	if err != nil {
		t.Fatal(err)
	}

	// Perform a GET request to retrieve customer 1.
	url := fmt.Sprintf("%s/customers/%d", ts.URL, c1.ID)
	res, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}

	// Read in the response from the api call.
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Validate we received the correct customer.
	body := string(b)
	if !strings.Contains(body, c1.Name) {
		t.Fatalf("Expected %s to contain %s", body, c1.Name)
	}
}

func TestCreateHandler(t *testing.T) {

	// Startup a server to handle processing these routes.
	ts := httptest.NewServer(App())
	defer ts.Close()

	customerName := "Jane Doe"

	// Perform a POST call to create a new customer.
	res, err := http.PostForm(ts.URL+"/customers", url.Values{"name": []string{customerName}})
	if err != nil {
		t.Fatal(err)
	}

	// Read in the response from the api call.
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Validate we received the correct customer back.
	body := string(b)
	if !strings.Contains(body, customerName) {
		t.Fatalf("Expected %s to contain %s", body, customerName)
	}

	// Retrieve customer 3, `Jane Doe`.
	c3, err := customer.Find(3)
	if err != nil {
		t.Fatal(err)
	}

	// Perform a GET request to retrieve customer 1.
	url := fmt.Sprintf("%s/customers/%d", ts.URL, c3.ID)
	res, err = http.Get(url)
	if err != nil {
		t.Fatal(err)
	}

	// Read in the response from the api call.
	b, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Validate we received the correct customer.
	body = string(b)
	if !strings.Contains(body, c3.Name) {
		t.Fatalf("Expected %s to contain %s", body, c3.Name)
	}
}
