package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/pat"
)

func indexHandler(res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode(Customers.All())
}

func showHandler(res http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get(":id")
	c, err := Customers.Find(id)
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(res).Encode(c)
}

func createHandler(res http.ResponseWriter, req *http.Request) {
	c := &Customer{}
	err := json.NewDecoder(req.Body).Decode(c)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	Customers.Save(c)
	json.NewEncoder(res).Encode(c)
}

func App() http.Handler {
	r := pat.New()

	// Order matters
	r.Get("/customers/{id}", showHandler)
	r.Get("/customers", indexHandler)
	r.Post("/customers", createHandler)

	// TODO: EXERCISE: Implement the PUT and PATCH response by accepting a
	// "name" form value, assigning it to the customer, saving it back
	// to the database, and then rendering the customer JSON.
	// r.Post("/customers/{id}", updateHandler)

	// TODO: EXERCISE: Implement the DELETE response by removing the
	// customer from the database.
	// r.Delete("/customers/{id}", deleteHandler)

	r.Get("/", indexHandler)

	return r
}

func main() {
	log.Fatal(http.ListenAndServe(":3000", App()))
}
