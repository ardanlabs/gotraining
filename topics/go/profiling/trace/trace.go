// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program that performs a series of I/O related tasks to
// better understand tracing in Go.
package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime/trace"
	"sync"
)

// LoadWrite reads a file from the network into memory and then
// writes it to disk.
func LoadWrite() {

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

	// Create a new file.
	f, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}

	// Defer the close and removal of the file.
	defer os.Remove(f.Name())
	defer f.Close()

	// Write the data to the file.
	_, err = f.Write(body)
	if err != nil {
		log.Fatal(err)
	}
}

// StreamWrite streams a file from the network, writing it to disk.
func StreamWrite() {

	// Download the tar file.
	r, err := http.Get("https://ftp.gnu.org/gnu/binutils/binutils-2.7.tar.gz")
	if err != nil {
		log.Fatal(err)
	}

	// Create a new file.
	f, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}

	// Defer the close and removal of the file.
	defer os.Remove(f.Name())
	defer f.Close()

	// Stream the file to disk.
	if _, err = io.Copy(f, r.Body); err != nil {
		if err != io.EOF {
			log.Fatal(err)
		}
	}
}

// Sort implements quick sort.
func Sort(values []int, l int, r int, calls int) {
	if l >= r {
		return
	}

	pivot := values[l]
	i := l + 1

	for j := l; j <= r; j++ {
		if pivot > values[j] {
			values[i], values[j] = values[j], values[i]
			i++
		}
	}

	values[l], values[i-1] = values[i-1], pivot

	if calls < 0 {
		calls++
		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			Sort(values, l, i-2, calls)
			wg.Done()
		}()
		go func() {
			Sort(values, i, r, calls)
			wg.Done()
		}()
		wg.Wait()
	} else {
		Sort(values, l, i-2, calls)
		Sort(values, i, r, calls)
	}
}

func main() {

	// Create a file to hold tracing data.
	tf, err := os.Create("trace.out")
	if err != nil {
		log.Fatal(err)
	}
	defer tf.Close()

	// Start gathering the tracing data.
	trace.Start(tf)
	defer trace.Stop()

	// LoadWrite()
	// StreamWrite()

	// rand.Seed(time.Now().UnixNano())
	// numbers := make([]int, 100000)
	// for i := range numbers {
	// 	numbers[i] = rand.Intn(10000000)
	// }
	// Sort(numbers, 0, len(numbers)-1, 0)
}
