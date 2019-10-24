// Copyright ©2015 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package draw // import "gonum.org/v1/plot/vg/draw"

import (
	"fmt"
	"image/color"
	"math"
	"strings"

	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/vgeps"
	"gonum.org/v1/plot/vg/vgimg"
	"gonum.org/v1/plot/vg/vgpdf"
	"gonum.org/v1/plot/vg/vgsvg"
)

// A Canvas is a vector graphics canvas along with
// an associated Rectangle defining a section of the canvas
// to which drawing should take place.
type Canvas struct {
	vg.Canvas
	vg.Rectangle
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
}

// XAlignment specifies text alignment in the X direction. Three preset
// options are available, but an arbitrary alignment
// can also be specified using XAlignment(desired number).
type XAlignment float64

const (
	// XLeft aligns the left edge of the text with the specified location.
	XLeft XAlignment = 0
	// XCenter aligns the horizontal center of the text with the specified location.
	XCenter XAlignment = -0.5
	// XRight aligns the right edge of the text with the specified location.
	XRight XAlignment = -1
)

// YAlignment specifies text alignment in the Y direction. Three preset
// options are available, but an arbitrary alignment
// can also be specified using YAlignment(desired number).
type YAlignment float64

const (
	// YTop aligns the top of of the text with the specified location.
	YTop YAlignment = -1
	// YCenter aligns the vertical center of the text with the specified location.
	YCenter YAlignment = -0.5
	// YBottom aligns the bottom of the text with the specified location.
	YBottom YAlignment = 0
)

// LineStyle describes what a line will look like.
type LineStyle struct {
	// Color is the color of the line.
	Color color.Color

	// Width is the width of the line.
	Width vg.Length

	Dashes   []vg.Length
	DashOffs vg.Length
}

// A GlyphStyle specifies the look of a glyph used to draw
// a point on a plot.
type GlyphStyle struct {
	// Color is the color used to draw the glyph.
	color.Color

	// Radius specifies the size of the glyph's radius.
	Radius vg.Length

	// Shape draws the shape of the glyph.
	Shape GlyphDrawer
}

// A GlyphDrawer wraps the DrawGlyph function.
type GlyphDrawer interface {
	// DrawGlyph draws the glyph at the given
	// point, with the given color and radius.
	DrawGlyph(*Canvas, GlyphStyle, vg.Point)
}

// DrawGlyph draws the given glyph to the draw
// area.  If the point is not within the Canvas
// or the sty.Shape is nil then nothing is drawn.
func (c *Canvas) DrawGlyph(sty GlyphStyle, pt vg.Point) {
	if sty.Shape == nil || !c.Contains(pt) {
		return
	}
	c.SetColor(sty.Color)
	sty.Shape.DrawGlyph(c, sty, pt)
}

// DrawGlyphNoClip draws the given glyph to the draw
// area.  If the sty.Shape is nil then nothing is drawn.
func (c *Canvas) DrawGlyphNoClip(sty GlyphStyle, pt vg.Point) {
	if sty.Shape == nil {
		return
	}
	c.SetColor(sty.Color)
	sty.Shape.DrawGlyph(c, sty, pt)
}

// Rectangle returns the rectangle surrounding this glyph,
// assuming that it is drawn centered at 0,0
func (g GlyphStyle) Rectangle() vg.Rectangle {
	return vg.Rectangle{
		Min: vg.Point{X: -g.Radius, Y: -g.Radius},
		Max: vg.Point{X: +g.Radius, Y: +g.Radius},
	}
}

// CircleGlyph is a glyph that draws a solid circle.
type CircleGlyph struct{}

// DrawGlyph implements the GlyphDrawer interface.
func (CircleGlyph) DrawGlyph(c *Canvas, sty GlyphStyle, pt vg.Point) {
	var p vg.Path
	p.Move(vg.Point{X: pt.X + sty.Radius, Y: pt.Y})
	p.Arc(pt, sty.Radius, 0, 2*math.Pi)
	p.Close()
	c.Fill(p)
}

// RingGlyph is a glyph that draws the outline of a circle.
type RingGlyph struct{}

// DrawGlyph implements the Glyph interface.
func (RingGlyph) DrawGlyph(c *Canvas, sty GlyphStyle, pt vg.Point) {
	c.SetLineStyle(LineStyle{Color: sty.Color, Width: vg.Points(0.5)})
	var p vg.Path
	p.Move(vg.Point{X: pt.X + sty.Radius, Y: pt.Y})
	p.Arc(pt, sty.Radius, 0, 2*math.Pi)
	p.Close()
	c.Stroke(p)
}

