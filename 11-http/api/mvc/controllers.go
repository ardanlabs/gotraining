// Package controllers contains the controller logic for processing requests.
package mvc

import (
	"log"
)

// GetUser returns the specified user.
func GetUser(c *Context) {
	log.Println(c.SessionID, ": mvc : GetUser : Started")

	d := struct {
		Name string
		Age  int
	}{
		Name: "Bill",
		Age:  45,
	}

	c.ServeJSON(d)

	log.Println(c.SessionID, ": mvc : GetUser : Completed")
}

// InsUser returns the specified user.
func InsUser(c *Context) {
	log.Println(c.SessionID, ": mvc : InsUser : Started")

	log.Println(c.SessionID, ": mvc : InsUser : Completed")
}

// UpdUser returns the specified user.
func UpdUser(c *Context) {
	log.Println(c.SessionID, ": mvc : UpdUser : Started")

	log.Println(c.SessionID, ": mvc : UpdUser : Completed")
}

// DelUser returns the specified user.
func DelUser(c *Context) {
	log.Println(c.SessionID, ": mvc : DelUser : Started")

	log.Println(c.SessionID, ": mvc : DelUser : Completed")
}
