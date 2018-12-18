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
	// fanOut()
	// fanOutSem()

	// waitForFinished()
	// poolingUnlimited()
	// poolingLimited()

	// cancellation()
	// drop()
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

// poolingUnlimited: Think about being a manager and hiring a team of employees.
// In this scenario, you want your new employees to perform tasks but they need
// to wait until you are ready. This is because you need to hand them a piece
// of paper before they start.
func poolingUnlimited() {
	ch := make(chan string)

	const emps = 2
	for e := 0; e < emps; e++ {
		go func(emp int) {
			for p := range ch {
				fmt.Printf("employee %d : recv'd signal : %s\n", emp, p)
			}
			fmt.Printf("employee %d : recv'd shutdown signal\n", emp)
		}(e)
	}

	const work = 10
	for w := 0; w < work; w++ {
		ch <- "paper"
		fmt.Println("manager : sent signal :", w)
	}

	close(ch)
	fmt.Println("manager : sent shutdown signal")

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}

// poolingLimited: Think about being a manager and hiring a team of employees.
// In this scenario, you want your new employees to perform tasks but they need
// to wait until you are ready. This is because you need to hand them a piece
// of paper before they start. You know exactly all the work that needs to get
// done before it is started.
func poolingLimited() {
	work := []string{"paper", "paper", "paper", "paper", "paper"}
	ch := make(chan string, len(work))
	for _, wrk := range work {
		ch <- wrk
	}

	fmt.Println("manager : sent shutdown signal but finish all work first")
	close(ch)

	const emps = 2
	for e := 0; e < emps; e++ {
		go func(emp int) {
			for p := range ch {
				fmt.Printf("employee %d : recv'd signal : %s\n", emp, p)
			}
			fmt.Printf("employee %d : recv'd shutdown signal\n", emp)
		}(e)
	}

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}

// fanOut: Think about being a manager and hiring a team of employees. In
// this scenario, you want your new employees to perform a task immediately when
// they are hired, and you need to wait for all the results of their work.
// You need to wait because you need the paper from each of them before you
// can continue.
func fanOut() {
	emps := 20
	ch := make(chan string, emps)

	for e := 0; e < emps; e++ {
		go func(emp int) {
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			ch <- "paper"
			fmt.Println("employee : sent signal :", emp)
		}(e)
	}

	for emps > 0 {
		p := <-ch
		emps--
		fmt.Println(p)
		fmt.Println("manager : recv'd signal :", emps)
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
	emps := 20
	ch := make(chan string, emps)

	const cap = 5
	sem := make(chan bool, cap)

	for e := 0; e < emps; e++ {
		go func(emp int) {
			sem <- true
			{
				time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
				ch <- "paper"
				fmt.Println("employee : sent signal :", emp)
			}
			<-sem
		}(e)
	}

	for emps > 0 {
		p := <-ch
		emps--
		fmt.Println(p)
		fmt.Println("manager : recv'd signal :", emps)
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

	const work = 20
	for w := 0; w < work; w++ {
		select {
		case ch <- "paper":
			fmt.Println("manager : sent signal :", w)
		default:
			fmt.Println("manager : dropped data :", w)
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
