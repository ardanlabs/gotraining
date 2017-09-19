// Copyright (C) 2011, Ross Light

package pdf

import (
	"bytes"
	"fmt"
	"image"
	"io"
	"math"
)

// writeCommand writes a PDF graphics command.
func writeCommand(w io.Writer, op string, args ...interface{}) error {
	for _, arg := range args {
		// TODO: Use the same buffer for all arguments
		if m, err := marshal(nil, arg); err == nil {
			if _, err := w.Write(append(m, ' ')); err != nil {
				return err
			}
		} else {
			return err
		}
	}
	if _, err := io.WriteString(w, op); err != nil {
		return err
	}
	if _, err := w.Write([]byte{'\n'}); err != nil {
		return err
	}
	return nil
}

// Canvas is a two-dimensional drawing region on a single page.  You can obtain
// a canvas once you have created a document.
type Canvas struct {
	doc          *Document
	page         *pageDict
	ref          Reference
	contents     *stream
	imageCounter uint
}

// Document returns the document the canvas is attached to.
func (canvas *Canvas) Document() *Document {
	return canvas.doc
}

// Close flushes the page's stream to the document.  This must be called once
// drawing has completed or else the document will be inconsistent.
func (canvas *Canvas) Close() error {
	return canvas.contents.Close()
}

// Size returns the page's media box (the size of the physical medium).
func (canvas *Canvas) Size() (width, height Unit) {
	mbox := canvas.page.MediaBox
	return mbox.Dx(), mbox.Dy()
}

// SetSize changes the page's media box (the size of the physical medium).
func (canvas *Canvas) SetSize(width, height Unit) {
	canvas.page.MediaBox = Rectangle{Point{0, 0}, Point{width, height}}
}

// CropBox returns the page's crop box.
func (canvas *Canvas) CropBox() Rectangle {
	return canvas.page.CropBox
}

// SetCropBox changes the page's crop box.
func (canvas *Canvas) SetCropBox(crop Rectangle) {
	canvas.page.CropBox = crop
}

// FillStroke fills then strokes the given path.  This operation has the same
// effect as performing a fill then a stroke, but does not repeat the path in
// the file.
func (canvas *Canvas) FillStroke(p *Path) {
	io.Copy(canvas.contents, &p.buf)
	writeCommand(canvas.contents, "B")
}

// Fill paints the area enclosed by the given path using the current fill color.
func (canvas *Canvas) Fill(p *Path) {
	io.Copy(canvas.contents, &p.buf)
	writeCommand(canvas.contents, "f")
}

// Stroke paints a line along the given path using the current stroke color.
func (canvas *Canvas) Stroke(p *Path) {
	io.Copy(canvas.contents, &p.buf)
	writeCommand(canvas.contents, "S")
}

// SetLineWidth changes the stroke width to the given value.
func (canvas *Canvas) SetLineWidth(w Unit) {
	writeCommand(canvas.contents, "w", w)
}

// SetLineDash changes the line dash pattern in the current graphics state.
// Examples:
//
//   c.SetLineDash(0, []Unit{})     // solid line
//   c.SetLineDash(0, []Unit{3})    // 3 units on, 3 units off...
//   c.SetLineDash(0, []Unit{2, 1}) // 2 units on, 1 unit off...
//   c.SetLineDash(1, []Unit{2})    // 1 unit on, 2 units off, 2 units on...
func (canvas *Canvas) SetLineDash(phase Unit, dash []Unit) {
	writeCommand(canvas.contents, "d", dash, phase)
}

// SetColor changes the current fill color to the given RGB triple (in device
// RGB space).
func (canvas *Canvas) SetColor(r, g, b float32) {
	writeCommand(canvas.contents, "rg", r, g, b)
}

// SetStrokeColor changes the current stroke color to the given RGB triple (in
// device RGB space).
func (canvas *Canvas) SetStrokeColor(r, g, b float32) {
	writeCommand(canvas.contents, "RG", r, g, b)
}

// Push saves a copy of the current graphics state.  The state can later be
// restored using Pop.
func (canvas *Canvas) Push() {
	writeCommand(canvas.contents, "q")
}

// Pop restores the most recently saved graphics state by popping it from the
// stack.
func (canvas *Canvas) Pop() {
	writeCommand(canvas.contents, "Q")
}

