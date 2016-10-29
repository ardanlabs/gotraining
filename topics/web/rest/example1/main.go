// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to use a regex to handle REST based
// URL schemas and routes.
package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/ardanlabs/gotraining/topics/web/customer"
)

// App handles the routing of all the incoming customer
// requests into the server.
func App() http.Handler {

	// The regex allows us to match `/customers/:id`
	rx := regexp.MustCompile(`^/([^/]+)/?(\d*)$`)

	h := func(res http.ResponseWriter, req *http.Request) {
		log.Printf("[%s] %s\n", req.Method, req.URL.Path)

		// Validate we have a customer url.
		m := rx.FindAllStringSubmatch(req.URL.Path, -1)
		if len(m) == 0 {

			// Redirect to `/customers` if there isn't a match.
			http.Redirect(res, req, "/customers", http.StatusPermanentRedirect)
			return
		}

		// Extract the id portion of the customer url, `/customers/:id`.
		idStr := m[0][2]

		switch {
		case req.Method == "GET":

			// Show the content of the requested customer.
			if idStr != "" {
				showHandler(res, req, idStr)
				return
			}

			// Show the base index page.
			indexHandler(res, req)
			return

		case req.Method == "POST":

			// Show the create customer page.
			createHandler(res, req)
			return

		// The request is not formatted properly.
		default:

			// The request does not conform to what we expect.
			err := errors.New("invalid")
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	return http.HandlerFunc(h)
}

// indexHandler returns the entire list of customers in the DB.
func indexHandler(res http.ResponseWriter, req *http.Request) {

	// Retrieve the list of customers and render the document.
	err := customer.T.ExecuteTemplate(res, "index.html", customer.All())
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

// createHandler adds new customers to the DB.
func createHandler(res http.ResponseWriter, req *http.Request) {

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

// showHandler provides information about the specified customer.
func showHandler(res http.ResponseWriter, req *http.Request, idStr string) {

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

func main() {

	// Start the http server to handle the request for
	// both versions of the API.
	log.Fatal(http.ListenAndServe(":3000", App()))
}
