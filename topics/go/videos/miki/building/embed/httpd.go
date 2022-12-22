package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	//go:embed static
	staticDir embed.FS
)

type Server struct {
	db *DB
}

func getTime(name string, r *http.Request) (time.Time, error) {
	s := r.URL.Query().Get(name)
	if s == "" {
		return time.Time{}, fmt.Errorf("missing %q", name)
	}

	return time.Parse("2006-01-02T15:04:05", s)
}

func (s *Server) getMetrics(w http.ResponseWriter, r *http.Request) {
	start, err := getTime("start", r)
	if err != nil {
		http.Error(w, "bad start time", http.StatusBadRequest)
		return
	}

	end, err := getTime("end", r)
	if err != nil {
		http.Error(w, "bad end time", http.StatusBadRequest)
		return
	}

	metrics, err := s.db.ByTime(start, end)
	if err != nil {
		log.Printf("can't query for %v-%v: %s", start, end, err)
		http.Error(w, "can't query", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(metrics); err != nil {
		log.Printf("error encoding: %s", err)
	}
}

func main() {
	dsn := os.Getenv("DSN")
	if dsn == "" {
		dsn = "metrics.db"
	}
	log.Printf("dsn: %q", dsn)

	db, err := NewDB(dsn)
	if err != nil {
		log.Fatal(err)
	}

	s := Server{db}
	http.HandleFunc("/metrics", s.getMetrics)
	http.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/static/doc.html", http.StatusTemporaryRedirect)
	})
	http.Handle("/static/", http.FileServer(http.FS(staticDir)))

	port := 8080
	if s := os.Getenv("PORT"); s != "" {
		var err error
		port, err = strconv.Atoi(s)
		if err != nil {
			log.Fatalf("error: bad port: %q - %s", s, err)
		}
	}

	addr := fmt.Sprintf(":%d", port)
	log.Printf("server starting on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
