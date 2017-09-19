// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"sort"
	"testing"

	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/palette"
	"gonum.org/v1/plot/vg"
)

var visualDebug = flag.Bool("visual", false, "output images for benchmarks and test data")

type unitGrid struct{ mat.Matrix }

func (g unitGrid) Dims() (c, r int)   { r, c = g.Matrix.Dims(); return c, r }
func (g unitGrid) Z(c, r int) float64 { return g.Matrix.At(r, c) }
func (g unitGrid) X(c int) float64 {
	_, n := g.Matrix.Dims()
	if c < 0 || c >= n {
		panic("index out of range")
	}
	return float64(c)
}
func (g unitGrid) Y(r int) float64 {
	m, _ := g.Matrix.Dims()
	if r < 0 || r >= m {
		panic("index out of range")
	}
	return float64(r)
}

func TestHeatMapWithContour(t *testing.T) {
	if !*visualDebug {
		return
	}
	m := unitGrid{mat.NewDense(3, 4, []float64{
		2, 1, 4, 3,
		6, 7, 2, 5,
		9, 10, 11, 12,
	})}
	h := NewHeatMap(m, palette.Heat(12, 1))

	levels := []float64{1.5, 2.5, 3.5, 4.5, 5.5, 6.5, 7.5, 8.5, 9.5, 10.5, 11.5}
	c := NewContour(m, levels, palette.Rainbow(10, palette.Blue, palette.Red, 1, 1, 1))
	c.LineStyles[0].Width *= 5

	plt, _ := plot.New()

	plt.Add(h)
	plt.Add(c)
	plt.Add(NewGlyphBoxes())

	plt.X.Padding = 0
	plt.Y.Padding = 0
	plt.X.Max = 3.5
	plt.Y.Max = 2.5
	plt.Save(7, 7, "heat.svg")
}

func TestComplexContours(t *testing.T) {
	rnd := rand.New(rand.NewSource(1))

	if !*visualDebug {
		return
	}
	for _, n := range []float64{0, 1, 2, 4, 8, 16, 32} {
		data := make([]float64, 6400)
		for i := range data {
			r := float64(i/80) - 40
			c := float64(i%80) - 40

			data[i] = rnd.NormFloat64()*n + math.Hypot(r, c)
		}

		m := unitGrid{mat.NewDense(80, 80, data)}

		levels := []float64{-1, 3, 7, 9, 13, 15, 19, 23, 27, 31}
		c := NewContour(m, levels, palette.Rainbow(10, palette.Blue, palette.Red, 1, 1, 1))

		plt, _ := plot.New()
		plt.Add(c)

		plt.X.Padding = 0
		plt.Y.Padding = 0
		plt.X.Max = 79.5
		plt.Y.Max = 79.5
		plt.Save(7, 7, fmt.Sprintf("complex_contour-%v.svg", n))
	}
}

func unity(f float64) vg.Length { return vg.Length(f) }

func BenchmarkComplexContour0(b *testing.B)  { complexContourBench(0, b) }
func BenchmarkComplexContour1(b *testing.B)  { complexContourBench(1, b) }
func BenchmarkComplexContour2(b *testing.B)  { complexContourBench(2, b) }
func BenchmarkComplexContour4(b *testing.B)  { complexContourBench(4, b) }
func BenchmarkComplexContour8(b *testing.B)  { complexContourBench(8, b) }
func BenchmarkComplexContour16(b *testing.B) { complexContourBench(16, b) }
func BenchmarkComplexContour32(b *testing.B) { complexContourBench(32, b) }

var cp map[float64][]vg.Path

func complexContourBench(noise float64, b *testing.B) {
	rnd := rand.New(rand.NewSource(1))

	data := make([]float64, 6400)
	for i := range data {
		r := float64(i/80) - 40
		c := float64(i%80) - 40

		data[i] = rnd.NormFloat64()*noise + math.Hypot(r, c)
	}

	m := unitGrid{mat.NewDense(80, 80, data)}

	levels := []float64{-1, 3, 7, 9, 13, 15, 19, 23, 27, 31}

	var p map[float64][]vg.Path

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p = contourPaths(m, levels, unity, unity)
	}

	cp = p
}

