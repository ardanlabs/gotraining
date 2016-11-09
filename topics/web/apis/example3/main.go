// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to create a basic CRUD based web api
// for customers.
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/ardanlabs/gotraining/topics/web/customer"
	"github.com/gorilla/pat"
)

// App loads the entire API set together for use.
func App() http.Handler {

	// Create a version of the pat router.
	r := pat.New()

	// Define the routes and order matters.
	r.Get("/customers/{id}", showHandler)
	r.Get("/customers", indexHandler)
	r.Post("/customers", createHandler)

	// Redirect requests from `/`` to `/customers`.
	r.Handle("/", http.RedirectHandler("/customers", http.StatusMovedPermanently))

	return r
}

// indexHandler returns the entire list of customers in the DB.
func indexHandler(res http.ResponseWriter, req *http.Request) {

	// Retrieve the list of customers, encode to JSON
	// and send the response.
	if err := json.NewEncoder(res).Encode(customer.All()); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

// showHandler returns a single specified customer.
func showHandler(res http.ResponseWriter, req *http.Request) {

	// Retrieve the customer id from the request.
	idStr := req.URL.Query().Get(":id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	// Retreive that customer from the DB.
	c, err := customer.Find(id)
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
		return
	}

	// Encode the customer to JSON and send the response.
	if err := json.NewEncoder(res).Encode(c); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

// createHandler adds new customers to the DB.
func createHandler(res http.ResponseWriter, req *http.Request) {

	// Create a customer value.
	var c customer.Customer

	// Encode the customer document received into the customer value.
	err := json.NewDecoder(req.Body).Decode(&c)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// Save the customer in the DB.
	c.ID, err = customer.Save(c)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode the customer to JSON and send the response.
	b, err := json.Marshal(&c)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusCreated)
	res.Write(b)
}

func main() {

	// Start the http server to handle the request for
	// both versions of the API.
	log.Fatal(http.ListenAndServe(":3000", App()))
}
