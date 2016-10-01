package main

import (
	"encoding/json"
	"fmt"
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

func main() {
	e := json.NewEncoder(os.Stdout)
	e.Encode(User{})

	fmt.Println("\n")
	e.Encode(User{FirstName: "Mary", LastName: "Jane"})
}
