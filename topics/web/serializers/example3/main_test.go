// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Tests for the sample program to show how to use the XML encoder.
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
	want := `<User><first_name></first_name><CreatedAt>0001-01-01T00:00:00Z</CreatedAt><Admin>false</Admin></User>`
	if got != want {
		t.Log("Wanted:", want)
		t.Log("Got   :", got)
		t.Fatal("Mismatch")
	}
}

func TestEncodeUser(t *testing.T) {

	// Create a bytes buffer for our writer.
	var bb bytes.Buffer

	// Create a user value for Mary Jane.
	u := User{
		FirstName: "Mary",
		LastName:  "Jane",
	}

	// Encode the user and write the XML to the bytes buffer.
	err := xml.NewEncoder(&bb).Encode(&u)
	if err != nil {
		t.Fatal(err)
	}

	// Validate we received the expected response.
	got := strings.TrimSpace(bb.String())
	want := `<User><first_name>Mary</first_name><LastName>Jane</LastName><CreatedAt>0001-01-01T00:00:00Z</CreatedAt><Admin>false</Admin></User>`
	if got != want {
		t.Log("Wanted:", want)
		t.Log("Got   :", got)
		t.Fatal("Mismatch")
	}
}
