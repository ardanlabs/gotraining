// Package controllers contains the controller logic for processing requests.
package controllers

import (
	"log"
	"net/http"
)

// Search processes the search api.
func Search(w http.ResponseWriter, r *http.Request) {
	log.Println("controllers : Search : Started")

	d := struct {
		Name string
		Age  int
	}{
		Name: "Bill",
		Age:  45,
	}

	ServeJSON(w, d)

	log.Println("controllers : Search : Completed")
}
