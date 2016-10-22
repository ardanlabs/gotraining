package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func App() http.Handler {
	r := echo.New()
	r.Use(middleware.Logger())

	r.SetRenderer(templates)

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

	st := standard.New("")
	st.SetHandler(r)
	return st
}

func indexHandler(ctx echo.Context) error {
	return ctx.Render(http.StatusOK, "index.html", Customers)
}

func showHandler(ctx echo.Context) error {
	id := ctx.Param("id")
	c, err := Customers.Find(id)
	if err != nil {
		ctx.Error(err)
		return err
	}
	return ctx.Render(http.StatusOK, "show.html", c)
}

func createHandler(ctx echo.Context) error {
	c := &Customer{}
	err := ctx.Bind(c)
	if err != nil {
		ctx.Error(err)
		return err
	}

	Customers.Save(c)
	return ctx.Redirect(http.StatusSeeOther, fmt.Sprintf("/customers/%s", c.ID))
}

func main() {
	log.Fatal(http.ListenAndServe(":3000", App()))
}
