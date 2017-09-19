// Copyright 2010 The draw2d Authors. All rights reserved.
// created: 13/12/2010 by Laurent Le Goff

package draw2dimg

import (
	"github.com/golang/freetype/raster"
	"golang.org/x/image/math/fixed"
)

type FtLineBuilder struct {
	Adder raster.Adder
}

func (liner FtLineBuilder) MoveTo(x, y float64) {
	liner.Adder.Start(fixed.Point26_6{X: fixed.Int26_6(x * 64), Y: fixed.Int26_6(y * 64)})
}

func (liner FtLineBuilder) LineTo(x, y float64) {
	liner.Adder.Add1(fixed.Point26_6{X: fixed.Int26_6(x * 64), Y: fixed.Int26_6(y * 64)})
}

func (liner FtLineBuilder) LineJoin() {
}

func (liner FtLineBuilder) Close() {
}

func (liner FtLineBuilder) End() {
}
