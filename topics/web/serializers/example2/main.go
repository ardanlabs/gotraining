package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	Age       int
	CreatedAt time.Time
	Admin     bool
	Bio       *string
}

func (u User) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"first_name": u.FirstName,
		"CreatedAt":  u.CreatedAt,
		"Admin":      u.Admin,
		"Bio":        u.Bio,
	}
	if u.LastName != "" {
		m["LastName"] = u.LastName
	}

	return json.Marshal(m)
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
