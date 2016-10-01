package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/pat"
)

func indexHandler(res http.ResponseWriter, req *http.Request) {
	err := templates.ExecuteTemplate(res, "index.html", Customers)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func showHandler(res http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get(":id")
	c, err := Customers.Find(id)
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
		return
	}
	err = templates.ExecuteTemplate(res, "show.html", c)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func createHandler(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	c := &Customer{Name: req.FormValue("name")}
	Customers.Save(c)
	http.Redirect(res, req, fmt.Sprintf("/customers/%s", c.ID), http.StatusSeeOther)
}

func App() http.Handler {
	r := pat.New()

	// Order matters
	r.Get("/customers/{id}", showHandler)
	r.Get("/customers", indexHandler)
	r.Post("/customers", createHandler)

	// TODO: EXERCISE: Implement the PUT and PATCH response by accepting a
	// "name" form value, assigning it to the customer, saving it back
	// to the database, and then redirecting to the customer show page.
	// r.Post("/customers/{id}", updateHandler)

	// TODO: EXERCISE: Implement the DELETE response by removing the
	// customer from the database and then redirecting back to the index page.
	// r.Delete("/customers/{id}", deleteHandler)

	r.Get("/", indexHandler)

	return r
}

func main() {
	log.Fatal(http.ListenAndServe(":3000", App()))
}
