// Copyright (C) 2011, Ross Light

package pdf

import (
	"image"
	"io"
	"strconv"
)

// Unit is a device-independent dimensional type.  On a new canvas, this
// represents 1/72 of an inch.
type Unit float32

func (unit Unit) String() string {
	return strconv.FormatFloat(float64(unit), 'f', marshalFloatPrec, 32)
}

// Common unit scales
const (
	Pt   Unit = 1
	Inch Unit = 72
	Cm   Unit = 28.35
)

// Common page sizes
const (
	USLetterWidth  Unit = 8.5 * Inch
	USLetterHeight Unit = 11.0 * Inch

	A4Width  Unit = 21.0 * Cm
	A4Height Unit = 29.7 * Cm
)

// Document provides a high-level drawing interface for the PDF format.
type Document struct {
	encoder
	catalog *catalog
	pages   []indirectObject
	fonts   map[name]Reference
}

// New creates a new document with no pages.
func New() *Document {
	doc := new(Document)
	doc.catalog = &catalog{
		Type: catalogType,
	}
	doc.root = doc.add(doc.catalog)
	doc.fonts = make(map[name]Reference, 14)
	return doc
}

// NewPage creates a new canvas with the given dimensions.
func (doc *Document) NewPage(width, height Unit) *Canvas {
	page := &pageDict{
		Type:     pageType,
		MediaBox: Rectangle{Point{0, 0}, Point{width, height}},
		CropBox:  Rectangle{Point{0, 0}, Point{width, height}},
		Resources: resources{
			ProcSet: []name{pdfProcSet, textProcSet, imageCProcSet},
			Font:    make(map[name]interface{}),
			XObject: make(map[name]interface{}),
		},
	}
	pageRef := doc.add(page)
	doc.pages = append(doc.pages, indirectObject{pageRef, page})

	stream := newStream(streamFlateDecode)
	page.Contents = doc.add(stream)

	return &Canvas{
		doc:      doc,
		page:     page,
		ref:      pageRef,
		contents: stream,
	}
}

// standardFont returns a reference to a standard font dictionary.  If there is
// no font dictionary for the font in the document yet, it is added
// automatically.
func (doc *Document) standardFont(fontName name) Reference {
	if ref, ok := doc.fonts[fontName]; ok {
		return ref
	}

	// TODO: check name is standard?
	ref := doc.add(standardFontDict{
		Type:     fontType,
		Subtype:  fontType1Subtype,
		BaseFont: fontName,
	})
	doc.fonts[fontName] = ref
	return ref
}

// AddImage encodes an image into the document's stream and returns its PDF
// file reference.  This reference can be used to draw the image multiple times
// without storing the image multiple times.
func (doc *Document) AddImage(img image.Image) Reference {
	bd := img.Bounds()
	st := newImageStream(streamFlateDecode, bd.Dx(), bd.Dy())
	defer st.Close()

	switch i := img.(type) {
	case *image.RGBA:
		encodeRGBAStream(st, i)
	case *image.NRGBA:
		encodeNRGBAStream(st, i)
	case *image.YCbCr:
		encodeYCbCrStream(st, i)
	default:
		encodeImageStream(st, i)
	}
	return doc.add(st)
}

// Encode writes the document to a writer in the PDF format.
func (doc *Document) Encode(w io.Writer) error {
	pageRoot := &pageRootNode{
		Type:  pageNodeType,
		Count: len(doc.pages),
	}
	doc.catalog.Pages = doc.add(pageRoot)
	for _, p := range doc.pages {
		page := p.Object.(*pageDict)
		page.Parent = doc.catalog.Pages
		pageRoot.Kids = append(pageRoot.Kids, p.Reference)
	}

	return doc.encoder.encode(w)
}

// PDF object types
const (
	catalogType  name = "Catalog"
	pageNodeType name = "Pages"
	pageType     name = "Page"
	fontType     name = "Font"
	xobjectType  name = "XObject"
)

// PDF object subtypes
const (
	imageSubtype name = "Image"

	fontType1Subtype name = "Type1"
)

type catalog struct {
	Type  name
	Pages Reference
}

type pageRootNode struct {
	Type  name
	Kids  []Reference
	Count int
}

type pageNode struct {
	Type   name
	Parent Reference
	Kids   []Reference
	Count  int
}

type pageDict struct {
	Type      name
	Parent    Reference
	Resources resources
	MediaBox  Rectangle
	CropBox   Rectangle
	Contents  Reference
}

// Point is a 2D point.
type Point struct {
	X, Y Unit
}

// A Rectangle defines a rectangle with two points.
type Rectangle struct {
	Min, Max Point
}

// Dx returns the rectangle's width.
func (r Rectangle) Dx() Unit {
	return r.Max.X - r.Min.X
}

// Dy returns the rectangle's height.
func (r Rectangle) Dy() Unit {
	return r.Max.Y - r.Min.Y
}

func (r Rectangle) marshalPDF(dst []byte) ([]byte, error) {
	dst = append(dst, '[', ' ')
	dst, _ = marshal(dst, r.Min.X)
	dst = append(dst, ' ')
	dst, _ = marshal(dst, r.Min.Y)
	dst = append(dst, ' ')
	dst, _ = marshal(dst, r.Max.X)
	dst = append(dst, ' ')
	dst, _ = marshal(dst, r.Max.Y)
	dst = append(dst, ' ', ']')
	return dst, nil
}

type resources struct {
	ProcSet []name
	Font    map[name]interface{}
	XObject map[name]interface{}
}

// Predefined procedure sets
const (
	pdfProcSet    name = "PDF"
	textProcSet   name = "Text"
	imageBProcSet name = "ImageB"
	imageCProcSet name = "ImageC"
	imageIProcSet name = "ImageI"
)

type standardFontDict struct {
	Type     name
	Subtype  name
	BaseFont name
}
