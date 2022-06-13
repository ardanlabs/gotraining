package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//channels transfer data between different go routines
//channels provide a means for diffent goroutines to communicate with each other

func main() {
	Loop()
	Wg()
	DoWork()
}

//loops and range in channels
func Loop() {
	//using for loops in channels
	makeChan := make(chan int)
	go func() { //this runs on a background thread
		for i := 0; i < 50; i++ {
			makeChan <- i //sending data into the channel
		}
		close(makeChan) //the go routine terminates well without a deadluck
	}()
	//receiving data from the channel
	for n := range makeChan { //this runs on the main thread
		fmt.Printf("n= %d\n", n) //prints from 0 to 49
	}
}

//working with wait grups
//wait groups wait for a collection of goroutines to finish
func Wg() {
	//using wait groups in channels
	newChan := make(chan int)
	go func() {
		wg := sync.WaitGroup{}    //using a wait group on each iteration of the loop(each go routine per iteration )
		for i := 0; i < 50; i++ { //for each iteration it creates a new goroutine
			wg.Add(1)
			//initializing another goroutine to make it faster
			go func() {
				defer wg.Done()
				result := DoWork() //a random integer of 100 from the function
				newChan <- result  //sending result into the channel
			}()
			wg.Wait() //wait until all the iterations are done
			close(newChan)
		}
	}()
}

//the function for the wg
func DoWork() int {
	time.Sleep(time.Second) //this stops the  execution of the goroutine for a second the returns the integer which will be used in the goroutine
	return rand.Intn(100)
}
