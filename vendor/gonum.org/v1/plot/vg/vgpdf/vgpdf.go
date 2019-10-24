// Copyright ©2015 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package vgpdf implements the vg.Canvas interface
// using gofpdf (github.com/jung-kurt/gofpdf).
package vgpdf // import "gonum.org/v1/plot/vg/vgpdf"

import (
	"bufio"
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"

	pdf "github.com/jung-kurt/gofpdf"

	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/fonts"
)

// DPI is the nominal resolution of drawing in PDF.
const DPI = 72

// Canvas implements the vg.Canvas interface,
// drawing to a PDF.
type Canvas struct {
	doc  *pdf.Fpdf
	w, h vg.Length

	dpi       int
	numImages int
	stack     []context
	fonts     map[vg.Font]struct{}

	// Switch to embed fonts in PDF file.
	// The default is to embed fonts.
	// This makes the PDF file more portable but also larger.
	embed bool
}

type context struct {
	fill  color.Color
	line  color.Color
	width vg.Length
}

// New creates a new PDF Canvas.
func New(w, h vg.Length) *Canvas {
	cfg := pdf.InitType{
		UnitStr: "pt",
		Size:    pdf.SizeType{Wd: w.Points(), Ht: h.Points()},
	}
	c := &Canvas{
		doc:   pdf.NewCustom(&cfg),
		w:     w,
		h:     h,
		dpi:   DPI,
		stack: make([]context, 1),
		fonts: make(map[vg.Font]struct{}),
		embed: true,
	}
	c.NextPage()
	vg.Initialize(c)
	return c
}

// EmbedFonts specifies whether the resulting PDF canvas should
// embed the fonts or not.
// EmbedFonts returns the previous value before modification.
func (c *Canvas) EmbedFonts(v bool) bool {
	prev := c.embed
	c.embed = v
	return prev
}

func (c *Canvas) DPI() float64 {
	return float64(c.dpi)
}

func (c *Canvas) context() *context {
	return &c.stack[len(c.stack)-1]
}

func (c *Canvas) Size() (w, h vg.Length) {
	return c.w, c.h
}

func (c *Canvas) SetLineWidth(w vg.Length) {
	c.context().width = w
	lw := c.unit(w)
	c.doc.SetLineWidth(lw)
}

func (c *Canvas) SetLineDash(dashes []vg.Length, offs vg.Length) {
	ds := make([]float64, len(dashes))
	for i, d := range dashes {
		ds[i] = c.unit(d)
	}
	c.doc.SetDashPattern(ds, c.unit(offs))
}

func (c *Canvas) SetColor(clr color.Color) {
	if clr == nil {
		clr = color.Black
	}
	c.context().line = clr
	c.context().fill = clr
	r, g, b, a := rgba(clr)
	c.doc.SetFillColor(r, g, b)
	c.doc.SetDrawColor(r, g, b)
	c.doc.SetTextColor(r, g, b)
	c.doc.SetAlpha(a, "Normal")
}

func (c *Canvas) Rotate(r float64) {
	c.doc.TransformRotate(-r*180/math.Pi, 0, 0)
}

func (c *Canvas) Translate(pt vg.Point) {
	xp, yp := c.pdfPoint(pt)
	c.doc.TransformTranslate(xp, yp)
}

func (c *Canvas) Scale(x float64, y float64) {
	c.doc.TransformScale(x*100, y*100, 0, 0)
}

func (c *Canvas) Push() {
	c.stack = append(c.stack, *c.context())
	c.doc.TransformBegin()
}

func (c *Canvas) Pop() {
	c.doc.TransformEnd()
	c.stack = c.stack[:len(c.stack)-1]
}

func (c *Canvas) Stroke(p vg.Path) {
	if c.context().width > 0 {
		c.pdfPath(p, "D")
	}
}

func (c *Canvas) Fill(p vg.Path) {
	c.pdfPath(p, "F")
}

