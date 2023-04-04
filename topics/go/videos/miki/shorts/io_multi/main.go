// Use io.MultiWriter to write to several destinations.

package main

import (
	"io"
	"log"
	"os"
)

func main() {
	flag := os.O_APPEND | os.O_CREATE | os.O_WRONLY
	file, err := os.OpenFile("app.log", flag, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	w := io.MultiWriter(file, os.Stdout)
	logger := log.New(w, "app: ", log.LstdFlags)
	logger.Printf("INFO: %q logged in", "eliot")
	// app: 2023/04/02 12:16:36 INFO: "eliot" logged in
	// both in stdout and app.log
}
