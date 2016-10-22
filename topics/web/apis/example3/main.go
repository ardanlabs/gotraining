// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// TODO: EXERCISE: Implement the PUT and PATCH response by accepting a
// "name" form value, assigning it to the customer, saving it back
// to the database, and then rendering the customer JSON.
// r.Post("/customers/{id}", updateHandler)

// TODO: EXERCISE: Implement the DELETE response by removing the
// customer from the database.
// r.Delete("/customers/{id}", deleteHandler)

// Sample program to show how to create a basic CRUD based
// web api for customers.
package main

import (
	"encoding/json"
	"log"
	"net/http"

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

	// Define the root route.
	r.Get("/", indexHandler)

	return r
}

// indexHandler returns the entire list of customers in the DB.
func indexHandler(res http.ResponseWriter, req *http.Request) {

	// Retrieve the list of customers, encode to JSON
	// and send the response.
	err := json.NewEncoder(res).Encode(DB.AllCustomers())
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

// showHandler returns a single specified customer.
func showHandler(res http.ResponseWriter, req *http.Request) {

	// Retrieve the customer id from the request.
	id := req.URL.Query().Get(":id")

	// Retreive that customer from the DB.
	c, err := DB.FindCustomer(id)
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
		return
	}

	// Encode the customer to JSON and send the response.
	err = json.NewEncoder(res).Encode(c)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

// createHandler adds new customers to the DB.
func createHandler(res http.ResponseWriter, req *http.Request) {

	// Create a customer value.
	var c Customer

	// Encode the customer document received into the customer value.
	err := json.NewDecoder(req.Body).Decode(&c)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// Save the customer in the DB.
	DB.SaveCustomer(c)

	// Encode the customer to JSON and send the response.
	err = json.NewEncoder(res).Encode(&c)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func main() {

	// Start the http server to handle the request for
	// both versions of the API.
	log.Fatal(http.ListenAndServe(":3000", App()))
}
