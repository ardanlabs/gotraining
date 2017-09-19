// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mat64

import (
	"fmt"
	"math/rand"
	"testing"
)

type dims struct{ r, c int }

var productTests = []struct {
	n       int
	factors []dims
	product dims
	panics  bool
}{
	{
		n:       1,
		factors: []dims{{3, 4}},
		product: dims{3, 4},
		panics:  false,
	},
	{
		n:       1,
		factors: []dims{{2, 4}},
		product: dims{3, 4},
		panics:  true,
	},
	{
		n:       3,
		factors: []dims{{10, 30}, {30, 5}, {5, 60}},
		product: dims{10, 60},
		panics:  false,
	},
	{
		n:       3,
		factors: []dims{{100, 30}, {30, 5}, {5, 60}},
		product: dims{10, 60},
		panics:  true,
	},
	{
		n:       7,
		factors: []dims{{60, 5}, {5, 5}, {5, 4}, {4, 10}, {10, 22}, {22, 45}, {45, 10}},
		product: dims{60, 10},
		panics:  false,
	},
	{
		n:       7,
		factors: []dims{{60, 5}, {5, 5}, {5, 400}, {4, 10}, {10, 22}, {22, 45}, {45, 10}},
		product: dims{60, 10},
		panics:  true,
	},
	{
		n:       3,
		factors: []dims{{1, 1000}, {1000, 2}, {2, 2}},
		product: dims{1, 2},
		panics:  false,
	},

	// Random chains.
	{
		n:       0,
		product: dims{0, 0},
		panics:  false,
	},
	{
		n:       2,
		product: dims{60, 10},
		panics:  false,
	},
	{
		n:       3,
		product: dims{60, 10},
		panics:  false,
	},
	{
		n:       4,
		product: dims{60, 10},
		panics:  false,
	},
	{
		n:       10,
		product: dims{60, 10},
		panics:  false,
	},
}

func TestProduct(t *testing.T) {
	for _, test := range productTests {
		dimensions := test.factors
		if dimensions == nil && test.n > 0 {
			dimensions = make([]dims, test.n)
			for i := range dimensions {
				if i != 0 {
					dimensions[i].r = dimensions[i-1].c
				}
				dimensions[i].c = rand.Intn(50) + 1
			}
			dimensions[0].r = test.product.r
			dimensions[test.n-1].c = test.product.c
		}
		factors := make([]Matrix, test.n)
		for i, d := range dimensions {
			data := make([]float64, d.r*d.c)
			for i := range data {
				data[i] = rand.Float64()
			}
			factors[i] = NewDense(d.r, d.c, data)
		}

		want := &Dense{}
		if !test.panics {
			a := &Dense{}
			for i, b := range factors {
				if i == 0 {
					want.Clone(b)
					continue
				}
				a, want = want, &Dense{}
				want.Mul(a, b)
			}
		}

		got := NewDense(test.product.r, test.product.c, nil)
		panicked, message := panics(func() {
			got.Product(factors...)
		})
		if test.panics {
			if !panicked {
				t.Errorf("fail to panic with product chain dimensions: %+v result dimension: %+v",
					dimensions, test.product)
			}
			continue
		} else if panicked {
			t.Errorf("unexpected panic %q with product chain dimensions: %+v result dimension: %+v",
				message, dimensions, test.product)
			continue
		}

		if len(factors) > 0 {
			p := newMultiplier(NewDense(test.product.r, test.product.c, nil), factors)
			p.optimize()
			gotCost := p.table.at(0, len(factors)-1).cost
			expr, wantCost, ok := bestExpressionFor(dimensions)
			if !ok {
				t.Fatal("unexpected number of expressions in brute force expression search")
			}
			if gotCost != wantCost {
				t.Errorf("unexpected cost for chain dimensions: %+v got: %v want: %v\n%s",
					dimensions, got, want, expr)
			}
		}

		if !EqualApprox(got, want, 1e-14) {
			t.Errorf("unexpected result from product chain dimensions: %+v", dimensions)
		}
	}
}

// node is a subexpression node.
type node struct {
	dims
	left, right *node
}

func (n *node) String() string {
	if n.left == nil || n.right == nil {
		rows, cols := n.shape()
		return fmt.Sprintf("[%d×%d]", rows, cols)
	}
	rows, cols := n.shape()
	return fmt.Sprintf("(%s * %s):[%d×%d]", n.left, n.right, rows, cols)
}

// shape returns the dimensions of the result of the subexpression.
func (n *node) shape() (rows, cols int) {
	if n.left == nil || n.right == nil {
		return n.r, n.c
	}
	rows, _ = n.left.shape()
	_, cols = n.right.shape()
	return rows, cols
}

// cost returns the cost to evaluate the subexpression.
func (n *node) cost() int {
	if n.left == nil || n.right == nil {
		return 0
	}
	lr, lc := n.left.shape()
	_, rc := n.right.shape()
	return lr*lc*rc + n.left.cost() + n.right.cost()
}

// expressionsFor returns a channel that can be used to iterate over all
// expressions of the given factor dimensions.
func expressionsFor(factors []dims) chan *node {
	if len(factors) == 1 {
		c := make(chan *node, 1)
		c <- &node{dims: factors[0]}
		close(c)
		return c
	}
	c := make(chan *node)
	go func() {
		for i := 1; i < len(factors); i++ {
			for left := range expressionsFor(factors[:i]) {
				for right := range expressionsFor(factors[i:]) {
					c <- &node{left: left, right: right}
				}
			}
		}
		close(c)
	}()
	return c
}

// catalan returns the nth 0-based Catalan number.
func catalan(n int) int {
	p := 1
	for k := n + 1; k < 2*n+1; k++ {
		p *= k
	}
	for k := 2; k < n+2; k++ {
		p /= k
	}
	return p
}

// bestExpressonFor returns the lowest cost expression for the given expression
// factor dimensions, the cost of the expression and whether the number of
// expressions searched matches the Catalan number for the number of factors.
func bestExpressionFor(factors []dims) (exp *node, cost int, ok bool) {
	const maxInt = int(^uint(0) >> 1)
	min := maxInt
	var best *node
	var n int
	for exp := range expressionsFor(factors) {
		n++
		cost := exp.cost()
		if cost < min {
			min = cost
			best = exp
		}
	}
	return best, min, n == catalan(len(factors)-1)
}
