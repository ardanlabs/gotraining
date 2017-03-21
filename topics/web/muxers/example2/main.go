// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to use the httprouter router.
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ardanlabs/gotraining/topics/web/customer"
	"github.com/julienschmidt/httprouter"
)

// App loads the entire API set together for use.
func App() http.Handler {

	// Create a version of the pat router.
	r := httprouter.New()

	// Define the routes and order matters.
	r.GET("/customers/:id", showHandler)
	r.GET("/customers", indexHandler)
	r.POST("/customers", createHandler)

	r.GET("/", indexHandler)

	return r
}

// indexHandler returns the entire list of customers in the DB.
func indexHandler(res http.ResponseWriter, req *http.Request, p httprouter.Params) {

	// Retrieve the list of customers and render the document.
	err := customer.T.ExecuteTemplate(res, "index.html", customer.All())
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

// showHandler provides information about the specified customer.
func showHandler(res http.ResponseWriter, req *http.Request, params httprouter.Params) {

	// Capture the id from the request.
	idStr := params.ByName("id")

	// Convert the id to an integer.
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	// Find that customer in the database. If that customer does
	// not exist, then return a 404 and stop processing the request.
	c, err := customer.Find(id)
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
		return
	}

	// Render the show.html template to display the customer.
	if err := customer.T.ExecuteTemplate(res, "show.html", c); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

// createHandler adds new customers to the DB.
func createHandler(res http.ResponseWriter, req *http.Request, p httprouter.Params) {

	// Parse the raw query from the URL and update r.Form.
	if err := req.ParseForm(); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a customer to save.
	c := customer.Customer{
		Name: req.FormValue("name"),
	}

	// Save the customer in the DB.
	var err error
	c.ID, err = customer.Save(c)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirect the user to the customer page.
	http.Redirect(res, req, fmt.Sprintf("/customers/%d", c.ID), http.StatusSeeOther)
}

func main() {

	// Start the http server to handle the request for
	// both versions of the API.
	log.Fatal(http.ListenAndServe(":3000", App()))
}