func (c *Canvas) FillString(fnt vg.Font, pt vg.Point, str string) {
	if fnt.Size == 0 {
		return
	}

	c.font(fnt, pt)
	c.doc.SetFont(fnt.Name(), "", c.unit(fnt.Size))

	c.Push()
	defer c.Pop()
	c.Translate(pt)
	// go-fpdf uses the top left corner as origin.
	c.Scale(1, -1)
	left, top, right, bottom := c.sbounds(fnt, str)
	w := right - left
	h := bottom - top
	margin := c.doc.GetCellMargin()

	c.doc.MoveTo(-left-margin, top)
	c.doc.CellFormat(w, h, str, "", 0, "BL", false, 0, "")
}

func (c *Canvas) sbounds(fnt vg.Font, txt string) (left, top, right, bottom float64) {
	_, h := c.doc.GetFontSize()
	d := c.doc.GetFontDesc("", "")
	if d.Ascent == 0 {
		// not defined (standard font?), use average of 81%
		top = 0.81 * h
	} else {
		top = -float64(d.Ascent) * h / float64(d.Ascent-d.Descent)
	}
	return 0, top, c.doc.GetStringWidth(txt), top + h
}

// DrawImage implements the vg.Canvas.DrawImage method.
func (c *Canvas) DrawImage(rect vg.Rectangle, img image.Image) {
	opts := pdf.ImageOptions{ImageType: "png", ReadDpi: true}
	name := c.imageName()

	buf := new(bytes.Buffer)
	err := png.Encode(buf, img)
	if err != nil {
		log.Panicf("error encoding image to PNG: %v", err)
	}
	c.doc.RegisterImageOptionsReader(name, opts, buf)

	xp, yp := c.pdfPoint(rect.Min)
	wp, hp := c.pdfPoint(rect.Size())

	c.doc.ImageOptions(name, xp, yp, wp, hp, false, opts, 0, "")
}

// font registers a font and a size with the PDF canvas.
func (c *Canvas) font(fnt vg.Font, pt vg.Point) {
	if _, ok := c.fonts[fnt]; ok {
		return
	}
	if n, ok := vg.FontMap[fnt.Name()]; ok {
		raw, err := fonts.Asset(n + ".ttf")
		if err != nil {
			log.Panicf("vgpdf: could not load TTF data from asset for TTF font %q: %v", n+".ttf", err)
		}

		enc, err := fonts.Asset("cp1252.map")
		if err != nil {
			log.Panicf("vgpdf: could not load encoding map: %v", err)
		}

		zdata, jdata, err := makeFont(raw, enc, c.embed)
		if err != nil {
			log.Panicf("vgpdf: could not generate font data for PDF: %v", err)
		}

		c.fonts[fnt] = struct{}{}
		c.doc.AddFontFromBytes(fnt.Name(), "", jdata, zdata)
		return
	}
	log.Panicf("vgpdf: could not find font %q in the pre-registered fonts map", fnt.Name())
}

// pdfPath processes a vg.Path and applies it to the canvas.
func (c *Canvas) pdfPath(path vg.Path, style string) {
	var (
		xp float64
		yp float64
	)
	for _, comp := range path {
		switch comp.Type {
		case vg.MoveComp:
			xp, yp = c.pdfPoint(comp.Pos)
			c.doc.MoveTo(xp, yp)
		case vg.LineComp:
			c.doc.LineTo(c.pdfPoint(comp.Pos))
		case vg.ArcComp:
			c.arc(comp, style)
		case vg.CurveComp:
			px, py := c.pdfPoint(comp.Pos)
			switch len(comp.Control) {
			case 1:
				cx, cy := c.pdfPoint(comp.Control[0])
				c.doc.CurveTo(cx, cy, px, py)
			case 2:
				cx, cy := c.pdfPoint(comp.Control[0])
				dx, dy := c.pdfPoint(comp.Control[1])
				c.doc.CurveBezierCubicTo(cx, cy, dx, dy, px, py)
			default:
				panic("vgpdf: invalid number of control points")
			}
		case vg.CloseComp:
			c.doc.LineTo(xp, yp)
			c.doc.ClosePath()
		default:
			panic(fmt.Sprintf("Unknown path component type: %d\n", comp.Type))
		}
	}
	c.doc.DrawPath(style)
	return
}

