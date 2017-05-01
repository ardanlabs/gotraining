// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to launch a server that uses TLS.
package main

import (
	"log"
	"net/http"
)

func main() {

	m := http.NewServeMux()

	m.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello, world"))
	})

	log.Print("Listening on localhost:3000")
	//http.ListenAndServe("localhost:3000", m)
	err := http.ListenAndServeTLS("localhost:3000", "cert.pem", "key.pem", m)
	if err != nil {
		log.Println(err)
	}
}
