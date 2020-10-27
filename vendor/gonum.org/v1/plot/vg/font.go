// Copyright ©2015 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Some of this code (namely the code for computing the
// width of a string in a given font) was copied from
// github.com/golang/freetype/ which includes
// the following copyright notice:
// Copyright 2010 The Freetype-Go Authors. All rights reserved.

package vg

import (
	"errors"
	"io/ioutil"
	"path/filepath"
	"sync"

	"gonum.org/v1/plot/vg/fonts"

	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"

	"github.com/golang/freetype/truetype"
)

var (
	// FontMap maps Postscript/PDF font names to compatible
	// free fonts (TrueType converted ghostscript fonts).
	// Fonts that are not keys of this map are not supported.
	FontMap = map[string]string{

		// We use fonts from RedHat's Liberation project:
		//  https://fedorahosted.org/liberation-fonts/

		"Courier":             "LiberationMono-Regular",
		"Courier-Bold":        "LiberationMono-Bold",
		"Courier-Oblique":     "LiberationMono-Italic",
		"Courier-BoldOblique": "LiberationMono-BoldItalic",

		"Helvetica":             "LiberationSans-Regular",
		"Helvetica-Bold":        "LiberationSans-Bold",
		"Helvetica-Oblique":     "LiberationSans-Italic",
		"Helvetica-BoldOblique": "LiberationSans-BoldItalic",

		"Times-Roman":      "LiberationSerif-Regular",
		"Times-Bold":       "LiberationSerif-Bold",
		"Times-Italic":     "LiberationSerif-Italic",
		"Times-BoldItalic": "LiberationSerif-BoldItalic",
	}

	// loadedFonts is indexed by a font name and it
	// caches the associated *truetype.Font.
	loadedFonts = make(map[string]*truetype.Font)

	// FontLock protects access to the loadedFonts map.
	fontLock sync.RWMutex
)

// A Font represents one of the supported font
// faces.
type Font struct {
	// Size is the size of the font.  The font size can
	// be used as a reasonable value for the vertical
	// distance between two successive lines of text.
	Size Length

	// name is the name of this font.
	name string

	// This is a little bit of a hack, but the truetype
	// font is currently only needed to determine the
	// dimensions of strings drawn in this font.
	// The actual drawing of the strings is handled
	// separately by different back-ends:
	// Both Postscript and PDF are capable of drawing
	// their own fonts and draw2d loads its own copy of
	// the truetype fonts for its own output.
	//
	// This isn't a necessity--some future backend is
	// free to use this field--however it is a consequence
	// of the fact that the current backends were
	// developed independently of this package.

	// font is the truetype font pointer for this
	// font.
	font *truetype.Font
}

// MakeFont returns a font object.  The name of the font must
// be a key of the FontMap.  The font file is located by searching
// the FontDirs slice for a directory containing the relevant font
// file.  The font file name is name mapped by FontMap with the
// .ttf extension.  For example, the font file for the font name
// Courier is LiberationMono-Regular.ttf.
func MakeFont(name string, size Length) (font Font, err error) {
	font.Size = size
	font.name = name
	font.font, err = getFont(name)
	return
}

// Name returns the name of the font.
func (f *Font) Name() string {
	return f.name
}

// Font returns the corresponding truetype.Font.
func (f *Font) Font() *truetype.Font {
	return f.font
}

func (f *Font) FontFace(dpi float64) font.Face {
	return truetype.NewFace(f.font, &truetype.Options{
		Size: f.Size.Points(),
		DPI:  dpi,
	})
}

// SetName sets the name of the font, effectively
// changing the font.  If an error is returned then
// the font is left unchanged.
func (f *Font) SetName(name string) error {
	font, err := getFont(name)
	if err != nil {
		return err
	}
	f.name = name
	f.font = font
	return nil
}

// FontExtents contains font metric information.
type FontExtents struct {
	// Ascent is the distance that the text
	// extends above the baseline.
	Ascent Length

	// Descent is the distance that the text
	// extends below the baseline.  The descent
	// is given as a negative value.
	Descent Length

	// Height is the distance from the lowest
	// descending point to the highest ascending
	// point.
	Height Length
}

// Extents returns the FontExtents for a font.
func (f *Font) Extents() FontExtents {
	bounds := f.font.Bounds(fixed.Int26_6(f.Font().FUnitsPerEm()))
	scale := f.Size / Points(float64(f.Font().FUnitsPerEm()))
	return FontExtents{
		Ascent:  Points(float64(bounds.Max.Y)) * scale,
		Descent: Points(float64(bounds.Min.Y)) * scale,
		Height:  Points(float64(bounds.Max.Y-bounds.Min.Y)) * scale,
	}
}

// Width returns width of a string when drawn using the font.
func (f *Font) Width(s string) Length {
	// scale converts truetype.FUnit to float64
	scale := f.Size / Points(float64(f.font.FUnitsPerEm()))

	width := 0
	prev, hasPrev := truetype.Index(0), false
	for _, rune := range s {
		index := f.font.Index(rune)
		if hasPrev {
			width += int(f.font.Kern(fixed.Int26_6(f.font.FUnitsPerEm()), prev, index))
		}
		width += int(f.font.HMetric(fixed.Int26_6(f.font.FUnitsPerEm()), index).AdvanceWidth)
		prev, hasPrev = index, true
	}
	return Points(float64(width)) * scale
}

// AddFont associates a truetype.Font with the given name.
func AddFont(name string, font *truetype.Font) {
	fontLock.Lock()
	loadedFonts[name] = font
	fontLock.Unlock()
}

// getFont returns the truetype.Font for the given font name or an error.
func getFont(name string) (*truetype.Font, error) {
	fontLock.RLock()
	f, ok := loadedFonts[name]
	fontLock.RUnlock()
	if ok {
		return f, nil
	}

	bytes, err := fontData(name)
	if err != nil {
		return nil, err
	}

	font, err := truetype.Parse(bytes)
	if err == nil {
		fontLock.Lock()
		loadedFonts[name] = font
		fontLock.Unlock()
	} else {
		err = errors.New("Failed to parse font file: " + err.Error())
	}

	return font, err
}

// fontData returns the []byte data for a font name or an error if it is not found.
func fontData(name string) ([]byte, error) {
	fname, err := fontFile(name)
	if err != nil {
		return nil, err
	}

	for _, d := range FontDirs {
		p := filepath.Join(d, fname)
		data, err := ioutil.ReadFile(p)
		if err != nil {
			continue
		}
		return data, nil
	}

	data, err := fonts.Asset(fname)
	if err == nil {
		return data, nil
	}

	return nil, errors.New("vg: failed to locate a font file " + fname + " for font name " + name)
}

// FontDirs is a slice of directories searched for font data files.
// If the first font file found is unreadable or cannot be parsed, then
// subsequent directories are not tried, and the font will fail to load.
//
// The default slice is initialised with the contents of the VGFONTPATH
// environment variable if it is defined.
// This slice may be changed to load fonts from different locations.
var FontDirs []string

// FontFile returns the font file name for a font name or an error
// if it is an unknown font (i.e., not in the FontMap).
func fontFile(name string) (string, error) {
	var err error
	n, ok := FontMap[name]
	if !ok {
		errStr := "Unknown font: " + name + ".  Available fonts are:"
		for n := range FontMap {
			errStr += " " + n
		}
		err = errors.New(errStr)
	}
	return n + ".ttf", err
}
