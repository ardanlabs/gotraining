// Use http.Flusher to stream data (chunked transfer encoding).

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Server struct {
	db *DB
}

func (s *Server) dailyReportHandler(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "no streaming support", http.StatusInternalServerError)
		return
	}

	rows := s.db.DailyReport(time.Now().UTC())
	enc := json.NewEncoder(w)
	var row Row
	for rows.Next(&row) {
		if err := enc.Encode(row); err != nil {
			// Can't set HTTP error
			log.Printf("error: encoding - %s", err)
			return
		}
		flusher.Flush()
	}
}

func main() {
	s := Server{db: &DB{}}
	http.HandleFunc("/daily", s.dailyReportHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("error: %s", err)
	}
}

// ---

type Row struct{}

type DB struct{}

type Rows struct {
	count int
	curr  int
}

func (db *DB) DailyReport(day time.Time) *Rows {
	return &Rows{count: 3}
}

func (r *Rows) Next(v any) bool {
	if r.curr == r.count {
		return false
	}
	r.curr++
	return true
}
