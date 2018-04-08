draw2d
======
[![Coverage](http://gocover.io/_badge/github.com/llgcode/draw2d?0)](http://gocover.io/github.com/llgcode/draw2d)
[![GoDoc](https://godoc.org/github.com/llgcode/draw2d?status.svg)](https://godoc.org/github.com/llgcode/draw2d)

Package draw2d is a pure [go](http://golang.org) 2D vector graphics library with support for multiple output devices such as [images](http://golang.org/pkg/image) (draw2d), pdf documents (draw2dpdf) and opengl (draw2dgl), which can also be used on the google app engine. It can be used as a pure go [Cairo](http://www.cairographics.org/) alternative. draw2d is released under the BSD license. See the [documentation](http://godoc.org/github.com/llgcode/draw2d) for more details.

[![geometry](https://raw.githubusercontent.com/llgcode/draw2d/master/output/samples/geometry.png)](https://raw.githubusercontent.com/llgcode/draw2d/master/resource/image/geometry.pdf)[![postscript](https://raw.githubusercontent.com/llgcode/draw2d/master/output/samples/postscript.png)](https://raw.githubusercontent.com/llgcode/draw2d/master/resource/image/postscript.pdf)

Click on an image above to get the pdf, generated with exactly the same draw2d code. The first image is the output of `samples/geometry`. The second image is the result of `samples/postcript`, which demonstrates that draw2d can draw postscript files into images or pdf documents with the [ps](https://github.com/llgcode/ps) package.

Features
--------

Operations in draw2d include stroking and filling polygons, arcs, Bézier curves, drawing images and text rendering with truetype fonts. All drawing operations can be transformed by affine transformations (scale, rotation, translation).

Package draw2d follows the conventions of the [HTML Canvas 2D Context](http://www.w3.org/TR/2dcontext/) for coordinate system, angles, etc...

Installation
------------

Install [golang](http://golang.org/doc/install). To install or update the package draw2d on your system, run:

Stable release
```
go get -u gopkg.in/llgcode/draw2d.v1
```

or Current release
```
go get -u github.com/llgcode/draw2d
```


Quick Start
-----------

The following Go code generates a simple drawing and saves it to an image file with package draw2d:

```go
package main

import (
	"github.com/llgcode/draw2d/draw2dimg"
	"image"
	"image/color"
)

func main() {
	// Initialize the graphic context on an RGBA image
	dest := image.NewRGBA(image.Rect(0, 0, 297, 210.0))
	gc := draw2dimg.NewGraphicContext(dest)

	// Set some properties
	gc.SetFillColor(color.RGBA{0x44, 0xff, 0x44, 0xff})
	gc.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
	gc.SetLineWidth(5)

	// Draw a closed shape
	gc.BeginPath() // Initialize a new path
	gc.MoveTo(10, 10) // Move to a position to start the new path
	gc.LineTo(100, 50)
	gc.QuadCurveTo(100, 10, 10, 10)
	gc.Close()
	gc.FillStroke()

	// Save to file
	draw2dimg.SaveToPngFile("hello.png", dest)
}
```

The same Go code can also generate a pdf document with package draw2dpdf:

```go
package main

import (
	"github.com/llgcode/draw2d/draw2dpdf"
	"image/color"
)

func main() {
	// Initialize the graphic context on an RGBA image
	dest := draw2dpdf.NewPdf("L", "mm", "A4")
	gc := draw2dpdf.NewGraphicContext(dest)

	// Set some properties
	gc.SetFillColor(color.RGBA{0x44, 0xff, 0x44, 0xff})
	gc.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
	gc.SetLineWidth(5)

	// Draw a closed shape
	gc.MoveTo(10, 10) // should always be called first for a new path
	gc.LineTo(100, 50)
	gc.QuadCurveTo(100, 10, 10, 10)
	gc.Close()
	gc.FillStroke()

	// Save to file
	draw2dpdf.SaveToPdfFile("hello.pdf", dest)
}
```

There are more examples here: https://github.com/llgcode/draw2d/tree/master/samples

Drawing on opengl is provided by the draw2dgl package.

Testing
-------

The samples are run as tests from the root package folder `draw2d` by:
```
go test ./...
```
Or if you want to run with test coverage:
```
go test -cover ./... | grep -v "no test"
```
This will generate output by the different backends in the output folder.

Acknowledgments
---------------

[Laurent Le Goff](https://github.com/llgcode) wrote this library, inspired by [Postscript](http://www.tailrecursive.org/postscript) and [HTML5 canvas](http://www.w3.org/TR/2dcontext/). He implemented the image and opengl backend with the [freetype-go](https://code.google.com/p/freetype-go/) package. Also he created a pure go [Postscript interpreter](https://github.com/llgcode/ps), which can read postscript images and draw to a draw2d graphic context. [Stani Michiels](https://github.com/stanim) implemented the pdf backend with the [gofpdf](https://github.com/jung-kurt/gofpdf) package.



Packages using draw2d
---------------------

 - [ps](https://github.com/llgcode/ps): Postscript interpreter written in Go
 - [gonum/plot](https://github.com/gonum/plot): drawing plots in Go
 - [go.uik](https://github.com/skelterjohn/go.uik): a concurrent UI kit written in pure go.
 - [smartcrop](https://github.com/muesli/smartcrop): content aware image cropping
 - [karta](https://github.com/peterhellberg/karta): drawing Voronoi diagrams
 - [chart](https://github.com/vdobler/chart): basic charts in Go
 - [hilbert](https://github.com/google/hilbert): package for drawing Hilbert curves

References
---------

 - [antigrain.com](http://www.antigrain.com)
 - [freetype-go](http://code.google.com/p/freetype-go)
 -
