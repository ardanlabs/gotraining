// +build gofuzz

// go-fuzz-build github.com/ardanlabs/gotraining/topics/fuzzing/example1
// go-fuzz -bin=./api-fuzz.zip -workdir=workdir/corpus

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
		return 0
	}

	return 1
}
