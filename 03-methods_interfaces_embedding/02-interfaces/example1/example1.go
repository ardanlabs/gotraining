// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/h5Q9dQgnzS

// Sample program to show how to use an interface in Go.
package main

import (
	"fmt"
)

type (
	// notifier is an interface that defined notification
	// type behavior.
	notifier interface {
		notify()
	}

	// user defines a user in the program.
	user struct {
		name  string
		email string
	}
)

// notify implements a method with a pointer receiver.
func (u *user) notify() {
	fmt.Printf("User: Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}

// main is the entry point for the application.
func main() {
	// Create a value of type User and send a notification.
	user1 := user{"Bill", "bill@email.com"}

	// Values of type user do not implement the interface because pointer
	// receivers don't belong to the method set of a value.

	sendNotification(user1)

	// ./example2.go:37: cannot use user1 (type User) as type notifier in argument to sendNotification:
	//   User does not implement notifier (notify method has pointer receiver)
}

// sendNotification accepts values that implement the notifier
// interface and sends notifications.
func sendNotification(notify notifier) {
	notify.notify()
}
