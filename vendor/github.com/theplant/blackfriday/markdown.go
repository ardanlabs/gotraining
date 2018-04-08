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
// Markdown parsing and processing
//
//

// Blackfriday markdown processor.
//
// Translates plain text with simple formatting rules into HTML or LaTeX.
package blackfriday

import (
	"bytes"
	"utf8"
)

const VERSION = "1.0"

// These are the supported markdown parsing extensions.
// OR these values together to select multiple extensions.
const (
	EXTENSION_NO_INTRA_EMPHASIS = 1 << iota // ignore emphasis markers inside words
	EXTENSION_TABLES                        // render tables
	EXTENSION_FENCED_CODE                   // render fenced code blocks
	EXTENSION_AUTOLINK                      // detect embedded URLs that are not explicitly marked
	EXTENSION_STRIKETHROUGH                 // strikethrough text using ~~test~~
	EXTENSION_LAX_HTML_BLOCKS               // loosen up HTML block parsing rules
	EXTENSION_SPACE_HEADERS                 // be strict about prefix header rules
	EXTENSION_HARD_LINE_BREAK               // translate newlines into line breaks
	EXTENSION_TAB_SIZE_EIGHT                // expand tabs to eight spaces instead of four
)

// These are the possible flag values for the link renderer.
// Only a single one of these values will be used; they are not ORed together.
// These are mostly of interest if you are writing a new output format.
const (
	LINK_TYPE_NOT_AUTOLINK = iota
	LINK_TYPE_NORMAL
	LINK_TYPE_EMAIL
)

// These are the possible flag values for the ListItem renderer.
// Multiple flag values may be ORed together.
// These are mostly of interest if you are writing a new output format.
const (
	LIST_TYPE_ORDERED = 1 << iota
	LIST_ITEM_CONTAINS_BLOCK
	LIST_ITEM_BEGINNING_OF_LIST
	LIST_ITEM_END_OF_LIST
)

// These are the possible flag values for the table cell renderer.
// Only a single one of these values will be used; they are not ORed together.
// These are mostly of interest if you are writing a new output format.
const (
	TABLE_ALIGNMENT_LEFT = 1 << iota
	TABLE_ALIGNMENT_RIGHT
	TABLE_ALIGNMENT_CENTER = (TABLE_ALIGNMENT_LEFT | TABLE_ALIGNMENT_RIGHT)
)

// The size of a tab stop.
const (
	TAB_SIZE_DEFAULT = 4
	TAB_SIZE_EIGHT   = 8
)

// These are the tags that are recognized as HTML block tags.
// Any of these can be included in markdown text without special escaping.
var blockTags = map[string]bool{
	"p":          true,
	"dl":         true,
	"h1":         true,
	"h2":         true,
	"h3":         true,
	"h4":         true,
	"h5":         true,
	"h6":         true,
	"ol":         true,
	"ul":         true,
	"del":        true,
	"div":        true,
	"ins":        true,
	"pre":        true,
	"form":       true,
	"math":       true,
	"table":      true,
	"iframe":     true,
	"script":     true,
	"fieldset":   true,
	"noscript":   true,
	"blockquote": true,
}

