// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0
//
// Exercise provided by Phil Pearl
// https://syslog.ravelin.com/making-something-faster-56dd6b772b83

package graphblog

import (
	"container/list"
)

type node struct {
	id  string
	adj graph
}

func (n *node) add(adjNode *node) {
	n.adj[adjNode.id] = adjNode
}

// =============================================================================

type graph map[string]*node

func newGraph() graph {
	return make(graph)
}

func (g graph) get(id string) *node {
	if n, found := g[id]; found {
		return n
	}

	n := node{
		id:  id,
		adj: make(graph),
	}
	g[id] = &n
	return &n
}

func (g graph) addEdge(a, b string) {
	an := g.get(a)
	bn := g.get(b)
	an.add(bn)
	bn.add(an)
}

func (g graph) diameter() int {
	var diameter int
	for id := range g {
		if df := g.longestShortestPath(id); df > diameter {
			diameter = df
		}
	}
	return diameter
}

func (g graph) longestShortestPath(startID string) int {
	type bfsNode struct {
		parent *node
		depth  int
	}
	bfsData := make(map[string]*bfsNode, len(g))

	l := list.New()
	curNode := g.get(startID)
	bfsData[curNode.id] = &bfsNode{parent: curNode, depth: 0}
	l.PushBack(curNode)

	for {
		elt := l.Front()
		if elt == nil {
			break
		}
		curNode = l.Remove(elt).(*node)

		for id, m := range curNode.adj {
			if bm := bfsData[id]; bm == nil || bm.parent == nil {
				bfsData[id] = &bfsNode{parent: curNode, depth: bfsData[curNode.id].depth + 1}
				l.PushBack(m)
			}
		}
	}

	return bfsData[curNode.id].depth
}
