package main

import (
	"log"
	"net/http"
	"path"
	"runtime"
)

func App() http.Handler {
	m := http.NewServeMux()
	m.Handle("/", http.FileServer(http.Dir(staticDir())))
	return m
}

func main() {
	log.Fatal(http.ListenAndServe(":3000", App()))
}

// staticDir builds a full path to the 'static' directory
// that is relative to this file.
func staticDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Join(path.Dir(filename), "static")
}
