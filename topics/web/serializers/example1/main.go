// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to encode and decode JSON.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

// User represents a user in the system.
type User struct {
	FirstName string
	LastName  string
}

func useStreams() {

	// Create a user value to encode as JSON
	u := User{
		FirstName: "Kell",
		LastName:  "Walker",
	}

	// Encode using an Encoder. The relevant signatures are
	//	func NewEncoder(w io.Writer) *Encoder
	//	func (enc *Encoder) Encode(v interface{}) error
	if err := json.NewEncoder(os.Stdout).Encode(u); err != nil {
		log.Fatal(err)
	}

	// Prints to stdout
	// {"FirstName":"Kell","LastName":"Walker"}

	// Create an io.Reader that provides some input
	input := strings.NewReader(`{"FirstName": "Eleanor", "LastName": "O'Shea"}`)

	// Decode input using a Decoder. The relevant signatures are
	//	func NewDecoder(r io.Reader) *Decoder
	//	func (dec *Decoder) Decode(v interface{}) error
	if err := json.NewDecoder(input).Decode(&u); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("User is now %s %s\n", u.FirstName, u.LastName)
}

func useBytes() {

	// Create a user value to encode as JSON
	u := User{
		FirstName: "Carter",
		LastName:  "Walker",
	}

	// Marshal to []byte then print them
	//	func Marshal(v interface{}) ([]byte, error)
	b, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

	// Prints to stdout
	// {"FirstName":"Carter","LastName":"Walker"}

	input := []byte(`{"FirstName": "Samuel", "LastName": "O'Shea"}`)

	// Decode input using Unmarshal. The relevant signature is
	//	func Unmarshal(data []byte, v interface{}) error
	if err := json.Unmarshal(input, &u); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("User is now %s %s\n", u.FirstName, u.LastName)
}

func main() {
	useStreams()
	fmt.Print("\n")
	useBytes()
}
