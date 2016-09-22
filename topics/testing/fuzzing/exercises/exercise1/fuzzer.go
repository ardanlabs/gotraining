// +build gofuzz

package fuzzprot

// Fuzz is executed by the go-fuzz tool. Input data modifications
// are provided and used to validate the UnpackUsers function.
func Fuzz(data []byte) int {
	if _, err := UnpackUsers(data); err != nil {
		return 0
	}

	return 1
}
