package binary

import (
	"fmt"
	"math"
)

// PrettyPrint takes a Tree value and displays a pretty print
// version of the tree.
func PrettyPrint(t Tree) {

	// Build an index map of positions for print layout.
	values := make(map[int]int)
	maxIdx := buildIndexMap(values, 0, 0, t.root)

	// Calculate the total number of levels based on
	// the max index provided.
	var levels int
	for {
		pow := math.Pow(2, float64(levels))
		if maxIdx < int(pow) {
			break
		}
		levels++
	}
	levels--

	// Capture the positional data to use.
	data := generateData(levels)

	// Set the edge of the top of the tree.
	for sp := 0; sp < data[0].edge; sp++ {
		fmt.Print(" ")
	}
	fmt.Printf("%02d", values[0])
	fmt.Print("\n")

	dataIdx := 1
	for i := 1; i < len(data); i = i + 2 {

		// Set the edge of this row.
		for sp := 0; sp < data[i].edge; sp++ {
			fmt.Print(" ")
		}

		// Draw the hashes for this row.
		dataHashIdx := dataIdx
		for h := 0; h < data[i].draw; h++ {
			if values[dataHashIdx] != maxInt {
				fmt.Printf("/")
			} else {
				fmt.Printf(" ")
			}
			for sp := 0; sp < data[i].padding; sp++ {
				fmt.Print(" ")
			}
			if values[dataHashIdx+1] != maxInt {
				fmt.Printf("\\")
			} else {
				fmt.Printf(" ")
			}
			dataHashIdx += 2

			if data[i].gaps != 0 && data[i].gaps > h {
				for sp := 0; sp < data[i].gapPad; sp++ {
					fmt.Print(" ")
				}
			}
		}
		fmt.Print("\n")

		// Set the edge of the next row.
		for sp := 0; sp < data[i+1].edge; sp++ {
			fmt.Print(" ")
		}

		// Draw the numbers for this row.
		for n := 0; n < data[i+1].draw; n++ {
			if values[dataIdx] != maxInt {
				fmt.Printf("%02d", values[dataIdx])
			} else {
				fmt.Printf("  ")
			}
			for sp := 0; sp < data[i+1].padding; sp++ {
				fmt.Print(" ")
			}
			if values[dataIdx+1] != maxInt {
				fmt.Printf("%02d", values[dataIdx+1])
			} else {
				fmt.Printf("  ")
			}
			dataIdx += 2

			if data[i+1].gaps != 0 && data[i+1].gaps > n {
				for sp := 0; sp < data[i+1].gapPad; sp++ {
					fmt.Print(" ")
				}
			}
		}
		fmt.Print("\n")
	}

	fmt.Print("\n")
}

const maxInt = int(^uint(0) >> 1)

// buildIndex traverses the tree and generates a map of index positions
// for each node in the tree for printing.
//          40
//       /      \
//      05      80
//     /  \    /  \
//    02  25  65  98
// values{0:40, 1:05, 2:80, 3:02, 4:25, 5:65, 6:98}
func buildIndexMap(values map[int]int, idx int, maxIdx int, n *node) int {

	// We need to keep track of the highest index position used
	// to help calculate tree depth.
	if idx > maxIdx {
		maxIdx = idx
	}

	// We have reached the end of a branch. Use the maxInt to mark
	// no value in that position.
	if n == nil {
		values[idx] = maxInt
		return maxIdx
	}

	// Save the value of this node in the map at the
	// calculated index position.
	values[idx] = n.data.Key

	// Check if there are still nodes to check down the left
	// branch. When we move down the tree, the next index doubles.
	if n.left != nil {
		nextidx := 2*idx + 1
		maxIdx = buildIndexMap(values, nextidx, maxIdx, n.left)
	}

	// Check if there are still nodes to check down the right
	// branch. When we move down the tree, the next index doubles.
	nextidx := 2*idx + 2
	maxIdx = buildIndexMap(values, nextidx, maxIdx, n.right)

	// We need to set missing indexes in the map to maxInt.
	// So they are ignored in the printing of the map.
	if idx == 0 {
		for i := 0; i < maxIdx; i++ {
			if _, ok := values[i]; !ok {
				values[i] = maxInt
			}
		}
	}

	return maxIdx
}

// pos provides positional data for printing a tree.
type pos struct {
	edge    int
	draw    int
	padding int
	gaps    int
	gapPad  int
}

// generateData generates all the positional data needed to display
// nodes at different levels.
func generateData(level int) []pos {
	totalData := (level * 2) - 1
	data := make([]pos, totalData)
	edge := 1
	draw := level - 2
	padding := 0
	gapPad := 2

	for i := totalData - 1; i >= 0; i = i - 2 {

		// Generate starting edge positions.
		data[i].edge = int(math.Pow(2, float64(edge)))
		if i > 0 {
			data[i-1].edge = data[i].edge + 1
		}
		edge++

		// Generate draw information.
		if draw > 0 {
			data[i].draw = int(math.Pow(2, float64(draw)))
			data[i-1].draw = data[i].draw
		} else {
			data[i].draw = 1
			if i > 0 {
				data[i-1].draw = 1
			}
		}
		draw--

		// Generate padding information.
		padding += data[i].edge
		data[i].padding = padding
		if i > 0 {
			data[i-1].padding = padding
		}

		// Generate gaps information.
		data[i].gaps = data[i].draw - 1
		if i > 0 {
			data[i-1].gaps = data[i].gaps
		}

		// Generate gap padding information.
		if i > 2 {
			data[i-1].gapPad = int(math.Pow(2, float64(gapPad)))
			data[i].gapPad = data[i-1].gapPad - 2
		}
		gapPad++
	}

	return data
}
