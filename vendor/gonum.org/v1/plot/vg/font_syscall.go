// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !js

package vg

import (
	"go/build"
	"os"
	"path/filepath"
)

const (
	// importString is the import string expected for
	// this package.  It is used to find the font
	// directory included with the package source.
	importString = "gonum.org/v1/plot/vg"
)

func init() {
	FontDirs = initFontDirs()
}

// InitFontDirs returns the initial value for the FontDirectories variable.
func initFontDirs() []string {
	dirs := filepath.SplitList(os.Getenv("VGFONTPATH"))

	if pkg, err := build.Import(importString, "", build.FindOnly); err == nil {
		p := filepath.Join(pkg.Dir, "fonts")
		if _, err := os.Stat(p); err == nil {
			dirs = append(dirs, p)
		}
	}

	if len(dirs) == 0 {
		dirs = []string{"./fonts"}
	}

	return dirs
}
