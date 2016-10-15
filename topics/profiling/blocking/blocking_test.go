// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to see what a trace will look like for basic
// channel latencies.
package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/plotutil"
	"github.com/gonum/plot/vg"
)

// data represents a set of bytes to process.
var data []byte

// init creates a data for processing.
func init() {
	f, err := os.Open("data.bytes")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data, err = ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Bytes", len(data))
}

// TestLatency runs a single stream so we can look at
// blocking profiles for different buffer sizes.
func TestLatency(t *testing.T) {
	bufSize := 0

	fmt.Println("BufSize:", bufSize)
	stream(bufSize)
}

// TestLatencies provides a test to profile and trace channel latencies
// with a little data science sprinkled in.
func TestLatencies(t *testing.T) {
	var bufSize int
	var count int
	var first time.Duration

	pts := make(plotter.XYs, 20)

	for {

		// Perform a stream with specified buffer size.
		since := stream(bufSize)

		// Calculate how long this took and the percent
		// of different from the unbuffered channel.
		if bufSize == 0 {
			first = since
		}
		dec := ((float64(first) - float64(since)) / float64(first)) * 100

		// Display the results.
		fmt.Printf("BufSize: %d\t%v\t%.2f%%\n", bufSize, since, dec)

		// Prepare the results for plotting.
		pts[count].X = float64(bufSize)
		pts[count].Y = dec
		count++

		// Want to look at a single buffer increment.
		if bufSize < 10 {
			bufSize++
			continue
		}

		// Increment by 10 moving forward.
		if bufSize == 100 {
			break
		}
		bufSize = bufSize + 10
	}

	// Make the plot of latencies.
	if err := makePlot(pts); err != nil {
		log.Fatal(err)
	}
}

// stream performs the moving of the data stream from
// one goroutine to the other.
func stream(bufSize int) time.Duration {

	// Create WaitGroup and channels.
	var wg sync.WaitGroup
	ch := make(chan int, bufSize)

	// Capture the reader for the input data.
	data := input()

	// Create the receiver goroutine.
	wg.Add(1)
	go func() {
		recv(ch)
		wg.Done()
	}()

	// Start the clock.
	st := time.Now()

	// Send all the data to the receiving goroutine.
	send(data, ch)

	// Close the channel and wait for the receiving
	// goroutine to terminate.
	close(ch)
	wg.Wait()

	// Calculate how long the send took.
	return time.Since(st)
}

// input returns a reader that can be used to read a stream
// of bytes.
func input() io.Reader {
	return bytes.NewBuffer(data)
}

// recv waits for bytes and adds them up.
func recv(ch chan int) {
	var total int

	for v := range ch {
		total = total + v
	}
}

// send reads the bytes and sends them through the channel.
func send(r io.Reader, ch chan int) {
	buf := make([]byte, 1)

	for {
		n, err := r.Read(buf)
		if n == 0 || err != nil {
			break
		}

		ch <- int(buf[0])
	}
}

// makePlot creates and saves a plot of the overall latencies
// differenced from the unbuffered channel.
func makePlot(xys plotter.XYs) error {

	// Create a new plot.
	p, err := plot.New()
	if err != nil {
		return err
	}

	// Label the new plot.
	p.Title.Text = "Latencies (differenced from the unbuffered channel)"
	p.X.Label.Text = "Buffer Length"
	p.Y.Label.Text = "Latency"

	// Add the prepared points to the plot.
	if err = plotutil.AddLinePoints(p, "Latencies", xys); err != nil {
		return err
	}

	// Save the plot to a PNG file.
	if err := p.Save(10*vg.Inch, 5*vg.Inch, "latencies.png"); err != nil {
		return err
	}

	return nil
}
