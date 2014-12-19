// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/c44Q5OiX5z

// go tool pprof http://localhost:9999/debug/pprof/heap
// go tool pprof http://localhost:9999/debug/pprof/profile
// while true; do curl http://localhost:9999/english; done

// Sample program to show off Go and check programming environment.
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/davecheney/profile"
)

type (
	// hello holds a message.
	hello struct {
		Message string
	}
)

// main is the entry point for the application.
func main() {
	cfg := profile.Config{
		MemProfile:     true,
		CPUProfile:     true,
		ProfilePath:    ".",  // store profiles in current directory
		NoShutdownHook: true, // do not hook SIGINT
	}

	// p.Stop() must be called before the program exits to
	// ensure profiling information is written to disk.
	p := profile.Start(&cfg)
	defer p.Stop()

	http.HandleFunc("/english", helloEnglish)
	http.HandleFunc("/chinese", helloChinese)

	go func() {
		http.ListenAndServe("localhost:9999", nil)
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan
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
