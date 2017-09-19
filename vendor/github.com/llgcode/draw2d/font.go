// Copyright 2010 The draw2d Authors. All rights reserved.
// created: 13/12/2010 by Laurent Le Goff

package draw2d

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/golang/freetype/truetype"
)

// FontStyle defines bold and italic styles for the font
// It is possible to combine values for mixed styles, eg.
//     FontData.Style = FontStyleBold | FontStyleItalic
type FontStyle byte

const (
	FontStyleNormal FontStyle = iota
	FontStyleBold
	FontStyleItalic
)

type FontFamily byte

const (
	FontFamilySans FontFamily = iota
	FontFamilySerif
	FontFamilyMono
)

type FontData struct {
	Name   string
	Family FontFamily
	Style  FontStyle
}

type FontFileNamer func(fontData FontData) string

func FontFileName(fontData FontData) string {
	fontFileName := fontData.Name
	switch fontData.Family {
	case FontFamilySans:
		fontFileName += "s"
	case FontFamilySerif:
		fontFileName += "r"
	case FontFamilyMono:
		fontFileName += "m"
	}
	if fontData.Style&FontStyleBold != 0 {
		fontFileName += "b"
	} else {
		fontFileName += "r"
	}

	if fontData.Style&FontStyleItalic != 0 {
		fontFileName += "i"
	}
	fontFileName += ".ttf"
	return fontFileName
}

func RegisterFont(fontData FontData, font *truetype.Font) {
	fontCache.Store(fontData, font)
}

func GetFont(fontData FontData) (font *truetype.Font) {
	var err error

	if font, err = fontCache.Load(fontData); err != nil {
		log.Println(err)
	}

	return
}

func GetFontFolder() string {
	return defaultFonts.folder
}

func SetFontFolder(folder string) {
	defaultFonts.folder = filepath.Clean(folder)
}

func SetFontNamer(fn FontFileNamer) {
	defaultFonts.namer = fn
}

// Types implementing this interface can be passed to SetFontCache to change the
// way fonts are being stored and retrieved.
type FontCache interface {
	// Loads a truetype font represented by the FontData object passed as
	// argument.
	// The method returns an error if the font could not be loaded, either
	// because it didn't exist or the resource it was loaded from was corrupted.
	Load(FontData) (*truetype.Font, error)

	// Sets the truetype font that will be returned by Load when given the font
	// data passed as first argument.
	Store(FontData, *truetype.Font)
}

// Changes the font cache backend used by the package. After calling this
// functionSetFontFolder and SetFontNamer will not affect anymore how fonts are
// loaded.
// To restore the default font cache, call this function passing nil as argument.
func SetFontCache(cache FontCache) {
	if cache == nil {
		fontCache = defaultFonts
	} else {
		fontCache = cache
	}
}

type defaultFontCache struct {
	fonts  map[string]*truetype.Font
	folder string
	namer  FontFileNamer
}

func (cache *defaultFontCache) Load(fontData FontData) (font *truetype.Font, err error) {
	if font = cache.fonts[cache.namer(fontData)]; font != nil {
		return font, nil
	}

	var data []byte
	var file = cache.namer(fontData)

	if data, err = ioutil.ReadFile(filepath.Join(cache.folder, file)); err != nil {
		return
	}

	if font, err = truetype.Parse(data); err != nil {
		return
	}

	cache.fonts[file] = font
	return
}

func (cache *defaultFontCache) Store(fontData FontData, font *truetype.Font) {
	cache.fonts[cache.namer(fontData)] = font
}

var (
	defaultFonts = &defaultFontCache{
		fonts:  make(map[string]*truetype.Font),
		folder: "../resource/font",
		namer:  FontFileName,
	}

	fontCache FontCache = defaultFonts
)
