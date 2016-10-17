package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/markbates/going/wait"
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
	wait.Wait(100, func(i int) {
		res, err := http.Get(fmt.Sprintf("http://localhost:3000/%d", i))
		if err != nil {
			log.Fatal(err)
		}
		if res.StatusCode != http.StatusTeapot {
			log.Fatal("Oops!")
		}
	})
	fmt.Printf("\nduration: %s\n", time.Now().Sub(start))
}
