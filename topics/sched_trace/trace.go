// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/9--g7HC0J5

// https://github.com/wg/wrk
// wrk -t8 -c500 -d5s http://localhost:4000/sendjson
// export GODEBUG=schedtrace=1000

// Sample program that implements a simple web service that will allow us to
// explore how to use the schedtrace.
// THIS CODE IS SHOWING MECHANICS. PLEASE DO NOT COPY/PASTE FOR PRODUCTION.
package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// main is the entry point for the application.
func main() {
	http.HandleFunc("/sendjson", sendJSON)

	log.Println("listener : Started : Listening on: http://localhost:4000")
	http.ListenAndServe(":4000", nil)
}

// sendJSON returns a simple JSON document.
func sendJSON(rw http.ResponseWriter, r *http.Request) {
	name := "bill"

	// Step 1:
	// Let's have each request do some pretend work and never allow the
	// goroutine to move out of a runnable state. DO NOT DO THIS AT HOME.
	//
	// Note we never see more than 500 goroutines. This is because we never
	// have more than 500 socket connections. One connection per goroutine.
	//
	// for i := 0; i < 1000; i++ {
	// 	runtime.Gosched()
	// }

	// Step 2:
	// Leak a goroutine on every request. Have the goroutine do a little
	// work in between sleeping on pretend IO.
	//
	// Note everything seems ok but then we see spikes in the number of
	// goroutines. This is because eventually a large number of these leaked
	// goroutines wake up.
	//
	// go func() {
	// 	for {
	// 		<-time.After(time.Second)
	// 		v1 := strings.Split("name:bill", ":")
	// 		name = v1[1]
	// 	}
	// }()

	u := struct {
		Name  string
		Email string
	}{
		Name:  name,
		Email: "bill@ardanlabs.com",
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	json.NewEncoder(rw).Encode(&u)
}
