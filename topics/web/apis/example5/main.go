package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func App() http.Handler {
	r := echo.New()
	r.Use(middleware.Logger())
	r.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			ctx.Request().Header().Set("Content-Type", "application/json")
			return h(ctx)
		}
	})

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

	st := standard.New("")
	st.SetHandler(r)
	return st
}

func indexHandler(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, Customers.All())
}

func showHandler(ctx echo.Context) error {
	id := ctx.Param("id")
	c, err := Customers.Find(id)
	if err != nil {
		ctx.Error(err)
		return err
	}
	return ctx.JSON(http.StatusOK, c)
}

func createHandler(ctx echo.Context) error {
	c := &Customer{}
	err := ctx.Bind(&c)
	if err != nil {
		ctx.Error(err)
		return err
	}

	Customers.Save(c)
	return ctx.JSON(http.StatusCreated, c)
}

func main() {
	log.Fatal(http.ListenAndServe(":3000", App()))
}
