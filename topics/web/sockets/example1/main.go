// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program on how to use the Google web socket package
// to bind HTTP requests.
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/gorilla/pat"
	"golang.org/x/net/websocket"
)

// Message represents the request we will receive on
// the web socket connection.
type Message struct {
	Original  string    `json:"original"`
	Formatted string    `json:"formatted"`
	Received  time.Time `json:"received"`
}

// socketHandler is created for each connect that is
// established on the server.
func socketHandler(ws *websocket.Conn) {

	log.Println("Connection established")
	defer log.Println("Connection dropped")

	// Create a buffer for use processing messages.
	msg := make([]byte, 512)

	// Maintain a read loop until the connection is
	// broken or lost.
	for {

		// Read a frame of data from the WebSocket connection.
		// If msg is not large enough for the frame data, it fills
		// the msg and next Read will read the rest of the frame data.
		n, err := ws.Read(msg)
		if err != nil {
			log.Println(err)
			break
		}

		// If we received any data process it.
		if n > 0 {

			// Convert the bytes we received to a string.
			data := string(msg[:n])

			// Create a message and store the data.
			msg := Message{
				Original:  data,
				Formatted: strings.ToUpper(data),
				Received:  time.Now(),
			}

			// Encode the message to JSON and send it back.
			if err := json.NewEncoder(ws).Encode(&msg); err != nil {
				log.Println(err)
				break
			}
		}
	}
}

// App returns a handler for handling requets with JWT.
func App() http.Handler {

	// Create a new Pat router.
	r := pat.New()

	// Bind a GET call for the `/socket` route. This will establish
	// a web socket connection.
	r.Get("/socket", func(res http.ResponseWriter, req *http.Request) {
		websocket.Handler(socketHandler).ServeHTTP(res, req)
	})

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
