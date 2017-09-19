package draw

import (
	"image/color"
	"reflect"
	"testing"

	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/recorder"
)

func TestCrop(t *testing.T) {
	ls := LineStyle{
		Color: color.NRGBA{0, 20, 0, 123},
		Width: 0.1 * vg.Inch,
	}
	var r1 recorder.Canvas
	c1 := NewCanvas(&r1, 6, 3)
	c11 := Crop(c1, 0, -3, 0, 0)
	c12 := Crop(c1, 3, 0, 0, 0)

	var r2 recorder.Canvas
	c2 := NewCanvas(&r2, 6, 3)
	c21 := Canvas{
		Canvas: c2.Canvas,
		Rectangle: vg.Rectangle{
			Min: vg.Point{X: 0, Y: 0},
			Max: vg.Point{X: 3, Y: 3},
		},
	}
	c22 := Canvas{
		Canvas: c2.Canvas,
		Rectangle: vg.Rectangle{
			Min: vg.Point{X: 3, Y: 0},
			Max: vg.Point{X: 6, Y: 3},
		},
	}
	str := "unexpected result: %+v != %+v"
	if c11.Rectangle != c21.Rectangle {
		t.Errorf(str, c11.Rectangle, c21.Rectangle)
	}
	if c12.Rectangle != c22.Rectangle {
		t.Errorf(str, c11.Rectangle, c21.Rectangle)
	}

	c11.StrokeLine2(ls, c11.Min.X, c11.Min.Y,
		c11.Min.X+3*vg.Inch, c11.Min.Y+3*vg.Inch)
	c12.StrokeLine2(ls, c12.Min.X, c12.Min.Y,
		c12.Min.X+3*vg.Inch, c12.Min.Y+3*vg.Inch)
	c21.StrokeLine2(ls, c21.Min.X, c21.Min.Y,
		c21.Min.X+3*vg.Inch, c21.Min.Y+3*vg.Inch)
	c22.StrokeLine2(ls, c22.Min.X, c22.Min.Y,
		c22.Min.X+3*vg.Inch, c22.Min.Y+3*vg.Inch)

	if !reflect.DeepEqual(r1.Actions, r2.Actions) {
		t.Errorf(str, r1.Actions, r2.Actions)
	}
}

func TestTile(t *testing.T) {
	var r recorder.Canvas
	c := NewCanvas(&r, 13, 7)
	const (
		rows = 2
		cols = 3
		pad  = 1
	)
	tiles := Tiles{
		Rows: rows, Cols: cols,
		PadTop: pad, PadBottom: pad,
		PadRight: pad, PadLeft: pad,
		PadX: pad, PadY: pad,
	}
	rectangles := [][]vg.Rectangle{
		{
			vg.Rectangle{
				Min: vg.Point{X: 1, Y: 4},
				Max: vg.Point{X: 4, Y: 6},
			},
			vg.Rectangle{
				Min: vg.Point{X: 5, Y: 4},
				Max: vg.Point{X: 8, Y: 6},
			},
			vg.Rectangle{
				Min: vg.Point{X: 9, Y: 4},
				Max: vg.Point{X: 12, Y: 6},
			},
		},
		{
			vg.Rectangle{
				Min: vg.Point{X: 1, Y: 1},
				Max: vg.Point{X: 4, Y: 3},
			},
			vg.Rectangle{
				Min: vg.Point{X: 5, Y: 1},
				Max: vg.Point{X: 8, Y: 3},
			},
			vg.Rectangle{
				Min: vg.Point{X: 9, Y: 1},
				Max: vg.Point{X: 12, Y: 3},
			},
		},
	}
	for j := 0; j < rows; j++ {
		for i := 0; i < cols; i++ {
			str := "row %d col %d unexpected result: %+v != %+v"
			tile := tiles.At(c, i, j)
			if tile.Rectangle != rectangles[j][i] {
				t.Errorf(str, j, i, tile.Rectangle, rectangles[j][i])
			}
		}
	}
}
