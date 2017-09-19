// Copyright ©2016 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"testing"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/internal/cmpimg"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
)

// ExampleSankey_sample creates a simple sankey diagram.
// The output can be found at https://github.com/gonum/plot/blob/master/plotter/testdata/sankeySimple_golden.png.
func ExampleSankey_simple() {
	p, err := plot.New()
	if err != nil {
		log.Panic(err)
	}

	// Define the stock categories
	const (
		treeType int = iota
		consumer
		fate
	)
	categoryLabels := []string{"Tree type", "Consumer", "Fate"}

	flows := []Flow{
		Flow{
			SourceCategory:   treeType,
			SourceLabel:      "Large",
			ReceptorCategory: consumer,
			ReceptorLabel:    "Mohamed",
			Value:            5,
		},
		Flow{
			SourceCategory:   treeType,
			SourceLabel:      "Small",
			ReceptorCategory: consumer,
			ReceptorLabel:    "Mohamed",
			Value:            2,
		},
		Flow{
			SourceCategory:   treeType,
			SourceLabel:      "Large",
			ReceptorCategory: consumer,
			ReceptorLabel:    "Sofia",
			Value:            3,
		},
		Flow{
			SourceCategory:   treeType,
			SourceLabel:      "Small",
			ReceptorCategory: consumer,
			ReceptorLabel:    "Sofia",
			Value:            1,
		},
		Flow{
			SourceCategory:   treeType,
			SourceLabel:      "Large",
			ReceptorCategory: consumer,
			ReceptorLabel:    "Wei",
			Value:            6,
		},
		Flow{
			SourceCategory:   consumer,
			SourceLabel:      "Mohamed",
			ReceptorCategory: fate,
			ReceptorLabel:    "Eaten",
			Value:            6,
		},
		Flow{
			SourceCategory:   consumer,
			SourceLabel:      "Mohamed",
			ReceptorCategory: fate,
			ReceptorLabel:    "Waste",
			Value:            1,
		},
		Flow{
			SourceCategory:   consumer,
			SourceLabel:      "Sofia",
			ReceptorCategory: fate,
			ReceptorLabel:    "Eaten",
			Value:            3,
		},
		Flow{
			SourceCategory:   consumer,
			SourceLabel:      "Sofia",
			ReceptorCategory: fate,
			ReceptorLabel:    "Waste",
			Value:            0.5, // An unbalanced flow
		},
		Flow{
			SourceCategory:   consumer,
			SourceLabel:      "Wei",
			ReceptorCategory: fate,
			ReceptorLabel:    "Eaten",
			Value:            5,
		},
		Flow{
			SourceCategory:   consumer,
			SourceLabel:      "Wei",
			ReceptorCategory: fate,
			ReceptorLabel:    "Waste",
			Value:            1,
		},
		Flow{
			SourceCategory:   treeType,
			SourceLabel:      "Large",
			ReceptorCategory: fate,
			ReceptorLabel:    "Waste",
			Value:            1,
		},
		Flow{
			SourceCategory:   treeType,
			SourceLabel:      "Small",
			ReceptorCategory: fate,
			ReceptorLabel:    "Waste",
			Value:            0.3,
		},
	}

	sankey, err := NewSankey(flows...)
	if err != nil {
		log.Panic(err)
	}
	p.Add(sankey)
	p.Y.Label.Text = "Number of apples"
	p.NominalX(categoryLabels...)
	err = p.Save(vg.Points(300), vg.Points(180), "testdata/sankeySimple.png")
	if err != nil {
		log.Panic(err)
	}
}

func TestSankey_simple(t *testing.T) {
	cmpimg.CheckPlot(ExampleSankey_simple, t, "sankeySimple.png")
}

