// Package controllers contains the controller logic for processing requests.
package controllers

import (
	"log"
	"net/http"
)

// Authentication is middleware for authenticating each request.
type Authentication struct{}

// Authentication handles the authentication of each request.
func (a Authentication) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Printf("controllers : Authentication : Started : Route[%s]\n", r.URL.RequestURI())

	// ServeError(w, errors.New("Auth Error"), http.StatusUnauthorized)

	log.Println("controllers : Authentication : Completed")
	next(w, r)
}

// BeforeRequest is middleware for setting up context prior to the request.
type BeforeRequest struct{}

// BeforeRequest handles the setup of processing the request.
func (b BeforeRequest) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Printf("controllers : BeforeRequest : Started")

	log.Println("controllers : BeforeRequest : Completed")
	next(w, r)
}

// AfterRequest is middleware for cleaning up context after to the request.
type AfterRequest struct{}

// BeforeRequest handles the setup of processing the request.
func (a AfterRequest) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Printf("controllers : AfterRequest : Started")

	log.Println("controllers : AfterRequest : Completed")
	next(w, r)
}
