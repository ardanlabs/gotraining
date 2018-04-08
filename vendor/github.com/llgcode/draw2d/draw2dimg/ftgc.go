// Copyright 2010 The draw2d Authors. All rights reserved.
// created: 21/11/2010 by Laurent Le Goff

package draw2dimg

import (
	"image"
	"image/color"
	"log"
	"math"

	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dbase"

	"github.com/golang/freetype/raster"
	"github.com/golang/freetype/truetype"

	"golang.org/x/image/draw"
	"golang.org/x/image/font"
	"golang.org/x/image/math/f64"
	"golang.org/x/image/math/fixed"
)

// Painter implements the freetype raster.Painter and has a SetColor method like the RGBAPainter
type Painter interface {
	raster.Painter
	SetColor(color color.Color)
}

// GraphicContext is the implementation of draw2d.GraphicContext for a raster image
type GraphicContext struct {
	*draw2dbase.StackGraphicContext
	img              draw.Image
	painter          Painter
	fillRasterizer   *raster.Rasterizer
	strokeRasterizer *raster.Rasterizer
	FontCache		 draw2d.FontCache
	glyphCache       draw2dbase.GlyphCache
	glyphBuf         *truetype.GlyphBuf
	DPI              int
}

// ImageFilter defines the type of filter to use
type ImageFilter int

const (
	// LinearFilter defines a linear filter
	LinearFilter ImageFilter = iota
	// BilinearFilter defines a bilinear filter
	BilinearFilter
	// BicubicFilter defines a bicubic filter
	BicubicFilter
)

// NewGraphicContext creates a new Graphic context from an image.
func NewGraphicContext(img draw.Image) *GraphicContext {

	var painter Painter
	switch selectImage := img.(type) {
	case *image.RGBA:
		painter = raster.NewRGBAPainter(selectImage)
	default:
		panic("Image type not supported")
	}
	return NewGraphicContextWithPainter(img, painter)
}

// NewGraphicContextWithPainter creates a new Graphic context from an image and a Painter (see Freetype-go)
func NewGraphicContextWithPainter(img draw.Image, painter Painter) *GraphicContext {
	width, height := img.Bounds().Dx(), img.Bounds().Dy()
	dpi := 92
	gc := &GraphicContext{
		draw2dbase.NewStackGraphicContext(),
		img,
		painter,
		raster.NewRasterizer(width, height),
		raster.NewRasterizer(width, height),
		draw2d.GetGlobalFontCache(),
		draw2dbase.NewGlyphCache(),
		&truetype.GlyphBuf{},
		dpi,
	}
	return gc
}

// GetDPI returns the resolution of the Image GraphicContext
func (gc *GraphicContext) GetDPI() int {
	return gc.DPI
}

// Clear fills the current canvas with a default transparent color
func (gc *GraphicContext) Clear() {
	width, height := gc.img.Bounds().Dx(), gc.img.Bounds().Dy()
	gc.ClearRect(0, 0, width, height)
}

// ClearRect fills the current canvas with a default transparent color at the specified rectangle
func (gc *GraphicContext) ClearRect(x1, y1, x2, y2 int) {
	imageColor := image.NewUniform(gc.Current.FillColor)
	draw.Draw(gc.img, image.Rect(x1, y1, x2, y2), imageColor, image.ZP, draw.Over)
}

// DrawImage draws an image into dest using an affine transformation matrix, an op and a filter
func DrawImage(src image.Image, dest draw.Image, tr draw2d.Matrix, op draw.Op, filter ImageFilter) {
	var transformer draw.Transformer
	switch filter {
	case LinearFilter:
		transformer = draw.NearestNeighbor
	case BilinearFilter:
		transformer = draw.BiLinear
	case BicubicFilter:
		transformer = draw.CatmullRom
	}
	transformer.Transform(dest, f64.Aff3{tr[0], tr[1], tr[4], tr[2], tr[3], tr[5]}, src, src.Bounds(), op, nil)
}

// DrawImage draws the raster image in the current canvas
func (gc *GraphicContext) DrawImage(img image.Image) {
	DrawImage(img, gc.img, gc.Current.Tr, draw.Over, BilinearFilter)
}

// FillString draws the text at point (0, 0)
func (gc *GraphicContext) FillString(text string) (width float64) {
	return gc.FillStringAt(text, 0, 0)
}

// FillStringAt draws the text at the specified point (x, y)
func (gc *GraphicContext) FillStringAt(text string, x, y float64) (width float64) {
	f, err := gc.loadCurrentFont()
	if err != nil {
		log.Println(err)
		return 0.0
	}
	startx := x
	prev, hasPrev := truetype.Index(0), false
	fontName := gc.GetFontName()
	for _, r := range text {
		index := f.Index(r)
		if hasPrev {
			x += fUnitsToFloat64(f.Kern(fixed.Int26_6(gc.Current.Scale), prev, index))
		}
		glyph := gc.glyphCache.Fetch(gc, fontName, r)
		x += glyph.Fill(gc, x, y)
		prev, hasPrev = index, true
	}
	return x - startx
}

