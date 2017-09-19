// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter

import (
	"log"
	"math/rand"
	"testing"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/internal/cmpimg"
	"gonum.org/v1/plot/vg/draw"
)

// ExampleErrors draws points and error bars.
func ExampleErrors() {
	rnd := rand.New(rand.NewSource(1))

	randomError := func(n int) Errors {
		err := make(Errors, n)
		for i := range err {
			err[i].Low = rnd.Float64()
			err[i].High = rnd.Float64()
		}
		return err
	}
	// randomPoints returns some random x, y points
	// with some interesting kind of trend.
	randomPoints := func(n int) XYs {
		pts := make(XYs, n)
		for i := range pts {
			if i == 0 {
				pts[i].X = rnd.Float64()
			} else {
				pts[i].X = pts[i-1].X + rnd.Float64()
			}
			pts[i].Y = pts[i].X + 10*rnd.Float64()
		}
		return pts
	}

	type errPoints struct {
		XYs
		YErrors
		XErrors
	}

	n := 15
	data := errPoints{
		XYs:     randomPoints(n),
		YErrors: YErrors(randomError(n)),
		XErrors: XErrors(randomError(n)),
	}

	p, err := plot.New()
	if err != nil {
		log.Panic(err)
	}
	scatter, err := NewScatter(data)
	if err != nil {
		log.Panic(err)
	}
	scatter.Shape = draw.CrossGlyph{}
	xerrs, err := NewXErrorBars(data)
	if err != nil {
		log.Panic(err)
	}
	yerrs, err := NewYErrorBars(data)
	if err != nil {
		log.Panic(err)
	}
	p.Add(scatter, xerrs, yerrs)

	err = p.Save(200, 200, "testdata/errorBars.png")
	if err != nil {
		log.Panic(err)
	}
}

func TestErrors(t *testing.T) {
	cmpimg.CheckPlot(ExampleErrors, t, "errorBars.png")
}
