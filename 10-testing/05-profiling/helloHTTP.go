// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/fcU9jQX2Qz

// Sample program to show how profiling works. This is the base
// code. Use the readme for changes needed to be made to the code.
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
)

type (
	// hello holds a message.
	hello struct {
		Message string
	}
)

// main is the entry point for the application.
func main() {
	http.HandleFunc("/english", helloEnglish)
	http.HandleFunc("/chinese", helloChinese)

	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan
}

// helloEnglish sends a greeting in English.
func helloEnglish(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(hello{"Hello World"}); err != nil {
		log.Println("Error encoding JSON", err)
		return
	}
}

// helloChinese sends a greeting in Chinese.
func helloChinese(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(hello{"你好世界"}); err != nil {
		log.Println("Error encoding JSON", err)
		return
	}
}
