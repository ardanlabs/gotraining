// Copyright ©2016 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter

import (
	"image/color"
	"log"
	"math/rand"
	"testing"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/internal/cmpimg"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

// Example_timeSeries draws a time series.
func Example_timeSeries() {
	rnd := rand.New(rand.NewSource(1))

	// xticks defines how we convert and display time.Time values.
	xticks := plot.TimeTicks{Format: "2006-01-02\n15:04"}

	// randomPoints returns some random x, y points
	// with some interesting kind of trend.
	randomPoints := func(n int) XYs {
		const (
			month = 1
			day   = 1
			hour  = 1
			min   = 1
			sec   = 1
			nsec  = 1
		)
		pts := make(XYs, n)
		for i := range pts {
			date := time.Date(2007+i, month, day, hour, min, sec, nsec, time.UTC).Unix()
			pts[i].X = float64(date)
			pts[i].Y = float64(pts[i].X+10*rnd.Float64()) * 1e-9
		}
		return pts
	}

	n := 10
	data := randomPoints(n)

	p, err := plot.New()
	if err != nil {
		log.Panic(err)
	}
	p.Title.Text = "Time Series"
	p.X.Tick.Marker = xticks
	p.Y.Label.Text = "Number of Gophers\n(Billions)"
	p.Add(NewGrid())

	line, points, err := NewLinePoints(data)
	if err != nil {
		log.Panic(err)
	}
	line.Color = color.RGBA{G: 255, A: 255}
	points.Shape = draw.CircleGlyph{}
	points.Color = color.RGBA{R: 255, A: 255}

	p.Add(line, points)

	err = p.Save(10*vg.Centimeter, 5*vg.Centimeter, "testdata/timeseries.png")
	if err != nil {
		log.Panic(err)
	}
}

func TestTimeSeries(t *testing.T) {
	cmpimg.CheckPlot(Example_timeSeries, t, "timeseries.png")
}
