// Copyright ©2020 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package draw // import "gonum.org/v1/plot/vg/draw"

import (
	"math"
	"strings"

	"gonum.org/v1/plot/vg"
)

// PlainTextHandler is a text/plain handler.
type PlainTextHandler struct{}

// Box returns the bounding box of the given text where:
//  - width is the horizontal space from the origin.
//  - height is the vertical space above the baseline.
//  - depth is the vertical space below the baseline, a negative number.
func (hdlr PlainTextHandler) Box(txt string, fnt vg.Font) (width, height, depth vg.Length) {
	ext := fnt.Extents()

	nl := hdlr.textNLines(txt)
	if nl != 0 {
		height = ext.Height*vg.Length(nl-1) + ext.Ascent
		depth = -ext.Descent
	}

	for _, line := range strings.Split(strings.TrimRight(txt, "\n"), "\n") {
		w := fnt.Width(line)
		if w > width {
			width = w
		}
	}

	return width, height, depth
}

// Draw renders the given text with the provided style and position
// on the canvas.
func (hdlr PlainTextHandler) Draw(c *Canvas, txt string, sty TextStyle, pt vg.Point) {
	txt = strings.TrimRight(txt, "\n")
	if len(txt) == 0 {
		return
	}

	c.SetColor(sty.Color)

	if sty.Rotation != 0 {
		c.Push()
		c.Rotate(sty.Rotation)
	}

	cos := vg.Length(math.Cos(sty.Rotation))
	sin := vg.Length(math.Sin(sty.Rotation))
	pt.X, pt.Y = pt.Y*sin+pt.X*cos, pt.Y*cos-pt.X*sin

	nl := hdlr.textNLines(txt)
	ht := sty.Height(txt)
	pt.Y += ht*vg.Length(sty.YAlign) - sty.Font.Extents().Ascent
	for i, line := range strings.Split(txt, "\n") {
		xoffs := vg.Length(sty.XAlign) * sty.Font.Width(line)
		n := vg.Length(nl - i)
		c.FillString(sty.Font, pt.Add(vg.Point{X: xoffs, Y: n * sty.Font.Size}), line)
	}

	if sty.Rotation != 0 {
		c.Pop()
	}
}

// textNLines returns the number of lines in the text.
func (PlainTextHandler) textNLines(txt string) int {
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

var (
	_ TextHandler = (*PlainTextHandler)(nil)
)
