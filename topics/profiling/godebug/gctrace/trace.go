// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to see what a memory leak looks like.
package main

import (
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

	// On a <ctrl> C shutdown the program.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig)
	<-sig
}