const (
	cosπover4 = vg.Length(.707106781202420)
	sinπover6 = vg.Length(.500000000025921)
	cosπover6 = vg.Length(.866025403769473)
)

// SquareGlyph is a glyph that draws the outline of a square.
type SquareGlyph struct{}

// DrawGlyph implements the Glyph interface.
func (SquareGlyph) DrawGlyph(c *Canvas, sty GlyphStyle, pt vg.Point) {
	c.SetLineStyle(LineStyle{Color: sty.Color, Width: vg.Points(0.5)})
	x := (sty.Radius-sty.Radius*cosπover4)/2 + sty.Radius*cosπover4
	var p vg.Path
	p.Move(vg.Point{X: pt.X - x, Y: pt.Y - x})
	p.Line(vg.Point{X: pt.X + x, Y: pt.Y - x})
	p.Line(vg.Point{X: pt.X + x, Y: pt.Y + x})
	p.Line(vg.Point{X: pt.X - x, Y: pt.Y + x})
	p.Close()
	c.Stroke(p)
}

// BoxGlyph is a glyph that draws a filled square.
type BoxGlyph struct{}

// DrawGlyph implements the Glyph interface.
func (BoxGlyph) DrawGlyph(c *Canvas, sty GlyphStyle, pt vg.Point) {
	x := (sty.Radius-sty.Radius*cosπover4)/2 + sty.Radius*cosπover4
	var p vg.Path
	p.Move(vg.Point{X: pt.X - x, Y: pt.Y - x})
	p.Line(vg.Point{X: pt.X + x, Y: pt.Y - x})
	p.Line(vg.Point{X: pt.X + x, Y: pt.Y + x})
	p.Line(vg.Point{X: pt.X - x, Y: pt.Y + x})
	p.Close()
	c.Fill(p)
}

// TriangleGlyph is a glyph that draws the outline of a triangle.
type TriangleGlyph struct{}

// DrawGlyph implements the Glyph interface.
func (TriangleGlyph) DrawGlyph(c *Canvas, sty GlyphStyle, pt vg.Point) {
	c.SetLineStyle(LineStyle{Color: sty.Color, Width: vg.Points(0.5)})
	r := sty.Radius + (sty.Radius-sty.Radius*sinπover6)/2
	var p vg.Path
	p.Move(vg.Point{X: pt.X, Y: pt.Y + r})
	p.Line(vg.Point{X: pt.X - r*cosπover6, Y: pt.Y - r*sinπover6})
	p.Line(vg.Point{X: pt.X + r*cosπover6, Y: pt.Y - r*sinπover6})
	p.Close()
	c.Stroke(p)
}

// PyramidGlyph is a glyph that draws a filled triangle.
type PyramidGlyph struct{}

// DrawGlyph implements the Glyph interface.
func (PyramidGlyph) DrawGlyph(c *Canvas, sty GlyphStyle, pt vg.Point) {
	r := sty.Radius + (sty.Radius-sty.Radius*sinπover6)/2
	var p vg.Path
	p.Move(vg.Point{X: pt.X, Y: pt.Y + r})
	p.Line(vg.Point{X: pt.X - r*cosπover6, Y: pt.Y - r*sinπover6})
	p.Line(vg.Point{X: pt.X + r*cosπover6, Y: pt.Y - r*sinπover6})
	p.Close()
	c.Fill(p)
}

// PlusGlyph is a glyph that draws a plus sign
type PlusGlyph struct{}

// DrawGlyph implements the Glyph interface.
func (PlusGlyph) DrawGlyph(c *Canvas, sty GlyphStyle, pt vg.Point) {
	c.SetLineStyle(LineStyle{Color: sty.Color, Width: vg.Points(0.5)})
	r := sty.Radius
	var p vg.Path
	p.Move(vg.Point{X: pt.X, Y: pt.Y + r})
	p.Line(vg.Point{X: pt.X, Y: pt.Y - r})
	c.Stroke(p)
	p = vg.Path{}
	p.Move(vg.Point{X: pt.X - r, Y: pt.Y})
	p.Line(vg.Point{X: pt.X + r, Y: pt.Y})
	c.Stroke(p)
}

// CrossGlyph is a glyph that draws a big X.
type CrossGlyph struct{}

