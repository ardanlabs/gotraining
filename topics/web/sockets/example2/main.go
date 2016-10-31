package main

import (
	"log"
	"net/http"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/gorilla/pat"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type Message struct {
	Original  string    `json:"original"`
	Formatted string    `json:"formatted"`
	Received  time.Time `json:"received"`
}

func SocketHandler(res http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(res, req, nil)
	if err != nil {
		log.Println(err)
		return
	}
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}

		m := string(msg)
		message := Message{
			Original:  m,
			Formatted: strings.ToUpper(m),
			Received:  time.Now(),
		}
		err = conn.WriteJSON(message)
		if err != nil {
			log.Println(err)
			break
		}
	}
}

func App() http.Handler {
	r := pat.New()

	r.Get("/socket", SocketHandler)
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
