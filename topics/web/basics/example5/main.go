package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	m := http.NewServeMux()
	m.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		log.Println(req.URL.Path)
		time.Sleep(1 * time.Second)
		res.WriteHeader(http.StatusTeapot)
	})
	go http.ListenAndServe(":3000", m)

	start := time.Now()
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			res, err := http.Get(fmt.Sprintf("http://localhost:3000/%d", n))
			if err != nil {
				log.Fatal(err)
			}
			if res.StatusCode != http.StatusTeapot {
				log.Fatal("Oops!")
			}
		}(i)
	}
	wg.Wait()
	fmt.Printf("\nduration: %s\n", time.Now().Sub(start))
}
