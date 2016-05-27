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

// Fuzz is morphing the input from the corpus.
func Fuzz(data []byte) int {
	r, _ := http.NewRequest("POST", "/process", bytes.NewBuffer(data))
	w := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(w, r)

	if w.Code != 200 {

		// This was not successful. It we panic it will be recorded.
		panic(w.Body.String())
	}

	// The data caused no issues and is not interesting.
	return 1
}
