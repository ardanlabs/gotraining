// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program that performs a series of I/O related tasks to
// better understand tracing in Go.
package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/exec"
	"os/signal"
	"runtime/pprof"
	"runtime/trace"
	"time"
)

// Download does network I/O.
func Download() []byte {

	// Download the tar file.
	r, err := http.Get("https://ftp.gnu.org/gnu/binutils/binutils-2.7.tar.gz")
	if err != nil {
		log.Fatal(err)
	}

	// Read in the entire contents of the file.
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	// Return the file.
	return body
}

// Write does disk I/O.
func Write(data []byte) {

	// Perform the disk I/O 50 times.
	for i := 0; i < 50; i++ {

		// Create a new file.
		tmpfile, err := ioutil.TempFile("", "example")
		if err != nil {
			log.Fatal(err)
		}
		defer os.Remove(tmpfile.Name())

		// Write the data to the file.
		_, err = tmpfile.Write(data)
		if err != nil {
			log.Fatal(err)
		}

		// Close the file and flush the final writes.
		tmpfile.Close()
	}
}

// Block waits on a channel for one second.
func Block() {
	<-time.NewTimer(time.Second).C
}

// Hash does CPU-intensive work.
func Hash(data []byte) {
	for i := 0; i < 50; i++ {
		sha256.Sum256(data)
	}
}

// Exec runs an external command.
func Exec() {
	if err := exec.Command("sleep", "1").Run(); err != nil {
		log.Fatal(err)
	}
}

// SlowHandler performs a set of actions that take long to complete.
func SlowHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Started Request")

	// Measure the time this takes.
	start := time.Now()

	// Download the file from the web.
	data := Download()

	// Write the file to disk.
	Write(data)

	// Block for a second.
	Block()

	// Hash the data we received.
	Hash(data)

	// Execute an out of process command.
	Exec()

	// Log how long all this took.
	s := fmt.Sprint("Request took:", time.Since(start))
	fmt.Println(s)
	fmt.Fprintf(w, s)
}

func main() {
	fmt.Println("Listening on 127.0.0.1:12345")

	// Start listening on the local ip.
	l, err := net.Listen("tcp", "127.0.0.1:12345")
	if err != nil {
		log.Fatal(err)
	}

	// Create a file to hold profiling data.
	pf, err := os.Create("cpu.pprof")
	if err != nil {
		log.Fatal(err)
	}
	defer pf.Close()

	// Start gathering the profiling data.
	pprof.StartCPUProfile(pf)
	defer pprof.StopCPUProfile()

	// Create a file to hold tracing data.
	tf, err := os.Create("trace.out")
	if err != nil {
		log.Fatal(err)
	}
	defer tf.Close()

	// Start gathering the tracing data.
	trace.Start(tf)
	defer trace.Stop()

	// Bind the handler to the work route.
	http.HandleFunc("/work", SlowHandler)

	// Set up the OS interrup to stop the service.
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt)
		<-sigChan
		l.Close()
	}()

	// This will server the website until we `/stop` the listener.
	// This will allow us to write the file profile and trace data.
	http.Serve(l, nil)
	fmt.Println("Shutdown")
}
