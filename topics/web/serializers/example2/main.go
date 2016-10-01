package main

import (
	"encoding/json"
	"fmt"
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

func main() {
	e := json.NewEncoder(os.Stdout)
	e.Encode(User{})

	fmt.Println("\n")
	e.Encode(User{FirstName: "Mary", LastName: "Jane"})
}
