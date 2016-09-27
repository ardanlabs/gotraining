package main

import (
	"log"
	"net/http"
	"path"
	"runtime"
)

func main() {
	m := http.NewServeMux()
	m.Handle("/", http.FileServer(http.Dir(staticDir())))
	log.Fatal(http.ListenAndServe(":3000", m))
}

// staticDir builds a full path to the 'static' directory
// that is relative to this file.
func staticDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Join(path.Dir(filename), "static")
}
