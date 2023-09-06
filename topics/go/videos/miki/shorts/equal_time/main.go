package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func main() {
	t1 := time.Now()
	data, err := json.Marshal(t1)
	if err != nil {
		log.Fatalf("error: can't marshal - %s", err)
	}

	var t2 time.Time
	if err := json.Unmarshal(data, &t2); err != nil {
		log.Fatalf("error: can't unmarshal - %s", err)
	}

	fmt.Println(t1 == t2)     // false
	fmt.Println(t1.Equal(t2)) // true
}