// Renderer is the rendering interface.
// This is mostly of interest if you are implementing a new rendering format.
//
// When a byte slice is provided, it contains the (rendered) contents of the
// element.
//
// When a callback is provided instead, it will write the contents of the
// respective element directly to the output buffer and return true on success.
// If the callback returns false, the rendering function should reset the
// output buffer as though it had never been called.
//
// Currently Html and Latex implementations are provided
type Renderer interface {
	// block-level callbacks
	BlockCode(out *bytes.Buffer, text []byte, lang string)
	BlockQuote(out *bytes.Buffer, text []byte)
	BlockHtml(out *bytes.Buffer, text []byte)
	Header(out *bytes.Buffer, text func() bool, level int)
	HRule(out *bytes.Buffer)
	List(out *bytes.Buffer, text func() bool, flags int)
	ListItem(out *bytes.Buffer, text []byte, flags int)
	Paragraph(out *bytes.Buffer, text func() bool)
	Table(out *bytes.Buffer, header []byte, body []byte, columnData []int)
	TableRow(out *bytes.Buffer, text []byte)
	TableCell(out *bytes.Buffer, text []byte, flags int)

	// Span-level callbacks
	AutoLink(out *bytes.Buffer, link []byte, kind int)
	CodeSpan(out *bytes.Buffer, text []byte)
	DoubleEmphasis(out *bytes.Buffer, text []byte)
	Emphasis(out *bytes.Buffer, text []byte)
	Image(out *bytes.Buffer, link []byte, title []byte, alt []byte)
	LineBreak(out *bytes.Buffer)
	Link(out *bytes.Buffer, link []byte, title []byte, content []byte)
	RawHtmlTag(out *bytes.Buffer, tag []byte)
	TripleEmphasis(out *bytes.Buffer, text []byte)
	StrikeThrough(out *bytes.Buffer, text []byte)

	// Low-level callbacks
	Entity(out *bytes.Buffer, entity []byte)
	NormalText(out *bytes.Buffer, text []byte)

	// Header and footer
	DocumentHeader(out *bytes.Buffer)
	DocumentFooter(out *bytes.Buffer)
}

// Callback functions for inline parsing. One such function is defined
// for each character that triggers a response when parsing inline data.
type inlineParser func(p *parser, out *bytes.Buffer, data []byte, offset int) int

// Parser holds runtime state used by the parser.
// This is constructed by the Markdown function.
type parser struct {
	r              Renderer
	refs           map[string]*reference
	inlineCallback [256]inlineParser
	flags          int
	nesting        int
	maxNesting     int
	insideLink     bool
}


//
//
// Public interface
//
//

// MarkdownBasic is a convenience function for simple rendering.
// It processes markdown input with no extensions enabled.
func MarkdownBasic(input []byte) []byte {
	// set up the HTML renderer
	htmlFlags := HTML_USE_XHTML
	renderer := HtmlRenderer(htmlFlags, "", "")

	// set up the parser
	extensions := 0

	return Markdown(input, renderer, extensions)
}

// Call Markdown with most useful extensions enabled
// MarkdownCommon is a convenience function for simple rendering.
// It processes markdown input with common extensions enabled, including:
//
// * Smartypants processing with smart fractions and LaTeX dashes
//
// * Intra-word emphasis supression
//
// * Tables
//
// * Fenced code blocks
//
// * Autolinking
//
// * Strikethrough support
//
// * Strict header parsing
func MarkdownCommon(input []byte) []byte {
	// set up the HTML renderer
	htmlFlags := 0
	htmlFlags |= HTML_USE_XHTML
	htmlFlags |= HTML_USE_SMARTYPANTS
	htmlFlags |= HTML_SMARTYPANTS_FRACTIONS
	htmlFlags |= HTML_SMARTYPANTS_LATEX_DASHES
	renderer := HtmlRenderer(htmlFlags, "", "")

	// set up the parser
	extensions := 0
	extensions |= EXTENSION_NO_INTRA_EMPHASIS
	extensions |= EXTENSION_TABLES
	extensions |= EXTENSION_FENCED_CODE
	extensions |= EXTENSION_AUTOLINK
	extensions |= EXTENSION_STRIKETHROUGH
	extensions |= EXTENSION_SPACE_HEADERS

	return Markdown(input, renderer, extensions)
}

// Markdown is the main rendering function.
// It parses and renders a block of markdown-encoded text.
// The supplied Renderer is used to format the output, and extensions dictates
// which non-standard extensions are enabled.
//
// To use the supplied Html or LaTeX renderers, see HtmlRenderer and
// LatexRenderer, respectively.
func Markdown(input []byte, renderer Renderer, extensions int) []byte {
	// no point in parsing if we can't render
	if renderer == nil {
		return nil
	}

	// fill in the render structure
	p := new(parser)
	p.r = renderer
	p.flags = extensions
	p.refs = make(map[string]*reference)
	p.maxNesting = 16
	p.insideLink = false

	// register inline parsers
	p.inlineCallback['*'] = emphasis
	p.inlineCallback['_'] = emphasis
	if extensions&EXTENSION_STRIKETHROUGH != 0 {
		p.inlineCallback['~'] = emphasis
	}
	p.inlineCallback['`'] = codeSpan
	p.inlineCallback['\n'] = lineBreak
	p.inlineCallback['['] = link
	p.inlineCallback['<'] = leftAngle
	p.inlineCallback['\\'] = escape
	p.inlineCallback['&'] = entity

	if extensions&EXTENSION_AUTOLINK != 0 {
		p.inlineCallback[':'] = autoLink
	}

	first := firstPass(p, input)
	second := secondPass(p, first)

	return second
}

