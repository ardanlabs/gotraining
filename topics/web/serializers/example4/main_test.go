// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Tests for the sample program to show how to implement the
// xml.Marshaler interface to dictate the marshaling.
package main

import (
	"bytes"
	"encoding/xml"
	"strings"
	"testing"
)

func TestEncodeZeroValueUser(t *testing.T) {

	// Create a bytes buffer for our writer.
	var bb bytes.Buffer

	// Encode a zero value user and write the XML
	// to the bytes buffer.
	err := xml.NewEncoder(&bb).Encode(&User{})
	if err != nil {
		t.Fatal(err)
	}

	// Validate we received the expected response.
	got := strings.TrimSpace(bb.String())
	want := []string{
		`<User><first_name></first_name><CreatedAt>0001-01-01T00:00:00Z</CreatedAt><Admin>false</Admin><Bio></Bio></User>`,
		`<User><Bio></Bio><first_name></first_name><CreatedAt>0001-01-01T00:00:00Z</CreatedAt><Admin>false</Admin></User>`,
		`<User><Admin>false</Admin><Bio></Bio><first_name></first_name><CreatedAt>0001-01-01T00:00:00Z</CreatedAt></User>`,
		`<User><CreatedAt>0001-01-01T00:00:00Z</CreatedAt><Admin>false</Admin><Bio></Bio><first_name></first_name></User>`,
	}

	var found bool
	for _, w := range want {
		if got == w {
			found = true
			break
		}
	}

	if !found {
		t.Log("Wanted:", want)
		t.Log("Got   :", got)
		t.Fatal("Mismatch")
	}
}

func TestEncodeUser(t *testing.T) {

	// Create a bytes buffer for our writer.
	var bb bytes.Buffer

	// Create a string variable so we can take its address.
	bio := "An Awesome Coder!"

	// Create a user value for Mary Jane.
	u := User{
		FirstName: "Mary",
		LastName:  "Jane",
		Bio:       &bio,
	}

	// Encode the user and write the XML to the bytes buffer.
	err := xml.NewEncoder(&bb).Encode(&u)
	if err != nil {
		t.Fatal(err)
	}

	// Validate we received the expected response.
	got := strings.TrimSpace(bb.String())
	want := []string{
		`<User><first_name>Mary</first_name><CreatedAt>0001-01-01T00:00:00Z</CreatedAt><Admin>false</Admin><Bio>An Awesome Coder!</Bio><LastName>Jane</LastName></User>`,
		`<User><Admin>false</Admin><Bio>An Awesome Coder!</Bio><LastName>Jane</LastName><first_name>Mary</first_name><CreatedAt>0001-01-01T00:00:00Z</CreatedAt></User>`,
		`<User><CreatedAt>0001-01-01T00:00:00Z</CreatedAt><Admin>false</Admin><Bio>An Awesome Coder!</Bio><LastName>Jane</LastName><first_name>Mary</first_name></User>`,
		`<User><Bio>An Awesome Coder!</Bio><LastName>Jane</LastName><first_name>Mary</first_name><CreatedAt>0001-01-01T00:00:00Z</CreatedAt><Admin>false</Admin></User>`,
	}

	var found bool
	for _, w := range want {
		if got == w {
			found = true
			break
		}
	}

	if !found {
		t.Log("Wanted:", want)
		t.Log("Got   :", got)
		t.Fatal("Mismatch")
	}
}
