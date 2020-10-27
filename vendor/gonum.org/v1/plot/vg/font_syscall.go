// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !js

package vg

import (
	"os"
	"path/filepath"
)

func init() {
	dirs := filepath.SplitList(os.Getenv("VGFONTPATH"))
	if len(dirs) > 0 {
		FontDirs = append(FontDirs, dirs...)
	}
}
