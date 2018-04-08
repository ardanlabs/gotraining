// Copyright 2010 The draw2d Authors. All rights reserved.
// created: 21/11/2010 by Laurent Le Goff

package draw2d

import (
	"fmt"
	"math"
)

// PathBuilder describes the interface for path drawing.
type PathBuilder interface {
	// LastPoint returns the current point of the current sub path
	LastPoint() (x, y float64)
	// MoveTo creates a new subpath that start at the specified point
	MoveTo(x, y float64)
	// LineTo adds a line to the current subpath
	LineTo(x, y float64)
	// QuadCurveTo adds a quadratic Bézier curve to the current subpath
	QuadCurveTo(cx, cy, x, y float64)
	// CubicCurveTo adds a cubic Bézier curve to the current subpath
	CubicCurveTo(cx1, cy1, cx2, cy2, x, y float64)
	// ArcTo adds an arc to the current subpath
	ArcTo(cx, cy, rx, ry, startAngle, angle float64)
	// Close creates a line from the current point to the last MoveTo
	// point (if not the same) and mark the path as closed so the
	// first and last lines join nicely.
	Close()
}

// PathCmp represents component of a path
type PathCmp int

const (
	// MoveToCmp is a MoveTo component in a Path
	MoveToCmp PathCmp = iota
	// LineToCmp is a LineTo component in a Path
	LineToCmp
	// QuadCurveToCmp is a QuadCurveTo component in a Path
	QuadCurveToCmp
	// CubicCurveToCmp is a CubicCurveTo component in a Path
	CubicCurveToCmp
	// ArcToCmp is a ArcTo component in a Path
	ArcToCmp
	// CloseCmp is a ArcTo component in a Path
	CloseCmp
)

// Path stores points
type Path struct {
	// Components is a slice of PathCmp in a Path and mark the role of each points in the Path
	Components []PathCmp
	// Points are combined with Components to have a specific role in the path
	Points []float64
	// Last Point of the Path
	x, y float64
}

func (p *Path) appendToPath(cmd PathCmp, points ...float64) {
	p.Components = append(p.Components, cmd)
	p.Points = append(p.Points, points...)
}

// LastPoint returns the current point of the current path
func (p *Path) LastPoint() (x, y float64) {
	return p.x, p.y
}

// MoveTo starts a new path at (x, y) position
func (p *Path) MoveTo(x, y float64) {
	p.appendToPath(MoveToCmp, x, y)
	p.x = x
	p.y = y
}

// LineTo adds a line to the current path
func (p *Path) LineTo(x, y float64) {
	if len(p.Components) == 0 { //special case when no move has been done
		p.MoveTo(x, y)
	} else {
		p.appendToPath(LineToCmp, x, y)
	}
	p.x = x
	p.y = y
}

// QuadCurveTo adds a quadratic bezier curve to the current path
func (p *Path) QuadCurveTo(cx, cy, x, y float64) {
	if len(p.Components) == 0 { //special case when no move has been done
		p.MoveTo(x, y)
	} else {
		p.appendToPath(QuadCurveToCmp, cx, cy, x, y)
	}
	p.x = x
	p.y = y
}

// CubicCurveTo adds a cubic bezier curve to the current path
func (p *Path) CubicCurveTo(cx1, cy1, cx2, cy2, x, y float64) {
	if len(p.Components) == 0 { //special case when no move has been done
		p.MoveTo(x, y)
	} else {
		p.appendToPath(CubicCurveToCmp, cx1, cy1, cx2, cy2, x, y)
	}
	p.x = x
	p.y = y
}

// ArcTo adds an arc to the path
func (p *Path) ArcTo(cx, cy, rx, ry, startAngle, angle float64) {
	endAngle := startAngle + angle
	clockWise := true
	if angle < 0 {
		clockWise = false
	}
	// normalize
	if clockWise {
		for endAngle < startAngle {
			endAngle += math.Pi * 2.0
		}
	} else {
		for startAngle < endAngle {
			startAngle += math.Pi * 2.0
		}
	}
	startX := cx + math.Cos(startAngle)*rx
	startY := cy + math.Sin(startAngle)*ry
	if len(p.Components) > 0 {
		p.LineTo(startX, startY)
	} else {
		p.MoveTo(startX, startY)
	}
	p.appendToPath(ArcToCmp, cx, cy, rx, ry, startAngle, angle)
	p.x = cx + math.Cos(endAngle)*rx
	p.y = cy + math.Sin(endAngle)*ry
}

// Close closes the current path
func (p *Path) Close() {
	p.appendToPath(CloseCmp)
}

// Copy make a clone of the current path and return it
func (p *Path) Copy() (dest *Path) {
	dest = new(Path)
	dest.Components = make([]PathCmp, len(p.Components))
	copy(dest.Components, p.Components)
	dest.Points = make([]float64, len(p.Points))
	copy(dest.Points, p.Points)
	dest.x, dest.y = p.x, p.y
	return dest
}

// Clear reset the path
func (p *Path) Clear() {
	p.Components = p.Components[0:0]
	p.Points = p.Points[0:0]
	return
}

// IsEmpty returns true if the path is empty
func (p *Path) IsEmpty() bool {
	return len(p.Components) == 0
}

// String returns a debug text view of the path
func (p *Path) String() string {
	s := ""
	j := 0
	for _, cmd := range p.Components {
		switch cmd {
		case MoveToCmp:
			s += fmt.Sprintf("MoveTo: %f, %f\n", p.Points[j], p.Points[j+1])
			j = j + 2
		case LineToCmp:
			s += fmt.Sprintf("LineTo: %f, %f\n", p.Points[j], p.Points[j+1])
			j = j + 2
		case QuadCurveToCmp:
			s += fmt.Sprintf("QuadCurveTo: %f, %f, %f, %f\n", p.Points[j], p.Points[j+1], p.Points[j+2], p.Points[j+3])
			j = j + 4
		case CubicCurveToCmp:
			s += fmt.Sprintf("CubicCurveTo: %f, %f, %f, %f, %f, %f\n", p.Points[j], p.Points[j+1], p.Points[j+2], p.Points[j+3], p.Points[j+4], p.Points[j+5])
			j = j + 6
		case ArcToCmp:
			s += fmt.Sprintf("ArcTo: %f, %f, %f, %f, %f, %f\n", p.Points[j], p.Points[j+1], p.Points[j+2], p.Points[j+3], p.Points[j+4], p.Points[j+5])
			j = j + 6
		case CloseCmp:
			s += "Close\n"
		}
	}
	return s
}

// Returns new Path with flipped y axes
func (path *Path) VerticalFlip() *Path {
	p := path.Copy()
	j := 0
	for _, cmd := range p.Components {
		switch cmd {
		case MoveToCmp, LineToCmp:
			p.Points[j+1] = -p.Points[j+1]
			j = j + 2
		case QuadCurveToCmp:
			p.Points[j+1] = -p.Points[j+1]
			p.Points[j+3] = -p.Points[j+3]
			j = j + 4
		case CubicCurveToCmp:
			p.Points[j+1] = -p.Points[j+1]
			p.Points[j+3] = -p.Points[j+3]
			p.Points[j+5] = -p.Points[j+5]
			j = j + 6
		case ArcToCmp:
			p.Points[j+1] = -p.Points[j+1]
			p.Points[j+3] = -p.Points[j+3]
			p.Points[j+4] = -p.Points[j+4] // start angle
			p.Points[j+5] = -p.Points[j+5] // angle
			j = j + 6
		case CloseCmp:
		}
	}
	p.y = -p.y
	return p
}