// Translate moves the canvas's coordinates system by the given offset.
func (canvas *Canvas) Translate(x, y Unit) {
	writeCommand(canvas.contents, "cm", 1, 0, 0, 1, x, y)
}

// Rotate rotates the canvas's coordinate system by a given angle (in radians).
func (canvas *Canvas) Rotate(theta float32) {
	s, c := math.Sin(float64(theta)), math.Cos(float64(theta))
	writeCommand(canvas.contents, "cm", c, s, -s, c, 0, 0)
}

// Scale multiplies the canvas's coordinate system by the given scalars.
func (canvas *Canvas) Scale(x, y float32) {
	writeCommand(canvas.contents, "cm", x, 0, 0, y, 0, 0)
}

// Transform concatenates a 3x3 matrix with the current transformation matrix.
// The arguments map to values in the matrix as shown below:
//
//  / a b 0 \
//  | c d 0 |
//  \ e f 1 /
//
// For more information, see Section 8.3.4 of ISO 32000-1.
func (canvas *Canvas) Transform(a, b, c, d, e, f float32) {
	writeCommand(canvas.contents, "cm", a, b, c, d, e, f)
}

// DrawText paints a text object onto the canvas.
func (canvas *Canvas) DrawText(text *Text) {
	for fontName := range text.fonts {
		if _, ok := canvas.page.Resources.Font[fontName]; !ok {
			canvas.page.Resources.Font[fontName] = canvas.doc.standardFont(fontName)
		}
	}
	writeCommand(canvas.contents, "BT")
	io.Copy(canvas.contents, &text.buf)
	writeCommand(canvas.contents, "ET")
}

// DrawImage paints a raster image at the given location and scaled to the
// given dimensions.  If you want to render the same image multiple times in
// the same document, use DrawImageReference.
func (canvas *Canvas) DrawImage(img image.Image, rect Rectangle) {
	canvas.DrawImageReference(canvas.doc.AddImage(img), rect)
}

// DrawImageReference paints the raster image referenced in the document at the
// given location and scaled to the given dimensions.
func (canvas *Canvas) DrawImageReference(ref Reference, rect Rectangle) {
	name := canvas.nextImageName()
	canvas.page.Resources.XObject[name] = ref

	canvas.Push()
	canvas.Transform(float32(rect.Dx()), 0, 0, float32(rect.Dy()), float32(rect.Min.X), float32(rect.Min.Y))
	writeCommand(canvas.contents, "Do", name)
	canvas.Pop()
}

// DrawLine paints a straight line from pt1 to pt2 using the current stroke
// color and line width.
func (canvas *Canvas) DrawLine(pt1, pt2 Point) {
	var path Path
	path.Move(pt1)
	path.Line(pt2)
	canvas.Stroke(&path)
}

const anonymousImageFormat = "__image%d__"

func (canvas *Canvas) nextImageName() name {
	var n name
	for {
		n = name(fmt.Sprintf(anonymousImageFormat, canvas.imageCounter))
		canvas.imageCounter++
		if _, ok := canvas.page.Resources.XObject[n]; !ok {
			break
		}
	}
	return n
}

// Path is a shape that can be painted on a canvas.  The zero value is an empty
// path.
type Path struct {
	buf bytes.Buffer
}

// Move begins a new subpath by moving the current point to the given location.
func (path *Path) Move(pt Point) {
	writeCommand(&path.buf, "m", pt.X, pt.Y)
}

// Line appends a line segment from the current point to the given location.
func (path *Path) Line(pt Point) {
	writeCommand(&path.buf, "l", pt.X, pt.Y)
}

// Curve appends a cubic Bezier curve to the path.
func (path *Path) Curve(pt1, pt2, pt3 Point) {
	writeCommand(&path.buf, "c", pt1.X, pt1.Y, pt2.X, pt2.Y, pt3.X, pt3.Y)
}

// Rectangle appends a complete rectangle to the path.
func (path *Path) Rectangle(rect Rectangle) {
	writeCommand(&path.buf, "re", rect.Min.X, rect.Min.Y, rect.Dx(), rect.Dy())
}

// Close appends a line segment from the current point to the starting point of
// the subpath.
func (path *Path) Close() {
	writeCommand(&path.buf, "h")
}
