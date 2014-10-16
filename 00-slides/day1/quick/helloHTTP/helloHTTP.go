// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/c44Q5OiX5z

// Sample program to show off Go and check programming environment.
package main

import (
	"encoding/json"
	"log"
	"net/http"
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
	http.ListenAndServe("localhost:9999", nil)
}

// helloEnglish sends a greeting in English.
func helloEnglish(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(hello{"Hello World"})
	if err != nil {
		log.Println("Error encoding JSON", err)
		return
	}
	log.Println("Sent English")
}

// helloChinese sends a greeting in Chinese.
func helloChinese(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(hello{"你好世界"})
	if err != nil {
		log.Println("Error encoding JSON", err)
		return
	}
	log.Println("Sent Chinese")
}
