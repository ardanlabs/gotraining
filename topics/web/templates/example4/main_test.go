package main

import (
	"bytes"
	"strings"
	"testing"
)

func Test_Exec(t *testing.T) {
	bb := &bytes.Buffer{}
	err := Exec(bb)
	if err != nil {
		t.Fatal(err)
	}

	act := bb.String()

	expectations := []string{
		"/foo?email=mark%40example.com",
		">mark@example.com</a>",
		`window.user = {"name":"Mark","email":"mark@example.com"};`,
	}

	for _, exp := range expectations {
		if !strings.Contains(act, exp) {
			t.Fatalf("expected %s to contain %s", act, exp)
		}
	}
}
