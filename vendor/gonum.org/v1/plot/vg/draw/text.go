// Copyright ©2020 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package draw // import "gonum.org/v1/plot/vg/draw"

import (
	"image/color"
	"math"
	"strings"

	"gonum.org/v1/plot/vg"
)

// TextHandler parses, formats and renders text.
type TextHandler interface {
	// Box returns the bounding box of the given text where:
	//  - width is the horizontal space from the origin.
	//  - height is the vertical space above the baseline.
	//  - depth is the vertical space below the baseline, a negative number.
	Box(txt string, fnt vg.Font) (width, height, depth vg.Length)

	// Draw renders the given text with the provided style and position
	// on the canvas.
	Draw(c *Canvas, txt string, sty TextStyle, pt vg.Point)
}

// TextStyle describes what text will look like.
type TextStyle struct {
	// Color is the text color.
	Color color.Color

	// Font is the font description.
	Font vg.Font

	// Rotation is the text rotation in radians, performed around the axis
	// defined by XAlign and YAlign.
	Rotation float64

	// XAlign and YAlign specify the alignment of the text.
	XAlign XAlignment
	YAlign YAlignment

	// TextHandler parses and formats text according to a given
	// dialect (Markdown, LaTeX, plain, ...)
	// The default is a plain text handler.
	Handler TextHandler
}

func (sty TextStyle) handler() TextHandler {
	if sty.Handler == nil {
		return PlainTextHandler{}
	}
	return sty.Handler
}

// Width returns the width of lines of text
// when using the given font before any text rotation is applied.
func (sty TextStyle) Width(txt string) (max vg.Length) {
	txt = strings.TrimRight(txt, "\n")
	for _, line := range strings.Split(txt, "\n") {
		if w := sty.Font.Width(line); w > max {
			max = w
		}
	}
	return
}

// Height returns the height of the text when using
// the given font before any text rotation is applied.
func (sty TextStyle) Height(txt string) vg.Length {
	nl := sty.textNLines(txt)
	if nl == 0 {
		return vg.Length(0)
	}
	e := sty.Font.Extents()
	return e.Height*vg.Length(nl-1) + e.Ascent - e.Descent
}

// Rectangle returns a rectangle giving the bounds of
// this text assuming that it is drawn at (0, 0).
func (sty TextStyle) Rectangle(txt string) vg.Rectangle {
	w := sty.Width(txt)
	h := sty.Height(txt)
	xoff := vg.Length(sty.XAlign) * w
	yoff := vg.Length(sty.YAlign) * h
	// lower left corner
	p1 := rotatePoint(sty.Rotation, vg.Point{X: xoff, Y: yoff})
	// upper left corner
	p2 := rotatePoint(sty.Rotation, vg.Point{X: xoff, Y: h + yoff})
	// lower right corner
	p3 := rotatePoint(sty.Rotation, vg.Point{X: w + xoff, Y: yoff})
	// upper right corner
	p4 := rotatePoint(sty.Rotation, vg.Point{X: w + xoff, Y: h + yoff})

	return vg.Rectangle{
		Max: vg.Point{
			X: max(p1.X, p2.X, p3.X, p4.X),
			Y: max(p1.Y, p2.Y, p3.Y, p4.Y),
		},
		Min: vg.Point{
			X: min(p1.X, p2.X, p3.X, p4.X),
			Y: min(p1.Y, p2.Y, p3.Y, p4.Y),
		},
	}
}

// textNLines returns the number of lines in the text.
func (sty TextStyle) textNLines(txt string) int {
	txt = strings.TrimRight(txt, "\n")
	if len(txt) == 0 {
		return 0
	}
	n := 1
	for _, r := range txt {
		if r == '\n' {
			n++
		}
	}
	return n
}

// rotatePoint applies rotation theta (in radians) about the origin to point p.
func rotatePoint(theta float64, p vg.Point) vg.Point {
	if theta == 0 {
		return p
	}
	x := float64(p.X)
	y := float64(p.Y)

	return vg.Point{
		X: vg.Length(x*math.Cos(theta) - y*math.Sin(theta)),
		Y: vg.Length(y*math.Cos(theta) + x*math.Sin(theta)),
	}
}
