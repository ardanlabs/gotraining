package main

import (
	"encoding/json"
	"log"
	"net/http"
	"path"
	"runtime"
	"strings"
	"time"

	"golang.org/x/net/websocket"

	"github.com/gorilla/pat"
)

type Message struct {
	Original  string    `json:"original"`
	Formatted string    `json:"formatted"`
	Received  time.Time `json:"received"`
}

func SocketHandler(ws *websocket.Conn) {
	enc := json.NewEncoder(ws)
	for {
		msg := make([]byte, 512)
		read, err := ws.Read(msg)
		if err != nil {
			log.Println(err)
			break
		}

		if read > 0 {
			m := string(msg[:read])
			message := Message{
				Original:  m,
				Formatted: strings.ToUpper(m),
				Received:  time.Now(),
			}
			enc.Encode(message)
		}
	}
}

func App() http.Handler {
	r := pat.New()

	r.Get("/socket", func(res http.ResponseWriter, req *http.Request) {
		websocket.Handler(SocketHandler).ServeHTTP(res, req)
	})
	// Bind the route for serving static files using the
	// default FileServer. This will load the home page.
	// r.Handle("/", http.FileServer(http.Dir(staticDir())))
	r.Get("/", func(res http.ResponseWriter, req *http.Request) {
		http.FileServer(http.Dir(staticDir())).ServeHTTP(res, req)
	})

	return r
}

// staticDir builds a full path to the 'static' directory
// that is relative to this file.
func staticDir() string {

	// Locate from the runtime the location of
	// the apps static files.
	_, filename, _, _ := runtime.Caller(1)

	// Return a path to the static folder.
	return path.Join(path.Dir(filename), "static")
}

func main() {
	log.Fatal(http.ListenAndServe(":3000", App()))
}
