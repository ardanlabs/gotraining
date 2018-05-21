// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show common JSON mistakes.
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// User represents a user in our system. email is not exported so it won't be
// marshaled.
type User struct {
	Name  string   `json:"name"`
	email string   `json:"email"`
	Roles []string `json:"roles"`
}

func main() {

	// Marshal a zero value User
	b, err := json.Marshal(User{})
	if err != nil {
		log.Fatal(err)
	}

	printMsgData("Zero value User", b)
	printMsgData("Note 'roles' is null", nil)

	// Initialize roles for an otherwise zeroed User
	u := User{
		Roles: []string{},
	}

	b, err = json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	printMsgData("User with empty roles slice", b)
	printMsgData("Note 'roles' is [] not null", nil)

	// Fill in data for user
	u = User{
		Name:  "Alice",
		email: "alice@example.com",
		Roles: []string{"admin", "agent"},
	}

	b, err = json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	printMsgData("User with data", b)

	fmt.Println("\nNote in all examples 'email' is missing.")
}

func printMsgData(msg string, data []byte) {
	fmt.Printf("%30s | %s\n", msg, string(data))
}
