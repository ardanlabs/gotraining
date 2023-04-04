// encoding/json will marshal []byte in base64

package main

import (
	"encoding/json"
	"os"
)

func main() {
	msg := map[string]interface{}{
		"label": 7,
		"shape": []int{8, 8},
		"image": []byte{
			0, 0, 3, 12, 16, 10, 0, 0,
			0, 2, 14, 12, 12, 12, 0, 0,
			0, 5, 10, 0, 10, 11, 0, 0,
			0, 0, 0, 1, 14, 9, 2, 0,
			0, 0, 8, 16, 16, 16, 10, 0,
			0, 0, 6, 16, 13, 7, 0, 0,
			0, 0, 0, 16, 5, 0, 0, 0,
			0, 0, 5, 13, 0, 0, 0, 0,
		},
	}
	json.NewEncoder(os.Stdout).Encode(msg)
}