// DrawGlyph implements the Glyph interface.
func (CrossGlyph) DrawGlyph(c *Canvas, sty GlyphStyle, pt vg.Point) {
	c.SetLineStyle(LineStyle{Color: sty.Color, Width: vg.Points(0.5)})
	r := sty.Radius * cosπover4
	var p vg.Path
	p.Move(vg.Point{X: pt.X - r, Y: pt.Y - r})
	p.Line(vg.Point{X: pt.X + r, Y: pt.Y + r})
	c.Stroke(p)
	p = vg.Path{}
	p.Move(vg.Point{X: pt.X - r, Y: pt.Y + r})
	p.Line(vg.Point{X: pt.X + r, Y: pt.Y - r})
	c.Stroke(p)
}

// New returns a new (bounded) draw.Canvas.
func New(c vg.CanvasSizer) Canvas {
	w, h := c.Size()
	return NewCanvas(c, w, h)
}

// NewFormattedCanvas creates a new vg.CanvasWriterTo with the specified
// image format.
//
// Supported formats are:
//
//  eps, jpg|jpeg, pdf, png, svg, and tif|tiff.
func NewFormattedCanvas(w, h vg.Length, format string) (vg.CanvasWriterTo, error) {
	var c vg.CanvasWriterTo
	switch format {
	case "eps":
		c = vgeps.New(w, h)

	case "jpg", "jpeg":
		c = vgimg.JpegCanvas{Canvas: vgimg.New(w, h)}

	case "pdf":
		c = vgpdf.New(w, h)

	case "png":
		c = vgimg.PngCanvas{Canvas: vgimg.New(w, h)}

	case "svg":
		c = vgsvg.New(w, h)

	case "tif", "tiff":
		c = vgimg.TiffCanvas{Canvas: vgimg.New(w, h)}

	default:
		return nil, fmt.Errorf("unsupported format: %q", format)
	}
	return c, nil
}

// NewCanvas returns a new (bounded) draw.Canvas of the given size.
func NewCanvas(c vg.Canvas, w, h vg.Length) Canvas {
	return Canvas{
		Canvas: c,
		Rectangle: vg.Rectangle{
			Min: vg.Point{X: 0, Y: 0},
			Max: vg.Point{X: w, Y: h},
		},
	}
}

// Center returns the center point of the area
func (c *Canvas) Center() vg.Point {
	return vg.Point{
		X: (c.Max.X-c.Min.X)/2 + c.Min.X,
		Y: (c.Max.Y-c.Min.Y)/2 + c.Min.Y,
	}
}

// Contains returns true if the Canvas contains the point.
func (c *Canvas) Contains(p vg.Point) bool {
	return c.ContainsX(p.X) && c.ContainsY(p.Y)
}

// ContainsX returns true if the Canvas contains the
// x coordinate.
func (c *Canvas) ContainsX(x vg.Length) bool {
	return x <= c.Max.X+slop && x >= c.Min.X-slop
}

// ContainsY returns true if the Canvas contains the
// y coordinate.
func (c *Canvas) ContainsY(y vg.Length) bool {
	return y <= c.Max.Y+slop && y >= c.Min.Y-slop
}

// X returns the value of x, given in the unit range,
// in the drawing coordinates of this draw area.
// A value of 0, for example, will return the minimum
// x value of the draw area and a value of 1 will
// return the maximum.
func (c *Canvas) X(x float64) vg.Length {
	return vg.Length(x)*(c.Max.X-c.Min.X) + c.Min.X
}

// Y returns the value of x, given in the unit range,
// in the drawing coordinates of this draw area.
// A value of 0, for example, will return the minimum
// y value of the draw area and a value of 1 will
// return the maximum.
func (c *Canvas) Y(y float64) vg.Length {
	return vg.Length(y)*(c.Max.Y-c.Min.Y) + c.Min.Y
}

// Crop returns a new Canvas corresponding to the Canvas
// c with the given lengths added to the minimum
// and maximum x and y values of the Canvas's Rectangle.
// Note that cropping the right and top sides of the canvas
// requires specifying negative values of right and top.
func Crop(c Canvas, left, right, bottom, top vg.Length) Canvas {
	minpt := vg.Point{
		X: c.Min.X + left,
		Y: c.Min.Y + bottom,
	}
	maxpt := vg.Point{
		X: c.Max.X + right,
		Y: c.Max.Y + top,
	}
	return Canvas{
		Canvas:    c,
		Rectangle: vg.Rectangle{Min: minpt, Max: maxpt},
	}
}

// Tiles creates regular subcanvases from a Canvas.
type Tiles struct {
	// Cols and Rows specify the number of rows and columns of tiles.
	Cols, Rows int
	// PadTop, PadBottom, PadRight, and PadLeft specify the padding
	// on the corresponding side of each tile.
	PadTop, PadBottom, PadRight, PadLeft vg.Length
	// PadX and PadY specify the padding between columns and rows
	// of tiles respectively..
	PadX, PadY vg.Length
}

