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
	// sendRecv()
	// signalClose()
	// closeRange()
	// fanOut()
	// fanOutSem()
	// selectCancel()
	// selectDrop()
}

type data struct {
	UserID string
}

// sendRecv shows the basics of a send and receive.
func sendRecv() {
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

// closeRange shows how to use range to receive values and
// how using close terminates the loop.
func closeRange() {
	fmt.Println("** closeRange")

	ch := make(chan data)

	const grs = 2
	for g := 0; g < grs; g++ {
		go func(id int) {
			for v := range ch {
				fmt.Println("g2 : received :", v, ": on :", id)
			}
			fmt.Println("g2 : received close : on :", id)
		}(g)
	}

	const reqs = 10
	for i := 0; i < reqs; i++ {
		ch <- data{fmt.Sprintf("%d", i)}
	}

	close(ch)
	fmt.Println("g1 : close ack")

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}

// fanOut shows how to use the fan out pattern to get work
// done concurrently.
func fanOut() {
	fmt.Println("** fanOut")

	grs := 20
	ch := make(chan data, grs)

	for g := 0; g < grs; g++ {
		go func(g int) {
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			ch <- data{fmt.Sprintf("%d", g)}
		}(g)
	}

	for grs > 0 {
		d := <-ch
		fmt.Println(d)
		grs--
	}

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}

// fanOutSem shows how to use the fan out pattern to get work
// done concurrently but limiting the number of active routines.
func fanOutSem() {
	fmt.Println("** fanOutSem")

	grs := 20
	ch := make(chan data, grs)

	const cap = 5
	sem := make(chan bool, cap)

	for g := 0; g < grs; g++ {
		go func(g int) {
			sem <- true
			{
				time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
				ch <- data{fmt.Sprintf("%d", g)}
			}
			<-sem
		}(g)
	}

	for grs > 0 {
		d := <-ch
		fmt.Println(d, len(sem))
		grs--
	}

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}

// selectCancel shows how to use the select statement to wait for a
// specified amount of time to receive a value.
func selectCancel() {
	fmt.Println("** selectRecv")

	ch := make(chan data, 1)

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- data{"123"}
		fmt.Println("g2 : send ack")
	}()

	tc := time.After(100 * time.Millisecond)

	select {
	case v := <-ch:
		fmt.Println("g1 : received :", v)
	case t := <-tc:
		fmt.Println("g1 : timed out :", t)
	}

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}

// selectDrop shows how to use the select to walk away from a channel
// operation if it will immediately block.
func selectDrop() {
	fmt.Println("** selectDrop")

	const cap = 5
	ch := make(chan data, cap)

	go func() {
		for v := range ch {
			fmt.Println("g2 : received :", v)
		}
	}()

	const reqs = 20
	for i := 0; i < reqs; i++ {
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
