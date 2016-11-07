// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to implement the xml.Marshaler interface
// to dictate the marshaling.
package main

import (
	"encoding/xml"
	"log"
	"os"
	"strconv"
	"time"
)

// User represents a user in the system.
type User struct {
	FirstName string
	LastName  string
	Age       int
	CreatedAt time.Time
	Admin     bool
	Bio       *string
}

// MarshalXML implements the xml.Marshaler interface so we
// can dictate how the user is marshaled.
func (u *User) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	// Call into MarshalText directly for the time value.
	ca, err := u.CreatedAt.MarshalText()
	if err != nil {
		return err
	}

	// Create a document of key/value pairs for each field
	m := map[string][]byte{
		"first_name": []byte(u.FirstName),
		"CreatedAt":  ca,
		"Admin":      []byte(strconv.FormatBool(u.Admin)),
		"Bio":        nil,
	}

	// Add the Bio value if we have one.
	if u.Bio != nil {
		m["Bio"] = []byte(*u.Bio)
	}

	// Omit the last name from the document unless
	// we have a value.
	if u.LastName != "" {
		m["LastName"] = []byte(u.LastName)
	}

	// Create a slice of tokens starting with the element
	// provided with the call.
	tokens := []xml.Token{start}

	// Range of the key/value pairs creating XML elements.
	for key, value := range m {

		// Declare the starting element.
		se := xml.StartElement{
			Name: xml.Name{
				Space: "",
				Local: key,
			},
		}

		// Declare the ending element.
		ee := xml.EndElement{
			Name: se.Name,
		}

		// Append the starting element, the value and the ending element.
		tokens = append(tokens, se, xml.CharData(value), ee)
	}

	// Append the final ending element with the slice.
	tokens = append(tokens, xml.EndElement{Name: start.Name})

	// Range over the tokens.
	for _, t := range tokens {

		// Encode each token into the encoder.
		if err := e.EncodeToken(t); err != nil {
			return err
		}
	}

	// Flush the encoder to ensure tokens are written.
	if err := e.Flush(); err != nil {
		return err
	}

	return nil
}

func main() {

	// Encode a zero valued version of a user and write to stdout.
	err := xml.NewEncoder(os.Stdout).Encode(&User{})
	if err != nil {
		log.Fatal(err)
	}

	// Create a string variable so we can take its address.
	bio := "An Awesome Coder!"

	// Create a user value for Mary Jane.
	u := User{
		FirstName: "Mary",
		LastName:  "Jane",
		Bio:       &bio,
	}

	// Encode the user value and write to stdout.
	err = xml.NewEncoder(os.Stdout).Encode(&u)
	if err != nil {
		log.Fatal(err)
	}
}
