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
	// waitForTask()
	// waitForResult()
	// waitForFinished()
	// pooling()
	// fanOut()
	// fanOutSem()
	// drop()
	// cancellation()
}

// waitForTask: Think about being a manager and hiring a new employee. In
// this scenario, you want your new employee to perform a task but they need
// to wait until you are ready. This is because you need to hand them a piece
// of paper before they start.
func waitForTask() {
	ch := make(chan string)

	go func() {
		p := <-ch
		fmt.Println("employee : recv'd signal :", p)
	}()

	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	ch <- "paper"
	fmt.Println("manager : sent signal")

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}

// waitForResult: Think about being a manager and hiring a new employee. In
// this scenario, you want your new employee to perform a task immediately when
// they are hired, and you need to wait for the result of their work. You need
// to wait because you need the paper from them before you can continue.
func waitForResult() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- "paper"
		fmt.Println("employee : sent signal")
	}()

	p := <-ch
	fmt.Println("manager : recv'd signal :", p)

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}

// waitForFinished: Think about being a manager and hiring a new employee. In
// this scenario, you want your new employee to perform a task immediately when
// they are hired, and you need to wait for the result of their work. You need
// to wait because you can't move on until you know they are but you don't need
// anything from them.
func waitForFinished() {
	ch := make(chan struct{})

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		close(ch)
		fmt.Println("employee : sent signal")
	}()

	_, wd := <-ch
	fmt.Println("manager : recv'd signal :", wd)

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}

// pooling: Think about being a manager and hiring a team of employees. In
// this scenario, you want your new employees to perform tasks but they need
// to wait until you are ready. This is because you need to hand them a piece
// of paper before they start.
func pooling() {
	ch := make(chan string)

	const grs = 2
	for g := 0; g < grs; g++ {
		go func(id int) {
			for p := range ch {
				fmt.Printf("employee %d : recv'd signal : %s\n", id, p)
			}
			fmt.Printf("employee %d : recv'd shutdown signal\n", id)
		}(g)
	}

	const reqs = 10
	for i := 0; i < reqs; i++ {
		ch <- "paper"
		fmt.Println("manager : sent signal :", i)
	}

	close(ch)
	fmt.Println("manager : sent shutdown signal")

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}

// fanOut: Think about being a manager and hiring a team of employees. In
// this scenario, you want your new employees to perform a task immediately when
// they are hired, and you need to wait for all the results of their work.
// You need to wait because you need the paper from each of them before you
// can continue.
func fanOut() {
	grs := 20
	ch := make(chan string, grs)

	for g := 0; g < grs; g++ {
		go func(g int) {
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			ch <- "paper"
			fmt.Println("employee : sent signal :", g)
		}(g)
	}

	for grs > 0 {
		p := <-ch
		fmt.Println(p)
		fmt.Println("manager : recv'd signal :", grs)
		grs--
	}

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}

// fanOutSem: Think about being a manager and hiring a team of employees. In
// this scenario, you want only some of the new employees performing a task
// immediately when they are hired. The rest wait until other employees finish.
// You need to wait for all the results of their work. You need to wait because
// you need the paper from each of them before you can continue.
func fanOutSem() {
	grs := 20
	ch := make(chan string, grs)

	const cap = 5
	sem := make(chan bool, cap)

	for g := 0; g < grs; g++ {
		go func(g int) {
			sem <- true
			{
				time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
				ch <- "paper"
				fmt.Println("employee : sent signal :", g)
			}
			<-sem
		}(g)
	}

	for grs > 0 {
		p := <-ch
		fmt.Println(p)
		fmt.Println("manager : recv'd signal :", grs)
		grs--
	}

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}

// drop: Think about being a manager and hiring a new employee. In
// this scenario, you want your new employee to perform a task but they need
// to wait until you are ready. As the employee finishes their task you don’t
// care to know they are done. All that’s important is whether you can or can’t
// send new work. If you can’t perform the send, then you know your employee is
// at capacity. At this point the new work needs to be discarded so things can
// keep moving.
func drop() {
	const cap = 5
	ch := make(chan string, cap)

	go func() {
		for p := range ch {
			fmt.Println("employee : recv'd signal :", p)
		}
	}()

	const reqs = 20
	for r := 0; r < reqs; r++ {
		select {
		case ch <- "paper":
			fmt.Println("manager : sent signal :", r)
		default:
			fmt.Println("manager : dropped data :", r)
		}
	}

	close(ch)
	fmt.Println("manager : sent shutdown signal")

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}

// cancellation: Think about being a manager and hiring a new employee. In
// this scenario, you want your new employee to perform a task immediately when
// they are hired, and you need to wait for the result of their work. This time
// you are not willing to wait for some unknown amount of time for the employee
// to finish. You are on a discrete deadline and if the employee doesn’t finish
// in time, you are not willing to wait.
func cancellation() {
	ch := make(chan string, 1)

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- "paper"
		fmt.Println("employee : sent signal")
	}()

	tc := time.After(100 * time.Millisecond)

	select {
	case p := <-ch:
		fmt.Println("manager : recv'd signal :", p)
	case t := <-tc:
		fmt.Println("manager : timedout :", t)
	}

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}
