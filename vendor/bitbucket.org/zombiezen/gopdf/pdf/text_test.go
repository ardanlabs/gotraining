// Copyright (C) 2011, Ross Light

package pdf

import (
	"math"
	"testing"
)

const textExpectedOutput = `/Helvetica 12.00000 Tf
14.40000 TL
14.00000 TL
(Hello, World!) Tj
T*
(This is SPARTA!!1!) Tj
`

func TestText(t *testing.T) {
	text := new(Text)
	text.SetFont(Helvetica, 12)
	text.SetLeading(14)
	text.Text("Hello, World!")
	text.NextLine()
	text.Text("This is SPARTA!!1!")

	if text.buf.String() != textExpectedOutput {
		t.Errorf("Output was %q, expected %q", text.buf.String(), textExpectedOutput)
	}

	if len(text.fonts) == 1 {
		if !text.fonts[Helvetica] {
			t.Error("Helvetica missing from fonts")
		}
	} else {
		t.Errorf("Got %d fonts, expected %d", len(text.fonts), 1)
	}
}

func floatEq(a, b, epsilon float64) bool {
	return math.Abs(a-b) < epsilon
}

func TestTextX(t *testing.T) {
	text := new(Text)
	if text.X() != 0 {
		t.Errorf("Text does not start at X=0 (got %.5f)", text.X())
	}

	text.SetFont(Helvetica, 12)

	text.Text("Hello!")
	if !floatEq(float64(text.X()), 30.672, 1e-5) {
		t.Errorf("\"Hello!\" has wrong X (=%.5f) when %.5f is desired", text.X(), 30.672)
	}

	text.NextLine()
	if text.X() != 0 {
		t.Errorf("Performing NextLine does not reset X (got %.5f)", text.X())
	}

	text.Text("Hello World")
	if !floatEq(float64(text.X()), 62.004, 1e-5) {
		t.Errorf("\"Hello World\" has wrong X (=%.5f) when %.5f is desired", text.X(), 62.004)
	}

	text.NextLineOffset(41.23, 55.555)
	if !floatEq(float64(text.X()), 41.23, 1e-3) {
		t.Errorf("NextLineOffset has wrong X (=%.5f) when %.5f is desired", text.X(), 41.23)
	}
}

func TestTextY(t *testing.T) {
	text := new(Text)
	if text.Y() != 0 {
		t.Errorf("Text does not start at Y=0 (got %.5f)", text.Y())
	}

	text.SetFont(Helvetica, 12)
	text.Text("Hello!")
	if text.Y() != 0 {
		t.Errorf("\"Hello!\" changes baseline (got %.5f)", text.Y())
	}

	text.NextLine()
	if !floatEq(float64(text.Y()), -14.400, 1e-4) {
		t.Errorf("NextLine y = %.5f (expected %.5f)", text.Y(), -14.400)
	}

	text.SetLeading(41.23)
	text.NextLine()
	if !floatEq(float64(text.Y()), -55.630, 1e-4) {
		t.Errorf("NextLine does not respect leading, y = %.5f (expected %.5f)", text.Y(), -55.630)
	}

	text.NextLineOffset(1.0, 5.5)
	if !floatEq(float64(text.Y()), -50.130, 1e-4) {
		t.Errorf("NextLineOffset does not set Y correctly, y = %.5f (expected %.5f)", text.Y(), -50.130)
	}
}
