// Copyright ©2020 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package text // import "gonum.org/v1/plot/text"

import "gonum.org/v1/plot/vg/draw"

// Plain is a text/plain handler.
type Plain = draw.PlainTextHandler

var _ draw.TextHandler = (*Plain)(nil)
