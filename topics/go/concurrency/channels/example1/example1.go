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
	basicSendRecv()
	signalClose()
	signalAck()
	closeRange()
	selectRecv()
	selectSend()
	selectDrop()
}

// basicSendRecv shows the basics of a send and receive.
func basicSendRecv() {
	fmt.Println("** basicSendRecv")

	ch := make(chan string)
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch <- "done"
		fmt.Println("g2 : send ack")
	}()

	v := <-ch
	fmt.Println("g1 : received :", v)

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}

// signalClose shows how to close a channel to signal an event.
func signalClose() {
	fmt.Println("** signalClose")

	ch := make(chan struct{})
	go func() {
		time.Sleep(100 * time.Millisecond)
		close(ch)
		fmt.Println("g2 : close ack")
	}()

	_, ok := <-ch
	fmt.Println("g1 : received :", ok)

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}

// signalAck shows how to signal an event and wait for an
// acknowledgment it is done.
func signalAck() {
	fmt.Println("** signalAck")

	ch := make(chan string)
	go func() {
		v := <-ch
		fmt.Println("g2 : received :", v)
		time.Sleep(100 * time.Millisecond)
		ch <- "done"
		fmt.Println("g2 : send ack")
	}()

	ch <- "work"
	fmt.Println("g1 : send ack")
	v := <-ch
	fmt.Println("g1 : received :", v)

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}

// closeRange shows how to use range to receive value and
// using close to terminate the loop.
func closeRange() {
	fmt.Println("** closeRange")

	ch := make(chan int, 5)
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)

	for v := range ch {
		fmt.Println("g1 : received :", v)
	}

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}

// selectRecv shows how to use the select statement to wait for a
// specified amount of time to receive a value.
func selectRecv() {
	fmt.Println("** selectRecv")

	ch := make(chan string, 1)
	go func() {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		ch <- "done"
		fmt.Println("g2 : send ack")
	}()

	select {
	case v := <-ch:
		fmt.Println("g1 : received :", v)
	case t := <-time.After(100 * time.Millisecond):
		fmt.Println("g1 : timed out :", t)
	}

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}

// selectSend shows how to use the select statement to attempt a
// send on a channel for a specified amount of time.
func selectSend() {
	fmt.Println("** selectSend")

	ch := make(chan string)
	go func() {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		v, ok := <-ch
		fmt.Println("g2 : received :", v, ok)
		if !ok {
			fmt.Println("g2 : cancelled")
			return
		}
	}()

	select {
	case ch <- "work":
		fmt.Println("g1 : send ack")
	case t := <-time.After(100 * time.Millisecond):
		fmt.Println("g1 : timed out :", t)
		close(ch)
	}

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}

// selectDrop shows how to use the select to walk away from a channel
// operation if it will immediately block.
func selectDrop() {
	fmt.Println("** selectDrop")

	ch := make(chan int, 5)
	go func() {
		for v := range ch {
			fmt.Println("g2 : received :", v)
		}
	}()

	for v := 0; v < 20; v++ {
		select {
		case ch <- v:
			fmt.Println("g1 : send ack")
		default:
			fmt.Println("g1 : drop")
		}
	}
	close(ch)

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}
