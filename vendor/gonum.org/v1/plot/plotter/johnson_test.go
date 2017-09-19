// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func linksTo(i ...int) set {
	if len(i) == 0 {
		return nil
	}
	s := make(set)
	for _, v := range i {
		s[v] = struct{}{}
	}
	return s
}

func (s set) String() string {
	a := make([]int, 0, len(s))
	for v := range s {
		a = append(a, v)
	}
	sort.Ints(a)
	return fmt.Sprint(a)
}

var graphTests = []struct {
	path path
	g    graph

	// Tarjan tests.
	orderIsAmbiguous bool
	wantSCCs         [][]int
	wantAdj          graph

	// Johnson tests.
	wantCycles     [][]int
	wantCyclePaths []path
}{
	{
		path: path{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}, {7, 7}, {8, 8}, {9, 9}},

		wantSCCs: [][]int{{9}, {8}, {7}, {6}, {5}, {4}, {3}, {2}, {1}, {0}},
		wantAdj:  nil,

		wantCycles:     nil,
		wantCyclePaths: nil,
	},
	{
		path: path{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}, {4, 4}, {7, 7}, {8, 8}, {9, 9}},

		wantSCCs: [][]int{
			{10}, {9}, {8},
			{4, 5, 6},
			{3}, {2}, {1}, {0},
			{7 /*second point{4, 4}*/},
		},
		wantAdj: graph{
			4:  linksTo(5),
			5:  linksTo(6),
			6:  linksTo(4),
			10: nil,
		},

		wantCycles: [][]int{
			{4, 5, 6, 4},
		},
		wantCyclePaths: []path{
			{{4, 4}, {5, 5}, {6, 6}, {4, 4}},
		},
	},
	{
		path: path{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}, {4, 4}, {7, 7}, {2, 2}, {9, 9}},

		wantSCCs: [][]int{
			{10},
			{2, 3, 4, 5, 6, 8},
			{1}, {0},
			{7 /*second point{4, 4}*/},
			{9 /*second point{2, 2}*/},
		},
		wantAdj: graph{
			2:  linksTo(3),
			3:  linksTo(4),
			4:  linksTo(5, 8),
			5:  linksTo(6),
			6:  linksTo(4),
			8:  linksTo(2),
			10: nil,
		},

		wantCycles: [][]int{
			{4, 5, 6, 4},
			{2, 3, 4, 8, 2},
		},
		wantCyclePaths: []path{
			{{4, 4}, {5, 5}, {6, 6}, {4, 4}},
			{{2, 2}, {3, 3}, {4, 4}, {7, 7}, {2, 2}},
		},
	},
	{
		path: path{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {2, 2}, {7, 7}, {2, 2}, {9, 9}},

		wantSCCs: [][]int{{9},
			{2, 3, 4, 5, 7},
			{1}, {0},
			{6 /*second point{2, 2}*/},
			{8 /*third point{2, 2}*/},
		},
		wantAdj: graph{
			2: linksTo(3, 7),
			3: linksTo(4),
			4: linksTo(5),
			5: linksTo(2),
			7: linksTo(2),
			9: nil,
		},

		wantCycles: [][]int{
			{2, 7, 2},
			{2, 3, 4, 5, 2},
		},
		wantCyclePaths: []path{
			{{2, 2}, {7, 7}, {2, 2}},
			{{2, 2}, {3, 3}, {4, 4}, {5, 5}, {2, 2}},
		},
	},
	{
		path: path{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}, {3, 3}, {8, 8}, {9, 9}, {5, 5}, {10, 10}},

		wantSCCs: [][]int{
			{11},
			{3, 4, 5, 6, 8, 9},
			{2}, {1}, {0},
			{7 /*second point{4, 4}*/},
			{10 /*second point{2, 2}*/},
		},
		wantAdj: graph{
			3:  linksTo(4, 8),
			4:  linksTo(5),
			5:  linksTo(6),
			6:  linksTo(3),
			8:  linksTo(9),
			9:  linksTo(5),
			11: nil,
		},

		wantCycles: [][]int{
			{3, 4, 5, 6, 3},
			{3, 8, 9, 5, 6, 3},
		},
		wantCyclePaths: []path{
			{{3, 3}, {4, 4}, {5, 5}, {6, 6}, {3, 3}},
			{{3, 3}, {8, 8}, {9, 9}, {5, 5}, {6, 6}, {3, 3}},
		},
	},
	{
		path: path{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {0, 0}, {5, 5}, {6, 6}, {7, 7}, {5, 5}, {9, 9}, {10, 10}, {9, 9}},

		wantSCCs: [][]int{
			{9, 10},
			{5, 6, 7},
			{0, 1, 2, 3},
			{4 /*second point{0, 0}*/},
			{8 /*second point{5, 5}*/},
			{11 /*second point{9, 9}*/},
		},
		wantAdj: graph{
			0:  linksTo(1),
			1:  linksTo(2),
			2:  linksTo(3),
			3:  linksTo(0),
			5:  linksTo(6),
			6:  linksTo(7),
			7:  linksTo(5),
			9:  linksTo(10),
			10: linksTo(9),
			11: nil,
		},

		wantCycles: [][]int{
			{9, 10, 9},
			{5, 6, 7, 5},
			{0, 1, 2, 3, 0},
		},
		wantCyclePaths: []path{
			{{9, 9}, {10, 10}, {9, 9}},
			{{5, 5}, {6, 6}, {7, 7}, {5, 5}},
			{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {0, 0}},
		},
	},

	{
		g: graph{
			0: linksTo(1),
			1: linksTo(2, 7),
			2: linksTo(3, 6),
			3: linksTo(4),
			4: linksTo(2, 5),
			6: linksTo(3, 5),
			7: linksTo(0, 6),
		},

		wantSCCs: [][]int{
			{5},
			{2, 3, 4, 6},
			{0, 1, 7},
		},
		wantAdj: graph{
			0: linksTo(1),
			1: linksTo(7),
			2: linksTo(3, 6),
			3: linksTo(4),
			4: linksTo(2),
			6: linksTo(3),
			7: linksTo(0),
		},

		wantCycles: [][]int{
			{0, 1, 7, 0},
			{2, 3, 4, 2},
			{2, 6, 3, 4, 2},
		},
	},
	{
		g: graph{
			0: linksTo(1, 2, 3),
			1: linksTo(2),
			2: linksTo(3),
			3: linksTo(1),
		},

		wantSCCs: [][]int{
			{1, 2, 3},
			{0},
		},
		wantAdj: graph{
			1: linksTo(2),
			2: linksTo(3),
			3: linksTo(1),
		},

		wantCycles: [][]int{
			{1, 2, 3, 1},
		},
	},
	{
		g: graph{
			0: linksTo(1),
			1: linksTo(0, 2),
			2: linksTo(1),
		},

		wantSCCs: [][]int{
			{0, 1, 2},
		},
		wantAdj: graph{
			0: linksTo(1),
			1: linksTo(0, 2),
			2: linksTo(1),
		},

		wantCycles: [][]int{
			{0, 1, 0},
			{1, 2, 1},
		},
	},
	{
		g: graph{
			0: linksTo(1),
			1: linksTo(2, 3),
			2: linksTo(4, 5),
			3: linksTo(4, 5),
			4: linksTo(6),
			5: nil,
			6: nil,
		},

		orderIsAmbiguous: true,
		wantSCCs: [][]int{
			// Node pairs (2, 3) and (4, 5) are not
			// relatively orderable within each pair.
			{6}, {5}, {4}, {3}, {2}, {1}, {0},
		},
		wantAdj: nil,

		wantCycles: nil,
	},
	{
		g: graph{
			0: linksTo(1),
			1: linksTo(2, 3, 4),
			2: linksTo(0, 3),
			3: linksTo(4),
			4: linksTo(3),
		},

		orderIsAmbiguous: true,
		wantSCCs: [][]int{
			// SCCs are not relatively ordable.
			{3, 4}, {0, 1, 2},
		},
		wantAdj: graph{
			0: linksTo(1),
			1: linksTo(2),
			2: linksTo(0),
			3: linksTo(4),
			4: linksTo(3),
		},

		wantCycles: [][]int{
			{3, 4, 3},
			{0, 1, 2, 0},
		},
	},
}

