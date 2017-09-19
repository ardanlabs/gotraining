package base

import (
	"bytes"
	"encoding/binary"
)

func sortXorOp(b []byte) []byte {
	ret := make([]byte, len(b))
	copy(ret, b)
	ret[0] ^= 0x80
	return ret
}

type sortSpec struct {
	r1 int
	r2 int
}

// Returns sortSpecs for inst in ascending order
func createSortSpec(inst FixedDataGrid, attrsArg []AttributeSpec) []sortSpec {

	attrs := make([]AttributeSpec, len(attrsArg))
	copy(attrs, attrsArg)
	// Reverse attribute order to be more intuitive
	for i, j := 0, len(attrs)-1; i < j; i, j = i+1, j-1 {
		attrs[i], attrs[j] = attrs[j], attrs[i]
	}

	_, rows := inst.Size()
	ret := make([]sortSpec, 0)

	// Create a buffer
	buf := bytes.NewBuffer(nil)
	ds := make([][]byte, rows)
	rs := make([]int, rows)
	rowSize := 0

	inst.MapOverRows(attrs, func(row [][]byte, rowNo int) (bool, error) {
		if rowSize == 0 {
			// Allocate a row buffer
			for _, r := range row {
				rowSize += len(r)
			}
		}

		byteBuf := make([]byte, rowSize)

		for i, r := range row {
			if i == 0 {
				binary.Write(buf, binary.LittleEndian, sortXorOp(r))
			} else {
				binary.Write(buf, binary.LittleEndian, r)
			}
		}

		buf.Read(byteBuf)
		ds[rowNo] = byteBuf
		rs[rowNo] = rowNo
		return true, nil
	})

	// Sort values
	valueBins := make([][][]byte, 256)
	rowBins := make([][]int, 256)
	for i := 0; i < rowSize; i++ {
		for j := 0; j < len(ds); j++ {
			// Address each row value by it's ith byte
			b := ds[j]
			valueBins[b[i]] = append(valueBins[b[i]], b)
			rowBins[b[i]] = append(rowBins[b[i]], rs[j])
		}
		j := 0
		for k := 0; k < 256; k++ {
			bs := valueBins[k]
			rc := rowBins[k]
			copy(ds[j:], bs)
			copy(rs[j:], rc)
			j += len(bs)
			valueBins[k] = bs[:0]
			rowBins[k] = rc[:0]
		}
	}

	done := make([]bool, rows)
	for index := range rs {
		if done[index] {
			continue
		}
		j := index
		for {
			done[j] = true
			if rs[j] != index {
				ret = append(ret, sortSpec{j, rs[j]})
				j = rs[j]
			} else {
				break
			}
		}
	}
	return ret
}

// Sort does a radix sort of DenseInstances, using SortDirection
// direction (Ascending or Descending) with attrs as a slice of Attribute
// indices that you want to sort by.
//
// IMPORTANT: Radix sort is not stable, so ordering outside
// the attributes used for sorting is arbitrary.
func Sort(inst FixedDataGrid, direction SortDirection, attrs []AttributeSpec) (FixedDataGrid, error) {
	sortInstructions := createSortSpec(inst, attrs)
	instUpdatable, ok := inst.(*DenseInstances)
	if ok {
		for _, i := range sortInstructions {
			instUpdatable.swapRows(i.r1, i.r2)
		}
		if direction == Descending {
			// Reverse the matrix
			_, rows := inst.Size()
			for i, j := 0, rows-1; i < j; i, j = i+1, j-1 {
				instUpdatable.swapRows(i, j)
			}
		}
	} else {
		panic("Sort is not supported for this yet!")
	}
	return instUpdatable, nil
}

// LazySort also does a sort, but returns an InstanceView and doesn't actually
// reorder the rows, just makes it look like they've been reordered
// See also: Sort
func LazySort(inst FixedDataGrid, direction SortDirection, attrs []AttributeSpec) (FixedDataGrid, error) {
	// Run the sort operation
	sortInstructions := createSortSpec(inst, attrs)

	// Build the row -> row mapping
	_, rows := inst.Size()      // Get the total row count
	rowArr := make([]int, rows) // Create an array of positions
	for i := 0; i < len(rowArr); i++ {
		rowArr[i] = i
	}
	for i := range sortInstructions {
		r1 := rowArr[sortInstructions[i].r1]
		r2 := rowArr[sortInstructions[i].r2]
		// Swap
		rowArr[sortInstructions[i].r1] = r2
		rowArr[sortInstructions[i].r2] = r1
	}
	if direction == Descending {
		for i, j := 0, rows-1; i < j; i, j = i+1, j-1 {
			tmp := rowArr[i]
			rowArr[i] = rowArr[j]
			rowArr[j] = tmp
		}
	}
	// Create a mapping dictionary
	rowMap := make(map[int]int)
	for i, a := range rowArr {
		if i == a {
			continue
		}
		rowMap[i] = a
	}
	// Create the return structure
	ret := NewInstancesViewFromRows(inst, rowMap)
	return ret, nil
}
