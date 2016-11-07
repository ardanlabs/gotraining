// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Tests for the sample program to show how to create a basic CRUD
// based web api for customers with a middleware component.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ardanlabs/gotraining/topics/web/customer"
)

func TestRoot(t *testing.T) {

	// Startup a server to handle processing these routes.
	ts := httptest.NewServer(App())
	defer ts.Close()

	// Request the root URL `/`.
	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	// Validate we were redirected to the customers url.
	got := res.Request.URL.String()
	want := ts.URL + "/customers"
	if got != want {
		t.Log("Wanted:", got)
		t.Log("Got   :", want)
		t.Fatal("Mismatch")
	}
}

func TestIndexHandler(t *testing.T) {

	// Startup a server to handle processing these routes.
	ts := httptest.NewServer(App())
	defer ts.Close()

	// Request all the customers in the DB.
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
	got := string(b)
	want := `[{"ID":1,"Name":"Mary Jane"},{"ID":2,"Name":"Bob Smith"}]`
	if got != want {
		t.Log("Wanted:", want)
		t.Log("Got   :", got)
		t.Fatal("Mismatch")
	}
}

func TestShowHandler(t *testing.T) {

	// Startup a server to handle processing these routes.
	ts := httptest.NewServer(App())
	defer ts.Close()

	// Find customer 1 in the DB.
	c1, err := customer.Find(1)
	if err != nil {
		t.Fatal(err)
	}

	// Request customer 1 from the DB.
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

	// Validate we received all the known customers.
	got := string(b)
	want := `{"ID":1,"Name":"Mary Jane"}`
	if got != want {
		t.Log("Wanted:", want)
		t.Log("Got   :", got)
		t.Fatal("Mismatch")
	}
}

func TestCreateHandler(t *testing.T) {

	// Startup a server to handle processing these routes.
	ts := httptest.NewServer(App())
	defer ts.Close()

	// Create a JSON document from this Customer value.
	b, err := json.Marshal(customer.Customer{Name: "Jane Doe"})
	if err != nil {
		t.Fatal(err)
	}

	// Save this customer into the database.
	res, err := http.Post(ts.URL+"/customers", "application/json", bytes.NewReader(b))
	if err != nil {
		t.Fatal(err)
	}

	// Read in the response from the api call.
	b, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Validate we received all the known customers.
	got := string(b)
	want := `{"ID":3,"Name":"Jane Doe"}`
	if got != want {
		t.Log("Wanted:", want)
		t.Log("Got   :", got)
		t.Fatal("Mismatch")
	}
}
