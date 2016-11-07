// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program on how to use the Gorilla web socket package
// to bind HTTP requests.
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

// upgrader provides configuration for the websocket.
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// Message represents the request we will receive on
// the web socket connection.
type Message struct {
	Original  string    `json:"original"`
	Formatted string    `json:"formatted"`
	Received  time.Time `json:"received"`
}

// socketHandler is created for each connect that is
// established on the server.
func socketHandler(res http.ResponseWriter, req *http.Request) {

	log.Println("Connection established")
	defer log.Println("Connection dropped")

	// Upgrade the HTTP server connection to the WebSocket protocol.
	conn, err := upgrader.Upgrade(res, req, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// Maintain a read loop until the connection is
	// broken or lost.
	for {

		// Read a message from the connection buffer.
		_, m, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}

		// Convert the bytes we received to a string.
		data := string(m)

		// Create a message and store the data.
		msg := Message{
			Original:  data,
			Formatted: strings.ToUpper(data),
			Received:  time.Now(),
		}

		// Encode the message to JSON and send it back.
		if err := conn.WriteJSON(msg); err != nil {
			log.Println(err)
			break
		}
	}
}

// App returns a handler for handling requets with JWT.
func App() http.Handler {

	// Create a new Pat router.
	r := pat.New()

	// Bind a GET call for the `/socket` route. This will establish
	// a web socket connection.
	r.Get("/socket", socketHandler)

	// Bind the route for serving static files using the
	// default FileServer. This will load the home page.
	// r.Handle("/", http.FileServer(http.Dir(staticDir())))
	r.Get("/", func(res http.ResponseWriter, req *http.Request) {
		http.FileServer(http.Dir(currentDirectory())).ServeHTTP(res, req)
	})

	return r
}

// Returns the current directory we are running in.
func currentDirectory() string {

	// Locate the current directory for the site.
	_, fn, _, _ := runtime.Caller(1)
	return path.Join(path.Dir(fn), "static")
}

func main() {

	// Start the http server to handle the request for
	// both versions of the API.
	log.Fatal(http.ListenAndServe(":3000", App()))
}
