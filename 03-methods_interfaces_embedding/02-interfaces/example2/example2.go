// http://play.golang.org/p/TEK2rfDrNx

// Sample program to show how to use an interface in Go. In this case,
// a pointer is used to support the interface call.
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
	// Create two values of type User.
	user1 := User{"Bill", "bill@email.com"}
	user2 := &User{"Jill", "jill@email.com"}

	// Pass a pointer of the values to support the interface.
	sendNotification(&user1)
	sendNotification(user2)
}

// sendNotification accepts values that implement the Notifier
// interface and sends notifications.
func sendNotification(notify Notifier) {
	notify.Notify()
}
