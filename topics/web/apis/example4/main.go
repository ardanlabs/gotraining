// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to create a basic CRUD based web api for
// customers with a middleware component.
package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ardanlabs/gotraining/topics/web/customer"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

// App loads the entire API set together for use.
func App() http.Handler {

	// Create an echo mux to handle routes and middleware.
	r := echo.New()

	// Add in the middleware for logging.
	r.Use(middleware.Logger())

	// Add in a custom middleware for setting the
	// context type in the request for later.
	r.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			ctx.Request().Header().Set("Content-Type", "application/json")
			return h(ctx)
		}
	})

	// Define the routes and order matters.
	r.GET("/customers/:id", showHandler)
	r.GET("/customers", indexHandler)
	r.POST("/customers", createHandler)

	// Redirect requests from `/`` to `/customers`.
	r.GET("/", func(ctx echo.Context) error {
		return ctx.Redirect(http.StatusMovedPermanently, "/customers")
	})

	// Create a standard echo server and bind
	// the echo mux as the handler.
	st := standard.New("")
	st.SetHandler(r)

	return st
}

// indexHandler returns the entire list of customers in the DB.
func indexHandler(ctx echo.Context) error {

	// Retrieve the list of customers, encode to JSON
	// and send the response.
	return ctx.JSON(http.StatusOK, customer.All())
}

// showHandler returns a single specified customer.
func showHandler(ctx echo.Context) error {

	// Retrieve the customer id from the request.
	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Retreive that customer from the DB.
	c, err := customer.Find(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	// Encode the customer to JSON and send the response.
	return ctx.JSON(http.StatusOK, c)
}

// createHandler adds new customers to the DB.
func createHandler(ctx echo.Context) error {

	// Create a customer value.
	var c customer.Customer

	// Encode the customer document received into the customer value.
	err := ctx.Bind(&c)
	if err != nil {
		ctx.Error(err)
		return err
	}

	// Save the customer in the DB.
	c.ID, err = customer.Save(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	// Encode the customer to JSON and send the response.
	return ctx.JSON(http.StatusOK, c)
}

func main() {

	// Start the http server to handle the request for
	// both versions of the API.
	log.Fatal(http.ListenAndServe(":3000", App()))
}
