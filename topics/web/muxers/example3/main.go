// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to use the echo toolkit.
package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/ardanlabs/gotraining/topics/web/customer"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

// render contains a pointer to the templates.
type render struct {
	*template.Template
}

// Render allows us to implement the echo.Render interface.
func (r render) Render(w io.Writer, name string, data interface{}, ctx echo.Context) error {
	return r.ExecuteTemplate(w, name, data)
}

// App loads the entire API set together for use.
func App() http.Handler {

	// Create an echo router.
	r := echo.New()

	// Add the logging middleware into the route.
	r.Use(middleware.Logger())

	// Load the customer templates.
	r.SetRenderer(&render{customer.T})

	// Define the routes and order matters.
	r.GET("/customers/:id", showHandler)
	r.GET("/customers", indexHandler)
	r.POST("/customers", createHandler)

	r.GET("/", indexHandler)

	// Create an echo server binding the
	// echo router.
	st := standard.New("")
	st.SetHandler(r)

	return st
}

// indexHandler returns the entire list of customers in the DB.
func indexHandler(ctx echo.Context) error {

	// Retrieve the list of customers and render the document.
	return ctx.Render(http.StatusOK, "index.html", customer.All())
}

// showHandler provides information about the specified customer.
func showHandler(ctx echo.Context) error {

	// Capture the id from the request.
	idStr := ctx.Param("id")

	// Convert the id to an integer.
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.Error(err)
		return err
	}

	// Find that customer in the database. If that customer does
	// not exist, then return a 404 and stop processing the request.
	c, err := customer.Find(id)
	if err != nil {
		ctx.Error(err)
		return err
	}

	// Render the show.html template to display the customer.
	return ctx.Render(http.StatusOK, "show.html", c)
}

// createHandler adds new customers to the DB.
func createHandler(ctx echo.Context) error {

	// Create a customer to save.
	var c customer.Customer

	// Bind the customer against the request body to
	// set the provided name.
	if err := ctx.Bind(&c); err != nil {
		ctx.Error(err)
		return err
	}

	// Save the customer in the DB.
	var err error
	c.ID, err = customer.Save(c)
	if err != nil {
		ctx.Error(err)
		return err
	}

	// Redirect the user to the customer page.
	return ctx.Redirect(http.StatusSeeOther, fmt.Sprintf("/customers/%d", c.ID))
}

func main() {

	// Start the http server to handle the request for
	// both versions of the API.
	log.Fatal(http.ListenAndServe(":3000", App()))
}
