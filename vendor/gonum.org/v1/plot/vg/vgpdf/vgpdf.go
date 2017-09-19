// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package vgpdf implements the vg.Canvas interface
// using gopdf (bitbucket.org/zombiezen/gopdf/pdf).
package vgpdf // import "gonum.org/v1/plot/vg/vgpdf"

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"io"
	"math"

	"bitbucket.org/zombiezen/gopdf/pdf"

	"gonum.org/v1/plot/vg"
)

// DPI is the nominal resolution of drawing in PDF.
const DPI = 72

// Canvas implements the vg.Canvas interface,
// drawing to a PDF.
type Canvas struct {
	doc         *pdf.Document
	w, h        vg.Length
	page        *pdf.Canvas
	lineVisible bool
}

// New creates a new PDF Canvas.
func New(w, h vg.Length) *Canvas {
	c := &Canvas{
		doc:         pdf.New(),
		w:           w,
		h:           h,
		lineVisible: true,
	}
	c.page = c.doc.NewPage(unit(w), unit(h))
	vg.Initialize(c)
	return c
}

func (c *Canvas) Size() (w, h vg.Length) {
	return c.w, c.h
}

func (c *Canvas) SetLineWidth(w vg.Length) {
	c.page.SetLineWidth(unit(w))
	c.lineVisible = w > 0
}

func (c *Canvas) SetLineDash(dashes []vg.Length, offs vg.Length) {
	ds := make([]pdf.Unit, len(dashes))
	for i, d := range dashes {
		ds[i] = unit(d)
	}
	c.page.SetLineDash(unit(offs), ds)
}

func (c *Canvas) SetColor(clr color.Color) {
	c.page.SetStrokeColor(pdfColor(clr))
	c.page.SetColor(pdfColor(clr))
}

func (c *Canvas) Rotate(r float64) {
	c.page.Rotate(float32(r))
}

func (c *Canvas) Translate(pt vg.Point) {
	c.page.Translate(unit(pt.X), unit(pt.Y))
}

func (c *Canvas) Scale(x float64, y float64) {
	c.page.Scale(float32(x), float32(y))
}

func (c *Canvas) Push() {
	c.page.Push()
}

func (c *Canvas) Pop() {
	c.page.Pop()
}

func (c *Canvas) Stroke(p vg.Path) {
	if c.lineVisible {
		c.page.Stroke(pdfPath(c, p))
	}
}

func (c *Canvas) Fill(p vg.Path) {
	c.page.Fill(pdfPath(c, p))
}

func (c *Canvas) FillString(fnt vg.Font, pt vg.Point, str string) {
	t := new(pdf.Text)
	t.SetFont(fnt.Name(), unit(fnt.Size))
	t.NextLineOffset(unit(pt.X), unit(pt.Y))
	t.Text(str)
	c.page.DrawText(t)
}

// DrawImage implements the vg.Canvas.DrawImage method.
func (c *Canvas) DrawImage(rect vg.Rectangle, img image.Image) {
	r := pdf.Rectangle{
		Min: pdfPoint(rect.Min.X, rect.Min.Y),
		Max: pdfPoint(rect.Max.X, rect.Max.Y),
	}
	c.page.DrawImage(img, r)
}

// pdfPath returns a pdf.Path from a vg.Path.
func pdfPath(c *Canvas, path vg.Path) *pdf.Path {
	p := new(pdf.Path)
	for _, comp := range path {
		switch comp.Type {
		case vg.MoveComp:
			p.Move(pdfPoint(comp.Pos.X, comp.Pos.Y))
		case vg.LineComp:
			p.Line(pdfPoint(comp.Pos.X, comp.Pos.Y))
		case vg.ArcComp:
			arc(p, comp)
		case vg.CloseComp:
			p.Close()
		default:
			panic(fmt.Sprintf("Unknown path component type: %d\n", comp.Type))
		}
	}
	return p
}

// Approximate a circular arc using multiple
// cubic Bézier curves, one for each π/2 segment.
//
// This is from:
// 	http://hansmuller-flex.blogspot.com/2011/04/approximating-circular-arc-with-cubic.html
func arc(p *pdf.Path, comp vg.PathComp) {
	x0 := comp.Pos.X + comp.Radius*vg.Length(math.Cos(comp.Start))
	y0 := comp.Pos.Y + comp.Radius*vg.Length(math.Sin(comp.Start))
	p.Line(pdfPoint(x0, y0))

	a1 := comp.Start
	end := a1 + comp.Angle
	sign := 1.0
	if end < a1 {
		sign = -1.0
	}
	left := math.Abs(comp.Angle)

	// Square root of the machine epsilon for IEEE 64-bit floating
	// point values.  This is the equality threshold recommended
	// in Numerical Recipes, if I recall correctly—it's small enough.
	const epsilon = 1.4901161193847656e-08

	for left > epsilon {
		a2 := a1 + sign*math.Min(math.Pi/2, left)
		partialArc(p, comp.Pos.X, comp.Pos.Y, comp.Radius, a1, a2)
		left -= math.Abs(a2 - a1)
		a1 = a2
	}
}

// Approximate a circular arc of fewer than π/2
// radians with cubic Bézier curve.
func partialArc(p *pdf.Path, x, y, r vg.Length, a1, a2 float64) {
	a := (a2 - a1) / 2
	x4 := r * vg.Length(math.Cos(a))
	y4 := r * vg.Length(math.Sin(a))
	x1 := x4
	y1 := -y4

	const k = 0.5522847498 // some magic constant
	f := k * vg.Length(math.Tan(a))
	x2 := x1 + f*y4
	y2 := y1 + f*x4
	x3 := x2
	y3 := -y2

	// Rotate and translate points into position.
	ar := a + a1
	sinar := vg.Length(math.Sin(ar))
	cosar := vg.Length(math.Cos(ar))
	x2r := x2*cosar - y2*sinar + x
	y2r := x2*sinar + y2*cosar + y
	x3r := x3*cosar - y3*sinar + x
	y3r := x3*sinar + y3*cosar + y
	x4 = r*vg.Length(math.Cos(a2)) + x
	y4 = r*vg.Length(math.Sin(a2)) + y
	p.Curve(pdfPoint(x2r, y2r), pdfPoint(x3r, y3r), pdfPoint(x4, y4))
}

func pdfPoint(x, y vg.Length) pdf.Point {
	return pdf.Point{X: unit(x), Y: unit(y)}
}

func pdfColor(clr color.Color) (float32, float32, float32) {
	if clr == nil {
		clr = color.Black
	}
	r, g, b, _ := clr.RGBA()
	return float32(r) / math.MaxUint16,
		float32(g) / math.MaxUint16,
		float32(b) / math.MaxUint16
}

// unit returns a pdf.Unit, converted from a vg.Length.
func unit(l vg.Length) pdf.Unit {
	return pdf.Unit(l.Points()) * pdf.Pt
}

// WriterCounter implements the io.Writer interface, and counts
// the total number of bytes written.
type writerCounter struct {
	io.Writer
	n int64
}

func (w *writerCounter) Write(p []byte) (int, error) {
	n, err := w.Writer.Write(p)
	w.n += int64(n)
	return n, err
}

// WriteTo writes the Canvas to an io.Writer.
// After calling Write, the canvas is closed
// and may no longer be used for drawing.
func (c *Canvas) WriteTo(w io.Writer) (int64, error) {
	c.page.Close()
	wc := writerCounter{Writer: w}
	b := bufio.NewWriter(&wc)
	if err := c.doc.Encode(b); err != nil {
		return wc.n, err
	}
	err := b.Flush()
	return wc.n, err
}
