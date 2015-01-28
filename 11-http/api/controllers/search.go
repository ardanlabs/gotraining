// Package controllers contains the controller logic for processing requests.
package controllers

import (
	"log"

	"github.com/ArdanStudios/gotraining/11-http/api/context"
)

// Search processes the search api.
func Search(c *context.Context) {
	log.Println(c.SessionID, ": controllers : Search : Started")

	d := struct {
		Name string
		Age  int
	}{
		Name: "Bill",
		Age:  45,
	}

	c.ServeJSON(d)

	log.Println(c.SessionID, ": controllers : Search : Completed")
}
