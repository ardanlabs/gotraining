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

// BeforeAfterRequest is middleware for setting up context prior to the request.
type BeforeAfterRequest struct{}

// BeforeRequest handles the setup of processing the request.
func (b BeforeAfterRequest) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	b.before()
	next(w, r)
	b.after()
}

// before is executed prior to the route begin executed.
func (b BeforeAfterRequest) before() {
	log.Printf("controllers : before : Started")

	log.Println("controllers : before : Completed")
}

// before is executed after to the route is executed.
func (b BeforeAfterRequest) after() {
	log.Printf("controllers : after : Started")

	log.Println("controllers : after : Completed")
}
