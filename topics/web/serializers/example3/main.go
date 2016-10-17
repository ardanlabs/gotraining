package main

import (
	"encoding/xml"
	"io"
	"log"
	"os"
	"time"
)

type User struct {
	FirstName string `xml:"first_name"`
	LastName  string `xml:",omitempty"`
	Age       int    `xml:"-"`
	CreatedAt time.Time
	Admin     bool
	Bio       *string
}

func EncodeUser(w io.Writer, u User) error {
	e := xml.NewEncoder(w)
	return e.Encode(User{})
}

func main() {
	err := EncodeUser(os.Stdout, User{})
	if err != nil {
		log.Fatal(err)
	}

	err = EncodeUser(os.Stdout, User{FirstName: "Mary", LastName: "Jane"})
	if err != nil {
		log.Fatal(err)
	}
}
