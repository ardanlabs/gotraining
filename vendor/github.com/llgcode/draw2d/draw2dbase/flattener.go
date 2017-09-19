// Copyright 2010 The draw2d Authors. All rights reserved.
// created: 06/12/2010 by Laurent Le Goff

package draw2dbase

import (
	"github.com/llgcode/draw2d"
)

// Liner receive segment definition
type Liner interface {
	// LineTo Draw a line from the current position to the point (x, y)
	LineTo(x, y float64)
}

// Flattener receive segment definition
type Flattener interface {
	// MoveTo Start a New line from the point (x, y)
	MoveTo(x, y float64)
	// LineTo Draw a line from the current position to the point (x, y)
	LineTo(x, y float64)
	// LineJoin use Round, Bevel or miter to join points
	LineJoin()
	// Close add the most recent starting point to close the path to create a polygon
	Close()
	// End mark the current line as finished so we can draw caps
	End()
}

// Flatten convert curves into straight segments keeping join segments info
func Flatten(path *draw2d.Path, flattener Flattener, scale float64) {
	// First Point
	var startX, startY float64 = 0, 0
	// Current Point
	var x, y float64 = 0, 0
	i := 0
	for _, cmp := range path.Components {
		switch cmp {
		case draw2d.MoveToCmp:
			x, y = path.Points[i], path.Points[i+1]
			startX, startY = x, y
			if i != 0 {
				flattener.End()
			}
			flattener.MoveTo(x, y)
			i += 2
		case draw2d.LineToCmp:
			x, y = path.Points[i], path.Points[i+1]
			flattener.LineTo(x, y)
			flattener.LineJoin()
			i += 2
		case draw2d.QuadCurveToCmp:
			TraceQuad(flattener, path.Points[i-2:], 0.5)
			x, y = path.Points[i+2], path.Points[i+3]
			flattener.LineTo(x, y)
			i += 4
		case draw2d.CubicCurveToCmp:
			TraceCubic(flattener, path.Points[i-2:], 0.5)
			x, y = path.Points[i+4], path.Points[i+5]
			flattener.LineTo(x, y)
			i += 6
		case draw2d.ArcToCmp:
			x, y = TraceArc(flattener, path.Points[i], path.Points[i+1], path.Points[i+2], path.Points[i+3], path.Points[i+4], path.Points[i+5], scale)
			flattener.LineTo(x, y)
			i += 6
		case draw2d.CloseCmp:
			flattener.LineTo(startX, startY)
			flattener.Close()
		}
	}
	flattener.End()
}

// Transformer apply the Matrix transformation tr
type Transformer struct {
	Tr        draw2d.Matrix
	Flattener Flattener
}

func (t Transformer) MoveTo(x, y float64) {
	u := x*t.Tr[0] + y*t.Tr[2] + t.Tr[4]
	v := x*t.Tr[1] + y*t.Tr[3] + t.Tr[5]
	t.Flattener.MoveTo(u, v)
}

func (t Transformer) LineTo(x, y float64) {
	u := x*t.Tr[0] + y*t.Tr[2] + t.Tr[4]
	v := x*t.Tr[1] + y*t.Tr[3] + t.Tr[5]
	t.Flattener.LineTo(u, v)
}

func (t Transformer) LineJoin() {
	t.Flattener.LineJoin()
}

func (t Transformer) Close() {
	t.Flattener.Close()
}

func (t Transformer) End() {
	t.Flattener.End()
}

type SegmentedPath struct {
	Points []float64
}

func (p *SegmentedPath) MoveTo(x, y float64) {
	p.Points = append(p.Points, x, y)
	// TODO need to mark this point as moveto
}

func (p *SegmentedPath) LineTo(x, y float64) {
	p.Points = append(p.Points, x, y)
}

func (p *SegmentedPath) LineJoin() {
	// TODO need to mark the current point as linejoin
}

func (p *SegmentedPath) Close() {
	// TODO Close
}

func (p *SegmentedPath) End() {
	// Nothing to do
}
