package kdtree

type heapNode struct {
	value    []float64
	length   float64
	srcRowNo int
}

type heap struct {
	tree []heapNode
}

// newHeap return a pointer of heap.
func newHeap() *heap {
	h := &heap{}
	h.tree = make([]heapNode, 0)
	return &heap{}
}

// maximum return the max heapNode in the heap.
func (h *heap) maximum() heapNode {
	if len(h.tree) == 0 {
		return heapNode{}
	}

	return h.tree[0]
}

// extractMax remove the Max heapNode in the heap.
func (h *heap) extractMax() {
	if len(h.tree) == 0 {
		return
	}

	h.tree[0] = h.tree[len(h.tree)-1]
	h.tree = h.tree[:len(h.tree)-1]

	target := 1
	for true {
		largest := target
		if target*2-1 >= len(h.tree) {
			break
		}
		if h.tree[target*2-1].length > h.tree[target-1].length {
			largest = target * 2
		}

		if target*2 < len(h.tree) {
			if h.tree[target*2].length > h.tree[largest-1].length {
				largest = target*2 + 1
			}
		}

		if largest == target {
			break
		}
		h.tree[largest-1], h.tree[target-1] = h.tree[target-1], h.tree[largest-1]
		target = largest
	}
}

// insert put a new heapNode into heap.
func (h *heap) insert(value []float64, length float64, srcRowNo int) {
	node := heapNode{}
	node.length = length
	node.srcRowNo = srcRowNo
	node.value = make([]float64, len(value))
	copy(node.value, value)
	h.tree = append(h.tree, node)

	target := len(h.tree)
	for target != 1 {
		if h.tree[(target/2)-1].length >= h.tree[target-1].length {
			break
		}
		h.tree[target-1], h.tree[(target/2)-1] = h.tree[(target/2)-1], h.tree[target-1]
		target /= 2
	}
}

func (h *heap) size() int {
	return len(h.tree)
}
