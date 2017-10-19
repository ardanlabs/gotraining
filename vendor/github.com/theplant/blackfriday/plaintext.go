//
// Blackfriday Markdown Processor
// Available at http://github.com/russross/blackfriday
//
// Copyright © 2011 Russ Ross <russ@russross.com>.
// Distributed under the Simplified BSD License.
// See README.md for details.
//

//
//
// HTML rendering backend
//
//

package blackfriday

import "bytes"

// PlainText renderer configuration options.
const (
	PLAIN_TEXT_SKIP_HTML = 1 << iota // skip preformatted HTML blocks
)

// PlainText is a type that implements the Renderer interface for HTML output.
//
// Do not create this directly, instead use the PlainTextRenderer function.
type PlainText struct {
	flags int    // HTML_* options
	title string // document title

	// table of contents data
	headerCount  int
	currentLevel int
}

// PlainTextRenderer creates and configures an PlainText object, which
// satisfies the Renderer interface.
//
// flags is a set of HTML_* options ORed together.
func PlainTextRenderer(flags int, title string) Renderer {
	// configure the rendering engine

	return &PlainText{
		flags:        flags,
		title:        title,
		headerCount:  0,
		currentLevel: 0,
	}
}

func (options *PlainText) Header(out *bytes.Buffer, text func() bool, level int) {
	marker := out.Len()
	doubleSpace(out)

	if !text() {
		out.Truncate(marker)
		return
	}
	out.WriteByte('\n')
}

func (options *PlainText) BlockHtml(out *bytes.Buffer, text []byte) {
	doubleSpace(out)
	out.Write(text)
	out.WriteByte('\n')
}

func (options *PlainText) HRule(out *bytes.Buffer) {
	doubleSpace(out)
	out.WriteString("\n")
}

func (options *PlainText) BlockCode(out *bytes.Buffer, text []byte, lang string) {
	options.BlockCodeNormal(out, text, lang)
}

func (options *PlainText) BlockCodeNormal(out *bytes.Buffer, text []byte, lang string) {
	doubleSpace(out)
	out.Write(text)
	out.WriteString("\n")
}

func (options *PlainText) BlockQuote(out *bytes.Buffer, text []byte) {
	doubleSpace(out)
	out.Write(text)
	out.WriteString("\n")
}

func (options *PlainText) Table(out *bytes.Buffer, header []byte, body []byte, columnData []int) {
	doubleSpace(out)
	out.WriteString("\n")
	out.Write(header)
	out.WriteString("\n")
	out.Write(body)
	out.WriteString("\n")
}

func (options *PlainText) TableRow(out *bytes.Buffer, text []byte) {
	doubleSpace(out)
	out.Write(text)
	out.WriteString("\n")
}

func (options *PlainText) TableCell(out *bytes.Buffer, text []byte, align int) {
	doubleSpace(out)
	out.Write(text)
	out.WriteString("|")
}

func (options *PlainText) Footnotes(out *bytes.Buffer, text func() bool) {
	options.HRule(out)
	options.List(out, text, LIST_TYPE_ORDERED)
}

func (options *PlainText) FootnoteItem(out *bytes.Buffer, name, text []byte, flags int) {
	doubleSpace(out)
	out.Write(text)
	out.WriteString("\n")
}

func (options *PlainText) List(out *bytes.Buffer, text func() bool, flags int) {
	marker := out.Len()
	doubleSpace(out)
	if !text() {
		out.Truncate(marker)
		return
	}
}

func (options *PlainText) ListItem(out *bytes.Buffer, text []byte, flags int) {
	if flags&LIST_ITEM_CONTAINS_BLOCK != 0 || flags&LIST_ITEM_BEGINNING_OF_LIST != 0 {
		doubleSpace(out)
	}
	out.Write(text)
	out.WriteString("\n")
}

func (options *PlainText) Paragraph(out *bytes.Buffer, text func() bool) {
	marker := out.Len()
	doubleSpace(out)
	if !text() {
		out.Truncate(marker)
		return
	}
	//out.WriteString("\n")
}

func (options *PlainText) AutoLink(out *bytes.Buffer, link []byte, kind int) {
	// Pretty print: if we get an email address as
	// an actual URI, e.g. `mailto:foo@bar.com`, we don't
	// want to print the `mailto:` prefix
	switch {
	case bytes.HasPrefix(link, []byte("mailto://")):
		out.Write(link[len("mailto://"):])
	case bytes.HasPrefix(link, []byte("mailto:")):
		out.Write(link[len("mailto:"):])
	default:
		out.Write(link)
	}
}

func (options *PlainText) CodeSpan(out *bytes.Buffer, text []byte) {
	out.WriteString(" ")
	out.Write(text)
	out.WriteString(" ")
}

func (options *PlainText) DoubleEmphasis(out *bytes.Buffer, text []byte) {
	out.Write(text)
}

func (options *PlainText) Emphasis(out *bytes.Buffer, text []byte) {
	if len(text) == 0 {
		return
	}
	out.Write(text)
}

func (options *PlainText) Image(out *bytes.Buffer, link []byte, title []byte, alt []byte) {
	if len(alt) == 0 {
		out.WriteString("IMAGE")
		return
	}
	out.Write(alt)
	return
}

func (options *PlainText) LineBreak(out *bytes.Buffer) {
	out.WriteString("\n")
}

func (options *PlainText) Link(out *bytes.Buffer, link []byte, title []byte, content []byte) {
	out.Write(content)
	return
}

func (options *PlainText) RawHtmlTag(out *bytes.Buffer, text []byte) {
	out.Write(text)
}

func (options *PlainText) TripleEmphasis(out *bytes.Buffer, text []byte) {
	out.Write(text)
}

func (options *PlainText) StrikeThrough(out *bytes.Buffer, text []byte) {
	out.Write(text)
}

func (options *PlainText) FootnoteRef(out *bytes.Buffer, ref []byte, id int) {
}

func (options *PlainText) Entity(out *bytes.Buffer, entity []byte) {
	out.Write(entity)
}

func (options *PlainText) NormalText(out *bytes.Buffer, text []byte) {
	out.Write(text)
}

func (options *PlainText) DocumentHeader(out *bytes.Buffer) {
}

func (options *PlainText) DocumentFooter(out *bytes.Buffer) {
}
