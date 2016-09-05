// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to see what a trace will look like for basic
// channel latencies.
package main

import (
	"bytes"
	"fmt"
	"io"
	"sync"
	"testing"
	"time"
)

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

func recv(ch chan int) {
	var total int

	for v := range ch {
		total = total + v
	}
}

func input() io.Reader {
	base := []byte{0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01}
	var data []byte
	for i := 0; i < 1000000; i++ {
		data = append(data, base...)
	}

	return bytes.NewBuffer(data)
}

// TestLatency provides a test to profile and trace channel latencies.
func TestLatency(t *testing.T) {
	for i := 0; i < 40; i++ {
		var wg sync.WaitGroup
		ch := make(chan int, i)

		wg.Add(1)
		go func() {
			recv(ch)
			wg.Done()
		}()

		st := time.Now()

		send(input(), ch)
		close(ch)
		wg.Wait()

		fmt.Println(i, time.Since(st))
	}
}
