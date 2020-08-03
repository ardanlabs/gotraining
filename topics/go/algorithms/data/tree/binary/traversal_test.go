package binary

import (
	"reflect"
	"testing"
)

var treeA Tree
var treeB Tree

func setTree() {
	sliceA := []int{40, 5, 10, 80, 62, 2, 45, 12, 23, 77, 3, 2}
	for _, n := range sliceA {
		treeA.Insert(n)
	}

	sliceB := []int{2, 1, 3}
	for _, n := range sliceB {
		treeB.Insert(n)
	}
}

func init() {
	setTree()
}

func TestInOrder(t *testing.T) {
	tests := map[string]struct {
		input Tree
		want  []int
	}{
		"Regular": {input: treeA, want: []int{2, 2, 3, 5, 10, 12, 23, 40, 45,
			62, 77, 80}},
		"Triangle": {input: treeB, want: []int{1, 2, 3}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := InOrder(&tc.input)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("In-order %s | expected: %v, got: %v", name, tc.want,
					got)
			}
		})
	}
}

func TestPreOrder(t *testing.T) {
	tests := map[string]struct {
		input Tree
		want  []int
	}{
		"Regular": {input: treeA, want: []int{40, 5, 2, 2, 3, 10, 12, 23, 80,
			62, 45, 77}},
		"Triangle": {input: treeB, want: []int{2, 1, 3}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := PreOrder(&tc.input)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("Pre-order %s | expected: %v, got: %v", name, tc.want,
					got)
			}
		})
	}
}

func TestPostOrder(t *testing.T) {
	tests := map[string]struct {
		input Tree
		want  []int
	}{
		"Regular": {input: treeA, want: []int{2, 3, 2, 23, 12, 10, 5, 45, 77,
			62, 80, 40}},
		"Triangle": {input: treeB, want: []int{1, 3, 2}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := PostOrder(&tc.input)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("Post-order %s | expected: %v, got: %v", name, tc.want,
					got)
			}
		})
	}
}
