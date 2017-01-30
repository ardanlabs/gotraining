// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This sample program demonstrates the basic channel mechanics
// for goroutine signaling.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// basicSendRecv()
	// signalClose()
	// signalAck()
	// closeRange()
	// selectRecv()
	// selectSend()
	// selectDrop()
}

// basicSendRecv shows the basics of a send and receive.
func basicSendRecv() {
	ch := make(chan string)
	go func() {
		ch <- "hello"
	}()

	fmt.Println(<-ch)
}

// signalClose shows how to close a channel to signal an event.
func signalClose() {
	ch := make(chan struct{})
	go func() {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("signal event")
		close(ch)
	}()

	<-ch
	fmt.Println("event received")
}

// signalAck shows how to signal an event and wait for an
// acknowledgment it is done.
func signalAck() {
	ch := make(chan string)
	go func() {
		fmt.Println(<-ch)
		ch <- "ok done"
	}()

	ch <- "do this"
	fmt.Println(<-ch)
}

// closeRange shows how to use range to receive value and
// using close to terminate the loop.
func closeRange() {
	ch := make(chan int, 5)
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)

	for v := range ch {
		fmt.Println(v)
	}
}

// selectRecv shows how to use the select statement to wait for a
// specified amount of time to receive a value.
func selectRecv() {
	ch := make(chan string)
	go func() {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		ch <- "work"
	}()

	select {
	case v := <-ch:
		fmt.Println(v)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("timed out")
	}
}

// selectRecv shows how to use the select statement to attempt a
// send on a channel for a specified amount of time.
func selectSend() {
	ch := make(chan string)
	go func() {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		fmt.Println(<-ch)
	}()

	select {
	case ch <- "work":
		fmt.Println("send work")
	case <-time.After(100 * time.Millisecond):
		fmt.Println("timed out")
	}
}

// selectDrop shows how to use the select to walk away from a channel
// operation if it will immediately block.
func selectDrop() {
	ch := make(chan int, 5)
	go func() {
		for v := range ch {
			fmt.Println("recv", v)
		}
	}()

	for i := 0; i < 20; i++ {
		select {
		case ch <- i:
		default:
			fmt.Println("drop", i)
		}
	}
	close(ch)
}
