package main

import (
	"encoding/json"
	"io"
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

func EncodeUser(w io.Writer, u User) {
	e := json.NewEncoder(w)
	e.Encode(u)
}

func main() {
	w := os.Stdout
	EncodeUser(w, User{})
	EncodeUser(w, User{FirstName: "Mary", LastName: "Jane"})
}
