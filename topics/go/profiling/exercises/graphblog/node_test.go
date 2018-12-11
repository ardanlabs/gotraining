// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0
//
// Exercise provided by Phil Pearl
// https://syslog.ravelin.com/making-something-faster-56dd6b772b83

package graphblog

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

type edge struct {
	a string
	b string
}

// =============================================================================

type edges []edge

func (e edges) build(g graph) {
	for _, edge := range e {
		g.addEdge(edge.a, edge.b)
	}
}

// =============================================================================

func TestDiameter(t *testing.T) {
	tests := []struct {
		name        string
		edges       edges
		expDiameter int
	}{
		{
			name: "empty",
		},
		{
			name:        "1edge",
			edges:       edges{{"a", "b"}},
			expDiameter: 1,
		},
		{
			name:        "3inline",
			edges:       edges{{"a", "b"}, {"b", "c"}},
			expDiameter: 2,
		},
		{
			name:        "4inline",
			edges:       edges{{"a", "b"}, {"b", "c"}, {"c", "d"}},
			expDiameter: 3,
		},
		{
			name:        "triangle",
			edges:       edges{{"a", "b"}, {"b", "c"}, {"a", "c"}},
			expDiameter: 1,
		},
		{
			name:        "square",
			edges:       edges{{"a", "b"}, {"b", "c"}, {"c", "d"}, {"a", "d"}},
			expDiameter: 2,
		},
		{
			name:        "2loops",
			edges:       edges{{"a", "b"}, {"b", "c"}, {"c", "a"}, {"c", "d"}, {"d", "e"}, {"e", "c"}},
			expDiameter: 2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g := newGraph()
			test.edges.build(g)
			diameter := g.diameter()
			if diameter != test.expDiameter {
				t.Errorf("expected %d diameter, got %d", test.expDiameter, diameter)
			}
		})
	}
}

// =============================================================================

var diameter int

func BenchmarkDiameter(b *testing.B) {
	fmt.Println("Warning, this benchmark at 1s takes > 38s in its current form.")
	g := newGraph()
	f, err := os.Open("edges.txt")
	if err != nil {
		b.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		edge := strings.Fields(line)
		if len(edge) != 2 {
			b.Fatalf("expected 2 edges, got %d", len(edge))
		}

		g.addEdge(edge[0], edge[1])
	}

	diameter = g.diameter()
	if diameter != 82 {
		b.Fatalf("expected 82 for diameter, got %d", diameter)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		diameter = g.diameter()
	}
}