func TestContourPaths(t *testing.T) {
	m := unitGrid{mat.NewDense(3, 4, []float64{
		2, 1, 4, 3,
		6, 7, 2, 5,
		9, 10, 11, 12,
	})}

	levels := []float64{1.5, 2.5, 3.5, 4.5, 5.5, 6.5, 7.5, 8.5, 9.5, 10.5}

	var (
		wantClosed = 2
		gotClosed  int
	)

	got := contourPaths(m, levels, unity, unity)
	for l, p := range got {
		sort.Sort(byLength(p))
		for i, c := range p {
			if isLoop(c) && isLoop(wantContours[l][i]) {
				if !circularPermutations(c[1:], wantContours[l][i][1:]) {
					t.Errorf("unexpected path:\n\tgot:%+v\n\twant:%+v", c, wantContours[l][i])
				}
			} else if !reflect.DeepEqual(c, wantContours[l][i]) && !reflect.DeepEqual(c, reverseOfPath(wantContours[l][i])) {
				t.Errorf("unexpected path:\n\tgot:%+v\n\twant:%+v", c, wantContours[l][i])
			}

			if isLoop(c) {
				gotClosed++
			}
		}
	}

	if gotClosed != wantClosed {
		t.Errorf("unexpected number of loops: got:%d want:%d", gotClosed, wantClosed)
	}
}

type byLength []vg.Path

func (p byLength) Len() int           { return len(p) }
func (p byLength) Less(i, j int) bool { return len(p[i]) < len(p[j]) }
func (p byLength) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func reverseOfPath(p vg.Path) vg.Path {
	rp := make(vg.Path, 0, len(p))
	for i := len(p) - 1; i >= 0; i-- {
		rp = append(rp, p[i])
	}
	rp[0].Type = vg.MoveComp
	rp[len(rp)-1].Type = vg.LineComp

	return rp
}

func circularPermutations(a, b vg.Path) bool {
	if len(a) != len(b) {
		return false
	}

	var off int

	var forward bool
	for i, pc := range b {
		if a[0] == pc {
			off = i
			forward = true
			break
		}
	}
	for i, pc := range a {
		if b[(off+i)%len(a)] != pc {
			forward = false
			break
		}
	}

	var reverse bool
	for i, pc := range b {
		if a[0] == pc {
			off = i
			reverse = true
			break
		}
	}
	for i, pc := range a {
		if b[(off-i+len(a))%len(a)] != pc {
			reverse = false
			break
		}
	}

	return forward || reverse
}

