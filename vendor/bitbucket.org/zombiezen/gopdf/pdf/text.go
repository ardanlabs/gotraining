// Copyright (C) 2011, Ross Light

package pdf

import (
	"bytes"
)

// Text is a PDF text object.  The zero value is an empty text object.
type Text struct {
	buf   bytes.Buffer
	fonts map[name]bool

	x, y        Unit
	currFont    name
	currSize    Unit
	currLeading Unit
}

// Text adds a string to the text object.
func (text *Text) Text(s string) {
	writeCommand(&text.buf, "Tj", s)
	if widths := getFontWidths(text.currFont); widths != nil {
		text.x += computeStringWidth(s, widths, text.currSize)
	}
}

const defaultLeadingScalar = 1.2

// SetFont changes the current font to a standard font.  This also changes the
// leading to 1.2 times the font size.
func (text *Text) SetFont(fontName string, size Unit) {
	if text.fonts == nil {
		text.fonts = make(map[name]bool)
	}
	text.fonts[name(fontName)] = true
	text.currFont, text.currSize = name(fontName), size
	writeCommand(&text.buf, "Tf", name(fontName), size)
	text.SetLeading(size * defaultLeadingScalar)
}

// SetLeading changes the amount of space between lines.
func (text *Text) SetLeading(leading Unit) {
	writeCommand(&text.buf, "TL", leading)
	text.currLeading = leading
}

// NextLine advances the current text position to the next line, based on the
// current leading.
func (text *Text) NextLine() {
	writeCommand(&text.buf, "T*")
	text.x = 0
	text.y -= text.currLeading
}

// NextLineOffset moves the current text position to an offset relative to the
// beginning of the line.
func (text *Text) NextLineOffset(tx, ty Unit) {
	writeCommand(&text.buf, "Td", tx, ty)
	text.x = tx
	text.y += ty
}

// X returns the current x position of the text cursor.
func (text *Text) X() Unit {
	return text.x
}

// Y returns the current y position of the text cursor.
func (text *Text) Y() Unit {
	return text.y
}

// Standard 14 fonts
const (
	Courier            = "Courier"
	CourierBold        = "Courier-Bold"
	CourierOblique     = "Courier-Oblique"
	CourierBoldOblique = "Courier-BoldOblique"

	Helvetica            = "Helvetica"
	HelveticaBold        = "Helvetica-Bold"
	HelveticaOblique     = "Helvetica-Oblique"
	HelveticaBoldOblique = "Helvetica-BoldOblique"

	Symbol = "Symbol"

	Times           = "Times-Roman"
	TimesBold       = "Times-Bold"
	TimesItalic     = "Times-Italic"
	TimesBoldItalic = "Times-BoldItalic"

	ZapfDingbats = "ZapfDingbats"
)

func getFontWidths(fontName name) []uint16 {
	switch fontName {
	case Courier:
		return courierWidths
	case CourierBold:
		return courierBoldWidths
	case CourierOblique:
		return courierObliqueWidths
	case CourierBoldOblique:
		return courierBoldObliqueWidths
	case Helvetica:
		return helveticaWidths
	case HelveticaBold:
		return helveticaBoldWidths
	case HelveticaOblique:
		return helveticaObliqueWidths
	case HelveticaBoldOblique:
		return helveticaBoldObliqueWidths
	case Symbol:
		return symbolWidths
	case Times:
		return timesRomanWidths
	case TimesBold:
		return timesBoldWidths
	case TimesItalic:
		return timesItalicWidths
	case TimesBoldItalic:
		return timesBoldItalicWidths
	case ZapfDingbats:
		return zapfDingbatsWidths
	}
	return nil
}

func computeStringWidth(s string, widths []uint16, fontSize Unit) Unit {
	width := Unit(0)
	for _, r := range s {
		if int(r) < len(widths) {
			width += Unit(widths[r])
		}
	}
	return width * fontSize / 1000
}
