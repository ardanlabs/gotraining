// Package controllers contains the controller logic for processing requests.
package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// NotFound handles the 404 response.
func NotFound(w http.ResponseWriter, r *http.Request) {
	log.Println("controllers : NotFound : Started")

	http.NotFound(w, r)

	log.Println("controllers : NotFound : Completed")
}

// ServeError handles application errors
func ServeError(w http.ResponseWriter, err error, statusCode int) {
	log.Printf("controllers : ServeError : Started : Error[%s]\n", err)

	e := struct {
		Err string
	}{
		Err: err.Error(),
	}

	data, err := json.MarshalIndent(&e, "", "    ")
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	http.Error(w, string(data), statusCode)
	log.Println("controllers : ServeError : Completed")
}

// ServeJSON handles serving values as JSON.
func ServeJSON(w http.ResponseWriter, v interface{}) {
	log.Printf("controllers : ServeJSON : Started : Error[%+v]\n", v)

	data, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, string(data))
	log.Println("controllers : ServeJSON : Completed")
}
