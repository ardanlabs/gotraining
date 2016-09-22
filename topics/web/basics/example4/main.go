package main

import (
	"fmt"
	"log"
	"net/http"
)

type App struct {
	h http.HandlerFunc
}

func (a App) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	a.h(res, req)
}

func myHander(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello World!")
}

func wrap(h http.HandlerFunc) http.Handler {
	return App{h: h}
}

func main() {
	m := http.NewServeMux()
	m.Handle("/", wrap(myHander))
	log.Panic(http.ListenAndServe(":3000", m))
}