// first pass:
// - extract references
// - expand tabs
// - normalize newlines
// - copy everything else
func firstPass(p *parser, input []byte) []byte {
	var out bytes.Buffer
	tabSize := TAB_SIZE_DEFAULT
	if p.flags&EXTENSION_TAB_SIZE_EIGHT != 0 {
		tabSize = TAB_SIZE_EIGHT
	}
	beg, end := 0, 0
	for beg < len(input) { // iterate over lines
		if end = isReference(p, input[beg:]); end > 0 {
			beg += end
		} else { // skip to the next line
			end = beg
			for end < len(input) && input[end] != '\n' && input[end] != '\r' {
				end++
			}

			// add the line body if present
			if end > beg {
				expandTabs(&out, input[beg:end], tabSize)
			}
			out.WriteByte('\n')

			if end < len(input) && input[end] == '\r' {
				end++
			}
			if end < len(input) && input[end] == '\n' {
				end++
			}

			beg = end
		}
	}

	// empty input?
	if out.Len() == 0 {
		out.WriteByte('\n')
	}

	return out.Bytes()
}

// second pass: actual rendering
func secondPass(p *parser, input []byte) []byte {
	var output bytes.Buffer

	p.r.DocumentHeader(&output)
	p.block(&output, input)
	p.r.DocumentFooter(&output)

	if p.nesting != 0 {
		panic("Nesting level did not end at zero")
	}

	return output.Bytes()
}


//
// Link references
//
// This section implements support for references that (usually) appear
// as footnotes in a document, and can be referenced anywhere in the document.
// The basic format is:
//
//    [1]: http://www.google.com/ "Google"
//    [2]: http://www.github.com/ "Github"
//
// Anywhere in the document, the reference can be linked by referring to its
// label, i.e., 1 and 2 in this example, as in:
//
//    This library is hosted on [Github][2], a git hosting site.

// References are parsed and stored in this struct.
type reference struct {
	link  []byte
	title []byte
}

