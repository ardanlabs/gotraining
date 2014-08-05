// Sample program to show off Go and check programming environment.
package main

import (
	"fmt"
	"net/http"
)

// main is the entry point for the application.
func main() {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/chinese", helloWorldChinese)
	http.ListenAndServe("localhost:9999", nil)
}

// helloWorld handles the index route.
func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<H1>Hello World</H1>")
}

// helloWorldChinese handles the chinese route.
func helloWorldChinese(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<H1>你好世界</H1>")
}
