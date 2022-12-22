package main

import (
	"database/sql"
	_ "embed"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var (
	//go:embed query.sql
	querySQL string
)

type DB struct {
	conn *sql.DB
}

func NewDB(dsn string) (*DB, error) {
	conn, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}
	db := DB{conn}
	return &db, nil
}

func (db *DB) Close() error {
	return db.conn.Close()
}

type Metric struct {
	Time  time.Time `json:"time"`
	Name  string    `json:"name"`
	Value float64   `json:"value"`
}

func (db *DB) ByTime(start, end time.Time) ([]Metric, error) {
	rows, err := db.conn.Query(querySQL, sql.Named("start", start), sql.Named("end", end))
	if err != nil {
		return nil, err
	}

	var metrics []Metric
	for rows.Next() {
		var m Metric
		if err := rows.Scan(&m.Time, &m.Name, &m.Value); err != nil {
			return nil, err
		}
		metrics = append(metrics, m)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return metrics, nil
}
