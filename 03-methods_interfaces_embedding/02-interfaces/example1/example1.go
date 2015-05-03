// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// https://play.golang.org/p/eY2Ms-UF-t

// Sample program to show how to use an interface in Go.
package main

import (
	"fmt"
)

// notifier is an interface that defines notification
// type behavior.
type notifier interface {
	notify()
}

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// notify implements the notifier interface with a pointer receiver.
func (u *user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}

// main is the entry point for the application.
func main() {
	// Create a value of type User and send a notification.
	u := user{"Bill", "bill@email.com"}

	// Values of type user do not implement the interface because pointer
	// receivers don't belong to the method set of a value.

	sendNotification(u)

	// ./example1.go:40: cannot use u (type user) as type notifier in argument to sendNotification:
	//   user does not implement notifier (notify method has pointer receiver)
}

// sendNotification accepts values that implement the notifier
// interface and sends notifications.
func sendNotification(n notifier) {
	n.notify()
}
