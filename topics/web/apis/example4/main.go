package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func App() http.Handler {
	r := httprouter.New()

	// Order matters
	r.GET("/customers/:id", showHandler)
	r.GET("/customers", indexHandler)
	r.POST("/customers", createHandler)

	// TODO: EXERCISE: Implement the PUT and PATCH response by accepting a
	// "name" form value, assigning it to the customer, saving it back
	// to the database, and then rendering the customer JSON.
	// r.POST("/customers/:id", updateHandler)

	// TODO: EXERCISE: Implement the DELETE response by removing the
	// customer from the database.
	// r.DELETE("/customers/:id", deleteHandler)

	r.GET("/", indexHandler)
	return r
}

func indexHandler(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := json.NewEncoder(res).Encode(Customers.All())
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func showHandler(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	c, err := Customers.Find(id)
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
		return
	}
	err = json.NewEncoder(res).Encode(c)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func createHandler(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	c := &Customer{}
	err := json.NewDecoder(req.Body).Decode(c)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	Customers.Save(c)
	err = json.NewEncoder(res).Encode(c)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	log.Fatal(http.ListenAndServe(":3000", App()))
}
