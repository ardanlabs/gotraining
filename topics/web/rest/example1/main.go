package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
)

type router struct{}

func (r router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// Log the request
	log.Printf("[%s] %s\n", req.Method, req.URL.Path)

	// Requests should match /customers/:id
	rx := regexp.MustCompile(`^/([^/]+)/?(\d*)$`)
	m := rx.FindAllStringSubmatch(req.URL.Path, -1)
	if len(m) == 0 {
		// Redirect to /customers if there isn't a match
		http.Redirect(res, req, "/customers", http.StatusPermanentRedirect)
		return
	}

	id := m[0][2]

	// If the request is a GET and there is no ID we want to
	// render the index template.
	if req.Method == "GET" && id == "" {
		indexHandler(res, req)
		return
	}

	// If the request is a POST we want to create a new customer
	if req.Method == "POST" {
		createHandler(res, req)
		return
	}

	// The rest of the actions all work on a specific customer.
	// Find that customer in the database. If that customer does
	// not exist, then return a 404 and stop processing the request.
	c, err := Customers.Find(id)
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
		return
	}

	// Switch on the request method (HTTP verb) and handle the verb
	// appropriately.
	switch req.Method {
	case "GET":
		// Render the show.html template to display the customer.
		err := templates.ExecuteTemplate(res, "show.html", c)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
		return
	case "PUT", "PATCH":
		// TODO: EXERCISE: Implement the PUT and PATCH response by accepting a
		// "name" form value, assigning it to the customer, saving it back
		// to the database, and then redirecting to the customer show page.
		return
	case "DELETE":
		// TODO: EXERCISE: Implement the DELETE response by removing the
		// customer from the database and then redirecting back to the index page.
		return
	default:
		http.Error(res, fmt.Sprintf("unable to handle the HTTP method: %s", req.Method), http.StatusBadRequest)
		return
	}
}

func App() http.Handler {
	return router{}
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	err := templates.ExecuteTemplate(res, "index.html", Customers)
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

func main() {
	log.Fatal(http.ListenAndServe(":3000", App()))
}
