// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

// Line implements the Plotter interface, drawing a line.
type Line struct {
	// XYs is a copy of the points for this line.
	XYs

	// LineStyle is the style of the line connecting
	// the points.
	draw.LineStyle

	// ShadeColor is the color of the shaded area.
	ShadeColor *color.Color
}

// NewLine returns a Line that uses the default line style and
// does not draw glyphs.
func NewLine(xys XYer) (*Line, error) {
	data, err := CopyXYs(xys)
	if err != nil {
		return nil, err
	}
	return &Line{
		XYs:       data,
		LineStyle: DefaultLineStyle,
	}, nil
}

// Plot draws the Line, implementing the plot.Plotter
// interface.
func (pts *Line) Plot(c draw.Canvas, plt *plot.Plot) {
	trX, trY := plt.Transforms(&c)
	ps := make([]vg.Point, len(pts.XYs))

	for i, p := range pts.XYs {
		ps[i].X = trX(p.X)
		ps[i].Y = trY(p.Y)
	}

	if pts.ShadeColor != nil && len(ps) > 0 {
		c.SetColor(*pts.ShadeColor)
		minY := trY(plt.Y.Min)
		var pa vg.Path
		pa.Move(vg.Point{X: ps[0].X, Y: minY})
		for i := range pts.XYs {
			pa.Line(ps[i])
		}
		pa.Line(vg.Point{X: ps[len(pts.XYs)-1].X, Y: minY})
		pa.Close()
		c.Fill(pa)
	}

	c.StrokeLines(pts.LineStyle, c.ClipLinesXY(ps)...)
}

// DataRange returns the minimum and maximum
// x and y values, implementing the plot.DataRanger
// interface.
func (pts *Line) DataRange() (xmin, xmax, ymin, ymax float64) {
	return XYRange(pts)
}

// Thumbnail the thumbnail for the Line,
// implementing the plot.Thumbnailer interface.
func (pts *Line) Thumbnail(c *draw.Canvas) {
	if pts.ShadeColor != nil {
		points := []vg.Point{
			{c.Min.X, c.Min.Y},
			{c.Min.X, c.Max.Y},
			{c.Max.X, c.Max.Y},
			{c.Max.X, c.Min.Y},
		}
		poly := c.ClipPolygonY(points)
		c.FillPolygon(*pts.ShadeColor, poly)

		points = append(points, vg.Point{X: c.Min.X, Y: c.Min.Y})
	} else {
		y := c.Center().Y
		c.StrokeLine2(pts.LineStyle, c.Min.X, y, c.Max.X, y)
	}
}

// NewLinePoints returns both a Line and a
// Points for the given point data.
func NewLinePoints(xys XYer) (*Line, *Scatter, error) {
	s, err := NewScatter(xys)
	if err != nil {
		return nil, nil, err
	}
	l := &Line{
		XYs:       s.XYs,
		LineStyle: DefaultLineStyle,
	}
	return l, s, nil
}
