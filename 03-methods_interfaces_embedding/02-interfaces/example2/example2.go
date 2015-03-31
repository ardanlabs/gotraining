// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// https://play.golang.org/p/tqy4dr2yYh

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

// notify implements the notifier interface with a value receiver.
func (u user) notify() {
	fmt.Printf("user: Sending user Email To %s<%s>\n",
		u.name,
		u.email)
}

// main is the entry point for the application.
func main() {
	// Create two values of type user.
	bill := user{"Bill", "bill@email.com"}
	jill := &user{"Jill", "jill@email.com"}

	// Values and pointers of type user implement the interface because value
	// receivers belong to the method sets of both values and pointers.

	// Pass a pointer of the values to support the interface.
	sendNotification(bill)
	sendNotification(jill)
}

// sendNotification accepts values that implement the notifier
// interface and sends notifications.
func sendNotification(notify notifier) {
	notify.notify()
}
