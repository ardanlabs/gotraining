// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter

import (
	"image/color"
	"log"
	"math/rand"
	"testing"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/internal/cmpimg"
	"gonum.org/v1/plot/vg"
)

func TestBubblesRadius(t *testing.T) {
	b := &Bubbles{
		MinRadius: vg.Length(0),
		MaxRadius: vg.Length(1),
	}

	tests := []struct {
		minz, maxz, z float64
		r             vg.Length
	}{
		{0, 0, 0, vg.Length(0.5)},
		{1, 1, 1, vg.Length(0.5)},
		{0, 1, 0, vg.Length(0)},
		{0, 1, 1, vg.Length(1)},
		{0, 1, 0.5, vg.Length(0.5)},
		{0, 2, 1, vg.Length(0.5)},
		{0, 4, 0, vg.Length(0)},
		{0, 4, 1, vg.Length(0.25)},
		{0, 4, 2, vg.Length(0.5)},
		{0, 4, 3, vg.Length(0.75)},
		{0, 4, 4, vg.Length(1)},
	}

	for _, test := range tests {
		b.MinZ, b.MaxZ = test.minz, test.maxz
		if r := b.radius(test.z); r != test.r {
			t.Errorf("Got incorrect radius (%g) on %v", r, test)
		}
	}
}

func ExampleBubbles() {
	rnd := rand.New(rand.NewSource(1))

	// randomTriples returns some random x, y, z triples
	// with some interesting kind of trend.
	randomTriples := func(n int) XYZs {
		data := make(XYZs, n)
		for i := range data {
			if i == 0 {
				data[i].X = rnd.Float64()
			} else {
				data[i].X = data[i-1].X + 2*rnd.Float64()
			}
			data[i].Y = data[i].X + 10*rnd.Float64()
			data[i].Z = data[i].X
		}
		return data
	}

	n := 10
	bubbleData := randomTriples(n)

	p, err := plot.New()
	if err != nil {
		log.Panic(err)
	}
	p.Title.Text = "Bubbles"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	bs, err := NewBubbles(bubbleData, vg.Points(1), vg.Points(20))
	if err != nil {
		log.Panic(err)
	}
	bs.Color = color.RGBA{R: 196, B: 128, A: 255}
	p.Add(bs)

	err = p.Save(200, 200, "testdata/bubbles.png")
	if err != nil {
		log.Panic(err)
	}
}

func TestBubbles(t *testing.T) {
	cmpimg.CheckPlot(ExampleBubbles, t, "bubbles.png")
}
