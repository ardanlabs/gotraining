// Copyright ©2016 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter

import (
	"image/png"
	"log"
	"os"
	"testing"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/internal/cmpimg"
	"gonum.org/v1/plot/vg"
)

const runImageLaTeX = false

// An example of embedding an image in a plot.
func ExampleImage() {
	p, err := plot.New()
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	p.Title.Text = "A Logo"

	// load an image
	f, err := os.Open("testdata/image_plot_input.png")
	if err != nil {
		log.Fatalf("error opening image file: %v\n", err)
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		log.Fatalf("error decoding image file: %v\n", err)
	}

	p.Add(NewImage(img, 100, 100, 200, 200))

	const (
		w = 5 * vg.Centimeter
		h = 5 * vg.Centimeter
	)

	err = p.Save(5*vg.Centimeter, 5*vg.Centimeter, "testdata/image_plot.png")
	if err != nil {
		log.Fatalf("error saving image plot: %v\n", err)
	}
}

func TestImagePlot(t *testing.T) {
	cmpimg.CheckPlot(ExampleImage, t, "image_plot.png")
}
