// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program that implements a simple web service that will allow us to
// explore how to use the schedtrace. This code is leaking a gorutine.
package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var leak bool

func main() {
	http.HandleFunc("/sendjson", sendJSON)

	// Leak goroutines if we have any argument.
	if len(os.Args) == 2 {
		leak = true
	}

	log.Printf("listener : Started : Listening on: http://localhost:4000 : Leak[%v]\n", leak)
	http.ListenAndServe(":4000", nil)
}

// sendJSON returns a simple JSON document.
func sendJSON(rw http.ResponseWriter, r *http.Request) {

	// Leak a goroutine every so often.
	if leak {
		if rand.Intn(100) == 5 {
			go func() {
				for {
					time.Sleep(time.Millisecond * 10)
				}
			}()
		}
	}

	u := struct {
		Name  string
		Email string
	}{
		Name:  "bill",
		Email: "bill@ardanlabs.com",
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	json.NewEncoder(rw).Encode(&u)
}
