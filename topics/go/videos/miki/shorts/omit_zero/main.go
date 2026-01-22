package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type LogQuery struct {
	StartTime time.Time `json:"start_time,omitempty"`
	EndTime   time.Time `json:"end_time,omitempty"`
	Level     string    `json:"level,omitempty"`
}

func main() {
	q := LogQuery{
		Level: "INFO",
	}

	data, err := json.Marshal(q)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: marshal - %v\n", err)
		os.Exit(1)
	}

	os.Stdout.Write(data)
}