// At returns the subcanvas within c that corresponds to the
// tile at column x, row y.
func (ts Tiles) At(c Canvas, x, y int) Canvas {
	tileH := (c.Max.Y - c.Min.Y - ts.PadTop - ts.PadBottom -
		vg.Length(ts.Rows-1)*ts.PadY) / vg.Length(ts.Rows)
	tileW := (c.Max.X - c.Min.X - ts.PadLeft - ts.PadRight -
		vg.Length(ts.Cols-1)*ts.PadX) / vg.Length(ts.Cols)

	ymax := c.Max.Y - ts.PadTop - vg.Length(y)*(ts.PadY+tileH)
	ymin := ymax - tileH
	xmin := c.Min.X + ts.PadLeft + vg.Length(x)*(ts.PadX+tileW)
	xmax := xmin + tileW

	return Canvas{
		Canvas: vg.Canvas(c),
		Rectangle: vg.Rectangle{
			Min: vg.Point{X: xmin, Y: ymin},
			Max: vg.Point{X: xmax, Y: ymax},
		},
	}
}

// SetLineStyle sets the current line style
func (c *Canvas) SetLineStyle(sty LineStyle) {
	c.SetColor(sty.Color)
	c.SetLineWidth(sty.Width)
	var dashDots []vg.Length
	for _, dash := range sty.Dashes {
		dashDots = append(dashDots, dash)
	}
	c.SetLineDash(dashDots, sty.DashOffs)
}

// StrokeLines draws a line connecting a set of points
// in the given Canvas.
func (c *Canvas) StrokeLines(sty LineStyle, lines ...[]vg.Point) {
	if len(lines) == 0 {
		return
	}

	c.SetLineStyle(sty)

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		var p vg.Path
		p.Move(l[0])
		for _, pt := range l[1:] {
			p.Line(pt)
		}
		c.Stroke(p)
	}
}

// StrokeLine2 draws a line between two points in the given
// Canvas.
func (c *Canvas) StrokeLine2(sty LineStyle, x0, y0, x1, y1 vg.Length) {
	c.StrokeLines(sty, []vg.Point{{x0, y0}, {x1, y1}})
}

// ClipLinesXY returns a slice of lines that
// represent the given line clipped in both
// X and Y directions.
func (c *Canvas) ClipLinesXY(lines ...[]vg.Point) [][]vg.Point {
	return c.ClipLinesY(c.ClipLinesX(lines...)...)
}

// ClipLinesX returns a slice of lines that
// represent the given line clipped in the
// X direction.
func (c *Canvas) ClipLinesX(lines ...[]vg.Point) (clipped [][]vg.Point) {
	var lines1 [][]vg.Point
	for _, line := range lines {
		ls := clipLine(isLeft, vg.Point{X: c.Max.X, Y: c.Min.Y}, vg.Point{X: -1, Y: 0}, line)
		lines1 = append(lines1, ls...)
	}
	for _, line := range lines1 {
		ls := clipLine(isRight, vg.Point{X: c.Min.X, Y: c.Min.Y}, vg.Point{X: 1, Y: 0}, line)
		clipped = append(clipped, ls...)
	}
	return
}

// ClipLinesY returns a slice of lines that
// represent the given line clipped in the
// Y direction.
func (c *Canvas) ClipLinesY(lines ...[]vg.Point) (clipped [][]vg.Point) {
	var lines1 [][]vg.Point
	for _, line := range lines {
		ls := clipLine(isAbove, vg.Point{X: c.Min.X, Y: c.Min.Y}, vg.Point{X: 0, Y: -1}, line)
		lines1 = append(lines1, ls...)
	}
	for _, line := range lines1 {
		ls := clipLine(isBelow, vg.Point{X: c.Min.X, Y: c.Max.Y}, vg.Point{X: 0, Y: 1}, line)
		clipped = append(clipped, ls...)
	}
	return
}

// clipLine performs clipping of a line by a single
// clipping line specified by the norm, clip point,
// and in function.
func clipLine(in func(vg.Point, vg.Point) bool, clip, norm vg.Point, pts []vg.Point) (lines [][]vg.Point) {
	var l []vg.Point
	for i := 1; i < len(pts); i++ {
		cur, next := pts[i-1], pts[i]
		curIn, nextIn := in(cur, clip), in(next, clip)
		switch {
		case curIn && nextIn:
			l = append(l, cur)

		case curIn && !nextIn:
			l = append(l, cur, isect(cur, next, clip, norm))
			lines = append(lines, l)
			l = []vg.Point{}

		case !curIn && !nextIn:
			// do nothing

		default: // !curIn && nextIn
			l = append(l, isect(cur, next, clip, norm))
		}
		if nextIn && i == len(pts)-1 {
			l = append(l, next)
		}
	}
	if len(l) > 1 {
		lines = append(lines, l)
	}
	return
}

