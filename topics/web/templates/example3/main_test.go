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

	exp := "Hello, Mark!"
	act := strings.TrimSpace(bb.String())

	if exp != act {
		t.Fatalf("expected %s, got %s", exp, act)
	}
}
