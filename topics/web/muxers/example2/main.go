package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func indexHandler(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := templates.ExecuteTemplate(res, "index.html", Customers)
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
	err = templates.ExecuteTemplate(res, "show.html", c)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func createHandler(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
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
	r := httprouter.New()

	// Order matters
	r.GET("/customers/:id", showHandler)
	r.GET("/customers", indexHandler)
	r.POST("/customers", createHandler)

	// TODO: EXERCISE: Implement the PUT and PATCH response by accepting a
	// "name" form value, assigning it to the customer, saving it back
	// to the database, and then redirecting to the customer show page.
	// r.POST("/customers/:id", updateHandler)

	// TODO: EXERCISE: Implement the DELETE response by removing the
	// customer from the database and then redirecting back to the index page.
	// r.DELETE("/customers/:id", deleteHandler)

	r.GET("/", indexHandler)
	return r
}

func main() {
	log.Fatal(http.ListenAndServe(":3000", App()))
}