// FillPolygon fills a polygon with the given color.
func (c *Canvas) FillPolygon(clr color.Color, pts []vg.Point) {
	if len(pts) == 0 {
		return
	}

	c.SetColor(clr)
	var p vg.Path
	p.Move(pts[0])
	for _, pt := range pts[1:] {
		p.Line(pt)
	}
	p.Close()
	c.Fill(p)
}

// ClipPolygonXY returns a slice of lines that
// represent the given polygon clipped in both
// X and Y directions.
func (c *Canvas) ClipPolygonXY(pts []vg.Point) []vg.Point {
	return c.ClipPolygonY(c.ClipPolygonX(pts))
}

// ClipPolygonX returns a slice of lines that
// represent the given polygon clipped in the
// X direction.
func (c *Canvas) ClipPolygonX(pts []vg.Point) []vg.Point {
	return clipPoly(isLeft, vg.Point{X: c.Max.X, Y: c.Min.Y}, vg.Point{X: -1, Y: 0},
		clipPoly(isRight, vg.Point{X: c.Min.X, Y: c.Min.Y}, vg.Point{X: 1, Y: 0}, pts))
}

// ClipPolygonY returns a slice of lines that
// represent the given polygon clipped in the
// Y direction.
func (c *Canvas) ClipPolygonY(pts []vg.Point) []vg.Point {
	return clipPoly(isBelow, vg.Point{X: c.Min.X, Y: c.Max.Y}, vg.Point{X: 0, Y: 1},
		clipPoly(isAbove, vg.Point{X: c.Min.X, Y: c.Min.Y}, vg.Point{X: 0, Y: -1}, pts))
}

// clipPoly performs clipping of a polygon by a single
// clipping line specified by the norm, clip point,
// and in function.
func clipPoly(in func(vg.Point, vg.Point) bool, clip, norm vg.Point, pts []vg.Point) (clipped []vg.Point) {
	for i := 0; i < len(pts); i++ {
		j := i + 1
		if i == len(pts)-1 {
			j = 0
		}
		cur, next := pts[i], pts[j]
		curIn, nextIn := in(cur, clip), in(next, clip)
		switch {
		case curIn && nextIn:
			clipped = append(clipped, cur)

		case curIn && !nextIn:
			clipped = append(clipped, cur, isect(cur, next, clip, norm))

		case !curIn && !nextIn:
			// do nothing

		default: // !curIn && nextIn
			clipped = append(clipped, isect(cur, next, clip, norm))
		}
	}
	return
}

// slop is some slop for floating point equality
const slop = 3e-8 // ≈ √1⁻¹⁵

func isLeft(p, clip vg.Point) bool {
	return p.X <= clip.X+slop
}

func isRight(p, clip vg.Point) bool {
	return p.X >= clip.X-slop
}

func isBelow(p, clip vg.Point) bool {
	return p.Y <= clip.Y+slop
}

func isAbove(p, clip vg.Point) bool {
	return p.Y >= clip.Y-slop
}

// isect returns the intersection of a line p0→p1 with the
// clipping line specified by the clip point and normal.
func isect(p0, p1, clip, norm vg.Point) vg.Point {
	// t = (norm · (p0 - clip)) / (norm · (p0 - p1))
	t := p0.Sub(clip).Dot(norm) / p0.Sub(p1).Dot(norm)

	// p = p0 + t*(p1 - p0)
	return p1.Sub(p0).Scale(t).Add(p0)
}

// FillText fills lines of text in the draw area.
// pt specifies the location where the text is to be drawn.
func (c *Canvas) FillText(sty TextStyle, pt vg.Point, txt string) {
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

	nl := textNLines(txt)
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
	nl := textNLines(txt)
	if nl == 0 {
		return vg.Length(0)
	}
	e := sty.Font.Extents()
	return e.Height*vg.Length(nl-1) + e.Ascent
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

func max(d ...vg.Length) vg.Length {
	o := vg.Length(math.Inf(-1))
	for _, dd := range d {
		if dd > o {
			o = dd
		}
	}
	return o
}

func min(d ...vg.Length) vg.Length {
	o := vg.Length(math.Inf(1))
	for _, dd := range d {
		if dd < o {
			o = dd
		}
	}
	return o
}

// textNLines returns the number of lines in the text.
func textNLines(txt string) int {
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
