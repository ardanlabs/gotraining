// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/_CbgE89cyO

// Sample program to learn how to identify memory leaks. This code is
// experimental. It is a treasure hunt at the end of the day.
package main

import (
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
)

func main() {
	// Create a goroutine that leaks memory. Since the map needs
	// to be constantly shuffled around, this becomes very expensive.
	go func() {
		m := make(map[int]int)

		for i := 0; ; i++ {
			m[i] = i
		}
	}()

	// Start a listener for the pprof support.
	go func() {
		http.ListenAndServe(":6060", nil)
	}()

	// On a <ctrl> C shutdown the program.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig)
	<-sig
}
