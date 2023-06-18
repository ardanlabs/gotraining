package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"time"
)

type Event struct {
	Time   time.Time `json:"time"`
	User   string    `json:"user"`
	Action string    `json:"action"`
	URI    string    `json:"uri"`
}

func main() {
	db := NewDB()
	rows := db.Query()

	var net bytes.Buffer

	// Server
	enc := json.NewEncoder(&net)
	for rows.Next() {
		evt := rows.Event()
		enc.Encode(evt)
	}
	if err := rows.Err(); err != nil {
		log.Fatalf("error: %s", err)
	}

	// Client
	dec := json.NewDecoder(&net)
	var e Event
	for {
		err := dec.Decode(&e)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error: %s", err)
		}
		fmt.Println(e)
	}
}
