// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/GQQXh3cf15

// package caching provides code to show why CPU caches matter and the way
// the hardware caches memory affects performance.
package caching

const rows = 64

// Set the size of each row to be 4k.
// const cols = 4*1024

// Set the size of each row to be 64k.
// const cols = 64 * 1024

// matrix represents a set of columns that each exist on
// their own cache line.
var matrix [rows][cols]byte

// init sets ~13% of the matrix to 0XFF.
func init() {
	var ctr int
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if col%8 == 0 {
				matrix[row][col] = 0xFF
				ctr++
			}
		}
	}
	println(ctr, "Elements set out of", cols*rows)
}

// rowTraverse traverses the matrix linearly by each column for each row.
func rowTraverse() int {
	var ctr int

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if matrix[row][col] == 0xFF {
				ctr++
			}
		}
	}

	return ctr
}

// colTraverse traverses the matrix linearly by each row for each column.
func colTraverse() int {
	var ctr int

	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			if matrix[row][col] == 0xFF {
				ctr++
			}
		}
	}

	return ctr
}
