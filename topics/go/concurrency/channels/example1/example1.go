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
	// fanOut()
	// selectDrop()
}

type data struct {
	UserID string
}

// basicSendRecv shows the basics of a send and receive.
func basicSendRecv() {
	fmt.Println("** basicSendRecv")

	ch := make(chan data)
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch <- data{"123"}
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

	_, wd := <-ch
	fmt.Println("g1 : received :", wd)

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}

// signalAck shows how to signal an event and wait for an
// acknowledgment it is done.
func signalAck() {
	fmt.Println("** signalAck")

	ch := make(chan data)
	go func() {
		v := <-ch
		fmt.Println("g2 : received :", v)
		time.Sleep(100 * time.Millisecond)
		ch <- data{"123"}
		fmt.Println("g2 : send ack")
	}()

	ch <- data{"123"}
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

	ch := make(chan data, 5)
	for i := 0; i < 5; i++ {
		ch <- data{fmt.Sprintf("%d", i)}
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

	ch := make(chan data, 1)
	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- data{"123"}
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

// fanOut shows how to use the fan out pattern to get work
// done concurrently.
func fanOut() {
	fmt.Println("** fanOut")

	const grs = 20
	ch := make(chan data, grs)

	for g := 0; g < grs; g++ {
		go func(g int) {
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			ch <- data{fmt.Sprintf("%d", g)}
		}(g)
	}

	for g := 0; g < grs; g++ {
		v := <-ch
		fmt.Println(v)
	}

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}

// selectDrop shows how to use the select to walk away from a channel
// operation if it will immediately block.
func selectDrop() {
	fmt.Println("** selectDrop")

	ch := make(chan data, 5)
	go func() {
		for v := range ch {
			fmt.Println("g2 : received :", v)
		}
	}()

	for i := 0; i < 20; i++ {
		select {
		case ch <- data{fmt.Sprintf("%d", i)}:
			fmt.Println("g1 : send ack")
		default:
			fmt.Println("g1 : drop")
		}
	}
	close(ch)

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}
