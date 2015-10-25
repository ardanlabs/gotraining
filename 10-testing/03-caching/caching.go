// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/QilJqGkQgb

// package caching provides code to show why CPU caches matter and the way
// the hardware caches memory affects performance.
package caching

import "fmt"

// Set the size of each row to be 64k.
const (
	cols = 64
	rows = 64 * 1024
)

// matrix represents a matrix with a large number
// of cache lines per row.
var matrix [cols][rows]byte

// init sets ~13% of the matrix to 0XFF.
func init() {
	var ctr int
	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			if row%8 == 0 {
				matrix[col][row] = 0xFF
				ctr++
			}
		}
	}

	fmt.Println(ctr, "Elements set out of", cols*rows)
}

// rowTraverse traverses the matrix linearly by each column for each row.
func rowTraverse() int {
	var ctr int

	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			if matrix[col][row] == 0xFF {
				ctr++
			}
		}
	}

	return ctr
}

// colTraverse traverses the matrix linearly by each row for each column.
func colTraverse() int {
	var ctr int

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if matrix[col][row] == 0xFF {
				ctr++
			}
		}
	}

	return ctr
}
