package main

import (
	"bytes"
	"strings"
	"testing"
)

func Test_EncodeUser_Blank(t *testing.T) {
	bb := &bytes.Buffer{}
	EncodeUser(bb, User{})

	act := strings.TrimSpace(bb.String())
	exp := `{"first_name":"","CreatedAt":"0001-01-01T00:00:00Z","Admin":false,"Bio":null}`

	if act != exp {
		t.Fatalf("expected %s to equal '%s'", exp, act)
	}
}

func Test_EncodeUser_WithData(t *testing.T) {
	bb := &bytes.Buffer{}
	EncodeUser(bb, User{FirstName: "Mary", LastName: "Jane"})

	act := strings.TrimSpace(bb.String())
	exp := `{"first_name":"Mary","LastName":"Jane","CreatedAt":"0001-01-01T00:00:00Z","Admin":false,"Bio":null}`

	if act != exp {
		t.Fatalf("expected %s to equal '%s'", exp, act)
	}
}
