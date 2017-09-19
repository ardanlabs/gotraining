// Copyright ©2016 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.7

package native

import (
	"testing"

	"github.com/gonum/lapack/testlapack"
)

func BenchmarkDgeev(b *testing.B) { testlapack.DgeevBenchmark(b, impl) }
