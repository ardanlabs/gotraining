// Copyright 2010 The draw2d Authors. All rights reserved.
// created: 13/12/2010 by Laurent Le Goff

// Package draw2d is a pure go 2D vector graphics library with support
// for multiple output devices such as images (draw2d), pdf documents
// (draw2dpdf) and opengl (draw2dgl), which can also be used on the
// google app engine. It can be used as a pure go Cairo alternative.
// draw2d is released under the BSD license.
//
// Features
//
// Operations in draw2d include stroking and filling polygons, arcs,
// Bézier curves, drawing images and text rendering with truetype fonts.
// All drawing operations can be transformed by affine transformations
// (scale, rotation, translation).
//
// Package draw2d follows the conventions of http://www.w3.org/TR/2dcontext for coordinate system, angles, etc...
//
// Installation
//
// To install or update the package draw2d on your system, run:
//   go get -u github.com/llgcode/draw2d
//
// Quick Start
//
// Package draw2d itself provides a graphic context that can draw vector
// graphics and text on an image canvas. The following Go code
// generates a simple drawing and saves it to an image file:
//		package main
//
// 		import (
// 			"github.com/llgcode/draw2d/draw2dimg"
// 			"image"
// 			"image/color"
// 		)
//
// 		func main() {
// 			// Initialize the graphic context on an RGBA image
// 			dest := image.NewRGBA(image.Rect(0, 0, 297, 210.0))
// 			gc := draw2dimg.NewGraphicContext(dest)
//
// 			// Set some properties
// 			gc.SetFillColor(color.RGBA{0x44, 0xff, 0x44, 0xff})
// 			gc.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
// 			gc.SetLineWidth(5)
//
// 			// Draw a closed shape
// 			gc.MoveTo(10, 10) // should always be called first for a new path
// 			gc.LineTo(100, 50)
// 			gc.QuadCurveTo(100, 10, 10, 10)
// 			gc.Close()
// 			gc.FillStroke()
//
// 			// Save to file
// 			draw2dimg.SaveToPngFile("hello.png", dest)
// 		}
//
//
// There are more examples here:
// https://github.com/llgcode/draw2d/tree/master/samples
//
// Drawing on pdf documents is provided by the draw2dpdf package.
// Drawing on opengl is provided by the draw2dgl package.
// See subdirectories at the bottom of this page.
//
// Testing
//
// The samples are run as tests from the root package folder `draw2d` by:
//   go test ./...
//
// Or if you want to run with test coverage:
//   go test -cover ./... | grep -v "no test"
//
// This will generate output by the different backends in the output folder.
//
// Acknowledgments
//
// Laurent Le Goff wrote this library, inspired by Postscript and
// HTML5 canvas. He implemented the image and opengl backend with the
// freetype-go package. Also he created a pure go Postscript
// interpreter, which can read postscript images and draw to a draw2d
// graphic context (https://github.com/llgcode/ps). Stani Michiels
// implemented the pdf backend with the gofpdf package.
//
// Packages using draw2d
//
// - https://github.com/llgcode/ps: Postscript interpreter written in Go
//
// - https://github.com/gonum/plot: drawing plots in Go
//
// - https://github.com/muesli/smartcrop: content aware image cropping
//
// - https://github.com/peterhellberg/karta: drawing Voronoi diagrams
//
// - https://github.com/vdobler/chart: basic charts in Go
package draw2d

import "image/color"

// FillRule defines the type for fill rules
type FillRule int

const (
	// FillRuleEvenOdd determines the "insideness" of a point in the shape
	// by drawing a ray from that point to infinity in any direction
	// and counting the number of path segments from the given shape that the ray crosses.
	// If this number is odd, the point is inside; if even, the point is outside.
	FillRuleEvenOdd FillRule = iota
	// FillRuleWinding determines the "insideness" of a point in the shape
	// by drawing a ray from that point to infinity in any direction
	// and then examining the places where a segment of the shape crosses the ray.
	// Starting with a count of zero, add one each time a path segment crosses
	// the ray from left to right and subtract one each time
	// a path segment crosses the ray from right to left. After counting the crossings,
	// if the result is zero then the point is outside the path. Otherwise, it is inside.
	FillRuleWinding
)

// LineCap is the style of line extremities
type LineCap int

const (
	// RoundCap defines a rounded shape at the end of the line
	RoundCap LineCap = iota
	// ButtCap defines a squared shape exactly at the end of the line
	ButtCap
	// SquareCap defines a squared shape at the end of the line
	SquareCap
)

// LineJoin is the style of segments joint
type LineJoin int

const (
	// BevelJoin represents cut segments joint
	BevelJoin LineJoin = iota
	// RoundJoin represents rounded segments joint
	RoundJoin
	// MiterJoin represents peaker segments joint
	MiterJoin
)

// StrokeStyle keeps stroke style attributes
// that is used by the Stroke method of a Drawer
type StrokeStyle struct {
	// Color defines the color of stroke
	Color color.Color
	// Line width
	Width float64
	// Line cap style rounded, butt or square
	LineCap LineCap
	// Line join style bevel, round or miter
	LineJoin LineJoin
	// offset of the first dash
	DashOffset float64
	// array represented dash length pair values are plain dash and impair are space between dash
	// if empty display plain line
	Dash []float64
}

// SolidFillStyle define style attributes for a solid fill style
type SolidFillStyle struct {
	// Color defines the line color
	Color color.Color
	// FillRule defines the file rule to used
	FillRule FillRule
}

// Valign Vertical Alignment of the text
type Valign int

const (
	// ValignTop top align text
	ValignTop Valign = iota
	// ValignCenter centered text
	ValignCenter
	// ValignBottom bottom aligned text
	ValignBottom
	// ValignBaseline align text with the baseline of the font
	ValignBaseline
)

// Halign Horizontal Alignment of the text
type Halign int

const (
	// HalignLeft Horizontally align to left
	HalignLeft = iota
	// HalignCenter Horizontally align to center
	HalignCenter
	// HalignRight Horizontally align to right
	HalignRight
)

// TextStyle describe text property
type TextStyle struct {
	// Color defines the color of text
	Color color.Color
	// Size font size
	Size float64
	// The font to use
	Font FontData
	// Horizontal Alignment of the text
	Halign Halign
	// Vertical Alignment of the text
	Valign Valign
}

// ScalingPolicy is a constant to define how to scale an image
type ScalingPolicy int

const (
	// ScalingNone no scaling applied
	ScalingNone ScalingPolicy = iota
	// ScalingStretch the image is stretched so that its width and height are exactly the given width and height
	ScalingStretch
	// ScalingWidth the image is scaled so that its width is exactly the given width
	ScalingWidth
	// ScalingHeight the image is scaled so that its height is exactly the given height
	ScalingHeight
	// ScalingFit the image is scaled to the largest scale that allow the image to fit within a rectangle width x height
	ScalingFit
	// ScalingSameArea the image is scaled so that its area is exactly the area of the given rectangle width x height
	ScalingSameArea
	// ScalingFill the image is scaled to the smallest scale that allow the image to fully cover a rectangle width x height
	ScalingFill
)

// ImageScaling style attributes used to display the image
type ImageScaling struct {
	// Horizontal Alignment of the image
	Halign Halign
	// Vertical Alignment of the image
	Valign Valign
	// Width Height used by scaling policy
	Width, Height float64
	// ScalingPolicy defines the scaling policy to applied to the image
	ScalingPolicy ScalingPolicy
}
