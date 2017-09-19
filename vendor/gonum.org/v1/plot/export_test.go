// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plot

import (
	"gonum.org/v1/plot/vg/draw"
)

// Draw exports the Legend draw method for testing.
func (l *Legend) Draw(c draw.Canvas) { l.draw(c) }
