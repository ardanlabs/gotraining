// Copyright 2010 The draw2d Authors. All rights reserved.
// created: 21/11/2010 by Laurent Le Goff

package draw2dbase

import (
	"fmt"
	"image"
	"image/color"

	"github.com/llgcode/draw2d"

	"github.com/golang/freetype/truetype"
)

var DefaultFontData = draw2d.FontData{Name: "luxi", Family: draw2d.FontFamilySans, Style: draw2d.FontStyleNormal}

type StackGraphicContext struct {
	Current *ContextStack
}

type ContextStack struct {
	Tr          draw2d.Matrix
	Path        *draw2d.Path
	LineWidth   float64
	Dash        []float64
	DashOffset  float64
	StrokeColor color.Color
	FillColor   color.Color
	FillRule    draw2d.FillRule
	Cap         draw2d.LineCap
	Join        draw2d.LineJoin
	FontSize    float64
	FontData    draw2d.FontData

	Font *truetype.Font
	// fontSize and dpi are used to calculate scale. scale is the number of
	// 26.6 fixed point units in 1 em.
	Scale float64

	Previous *ContextStack
}

/**
 * Create a new Graphic context from an image
 */
func NewStackGraphicContext() *StackGraphicContext {
	gc := &StackGraphicContext{}
	gc.Current = new(ContextStack)
	gc.Current.Tr = draw2d.NewIdentityMatrix()
	gc.Current.Path = new(draw2d.Path)
	gc.Current.LineWidth = 1.0
	gc.Current.StrokeColor = image.Black
	gc.Current.FillColor = image.White
	gc.Current.Cap = draw2d.RoundCap
	gc.Current.FillRule = draw2d.FillRuleEvenOdd
	gc.Current.Join = draw2d.RoundJoin
	gc.Current.FontSize = 10
	gc.Current.FontData = DefaultFontData
	return gc
}

func (gc *StackGraphicContext) GetMatrixTransform() draw2d.Matrix {
	return gc.Current.Tr
}

func (gc *StackGraphicContext) SetMatrixTransform(Tr draw2d.Matrix) {
	gc.Current.Tr = Tr
}

func (gc *StackGraphicContext) ComposeMatrixTransform(Tr draw2d.Matrix) {
	gc.Current.Tr.Compose(Tr)
}

func (gc *StackGraphicContext) Rotate(angle float64) {
	gc.Current.Tr.Rotate(angle)
}

func (gc *StackGraphicContext) Translate(tx, ty float64) {
	gc.Current.Tr.Translate(tx, ty)
}

func (gc *StackGraphicContext) Scale(sx, sy float64) {
	gc.Current.Tr.Scale(sx, sy)
}

func (gc *StackGraphicContext) SetStrokeColor(c color.Color) {
	gc.Current.StrokeColor = c
}

func (gc *StackGraphicContext) SetFillColor(c color.Color) {
	gc.Current.FillColor = c
}

func (gc *StackGraphicContext) SetFillRule(f draw2d.FillRule) {
	gc.Current.FillRule = f
}

func (gc *StackGraphicContext) SetLineWidth(lineWidth float64) {
	gc.Current.LineWidth = lineWidth
}

func (gc *StackGraphicContext) SetLineCap(cap draw2d.LineCap) {
	gc.Current.Cap = cap
}

func (gc *StackGraphicContext) SetLineJoin(join draw2d.LineJoin) {
	gc.Current.Join = join
}

func (gc *StackGraphicContext) SetLineDash(dash []float64, dashOffset float64) {
	gc.Current.Dash = dash
	gc.Current.DashOffset = dashOffset
}

func (gc *StackGraphicContext) SetFontSize(fontSize float64) {
	gc.Current.FontSize = fontSize
}

func (gc *StackGraphicContext) GetFontSize() float64 {
	return gc.Current.FontSize
}

func (gc *StackGraphicContext) SetFontData(fontData draw2d.FontData) {
	gc.Current.FontData = fontData
}

func (gc *StackGraphicContext) GetFontData() draw2d.FontData {
	return gc.Current.FontData
}

func (gc *StackGraphicContext) BeginPath() {
	gc.Current.Path.Clear()
}

func (gc *StackGraphicContext) GetPath() draw2d.Path {
	return *gc.Current.Path.Copy()
}

func (gc *StackGraphicContext) IsEmpty() bool {
	return gc.Current.Path.IsEmpty()
}

func (gc *StackGraphicContext) LastPoint() (float64, float64) {
	return gc.Current.Path.LastPoint()
}

func (gc *StackGraphicContext) MoveTo(x, y float64) {
	gc.Current.Path.MoveTo(x, y)
}

func (gc *StackGraphicContext) LineTo(x, y float64) {
	gc.Current.Path.LineTo(x, y)
}

func (gc *StackGraphicContext) QuadCurveTo(cx, cy, x, y float64) {
	gc.Current.Path.QuadCurveTo(cx, cy, x, y)
}

func (gc *StackGraphicContext) CubicCurveTo(cx1, cy1, cx2, cy2, x, y float64) {
	gc.Current.Path.CubicCurveTo(cx1, cy1, cx2, cy2, x, y)
}

func (gc *StackGraphicContext) ArcTo(cx, cy, rx, ry, startAngle, angle float64) {
	gc.Current.Path.ArcTo(cx, cy, rx, ry, startAngle, angle)
}

func (gc *StackGraphicContext) Close() {
	gc.Current.Path.Close()
}

func (gc *StackGraphicContext) Save() {
	context := new(ContextStack)
	context.FontSize = gc.Current.FontSize
	context.FontData = gc.Current.FontData
	context.LineWidth = gc.Current.LineWidth
	context.StrokeColor = gc.Current.StrokeColor
	context.FillColor = gc.Current.FillColor
	context.FillRule = gc.Current.FillRule
	context.Dash = gc.Current.Dash
	context.DashOffset = gc.Current.DashOffset
	context.Cap = gc.Current.Cap
	context.Join = gc.Current.Join
	context.Path = gc.Current.Path.Copy()
	context.Font = gc.Current.Font
	context.Scale = gc.Current.Scale
	copy(context.Tr[:], gc.Current.Tr[:])
	context.Previous = gc.Current
	gc.Current = context
}

func (gc *StackGraphicContext) Restore() {
	if gc.Current.Previous != nil {
		oldContext := gc.Current
		gc.Current = gc.Current.Previous
		oldContext.Previous = nil
	}
}

func (gc *StackGraphicContext) GetFontName() string {
	fontData := gc.Current.FontData
	return fmt.Sprintf("%s:%d:%d:%d", fontData.Name, fontData.Family, fontData.Style, gc.Current.FontSize)
}
