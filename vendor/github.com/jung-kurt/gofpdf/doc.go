/*
 * Copyright (c) 2013-2017 Kurt Jung (Gmail: kurt.w.jung)
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

/*
Package gofpdf implements a PDF document generator with high level support for
text, drawing and images.

Features

• Choice of measurement unit, page format and margins

• Page header and footer management

• Automatic page breaks, line breaks, and text justification

• Inclusion of JPEG, PNG, GIF, TIFF and basic path-only SVG images

• Colors, gradients and alpha channel transparency

• Outline bookmarks

• Internal and external links

• TrueType, Type1 and encoding support

• Page compression

• Lines, Bézier curves, arcs, and ellipses

• Rotation, scaling, skewing, translation, and mirroring

• Clipping

• Document protection

• Layers

• Templates

• Barcodes

• Charting facility

gofpdf has no dependencies other than the Go standard library. All tests pass
on Linux, Mac and Windows platforms.

Like FPDF version 1.7, from which gofpdf is derived, this package does not yet
support UTF-8 fonts. In particular, languages that require more than one code
page such as Chinese, Japanese, and Arabic are not currently supported. This is
explained in issue 109. However, support is provided to automatically translate
UTF-8 runes to code page encodings for languages that have fewer than 256
glyphs.

Installation

To install the package on your system, run

	go get github.com/jung-kurt/gofpdf

Later, to receive updates, run

	go get -u -v github.com/jung-kurt/gofpdf/...

Quick Start

The following Go code generates a simple PDF file.

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, world")
	err := pdf.OutputFileAndClose("hello.pdf")

See the functions in the fpdf_test.go file (shown as examples in this
documentation) for more advanced PDF examples.

Errors

If an error occurs in an Fpdf method, an internal error field is set. After
this occurs, Fpdf method calls typically return without performing any
operations and the error state is retained. This error management scheme
facilitates PDF generation since individual method calls do not need to be
examined for failure; it is generally sufficient to wait until after Output()
is called. For the same reason, if an error occurs in the calling application
during PDF generation, it may be desirable for the application to transfer the
error to the Fpdf instance by calling the SetError() method or the SetErrorf()
method. At any time during the life cycle of the Fpdf instance, the error state
can be determined with a call to Ok() or Err(). The error itself can be
retrieved with a call to Error().

Conversion Notes

This package is a relatively straightforward translation from the original FPDF
library written in PHP (despite the caveat in the introduction to Effective
Go). The API names have been retained even though the Go idiom would suggest
otherwise (for example, pdf.GetX() is used rather than simply pdf.X()). The
similarity of the two libraries makes the original FPDF website a good source
of information. It includes a forum and FAQ.

However, some internal changes have been made. Page content is built up using
buffers (of type bytes.Buffer) rather than repeated string concatenation.
Errors are handled as explained above rather than panicking. Output is
generated through an interface of type io.Writer or io.WriteCloser. A number of
the original PHP methods behave differently based on the type of the arguments
that are passed to them; in these cases additional methods have been exported
to provide similar functionality. Font definition files are produced in JSON
rather than PHP.

Example PDFs

A side effect of running "go test ./..." is the production of a number of
example PDFs. These can be found in the gofpdf/pdf directory after the tests
complete.

Please note that these examples run in the context of a test. In order run an
example as a standalone application, you'll need to examine fpdf_test.go for
some helper routines, for example exampleFilename() and summary().

Example PDFs can be compared with reference copies in order to verify that they
have been generated as expected. This comparison will be performed if a PDF
with the same name as the example PDF is placed in the gofpdf/pdf/reference
directory and if the third argument to ComparePDFFiles() in
internal/example/example.go is true. (By default it is false.) The routine that
summarizes an example will look for this file and, if found, will call
ComparePDFFiles() to check the example PDF for equality with its reference PDF.
If differences exist between the two files they will be printed to standard
output and the test will fail. If the reference file is missing, the comparison
is considered to succeed. In order to successfully compare two PDFs, the
placement of internal resources must be consistent and the internal creation
timestamps must be the same. To do this, the methods SetCatalogSort() and
SetCreationDate() need to be called for both files. This is done automatically
for all examples.

Nonstandard Fonts

Nothing special is required to use the standard PDF fonts (courier, helvetica,
times, zapfdingbats) in your documents other than calling SetFont().

In order to use a different TrueType or Type1 font, you will need to generate a
font definition file and, if the font will be embedded into PDFs, a compressed
version of the font file. This is done by calling the MakeFont function or
using the included makefont command line utility. To create the utility, cd
into the makefont subdirectory and run "go build". This will produce a
standalone executable named makefont. Select the appropriate encoding file from
the font subdirectory and run the command as in the following example.

	./makefont --embed --enc=../font/cp1252.map --dst=../font ../font/calligra.ttf

In your PDF generation code, call AddFont() to load the font and, as with the
standard fonts, SetFont() to begin using it. Most examples, including the
package example, demonstrate this method. Good sources of free, open-source
fonts include http://www.google.com/fonts/ and http://dejavu-fonts.org/.

Related Packages

The draw2d package (https://github.com/llgcode/draw2d) is a two dimensional
vector graphics library that can generate output in different forms. It uses
gofpdf for its document production mode.

Contributing Changes

gofpdf is a global community effort and you are invited to make it even better.
If you have implemented a new feature or corrected a problem, please consider
contributing your change to the project. A contribution that does not directly
pertain to the core functionality of gofpdf should be placed in its own
directory directly beneath the `contrib` directory.

Here are guidelines for making submissions. Your change should

• be compatible with the MIT License

• be properly documented

• be formatted with `go fmt`

• include an example in fpdf_test.go if appropriate

• conform to the standards of golint (https://github.com/golang/lint) and
go vet (https://godoc.org/golang.org/x/tools/cmd/vet), that is, `golint .` and
`go vet .` should not generate any warnings

• not diminish test coverage (https://blog.golang.org/cover)

Pull requests (https://help.github.com/articles/using-pull-requests/) work
nicely as a means of contributing your changes.

License

gofpdf is released under the MIT License. It is copyrighted by Kurt Jung and
the contributors acknowledged below.

Acknowledgments

This package's code and documentation are closely derived from the FPDF library
(http://www.fpdf.org/) created by Olivier Plathey, and a number of font and
image resources are copied directly from it. Bruno Michel has provided valuable
assistance with the code. Drawing support is adapted from the FPDF geometric
figures script by David Hernández Sanz. Transparency support is adapted from
the FPDF transparency script by Martin Hall-May. Support for gradients and
clipping is adapted from FPDF scripts by Andreas Würmser. Support for outline
bookmarks is adapted from Olivier Plathey by Manuel Cornes. Layer support is
adapted from Olivier Plathey. Support for transformations is adapted from the
FPDF transformation script by Moritz Wagner and Andreas Würmser. PDF
protection is adapted from the work of Klemen Vodopivec for the FPDF product.
Lawrence Kesteloot provided code to allow an image's extent to be determined
prior to placement. Support for vertical alignment within a cell was provided
by Stefan Schroeder. Ivan Daniluk generalized the font and image loading code
to use the Reader interface while maintaining backward compatibility. Anthony
Starks provided code for the Polygon function. Robert Lillack provided the
Beziergon function and corrected some naming issues with the internal curve
function. Claudio Felber provided implementations for dashed line drawing and
generalized font loading. Stani Michiels provided support for multi-segment
path drawing with smooth line joins, line join styles, enhanced fill modes, and
has helped greatly with package presentation and tests. Templating is adapted
by Marcus Downing from the FPDF_Tpl library created by Jan Slabon and Setasign.
Jelmer Snoeck contributed packages that generate a variety of barcodes and help
with registering images on the web. Jelmer Snoek and Guillermo Pascual
augmented the basic HTML functionality with aligned text. Kent Quirk
implemented backwards-compatible support for reading DPI from images that
support it, and for setting DPI manually and then having it properly taken into
account when calculating image size. Paulo Coutinho provided support for static
embedded fonts. Dan Meyers added support for embedded JavaScript. David Fish
added a generic alias-replacement function to enable, among other things, table
of contents functionality. Andy Bakun identified and corrected a problem in
which the internal catalogs were not sorted stably. Paul Montag added encoding
and decoding functionality for templates, including images that are embedded in
templates; this allows templates to be stored independently of gofpdf. Paul
also added support for page boxes used in printing PDF documents. Wojciech
Matusiak added supported for word spacing.

Roadmap

• Handle UTF-8 source text natively. Until then, automatic translation of
UTF-8 runes to code page bytes is provided.

• Improve test coverage as reported by the coverage tool.

*/
package gofpdf
