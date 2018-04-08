// Copyright 2010 The draw2d Authors. All rights reserved.
// created: 13/12/2010 by Laurent Le Goff

package draw2d

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/golang/freetype/truetype"
	"sync"
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
	defaultFonts.setFolder(filepath.Clean(folder))
}

func GetGlobalFontCache() FontCache {
	return fontCache
}

func SetFontNamer(fn FontFileNamer) {
	defaultFonts.setNamer(fn)
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

// FolderFontCache can Load font from folder
type FolderFontCache struct {
	fonts  map[string]*truetype.Font
	folder string
	namer  FontFileNamer
}

// NewFolderFontCache creates FolderFontCache
func NewFolderFontCache(folder string) *FolderFontCache {
	return &FolderFontCache{
		fonts:  make(map[string]*truetype.Font),
		folder: folder,
		namer:  FontFileName,
	}
}

// Load a font from cache if exists otherwise it will load the font from file
func (cache *FolderFontCache) Load(fontData FontData) (font *truetype.Font, err error) {
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

// Store a font to this cache
func (cache *FolderFontCache) Store(fontData FontData, font *truetype.Font) {
	cache.fonts[cache.namer(fontData)] = font
}

// SyncFolderFontCache can Load font from folder
type SyncFolderFontCache struct {
	sync.RWMutex
	fonts  map[string]*truetype.Font
	folder string
	namer  FontFileNamer
}

// NewSyncFolderFontCache creates SyncFolderFontCache
func NewSyncFolderFontCache(folder string) *SyncFolderFontCache {
	return &SyncFolderFontCache{
		fonts:  make(map[string]*truetype.Font),
		folder: folder,
		namer:  FontFileName,
	}
}

func (cache *SyncFolderFontCache) setFolder(folder string) {
	cache.Lock()
	cache.folder = folder
	cache.Unlock()
}

func (cache *SyncFolderFontCache) setNamer(namer FontFileNamer) {
	cache.Lock()
	cache.namer = namer
	cache.Unlock()
}

// Load a font from cache if exists otherwise it will load the font from file
func (cache *SyncFolderFontCache) Load(fontData FontData) (font *truetype.Font, err error) {
	cache.RLock()
	font = cache.fonts[cache.namer(fontData)]
	cache.RUnlock()

	if font != nil {
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
	cache.Lock()
	cache.fonts[file] = font
	cache.Unlock()
	return
}

// Store a font to this cache
func (cache *SyncFolderFontCache) Store(fontData FontData, font *truetype.Font) {
	cache.Lock()
	cache.fonts[cache.namer(fontData)] = font
	cache.Unlock()
}

var (
	defaultFonts = NewSyncFolderFontCache("../resource/font")

	fontCache FontCache = defaultFonts
)