// Contour paths sorted by path length.
var wantContours = map[float64][]vg.Path{
	1.5: []vg.Path{
		vg.Path{
			{Type: vg.MoveComp, Pos: vg.Point{X: 1.1666666666666667, Y: 0}},
			{Type: vg.LineComp, Pos: vg.Point{X: 1.1, Y: 0.1}},
			{Type: vg.LineComp, Pos: vg.Point{X: 1, Y: 0.08333333333333333}},
			{Type: vg.LineComp, Pos: vg.Point{X: 0.9166666666666666, Y: 0.08333333333333333}},
			{Type: vg.LineComp, Pos: vg.Point{X: 0.5, Y: 0}},
		},
	},
	2.5: []vg.Path{
		vg.Path{
			{Type: vg.MoveComp, Pos: vg.Point{X: 1.5, Y: 0}},
			{Type: vg.LineComp, Pos: vg.Point{X: 1.3, Y: 0.3}},
			{Type: vg.LineComp, Pos: vg.Point{X: 1, Y: 0.25}},
			{Type: vg.LineComp, Pos: vg.Point{X: 0.75, Y: 0.25}},
			{Type: vg.LineComp, Pos: vg.Point{X: 0.125, Y: 0.125}},
			{Type: vg.LineComp, Pos: vg.Point{X: 0, Y: 0.125}},
		},
		vg.Path{
			{Type: vg.MoveComp, Pos: vg.Point{X: 2, Y: 1.0555555555555556}},
			{Type: vg.LineComp, Pos: vg.Point{X: 1.9545454545454546, Y: 1.0454545454545454}},
			{Type: vg.LineComp, Pos: vg.Point{X: 1.9, Y: 1}},
			{Type: vg.LineComp, Pos: vg.Point{X: 1.8333333333333333, Y: 0.8333333333333334}},
			{Type: vg.LineComp, Pos: vg.Point{X: 2, Y: 0.75}},
			{Type: vg.LineComp, Pos: vg.Point{X: 2.1666666666666665, Y: 0.8333333333333334}},
			{Type: vg.LineComp, Pos: vg.Point{X: 2.1666666666666665, Y: 1}},
			{Type: vg.LineComp, Pos: vg.Point{X: 2.0454545454545454, Y: 1.0454545454545454}},
			{Type: vg.LineComp, Pos: vg.Point{X: 2, Y: 1.0555555555555556}},
		},
	},
	3.5: []vg.Path{
		vg.Path{
			{Type: vg.MoveComp, Pos: vg.Point{X: 3, Y: 0.25}},
			{Type: vg.LineComp, Pos: vg.Point{X: 2.5, Y: 0.5}},
			{Type: vg.LineComp, Pos: vg.Point{X: 2.5, Y: 0}},
		},
		vg.Path{
			{Type: vg.MoveComp, Pos: vg.Point{X: 1.8333333333333333, Y: 0}},
			{Type: vg.LineComp, Pos: vg.Point{X: 1.5, Y: 0.5}},
			{Type: vg.LineComp, Pos: vg.Point{X: 1, Y: 0.4166666666666667}},
			{Type: vg.LineComp, Pos: vg.Point{X: 0.5833333333333334, Y: 0.4166666666666667}},
			{Type: vg.LineComp, Pos: vg.Point{X: 0.375, Y: 0.375}},
			{Type: vg.LineComp, Pos: vg.Point{X: 0, Y: 0.375}},
		},
		vg.Path{
			{Type: vg.MoveComp, Pos: vg.Point{X: 1.5, Y: 0.5}},
			{Type: vg.LineComp, Pos: vg.Point{X: 2, Y: 0.25}},
			{Type: vg.LineComp, Pos: vg.Point{X: 2.5, Y: 0.5}},
			{Type: vg.LineComp, Pos: vg.Point{X: 2.5, Y: 1}},
			{Type: vg.LineComp, Pos: vg.Point{X: 2.1363636363636362, Y: 1.1363636363636365}},
			{Type: vg.LineComp, Pos: vg.Point{X: 2, Y: 1.1666666666666667}},
			{Type: vg.LineComp, Pos: vg.Point{X: 1.8636363636363635, Y: 1.1363636363636365}},
			{Type: vg.LineComp, Pos: vg.Point{X: 1.7, Y: 1}},
			{Type: vg.LineComp, Pos: vg.Point{X: 1.5, Y: 0.5}},
		},
	},
	4.5: []vg.Path{
		vg.Path{
			{Type: vg.MoveComp, Pos: vg.Point{X: 0, Y: 0.625}},
			{Type: vg.LineComp, Pos: vg.Point{X: 0.375, Y: 0.625}},
			{Type: vg.LineComp, Pos: vg.Point{X: 0.5833333333333334, Y: 0.5833333333333334}},
			{Type: vg.LineComp, Pos: vg.Point{X: 1, Y: 0.5833333333333334}},
			{Type: vg.LineComp, Pos: vg.Point{X: 1.3571428571428572, Y: 0.6428571428571429}},
			{Type: vg.LineComp, Pos: vg.Point{X: 1.5, Y: 1}},
			{Type: vg.LineComp, Pos: vg.Point{X: 1.7727272727272727, Y: 1.2272727272727273}},
			{Type: vg.LineComp, Pos: vg.Point{X: 2, Y: 1.2777777777777777}},
			{Type: vg.LineComp, Pos: vg.Point{X: 2.227272727272727, Y: 1.2272727272727273}},
			{Type: vg.LineComp, Pos: vg.Point{X: 2.8333333333333335, Y: 1}},
			{Type: vg.LineComp, Pos: vg.Point{X: 2.8333333333333335, Y: 0.8333333333333334}},
			{Type: vg.LineComp, Pos: vg.Point{X: 3, Y: 0.75}},
		},
	},
	5.5: []vg.Path{
		vg.Path{
			{Type: vg.MoveComp, Pos: vg.Point{X: 0, Y: 0.875}},
			{Type: vg.LineComp, Pos: vg.Point{X: 0.125, Y: 0.875}},
			{Type: vg.LineComp, Pos: vg.Point{X: 0.75, Y: 0.75}},
			{Type: vg.LineComp, Pos: vg.Point{X: 1, Y: 0.75}},
			{Type: vg.LineComp, Pos: vg.Point{X: 1.2142857142857142, Y: 0.7857142857142857}},
			{Type: vg.LineComp, Pos: vg.Point{X: 1.3, Y: 1}},
			{Type: vg.LineComp, Pos: vg.Point{X: 1.6818181818181819, Y: 1.3181818181818181}},
			{Type: vg.LineComp, Pos: vg.Point{X: 2, Y: 1.3888888888888888}},
			{Type: vg.LineComp, Pos: vg.Point{X: 2.3181818181818183, Y: 1.3181818181818181}},
			{Type: vg.LineComp, Pos: vg.Point{X: 2.9, Y: 1.1}},
			{Type: vg.LineComp, Pos: vg.Point{X: 3, Y: 1.0714285714285714}},
		},
	},
	6.5: []vg.Path{
		vg.Path{
			{Type: vg.MoveComp, Pos: vg.Point{X: 0, Y: 1.1666666666666667}},
			{Type: vg.LineComp, Pos: vg.Point{X: 0.125, Y: 1.125}},
			{Type: vg.LineComp, Pos: vg.Point{X: 0.5, Y: 1}},
			{Type: vg.LineComp, Pos: vg.Point{X: 0.9166666666666666, Y: 0.9166666666666666}},
			{Type: vg.LineComp, Pos: vg.Point{X: 1, Y: 0.9166666666666666}},
			{Type: vg.LineComp, Pos: vg.Point{X: 1.0714285714285714, Y: 0.9285714285714286}},
			{Type: vg.LineComp, Pos: vg.Point{X: 1.1, Y: 1}},
			{Type: vg.LineComp, Pos: vg.Point{X: 1.5909090909090908, Y: 1.4090909090909092}},
			{Type: vg.LineComp, Pos: vg.Point{X: 2, Y: 1.5}},
			{Type: vg.LineComp, Pos: vg.Point{X: 2.409090909090909, Y: 1.4090909090909092}},
			{Type: vg.LineComp, Pos: vg.Point{X: 2.7, Y: 1.3}},
			{Type: vg.LineComp, Pos: vg.Point{X: 3, Y: 1.2142857142857142}},
		},
	},
	7.5: []vg.Path{
		vg.Path{
			{Type: vg.MoveComp, Pos: vg.Point{X: 0, Y: 1.5}},
			{Type: vg.LineComp, Pos: vg.Point{X: 0.375, Y: 1.375}},
			{Type: vg.LineComp, Pos: vg.Point{X: 0.75, Y: 1.25}},
			{Type: vg.LineComp, Pos: vg.Point{X: 1, Y: 1.1666666666666667}},
			{Type: vg.LineComp, Pos: vg.Point{X: 1.5, Y: 1.5}},
			{Type: vg.LineComp, Pos: vg.Point{X: 2, Y: 1.6111111111111112}},
			{Type: vg.LineComp, Pos: vg.Point{X: 2.5, Y: 1.5}},
			{Type: vg.LineComp, Pos: vg.Point{X: 3, Y: 1.3571428571428572}},
		},
	},
	8.5: []vg.Path{
		vg.Path{
			{Type: vg.MoveComp, Pos: vg.Point{X: 0, Y: 1.8333333333333333}},
			{Type: vg.LineComp, Pos: vg.Point{X: 0.25, Y: 1.75}},
			{Type: vg.LineComp, Pos: vg.Point{X: 0.625, Y: 1.625}},
			{Type: vg.LineComp, Pos: vg.Point{X: 1, Y: 1.5}},
			{Type: vg.LineComp, Pos: vg.Point{X: 1.3, Y: 1.7}},
			{Type: vg.LineComp, Pos: vg.Point{X: 1.6428571428571428, Y: 1.6428571428571428}},
			{Type: vg.LineComp, Pos: vg.Point{X: 2, Y: 1.7222222222222223}},
			{Type: vg.LineComp, Pos: vg.Point{X: 2.357142857142857, Y: 1.6428571428571428}},
			{Type: vg.LineComp, Pos: vg.Point{X: 2.611111111111111, Y: 1.6111111111111112}},
			{Type: vg.LineComp, Pos: vg.Point{X: 3, Y: 1.5}},
		},
	},
	9.5: []vg.Path{
		vg.Path{
			{Type: vg.MoveComp, Pos: vg.Point{X: 0.5, Y: 2}},
			{Type: vg.LineComp, Pos: vg.Point{X: 0.875, Y: 1.875}},
			{Type: vg.LineComp, Pos: vg.Point{X: 1, Y: 1.8333333333333333}},
			{Type: vg.LineComp, Pos: vg.Point{X: 1.1, Y: 1.9}},
			{Type: vg.LineComp, Pos: vg.Point{X: 1.7857142857142858, Y: 1.7857142857142858}},
			{Type: vg.LineComp, Pos: vg.Point{X: 2, Y: 1.8333333333333333}},
			{Type: vg.LineComp, Pos: vg.Point{X: 2.2142857142857144, Y: 1.7857142857142858}},
			{Type: vg.LineComp, Pos: vg.Point{X: 2.7222222222222223, Y: 1.7222222222222223}},
			{Type: vg.LineComp, Pos: vg.Point{X: 3, Y: 1.6428571428571428}},
		},
	},
	10.5: []vg.Path{
		vg.Path{
			{Type: vg.MoveComp, Pos: vg.Point{X: 1.5, Y: 2}},
			{Type: vg.LineComp, Pos: vg.Point{X: 1.9285714285714286, Y: 1.9285714285714286}},
			{Type: vg.LineComp, Pos: vg.Point{X: 2, Y: 1.9444444444444444}},
			{Type: vg.LineComp, Pos: vg.Point{X: 2.0714285714285716, Y: 1.9285714285714286}},
			{Type: vg.LineComp, Pos: vg.Point{X: 2.8333333333333335, Y: 1.8333333333333333}},
			{Type: vg.LineComp, Pos: vg.Point{X: 3, Y: 1.7857142857142858}},
		},
	},
}