// StrokeString draws the contour of the text at point (0, 0)
func (gc *GraphicContext) StrokeString(text string) (width float64) {
	return gc.StrokeStringAt(text, 0, 0)
}

// StrokeStringAt draws the contour of the text at point (x, y)
func (gc *GraphicContext) StrokeStringAt(text string, x, y float64) (width float64) {
	f, err := gc.loadCurrentFont()
	if err != nil {
		log.Println(err)
		return 0.0
	}
	startx := x
	prev, hasPrev := truetype.Index(0), false
	fontName := gc.GetFontName()
	for _, r := range text {
		index := f.Index(r)
		if hasPrev {
			x += fUnitsToFloat64(f.Kern(fixed.Int26_6(gc.Current.Scale), prev, index))
		}
		glyph := gc.glyphCache.Fetch(gc, fontName, r)
		x += glyph.Stroke(gc, x, y)
		prev, hasPrev = index, true
	}
	return x - startx
}

func (gc *GraphicContext) loadCurrentFont() (*truetype.Font, error) {
	font, err := gc.FontCache.Load(gc.Current.FontData)
	if err != nil {
		font, err = gc.FontCache.Load(draw2dbase.DefaultFontData)
	}
	if font != nil {
		gc.SetFont(font)
		gc.SetFontSize(gc.Current.FontSize)
	}
	return font, err
}

// p is a truetype.Point measured in FUnits and positive Y going upwards.
// The returned value is the same thing measured in floating point and positive Y
// going downwards.

func (gc *GraphicContext) drawGlyph(glyph truetype.Index, dx, dy float64) error {
	if err := gc.glyphBuf.Load(gc.Current.Font, fixed.Int26_6(gc.Current.Scale), glyph, font.HintingNone); err != nil {
		return err
	}
	e0 := 0
	for _, e1 := range gc.glyphBuf.Ends {
		DrawContour(gc, gc.glyphBuf.Points[e0:e1], dx, dy)
		e0 = e1
	}
	return nil
}

// CreateStringPath creates a path from the string s at x, y, and returns the string width.
// The text is placed so that the left edge of the em square of the first character of s
// and the baseline intersect at x, y. The majority of the affected pixels will be
// above and to the right of the point, but some may be below or to the left.
// For example, drawing a string that starts with a 'J' in an italic font may
// affect pixels below and left of the point.
func (gc *GraphicContext) CreateStringPath(s string, x, y float64) float64 {
	f, err := gc.loadCurrentFont()
	if err != nil {
		log.Println(err)
		return 0.0
	}
	startx := x
	prev, hasPrev := truetype.Index(0), false
	for _, rune := range s {
		index := f.Index(rune)
		if hasPrev {
			x += fUnitsToFloat64(f.Kern(fixed.Int26_6(gc.Current.Scale), prev, index))
		}
		err := gc.drawGlyph(index, x, y)
		if err != nil {
			log.Println(err)
			return startx - x
		}
		x += fUnitsToFloat64(f.HMetric(fixed.Int26_6(gc.Current.Scale), index).AdvanceWidth)
		prev, hasPrev = index, true
	}
	return x - startx
}

// GetStringBounds returns the approximate pixel bounds of the string s at x, y.
// The the left edge of the em square of the first character of s
// and the baseline intersect at 0, 0 in the returned coordinates.
// Therefore the top and left coordinates may well be negative.
func (gc *GraphicContext) GetStringBounds(s string) (left, top, right, bottom float64) {
	f, err := gc.loadCurrentFont()
	if err != nil {
		log.Println(err)
		return 0, 0, 0, 0
	}
	top, left, bottom, right = 10e6, 10e6, -10e6, -10e6
	cursor := 0.0
	prev, hasPrev := truetype.Index(0), false
	for _, rune := range s {
		index := f.Index(rune)
		if hasPrev {
			cursor += fUnitsToFloat64(f.Kern(fixed.Int26_6(gc.Current.Scale), prev, index))
		}
		if err := gc.glyphBuf.Load(gc.Current.Font, fixed.Int26_6(gc.Current.Scale), index, font.HintingNone); err != nil {
			log.Println(err)
			return 0, 0, 0, 0
		}
		e0 := 0
		for _, e1 := range gc.glyphBuf.Ends {
			ps := gc.glyphBuf.Points[e0:e1]
			for _, p := range ps {
				x, y := pointToF64Point(p)
				top = math.Min(top, y)
				bottom = math.Max(bottom, y)
				left = math.Min(left, x+cursor)
				right = math.Max(right, x+cursor)
			}
		}
		cursor += fUnitsToFloat64(f.HMetric(fixed.Int26_6(gc.Current.Scale), index).AdvanceWidth)
		prev, hasPrev = index, true
	}
	return left, top, right, bottom
}

