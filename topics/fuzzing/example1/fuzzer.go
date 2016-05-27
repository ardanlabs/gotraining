// +build gofuzz

package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
)

func init() {
	Routes()
}

// Fuzz is executed by the go-fuzz tool. Input data modifications
// are provided and used to validate API call.
func Fuzz(data []byte) int {
	r, _ := http.NewRequest("POST", "/process", bytes.NewBuffer(data))
	w := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(w, r)

	// Data which has been corrupted but "looks valid" is more likely to be
	// able to get into other valid parts of your program without being
	// discarded as junk.

	if w.Code != 200 {

		// Report the data that produced this error as not interesting.
		return 0
	}

	// Report the data that did not cause an error as interesting.
	return 1
}
