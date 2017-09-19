// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package vgimg implements the vg.Canvas interface using
// draw2d (github.com/llgcode/draw2d)
// as a backend to output raster images.
package vgimg // import "gonum.org/v1/plot/vg/vgimg"

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io"

	"golang.org/x/image/tiff"

	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"

	"gonum.org/v1/plot/vg"
)

// Canvas implements the vg.Canvas interface,
// drawing to an image.Image using draw2d.
type Canvas struct {
	gc    draw2d.GraphicContext
	img   draw.Image
	w, h  vg.Length
	color []color.Color

	// dpi is the number of dots per inch for this canvas.
	dpi int

	// width is the current line width.
	width vg.Length
}

const (
	// DefaultDPI is the default dot resolution for image
	// drawing in dots per inch.
	DefaultDPI = 96

	// DefaultWidth and DefaultHeight are the default canvas
	// dimensions.
	DefaultWidth  = 4 * vg.Inch
	DefaultHeight = 4 * vg.Inch
)

// New returns a new image canvas.
func New(w, h vg.Length) *Canvas {
	return NewWith(UseWH(w, h))
}

// NewWith returns a new image canvas created according to the specified
// options. The currently accepted options are UseWH,
// UseDPI, UseImage, and UseImageWithContext.
// Each of the options specifies the size of the canvas (UseWH, UseImage),
// the resolution of the canvas (UseDPI), or both (useImageWithContext).
// If size or resolution are not specified, defaults are used.
// It panics if size and resolution are overspecified (i.e., too many options are
// passed).
func NewWith(o ...option) *Canvas {
	c := new(Canvas)
	var g uint32
	for _, opt := range o {
		f := opt(c)
		if g&f != 0 {
			panic("incompatible options")
		}
		g |= f
	}
	if c.dpi == 0 {
		c.dpi = DefaultDPI
	}
	if c.w == 0 { // h should also == 0.
		if c.img == nil {
			c.w = DefaultWidth
			c.h = DefaultHeight
		} else {
			w := float64(c.img.Bounds().Max.X - c.img.Bounds().Min.X)
			h := float64(c.img.Bounds().Max.Y - c.img.Bounds().Min.Y)
			c.w = vg.Length(w/float64(c.dpi)) * vg.Inch
			c.h = vg.Length(h/float64(c.dpi)) * vg.Inch
		}
	}
	if c.img == nil {
		w := c.w / vg.Inch * vg.Length(c.dpi)
		h := c.h / vg.Inch * vg.Length(c.dpi)
		c.img = draw.Image(image.NewRGBA(image.Rect(0, 0, int(w+0.5), int(h+0.5))))
	}
	if c.gc == nil {
		h := float64(c.img.Bounds().Max.Y - c.img.Bounds().Min.Y)
		c.gc = draw2dimg.NewGraphicContext(c.img)
		c.gc.SetDPI(c.dpi)
		c.gc.Scale(1, -1)
		c.gc.Translate(0, -h)
	}
	draw.Draw(c.img, c.img.Bounds(), image.White, image.ZP, draw.Src)
	c.color = []color.Color{color.Black}
	vg.Initialize(c)
	return c
}

// These constants are used to ensure that the options
// used when initializing a canvas are compatible with
// each other.
const (
	setsDPI = 1 << iota
	setsSize
)

type option func(*Canvas) uint32

// UseWH specifies the width and height of the canvas.
// The size is rounded up to the nearest pixel.
func UseWH(w, h vg.Length) option {
	return func(c *Canvas) uint32 {
		if w <= 0 || h <= 0 {
			panic("w and h must both be > 0.")
		}
		c.w, c.h = w, h
		return setsSize
	}
}

// UseDPI sets the dots per inch of a canvas. It should only be
// used as an option argument when initializing a new canvas.
func UseDPI(dpi int) option {
	if dpi <= 0 {
		panic("DPI must be > 0.")
	}
	return func(c *Canvas) uint32 {
		c.dpi = dpi
		return setsDPI
	}
}

// UseImage specifies an image to create
// the canvas from. The
// minimum point of the given image
// should probably be 0,0.
func UseImage(img draw.Image) option {
	return func(c *Canvas) uint32 {
		c.img = img
		return setsSize
	}
}

// UseImageWithContext specifies both an image
// and a graphic context to create the canvas from.
// The minimum point of the given image
// should probably be 0,0.
func UseImageWithContext(img draw.Image, gc draw2d.GraphicContext) option {
	return func(c *Canvas) uint32 {
		c.img = img
		c.gc = gc
		c.dpi = gc.GetDPI()
		return setsDPI | setsSize
	}
}

// Image returns the image the canvas is drawing to.
//
// The dimensions of the returned image must not be modified.
func (c *Canvas) Image() draw.Image {
	return c.img
}

func (c *Canvas) Size() (w, h vg.Length) {
	return c.w, c.h
}

func (c *Canvas) SetLineWidth(w vg.Length) {
	c.width = w
	c.gc.SetLineWidth(w.Dots(c.DPI()))
}

func (c *Canvas) SetLineDash(ds []vg.Length, offs vg.Length) {
	dashes := make([]float64, len(ds))
	for i, d := range ds {
		dashes[i] = d.Dots(c.DPI())
	}
	c.gc.SetLineDash(dashes, offs.Dots(c.DPI()))
}

