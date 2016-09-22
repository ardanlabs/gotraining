package main

import (
	"fmt"
	"log"
	"net/http"
)

type App struct{}

func (a App) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello World!")
}

func main() {
	log.Panic(http.ListenAndServe(":3000", App{}))
}
