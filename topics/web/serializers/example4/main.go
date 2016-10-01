package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
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

func (u User) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	ca, _ := u.CreatedAt.MarshalText()
	m := map[string][]byte{
		"first_name": []byte(u.FirstName),
		"CreatedAt":  ca,
		"Admin":      []byte(strconv.FormatBool(u.Admin)),
		"Bio":        nil,
	}
	if u.Bio != nil {
		m["Bio"] = []byte(*u.Bio)
	}
	if u.LastName != "" {
		m["LastName"] = []byte(u.LastName)
	}

	tokens := []xml.Token{start}

	for key, value := range m {
		t := xml.StartElement{Name: xml.Name{Space: "", Local: key}}
		tokens = append(tokens, t, xml.CharData(value), xml.EndElement{t.Name})
	}

	tokens = append(tokens, xml.EndElement{Name: start.Name})

	for _, t := range tokens {
		err := e.EncodeToken(t)
		if err != nil {
			return err
		}
	}

	// flush to ensure tokens are written
	err := e.Flush()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	e := xml.NewEncoder(os.Stdout)
	e.Encode(User{})

	fmt.Println("\n")
	bio := "An Awesome Coder!"
	e.Encode(User{FirstName: "Mary", LastName: "Jane", Bio: &bio})
}
