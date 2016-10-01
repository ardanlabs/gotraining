package main

import (
	"encoding/xml"
	"fmt"
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

func main() {
	e := xml.NewEncoder(os.Stdout)
	e.Encode(User{})
	fmt.Println("\n")
	e.Encode(User{FirstName: "Mary", LastName: "Jane"})
}
