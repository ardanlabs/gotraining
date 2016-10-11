package main

import (
	"encoding/json"
	"io"
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

func EncodeUser(w io.Writer, u User) {
	e := json.NewEncoder(w)
	e.Encode(u)
}

func main() {
	w := os.Stdout
	EncodeUser(w, User{})

	EncodeUser(w, User{FirstName: "Mary", LastName: "Jane"})
}
