// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Tests for the sample program to show how to use create, parse and execute
// a template with simple data processing. This example uses a struct type
// value with a slice and method for generating HTML markup.
package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestExec(t *testing.T) {

	// Create a bytes Buffer for our Writer.
	var bb bytes.Buffer

	// Execute the template putting the output
	// into our bytes buffer.
	if err := Exec(&bb); err != nil {
		t.Fatal(err)
	}

	// Validate we received the correct version.
	got := strings.TrimSpace(bb.String())
	want := strings.TrimSpace(`
<h1>Mary Smith</h1>
<h2>MARY SMITH</h2>

Aliases:
<ul>
	<li>Scarface</li>
	<li>MC Skat Kat</li>
	
</ul>`)

	// NOTE: The above test string has a TAB on line 36
	//       to match the HTML string being produced.

	if got != want {
		t.Logf("Wanted: %v", want)
		t.Logf("Got   : %v", got)
		t.Fatal("Mismatch")
	}
}
