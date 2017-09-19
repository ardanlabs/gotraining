// Copyright (C) 2011, Ross Light

package pdf

import (
	"testing"
)

const pathExpectedOutput = `12.00000 34.00000 m
-56.00000 78.00000 l
h
3.10000 -5.90000 21.10000 80.90000 re
`

func TestPath(t *testing.T) {
	path := new(Path)
	path.Move(Point{12, 34})
	path.Line(Point{-56, 78})
	path.Close()
	path.Rectangle(Rectangle{Point{3.1, -5.9}, Point{24.2, 75.0}})

	if path.buf.String() != pathExpectedOutput {
		t.Errorf("Output was %q, expected %q", path.buf.String(), pathExpectedOutput)
	}
}