// Check whether or not data starts with a reference link.
// If so, it is parsed and stored in the list of references
// (in the render struct).
// Returns the number of bytes to skip to move past it,
// or zero if the first line is not a reference.
func isReference(p *parser, data []byte) int {
	// up to 3 optional leading spaces
	if len(data) < 4 {
		return 0
	}
	i := 0
	for i < 3 && data[i] == ' ' {
		i++
	}

	// id part: anything but a newline between brackets
	if data[i] != '[' {
		return 0
	}
	i++
	idOffset := i
	for i < len(data) && data[i] != '\n' && data[i] != '\r' && data[i] != ']' {
		i++
	}
	if i >= len(data) || data[i] != ']' {
		return 0
	}
	idEnd := i

	// spacer: colon (space | tab)* newline? (space | tab)*
	i++
	if i >= len(data) || data[i] != ':' {
		return 0
	}
	i++
	for i < len(data) && (data[i] == ' ' || data[i] == '\t') {
		i++
	}
	if i < len(data) && (data[i] == '\n' || data[i] == '\r') {
		i++
		if i < len(data) && data[i] == '\n' && data[i-1] == '\r' {
			i++
		}
	}
	for i < len(data) && (data[i] == ' ' || data[i] == '\t') {
		i++
	}
	if i >= len(data) {
		return 0
	}

	// link: whitespace-free sequence, optionally between angle brackets
	if data[i] == '<' {
		i++
	}
	linkOffset := i
	for i < len(data) && data[i] != ' ' && data[i] != '\t' && data[i] != '\n' && data[i] != '\r' {
		i++
	}
	linkEnd := i
	if data[linkOffset] == '<' && data[linkEnd-1] == '>' {
		linkOffset++
		linkEnd--
	}

	// optional spacer: (space | tab)* (newline | '\'' | '"' | '(' )
	for i < len(data) && (data[i] == ' ' || data[i] == '\t') {
		i++
	}
	if i < len(data) && data[i] != '\n' && data[i] != '\r' && data[i] != '\'' && data[i] != '"' && data[i] != '(' {
		return 0
	}

	// compute end-of-line
	lineEnd := 0
	if i >= len(data) || data[i] == '\r' || data[i] == '\n' {
		lineEnd = i
	}
	if i+1 < len(data) && data[i] == '\r' && data[i+1] == '\n' {
		lineEnd++
	}

	// optional (space|tab)* spacer after a newline
	if lineEnd > 0 {
		i = lineEnd + 1
		for i < len(data) && (data[i] == ' ' || data[i] == '\t') {
			i++
		}
	}

	// optional title: any non-newline sequence enclosed in '"() alone on its line
	titleOffset, titleEnd := 0, 0
	if i+1 < len(data) && (data[i] == '\'' || data[i] == '"' || data[i] == '(') {
		i++
		titleOffset = i

		// look for EOL
		for i < len(data) && data[i] != '\n' && data[i] != '\r' {
			i++
		}
		if i+1 < len(data) && data[i] == '\n' && data[i+1] == '\r' {
			titleEnd = i + 1
		} else {
			titleEnd = i
		}

		// step back
		i--
		for i > titleOffset && (data[i] == ' ' || data[i] == '\t') {
			i--
		}
		if i > titleOffset && (data[i] == '\'' || data[i] == '"' || data[i] == ')') {
			lineEnd = titleEnd
			titleEnd = i
		}
	}
	if lineEnd == 0 { // garbage after the link
		return 0
	}

	// a valid ref has been found

	// id matches are case-insensitive
	id := string(bytes.ToLower(data[idOffset:idEnd]))
	p.refs[id] = &reference{
		link:  data[linkOffset:linkEnd],
		title: data[titleOffset:titleEnd],
	}

	return lineEnd
}


//
//
// Miscellaneous helper functions
//
//


// Test if a character is a punctuation symbol.
// Taken from a private function in regexp in the stdlib.
func ispunct(c byte) bool {
	for _, r := range []byte("!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~") {
		if c == r {
			return true
		}
	}
	return false
}

// Test if a character is a whitespace character.
func isspace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\n' || c == '\r' || c == '\f' || c == '\v'
}

// Test if a character is a letter or a digit.
// TODO: check when this is looking for ASCII alnum and when it should use unicode
func isalnum(c byte) bool {
	return (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

// Replace tab characters with spaces, aligning to the next TAB_SIZE column.
// always ends output with a newline
func expandTabs(out *bytes.Buffer, line []byte, tabSize int) {
	// first, check for common cases: no tabs, or only tabs at beginning of line
	i, prefix := 0, 0
	slowcase := false
	for i = 0; i < len(line); i++ {
		if line[i] == '\t' {
			if prefix == i {
				prefix++
			} else {
				slowcase = true
				break
			}
		}
	}

	// no need to decode runes if all tabs are at the beginning of the line
	if !slowcase {
		for i = 0; i < prefix*tabSize; i++ {
			out.WriteByte(' ')
		}
		out.Write(line[prefix:])
		return
	}

	// the slow case: we need to count runes to figure out how
	// many spaces to insert for each tab
	column := 0
	i = 0
	for i < len(line) {
		start := i
		for i < len(line) && line[i] != '\t' {
			_, size := utf8.DecodeRune(line[i:])
			i += size
			column++
		}

		if i > start {
			out.Write(line[start:i])
		}

		if i >= len(line) {
			break
		}

		for {
			out.WriteByte(' ')
			column++
			if column%tabSize == 0 {
				break
			}
		}

		i++
	}
}
