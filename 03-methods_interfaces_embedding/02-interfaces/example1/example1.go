// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/7q3zw-sVwn

// Sample program to show how to use an interface in Go. In this case,
// Go will NOT deference a pointer value to support the interface.
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

// notify implements a method that can be called via
// a value of type user.
func (u *user) notify() {
	fmt.Printf("User: Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}

// main is the entry point for the application.
func main() {
	// Create a value of type User and send a notification.
	user1 := user{"Bill", "bill@email.com"}

	sendNotification(user1)

	// ./example2.go:38: cannot use user1 (type User) as type notifier in argument to sendNotification:
	//   User does not implement notifier (notify method has pointer receiver)
}

// sendNotification accepts values that implement the notifier
// interface and sends notifications.
func sendNotification(notify notifier) {
	notify.notify()
}