var loopTests = []struct {
	c *contour

	want []*contour
}{
	{
		c: &contour{backward: path{{0, 0}}, forward: path{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}, {7, 7}, {8, 8}, {9, 9}}},
		want: []*contour{
			&contour{backward: path{{0, 0}}, forward: path{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}, {7, 7}, {8, 8}, {9, 9}}},
		},
	},
	{
		c: &contour{backward: path{{0, 0}}, forward: path{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}, {4, 4}, {7, 7}, {8, 8}, {9, 9}}},
		want: []*contour{
			&contour{backward: path{{4, 4}}, forward: path{{5, 5}, {6, 6}, {4, 4}}},
			&contour{backward: path{{0, 0}}, forward: path{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {7, 7}, {8, 8}, {9, 9}}},
		},
	},
	{
		c: &contour{backward: path{{0, 0}}, forward: path{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {3, 3}, {7, 7}, {1, 1}, {9, 9}}},
		want: []*contour{
			&contour{backward: path{{0, 0}}, forward: path{{1, 1}, {9, 9}}},
			&contour{backward: path{{3, 3}}, forward: path{{4, 4}, {5, 5}, {3, 3}}},
			&contour{backward: path{{1, 1}}, forward: path{{2, 2}, {3, 3}, {7, 7}, {1, 1}}},
		},
	},
	{
		c: &contour{backward: path{{0, 0}}, forward: path{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {2, 2}, {7, 7}, {2, 2}, {9, 9}}},
		want: []*contour{
			&contour{backward: path{{2, 2}}, forward: path{{7, 7}, {2, 2}}},
			&contour{backward: path{{0, 0}}, forward: path{{1, 1}, {2, 2}, {9, 9}}},
			&contour{backward: path{{2, 2}}, forward: path{{3, 3}, {4, 4}, {5, 5}, {2, 2}}},
		},
	},
	{
		// This test is a known failing case for exciseQuick.
		c: &contour{backward: path{{0, 0}}, forward: path{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}, {3, 3}, {8, 8}, {9, 9}, {5, 5}, {10, 10}}},
		want: []*contour{
			&contour{backward: path{{5, 5}}, forward: path{{10, 10}}},
			&contour{backward: path{{0, 0}}, forward: path{{1, 1}, {2, 2}, {3, 3}}},
			&contour{backward: path{{3, 3}}, forward: path{{4, 4}, {5, 5}, {6, 6}, {3, 3}}},
			&contour{backward: path{{3, 3}}, forward: path{{8, 8}, {9, 9}, {5, 5}, {6, 6}, {3, 3}}},
		},
	},
}

func (c testContour) String() string {
	var s string
	for i, p := range c {
		if i != 0 {
			s += ", "
		}
		s += fmt.Sprintf("%v", append(p.backward.reverse(), p.forward...))
		p.backward.reverse()
	}
	return s
}

func TestExciseLoops(t *testing.T) {
	for _, quick := range []bool{true, false} {
		for i, test := range loopTests {
			gotSet := make(contourSet)
			c := &contour{
				backward: append(path(nil), test.c.backward...),
				forward:  append(path(nil), test.c.forward...),
			}
			gotSet[c] = struct{}{}
			c.exciseLoops(gotSet, quick)
			var got []*contour
			for c := range gotSet {
				got = append(got, c)
			}
			sort.Sort(testContour(got))
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("unexpected loop excision result for %d quick=%t:\n\tgot:%v\n\twant:%v",
					i, quick, testContour(got), testContour(test.want))
			}
		}
	}
}

type testContour []*contour

func (c testContour) Len() int           { return len(c) }
func (c testContour) Less(i, j int) bool { return len(c[i].forward) < len(c[j].forward) }
func (c testContour) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
