package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"time"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:",omitempty"`
	Age       int    `json:"-"`
	CreatedAt time.Time
	Admin     bool
	Bio       *string
}

func EncodeUser(w io.Writer, u User) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

func main() {
	w := os.Stdout
	err := EncodeUser(w, User{})
	if err != nil {
		log.Fatal(err)
	}
	err = EncodeUser(w, User{FirstName: "Mary", LastName: "Jane"})
	if err != nil {
		log.Fatal(err)
	}
}
