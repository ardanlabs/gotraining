package binary

import (
	"reflect"
	"testing"
)

func maketree(numbers []int) Tree {
	var tree Tree
	for _, n := range numbers {
		tree.Insert(n)
	}
	return tree

}

func TestInOrder(t *testing.T) {
	tree := maketree([]int{40, 5, 10, 80, 62, 2, 45, 12, 23, 77, 3, 2})
	tree1 := maketree([]int{13})
	tests := map[string]struct {
		input Tree
		want  []int
	}{
		"normal":      {input: tree, want: []int{2, 2, 3, 5, 10, 12, 23, 40, 45, 62, 77, 80}},
		"single node": {input: tree1, want: []int{13}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := InOrder(&tc.input)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s | expected: %v, got: %v", name, tc.want, got)
			}
		})
	}
}

func TestPreOrder(t *testing.T) {
	tree := maketree([]int{40, 5, 10, 80, 62, 2, 45, 12, 23, 77, 3, 2})
	tree1 := maketree([]int{13})
	tests := map[string]struct {
		input Tree
		want  []int
	}{
		"normal":      {input: tree, want: []int{40, 5, 2, 2, 3, 10, 12, 23, 80, 62, 45, 77}},
		"single node": {input: tree1, want: []int{13}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := PreOrder(&tc.input)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s | expected: %v, got: %v", name, tc.want, got)
			}
		})
	}
}

func TestPostOrder(t *testing.T) {
	tree := maketree([]int{40, 5, 10, 80, 62, 2, 45, 12, 23, 77, 3, 2})
	tree1 := maketree([]int{13})
	tests := map[string]struct {
		input Tree
		want  []int
	}{
		"normal": {input: tree, want: []int{2, 3, 2, 23, 12, 10, 5, 45, 77, 62, 80, 40}},
		"1 node": {input: tree1, want: []int{13}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := PostOrder(&tc.input)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s | expected: %v, got: %v", name, tc.want, got)
			}
		})
	}
}
