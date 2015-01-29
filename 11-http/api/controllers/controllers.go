// Package controllers contains the controller logic for processing requests.
package controllers

import (
	"log"

	"github.com/ArdanStudios/gotraining/11-http/api/context"
)

// GetUser returns the specified user.
func GetUser(c *context.Context) {
	log.Println(c.SessionID, ": controllers : GetUser : Started")

	d := struct {
		Name string
		Age  int
	}{
		Name: "Bill",
		Age:  45,
	}

	c.ServeJSON(d)

	log.Println(c.SessionID, ": controllers : GetUser : Completed")
}

// InsUser returns the specified user.
func InsUser(c *context.Context) {
	log.Println(c.SessionID, ": controllers : InsUser : Started")

	log.Println(c.SessionID, ": controllers : InsUser : Completed")
}

// UpdUser returns the specified user.
func UpdUser(c *context.Context) {
	log.Println(c.SessionID, ": controllers : UpdUser : Started")

	log.Println(c.SessionID, ": controllers : UpdUser : Completed")
}

// DelUser returns the specified user.
func DelUser(c *context.Context) {
	log.Println(c.SessionID, ": controllers : DelUser : Started")

	log.Println(c.SessionID, ": controllers : DelUser : Completed")
}
