package gofpdf

/*
 * Copyright (c) 2015 Kurt Jung (Gmail: kurt.w.jung),
 *   Marcus Downing, Jan Slabon (Setasign)
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

// newTpl creates a template, copying graphics settings from a template if one is given
func newTpl(corner PointType, size SizeType, unitStr, fontDirStr string, fn func(*Tpl), copyFrom *Fpdf) Template {
	orientationStr := "p"
	if size.Wd > size.Ht {
		orientationStr = "l"
	}
	sizeStr := ""

	fpdf := fpdfNew(orientationStr, unitStr, sizeStr, fontDirStr, size)
	tpl := Tpl{*fpdf}
	if copyFrom != nil {
		tpl.loadParamsFromFpdf(copyFrom)
	}
	tpl.Fpdf.SetAutoPageBreak(false, 0)
	tpl.Fpdf.AddPage()
	fn(&tpl)
	bytes := tpl.Fpdf.pages[tpl.Fpdf.page].Bytes()
	templates := make([]Template, 0, len(tpl.Fpdf.templates))
	for _, key := range templateKeyList(tpl.Fpdf.templates, true) {
		templates = append(templates, tpl.Fpdf.templates[key])
	}
	images := tpl.Fpdf.images

	id := GenerateTemplateID()
	template := FpdfTpl{id, corner, size, bytes, images, templates}
	return &template
}

// FpdfTpl is a concrete implementation of the Template interface.
type FpdfTpl struct {
	id        int64
	corner    PointType
	size      SizeType
	bytes     []byte
	images    map[string]*ImageInfoType
	templates []Template
}

// ID returns the global template identifier
func (t *FpdfTpl) ID() int64 {
	return t.id
}

// Size gives the bounding dimensions of this template
func (t *FpdfTpl) Size() (corner PointType, size SizeType) {
	return t.corner, t.size
}

// Bytes returns the actual template data, not including resources
func (t *FpdfTpl) Bytes() []byte {
	return t.bytes
}

// Images returns a list of the images used in this template
func (t *FpdfTpl) Images() map[string]*ImageInfoType {
	return t.images
}

// Templates returns a list of templates used in this template
func (t *FpdfTpl) Templates() []Template {
	return t.templates
}

// Tpl is an Fpdf used for writing a template. It has most of the facilities of
// an Fpdf, but cannot add more pages. Tpl is used directly only during the
// limited time a template is writable.
type Tpl struct {
	Fpdf
}

func (t *Tpl) loadParamsFromFpdf(f *Fpdf) {
	t.Fpdf.compress = false

	t.Fpdf.k = f.k
	t.Fpdf.x = f.x
	t.Fpdf.y = f.y
	t.Fpdf.lineWidth = f.lineWidth
	t.Fpdf.capStyle = f.capStyle
	t.Fpdf.joinStyle = f.joinStyle

	t.Fpdf.color.draw = f.color.draw
	t.Fpdf.color.fill = f.color.fill
	t.Fpdf.color.text = f.color.text

	t.Fpdf.fonts = f.fonts
	t.Fpdf.currentFont = f.currentFont
	t.Fpdf.fontFamily = f.fontFamily
	t.Fpdf.fontSize = f.fontSize
	t.Fpdf.fontSizePt = f.fontSizePt
	t.Fpdf.fontStyle = f.fontStyle
	t.Fpdf.ws = f.ws
}

// AddPage does nothing because you cannot add pages to a template
func (t *Tpl) AddPage() {
}

// AddPageFormat does nothign because you cannot add pages to a template
func (t *Tpl) AddPageFormat(orientationStr string, size SizeType) {
}

// SetAutoPageBreak does nothing because you cannot add pages to a template
func (t *Tpl) SetAutoPageBreak(auto bool, margin float64) {
}