func TestTarjan(t *testing.T) {
	for i, test := range graphTests {
		var g graph
		if test.path != nil {
			g = graphFrom(test.path)
		} else {
			g = test.g
		}
		tar := newTarjan(g)
		gotSCCs := tar.sccs
		if test.orderIsAmbiguous {
			// We lose topological order here, but that
			// is not important for this use case.
			sort.Sort(byComponentLengthOrStart(test.wantSCCs))
			sort.Sort(byComponentLengthOrStart(gotSCCs))
		}
		// tarjan.strongconnect does range iteration over maps,
		// so sort SCC members to ensure consistent ordering.
		for _, scc := range gotSCCs {
			sort.Ints(scc)
		}
		if !reflect.DeepEqual(gotSCCs, test.wantSCCs) {
			t.Errorf("unexpected tarjan scc result for %d:\n\tgot:%v\n\twant:%v", i, gotSCCs, test.wantSCCs)
		}
		gotAdj := tar.sccSubGraph(2)
		if !reflect.DeepEqual(gotAdj, test.wantAdj) {
			t.Errorf("unexpected tarjan sccSubGraph(2) result for %d:\n\tgot:%#v\n\twant:%#v", i, gotAdj, test.wantAdj)
		}
	}
}

func TestJohnson(t *testing.T) {
	for i, test := range graphTests {
		var g graph
		if test.path != nil {
			g = graphFrom(test.path)
		} else {
			g = test.g
		}
		gotCycles := cyclesIn(g)
		// johnson.circuit does range iteration over maps,
		// so sort to ensure consistent ordering.
		sort.Sort(byComponentLengthOrStart(gotCycles))
		if !reflect.DeepEqual(gotCycles, test.wantCycles) {
			t.Errorf("unexpected johnson result for %d:\n\tgot:%#v\n\twant:%#v", i, gotCycles, test.wantCycles)
		}

		// Don't do path reconstruction tests without a path definition.
		if test.path == nil {
			continue
		}

		// Test we reconstruct paths correctly from cycles.
		var gotPaths []path
		for _, pi := range gotCycles {
			gotPaths = append(gotPaths, test.path.subpath(pi))
		}
		if !reflect.DeepEqual(gotPaths, test.wantCyclePaths) {
			t.Errorf("unexpected johnson path result for %d:\n\tgot:%#v\n\twant:%#v", i, gotPaths, test.wantCyclePaths)
		}
	}
}

type byComponentLengthOrStart [][]int

func (c byComponentLengthOrStart) Len() int { return len(c) }
func (c byComponentLengthOrStart) Less(i, j int) bool {
	return len(c[i]) < len(c[j]) || (len(c[i]) == len(c[j]) && c[i][0] < c[j][0])
}
func (c byComponentLengthOrStart) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