func (c *Canvas) SetColor(clr color.Color) {
	if clr == nil {
		clr = color.Black
	}
	c.gc.SetFillColor(clr)
	c.gc.SetStrokeColor(clr)
	c.color[len(c.color)-1] = clr
}

func (c *Canvas) Rotate(t float64) {
	c.gc.Rotate(t)
}

func (c *Canvas) Translate(pt vg.Point) {
	c.gc.Translate(pt.X.Dots(c.DPI()), pt.Y.Dots(c.DPI()))
}

func (c *Canvas) Scale(x, y float64) {
	c.gc.Scale(x, y)
}

func (c *Canvas) Push() {
	c.color = append(c.color, c.color[len(c.color)-1])
	c.gc.Save()
}

func (c *Canvas) Pop() {
	c.color = c.color[:len(c.color)-1]
	c.gc.Restore()
}

func (c *Canvas) Stroke(p vg.Path) {
	if c.width <= 0 {
		return
	}
	c.outline(p)
	c.gc.Stroke()
}

func (c *Canvas) Fill(p vg.Path) {
	c.outline(p)
	c.gc.Fill()
}

func (c *Canvas) outline(p vg.Path) {
	c.gc.BeginPath()
	for _, comp := range p {
		switch comp.Type {
		case vg.MoveComp:
			c.gc.MoveTo(comp.Pos.X.Dots(c.DPI()), comp.Pos.Y.Dots(c.DPI()))

		case vg.LineComp:
			c.gc.LineTo(comp.Pos.X.Dots(c.DPI()), comp.Pos.Y.Dots(c.DPI()))

		case vg.ArcComp:
			c.gc.ArcTo(comp.Pos.X.Dots(c.DPI()), comp.Pos.Y.Dots(c.DPI()),
				comp.Radius.Dots(c.DPI()), comp.Radius.Dots(c.DPI()),
				comp.Start, comp.Angle)

		case vg.CloseComp:
			c.gc.Close()

		default:
			panic(fmt.Sprintf("Unknown path component: %d", comp.Type))
		}
	}
}

func (c *Canvas) DPI() float64 {
	return float64(c.gc.GetDPI())
}

func (c *Canvas) FillString(font vg.Font, pt vg.Point, str string) {
	c.gc.Save()
	defer c.gc.Restore()

	data := draw2d.FontData{Name: font.Name()}
	if !registeredFont[font.Name()] {
		draw2d.RegisterFont(data, font.Font())
		registeredFont[font.Name()] = true
	}
	c.gc.SetFontData(data)
	c.gc.SetFontSize(font.Size.Points())
	c.gc.Translate(pt.X.Dots(c.DPI()), pt.Y.Dots(c.DPI()))
	c.gc.Scale(1, -1)
	c.gc.FillString(str)
}

// DrawImage implements the vg.Canvas.DrawImage method.
func (c *Canvas) DrawImage(rect vg.Rectangle, img image.Image) {
	var (
		dpi    = c.DPI()
		min    = rect.Min
		xmin   = min.X.Dots(dpi)
		ymin   = min.Y.Dots(dpi)
		rsz    = rect.Size()
		width  = rsz.X.Dots(dpi)
		height = rsz.Y.Dots(dpi)
		dx     = float64(img.Bounds().Dx())
		dy     = float64(img.Bounds().Dy())
	)
	c.gc.Save()
	c.gc.Scale(1, -1)
	c.gc.Translate(xmin, -ymin-height)
	c.gc.Scale(width/dx, height/dy)
	c.gc.DrawImage(img)
	c.gc.Restore()
}

var (
	// RegisteredFont contains the set of font names
	// that have already been registered with draw2d.
	registeredFont = map[string]bool{}
)

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

// A JpegCanvas is an image canvas with a WriteTo method
// that writes a jpeg image.
type JpegCanvas struct {
	*Canvas
}

// WriteTo implements the io.WriterTo interface, writing a jpeg image.
func (c JpegCanvas) WriteTo(w io.Writer) (int64, error) {
	wc := writerCounter{Writer: w}
	b := bufio.NewWriter(&wc)
	if err := jpeg.Encode(b, c.img, nil); err != nil {
		return wc.n, err
	}
	err := b.Flush()
	return wc.n, err
}

// A PngCanvas is an image canvas with a WriteTo method that
// writes a png image.
type PngCanvas struct {
	*Canvas
}

// WriteTo implements the io.WriterTo interface, writing a png image.
func (c PngCanvas) WriteTo(w io.Writer) (int64, error) {
	wc := writerCounter{Writer: w}
	b := bufio.NewWriter(&wc)
	if err := png.Encode(b, c.img); err != nil {
		return wc.n, err
	}
	err := b.Flush()
	return wc.n, err
}

// A TiffCanvas is an image canvas with a WriteTo method that
// writes a tiff image.
type TiffCanvas struct {
	*Canvas
}

// WriteTo implements the io.WriterTo interface, writing a tiff image.
func (c TiffCanvas) WriteTo(w io.Writer) (int64, error) {
	wc := writerCounter{Writer: w}
	b := bufio.NewWriter(&wc)
	if err := tiff.Encode(b, c.img, nil); err != nil {
		return wc.n, err
	}
	err := b.Flush()
	return wc.n, err
}
