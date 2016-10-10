package main

import (
	"bytes"
	"strings"
	"testing"
)

func Test_Exec(t *testing.T) {
	bb := &bytes.Buffer{}
	Exec(bb)

	act := bb.String()

	expectations := []string{
		"<h1>Mary Smith</h1>",
		"<li>Scarface</li>",
		"<li>MC Skat Kat</li>",
	}

	for _, exp := range expectations {
		if !strings.Contains(act, exp) {
			t.Fatalf("expected %s to contain %s", act, exp)
		}
	}
}
