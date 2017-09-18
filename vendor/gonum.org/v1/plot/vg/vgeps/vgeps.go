// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package vgeps implements the vg.Canvas interface using
// encapsulated postscript.
package vgeps // import "gonum.org/v1/plot/vg/vgeps"

import (
	"bufio"
	"bytes"
	"fmt"
	"image"
	"image/color"
	"io"
	"math"
	"time"

	"gonum.org/v1/plot/vg"
)

// DPI is the nominal resolution of drawing in EPS.
const DPI = 72

type Canvas struct {
	stk  []ctx
	w, h vg.Length
	buf  *bytes.Buffer
}

type ctx struct {
	color  color.Color
	width  vg.Length
	dashes []vg.Length
	offs   vg.Length
	font   string
	fsize  vg.Length
}

// pr is the amount of precision to use when outputting float64s.
const pr = 5

// New returns a new Canvas.
func New(w, h vg.Length) *Canvas {
	return NewTitle(w, h, "")
}

// NewTitle returns a new Canvas with the given title string.
func NewTitle(w, h vg.Length, title string) *Canvas {
	c := &Canvas{
		stk: []ctx{ctx{}},
		w:   w,
		h:   h,
		buf: new(bytes.Buffer),
	}
	c.buf.WriteString("%%!PS-Adobe-3.0 EPSF-3.0\n")
	c.buf.WriteString("%%Creator gonum.org/v1/plot/vg/vgeps\n")
	c.buf.WriteString("%%Title: " + title + "\n")
	c.buf.WriteString(fmt.Sprintf("%%%%BoundingBox: 0 0 %.*g %.*g\n",
		pr, w.Dots(DPI),
		pr, h.Dots(DPI)))
	c.buf.WriteString(fmt.Sprintf("%%%%CreationDate: %s\n", time.Now()))
	c.buf.WriteString("%%Orientation: Portrait\n")
	c.buf.WriteString("%%EndComments\n")
	c.buf.WriteString("\n")
	vg.Initialize(c)
	return c
}

func (c *Canvas) Size() (w, h vg.Length) {
	return c.w, c.h
}

// cur returns the top context on the stack.
func (e *Canvas) cur() *ctx {
	return &e.stk[len(e.stk)-1]
}

func (e *Canvas) SetLineWidth(w vg.Length) {
	if e.cur().width != w {
		e.cur().width = w
		fmt.Fprintf(e.buf, "%.*g setlinewidth\n", pr, w.Dots(DPI))
	}
}

func (e *Canvas) SetLineDash(dashes []vg.Length, o vg.Length) {
	cur := e.cur().dashes
	dashEq := len(dashes) == len(cur)
	for i := 0; dashEq && i < len(dashes); i++ {
		if dashes[i] != cur[i] {
			dashEq = false
		}
	}
	if !dashEq || e.cur().offs != o {
		e.cur().dashes = dashes
		e.cur().offs = o
		e.buf.WriteString("[")
		for _, d := range dashes {
			fmt.Fprintf(e.buf, " %.*g", pr, d.Dots(DPI))
		}
		e.buf.WriteString(" ] ")
		fmt.Fprintf(e.buf, "%.*g setdash\n", pr, o.Dots(DPI))
	}
}

func (e *Canvas) SetColor(c color.Color) {
	if c == nil {
		c = color.Black
	}
	if e.cur().color != c {
		e.cur().color = c
		r, g, b, _ := c.RGBA()
		mx := float64(math.MaxUint16)
		fmt.Fprintf(e.buf, "%.*g %.*g %.*g setrgbcolor\n", pr, float64(r)/mx,
			pr, float64(g)/mx, pr, float64(b)/mx)
	}
}

func (e *Canvas) Rotate(r float64) {
	fmt.Fprintf(e.buf, "%.*g rotate\n", pr, r*180/math.Pi)
}

func (e *Canvas) Translate(pt vg.Point) {
	fmt.Fprintf(e.buf, "%.*g %.*g translate\n",
		pr, pt.X.Dots(DPI), pr, pt.Y.Dots(DPI))
}

func (e *Canvas) Scale(x, y float64) {
	fmt.Fprintf(e.buf, "%.*g %.*g scale\n", pr, x, pr, y)
}

func (e *Canvas) Push() {
	e.stk = append(e.stk, *e.cur())
	e.buf.WriteString("gsave\n")
}

func (e *Canvas) Pop() {
	e.stk = e.stk[:len(e.stk)-1]
	e.buf.WriteString("grestore\n")
}

func (e *Canvas) Stroke(path vg.Path) {
	if e.cur().width <= 0 {
		return
	}
	e.trace(path)
	e.buf.WriteString("stroke\n")
}

func (e *Canvas) Fill(path vg.Path) {
	e.trace(path)
	e.buf.WriteString("fill\n")
}

func (e *Canvas) trace(path vg.Path) {
	e.buf.WriteString("newpath\n")
	for _, comp := range path {
		switch comp.Type {
		case vg.MoveComp:
			fmt.Fprintf(e.buf, "%.*g %.*g moveto\n", pr, comp.Pos.X, pr, comp.Pos.Y)
		case vg.LineComp:
			fmt.Fprintf(e.buf, "%.*g %.*g lineto\n", pr, comp.Pos.X, pr, comp.Pos.Y)
		case vg.ArcComp:
			end := comp.Start + comp.Angle
			arcOp := "arc"
			if comp.Angle < 0 {
				arcOp = "arcn"
			}
			fmt.Fprintf(e.buf, "%.*g %.*g %.*g %.*g %.*g %s\n", pr, comp.Pos.X, pr, comp.Pos.Y,
				pr, comp.Radius, pr, comp.Start*180/math.Pi, pr,
				end*180/math.Pi, arcOp)
		case vg.CloseComp:
			e.buf.WriteString("closepath\n")
		default:
			panic(fmt.Sprintf("Unknown path component type: %d\n", comp.Type))
		}
	}
}

func (e *Canvas) FillString(fnt vg.Font, pt vg.Point, str string) {
	if e.cur().font != fnt.Name() || e.cur().fsize != fnt.Size {
		e.cur().font = fnt.Name()
		e.cur().fsize = fnt.Size
		fmt.Fprintf(e.buf, "/%s findfont %.*g scalefont setfont\n",
			fnt.Name(), pr, fnt.Size)
	}
	fmt.Fprintf(e.buf, "%.*g %.*g moveto\n", pr, pt.X.Dots(DPI), pr, pt.Y.Dots(DPI))
	fmt.Fprintf(e.buf, "(%s) show\n", str)
}

// DrawImage implements the vg.Canvas.DrawImage method.
func (c *Canvas) DrawImage(rect vg.Rectangle, img image.Image) {
	// FIXME: https://github.com/gonum/plot/issues/271
	panic("vgeps: DrawImage not implemented")
}

// WriteTo writes the canvas to an io.Writer.
func (e *Canvas) WriteTo(w io.Writer) (int64, error) {
	b := bufio.NewWriter(w)
	n, err := e.buf.WriteTo(b)
	if err != nil {
		return n, err
	}
	m, err := fmt.Fprintln(b, "showpage")
	n += int64(m)
	if err != nil {
		return n, err
	}
	return n, b.Flush()
}