// ExampleSankey_grouped creates a sankey diagram with grouped flows.
// The output can be found at https://github.com/gonum/plot/blob/master/plotter/testdata/sankeyGrouped_golden.png.
func ExampleSankey_grouped() {
	p, err := plot.New()
	if err != nil {
		log.Panic(err)
	}
	c := vgimg.New(vg.Points(300), vg.Points(180))
	dc := draw.New(c)

	// Define the stock categories
	const (
		treeType int = iota
		consumer
		fate
	)
	categoryLabels := []string{"Tree type", "Consumer", "Fate"}

	flows := []Flow{
		Flow{
			SourceCategory:   treeType,
			SourceLabel:      "LargeLargeLargeLargeLargeLargeLargeLargeLarge",
			ReceptorCategory: consumer,
			ReceptorLabel:    "Mohamed",
			Group:            "Apples",
			Value:            5,
		},
		Flow{
			SourceCategory:   treeType,
			SourceLabel:      "LargeLargeLargeLargeLargeLargeLargeLargeLarge",
			ReceptorCategory: consumer,
			ReceptorLabel:    "Mohamed",
			Group:            "Dates",
			Value:            3,
		},
		Flow{
			SourceCategory:   treeType,
			SourceLabel:      "Small",
			ReceptorCategory: consumer,
			ReceptorLabel:    "Mohamed",
			Group:            "Lychees",
			Value:            2,
		},
		Flow{
			SourceCategory:   treeType,
			SourceLabel:      "LargeLargeLargeLargeLargeLargeLargeLargeLarge",
			ReceptorCategory: consumer,
			ReceptorLabel:    "Sofia",
			Group:            "Apples",
			Value:            3,
		},
		Flow{
			SourceCategory:   treeType,
			SourceLabel:      "LargeLargeLargeLargeLargeLargeLargeLargeLarge",
			ReceptorCategory: consumer,
			ReceptorLabel:    "Sofia",
			Group:            "Dates",
			Value:            4,
		},
		Flow{
			SourceCategory:   treeType,
			SourceLabel:      "Small",
			ReceptorCategory: consumer,
			ReceptorLabel:    "Sofia",
			Group:            "Apples",
			Value:            1,
		},
		Flow{
			SourceCategory:   treeType,
			SourceLabel:      "LargeLargeLargeLargeLargeLargeLargeLargeLarge",
			ReceptorCategory: consumer,
			ReceptorLabel:    "Wei",
			Group:            "Lychees",
			Value:            6,
		},
		Flow{
			SourceCategory:   treeType,
			SourceLabel:      "Small",
			ReceptorCategory: consumer,
			ReceptorLabel:    "Wei",
			Group:            "Apples",
			Value:            3,
		},
		Flow{
			SourceCategory:   consumer,
			SourceLabel:      "Mohamed",
			ReceptorCategory: fate,
			ReceptorLabel:    "Eaten",
			Group:            "Apples",
			Value:            4,
		},
		Flow{
			SourceCategory:   consumer,
			SourceLabel:      "Mohamed",
			ReceptorCategory: fate,
			ReceptorLabel:    "Waste",
			Group:            "Apples",
			Value:            1,
		},
		Flow{
			SourceCategory:   consumer,
			SourceLabel:      "Mohamed",
			ReceptorCategory: fate,
			ReceptorLabel:    "Eaten",
			Group:            "Dates",
			Value:            3,
		},
		Flow{
			SourceCategory:   consumer,
			SourceLabel:      "Mohamed",
			ReceptorCategory: fate,
			ReceptorLabel:    "Waste",
			Group:            "Lychees",
			Value:            2,
		},
		Flow{
			SourceCategory:   consumer,
			SourceLabel:      "Sofia",
			ReceptorCategory: fate,
			ReceptorLabel:    "Eaten",
			Group:            "Apples",
			Value:            4,
		},
		Flow{
			SourceCategory:   consumer,
			SourceLabel:      "Sofia",
			ReceptorCategory: fate,
			ReceptorLabel:    "Eaten",
			Group:            "Dates",
			Value:            3,
		},
		Flow{
			SourceCategory:   consumer,
			SourceLabel:      "Sofia",
			ReceptorCategory: fate,
			ReceptorLabel:    "Waste",
			Group:            "Dates",
			Value:            1,
		},
		Flow{
			SourceCategory:   consumer,
			SourceLabel:      "Wei",
			ReceptorCategory: fate,
			ReceptorLabel:    "Eaten",
			Group:            "Lychees",
			Value:            6,
		},
		Flow{
			SourceCategory:   consumer,
			SourceLabel:      "Wei",
			ReceptorCategory: fate,
			ReceptorLabel:    "Eaten",
			Group:            "Apples",
			Value:            2,
		},
		Flow{
			SourceCategory:   consumer,
			SourceLabel:      "Wei",
			ReceptorCategory: fate,
			ReceptorLabel:    "Waste",
			Group:            "Apples",
			Value:            1,
		},
		Flow{
			SourceCategory:   treeType,
			SourceLabel:      "LargeLargeLargeLargeLargeLargeLargeLargeLarge",
			ReceptorCategory: fate,
			ReceptorLabel:    "Waste",
			Group:            "Apples",
			Value:            1,
		},
		Flow{
			SourceCategory:   treeType,
			SourceLabel:      "LargeLargeLargeLargeLargeLargeLargeLargeLarge",
			ReceptorCategory: fate,
			ReceptorLabel:    "Waste",
			Group:            "Dates",
			Value:            1,
		},
		Flow{
			SourceCategory:   treeType,
			SourceLabel:      "Small",
			ReceptorCategory: fate,
			ReceptorLabel:    "Waste",
			Group:            "Lychees",
			Value:            0.3,
		},
	}

	sankey, err := NewSankey(flows...)
	if err != nil {
		log.Panic(err)
	}

	// Here we specify the FLowStyle function to set the
	// colors of the different fruit groups.
	sankey.FlowStyle = func(group string) (color.Color, draw.LineStyle) {
		switch group {
		case "Lychees":
			return color.NRGBA{R: 242, G: 169, B: 178, A: 100}, sankey.LineStyle
		case "Apples":
			return color.NRGBA{R: 91, G: 194, B: 54, A: 100}, sankey.LineStyle
		case "Dates":
			return color.NRGBA{R: 112, G: 22, B: 0, A: 100}, sankey.LineStyle
		default:
			panic(fmt.Errorf("invalid group %s", group))
		}
	}

	// Here we set the StockStyle function to give an example of
	// setting a custom style for one of the stocks.
	sankey.StockStyle = func(label string, category int) (string, draw.TextStyle, vg.Length, vg.Length, color.Color, draw.LineStyle) {
		if label == "Small" && category == treeType {
			// Here we demonstrate how to rotate the label text
			// and change the style of the stock bar.
			ts := sankey.TextStyle
			ts.Rotation = 0.0
			ts.XAlign = draw.XRight
			ls := sankey.LineStyle
			ls.Color = color.White
			xOff := -sankey.StockBarWidth / 2
			yOff := vg.Length(0)
			return "small", ts, xOff, yOff, color.Black, ls
		}
		if label == "LargeLargeLargeLargeLargeLargeLargeLargeLarge" && category == treeType {
			// Here we demonstrate how to replace a long label that doesn't fit
			// in the existing space with a shorter version. Note that because
			// we are not able to account for the difference between the overall
			// canvas size and the size of the plotting area here, if a label
			// was only slightly larger than the available space, it would not
			// be caught and replaced.
			min, max, err := sankey.StockRange(label, category)
			if err != nil {
				log.Panic(err)
			}
			_, yTr := p.Transforms(&dc)
			barHeight := yTr(max) - yTr(min)
			if sankey.TextStyle.Font.Width(label) > barHeight {
				return "large", sankey.TextStyle, 0, 0, sankey.Color, sankey.LineStyle
			}
		}
		return label, sankey.TextStyle, 0, 0, sankey.Color, sankey.LineStyle
	}

	// Here we set the backgroud color for stocks from grey to white.
	sankey.Color = color.White

	p.Add(sankey)
	p.Y.Label.Text = "Number of fruit pieces"
	p.NominalX(categoryLabels...)

	legendLabels, thumbs := sankey.Thumbnailers()
	for i, l := range legendLabels {
		t := thumbs[i]
		p.Legend.Add(l, t)
	}
	p.Legend.Top = true
	p.X.Max = 3.05 // give room for the legend

	// Add boundary boxes for debugging.
	p.Add(NewGlyphBoxes())

	p.Draw(dc)
	pngimg := vgimg.PngCanvas{Canvas: c}
	f, err := os.Create("testdata/sankeyGrouped.png")
	if err != nil {
		log.Panic(err)
	}
	if _, err = pngimg.WriteTo(f); err != nil {
		log.Panic(err)
	}
}

func TestSankey_grouped(t *testing.T) {
	cmpimg.CheckPlot(ExampleSankey_grouped, t, "sankeyGrouped.png")
}