// recalc recalculates scale and bounds values from the font size, screen
// resolution and font metrics, and invalidates the glyph cache.
func (gc *GraphicContext) recalc() {
	gc.Current.Scale = gc.Current.FontSize * float64(gc.DPI) * (64.0 / 72.0)
}

// SetDPI sets the screen resolution in dots per inch.
func (gc *GraphicContext) SetDPI(dpi int) {
	gc.DPI = dpi
	gc.recalc()
}

// SetFont sets the font used to draw text.
func (gc *GraphicContext) SetFont(font *truetype.Font) {
	gc.Current.Font = font
}

// SetFontSize sets the font size in points (as in ``a 12 point font'').
func (gc *GraphicContext) SetFontSize(fontSize float64) {
	gc.Current.FontSize = fontSize
	gc.recalc()
}

func (gc *GraphicContext) paint(rasterizer *raster.Rasterizer, color color.Color) {
	gc.painter.SetColor(color)
	rasterizer.Rasterize(gc.painter)
	rasterizer.Clear()
	gc.Current.Path.Clear()
}

// Stroke strokes the paths with the color specified by SetStrokeColor
func (gc *GraphicContext) Stroke(paths ...*draw2d.Path) {
	paths = append(paths, gc.Current.Path)
	gc.strokeRasterizer.UseNonZeroWinding = true

	stroker := draw2dbase.NewLineStroker(gc.Current.Cap, gc.Current.Join, draw2dbase.Transformer{Tr: gc.Current.Tr, Flattener: FtLineBuilder{Adder: gc.strokeRasterizer}})
	stroker.HalfLineWidth = gc.Current.LineWidth / 2

	var liner draw2dbase.Flattener
	if gc.Current.Dash != nil && len(gc.Current.Dash) > 0 {
		liner = draw2dbase.NewDashConverter(gc.Current.Dash, gc.Current.DashOffset, stroker)
	} else {
		liner = stroker
	}
	for _, p := range paths {
		draw2dbase.Flatten(p, liner, gc.Current.Tr.GetScale())
	}

	gc.paint(gc.strokeRasterizer, gc.Current.StrokeColor)
}

// Fill fills the paths with the color specified by SetFillColor
func (gc *GraphicContext) Fill(paths ...*draw2d.Path) {
	paths = append(paths, gc.Current.Path)
	gc.fillRasterizer.UseNonZeroWinding = gc.Current.FillRule == draw2d.FillRuleWinding

	/**** first method ****/
	flattener := draw2dbase.Transformer{Tr: gc.Current.Tr, Flattener: FtLineBuilder{Adder: gc.fillRasterizer}}
	for _, p := range paths {
		draw2dbase.Flatten(p, flattener, gc.Current.Tr.GetScale())
	}

	gc.paint(gc.fillRasterizer, gc.Current.FillColor)
}

// FillStroke first fills the paths and than strokes them
func (gc *GraphicContext) FillStroke(paths ...*draw2d.Path) {
	paths = append(paths, gc.Current.Path)
	gc.fillRasterizer.UseNonZeroWinding = gc.Current.FillRule == draw2d.FillRuleWinding
	gc.strokeRasterizer.UseNonZeroWinding = true

	flattener := draw2dbase.Transformer{Tr: gc.Current.Tr, Flattener: FtLineBuilder{Adder: gc.fillRasterizer}}

	stroker := draw2dbase.NewLineStroker(gc.Current.Cap, gc.Current.Join, draw2dbase.Transformer{Tr: gc.Current.Tr, Flattener: FtLineBuilder{Adder: gc.strokeRasterizer}})
	stroker.HalfLineWidth = gc.Current.LineWidth / 2

	var liner draw2dbase.Flattener
	if gc.Current.Dash != nil && len(gc.Current.Dash) > 0 {
		liner = draw2dbase.NewDashConverter(gc.Current.Dash, gc.Current.DashOffset, stroker)
	} else {
		liner = stroker
	}

	demux := draw2dbase.DemuxFlattener{Flatteners: []draw2dbase.Flattener{flattener, liner}}
	for _, p := range paths {
		draw2dbase.Flatten(p, demux, gc.Current.Tr.GetScale())
	}

	// Fill
	gc.paint(gc.fillRasterizer, gc.Current.FillColor)
	// Stroke
	gc.paint(gc.strokeRasterizer, gc.Current.StrokeColor)
}

func toFtCap(c draw2d.LineCap) raster.Capper {
	switch c {
	case draw2d.RoundCap:
		return raster.RoundCapper
	case draw2d.ButtCap:
		return raster.ButtCapper
	case draw2d.SquareCap:
		return raster.SquareCapper
	}
	return raster.RoundCapper
}

func toFtJoin(j draw2d.LineJoin) raster.Joiner {
	switch j {
	case draw2d.RoundJoin:
		return raster.RoundJoiner
	case draw2d.BevelJoin:
		return raster.BevelJoiner
	}
	return raster.RoundJoiner
}