func (c *Canvas) arc(comp vg.PathComp, style string) {
	x0 := comp.Pos.X + comp.Radius*vg.Length(math.Cos(comp.Start))
	y0 := comp.Pos.Y + comp.Radius*vg.Length(math.Sin(comp.Start))
	c.doc.LineTo(c.pdfPointXY(x0, y0))
	r := c.unit(comp.Radius)
	const deg = 180 / math.Pi
	angle := comp.Angle * deg
	beg := comp.Start * deg
	end := beg + angle
	x := c.unit(comp.Pos.X)
	y := c.unit(comp.Pos.Y)
	c.doc.Arc(x, y, r, r, angle, beg, end, style)
	x1 := comp.Pos.X + comp.Radius*vg.Length(math.Cos(comp.Start+comp.Angle))
	y1 := comp.Pos.Y + comp.Radius*vg.Length(math.Sin(comp.Start+comp.Angle))
	c.doc.MoveTo(c.pdfPointXY(x1, y1))
}

func (c *Canvas) pdfPointXY(x, y vg.Length) (float64, float64) {
	return c.unit(x), c.unit(y)
}

func (c *Canvas) pdfPoint(pt vg.Point) (float64, float64) {
	return c.unit(pt.X), c.unit(pt.Y)
}

// unit returns a fpdf.Unit, converted from a vg.Length.
func (c *Canvas) unit(l vg.Length) float64 {
	return l.Dots(c.DPI())
}

// imageName generates a unique image name for this PDF canvas
func (c *Canvas) imageName() string {
	c.numImages++
	return fmt.Sprintf("image_%03d.png", c.numImages)
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
	c.Pop()
	c.doc.Close()
	wc := writerCounter{Writer: w}
	b := bufio.NewWriter(&wc)
	if err := c.doc.Output(b); err != nil {
		return wc.n, err
	}
	err := b.Flush()
	return wc.n, err
}

// rgba converts a Go color into a gofpdf 3-tuple int + 1 float64
func rgba(c color.Color) (int, int, int, float64) {
	if c == nil {
		c = color.Black
	}
	r, g, b, a := c.RGBA()
	return int(r >> 8), int(g >> 8), int(b >> 8), float64(a) / math.MaxUint16
}

func makeFont(font, encoding []byte, embed bool) (z, j []byte, err error) {
	tmpdir, err := ioutil.TempDir("", "gofpdf-makefont-")
	if err != nil {
		return z, j, err
	}
	defer os.RemoveAll(tmpdir)

	indir := filepath.Join(tmpdir, "input")
	err = os.Mkdir(indir, 0755)
	if err != nil {
		return z, j, err
	}

	outdir := filepath.Join(tmpdir, "output")
	err = os.Mkdir(outdir, 0755)
	if err != nil {
		return z, j, err
	}

	fname := filepath.Join(indir, "font.ttf")
	encname := filepath.Join(indir, "cp1252.map")

	err = ioutil.WriteFile(fname, font, 0644)
	if err != nil {
		return z, j, err
	}

	err = ioutil.WriteFile(encname, encoding, 0644)
	if err != nil {
		return z, j, err
	}

	err = pdf.MakeFont(fname, encname, outdir, ioutil.Discard, embed)
	if err != nil {
		return z, j, err
	}

	if embed {
		z, err = ioutil.ReadFile(filepath.Join(outdir, "font.z"))
		if err != nil {
			return z, j, err
		}
	}
	j, err = ioutil.ReadFile(filepath.Join(outdir, "font.json"))

	return z, j, err
}

// NextPage creates a new page in the final PDF document.
// The new page is the new current page.
// Modifications applied to the canvas will only be applied to that new page.
func (c *Canvas) NextPage() {
	if c.doc.PageNo() > 0 {
		c.Pop()
	}
	c.doc.SetMargins(0, 0, 0)
	c.doc.AddPage()
	c.Push()
	c.Translate(vg.Point{0, c.h})
	c.Scale(1, -1)
}
