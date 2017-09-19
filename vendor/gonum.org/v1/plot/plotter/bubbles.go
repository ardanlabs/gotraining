// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter

import (
	"errors"
	"image/color"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

// Bubbles implements the Plotter interface, drawing
// a bubble plot of x, y, z triples where the z value
// determines the radius of the bubble.
type Bubbles struct {
	XYZs

	// Color is the color of the bubbles.
	color.Color

	// MinRadius and MaxRadius give the minimum
	// and maximum bubble radius respectively.
	// The radii of each bubble is interpolated linearly
	// between these two values.
	MinRadius, MaxRadius vg.Length

	// MinZ and MaxZ are the minimum and
	// maximum Z values from the data.
	MinZ, MaxZ float64
}

// NewBubbles creates as new bubble plot plotter for
// the given data, with a minimum and maximum
// bubble radius.
func NewBubbles(xyz XYZer, min, max vg.Length) (*Bubbles, error) {
	cpy, err := CopyXYZs(xyz)
	if err != nil {
		return nil, err
	}
	if min > max {
		return nil, errors.New("Min bubble radius is greater than the max radius")
	}
	minz := cpy[0].Z
	maxz := cpy[0].Z
	for _, d := range cpy {
		minz = math.Min(minz, d.Z)
		maxz = math.Max(maxz, d.Z)
	}
	return &Bubbles{
		XYZs:      cpy,
		MinRadius: min,
		MaxRadius: max,
		MinZ:      minz,
		MaxZ:      maxz,
	}, nil
}

// Plot implements the Plot method of the plot.Plotter interface.
func (bs *Bubbles) Plot(c draw.Canvas, plt *plot.Plot) {
	trX, trY := plt.Transforms(&c)

	c.SetColor(bs.Color)

	for _, d := range bs.XYZs {
		x := trX(d.X)
		y := trY(d.Y)
		pt := vg.Point{X: x, Y: y}
		if !c.Contains(pt) {
			continue
		}

		rad := bs.radius(d.Z)

		// draw a circle centered at x, y
		var p vg.Path
		p.Move(vg.Point{X: x + rad, Y: y})
		p.Arc(pt, rad, 0, 2*math.Pi)
		p.Close()
		c.Fill(p)
	}
}

// radius returns the radius of a bubble by linear interpolation.
func (bs *Bubbles) radius(z float64) vg.Length {
	rng := bs.MaxRadius - bs.MinRadius
	if bs.MaxZ == bs.MinZ {
		return rng/2 + bs.MinRadius
	}
	d := (z - bs.MinZ) / (bs.MaxZ - bs.MinZ)
	return vg.Length(d)*rng + bs.MinRadius
}

// DataRange implements the DataRange method
// of the plot.DataRanger interface.
func (bs *Bubbles) DataRange() (xmin, xmax, ymin, ymax float64) {
	return XYRange(XYValues{bs.XYZs})
}

// GlyphBoxes implements the GlyphBoxes method
// of the plot.GlyphBoxer interface.
func (bs *Bubbles) GlyphBoxes(plt *plot.Plot) []plot.GlyphBox {
	boxes := make([]plot.GlyphBox, len(bs.XYZs))
	for i, d := range bs.XYZs {
		boxes[i].X = plt.X.Norm(d.X)
		boxes[i].Y = plt.Y.Norm(d.Y)
		r := bs.radius(d.Z)
		boxes[i].Rectangle = vg.Rectangle{
			Min: vg.Point{X: -r, Y: -r},
			Max: vg.Point{X: +r, Y: +r},
		}
	}
	return boxes
}
