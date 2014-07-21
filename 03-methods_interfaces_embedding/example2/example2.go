// Sample program to show how to use an interface in Go. In this case,
// Go will not deference a value to support the interface.
package main

import (
	"fmt"
)

type (
	// Notifier is an interface that defined notification
	// type behavior.
	Notifier interface {
		Notify()
	}

	// User defines a user in the program.
	User struct {
		Name  string
		Email string
	}
)

// Notify implements a method that can be called via
// a value of type User.
func (u *User) Notify() {
	fmt.Printf("User: Sending User Email To %s<%s>\n",
		u.Name,
		u.Email)
}

// main is the entry point for the application.
func main() {
	// Create a value of type User and send a notification.
	user1 := User{"Bill", "bill@email.com"}

	// ./example2.go:38: cannot use user1 (type User) as type Notifier in argument to sendNotification:
	//   User does not implement Notifier (Notify method has pointer receiver)
	sendNotification(user1)
}

// sendNotification accepts values that implement the Notifier
// interface and sends notifications.
func sendNotification(notify Notifier) {
	notify.Notify()
}
